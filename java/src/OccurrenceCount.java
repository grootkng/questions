import java.io.ByteArrayInputStream;
import java.io.InputStream;
import java.util.Scanner;

public class OccurrenceCount {

  public static int getOccurrenceCount(String toSearch, InputStream stream) throws Exception {
    if (stream == null) {
      return 0;
    }

    int count = 0;
    try (Scanner scanner = new Scanner(stream)) {
      while (scanner.hasNextLine()) {
        final String line = scanner.nextLine();
        int index = line.indexOf(toSearch);
        if (index != -1) {
          count++;
        }
      }
    }
    return count;
  }

  public static void main(String[] args) throws Exception {
    String msg = "Hey! How are you?\nI am good, how about you?\nI am good too.";

    try (InputStream stream = new ByteArrayInputStream(msg.getBytes())) {
      System.out.println(OccurrenceCount.getOccurrenceCount("good", stream));
    }
  }

}
