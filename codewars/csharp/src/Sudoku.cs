using System.Collections.Generic;

public class Sudoku
{
    public static bool ValidateSolution(int[][] board)
    {
        var target = new HashSet<int>() { 1, 2, 3, 4, 5, 6, 7, 8, 9 };
        for (int o1 = 0; o1 < 9; o1 += 3)
        {
            for (int o2 = 0; o2 < 9; o2 += 3)
            {
                var numbers = new HashSet<int>();
                for (int i = 0; i < 3; i++)
                {
                    for (int j = 0; j < 3; j++)
                    {
                        numbers.Add(board[o1 + i][o2 + j]);
                    }
                }

                if (!numbers.SetEquals(target))
                {
                    return false;
                }
            }
        }
        for (int i = 0; i < 9; i++)
        {
            var numbers = new HashSet<int>(board[i]);
            if (!numbers.SetEquals(target))
            {
                return false;
            }
        }

        for (int i = 0; i < 9; i++)
        {
            var numbers = new HashSet<int>();
            for (int j = 0; j < 9; j++)
            {
                numbers.Add(board[j][i]);
            }

            if (!numbers.SetEquals(target))
            {
                return false;
            }
        }
        return true;
    }
}