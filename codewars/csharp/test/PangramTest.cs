namespace Pangram
{
  using Xunit;
  
  public class Tests
  {
    [Fact]
    public void SampleTests() 
    {
      Assert.Equal(true, Kata.IsPangram("The quick brown fox jumps over the lazy dog."));
    }
  }
}