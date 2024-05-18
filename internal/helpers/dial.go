package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"sync"
)

type Message struct {
	Mail string
	JWT  string
}

func SendMail(jwt, mail string, wg *sync.WaitGroup, ch chan<- error) {
	defer close(ch)
	defer wg.Done()
	var message = &Message{
		Mail: mail,
		JWT:  jwt,
	}
	var result, err = json.Marshal(message)

	if err != nil {
		ch <- err
		return
	}

	var con, conErr = net.Dial("tcp", ":6942")

	if conErr != nil {
		ch <- conErr
		return
	}

	defer con.Close()

	_, err = con.Write(result)

	if err != nil {
		ch <- err
		return
	}
	r, err := io.ReadAll(con)
	fmt.Println(string(r), err)
}
