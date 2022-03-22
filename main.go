package main

import (
	"fmt"
	"os/exec"
)

func main() {

	f, err := exec.Command("ls", "/").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(f))
}
