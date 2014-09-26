package classes

import server "github.com/towski/artery/server"
import "log"

type Dwarf struct {
    Id int
    Name string
    Job string
    Mood string
}

func (*Dwarf) Insert(dwarf Dwarf, result *int) error {
    err := server.Dbmap_global.Insert(&dwarf)
    if (err != nil){
        log.Println("no insert ", err)
    }
    return nil
}
