package iocopy

import (
	"io"
	"sync"
)

// CopyFunc is type of io copy function
type CopyFunc func(dst io.Writer, src io.Reader)

// CreatePooledIoCopyFunc return a pooled io copy function
func CreatePooledIoCopyFunc(size int) CopyFunc {
	createBuffer := func() interface{} {
		return make([]byte, 0, size)
	}

	bufferPool := sync.Pool{New: createBuffer}

	return func(dst io.Writer, src io.Reader) {
		buf := bufferPool.Get().([]byte)
		defer bufferPool.Put(buf)

		bufCap := cap(buf)
		io.CopyBuffer(dst, src, buf[0:bufCap:bufCap])
	}
}

// DefaultPolledIoCopyFunc return a pooled io copy function with default size
var DefaultPolledIoCopyFunc = CreatePooledIoCopyFunc(10 * 1024)
