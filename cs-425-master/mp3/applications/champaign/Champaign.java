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

public class Champaign {

  public static class TokenizerMapper extends Mapper<Object, Text, Text, IntWritable> {

    private final static IntWritable one = new IntWritable(1);
    private Text word = new Text();

    // SplitAtCommas - https://stackoverflow.com/questions/59297737/go-split-string-by-comma-but-ignore-comma-within-double-quotes
    // private List<String> splitAtCommas(String s) {
    //   List<String> res = new ArrayList<>();
      
    //   int beg = 0;
    //   boolean inString = false;
    
    //   for (int i = 0; i < s.length(); i++) {
    //     if (s.charAt(i) == ',' && !inString) {
    //       res.add(s.substring(beg, i));
    //     } else if (s.charAt(i) == '"') {
    //       if (!inString) {
    //         inString = true;
    //       } else if (i > 0 && s.charAt(i - 1) != '\\') {
    //         inString = false;
    //       }
    //     }
    //   }

    //   res.add(s.substring(beg));
      
    //   return res;
    // }

    public void map(Object key, Text value, Context context) throws IOException, InterruptedException {
      String line = value.toString();
      // split csv columns, but ignore commas inside quotes 
      // https://stackoverflow.com/questions/15738918/splitting-a-csv-file-with-quotes-as-text-delimiter-using-string-split
      String[] split = line.split(",(?=([^\"]*\"[^\"]*\")*[^\"]*$)");
  
      String zipCode = split[40];
      boolean isResidential = split[13].equals("Y");
      boolean isMailable = split[14].equals("Y");
      boolean isActive = split[18].equals("Active");
  
      if (isResidential && isMailable && isActive) {
        word.set(zipCode);
        context.write(word, one);
      }
    }
  }

  public static class IntSumReducer extends Reducer<Text, IntWritable, Text, IntWritable> {
    private IntWritable result = new IntWritable();

    public void reduce(Text key, Iterable<IntWritable> values, Context context)
        throws IOException, InterruptedException {
      int sum = 0;
      for (IntWritable val : values) {
        sum += val.get();
      }

      result.set(sum);
      context.write(key, result);
    }
  }

  public static void main(String[] args) throws Exception {
    Configuration conf = new Configuration();
    Job job = Job.getInstance(conf, "champaign stage 1");
    job.setJarByClass(Champaign.class);
    job.setMapperClass(TokenizerMapper.class);
    job.setCombinerClass(IntSumReducer.class);
    job.setReducerClass(IntSumReducer.class);
    job.setOutputKeyClass(Text.class);
    job.setOutputValueClass(IntWritable.class);
    FileInputFormat.addInputPath(job, new Path(args[0]));
    FileOutputFormat.setOutputPath(job, new Path(args[1]));
    System.exit(job.waitForCompletion(true) ? 0 : 1);
  }
}