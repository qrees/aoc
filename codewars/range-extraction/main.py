def solution(int_range):
    pairs_list = []
    current = [int_range[0], int_range[0]]
    for item in int_range[1:]:
        print(item, current)
        if item == current[1] + 1:
            current[1] = item
        else:
            pairs_list.append(current)
            current = [item, item]
    pairs_list.append(current)
    
    str_list = []
    for pair in pairs_list:
        if pair[0] == pair[1] - 1:
            str_list.append(f"{pair[0]},{pair[1]}")
        elif pair[0] == pair[1]:
            str_list.append(str(pair[0]))
        else:
            str_list.append(f"{pair[0]}-{pair[1]}")
    return ",".join(str_list)


print(solution([-6,-3,-2,-1,0,1,3,4,5,7,8,9,10,11,14,15,17,18,19,20]))
