namespace SortBinaryTreeByLevel
{
    using Xunit;
    using System.Collections.Generic;

    public class SolutionTest
    {
        [Fact]
        public void NullTest()
        {
            Assert.Equal(new List<int>(), Kata.TreeByLevels(null));
        }

        [Fact]
        public void BasicTest()
        {
            Assert.Equal(new List<int>() { 1, 2, 3, 4, 5, 6 }, Kata.TreeByLevels(new Node(new Node(null, new Node(null, null, 4), 2), new Node(new Node(null, null, 5), new Node(null, null, 6), 3), 1)));
        }
    }
}