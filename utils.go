package burn

import (
	"unsafe"
)

func GetString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
