package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func main() {
	arr := [10]byte{0, 1, 2, 3, 4, 5, 5, 6, 7, 8}
	size := len(arr)
	// 获取到地址,并把他转换成正整数1
	p := uintptr(unsafe.Pointer(&arr))

	var data []byte

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&data))

	sh.Data = p
	sh.Len = size
	sh.Cap = size

	fmt.Println(data)
	runtime.KeepAlive(arr)
}
