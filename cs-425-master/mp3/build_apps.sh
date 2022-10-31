BUILDDIR=build
NETID=heesooy2
GROUP=g08

mkdir -p $BUILDDIR

# condorcet
rm -f $BUILDDIR/condorcet*
GOOS=linux GOARCH=amd64 go build -o $BUILDDIR/ -v applications/condorcet/condorcet_maple_1.go
GOOS=linux GOARCH=amd64 go build -o $BUILDDIR/ -v applications/condorcet/condorcet_maple_2.go
GOOS=linux GOARCH=amd64 go build -o $BUILDDIR/ -v applications/condorcet/condorcet_maple_3.go
GOOS=linux GOARCH=amd64 go build -o $BUILDDIR/ -v applications/condorcet/condorcet_reduce_1.go
GOOS=linux GOARCH=amd64 go build -o $BUILDDIR/ -v applications/condorcet/condorcet_reduce_2.go
GOOS=linux GOARCH=amd64 go build -o $BUILDDIR/ -v applications/condorcet/condorcet_reduce_3.go

for vm in $(seq -f "%02g" 1 10)
do
    # executables
    for i in {1..3}
    do
        scp $BUILDDIR/condorcet_maple_$i $NETID@fa20-cs425-$GROUP-$vm.cs.illinois.edu:/home/$NETID/ &
        scp $BUILDDIR/condorcet_reduce_$i $NETID@fa20-cs425-$GROUP-$vm.cs.illinois.edu:/home/$NETID/ &
    done

    # python data generation scripts
    scp applications/condorcet/condorcet_generate_dataset.py $NETID@fa20-cs425-$GROUP-$vm.cs.illinois.edu:/home/$NETID/ &
    scp applications/condorcet/condorcet_ground_truth.py $NETID@fa20-cs425-$GROUP-$vm.cs.illinois.edu:/home/$NETID/ &
done

# # champaign
# rm -f $BUILDDIR/champaign*
# GOOS=linux GOARCH=amd64 go build -o $BUILDDIR/ -v applications/champaign/champaign_maple.go
# GOOS=linux GOARCH=amd64 go build -o $BUILDDIR/ -v applications/champaign/champaign_reduce.go

# for vm in $(seq -f "%02g" 1 10)
# do
#     scp $BUILDDIR/champaign_maple $NETID@fa20-cs425-$GROUP-$vm.cs.illinois.edu:/home/$NETID/ &
#     scp $BUILDDIR/champaign_reduce $NETID@fa20-cs425-$GROUP-$vm.cs.illinois.edu:/home/$NETID/ &
# done

# # basic
# rm -f $BUILDDIR/basic*
# cp applications/basic/basic_maple $BUILDDIR/
# cp applications/basic/basic_reduce $BUILDDIR/

# for vm in $(seq -f "%02g" 1 10)
# do
#     scp $BUILDDIR/basic_maple $NETID@fa20-cs425-$GROUP-$vm.cs.illinois.edu:/home/$NETID/ &
#     scp $BUILDDIR/basic_reduce $NETID@fa20-cs425-$GROUP-$vm.cs.illinois.edu:/home/$NETID/ &
# done