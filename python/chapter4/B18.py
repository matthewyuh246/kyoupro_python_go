import sys

N, S = map(int, input().split())
A = list(map(int, input().split()))

dp = [[ None ] * (S + 1) for i in range(N + 1)]
dp[0][0] = True
for i in range(1, S+1):
    dp[0][i] = False

for i in range(1, N+1):
    for j in range(0, S+1):
        if j < A[i-1]:
            if dp[i-1][j] == True:
                dp[i][j] = True
            else:
                dp[i][j] = False
        if j >= A[i-1]:
            if dp[i-1][j] == True or dp[i-1][j-A[i-1]] == True:
                dp[i][j] = True
            else:
                dp[i][j] = False

if dp[N][S] == False:
    print("-1")
    sys.exit(0)

Answer = []
Place = S
for i in reversed(range(1, N+1)):
    if dp[i-1][Place] == True:
        Place = Place - 0
    else:
        Place = Place - A[i-1]
        Answer.append(i)
Answer.reverse()

Answer2 = [str(i) for i in Answer]
print(len(Answer))
print(" ".join(Answer2))
