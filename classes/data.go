package classes

import server "github.com/towski/artery/server"
import _ "github.com/coopernurse/gorp"
import "log"
import _ "errors"
import _ "strconv"
import "reflect"
import "strings"
import "fmt"
import "database/sql"
import "code.google.com/p/go.crypto/pbkdf2"
import "crypto/sha256"

var Salt []byte


func HashPassword(password []byte) string {
    return fmt.Sprintf("%x", pbkdf2.Key(password, Salt, 4096, sha256.Size, sha256.New))
}

type User struct {
    Name string
    Ip string
    Password string
    Id int
}

func (u *User) Insert() bool {
    err := server.Dbmap_global.Insert(u)
    if (err != nil){
        log.Println("no insert ", err)
        return false
    }
    return true
}

func UserFindById(id string) (*User) {
    user := &User{}
    err := server.Dbmap_global.SelectOne(user, "select * from User where Id=?", id)
    if (err != nil){
        log.Println("no find ", err)
        return nil
    }
    return user
}

func UserFindByName(name string) (*User) {
    user := &User{}
    err := server.Dbmap_global.SelectOne(user, "select * from User where Name=?", name)
    if (err != nil){
        log.Println("no find ", err)
        return nil
    }
    return user
}


func (u *User) Authenticate(password string) bool {
    return HashPassword([]byte(password)) == u.Password
}

type CustomIntType int

type Dwarf struct {
    Id int
    HistFigureId sql.NullInt64
    UserId sql.NullInt64
    Name string
    Job string
    Mood string
    Thoughts string
}

func FindDwarfByUserId(id int) (*Dwarf) {
    return FindDwarfByX("UserId", id)
}

func FindDwarfById(id int) (*Dwarf) {
    return FindDwarfByX("Id", id)
}

func FindDwarfByX(thing string, data int) (*Dwarf) {
    dwarf := &Dwarf{}
    err := server.Dbmap_global.SelectOne(dwarf, "select * from Dwarf where " + thing + "=?", data)
    if (err != nil){
        log.Println("no find ", err)
        return nil
    }
    return dwarf
}

func (d *Dwarf) SetUserId(user_id int) {
    d.UserId = sql.NullInt64{Int64:int64(user_id),Valid:true}
    count, err := server.Dbmap_global.Update(d)
    if (err != nil){
        log.Println("no update ", err)
    }
    if (count != 1){
        log.Println("update ", count, err)
    }
}

func (d *Dwarf) GetUserId() sql.NullInt64 {
    result, err := server.Dbmap_global.SelectNullInt("select UserId from Dwarf where Id=? ", d.Id)
    if (err != nil){
        log.Println("no select ", err)
    }
    return result
}

func (*Dwarf) Insert(dwarf Dwarf, result *int) error {
    existing := &Dwarf{}
    FindById(existing, dwarf.Id)
    if existing != nil {
        dwarf.UserId = existing.UserId
    }
    InsertOrUpdate(&dwarf, existing.Id)
    return nil
}

type Model struct {
    Id int
}

type Item struct {
    Model
    Type string
    Quality string
    Description string
    HistFigureId sql.NullInt64
    ItemOwnerId sql.NullInt64
    ItemHolderId sql.NullInt64
    Age int
    Wear int
    DeadDwarf bool
    InChest bool
    OnGround bool
    InInventory bool
    Owned bool
}

func FindById(data interface{}, id int){
    typeOf := reflect.TypeOf(data)
    className := strings.Split(typeOf.String(), ".")[1]
    err := server.Dbmap_global.SelectOne(data, "select * from " + className +" where Id=?", id)
    if (err != nil){
        log.Println("no find ", err)
    }
}

func InsertOrUpdate(data interface{}, id int){
    if id != 0 {
        _, err := server.Dbmap_global.Update(data)
        if (err != nil){
            log.Println("no update ", err)
        }
    } else {
        err := server.Dbmap_global.Insert(data)
        if (err != nil){
            log.Println("no insert ", err)
        }
    }
}

func (*Item) Insert(item Item, result *int) error {
    existing := &Item{}
    FindById(existing, item.Id)
    InsertOrUpdate(&item, existing.Id)
    /*Go("GenericBuildServer.BuildHTML",
        :Id => unit.id,
        :Class => "item",
        :File => "item.html",
        :Data => {:items => items} */
    return nil
}
