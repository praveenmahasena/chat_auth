package helpers

import (
	"encoding/json"
	"net"
	"sync"
)

type Message struct{
    Mail string
    JWT string
}

func SendMail(jwt,mail string,wg *sync.WaitGroup,ch chan<-error){
    defer wg.Done()
    var message=&Message{
        Mail: mail,
        JWT: jwt,
    }
    var result,err=json.Marshal(message)

    if err!=nil{
        ch<-err
        return
    }

    var con,conErr=net.Dial("tcp",":6942")

    if conErr!=nil{
        ch<-conErr
        return
    }

    defer con.Close()

    _,err=con.Write(result)

    if err!=nil{
        ch<-err
        return
    }
}
