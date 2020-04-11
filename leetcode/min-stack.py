class Node:
    def __init__(self, value, min_value):
        self.value = value
        self.min_value = min_value


class MinStack:

    def __init__(self):
        self.stack = []
        
    def push(self, x: int) -> None:
        min_value = x
        if len(self.stack) > 0:
            min_value = min(self.stack[-1].min_value, x)

        self.stack.append(Node(x, min_value))

    def pop(self) -> None:
        self.stack = self.stack[:-1]

    def top(self) -> int:
        return self.stack[-1].value

    def getMin(self) -> int:
        return self.stack[-1].min_value


if __name__ == "__main__":
    minStack = MinStack();
    minStack.push(-2);
    minStack.push(0);
    minStack.push(-3);
    assert minStack.getMin() == -3;
    minStack.pop();
    assert minStack.top() == 0;
    assert minStack.getMin() == -2;
