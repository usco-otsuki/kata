using Xunit;

public class DigitalRootTest 
{
    [Fact]
    public void Test1()
    {
        var solver = new DigitalRootSolver();
        Assert.Equal(7, solver.DigitalRoot(16));       
        Assert.Equal(6, solver.DigitalRoot(456)); 
    }
}