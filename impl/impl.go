package impl

type (
	Impl struct {
		name  string
		actor func([]byte) ([]byte, error)
	}
)

func New(name, impl string) *Impl {
	switch impl {
	case "units.throttler.Throttler":
		return Throttler(name)
	case "units.rule_filter.RuleFilter":
		return Echo(name)
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
