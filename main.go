package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("lotus-miner", "sealing", "/root/go/src/new_rust")
	//err := cmd.Run()
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	if buf != nil {
		fmt.Println(string(buf))
	}

}
