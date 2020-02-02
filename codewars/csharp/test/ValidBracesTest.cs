using System;
using Xunit;

public class ValidBracesTest 
{
    [Fact]
    public void Test1()
    {
        Assert.Equal(true, ValidBraces.validBraces( "()" ));
        Assert.Equal(false, ValidBraces.validBraces("[(])"));
    }
}

