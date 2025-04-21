def isPrime(N):
    LIMIT = int(N ** 0.5)
    for i in range(2, LIMIT + 1):
        if N % i == 0:
            return False
        return True
    
Q = int(input())
X = [ None ] * Q
for i in range(Q):
    X[i] = int(input())

for i in range(Q):
    Answer = isPrime(X[i])
    if Answer == True:
        print("Yes")
    else:
        print("No")
