package rw

import "encoding/binary"

type (
	state = uint16
)

const (
	unlocked state = 0
	locked   state = 1
)

var (
	Locked   []byte = make([]byte, 2)
	Unlocked []byte = make([]byte, 2)
)

func init() {
	binary.LittleEndian.PutUint16(Locked, locked)
	binary.LittleEndian.PutUint16(Unlocked, unlocked)
}
