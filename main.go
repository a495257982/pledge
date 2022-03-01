package main

import (
	"fmt"
	"os/exec"
)

func main() {

	cmd := exec.Command("echo", "'helloworld!'")
	out, _ := cmd.Output()
	fmt.Println(out)
	fmt.Println("aaa")
}
