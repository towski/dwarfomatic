package main

// /home/towski/save/df_linux/hack/libdfhack.so /home/towski/save/df_linux/hack/liblua.so /home/towski/save/df_linux/hack/libprotobuf-lite.so /home/towski/code/howto-go-with-cpp/libfoo.a /home/towski/save/df_linux/hack/libdfhack-client.so -lstdc++

// #cgo LDFLAGS: /home/towski/code/dwarfomatic/libfoo.a /home/towski/save/df_linux/hack/libprotobuf-lite.so /home/towski/save/df_linux/hack/libdfhack-client.so -lstdc++ 
// #include "foo.h"
import "C"
import _ "fmt"
import "unicode"
import _ "log"
import "os"
import "time"
import "net/rpc"
import "os/exec"
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

func (f GoFoo) Update() {
    C.Update()
}

func (f GoFoo) Exit() {
    C.Update()
}

func (f GoFoo) Size() int {
    return (int)(C.Size())
}

func main(){
	foo := New()
    client := write.Client()
    foo.Update()
    ProcessManagerOrders(client, foo)
    ProcessData(client, foo)
    stonesense_window := os.Args[1]
    dwarf_fortress_window := os.Args[2]
    go func(){
        var cmd *exec.Cmd
        for {
            time.Sleep(2000 * time.Millisecond)
            cmd = exec.Command("./capture_screenshot.sh", stonesense_window, "stonesense.png")
            cmd.Run()
            cmd = exec.Command("./capture_screenshot.sh", dwarf_fortress_window, "screenshot.png")
            cmd.Run()
        }
    }()
    go func(){
        var cmd *exec.Cmd
        for {
            time.Sleep(10000 * time.Millisecond)
            cmd = exec.Command("/home/towski/save/df_linux/dfhack-run", "show_orders")
            cmd.Run()
        }
    }()
    for {
        time.Sleep(500 * time.Millisecond)
        foo.Update()
        ProcessData(client, foo)
    }
	foo.Free()
}

func ProcessManagerOrders(client *rpc.Client, foo GoFoo){
    job_html := post.JobHtml{}
    job_html.Captions = make([]string, 0)
    i := 0
    for i < 2858 {
        job_html.Captions = append(job_html.Captions, foo.GetJobType(i))
        i += 1
    }
    client.Call("BuildServer.BuildJobs", job_html, nil)
}

func ProcessData(client *rpc.Client, foo GoFoo) {
    i := 0
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
    client.Call("BuildServer.BuildDF", df_html, nil)
    i = 0
    cmd := exec.Command("update_game_log.sh")
    _ = cmd.Run()
}
