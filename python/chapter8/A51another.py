from typing import Any

class Stack(object):

    def __init__(self) -> None:
        self.stack = []

    def push(self, data) -> None:
        self.stack.append(data)

    def pop(self) -> Any:
        if self.stack:
            return self.stack.pop()

    def print(self) -> None:
        print(self.stack[-1])

if __name__ == '__main__':
    Q = int(input())
    queries = [ input().split() for i in range(Q)]

    stack = Stack()

    for q in queries:
        if q[0] == "1":
            stack.push(q[1])
        if q[0] == "2":
            stack.print()
        if q[0] == "3":
            stack.pop()
