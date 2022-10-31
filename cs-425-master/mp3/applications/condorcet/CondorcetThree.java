import java.io.IOException;
import java.util.StringTokenizer;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.fs.Path;
import org.apache.hadoop.io.IntWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Job;
import org.apache.hadoop.mapreduce.Mapper;
import org.apache.hadoop.mapreduce.Reducer;
import org.apache.hadoop.mapreduce.lib.input.FileInputFormat;
import org.apache.hadoop.mapreduce.lib.output.FileOutputFormat;

public class CondorcetThree {

  public static class TokenizerMapper extends Mapper<Object, Text, Text, Text> {

    public void map(Object key, Text value, Context context) throws IOException, InterruptedException {
      String line = value.toString();
      String[] split = line.split("\\s+");

      String keyString = split[0];
      String valueString = split[1];
      
      String A = keyString;
      int sum = Integer.parseInt(valueString);

      String out = A + "," + sum;
      context.write(new Text("1"), new Text(out));
    }
  }

  public static class IntSumReducer extends Reducer<Text, Text, Text, Text> {
    private Text result = new Text();

    public void reduce(Text key, Iterable<Text> values, Context context)
        throws IOException, InterruptedException {

      // get max condorcet winners
      int curMax = -1;
      String winnerString = "";

      for (Text val : values) {
        String line = val.toString();
        String[] sep = line.split(",");

        String A = sep[0];
        int sum = Integer.parseInt(sep[1]);

        if (sum > curMax) {
          curMax = sum;
          winnerString = A;
        } else if (sum == curMax) {
          winnerString += "," + A;
        }
      }

      // output net wins
      result.set(winnerString);
      context.write(new Text("winner(s):"), result);
    }
  }

  public static void main(String[] args) throws Exception {
    Configuration conf = new Configuration();
    Job job = Job.getInstance(conf, "condorcet stage 3");
    job.setJarByClass(CondorcetThree.class);
    job.setMapperClass(TokenizerMapper.class);
    job.setReducerClass(IntSumReducer.class);

    job.setOutputKeyClass(Text.class);
    job.setOutputValueClass(Text.class);
    job.setMapOutputKeyClass(Text.class);
    job.setMapOutputValueClass(Text.class);

    FileInputFormat.addInputPath(job, new Path(args[0]));
    FileOutputFormat.setOutputPath(job, new Path(args[1]));
    System.exit(job.waitForCompletion(true) ? 0 : 1);
  }
}