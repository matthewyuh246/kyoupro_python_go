N = int(input())
H = list(map(int, input().split()))

dp = [ None ] * N

dp[0] = 0
dp[1] = abs(H[0] - H[1])
for i in range(2, N):
    dp[i] = min(dp[i-1] + abs(H[i-1]-H[i]), dp[i-2] + abs(H[i-2]-H[i]))

print(dp[N-1])