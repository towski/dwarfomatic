package main

import "net/http"
import "log"
import artery "github.com/towski/artery/server"
import dwarfomatic "github.com/towski/dwarfomatic/classes"

func DataInit()  {
    artery.Dbmap_global.AddTableWithName(dwarfomatic.Dwarf{}, "Dwarf").SetKeys(false, "Id")
    data_server := artery.NewDataServer()
    data_server.Register(&dwarfomatic.DwarfServer{})
    go data_server.Start()
}

func BuildInit(){
    build_server := artery.NewBuildServer()
    build_server.Register(&dwarfomatic.ABuildServer{})
    go build_server.Start()
}

func HttpInit(){
    http.Handle("/df/artery/bar", &dwarfomatic.BarHandler{})
    http.Handle("/df/artery/submit_reaction", &dwarfomatic.SubmitReactionHandler{})
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main (){
    artery.Init()
    DataInit()
    BuildInit()
    HttpInit()
}
