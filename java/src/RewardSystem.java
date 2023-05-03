import java.util.HashMap;

public class RewardSystem {
  final private HashMap<String, Integer> customers;

  public RewardSystem() {
    this.customers = new HashMap<String, Integer>();
  }

  public void earnPoints(String customerName, int points) {
    if (points <= 0) {
      return;
    }

    int currentPoints = this.customers.getOrDefault(customerName, 0);
    if (currentPoints == 0) { // customer is new, award 500 points
      this.customers.put(customerName, 500 + points);
    } else {
      this.customers.put(customerName, currentPoints + points);
    }
  }

  public int spendPoints(String customerName, int points) {
    if (points <= 0) {
      return this.customers.getOrDefault(customerName, 0);
    }

    int currentPoints = this.customers.getOrDefault(customerName, 0);
    if (currentPoints < points) {
      return 0;
    } else {
      this.customers.put(customerName, currentPoints - points);
      return this.customers.getOrDefault(customerName, 0);
    }
  }

  public static void main(String[] args) {
    RewardSystem rewardPoints = new RewardSystem();
    rewardPoints.earnPoints("John", 520);
    System.out.println(rewardPoints.spendPoints("John", 200));
  }
}
