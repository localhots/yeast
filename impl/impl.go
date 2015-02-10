package impl

type (
	Impl struct {
		actor func([]byte) ([]byte, error)
	}
)

func New(name string) *Impl {
	switch name {
	case "units.throttler.Throttler":
		return Throttler()
	default:
		return nil
	}
}

func (i *Impl) Call(data []byte) (resp []byte, err error) {
	return i.actor(data)
}

func (i *Impl) Units() []string {
	return []string{}
}
