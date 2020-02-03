using Xunit;

public class PileOfCubesTest
{
    [Fact]
    public void Test1()
    {
        Assert.Equal(2022, PileOfCubes.findNb(4183059834009));
        Assert.Equal(-1, PileOfCubes.findNb(24723578342962));
        Assert.Equal(4824, PileOfCubes.findNb(135440716410000));
        Assert.Equal(3568, PileOfCubes.findNb(40539911473216));
    }
}