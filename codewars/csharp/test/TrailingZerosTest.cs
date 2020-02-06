using System;
using Xunit;

public class TrailingZerosTest 
{
    [Fact]
    public void Test1()
    {
      Assert.Equal(1, TrailingZeros.Solve(5));
      Assert.Equal(6, TrailingZeros.Solve(25));
      Assert.Equal(131, TrailingZeros.Solve(531));
    }
}

