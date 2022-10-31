# CS 425 MP2 - SDFS Distributed File System

## Prerequisites for compilation
- The Go programming language: https://golang.org/
- Google Protocol Buffers (for Go): https://developers.google.com/protocol-buffers/docs/gotutorial

## Compilation & Running
1. Clone the repo and change directory into `mp2/`
2. Compile the code with `go build main.go`, which will produce the executable `./main`
3. Run the code with the following commands:
    - `./main -server -masterIP=123.123.10.1` (join the SDFS system with master at `masterIP` as a server)
    - `./main -server -first` (start up an SDFS system as the first machine in the system, which means you are the master)
    - `./main -client -masterIP=123.123.10.1` (run as a client, without being a part of the system you can still communicate with the master through get/put/delete/ls commands)

Alternatively, you can run the code without building an exectuable using `go run main.go [args]`, such as `go run main.go -server -masterIP=123.123.10.1`

## Using
After following the above steps, you can simply type into the terminal the commands that you would like to execute. Such as:
  - `put localfile sdfsfile` (put the file with name `localfile` from your local machine into the SDFS system with name `sdfsfile`)
  - `get sdfsfile localfile` (get the file with name `sdfsfile` from the SDFS system into your local machine with name `localfile`)
  - `delete sdfsfile` (delete the file with name `sdfsfile` from the SDFS system)
  - `ls filename` (get a list of IPs where this file is stored)

We have included an executable file (named `./vm_main`) for Linux AMD64 machines, which can be run out of the box without installing any of the prerequisites.