package common

import (
	"io"
	"strconv"
	"sync"
)

var bufpool *sync.Pool

func InitBufPool() {
	bufpool = &sync.Pool{}
	bufpool.New = func() interface{} {
		return make([]byte, 64*1024)
	}
}

func Copy(dst io.Writer, src io.Reader) (written int64, err error) {
	if wt, ok := src.(io.WriterTo); ok {
		return wt.WriteTo(dst)
	}
	if rt, ok := dst.(io.ReaderFrom); ok {
		return rt.ReadFrom(src)
	}

	buf := bufpool.Get().([]byte)
	defer bufpool.Put(buf)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er == io.EOF {
			break
		}
		if er != nil {
			err = er
			break
		}
	}
	return written, err
}

// Get Addr
func GetAddr(ip string, port uint) string {
	return ip + ":" + strconv.Itoa(int(port))
}
