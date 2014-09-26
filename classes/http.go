package classes

import "net/http"
import "strings"
import "os/exec"
import "regexp"
import "log"

type BarHandler struct {
}

func (*BarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/df/index.html", http.StatusFound)
}

type SubmitReactionHandler struct {
}

func (*SubmitReactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
}

