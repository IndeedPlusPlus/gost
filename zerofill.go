package gost

import "runtime"

func zeroBuffer(buf []byte) {
	// This loop will be optimized to use memclr
	for i := range buf {
		buf[i] = 0
	}
}

func setZeroingFinalizer(buf []byte) {
	runtime.SetFinalizer(&buf, func(buf *[]byte) {
		zeroBuffer(*buf)
	})
}
