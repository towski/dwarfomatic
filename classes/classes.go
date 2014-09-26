package classes

import server "github.com/towski/artery/server"
import "strconv"
import "log"

type Dwarf struct {
    Id int
    Name string
    Job string
    Mood string
}

type DFHtml struct {
    Paused int
    Dwarves []Dwarf
}

//pwd, err := os.Getwd()
//if err != nil {
//    fmt.Println(err)
//    os.Exit(1)
//}

type DwarfHtml struct {
    Thoughts string
    Name string
    CurrentJob string
    Happiness string
    Id int
}

type JobHtml struct {
    Captions []string
}

type DwarfServer struct {
}

func (*DwarfServer) Insert(dwarf Dwarf, result *int) error {
    err := server.Dbmap_global.Insert(&dwarf)
    log.Println("insert")
    if (err != nil){
        log.Println("no insert ", err)
    }
    return nil
}

type ABuildServer struct {
}

func (*ABuildServer) BuildDF(dfhtml DFHtml, ret *int) error {
    server.WriteToFile("templates/index.html", "public", dfhtml)
    return nil
}

func (*ABuildServer) BuildDwarf(dwarf_html DwarfHtml, ret *int) error {
    server.WriteToFile("templates/dwarf/id.html", "public/dwarf/" + strconv.Itoa(dwarf_html.Id) + ".html", dwarf_html)
    return nil
}

func (*ABuildServer) BuildJobs(job_html JobHtml, ret *int) error {
    server.WriteToFile("templates/jobs.html", "public", job_html)
    return nil
}
