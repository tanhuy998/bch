package logCmd

import (
	"fmt"
	"log"
	"net/rpc"
)

func toggleTraceLog(mode string) error {

	rpcClient, err := rpc.Dial("tcp", fmt.Sprintf(":%v", HOST_PORT))

	defer func() {

		if rpcClient != nil {

			rpcClient.Close()
		}
	}()

	if err != nil {

		log.Fatal(err)
	}

	var rep string

	err = rpcClient.Call("LogTools.ToggleTraceLog", mode, &rep)

	if err != nil {

		return err
	}

	log.Println(rep)

	return nil
}
