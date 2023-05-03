import java.util.ArrayList;
import java.util.List;

public class CreateDataStructure {
  List<Integer> arr = new ArrayList<>();

  /**
   * Adds/appends list of integers at the end of internal list.
   */
  public void append(int[] list) {
    for (int i = 0; i < list.length; i++) {
      this.arr.add(list[i]);
    }
  }

  /**
   * Returns boolean representing if any three consecutive integers in the
   * internal list have given total.
   */
  public boolean contains(int total) {
    int n = arr.size();
    for (int i = 0; i < n - 2; i++) {
      int sum = arr.get(i) + arr.get(i+1) + arr.get(i+2);
      if (sum == total) {
        return true;
      }
    }
    return false;
  }

  public static void main(String[] args) {
    CreateDataStructure movingTotal = new CreateDataStructure();

    movingTotal.append(new int[] { 1, 2, 3, 4 });

    System.out.println(movingTotal.contains(6));
    System.out.println(movingTotal.contains(9));
    System.out.println(movingTotal.contains(12));
    System.out.println(movingTotal.contains(7));

    movingTotal.append(new int[] { 5 });

    System.out.println(movingTotal.contains(6));
    System.out.println(movingTotal.contains(9));
    System.out.println(movingTotal.contains(12));
    System.out.println(movingTotal.contains(7));
  }
}
