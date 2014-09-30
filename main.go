package main

import "net/http"
import "log"
import "strings"
import "strconv"
import "time"
import "net"
import "os"
import "encoding/json"
import artery "github.com/towski/artery/server"
import dwarfomatic "github.com/towski/dwarfomatic/classes"
import "github.com/gorilla/securecookie"


var s *securecookie.SecureCookie

func Authenticate(w http.ResponseWriter, r *http.Request) {
    user := dwarfomatic.UserFindByName(r.FormValue("name"))
    if user != nil {
        if user.Authenticate(r.FormValue("password")) {
            SetCookieHandler(w, r, strconv.Itoa(user.Id))
            http.Redirect(w, r, "/df/user/index.html", http.StatusFound)
        } else {
            http.Redirect(w, r, "/df/user/login.html", http.StatusFound)
        }
    } else {
        http.Redirect(w, r, "/df/user/login.html", http.StatusFound)
    }
}

func AuthenticatedOnly(w http.ResponseWriter, r *http.Request) {
    user := ReadCookieHandler(w, r)
    var data interface{}
    if user != nil {
        log.Println("got a user")
        dwarf := dwarfomatic.FindDwarfByUserId(user.Id)
        if dwarf == nil {
            log.Println("No dwarf!")
            holder := make(map[string]string)
            holder["UserId"] = strconv.Itoa(user.Id)
            data = holder
        } else {
            data = dwarf
        }
    } else {
        log.Println("not a user")
        data = nil
    }
    b, _ := json.Marshal(data)
    log.Println(string(b))
    http.ServeContent(w, r, "name", time.Now(), strings.NewReader(string(b)))
}

func SetUserDwarf(w http.ResponseWriter, r *http.Request) {
    user := ReadCookieHandler(w, r)
    if user != nil {
        dwarf := dwarfomatic.FindDwarfByUserId(user.Id)
        if dwarf == nil{
            log.Println("dwarf id " + r.FormValue("dwarf_id"))
            id, err := strconv.Atoi(r.FormValue("dwarf_id"))
            if err != nil {
                log.Println("err atoing")
            }
            dwarf = dwarfomatic.FindDwarfById(id)
            if dwarf != nil && dwarf.UserId.Valid == false {
                dwarf.SetUserId(user.Id)
                http.Redirect(w, r, "/df/user/index.html", http.StatusFound)
            } else {
                http.Redirect(w, r, "/df/index.html", http.StatusFound)
            }
        } else {
            http.Redirect(w, r, "/df/index.html", http.StatusFound)
        }
    } else {
        http.Redirect(w, r, "/df/user/new.html", http.StatusFound)
    }
}

func SetCookieHandler(w http.ResponseWriter, r *http.Request, user_id string) {
    value := map[string]string{
        "user_id": user_id,
    }
    if encoded, err := s.Encode("cookie-name", value); err == nil {
        cookie := &http.Cookie{
            Name:  "cookie-name",
            Value: encoded,
            Path:  "/",
        }
        http.SetCookie(w, cookie)
    }
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
    user := &dwarfomatic.User{Name:r.FormValue("name"),Password:r.FormValue("password")}
    user.Password = string(dwarfomatic.HashPassword([]byte(user.Password)))
    _, _, _ = net.SplitHostPort(r.RemoteAddr)
    //user.Ip = ip
    if(user.Insert()){
        SetCookieHandler(w, r, strconv.Itoa(user.Id))
        http.Redirect(w, r, "/df/index.html", http.StatusFound)
    } else {
        http.Redirect(w, r, "/df/user/new.html", http.StatusFound)
    }
}

func ReadCookieHandler(w http.ResponseWriter, r *http.Request) *dwarfomatic.User {
    if cookie, err := r.Cookie("cookie-name"); err == nil {
        value := make(map[string]string)
        if err = s.Decode("cookie-name", cookie.Value, &value); err == nil {
            log.Printf("The value of user_id is %q", value["user_id"])
            return dwarfomatic.UserFindById(value["user_id"])
        }
    }
    return nil
}

func DataInit()  {
    artery.Dbmap_global.AddTableWithName(dwarfomatic.Dwarf{}, "Dwarf").SetKeys(false, "Id")
    artery.Dbmap_global.AddTableWithName(dwarfomatic.User{}, "User").SetKeys(true, "Id")
    artery.Dbmap_global.AddTableWithName(dwarfomatic.Item{}, "Item").SetKeys(false, "Id")
    data_server := artery.NewDataServer()
    data_server.Register(&dwarfomatic.Dwarf{})
    data_server.Register(&dwarfomatic.Item{})
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
    http.HandleFunc("/df/artery/user", UserHandler)
    http.HandleFunc("/df/artery/authenticate", Authenticate)
    http.HandleFunc("/df/artery/set_user_dwarf", SetUserDwarf)
    http.HandleFunc("/df/artery/user.json", AuthenticatedOnly)
}

func main (){
    artery.Init("root:mysql@/dwarfomatic")
    secret := os.Getenv("DWARFOMATIC_COOKIE_SECRET")
    if(secret == ""){
        log.Fatal("must set a DWARFOMATIC_COOKIE_SECRET")
    }
    pass_salt := os.Getenv("DWARFOMATIC_SALT")
    if(pass_salt == ""){
        log.Fatal("must set a DWARFOMATIC_SALT")
    }
    dwarfomatic.Salt = []byte(pass_salt)
    var hashKey = []byte(secret)
    s = securecookie.New(hashKey, nil)
    DataInit()
    BuildInit()
    HttpInit()
    log.Fatal(http.ListenAndServe(":8081", nil))
}
