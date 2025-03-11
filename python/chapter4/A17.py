N = int(input())
A = list(map(int, input().split()))
B = list(map(int, input().split()))

dp = [ None ] * (N+1)
dp[1] = 0
dp[2] = A[0]
for i in range(3, N+1):
    dp[i] = min(dp[i-1]+A[i-2], dp[i-2]+B[i-3])

Answer = []
Place = N
while True:
    Answer.append(Place)
    if Place == 1:
        break

    if dp[Place-1] + A[Place-2] == dp[Place]:
        Place = Place - 1
    else:
        Place = Place - 2
Answer.reverse()

Answer2 = [str(i) for i in Answer]
print(len(Answer))
print(" ".join(Answer2))