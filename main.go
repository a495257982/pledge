package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmdArgs := []string{"/root/go/src/new_rust", "sealing", "jobs"}

	cmd := exec.Command("lotus-miner", cmdArgs...)
	//err := cmd.Run()
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	if buf != nil {
		fmt.Println(string(buf))
	}
}
