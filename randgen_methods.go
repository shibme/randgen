package randgen

import (
	"bytes"
	"errors"
	"io"
	"os"

	"github.com/shirou/gopsutil/v4/mem"
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

func VerifyFile(file string) (string, error) {
	src, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer src.Close()
	return Verify(src)
}

func GetData(size int, secure bool) ([]byte, error) {
	virtMem, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	usable := virtMem.Available / 2
	if size > int(usable) {
		return nil, errors.New("requested size is greater than available memory")
	}
	reader, err := NewRandReader(size, secure)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(reader)
}

func VerifyData(data []byte) (string, error) {
	reader := bytes.NewReader(data)
	return Verify(reader)
}
