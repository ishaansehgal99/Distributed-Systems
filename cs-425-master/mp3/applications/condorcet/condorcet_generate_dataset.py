# generate-dataset.py NUM_LINES
import os
import sys
import random

args = sys.argv
n = int(args[1]) # 10 000 000 for 135MB
outfile = "condorcet-{}.txt".format(n) 
voters = ['A', 'B', 'C']

# generate
out = []
for i in range(n):
    random.shuffle(voters)
    line = str(i) + " " + ','.join(voters)
    out.append(line)
    
# save lines to file
outString = '\n'.join(out)
with open(outfile, 'w') as text_file:
    text_file.write("{}".format(outString))
    
print("Wrote {} lines to {}.".format(n, outfile))

"""
./vm_main server -first
./vm_main -server -masterIP=172.22.158.25

put condorcet-10000000.txt p1/1
OR
put condorcet-1000.txt p1/1

maple condorcet_maple_1 6 inter1 p1
juice condorcet_reduce_1 6 inter1 p2/1 delete_input=0 partition_type=hash

maple condorcet_maple_2 6 inter2 p2
juice condorcet_reduce_2 6 inter2 p3/1 delete_input=0 partition_type=hash

maple condorcet_maple_3 6 inter3 p3
juice condorcet_reduce_3 6 inter3 dest delete_input=0 partition_type=hash
"""

"""
HDFS browser: http://172.22.158.25:9870/explorer.html#/
Hadoop browser: http://172.22.158.25:8088/cluster/apps

hadoop/sbin/start-dfs.sh
hadoop/sbin/start-yarn.sh

hdfs dfs -rm -r condorcet/output*

cd Condorcet
hadoop jar CondorcetOne.jar CondorcetOne condorcet/input/condorcet-10000000.txt condorcet/output1
hadoop jar CondorcetTwo.jar CondorcetTwo condorcet/output1 condorcet/output2
hadoop jar CondorcetThree.jar CondorcetThree condorcet/output2 condorcet/output3

hdfs dfs -cat condorcet/output3/*

hadoop/sbin/stop-all.sh
"""