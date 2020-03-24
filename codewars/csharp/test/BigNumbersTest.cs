namespace BigNumbers {
  using Xunit;

  public class SolutionTest
  {
    [Fact]
    public void MyTest()
    {
      Assert.Equal("3000000000000000000000001",Solver.Add("1000000000000000000000000", "2000000000000000000000001"));
    }
  }
}
