package parser

type Instruction int

// Instructions
const (
	// Stop and Arithmetic Operations
	STOP       Instruction = 0x00
	ADD        Instruction = 0x01
	MUL        Instruction = 0x02
	SUB        Instruction = 0x03
	DIV        Instruction = 0x04
	SDIV       Instruction = 0x05
	MOD        Instruction = 0x06
	SMOD       Instruction = 0x07
	ADDMOD     Instruction = 0x08
	MULMOD     Instruction = 0x09
	EXP        Instruction = 0x0a
	SIGNEXTEND Instruction = 0x0b
	// 0x0c to 0x0f invalid
	// Comparison and Bitwise Logic operations
	LT     Instruction = 0x10
	GT     Instruction = 0x11
	SLT    Instruction = 0x12
	SGT    Instruction = 0x13
	EQ     Instruction = 0x14
	ISZERO Instruction = 0x15
	AND    Instruction = 0x16
	OR     Instruction = 0x17
	XOR    Instruction = 0x18
	NOT    Instruction = 0x19
	BYTE   Instruction = 0x1a
	SHL    Instruction = 0x1b
	SHR    Instruction = 0x1c
	SAR    Instruction = 0x1d
	// 0x1e and 0x1f are invalid instructions
	// KECCAK256 hash method
	KECCAK256 Instruction = 0x20
	// 0x21 and 0x2f are invalid instructions
	// Environmental Information
	ADDRESS        Instruction = 0x30
	BALANCE        Instruction = 0x31
	ORIGIN         Instruction = 0x32
	CALLER         Instruction = 0x33
	CALLVALUE      Instruction = 0x34
	CALLDATALOAD   Instruction = 0x35
	CALLDATASIZE   Instruction = 0x36
	CALLDATACOPY   Instruction = 0x37
	CODESIZE       Instruction = 0x38
	CODECOPY       Instruction = 0x39
	GASPRICE       Instruction = 0x3a
	EXTCODESIZE    Instruction = 0x3b
	EXTCODECOPY    Instruction = 0x3c
	RETURNDATASIZE Instruction = 0x3d
	RETURNDATACOPY Instruction = 0x3e
	EXTCODEHASH    Instruction = 0x3f
	// Block Information
	BLOCKHASH   Instruction = 0x40
	COINBASE    Instruction = 0x41
	TIMESTAMP   Instruction = 0x42
	NUMBER      Instruction = 0x43
	PREVRANDAO  Instruction = 0x44
	GASLIMIT    Instruction = 0x45
	CHAINID     Instruction = 0x46
	SELFBALANCE Instruction = 0x47
	BASEFEE     Instruction = 0x48
	BLOBHASH    Instruction = 0x49
	BLOBBASE    Instruction = 0x4a
	// Stack, Memory, Storage and Flow Operations
	POP      Instruction = 0x50
	MLOAD    Instruction = 0x51
	MSTORE   Instruction = 0x52
	MSTORE8  Instruction = 0x53
	SLOAD    Instruction = 0x54
	SSTORE   Instruction = 0x55
	JUMP     Instruction = 0x56
	JUMPI    Instruction = 0x57
	PC       Instruction = 0x58
	MSIZE    Instruction = 0x59
	GAS      Instruction = 0x5a
	JUMPDEST Instruction = 0x5b
	TLOAD    Instruction = 0x5c
	TSTORE   Instruction = 0x5d
	MCOPY    Instruction = 0x5e
	// 0x5c to 0x5f are invalid instructions
	// Push operations
	PUSH0  Instruction = 0x5f
	PUSH1  Instruction = 0x60
	PUSH2  Instruction = 0x61
	PUSH3  Instruction = 0x62
	PUSH4  Instruction = 0x63
	PUSH5  Instruction = 0x64
	PUSH6  Instruction = 0x65
	PUSH7  Instruction = 0x66
	PUSH8  Instruction = 0x67
	PUSH9  Instruction = 0x68
	PUSH10 Instruction = 0x69
	PUSH11 Instruction = 0x6a
	PUSH12 Instruction = 0x6b
	PUSH13 Instruction = 0x6c
	PUSH14 Instruction = 0x6d
	PUSH15 Instruction = 0x6e
	PUSH16 Instruction = 0x6f
	PUSH17 Instruction = 0x70
	PUSH18 Instruction = 0x71
	PUSH19 Instruction = 0x72
	PUSH20 Instruction = 0x73
	PUSH21 Instruction = 0x74
	PUSH22 Instruction = 0x75
	PUSH23 Instruction = 0x76
	PUSH24 Instruction = 0x77
	PUSH25 Instruction = 0x78
	PUSH26 Instruction = 0x79
	PUSH27 Instruction = 0x7a
	PUSH28 Instruction = 0x7b
	PUSH29 Instruction = 0x7c
	PUSH30 Instruction = 0x7d
	PUSH31 Instruction = 0x7e
	PUSH32 Instruction = 0x7f
	// Duplication operations
	DUP1  Instruction = 0x80
	DUP2  Instruction = 0x81
	DUP3  Instruction = 0x82
	DUP4  Instruction = 0x83
	DUP5  Instruction = 0x84
	DUP6  Instruction = 0x85
	DUP7  Instruction = 0x86
	DUP8  Instruction = 0x87
	DUP9  Instruction = 0x88
	DUP10 Instruction = 0x89
	DUP11 Instruction = 0x8a
	DUP12 Instruction = 0x8b
	DUP13 Instruction = 0x8c
	DUP14 Instruction = 0x8d
	DUP15 Instruction = 0x8e
	DUP16 Instruction = 0x8f
	// Exchange operations
	SWAP1  Instruction = 0x90
	SWAP2  Instruction = 0x91
	SWAP3  Instruction = 0x92
	SWAP4  Instruction = 0x93
	SWAP5  Instruction = 0x94
	SWAP6  Instruction = 0x95
	SWAP7  Instruction = 0x96
	SWAP8  Instruction = 0x97
	SWAP9  Instruction = 0x98
	SWAP10 Instruction = 0x99
	SWAP11 Instruction = 0x9a
	SWAP12 Instruction = 0x9b
	SWAP13 Instruction = 0x9c
	SWAP14 Instruction = 0x9d
	SWAP15 Instruction = 0x9e
	SWAP16 Instruction = 0x9f

	// Logging Operations
	LOG0 Instruction = 0xa0
	LOG1 Instruction = 0xa1
	LOG2 Instruction = 0xa2
	LOG3 Instruction = 0xa3
	LOG4 Instruction = 0xa4

	// System operations
	CREATE       Instruction = 0xf0
	CALL         Instruction = 0xf1
	CALLCODE     Instruction = 0xf2
	RETURN       Instruction = 0xf3
	DELEGATECALL Instruction = 0xf4
	CREATE2      Instruction = 0xf5
	// 0xf6 to 0xf9 invalid
	STATICCALL Instruction = 0xfa
	// 0xfb and 0xfc
	REVERT       Instruction = 0xfd
	INVALID      Instruction = 0xfe
	SELFDESTRUCT Instruction = 0xff
)

