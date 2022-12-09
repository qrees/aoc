# def next_bigger(n):
#     str_n = str(n)
#     str_n = list(str_n[::-1])
#     letter = str_n[0]
#     for index, next_letter in enumerate(str_n[1:]):
#         index = index + 1
#         print(next_letter)
#         if next_letter < letter:
#             print(next_letter, "<", letter)
#             previous_letters = str_n[:index]
#             print("previous:", previous_letters)
#             for letter in sorted(previous_letters):
#                 if letter > next_letter:
#                     print("chosen:", letter)
#                     previous_letters.remove(letter)
#                     previous_letters.append(next_letter)
#                     result = str_n[index + 1:]
#                     result = result[::-1]
#                     result.append(letter)
#                     for letter in sorted(previous_letters):
#                         result.append(letter)
#                     return result
#         letter = next_letter
#     return -1

def next_bigger(n):
    str_n = list(str(n))
    letter = "0"
    for index, next_letter in list(enumerate(str_n))[::-1]:
        if next_letter < letter:
            previous_letters = str_n[index:]
            for letter in sorted(previous_letters):
                if letter > next_letter:
                    previous_letters.remove(letter)
                    result = str_n[0:index]
                    result.append(letter)
                    for letter in sorted(previous_letters):
                        result.append(letter)
                    return int("".join(result))
        letter = next_letter
    return -1


print(next_bigger(247775000))