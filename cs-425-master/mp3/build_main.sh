EXE=vm_main
NETID=heesooy2
GROUP=g08

GOOS=linux GOARCH=amd64 go build -o $EXE -v main.go

for vm in $(seq -f "%02g" 1 10)
do
    scp $EXE $NETID@fa20-cs425-$GROUP-$vm.cs.illinois.edu:/home/$NETID/ &
done