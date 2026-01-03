package status

type Status int8

const (
	Active  Status = iota + 1 // 正常
	Disable                   // 禁用
	Locked                    // 上锁
)
