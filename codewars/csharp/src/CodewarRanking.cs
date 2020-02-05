// TODO: create the User class/object
// it must support rank, progress, and the incProgress(int rank) method

using System;

namespace CodewarRanking
{
    public class User
    {

        private int _progress;
        private int _rank;

        public int rank
        {
            get
            {
                if (_rank >= 0)
                {
                    return _rank + 1;
                }
                return _rank;
            }
            set
            {
                _rank = value;
            }
        }

        public int progress
        {
            get
            {
                if (rank >= 8) {
                    return 0;
                }
                return _progress;
            }
            set
            {
                _progress = value;
            }
        }

        public User()
        {
            this._rank = -8;
            this._progress = 0;
        }

        public void incProgress(int rank)
        {
            if (rank > 8 || rank < -8 || rank == 0) {
                throw new ArgumentException();
            }
            if (this.rank > 7)
            {
                return;
            }
            if (rank > 0)
            {
                rank--;
            }
            var diff = rank - this._rank;
            if (diff == 0)
            {
                progress += 3;
            }
            else if (diff > 0)
            {
                progress += 10 * diff * diff;
            }
            else if (diff == -1)
            {
                progress += 1;
            }
            System.Console.WriteLine("=========");
            System.Console.WriteLine(progress + " " + this.rank);
            while (progress >= 100 && this.rank <= 7)
            {
                this.rank = this._rank + 1;
                progress -= 100;
                System.Console.WriteLine("reduced " + progress + " " + this.rank);
            }
        }
    }
}