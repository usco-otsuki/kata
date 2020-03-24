using System.Collections.Generic;

namespace SortBinaryTreeByLevel
{
    public class Node
    {
        public Node Left;
        public Node Right;
        public int Value;

        public Node(Node l, Node r, int v)
        {
            Left = l;
            Right = r;
            Value = v;
        }
    }
    public class Solution
    {
        public static List<int> TreeByLevels(Node root)
        {
            var sorted = new List<int>();
            var queue = new Queue<Node>();
            queue.Enqueue(root);
            while (queue.Count > 0)
            {
                var node = queue.Dequeue();
                if (node == null)
                {
                    continue;
                }
                sorted.Add(node.Value);
                queue.Enqueue(node.Left);
                queue.Enqueue(node.Right);
            }
            return sorted;
        }
    }
}
