# Our system

## dataset

generation (3 candidates)

```
python3 condorcet_generate_data.py NUM_LINES
```

check ground truth on 1 cpu

```
python3 condorcet_ground_truth.py DATA_FILE
```

## usage

```
./vm_main server -first
./vm_main -server -masterIP=172.22.158.25

put condorcet-10000000.txt p1/1
OR
put condorcet-1000.txt p1/1

maple condorcet_maple_1 5 inter1 p1
juice condorcet_reduce_1 5 inter1 p2/1 delete_input=0 partition_type=hash

maple condorcet_maple_2 5 inter2 p2
juice condorcet_reduce_2 5 inter2 p3/1 delete_input=0 partition_type=hash

maple condorcet_maple_3 5 inter3 p3
juice condorcet_reduce_3 5 inter3 dest delete_input=0 partition_type=hash
```

# Hadoop

## helpful

HDFS browser: http://172.22.158.25:9870/explorer.html#/
Hadoop browser: http://172.22.158.25:8088/cluster/apps

Hadoop setup guide: https://www.linode.com/docs/guides/how-to-install-and-set-up-hadoop-cluster/

## usage

Start up HDFS and YARN and clean old output

```
hadoop/sbin/start-dfs.sh
hadoop/sbin/start-yarn.sh

hdfs dfs -rm -r condorcet/output*
```

Uploading and compiling dataset and JARs:

* see applications/champaign/README.md

```
hadoop jar CondorcetOne.jar CondorcetOne condorcet/input/condorcet-10000000.txt condorcet/output1
hadoop jar CondorcetTwo.jar CondorcetTwo condorcet/output1 condorcet/output2
hadoop jar CondorcetThree.jar CondorcetThree condorcet/output2 condorcet/output3

hdfs dfs -cat condorcet/output3/*
```

Clean up

```
hadoop/sbin/stop-all.sh
```