S = input()

Stack = []
for i in range(len(S)):
    if S[i] == '(':
        Stack.append(i + 1)
    if S[i] == ')':
        print(Stack.pop(), i + 1)