package cat

type enemy struct {
	enType enType
}

type enType int

const (
	snake enType = iota
)
