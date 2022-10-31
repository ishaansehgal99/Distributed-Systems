# Our system
See README in applications/condorcet

# Hadoop

## dataset
https://gis-cityofchampaign.opendata.arcgis.com/datasets/address-points

## compilation:
(instructions: https://hadoop.apache.org/docs/current/hadoop-mapreduce-client/hadoop-mapreduce-client-core/MapReduceTutorial.html)

Env variables setup:

```
export JAVA_HOME=/usr/java/default
export PATH=${JAVA_HOME}/bin:${PATH}
export HADOOP_CLASSPATH=${JAVA_HOME}/lib/tools.jar
```

Upload dataset into HDFS

```
hdfs dfs -mkdir -p champaign/input
hdfs dfs -put champaign-dataset.csv champaign/input
```

Compile and run the hadoop program

```
hadoop com.sun.tools.javac.Main Champaign.java
jar cf Champaign.jar Champaign*.class

hadoop jar Champaign.jar Champaign champaign/input champaign/output

hdfs dfs -cat champaign/otput/*
```

## expected output
```
61820 - 19789

61821 - 13647

61822 - 10751
```