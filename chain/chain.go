package chain

import (
	"bytes"
	"sync"

	"github.com/localhots/yeast/unit"
)

type (
	Chain struct {
		Flow  Flow
		Links []unit.Caller
	}
)

const (
	LF = byte(10)
)

func New(name string) *Chain {
	c, _ := chains[name]
	return c
}

func (c *Chain) Call(data []byte) (resp []byte, err error) {
	switch c.Flow {
	case SequentialFlow:
		return c.processSequentially(data)
	case ParallelFlow:
		return c.processInParallel(data)
	case DelayedFlow:
		return c.processDelayed(data)
	default:
		panic("Unreachable")
	}
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

func (c *Chain) processSequentially(data []byte) (resp []byte, err error) {
	for _, caller := range c.Links {
		if resp, err = caller.Call(data); err != nil {
			return
		}
	}
	return
}

func (c *Chain) processInParallel(data []byte) (resp []byte, err error) {
	var (
		inbox = make(chan []byte) // This channel must be unbuffered
		buf   bytes.Buffer
		wg    sync.WaitGroup
	)

	for _, caller := range c.Links {
		wg.Add(1)
		go func() {
			if res, err := caller.Call(data); err == nil {
				inbox <- res
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(inbox)
	}()

	for {
		if res, ok := <-inbox; ok {
			buf.Write(res)
			buf.WriteByte(LF) // Add linebreak
		} else {
			break
		}
	}

	return buf.Bytes(), nil
}

func (c *Chain) processDelayed(data []byte) (resp []byte, err error) {
	for _, caller := range c.Links {
		go caller.Call(data)
	}
	return data, nil
}
