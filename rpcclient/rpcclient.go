package rpcclient

import (
	"fmt"
	"net/rpc"
	"os"
	"os/exec"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Response struct {
	Message string
}

type Request struct {
	Name string
}

var (
	ip      = "127.0.0.1"
	port    = "4555"
	address = fmt.Sprintf("%s:%s", ip, port)
	outfile *os.File
)

func Client() {
	wg.Add(1)
	c, err := rpc.Dial("tcp", "127.0.0.1:4555")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result string
	err = c.Call("Server.Add", [2]int64{10, 20}, &result)
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()
}

func StartServer() {
	wg.Add(1)
	var err error
	cmd := exec.Command("./goverifier", "-rpc=true", fmt.Sprintf("-port=%s", port))
	outfile, err = os.Create("../logs/verifier.log")
	cmd.Stdout = outfile
	cmd.Stderr = outfile
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	err = cmd.Start()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}
	wg.Done()

	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	wg.Add(1)

	c, err := rpc.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	var result *Response
	err = c.Call("Server.StartServer", Request{"StartServer"}, &result)
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()

}

func StopServer() {
	wg.Add(1)
	c, err := rpc.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	var result *Response
	err = c.Call("Server.StopServer", Request{"StopServer"}, &result)
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()
}

func RestartServer() {
	wg.Add(1)
	c, err := rpc.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println("Connected...")
	var result *Response
	err = c.Call("Server.RestartServer", Request{"RestartServer"}, &result)
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()
}
