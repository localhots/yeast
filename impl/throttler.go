package impl

func Throttler() *Impl {
	return &Impl{
		Call: func(data []byte) (resp []byte, err error) {
			return data, nil
		},
	}
}
