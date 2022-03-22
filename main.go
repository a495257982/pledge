package main

import (
	"fmt"
	"os/exec"
)

func main() {

	f, err := exec.Command("/root/go/src/pledge/lotus-miner ", "sealing jobs").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(f))
}
