from collections import deque

Q = int(input())
queries = [ input().split() for i in range(Q) ]

S = deque()
for q in queries:
    if q[0] == "1":
        S.append(q[1])
    if q[0] == "2":
        print(S[-1])
    if q[0] == "3":
        S.pop()