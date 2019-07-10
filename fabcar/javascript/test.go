package main

import (
	"fmt"
	"os/exec"
)

func main() {
	c := "node query.js"
	cmd := exec.Command("sh","-c", c)
	f,err := cmd.Output()
	_ = err
	fmt.Println(string(f))
}
