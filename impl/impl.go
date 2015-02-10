package impl

type (
	Impl struct {
		Call func([]byte) ([]byte, error)
	}
)

func (i *Impl) Units() []string {
	return []string{}
}

func New(name string) *Impl {
	switch name {
	case "units.throttler.Throttler":
		return Throttler()
	default:
		return nil
	}
}
