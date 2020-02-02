public class DigitalRootSolver
{
  private long Sum(long n) {
      long sum = 0; 
      while (n > 0) {
          sum += n % 10;
          n /= 10;
      }
      return sum;
  }
  public int DigitalRoot(long n)
  {
    while (n >= 10) {
        n = Sum(n);
    }      
    return (int)n;
  }
}