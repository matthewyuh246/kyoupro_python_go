from collections import deque

N, X = map(int, input().split())
X -= 1
A = list(input())

Q = deque([X])
A[X] = '@'
while Q:
    pos = Q.popleft()
    if pos - 1 >= 0 and A[pos - 1] == '.':
        A[pos - 1] = '@'
        Q.append(pos - 1)
    if pos + 1 < N and A[pos + 1] == '.':
        A[pos + 1] = '@'
        Q.append(pos + 1)

print("".join(A))
