package main

import "net/rpc"
import "net"
import "log"
import "bufio"
import "fmt"
import "sync"
import "encoding/gob"
import _ "errors"
import "code.google.com/p/goprotobuf/proto"
import _ "github.com/towski/protodwarf/BasicApi.pb"
import "github.com/chai2010/protorpc"

type Header struct {
    magic []byte
    version int
}

func Client() {
    fmt.Println("heyread")
    log.Println("heyread")
    conn, err := net.Dial("tcp", "127.0.0.1:5000")
    if err != nil {
        log.Fatal("Couldn't connect", err)
    }
    encoder := gob.NewEncoder(conn)
    header := Header{}
    header.version = 1
    header.magic = []byte("DFHack?\n")
    if false {
        encoder.Encode(header)
    }
    //encoder.Encode(111)
    conn.Write([]byte{'D', 'F', 'H', 'a', 'c', 'k', '?', '\n', 1, 0, 0, 0})
    //conn.Write('D')
    conn.Write([]byte{1, 0, 0, 0, 1, 0, 0, 0})
    client := rpc.NewClientWithCodec(protorpc.NewClientCodec(conn))
    defer client.Close()
    reader := bufio.NewReader(conn)
    conn.Write([]byte{1})
    //err = client.Call("RunCommand", nil, nil)
        _ = proto.Int32(2)
    for {
        status, _ := reader.ReadString('\n')
        fmt.Println(status)
        fmt.Println("READING SMTHING")
    }
}

bool sendRemoteMessage(CSimpleSocket *socket, int16_t id, const MessageLite *msg, bool size_ready)
{
    int size = size_ready ? msg->GetCachedSize() : msg->ByteSize();
    int fullsz = size + sizeof(RPCMessageHeader);

    uint8_t *data = new uint8_t[fullsz];
    RPCMessageHeader *hdr = (RPCMessageHeader*)data;

    hdr->id = id;
    hdr->size = size;

    uint8_t *pstart = data + sizeof(RPCMessageHeader);
    uint8_t *pend = msg->SerializeWithCachedSizesToArray(pstart);
    assert((pend - pstart) == size);

    int got = socket->Send(data, fullsz);
    delete[] data; 
    return (got == fullsz);
}


func main(){
    messages := make(chan int)
    var wg sync.WaitGroup
    wg.Add(1)
    // Synchronous call
    go func(){
        defer wg.Done()
        Client()
        messages <- 3
    }()
    fmt.Println("heyread")
        //if err != nil {
        //    log.Fatal("error:", err)
        //}
    go func() {
        wg.Wait()
        close(messages)
    }()
    for i := range messages {
        fmt.Println(i)
    }
}
