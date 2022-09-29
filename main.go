package main

import (
	"bufio"
	"fmt"
	"go-atm/config"
	"go-atm/modules/controller"
	"os"
	"strings"
)

func main() {
	config.Connect()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		command := strings.Split(line, " ")
		Router(command)
		if line == "quit" {
			fmt.Println("Quitting...")
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error encountered:", err)
	}
}

func Router(cmd []string) {
	switch cmd[0] {
	case "login":
		controller.Login(cmd[1])
	case "logout":
		controller.Logout()
	case "deposit":
		controller.Deposit(cmd[1])
	case "transfer":
		controller.Transfer()
	case "withdraw":
		controller.Withdraw(cmd[1])
	}
}
