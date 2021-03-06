package arch8

import (
	"bytes"
	"os"
)

// RunImageFile loads and run a raw image on a single core machine
// with 1GB physical memory until it runs into an exception.
func RunImageFile(path string) error {
	f, e := os.Open(path)
	if e != nil {
		return e
	}

	m := NewMachine(uint32(1<<30), 1)
	if e := m.LoadImage(f, InitPC); e != nil {
		return e
	}
	if e := f.Close(); e != nil {
		return e
	}

	_, exp := m.Run(0)
	if exp != nil {
		return exp
	}

	return nil
}

// RunImage runs a series of bytes as a VM image
// with 1GB physical memory for maximum n cycles.
// It returns the number of cycles, and the exit error
// if any.
func RunImage(bs []byte, n int) (int, error) {
	r := bytes.NewBuffer(bs)
	m := NewMachine(uint32(1<<30), 1)
	if e := m.LoadImage(r, InitPC); e != nil {
		return 0, e
	}

	ret, exp := m.Run(n)
	if exp == nil {
		return ret, nil
	}
	return ret, exp
}

// RunImageOutput runs a image. It is similar to
// RunImage() but also returns the output.
func RunImageOutput(bs []byte, n int) (int, string, error) {
	r := bytes.NewBuffer(bs)
	m := NewMachine(uint32(1<<30), 1)
	if e := m.LoadImage(r, InitPC); e != nil {
		return 0, "", e
	}

	out := new(bytes.Buffer)
	m.SetOutput(out)

	ret, exp := m.Run(n)
	if exp == nil {
		return ret, out.String(), nil
	}
	return ret, out.String(), exp
}

// IsHalt returns true only when the error is a halt exception
func IsHalt(e error) bool {
	if e, ok := e.(*Excep); ok {
		return e.Code == ErrHalt
	}
	if e, ok := e.(*CoreExcep); ok {
		return e.Excep.Code == ErrHalt
	}

	return false
}
