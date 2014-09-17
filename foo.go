package main

// /home/towski/save/df_linux/hack/libdfhack.so /home/towski/save/df_linux/hack/liblua.so /home/towski/save/df_linux/hack/libprotobuf-lite.so /home/towski/code/howto-go-with-cpp/libfoo.a /home/towski/save/df_linux/hack/libdfhack-client.so -lstdc++

// #cgo LDFLAGS: /home/towski/code/howto-go-with-cpp/libfoo.a /home/towski/save/df_linux/hack/libprotobuf-lite.so /home/towski/save/df_linux/hack/libdfhack-client.so -lstdc++ 
// #include "foo.h"
import "C"
import _ "fmt"
import "unicode"
import "log"
import _ "os/exec"
import "github.com/towski/artery/write"
import "github.com/towski/artery/post"

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

func (f GoFoo) GetJobType(i int) string {
    return C.GoString(C.GetJobType(C.int(i)))
}

func (f GoFoo) GetFirstName(i int) string {
    return C.GoString(C.GetFirstName(C.int(i)))
}

func (f GoFoo) GetThoughts(i int) string {
    return C.GoString(C.GetThoughts(C.int(i)))
}

func (f GoFoo) GetCurrentJob(i int) string {
    return C.GoString(C.GetCurrentJob(C.int(i)))
}

func (f GoFoo) GetHappiness(i int) string {
    return C.GoString(C.GetHappiness(C.int(i)))
}

func (f GoFoo) Size() int {
    return (int)(C.Size())
}

func main() {
    log.Println("Getting foo...")
	foo := New()
    log.Println("Getting client...")
    client := write.Client()
    i := 0
    job_html := post.JobHtml{}
    job_html.Captions = make([]string, 0)
    df_html := post.DFHtml{}
    df_html.Names = make([]string, 0)
    df_html.Jobs = make([]string, 0)
    df_html.Moods = make([]string, 0)
    for i < foo.Size() {
        dwarf_html := post.DwarfHtml{}
        dwarf_html.Thoughts = foo.GetThoughts(i)
        current_job := foo.GetCurrentJob(i)
        new_string := ""
        for _, charo := range current_job {
            if(unicode.IsUpper(charo)){
                new_string += " "
            }
            new_string += string(charo)
        }
        dwarf_html.CurrentJob = new_string
        df_html.Names = append(df_html.Names, foo.GetFirstName(i))
        df_html.Jobs = append(df_html.Jobs, new_string)
        df_html.Moods = append(df_html.Moods, foo.GetHappiness(i))
        dwarf_html.Id = i
        dwarf_html.Name = foo.GetFirstName(i)
        dwarf_html.Happiness = foo.GetHappiness(i)
        client.Call("BuildServer.BuildDwarf", dwarf_html, nil)
        i++;
    }
    log.Println("Last call...")
    client.Call("BuildServer.BuildDF", df_html, nil)
    i = 0
    for i < 2870 {
        job_html.Captions = append(job_html.Captions, foo.GetJobType(i))
        i += 1
    }
    client.Call("BuildServer.BuildJobs", job_html, nil)
	foo.Free()
    log.Println("done")
}
