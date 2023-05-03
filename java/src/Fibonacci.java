import java.util.ArrayList;
import java.util.List;

public class Fibonacci {
  List<Integer> list = new ArrayList<Integer>();

  void calculate(int range) {
    this.list.add(0);
    this.list.add(1);

    for (int i = 1; i < range; i++) {
      this.list.add(this.list.get(i-1) + this.list.get(i));
    }
  }

  void showList() {
    for (int i = 0; i < list.size(); i++) {
      System.out.println(list.get(i));
    }
  }

  public static void main(String[] args) {
    final Fibonacci fibonacci = new Fibonacci();
    fibonacci.calculate(20);
    fibonacci.showList();
  }
}
