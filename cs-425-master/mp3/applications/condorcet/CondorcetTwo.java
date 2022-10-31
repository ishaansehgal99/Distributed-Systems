import java.io.IOException;
import java.util.StringTokenizer;
import java.util.HashMap;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.fs.Path;
import org.apache.hadoop.io.IntWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Job;
import org.apache.hadoop.mapreduce.Mapper;
import org.apache.hadoop.mapreduce.Reducer;
import org.apache.hadoop.mapreduce.lib.input.FileInputFormat;
import org.apache.hadoop.mapreduce.lib.output.FileOutputFormat;

public class CondorcetTwo {

  public static class TokenizerMapper extends Mapper<Object, Text, Text, Text> {

    public void map(Object key, Text value, Context context) throws IOException, InterruptedException {
      String line = value.toString();
      String[] split = line.split("\\s+");

      String keyString = split[0];
      String valueString = split[1];
      
      String A = Character.toString(keyString.charAt(0));
      String B = Character.toString(keyString.charAt(1));

      int sum = Integer.parseInt(valueString);

      // first output (key=A, out=(B, sum, 0))
      String out = B + "," + sum + "," + 0;
      context.write(new Text(A), new Text(out));

      // second output (key=B, out=(A, 0, sum))
      out = A + "," + 0 + "," + sum;
      context.write(new Text(B), new Text(out));
    }
  }

  public static class IntSumReducer extends Reducer<Text, Text, Text, IntWritable> {
    private IntWritable result = new IntWritable();

    public void reduce(Text key, Iterable<Text> values, Context context)
        throws IOException, InterruptedException {

      // populate map of net votes over opponents
      HashMap<String, Integer> m = new HashMap<>();
      for (Text val : values) {
        // val = B,AoverB,BoverA
        String line = val.toString();
        String[] sep = line.split(",");

        String opp = sep[0];
        int pos = Integer.parseInt(sep[1]);
        int neg = Integer.parseInt(sep[2]); 

        int cur = m.containsKey(opp) ? m.get(opp) : 0;
        m.put(opp, cur + pos - neg);
      }

      // count number of wins over other opponents
      int sum = 0;
      for (int val : m.values()) {
        if (val > 0) {
          sum++;
        }
      }

      // output net wins
      result.set(sum);
      context.write(key, result);
    }
  }

  public static void main(String[] args) throws Exception {
    Configuration conf = new Configuration();
    Job job = Job.getInstance(conf, "condorcet stage 2");
    job.setJarByClass(CondorcetTwo.class);
    job.setMapperClass(TokenizerMapper.class);
    job.setReducerClass(IntSumReducer.class);

    job.setOutputKeyClass(Text.class);
    job.setOutputValueClass(IntWritable.class);
    job.setMapOutputKeyClass(Text.class);
    job.setMapOutputValueClass(Text.class);

    FileInputFormat.addInputPath(job, new Path(args[0]));
    FileOutputFormat.setOutputPath(job, new Path(args[1]));
    System.exit(job.waitForCompletion(true) ? 0 : 1);
  }
}