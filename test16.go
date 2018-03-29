package main

import (
	"os/exec"
	"os"
	"fmt"
)

func main() {
	cmd := exec.Command("/bin/sh", "-c", "git clone --progress http://github.com/leanote/leanote /Users/life/Desktop/tmp")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("cmd.Output: ", err)
		return
	}
}