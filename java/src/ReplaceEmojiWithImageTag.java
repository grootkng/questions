import java.util.HashMap;
import java.util.Map;

public class ReplaceEmojiWithImageTag {
	private static final Map<String, String> emoji = new HashMap<>();

	static {
		emoji.put(":smile:", "smile");
		emoji.put(":heart:", "heart");
		emoji.put(":thumbs_up:", "thumbs_up");
	}

	private static String replaceEmojiWithImageTags(String input) {
		int startIndex = 0;

		for (int i = 0; i < input.length(); i++) {
			char currentChar = input.charAt(i);

			if (currentChar == ':') {
				if (startIndex == 0) {
					startIndex = i;
				} else {
					String emojiName = input.substring(startIndex, i + 1);
					String emojiHtmlElement = "<img src='" + emoji.getOrDefault(emojiName, "") + ".png'/>";
					input = input.replace(emojiName, emojiHtmlElement);
					startIndex = 0;
				}
			}
		}

		return input;
	}

	public static void main(String[] args) {
		System.out.println(replaceEmojiWithImageTags("I'm feeling :smile: today! :heart:"));
	}
}
