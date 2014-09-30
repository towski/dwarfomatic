package main

import "net/http"
import "testing"
import "fmt"
import "strings"
import "io/ioutil"
import "net/url"
import "net/http/httptest"
import dwarfomatic "github.com/towski/dwarfomatic"
import classes "github.com/towski/dwarfomatic/classes"
import server "github.com/towski/artery/server"

var inited = false

func initIt() {
  if inited == true {
    return
  }
  server.Init("root:mysql@/dwarfomatic_test")
  dwarfomatic.HttpInit()
  dwarfomatic.DataInit()
  inited = true
}

func TestHeader(t *testing.T) {
    initIt()
    resp := httptest.NewRecorder()

    uri := "/df/artery/bar"
    path := "/home/test"
    unlno := "997225821"

    param := make(url.Values)
    param["param1"] = []string{path}
    param["param2"] = []string{unlno}

    //req, err := http.NewRequest("GET", uri+param.Encode(), nil)
    req, err := http.NewRequest("GET", uri, nil)
    if err != nil {
        t.Fatal(err)
    }

    http.DefaultServeMux.ServeHTTP(resp, req)
    if p, err := ioutil.ReadAll(resp.Body); err != nil {
        t.Fail()
    } else {
        if strings.Contains(string(p), "Error") {
            t.Errorf("header response shouldn't return error: %s", p)
        } else if !strings.Contains(string(p), `index.html`) {
            t.Errorf("header response doen't match:\n%s", p)
        }
    }
}

func TestAuthentication(t *testing.T) {
    initIt()
    resp := httptest.NewRecorder()

    uri := "/df/artery/authenticate"
    name := "tracy"
    password := "dicks"
    user := classes.User{Name:name,Password:password}
    if(!user.Insert()){
        t.Errorf("no insert")
    }
    found_user := classes.UserFindByName(name)
    if(found_user == nil){
        t.Errorf("no user")
    }

    param := make(url.Values)
    param["name"] = []string{name}
    param["password"] = []string{password}
    fmt.Println(param.Encode())

    req, err := http.NewRequest("POST", uri, strings.NewReader(param.Encode()))
    if err != nil {
        t.Fatal(err)
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

    http.DefaultServeMux.ServeHTTP(resp, req)
    if p, err := ioutil.ReadAll(resp.Body); err != nil {
        t.Fail()
    } else {
        fmt.Println(string(p))
        if strings.Contains(string(p), "Error") {
            t.Errorf("header response shouldn't return error: %s", p)
        } else if !strings.Contains(string(p), `authenticated`) {
            t.Errorf("header response doesn't match:\n%s", p)
        }
    }

    req, err = http.NewRequest("GET", "/df/artery/user/home", nil)
    if err != nil {
        t.Fatal(err)
    }
    http.DefaultServeMux.ServeHTTP(resp, req)
    if p, err := ioutil.ReadAll(resp.Body); err != nil {
        t.Fail()
    } else {
        fmt.Println(string(p))
    }
}
