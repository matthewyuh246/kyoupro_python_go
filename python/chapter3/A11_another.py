from typing import List, NewType

IndexNum = NewType('IndexNum', int)

def binary_search(numbers: List[int], value: int) -> IndexNum:
    def _binary_search(numbers: List[int], value: int, left: IndexNum, right: IndexNum) -> IndexNum:
        if left > right:
            return -1
        
        mid = (left + right) // 2
        if numbers[mid] == value:
            return mid
        elif numbers[mid] < value:
            return _binary_search(numbers, value, mid + 1, right)
        else:
            return _binary_search(numbers, value, left, mid - 1)
        
    return _binary_search(numbers, value, 0, len(numbers))

N, X = map(int, input().split())
A = list(map(int, input().split()))

answer = binary_search(A, X)
print(answer + 1)