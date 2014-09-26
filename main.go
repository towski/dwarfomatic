package main

import "net/http"
import "log"
import server "github.com/towski/artery/server"
import dwarfomatic "github.com/towski/dwarfomatic/classes"

func Init()  {
    server.Dbmap_global.AddTableWithName(dwarfomatic.Dwarf{}, "Dwarf").SetKeys(false, "Id")
    data_server := server.NewDataServer()
    data_server.Register(&dwarfomatic.DwarfServer{})
    build_server := server.NewBuildServer()
    build_server.Register(&dwarfomatic.ABuildServer{})
    go data_server.Start()
    go build_server.Start()
}

func main (){
    http.Handle("/df/artery/bar", &dwarfomatic.BarHandler{})
    http.Handle("/df/artery/submit_reaction", &dwarfomatic.SubmitReactionHandler{})
    server.Init()
    Init()
    log.Fatal(http.ListenAndServe(":8081", nil))
}
