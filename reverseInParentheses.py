def solution(input_string):
    output, _ = get_string_portion(input_string, 0, False)
    return output


def get_string_portion(input_string, start_index, reverse):
    portion = ""
    index = start_index
    while index < len(input_string):
        char = input_string[index]
        if char == ')':
            return portion, index
        if char == '(':
            new_portion, index = get_string_portion(input_string, index + 1, not reverse)
        else:
            new_portion = char
        if reverse:
            portion = new_portion + portion
        else:
            portion = portion + new_portion
        index += 1
    return portion, len(input_string)


print(solution("foo(bar(baz))blim"))
