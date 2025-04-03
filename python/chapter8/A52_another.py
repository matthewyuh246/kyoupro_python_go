from typing import Any

class Queue(object):

    def __init__(self) -> None:
        self.queue = []

    def push(self, data) -> None:
        self.queue.append(data)
    
    def popleft(self) -> Any:
        if self.queue:
            return self.queue.pop(0)
    
    def print(self) -> None:
        print(self.queue[0])

if __name__ == '__main__':
    Q = int(input())
    queries = [ input().split() for i in range(Q)]

    queue = Queue()

    for q in queries:
        if q[0] == "1":
            queue.push(q[1])
        if q[0] == "2":
            queue.print()
        if q[0] == "3":
            queue.popleft()