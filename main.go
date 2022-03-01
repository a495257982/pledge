package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("/root/go/src/new_rust/lotus-miner sealing jobs")
	out, _ := cmd.Output()
	fmt.Println(string(out))
	fmt.Println("aaa")
}
