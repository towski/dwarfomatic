package main

import "net/http"
import _ "html"
import "time"
import "fmt"
import "strings"
import "regexp"
import "os/exec"
import "log"
import _ "reflect"
import _ "unsafe"
import server "github.com/towski/artery/server"
import dwarfomatic "github.com/towski/dwarfomatic/classes"

type timeHandler struct {
  zone *time.Location
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  tm := time.Now().In(th.zone).Format(time.RFC1123)
  w.Write([]byte("The time is: " + tm))
}

func newTimeHandler(name string) *timeHandler {
  return &timeHandler{zone: time.FixedZone(name, 0)}
}

func fooHandler(){
    log.Fatal("hey")
}

//var data_client *rpc.Client
//var build_client *rpc.Client
func Init()  {
    server.Dbmap_global.AddTableWithName(dwarfomatic.Dwarf{}, "Dwarf").SetKeys(false, "Id")
    data_server := server.NewDataServer()
    data_server.Register(&dwarfomatic.DwarfServer{})
    //server.Register(&Dwarf)
    build_server := server.NewBuildServer()
    build_server.Register(&dwarfomatic.ABuildServer{})
    go data_server.Start()
    go build_server.Start()
   // go 
    //build_client, _ = rpc.Dial("unix", "/tmp/build.sock")
//    StartDBWriter(Post_channel, dbmap)
    // Returns the user with the given id 
}

func main (){
    http.Handle("/foo", newTimeHandler("EST"))
    http.HandleFunc("/df/artery/foo.json", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Your changes will appear shortly")
    })
    http.HandleFunc("/df/artery/bar", func(w http.ResponseWriter, r *http.Request) {
        //fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
        //p := &server.Post{Title: "yolo"}
        //fmt.Println(p)
        http.Redirect(w, r, "/df/index.html", http.StatusFound)
    })
    http.HandleFunc("/df/artery/submit_reaction", func(w http.ResponseWriter, r *http.Request) {
        //title := r.URL.Path[len("/edit/"):]
        job_name := strings.Trim(strings.Replace(r.FormValue("name"), "\n", " ", -1), " ")
        re := regexp.MustCompile("[^A-Za-z' _]")
        job_name = re.ReplaceAllString(job_name, " ")
        command := re.ReplaceAllString(r.FormValue("command"), " ")
        command = strings.Trim(command, " ")
        log.Println(command)
        if !(command == "submit_reaction" || command == "fpause" || command == "unpause") {
            command = "ls"
        }
        log.Println(command + " " + job_name)
        cmd := exec.Command("/home/towski/save/df_linux/dfhack-run", command, job_name)
         _ = cmd.Start()
        cmd = exec.Command("/home/towski/save/df_linux/dfhack-run", "show_orders")
         _ = cmd.Start()
        http.Redirect(w, r, "/df/index.html", http.StatusFound)

    })
    server.Init()
    Init()
    log.Fatal(http.ListenAndServe(":8081", nil))
}
