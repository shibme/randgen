package randgen

import (
	"errors"
	"io"
	"os"
)

var (
	errVerificationFailed = errors.New("the data integrity check failed")
)

func WriteRand(w io.Writer, size int, secure bool) error {
	reader, err := NewRandReader(size, secure)
	if err != nil {
		return err
	}
	if _, err = io.Copy(w, reader); err != nil {
		return err
	}
	return nil
}

func CreateFile(file string, size int, secure bool) error {
	dst, err := os.Create(file)
	if err != nil {
		return err
	}
	defer dst.Close()
	return WriteRand(dst, size, secure)
}

func VerifyFile(file string) error {
	src, err := os.Open(file)
	if err != nil {
		return err
	}
	defer src.Close()
	return Verify(src)
}
