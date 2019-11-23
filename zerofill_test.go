package gost

import "testing"

func TestZeroBuffer(t *testing.T) {
	buf := make([]byte, 512)
	buf[0] = 1
	buf[1] = 2
	buf[2] = 3
	zeroBuffer(buf)
	if buf[0] != 0 || buf[1] != 0 || buf[2] != 0 {
		t.Fatal("did not clear buf")
	}
}
