package main

// /home/towski/save/df_linux/hack/libdfhack.so /home/towski/save/df_linux/hack/liblua.so /home/towski/save/df_linux/hack/libprotobuf-lite.so /home/towski/code/howto-go-with-cpp/libfoo.a /home/towski/save/df_linux/hack/libdfhack-client.so -lstdc++

// #cgo LDFLAGS: /home/towski/code/howto-go-with-cpp/libfoo.a /home/towski/save/df_linux/hack/libprotobuf-lite.so /home/towski/save/df_linux/hack/libdfhack-client.so -lstdc++
// #include "foo.h"
import "C"
import "fmt"

// #cgo CFLAGS: -Ilibrary/include/df -Ilibrary/include/ -Ilibrary/proto -Idepends/protobuf/google/protobuf/ -Idepends/protobuf -std=c++11

type Unit struct {
    FirstName string
}

var units []Unit

type GoFoo struct {
	foo C.Foo
}

func New() GoFoo {
	var ret GoFoo
	ret.foo = C.FooInit()
	return ret
}
func (f GoFoo) Free() {
	C.FooFree(f.foo)
}
func (f GoFoo) Bar(i int) string {
    return C.GoString(C.FooBar(C.int(i)))
}

func main() {
	foo := New()
	fmt.Println(foo.Bar(1))
	foo.Free()
}
