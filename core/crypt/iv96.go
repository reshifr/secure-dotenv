package crypt

import (
	"encoding/binary"
)

const (
	iv96Len int = 12
)

type IV96 struct {
	id        uint32
	increment uint64
}

func MakeIV96(id uint32) *IV96 {
	iv := &IV96{id: id, increment: 0}
	return iv
}

func LoadIV96(rawIV []byte) (*IV96, error) {
	if len(rawIV) != iv96Len {
		return nil, ErrInvalidRawIVLen
	}
	iv := &IV96{
		id:        binary.BigEndian.Uint32(rawIV[0:4]),
		increment: binary.BigEndian.Uint64(rawIV[4:iv96Len]),
	}
	return iv, nil
}

func (*IV96) Len() int {
	return iv96Len
}

func (iv *IV96) Invoke() {
	iv.increment++
}

func (iv *IV96) Raw() []byte {
	rawIV := [iv96Len]byte{}
	binary.BigEndian.PutUint32(rawIV[0:4], iv.id)
	binary.BigEndian.PutUint64(rawIV[4:iv96Len], iv.increment)
	return rawIV[:]
}
