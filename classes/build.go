package classes

import server "github.com/towski/artery/server"
import "strconv"

type DFHtml struct {
    Paused int
    Dwarves []Dwarf
}

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
