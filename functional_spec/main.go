package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	cmdmanager "./cmdmanager"
)

const (
	ps1 = "\x1b[32;1m>> \x1b[0m"
)

func main() {
	fmt.Print(ps1)
	cmd := cmdmanager.InitCommand()

	// read if file input exist
	if len(os.Args) > 1 && os.Args[1] != "" {
		cmdFile, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer cmdFile.Close()
		cmdScanner := bufio.NewScanner(cmdFile)
		for cmdScanner.Scan() {
			cmdInput := cmdScanner.Text()
			cmdInput = strings.TrimRight(cmdInput, "\n")
			if cmdInput != "" {
				output := cmd.Run(cmdInput)
				fmt.Println(output)
			}
		}
		fmt.Print(ps1)
		if err := cmdScanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	reader := bufio.NewReader(os.Stdin)
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
