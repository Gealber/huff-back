package parser

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

// Parse given a filename, parse bytecode in given file
// displaying
func Parse(filename, outputFile string) error {
	b, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}

	// strip 0x prefix
	b, _ = bytes.CutPrefix(b, []byte("0x"))
	b, _ = bytes.CutSuffix(b, []byte("\n"))
	byteArr, err := transformByteArr(b)
	if err != nil {
		return err
	}

	// reading bytes
	size := len(byteArr)
	counter := 0
	for counter < size {
		// reading instruction byte
		op, ok := InstructionsMap[Instruction(byteArr[counter])]
		if !ok {
			f.Write([]byte("INVALID_"))
			counter++
			continue
		}

		// for PUSH instructions we need to read following bytes
		if op.SizeToRead > 0 {
			if counter+1+op.SizeToRead > size {
				return fmt.Errorf("insufficient bytes to read for push instruction: %s", op.Mnemonic)
			}

			next := counter + 1 + op.SizeToRead

			if next < size {
				if nexOp, ok := InstructionsMap[Instruction(byteArr[next])]; ok && nexOp.Mnemonic == "JUMPI" {
					n, err := strconv.ParseInt(hex.EncodeToString(byteArr[counter+1:counter+1+op.SizeToRead]), 16, 64)
					if err != nil {
						return err
					}

					line := fmt.Sprintf("JUMP_%d\n", n)
					f.Write([]byte(line))
					counter += 1 + op.SizeToRead + 1
					continue
				}
			}

			// read subsequents bytes to push as number
			// ignoring current position of counter to not print "0X01 PUSH1" for example
			n := "0x" + hex.EncodeToString(byteArr[counter+1:counter+1+op.SizeToRead]) + "\n"
			f.Write([]byte(n))
			counter += 1 + op.SizeToRead
		} else {
			if op.Mnemonic == "PUSH0" {
				n := "0x" + hex.EncodeToString([]byte{0}) + "\n"
				f.Write([]byte(n))
			} else if op.Mnemonic == "JUMPDEST" {
				f.Write([]byte("JUMP" + "_" + strconv.Itoa(counter) + ":" + "\n"))
			} else {
				f.Write([]byte(op.Mnemonic + "\n"))
			}
			counter++
		}
	}

	return nil
}

// transformByteArr transform the readed byte array
// interpreting each pair of bytes as a one byte hex number
// for example: [54 48] => 0x60, taking that 54 => '6' and 48 => '0'
func transformByteArr(arr []byte) ([]byte, error) {
	result := make([]byte, 0)
	if len(arr)%2 != 0 {
		return nil, fmt.Errorf("byte amount is not even: %d", len(arr))
	}

	for i := 0; i < len(arr)-1; i += 2 {
		r, err := hex.DecodeString(string(arr[i : i+2]))
		if err != nil {
			return nil, err
		}
		result = append(result, r...)
	}

	return result, nil
}
