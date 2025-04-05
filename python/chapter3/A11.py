from typing import List, NewType

IndexNum = NewType('IndexNum', int)

def binary_search(numbers: List[int], value: int) -> IndexNum:
    left, right = 0, len(numbers) - 1
    while left <= right:
        mid = (left + right) // 2
        if numbers[mid] == value:
            return mid
        elif numbers[mid] < value:
            left = mid + 1
        else:
            right = mid - 1
    return -1

N, X = map(int, input().split())
A = list(map(int, input().split()))

answer = binary_search(A, X)
print(answer + 1)