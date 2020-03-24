namespace Pangram
{
  using Xunit;
  
  public class Tests
  {
    [Fact]
    public void SampleTests() 
    {
      Assert.Equal(true, Solution.IsPangram("The quick brown fox jumps over the lazy dog."));
    }
  }
}
