package randgen

import (
	"bytes"
	cryptorand "crypto/rand"
	"io"
	mathrand "math/rand"
	"os"
	"time"
)

const (
	blockSize = 1024 // 1 kB
)

type RandReader struct {
	reader io.Reader
	block  []byte
	buf    bytes.Buffer
	limit  uint64
}

func NewRandReader(limit uint64, secureRandom bool) *RandReader {
	randReader := &RandReader{
		block: make([]byte, blockSize),
		buf:   bytes.Buffer{},
		limit: limit,
	}
	if secureRandom {
		randReader.reader = cryptorand.Reader
	} else {
		randReader.reader = mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	}
	return randReader
}

func (r *RandReader) Read(p []byte) (int, error) {
	if r.limit == 0 {
		return 0, io.EOF
	}
	if r.limit < uint64(len(p)) {
		p = p[:r.limit]
	}
	if r.buf.Len() >= len(p) {
		r.limit -= uint64(len(p))
		return r.buf.Read(p)
	} else {
		for r.buf.Len() < len(p) {
			if _, err := io.ReadFull(r.reader, r.block); err == nil {
				r.buf.Write(r.block)
			} else if err == io.EOF {
				return r.buf.Read(p)
			} else {
				return 0, err
			}
		}
		return r.Read(p)
	}
}

func WriteBlob(w io.Writer, size uint64, secureRandom bool) error {
	reader := NewRandReader(size, secureRandom)
	if _, err := io.CopyN(w, reader, int64(size)); err != nil {
		return err
	}
	return nil
}

func CreateFile(fileName string, size uint64, secureRandom bool) error {
	dst, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer dst.Close()
	return WriteBlob(dst, size, secureRandom)
}
