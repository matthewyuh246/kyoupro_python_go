from typing import List

def bubble_sort(numbers: List[int]) -> List[int]:
    len_numbers = len(numbers)
    for i in range(len_numbers):
        for j in range(len_numbers-1-i):
            if numbers[j] > numbers[j+1]:
                numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
    return numbers

def binary_search_left(numbers: List[int], value: int) -> int:
    left, right = 0, len(numbers)
    while left < right:
        mid = (left + right) // 2
        if numbers[mid] < value:
            left = mid + 1
        else:
            right = mid
    return left

N = int(input())
A = list(map(int, input().split()))
Q = int(input())
X = [ None ] * Q
for i in range(Q):
    X[i] = int(input())

numbers = bubble_sort(A)

for i in range(Q):
    pos1 = binary_search_left(numbers, X[i])
    print(pos1)
