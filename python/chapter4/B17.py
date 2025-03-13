N = int(input())
H = list(map(int, input().split()))

dp = [ None ] * N

dp[0] = 0
dp[1] = abs(H[0] - H[1])
for i in range(2, N):
    dp[i] = min(dp[i-1] + abs(H[i-1]-H[i]), dp[i-2] + abs(H[i-2]-H[i]))

Answer = []
Place = N-1
while True:
    Answer.append(Place + 1)
    if Place == 0:
        break

    if dp[Place-1] + abs(H[Place-1]-H[Place]) == dp[Place]:
        Place = Place - 1
    else:
        Place = Place - 2
Answer.reverse()

Answer2 = [str(i) for i in Answer]
print(len(Answer))
print(" ".join(Answer2))