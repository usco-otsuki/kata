namespace SuperPow
{
    using Xunit;

    public class SuperPowTest 
    {
        [Fact]
        public void Test1()
        {
            var solver = new Solution();
            Assert.Equal(8, solver.SuperPow(2, new []{3}));
            Assert.Equal(1024, solver.SuperPow(2, new []{1,0}));
        }
    }
}
