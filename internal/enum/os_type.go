package enum

type OSType int

const (
	Windows OSType = iota
	Mac            = iota
	Android        = iota
	IOS            = iota // unused
)
