package gost

import "testing"

const (
	bufSize = 16384
)

func TestBufferPoolZeroOut(t *testing.T) {
	p := newBufferPool(bufSize)
	buf := p.Get().([]byte)
	buf[0] = 1
	p.Put(buf)
	buf = p.Get().([]byte)
	if buf[0] != 0 {
		t.Errorf("buffer not filled with zero, got %v", buf[0])
	}
}
