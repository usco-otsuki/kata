
def solve():
    n = int(input())
    board = []
    d = [(0, 0), (1, 0), (-1, 0), (0, 1), (0, -1)]
    for i in range(n):
        board.append(list(input().strip()))
    
    for row in range(1, n-1):
        for column in range(1, n-1):
            if all(board[row + dx][column + dy] == "." for (dx, dy) in d):
                for (dx, dy) in d:
                    board[row + dx][column + dy] = "#"

    if all(board[row][column] == "#" for row in range(n) for column in range(n)):
        print("YES")
    else:
        print("NO")
    
solve()
