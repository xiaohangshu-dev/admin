package gender

type Gender int8

const (
	Unknown Gender = iota + 1 // 保密
	Male                      // 男
	Female                    // 女
)
