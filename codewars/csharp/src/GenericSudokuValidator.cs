using System;
using System.Collections.Generic;
using System.Linq;

namespace GenericSudokuValidator
{
    public class Sudoku
    {
        private int[][] sudokuData;
        public Sudoku(int[][] sudokuData)
        {
            this.sudokuData = sudokuData;
        }

        public bool IsValid()
        {
            int n = sudokuData.Count();
            int sqrtn = (int)Math.Sqrt(n);
            var target = new HashSet<int>(Enumerable.Range(1, n));
            for (int o1 = 0; o1 < n; o1 += sqrtn)
            {
                for (int o2 = 0; o2 < n; o2 += sqrtn)
                {
                    var numbers = new HashSet<int>();
                    for (int i = 0; i < sqrtn; i++)
                    {
                        for (int j = 0; j < sqrtn; j++)
                        {
                            numbers.Add(sudokuData[o1 + i][o2 + j]);
                        }
                    }

                    if (!numbers.SetEquals(target))
                    {
                        return false;
                    }
                }
            }
            for (int i = 0; i < n; i++)
            {
                var numbers = new HashSet<int>(sudokuData[i]);
                if (!numbers.SetEquals(target))
                {
                    return false;
                }
            }

            for (int i = 0; i < n; i++)
            {
                var numbers = new HashSet<int>();
                for (int j = 0; j < n; j++)
                {
                    numbers.Add(sudokuData[j][i]);
                }

                if (!numbers.SetEquals(target))
                {
                    return false;
                }
            }
            return true;
        }
    }
}