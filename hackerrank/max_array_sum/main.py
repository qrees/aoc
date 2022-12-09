# Complete the maxSubsetSum function below.

def maxSubsetSum(arr):
    arr_len = len(arr)
    dyn_arr = [0] * arr_len
    dyn_arr[arr_len - 1] = arr[arr_len - 1]
    if arr_len == 1:
        return dyn_arr[0]
    dyn_arr[arr_len - 2] = max(arr[arr_len - 2], arr[arr_len - 1])
    if arr_len == 2:
        return dyn_arr[0]
    for i in range(3, arr_len + 1):
        cur = arr[arr_len - i]
        max_so_far = dyn_arr[arr_len - i + 1]
        max_and_cur = cur + dyn_arr[arr_len - i + 2]
        if cur > max_so_far and cur > max_and_cur:
            max_value = cur
        elif max_so_far > max_and_cur:
            max_value = max_so_far
        else:
            max_value = max_and_cur
        dyn_arr[arr_len - i] = max_value
    return dyn_arr[0]



if __name__ == '__main__':
    import sys
    fptr = sys.stdout

    n = int(input())
    arr = list(map(int, input().rstrip().split()))

    res = maxSubsetSum(arr)

    fptr.write(str(res) + '\n')
