// by Shawn
package main

import (
	"bytes"
	"fmt"
	"os/exec"
	)

func main() {
    in := bytes.NewBuffer(nil)
    cmd := exec.Command("bash")
    cmd.Stdin = in
    err := cmd.Run()
    go func() {
        in.WriteString("")
        fmt.Printf("it's begin")
    }
    if err != nil {
        fmt.Println(err)
        return
    }
}
