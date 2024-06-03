package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"time"
)

type ShuffleInitRequest struct {
}

type ShuffleInitReply struct {
	Status bool
}

func main() {
	args := os.Args[1:]

	server_address := args[0]

	network_interface, err := rpc.DialHTTP("tcp", server_address)
	if err != nil {
		log.Fatal("dialing:", err)
	}

	ping_successful := false

	req := ShuffleInitRequest{}

	var reply ShuffleInitReply

	fmt.Println(server_address)
	for !ping_successful {
		err = network_interface.Call("CTLogCheckerAuditor.PingStartShuffle", req, &reply)
		if err != nil || !reply.Status {
			fmt.Println(err)
			time.Sleep(2 * time.Second)
		}
		if reply.Status {
			ping_successful = true
		}

	}

}
