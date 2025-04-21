N, X = map(int, input().split())
A = list(map(int, input().split()))
Answer = False

for i in range(N):
    if A[i] == X:
        Answer = True

if Answer == True:
    print("Yes")
else:
    print("No")