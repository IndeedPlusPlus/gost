package gost

func zeroBuffer(buf []byte) {
	// This loop will be optimized to use memclr
	for i := range buf {
		buf[i] = 0
	}
}
