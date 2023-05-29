import java.util.HashMap;
import java.util.Map;

public class ReplaceEmojiWithImageTag {
  private static final Map<String, String> emoji = new HashMap<>();

  static {
    emoji.put(":smile:", "smile");
    emoji.put(":heart:", "heart");
  }

  private static String replaceEmojiWithImageTags(String input) {
    StringBuilder result = new StringBuilder();
    int startIndex = 0;
    int endIndex;

    while (startIndex < input.length()) {
      int emojiStartIndex = input.indexOf(':', startIndex);
      if (emojiStartIndex == -1) {
        result.append(input.substring(startIndex));
        break;
      }

      int emojiEndIndex = input.indexOf(':', emojiStartIndex + 1);
      if (emojiEndIndex == -1) {
        result.append(input.substring(startIndex));
        break;
      }

      endIndex = emojiEndIndex + 1;

      String emojiName = input.substring(emojiStartIndex, endIndex);
      String imageUrl = emoji.getOrDefault(emojiName, "") + ".png";
      String imgTag = "<img src='" + imageUrl + "'>";

      result.append(input.substring(startIndex, emojiStartIndex));
      result.append(imgTag);

      startIndex = endIndex;
    }

    return result.toString();
  }

  public static void main(String[] args) {
    System.out.println(replaceEmojiWithImageTags("I'm feeling :smile: today! :heart:"));
  }
}
