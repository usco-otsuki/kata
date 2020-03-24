using System;
using Xunit;

namespace GreedIsGood
{
    public static class ScoreCheckerTest
    {
        [Fact]
        public static void ShouldBeWorthless()
        {
            Assert.Equal(0, Solution.Score(new int[] { 2, 3, 4, 6, 2 }));
        }

        [Fact]
        public static void ShouldValueTriplets()
        {
            Assert.Equal(400, Solution.Score(new int[] { 4, 4, 4, 3, 3 }));
        }

        [Fact]
        public static void ShouldValueMixedSets()
        {
            Assert.Equal(450, Solution.Score(new int[] { 2, 4, 4, 5, 4 }));
        }

        [Fact]
        public static void Test4()
        {
            Assert.Equal(1100, Solution.Score(new int[] { 1, 1, 1, 3, 1 }));
        }

        [Fact]
        public static void Test5()
        {
            Assert.Equal(250, Solution.Score(new int[] { 5, 1, 3, 4, 1 }));
        }

    }
}