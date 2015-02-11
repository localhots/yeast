package impl

// Implementation of a echo unit
// Comes handy when you need to mock a unit
func Echo(name string) *Impl {
	return &Impl{
		name: name,
		actor: func(data []byte) (resp []byte, err error) {
			return data, nil
		},
	}
}
