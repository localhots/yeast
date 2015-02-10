package impl

func Throttler() *Impl {
	return &Impl{
		actor: func(data []byte) (resp []byte, err error) {
			return data, nil
		},
	}
}
