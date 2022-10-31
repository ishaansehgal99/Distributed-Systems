package main

import (
	"./client"
	"./logger"
	"./server"
	"flag"
	"os"
)

// Entry point of program, parses flags and initiates the correct function
func main() {
	isServer := flag.Bool("server", false, "whether this machine is a server")
	isClient := flag.Bool("client", false, "whether this machine is a client")

	isFirst := flag.Bool("first", false, "whether this machine is the first machine in the SDFS system")

	masterIP := flag.String("masterIP", "", "the string of the master to connect to")
	flag.Parse()

	if (!*isClient && !*isServer) || (*isServer && *isClient) {
		logger.PrintError("Machine must either be either client or server.\nUse the following flags: -server OR -client")
		os.Exit(1)
	}

	if *isClient {
		if *masterIP == "" {
			logger.PrintError("Client masterIP must be non-empty.\nSpecify with the following flag: -masterIP=...")
			os.Exit(1)
		}

		client.Run(*masterIP)
		return
	}

	if (!*isFirst && *masterIP == "") || (*isFirst && *masterIP != "") {
		logger.PrintError("Machine must either be first or have IP address of the master to connect to, but not both.\nUse the following flags: -first -masterIp=<ip>")
		os.Exit(1)
	}

	logger.InfoLogger.Println("Starting the application...")
	server.Run(*isFirst, *masterIP)
}
