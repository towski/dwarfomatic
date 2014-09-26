package main

// /home/towski/save/df_linux/hack/libdfhack.so /home/towski/save/df_linux/hack/liblua.so /home/towski/save/df_linux/hack/libprotobuf-lite.so /home/towski/code/howto-go-with-cpp/libfoo.a /home/towski/save/df_linux/hack/libdfhack-client.so -lstdc++

// #cgo LDFLAGS: /home/towski/code/dwarfomatic/libfoo.a /home/towski/save/df_linux/hack/libprotobuf-lite.so /home/towski/save/df_linux/hack/libdfhack-client.so -lstdc++ 
// #include "foo.h"
import "C"
import _ "fmt"
import "unicode"
import "log"
import "sync"
import "os/signal"
import "os"
import "strconv"
import "time"
import "net/rpc"
import "os/exec"
import "github.com/towski/artery/write"
import server "github.com/towski/artery/server"
import dwarfomatic "github.com/towski/dwarfomatic/classes"

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

var mutex sync.Mutex

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

func (f GoFoo) GetId(i int) int {
    return (int)(C.GetId(C.int(i)))
}

func (f GoFoo) Update() {
    mutex.Lock()
    C.Update()
    mutex.Unlock()
}

func (f GoFoo) Exit() {
    C.Exit()
}

func (f GoFoo) Size() int {
    return (int)(C.Size())
}

func (f GoFoo) Paused() int {
    return (int)(C.Paused())
}

var running bool = true
var data_client *rpc.Client
var foo GoFoo

func main(){
    var wg sync.WaitGroup
	foo = New()
    client := write.Client()
    data_client = server.DataClient()
    foo.Update()
    ProcessManagerOrders(client)
    ProcessData(client)
    stonesense_window := os.Args[1]
    dwarf_fortress_window := os.Args[2]
    wg.Add(1)
    go func(){
        var cmd *exec.Cmd
        for running == true{
            time.Sleep(2000 * time.Millisecond)
            cmd = exec.Command("./capture_screenshot.sh", stonesense_window, "stonesense.png")
            cmd.Run()
            cmd = exec.Command("./capture_screenshot.sh", dwarf_fortress_window, "screenshot.png")
            cmd.Run()
        }
        wg.Done()
    }()
    go func(){
        wg.Add(1)
        var cmd *exec.Cmd
        for running == true{
            time.Sleep(10000 * time.Millisecond)
            cmd = exec.Command("/home/towski/save/df_linux/dfhack-run", "show_orders")
            _ = cmd
            //cmd.Run()
        }
        wg.Done()
    }()
    go func(){
        wg.Add(1)
        for running == true{
            time.Sleep(500 * time.Millisecond)
            foo.Update()
            ProcessData(client)
        }
        wg.Done()
    }()
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
        wg.Add(1)
        for sig := range c {
            log.Printf("Cancelling", sig)
            foo.Free()
            running = false
            break
        }
        wg.Done()
    }()
    go func(){
        wg.Add(1)
        for running == true{
            time.Sleep(10000 * time.Millisecond)
            i := 0
            mutex.Lock()
            for i < foo.Size() {
                ProcessAvatars(i)
                i += 1
            }
            mutex.Unlock()
        }
        wg.Done()
    }()
    go func(){
        wg.Add(1)
        for running == true{
            i := 0
            mutex.Lock()
            for i < foo.Size() {
                data_client.Call("DwarfServer.Insert", GetDwarf(i), nil)
                i += 1
            }
            mutex.Unlock()
            time.Sleep(100000 * time.Millisecond)
        }
        wg.Done()
    }()
    wg.Wait()
    foo.Exit()
}

func GetDwarf(i int) (*dwarfomatic.Dwarf){
    current_job := foo.GetCurrentJob(i)
    new_string := ""
    for _, charo := range current_job {
        if(unicode.IsUpper(charo)){
            new_string += " "
        }
        new_string += string(charo)
    }
    dwarf := dwarfomatic.Dwarf{}
    dwarf.Name = foo.GetFirstName(i)
    dwarf.Job = new_string
    dwarf.Mood = foo.GetHappiness(i)
    dwarf.Id = foo.GetId(i)
    return &dwarf
}

func ProcessAvatars(i int){
    image_cmd := exec.Command("convert", "/tmp/output_dwarf" + strconv.Itoa(foo.GetId(i)) + ".bmp", "./public/" + strconv.Itoa(foo.GetId(i)) + ".jpg")
    _ = image_cmd.Run()
    image_cmd = exec.Command("convert", "-resize", "32x32", "/tmp/output_dwarf" + strconv.Itoa(foo.GetId(i)) + ".bmp", "./public/" + strconv.Itoa(foo.GetId(i)) + "_thumb.jpg")
    _ = image_cmd.Run()
}

func ProcessManagerOrders(client *rpc.Client){
    job_html := dwarfomatic.JobHtml{}
    job_html.Captions = make([]string, 0)
    i := 0
    for i < 2858 {
        job_html.Captions = append(job_html.Captions, foo.GetJobType(i))
        i += 1
    }
    client.Call("ABuildServer.BuildJobs", job_html, nil)
}

func ProcessData(client *rpc.Client) {
    i := 0
    df_html := dwarfomatic.DFHtml{}
    df_html.Paused = foo.Paused()
    df_html.Dwarves = make([]dwarfomatic.Dwarf, 0)
    for i < foo.Size() {
        dwarf_html := dwarfomatic.DwarfHtml{}
        dwarf_html.Thoughts = foo.GetThoughts(i)
        dwarf := GetDwarf(i)
        df_html.Dwarves = append(df_html.Dwarves, *dwarf)
        dwarf_html.CurrentJob = dwarf.Job
        dwarf_html.Id = foo.GetId(i)
        dwarf_html.Name = foo.GetFirstName(i)
        dwarf_html.Happiness = foo.GetHappiness(i)
        client.Call("ABuildServer.BuildDwarf", dwarf_html, nil)
        i++;
    }
    client.Call("ABuildServer.BuildDF", df_html, nil)
    i = 0
    cmd := exec.Command("update_game_log.sh")
    _ = cmd.Run()
}
