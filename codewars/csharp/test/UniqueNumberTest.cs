namespace Solution
{
    using Xunit;
    public class SolutionTest
    {
        [Fact]
        public void Test()
        {
            Assert.Equal(1, UniqueNumber.GetUnique(new[] { 1, 2, 2, 2 }));
            Assert.Equal(-2, UniqueNumber.GetUnique(new[] { -2, 2, 2, 2 }));
            Assert.Equal(14, UniqueNumber.GetUnique(new[] { 11, 11, 14, 11, 11 }));
        }
    }
}