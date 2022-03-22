package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("lotus-miner", "/c", "/root/go/src/new_rust", "version")
	//err := cmd.Run()
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	if buf != nil {
		fmt.Println(buf)
	}

}
