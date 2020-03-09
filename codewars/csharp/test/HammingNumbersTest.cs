using Xunit;
public class HammingTests
{
    [Fact]
    public void Test1()
    {
        Assert.Equal(1, Hamming.hamming(1));
        Assert.Equal(2, Hamming.hamming(2));
        Assert.Equal(3, Hamming.hamming(3));
        Assert.Equal(4, Hamming.hamming(4));
        Assert.Equal(5, Hamming.hamming(5));
        Assert.Equal(6, Hamming.hamming(6));
        Assert.Equal(8, Hamming.hamming(7));
        Assert.Equal(9, Hamming.hamming(8));
        Assert.Equal(10, Hamming.hamming(9));
        Assert.Equal(12, Hamming.hamming(10));
        Assert.Equal(15, Hamming.hamming(11));
        Assert.Equal(16, Hamming.hamming(12));
        Assert.Equal(18, Hamming.hamming(13));
        Assert.Equal(20, Hamming.hamming(14));
        Assert.Equal(24, Hamming.hamming(15));
        Assert.Equal(25, Hamming.hamming(16));
        Assert.Equal(27, Hamming.hamming(17));
        Assert.Equal(30, Hamming.hamming(18));
        Assert.Equal(32, Hamming.hamming(19));
    }
}