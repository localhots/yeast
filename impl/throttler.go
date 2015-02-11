package impl

func Throttler(name string) *Impl {
	return &Impl{
		name: name,
		actor: func(data []byte) (resp []byte, err error) {
			return data, nil
		},
	}
}
