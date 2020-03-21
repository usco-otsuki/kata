using System.Linq;
public static class CountCombi
{
   public static int CountCombinations(int money, int[] coins)
   {
       if (money == 0) {
           return 1;
       }
       if (money < 0 || coins.Length == 0) {
           return 0;
       } 
       int count = 0;
       count += CountCombinations(money - coins[0], coins) + CountCombinations(money , coins.Skip(1).ToArray());
       return count;
   }
}