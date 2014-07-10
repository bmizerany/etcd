package wal

import (
	"encoding/binary"
	"errors"
	"io"
)

var ErrInvalidChecksum = errors.New("wal: invalid checksum")

type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r}
}

func (d *Decoder) Decode(rec *Record) error {
	// TODO(bmizerany): reuse cached buf?
	b := make([]byte, 8)
	var l int64
	if err := binary.Read(d.r, binary.BigEndian, &l); err != nil {
		return err
	}

	// TODO(bmizerany): reuse cached buf?
	b = make([]byte, l)
	if _, err := io.ReadFull(d.r, b); err != nil {
		return err
	}
	if err := rec.Unmarshal(b); err != nil {
		return err
	}
	if !rec.isValid() {
		return ErrInvalidChecksum
	}
	return nil
}
