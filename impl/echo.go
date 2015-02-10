package impl

// Implementation of a echo unit
// Comes handy when you need to mock a unit
func Echo() *Impl {
	return &Impl{
		Call: func(data []byte) (resp []byte, err error) {
			return data, nil
		},
	}
}
