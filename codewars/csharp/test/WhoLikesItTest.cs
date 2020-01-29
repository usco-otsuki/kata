using System;
using Xunit;

public class WhoLikesItTest 
{
    [Fact]
    public void Test1()
    {
        Assert.Equal("no one likes this", WhoLikesIt.Likes(new string[0]));
        Assert.Equal("Peter likes this", WhoLikesIt.Likes(new string[] { "Peter" }));
        Assert.Equal("Jacob and Alex like this", WhoLikesIt.Likes(new string[] { "Jacob", "Alex" }));
        Assert.Equal("Max, John and Mark like this", WhoLikesIt.Likes(new string[] { "Max", "John", "Mark" }));
        Assert.Equal("Alex, Jacob and 2 others like this", WhoLikesIt.Likes(new string[] { "Alex", "Jacob", "Mark", "Max" }));
    }
}