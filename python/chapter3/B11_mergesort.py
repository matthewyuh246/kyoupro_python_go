from typing import List

def merge_sort(numbers: List[int]) -> List[int]:
    if len(numbers) <= 1:
        return numbers
    
    center = len(numbers) // 2
    left = numbers[:center]
    right = numbers[center:]

    merge_sort(left)
    merge_sort(right)

    i = j = k = 0
    while i < len(left) and j < len(right):
        if left[i] <= right[j]:
            numbers[k] = left[i]
            i += 1
        else:
            numbers[k] = right[j]
            j += 1
        k += 1

    while i < len(left):
        numbers[k] = left[i]
        i += 1
        k += 1
    
    while j < len(right):
        numbers[k] = right[j]
        j += 1
        k += 1

    return numbers

def binary_search_left(numbers: List[int], value: int) -> int:
    left, right = 0, len(numbers) - 1
    while left <= right:
        mid = (left + right) // 2
        if numbers[mid] < value:
            left = mid + 1
        else:
            right = mid - 1
    return left


N = int(input())
A = list(map(int, input().split()))
Q = int(input())
X = [ None ] * Q
for i in range(Q):
    X[i] = int(input())

numbers = merge_sort(A)

for i in range(Q):
    pos1 = binary_search_left(numbers, X[i])
    print(pos1)