type InstructionOp struct {
	Mnemonic   string
	SizeToRead int
}

var (
	// Instructions map
	InstructionsMap = map[Instruction]InstructionOp{
		STOP:           {Mnemonic: "STOP", SizeToRead: 0},
		ADD:            {Mnemonic: "ADD", SizeToRead: 0},
		MUL:            {Mnemonic: "MUL", SizeToRead: 0},
		SUB:            {Mnemonic: "SUB", SizeToRead: 0},
		DIV:            {Mnemonic: "DIV", SizeToRead: 0},
		SDIV:           {Mnemonic: "SDIV", SizeToRead: 0},
		MOD:            {Mnemonic: "MOD", SizeToRead: 0},
		SMOD:           {Mnemonic: "SMOD", SizeToRead: 0},
		ADDMOD:         {Mnemonic: "ADDMOD", SizeToRead: 0},
		MULMOD:         {Mnemonic: "MULMOD", SizeToRead: 0},
		EXP:            {Mnemonic: "EXP", SizeToRead: 0},
		SIGNEXTEND:     {Mnemonic: "SIGNEXTEND", SizeToRead: 0},
		LT:             {Mnemonic: "LT", SizeToRead: 0},
		GT:             {Mnemonic: "GT", SizeToRead: 0},
		SLT:            {Mnemonic: "SLT", SizeToRead: 0},
		SGT:            {Mnemonic: "SGT", SizeToRead: 0},
		EQ:             {Mnemonic: "EQ", SizeToRead: 0},
		ISZERO:         {Mnemonic: "ISZERO", SizeToRead: 0},
		AND:            {Mnemonic: "AND", SizeToRead: 0},
		OR:             {Mnemonic: "OR", SizeToRead: 0},
		XOR:            {Mnemonic: "XOR", SizeToRead: 0},
		NOT:            {Mnemonic: "NOT", SizeToRead: 0},
		BYTE:           {Mnemonic: "BYTE", SizeToRead: 0},
		SHL:            {Mnemonic: "SHL", SizeToRead: 0},
		SHR:            {Mnemonic: "SHR", SizeToRead: 0},
		SAR:            {Mnemonic: "SAR", SizeToRead: 0},
		KECCAK256:      {Mnemonic: "KECCAK256", SizeToRead: 0},
		ADDRESS:        {Mnemonic: "ADDRESS", SizeToRead: 0},
		BALANCE:        {Mnemonic: "BALANCE", SizeToRead: 0},
		ORIGIN:         {Mnemonic: "ORIGIN", SizeToRead: 0},
		CALLER:         {Mnemonic: "CALLER", SizeToRead: 0},
		CALLVALUE:      {Mnemonic: "CALLVALUE", SizeToRead: 0},
		CALLDATALOAD:   {Mnemonic: "CALLDATALOAD", SizeToRead: 0},
		CALLDATASIZE:   {Mnemonic: "CALLDATASIZE", SizeToRead: 0},
		CALLDATACOPY:   {Mnemonic: "CALLDATACOPY", SizeToRead: 0},
		CODESIZE:       {Mnemonic: "CODESIZE", SizeToRead: 0},
		CODECOPY:       {Mnemonic: "CODECOPY", SizeToRead: 0},
		GASPRICE:       {Mnemonic: "GASPRICE", SizeToRead: 0},
		EXTCODESIZE:    {Mnemonic: "EXTCODESIZE", SizeToRead: 0},
		EXTCODECOPY:    {Mnemonic: "EXTCODECOPY", SizeToRead: 0},
		RETURNDATASIZE: {Mnemonic: "RETURNDATASIZE", SizeToRead: 0},
		RETURNDATACOPY: {Mnemonic: "RETURNDATACOPY", SizeToRead: 0},
		EXTCODEHASH:    {Mnemonic: "EXTCODEHASH", SizeToRead: 0},
		BLOCKHASH:      {Mnemonic: "BLOCKHASH", SizeToRead: 0},
		COINBASE:       {Mnemonic: "COINBASE", SizeToRead: 0},
		TIMESTAMP:      {Mnemonic: "TIMESTAMP", SizeToRead: 0},
		NUMBER:         {Mnemonic: "NUMBER", SizeToRead: 0},
		PREVRANDAO:     {Mnemonic: "PREVRANDAO", SizeToRead: 0},
		GASLIMIT:       {Mnemonic: "GASLIMIT", SizeToRead: 0},
		CHAINID:        {Mnemonic: "CHAINID", SizeToRead: 0},
		SELFBALANCE:    {Mnemonic: "SELFBALANCE", SizeToRead: 0},
		BASEFEE:        {Mnemonic: "BASEFEE", SizeToRead: 0},
		BLOBHASH:       {Mnemonic: "BLOBHASH", SizeToRead: 0},
		BLOBBASE:       {Mnemonic: "BLOBBASE", SizeToRead: 0},
		POP:            {Mnemonic: "POP", SizeToRead: 0},
		MLOAD:          {Mnemonic: "MLOAD", SizeToRead: 0},
		MSTORE:         {Mnemonic: "MSTORE", SizeToRead: 0},
		MSTORE8:        {Mnemonic: "MSTORE8", SizeToRead: 0},
		SLOAD:          {Mnemonic: "SLOAD", SizeToRead: 0},
		SSTORE:         {Mnemonic: "SSTORE", SizeToRead: 0},
		JUMP:           {Mnemonic: "JUMP", SizeToRead: 0},
		JUMPI:          {Mnemonic: "JUMPI", SizeToRead: 0},
		PC:             {Mnemonic: "PC", SizeToRead: 0},
		MSIZE:          {Mnemonic: "MSIZE", SizeToRead: 0},
		GAS:            {Mnemonic: "GAS", SizeToRead: 0},
		JUMPDEST:       {Mnemonic: "JUMPDEST", SizeToRead: 0},
		TLOAD:          {Mnemonic: "TLOAD", SizeToRead: 0},
		TSTORE:         {Mnemonic: "TSTORE", SizeToRead: 0},
		MCOPY:          {Mnemonic: "MCOPY", SizeToRead: 0},
		PUSH0:          {Mnemonic: "PUSH0", SizeToRead: 0},
		PUSH1:          {Mnemonic: "PUSH1", SizeToRead: 1},
		PUSH2:          {Mnemonic: "PUSH2", SizeToRead: 2},
		PUSH3:          {Mnemonic: "PUSH3", SizeToRead: 3},
		PUSH4:          {Mnemonic: "PUSH4", SizeToRead: 4},
		PUSH5:          {Mnemonic: "PUSH5", SizeToRead: 5},
		PUSH6:          {Mnemonic: "PUSH6", SizeToRead: 6},
		PUSH7:          {Mnemonic: "PUSH7", SizeToRead: 7},
		PUSH8:          {Mnemonic: "PUSH8", SizeToRead: 8},
		PUSH9:          {Mnemonic: "PUSH9", SizeToRead: 9},
		PUSH10:         {Mnemonic: "PUSH10", SizeToRead: 10},
		PUSH11:         {Mnemonic: "PUSH11", SizeToRead: 11},
		PUSH12:         {Mnemonic: "PUSH12", SizeToRead: 12},
		PUSH13:         {Mnemonic: "PUSH13", SizeToRead: 13},
		PUSH14:         {Mnemonic: "PUSH14", SizeToRead: 14},
		PUSH15:         {Mnemonic: "PUSH15", SizeToRead: 15},
		PUSH16:         {Mnemonic: "PUSH16", SizeToRead: 16},
		PUSH17:         {Mnemonic: "PUSH17", SizeToRead: 17},
		PUSH18:         {Mnemonic: "PUSH18", SizeToRead: 18},
		PUSH19:         {Mnemonic: "PUSH19", SizeToRead: 19},
		PUSH20:         {Mnemonic: "PUSH20", SizeToRead: 20},
		PUSH21:         {Mnemonic: "PUSH21", SizeToRead: 21},
		PUSH22:         {Mnemonic: "PUSH22", SizeToRead: 22},
		PUSH23:         {Mnemonic: "PUSH23", SizeToRead: 23},
		PUSH24:         {Mnemonic: "PUSH24", SizeToRead: 24},
		PUSH25:         {Mnemonic: "PUSH25", SizeToRead: 25},
		PUSH26:         {Mnemonic: "PUSH26", SizeToRead: 26},
		PUSH27:         {Mnemonic: "PUSH27", SizeToRead: 27},
		PUSH28:         {Mnemonic: "PUSH28", SizeToRead: 28},
		PUSH29:         {Mnemonic: "PUSH29", SizeToRead: 29},
		PUSH30:         {Mnemonic: "PUSH30", SizeToRead: 30},
		PUSH31:         {Mnemonic: "PUSH31", SizeToRead: 31},
		PUSH32:         {Mnemonic: "PUSH32", SizeToRead: 32},
		DUP1:           {Mnemonic: "DUP1", SizeToRead: 0},
		DUP2:           {Mnemonic: "DUP2", SizeToRead: 0},
		DUP3:           {Mnemonic: "DUP3", SizeToRead: 0},
		DUP4:           {Mnemonic: "DUP4", SizeToRead: 0},
		DUP5:           {Mnemonic: "DUP5", SizeToRead: 0},
		DUP6:           {Mnemonic: "DUP6", SizeToRead: 0},
		DUP7:           {Mnemonic: "DUP7", SizeToRead: 0},
		DUP8:           {Mnemonic: "DUP8", SizeToRead: 0},
		DUP9:           {Mnemonic: "DUP9", SizeToRead: 0},
		DUP10:          {Mnemonic: "DUP10", SizeToRead: 0},
		DUP11:          {Mnemonic: "DUP11", SizeToRead: 0},
		DUP12:          {Mnemonic: "DUP12", SizeToRead: 0},
		DUP13:          {Mnemonic: "DUP13", SizeToRead: 0},
		DUP14:          {Mnemonic: "DUP14", SizeToRead: 0},
		DUP15:          {Mnemonic: "DUP15", SizeToRead: 0},
		DUP16:          {Mnemonic: "DUP16", SizeToRead: 0},
		SWAP1:          {Mnemonic: "SWAP1", SizeToRead: 0},
		SWAP2:          {Mnemonic: "SWAP2", SizeToRead: 0},
		SWAP3:          {Mnemonic: "SWAP3", SizeToRead: 0},
		SWAP4:          {Mnemonic: "SWAP4", SizeToRead: 0},
		SWAP5:          {Mnemonic: "SWAP5", SizeToRead: 0},
		SWAP6:          {Mnemonic: "SWAP6", SizeToRead: 0},
		SWAP7:          {Mnemonic: "SWAP7", SizeToRead: 0},
		SWAP8:          {Mnemonic: "SWAP8", SizeToRead: 0},
		SWAP9:          {Mnemonic: "SWAP9", SizeToRead: 0},
		SWAP10:         {Mnemonic: "SWAP10", SizeToRead: 0},
		SWAP11:         {Mnemonic: "SWAP11", SizeToRead: 0},
		SWAP12:         {Mnemonic: "SWAP12", SizeToRead: 0},
		SWAP13:         {Mnemonic: "SWAP13", SizeToRead: 0},
		SWAP14:         {Mnemonic: "SWAP14", SizeToRead: 0},
		SWAP15:         {Mnemonic: "SWAP15", SizeToRead: 0},
		SWAP16:         {Mnemonic: "SWAP16", SizeToRead: 0},
		LOG0:           {Mnemonic: "LOG0", SizeToRead: 0},
		LOG1:           {Mnemonic: "LOG1", SizeToRead: 0},
		LOG2:           {Mnemonic: "LOG2", SizeToRead: 0},
		LOG3:           {Mnemonic: "LOG3", SizeToRead: 0},
		LOG4:           {Mnemonic: "LOG4", SizeToRead: 0},
		CREATE:         {Mnemonic: "CREATE", SizeToRead: 0},
		CALL:           {Mnemonic: "CALL", SizeToRead: 0},
		CALLCODE:       {Mnemonic: "CALLCODE", SizeToRead: 0},
		RETURN:         {Mnemonic: "RETURN", SizeToRead: 0},
		DELEGATECALL:   {Mnemonic: "DELEGATECALL", SizeToRead: 0},
		CREATE2:        {Mnemonic: "CREATE2", SizeToRead: 0},
		STATICCALL:     {Mnemonic: "STATICCALL", SizeToRead: 0},
		REVERT:         {Mnemonic: "REVERT", SizeToRead: 0},
		INVALID:        {Mnemonic: "INVALID", SizeToRead: 0},
		SELFDESTRUCT:   {Mnemonic: "SELFDESTRUCT", SizeToRead: 0},
	}
)
