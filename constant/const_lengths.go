package constant

// 常用的常量长度值
const (
	LEN_0  = 0               // iota = 0, 用于跳过 LEN_0
	LEN_1  = 1 << (iota - 1) // 1 << 0 = 1
	LEN_2                    // 1 << 1 = 2
	LEN_4                    // 1 << 2 = 4
	LEN_8                    // 1 << 3 = 8
	LEN_16                   // ...
	LEN_32
	LEN_64
	LEN_128
	LEN_256
	LEN_512
	LEN_1024
	LEN_2048
	LEN_4096
	LEN_8192
	LEN_16384
	LEN_32768
	LEN_65536
)
