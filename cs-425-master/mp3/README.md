# CS 425 MP3 - Simplified MapReduce (MapleJuice)

Group g08 - heesooy2, asdale2, isehgal2

## Prerequisites for compilation
- The Go programming language: https://golang.org/
- Google Protocol Buffers (for Go): https://developers.google.com/protocol-buffers/docs/gotutorial

## Compilation & Running
1. Clone the repo and change directory into `mp3/`

2. Compile the code with `go build main.go`, which will produce the executable `./main` (we have also provided sh scripts for deploying the system and MapleJuice applications to the VMs).

3. Run the code with the following commands:
    - `./main -server -first` (start up MapleJuice system as the first machine in the system, which means you are the master)
    - `./main -server -masterIP=123.123.10.1` (join MapleJuice system with master at `masterIP` as a worker node)
    
Alternatively, you can run the code without building an exectuable using `go run main.go [args]`, such as `go run main.go -server -masterIP=123.123.10.1`

## Usage

### Membership & Filesystem

You can use the command prompt to interact with the membership protocol. The commands are the same as MP1.

To upload/modify files to the system, once again, the commands are the same as MP2:
  - `put localfile sdfsfile` (put the file with name `localfile` from your local machine into the SDFS system with name `sdfsfile`)
  - `get sdfsfile localfile` (get the file with name `sdfsfile` from the SDFS system into your local machine with name `localfile`)
  - `delete sdfsfile` (delete the file with name `sdfsfile` from the SDFS system)
  - `ls filename` (get a list of IPs where this file is stored)

### MapleJuice

Once the input data is in the SDFS, you can use the following commands to initiate Maple and Juice tasks. These match the format in the MP instructions:
  - `maple <maple_exe> <num_workers> <intermediate_prefix> <input_folder>`
    - start maple task with the <maple_exe> executable using data from <input_folder>, saving output to the folder named <intermediate_prefix>
  - `juice <juice_exe> <num_workers> <intermediate_prefix> <output_filename> <delete_input={0, 1}> <partition_type={range,hash}>`
    - start juice task from files with <intermediate_prefix> and output to <output_filename>. Optionally choose to delete intermediate files and choose key partitioning strategy

We have included an executable file (named `./vm_main`) for Linux AMD64 machines, which can be run out of the box without installing any of the prerequisites. This is compiled using the command in the build_main.sh script. Lastly, we have also included example MapleJuice programs under the `applications/` folder with READMEs explaining detailed usage.