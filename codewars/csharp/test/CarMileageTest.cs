using Xunit;
using System.Collections.Generic;

public class CarMileageTest
{
    [Fact]
    public void Test1()
    {
        Assert.Equal(0, CarMileage.IsInteresting(3, new List<int>() { 1337, 256 }));
        Assert.Equal(1, CarMileage.IsInteresting(1336, new List<int>() { 1337, 256 }));
        Assert.Equal(2, CarMileage.IsInteresting(1337, new List<int>() { 1337, 256 }));
        Assert.Equal(0, CarMileage.IsInteresting(11208, new List<int>() { 1337, 256 }));
        Assert.Equal(1, CarMileage.IsInteresting(11209, new List<int>() { 1337, 256 }));
        Assert.Equal(2, CarMileage.IsInteresting(11211, new List<int>() { 1337, 256 }));
    }
}