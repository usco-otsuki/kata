using Xunit;

public class Rot13Test 
{
    [Fact]
    public void Test1()
    {
        Assert.Equal("ROT13 example.", Rot13.Solve("EBG13 rknzcyr."));
    }
}