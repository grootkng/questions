import java.util.Arrays;

public class SecondLargestNumber {

  public static int findSecondLargest(int[] arr) {
    int largest = -1;
    int secondLargest = largest;

    if (arr.length == 0 || arr.length == 1) {
      return -1;
    }

    Arrays.sort(arr);
    for (int i = 0; i < arr.length; i++) {
      if (arr[i] > largest) {
        secondLargest = largest;
        largest = arr[i];
      }
    }

    return secondLargest;
  }

  public static void main(String[] args) {
    final int[] f = {3, 1, 4, 7, 5, 9, 2};
    final int[] s = {2, 2, 2};
    System.out.println(findSecondLargest(f)); // should return 7
    System.out.println(findSecondLargest(s)); // should return -1
  }
}
