package event

type PubsubType int32

const (
	PubsubTypeDefault PubsubType = 0
)

func (t PubsubType) ToString() string {
	switch t {
	case PubsubTypeDefault:
		return ""
	}
	return ""
}
