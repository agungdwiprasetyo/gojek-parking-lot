package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	cmdmanager "./cmdmanager"
)

const (
	ps1 = "\x1b[32;1m>> \x1b[0m"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(ps1)
	cmd := cmdmanager.InitCommand()
	for {
		cmdInput, _ := reader.ReadString('\n')
		cmdInput = strings.TrimRight(cmdInput, "\n")
		if cmdInput != "" {
			output := cmd.Run(cmdInput)
			fmt.Println(output)
		}
		fmt.Print(ps1)
	}
}
