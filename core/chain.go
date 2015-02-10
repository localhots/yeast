package core

type (
	Chain struct {
		Flow  Flow
		Links []Caller
	}
	Caller interface {
		Call([]byte) ([]byte, error)
		Units() []string
	}
)

func (c *Chain) Call(data []byte) (resp []byte, err error) {
	return data, nil
}

func (c *Chain) Units() []string {
	// Collecting unique unit names using map
	units := map[string]*struct{}{}
	for _, caller := range c.Links {
		for _, unit := range caller.Units() {
			units[unit] = nil
		}
	}

	// Extracting names to a slice
	uniq := []string{}
	for unit, _ := range units {
		uniq = append(uniq, unit)
	}

	return uniq
}
