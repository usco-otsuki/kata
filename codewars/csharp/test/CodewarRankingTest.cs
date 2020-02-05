using Xunit;

using CodewarRanking;

public class CodewarRaningTest
{
    [Fact]
    public void Test1()
    {
        var user = new User();
        Assert.Equal(-8, user.rank);
        user.incProgress(-8);
        Assert.Equal(3, user.progress);
        user.incProgress(-7);
        Assert.Equal(13, user.progress);
        user.incProgress(-5);
        Assert.Equal(3, user.progress);
        Assert.Equal(-7, user.rank);
        user.incProgress(1);
        Assert.Equal(93, user.progress);
        Assert.Equal(-3, user.rank);
        // user.incProgress(8);
    }

    [Fact]
    public void Test2()
    {
        var user = new User();
        var path = new[]
        {
          new { rank = 1, expRank = -2, expProg = 40 },
          new { rank = 1, expRank = -2, expProg = 80 },
          new { rank = 1, expRank = -1, expProg = 20 },
          new { rank = 1, expRank = -1, expProg = 30 },
          new { rank = 1, expRank = -1, expProg = 40 },
          new { rank = 2, expRank = -1, expProg = 80 },
          new { rank = 2, expRank = 1,  expProg =20 },
          new { rank = -1,expRank =  1, expProg = 21 },
          new { rank = 3, expRank = 1,  expProg =61 },
          new { rank = 8, expRank = 6,  expProg =51 },
          new { rank = 8, expRank = 6,  expProg =91 },
          new { rank = 8, expRank = 7,  expProg =31 },
          new { rank = 8, expRank = 7,  expProg =41 },
          new { rank = 8, expRank = 7,  expProg =51 },
          new { rank = 8, expRank = 7,  expProg =61 },
          new { rank = 8, expRank = 7,  expProg =71 },
          new { rank = 8, expRank = 7,  expProg =81 },
          new { rank = 8, expRank = 7,  expProg =91 },
          new { rank = 8, expRank = 8,  expProg =0 },
          new { rank = 8, expRank = 8,  expProg =0 },
      };

        // Go through path specified and assert rank at each step
        foreach (var element in path)
        {
            user.incProgress(element.rank);
            // System.Console.WriteLine($"{user.progress} {user.rank}");
            Assert.Equal(element.expProg, user.progress);
            Assert.Equal(element.expRank, user.rank);
        }
    }
}