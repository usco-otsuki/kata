using System;
using Xunit; 

public class CodeWarsTest
{
    [Fact]
    public static void SimpleCase()
    {
        Assert.Equal(3, CountCombi.CountCombinations(4, new[] { 1, 2 }));
    }

    [Fact]
    public static void AnotherSimpleCase()
    {
        Assert.Equal(4, CountCombi.CountCombinations(10, new[] { 5, 2, 3 }));
    }

    [Fact]
    public static void NoChange()
    {
        Assert.Equal(0, CountCombi.CountCombinations(11, new[] { 5, 7 }));
    }
}