package randgen

import (
	"bytes"
	cryptorand "crypto/rand"
	"errors"
	"hash"
	"hash/crc32"
	"io"
	mathrand "math/rand"
	"strconv"
	"time"
)

const (
	blockSize = 1024
	hashSize  = crc32.Size
)

type RandReader struct {
	randomizer io.Reader
	block      []byte
	buf        bytes.Buffer
	limit      int
	hasher     hash.Hash
}

func NewRandReader(limit int, secureRandom bool) (*RandReader, error) {
	if limit <= hashSize {
		return nil, errors.New("size must be greater than " + strconv.Itoa(hashSize))
	}
	limit -= hashSize
	randReader := &RandReader{
		block:  make([]byte, blockSize),
		buf:    bytes.Buffer{},
		limit:  limit,
		hasher: crc32.NewIEEE(),
	}
	if secureRandom {
		randReader.randomizer = cryptorand.Reader
	} else {
		randReader.randomizer = mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	}
	return randReader, nil
}

func (r *RandReader) fillBuf(minSize int) error {
	for r.limit > 0 && r.buf.Len() < minSize {
		if r.limit < len(r.block) {
			r.block = r.block[:r.limit]
		}
		if _, err := io.ReadFull(r.randomizer, r.block); err != nil {
			return err
		}
		r.buf.Write(r.block)
		r.hasher.Write(r.block)
		r.limit -= len(r.block)
		if r.limit == 0 {
			digest := r.hasher.Sum(nil)
			r.buf.Write(digest)
		}
	}
	return nil
}

func (r *RandReader) Read(p []byte) (int, error) {
	if r.buf.Len() < len(p) {
		if err := r.fillBuf(len(p)); err != nil {
			return 0, err
		}
	}
	n, err := r.buf.Read(p)
	return n, err
}

func Verify(r io.Reader) error {
	hasher := crc32.NewIEEE()
	block := make([]byte, blockSize)
	var residue []byte
	var err error
	var n int
	for n, err = r.Read(block); err == nil; n, err = r.Read(block) {
		residue = append(residue, block[:n]...)
		hashableLen := len(residue) - hashSize
		hasher.Write(residue[:hashableLen])
		residue = residue[hashableLen:]
	}
	if err != io.EOF {
		return err
	}
	if !bytes.Equal(hasher.Sum(nil), residue) {
		return errVerificationFailed
	}
	return nil
}
