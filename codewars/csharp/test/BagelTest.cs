namespace Solution
{

    using Xunit;
    public class BagelTest
    {
        [Fact]
        public void TestBagel()
        {
            Bagel bagel = BagelSolver.Bagel;
            Assert.Equal(4, bagel.Value);
        }
    }
}