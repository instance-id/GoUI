package rpcclient

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"

	. "github.com/instance-id/GoUI/components"
	"github.com/smallnest/rpcx/client"
)

var (
	wg             sync.WaitGroup
	ad             AccessData
	verifierStatus bool
	rpcStatus      bool
	xC             client.XClient
)

type AccessData struct {
	key []byte
}

type Reply struct {
	Message   string
	RunCheck  bool
	RPCCheck  bool
	ProcessID int
	Key       []byte
}

type Args struct {
	Key     []byte
	Name    string
	Message string
	Ctx     map[string]interface{}
}

var (
	ip      = "localhost"
	port    = "14555"
	address = fmt.Sprintf("%s:%s", ip, port)
	outfile *os.File
	dev     *bool
)

type RpcClient struct {
}

func ReturnRunning() (bool, bool) {
	return verifierStatus, rpcStatus
}

func send(args *Args, reply *Reply) (*Args, *Reply) {
	err := xC.Call(context.Background(), "RpcRequestHandler", args, reply)
	if err != nil {
		_ = fmt.Errorf("failed to call: %v", err)
	}
	return args, reply
}

func CheckRunning() (bool, bool) {

	reply := &Reply{}
	args := &Args{Name: RPC_STATUS, Key: ad.key, Message: "You got this!"}

	err := xC.Call(context.Background(), "RpcRequestHandler", args, reply)
	if err != nil {
		_ = fmt.Errorf("failed to call: %v", err)
	}
	fmt.Printf("Response: %s : Pid: %v \n", reply.Message, reply.ProcessID)

	if reply.RunCheck == false {
		verifierStatus = false
		rpcStatus = false
	} else {
		verifierStatus = reply.RunCheck
		rpcStatus = reply.RPCCheck
	}

	fmt.Printf("Runcheck: %v RPCCheck: %v \n", reply.RunCheck, reply.RPCCheck)
	fmt.Printf("Verifier: %v RPCServer: %v \n", verifierStatus, rpcStatus)
	fmt.Printf("Whats the message?? %s", reply.Message)
	return verifierStatus, rpcStatus
}

func StartVerifier() {
	wg.Add(1)
	var err error
	var path string
	if *dev {
		path = "/home/mosthated/_dev/programming/go/src/github.com/instance-id/GoVerifier-dgo/main"

	} else {
		path = "./goverifier"
	}
	cmd := exec.Command(path, "-rpc=true", fmt.Sprintf("-port=%s", port))
	err = cmd.Start()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}
	rpcStatus = true
	err = cmd.Process.Release()
	if err != nil {

	}
	wg.Done()
	fmt.Printf("Running Timer to allow Verifier to load...")

}

func StartServer() bool {

	if !rpcStatus {
		go StartVerifier()
		runtime.Gosched()
		timer1 := time.NewTimer(5 * time.Second)
		<-timer1.C
		fmt.Printf("Starting Verifier discord services...")
	}

	wg.Add(1)

	// --- Send RPC request -----------
	reply := &Reply{}
	args := &Args{Name: RPC_START, Key: ad.key}
	args, reply = send(args, reply)
	fmt.Printf("Response: %s ", reply.Message)

	verifierStatus = reply.RunCheck
	wg.Done()
	return verifierStatus
}

//
func RestartServer() {

	wg.Add(1)

	reply := &Reply{}
	args := &Args{Name: RPC_RESTART, Key: ad.key}
	args, reply = send(args, reply)
	fmt.Printf("Response: %s ", reply.Message)

	wg.Done()
}

//
func StopServer() {

	wg.Add(1)

	reply := &Reply{}
	args := &Args{Name: RPC_STOP, Key: ad.key}
	args, reply = send(args, reply)
	fmt.Printf("Response: %s ", reply.Message)

	wg.Done()
}

func CreateClient(phrase string, key string, isDev *bool) {
	dev = isDev
	GetKey(phrase, key)
	go runClient()
	runtime.Gosched()
	CheckRunning()

}

func runClient() {

	wg.Add(1)

	cli := client.NewPeer2PeerDiscovery("tcp@"+address, "")
	xClient := client.NewXClient("Server", client.Failtry, client.RandomSelect, cli, client.DefaultOption)
	xC = xClient
	defer xC.Close()

	wg.Wait()
	fmt.Printf("Run Client Closing")

}

func GetKey(phrase string, key string) {
	keyChar := key[len(key)-13:]
	encrypted := encrypt([]byte(keyChar), phrase)
	ad.key = encrypted
}

func CloseConnection() {
	wg.Done()
}

func Initialize() {

	isDev := flag.Bool("dev", false, "Run with developer paths? (Looks in different folder for GoVerifier)")
	flag.Parse()
	key := Cntnrs.Dac.System.Token
	phrase := Cntnrs.Dac.Discord.Guild
	CreateClient(phrase, key, isDev)
}
