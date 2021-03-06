package main

import (
	"fmt"

	"github.com/sonm-io/core/fusrodah/miner"
)

func main() {

	srv, err := miner.NewServer(nil)
	if err != nil {
		fmt.Printf("Error while initialize instanse: %s \r\n", err)
		return
	}
	err = srv.Start()
	if err != nil {
		fmt.Printf("Error while start instanse: %s \r\n", err)
		return
	}

	srv.Serve()

	fmt.Println(srv.GetHubIp())
}
