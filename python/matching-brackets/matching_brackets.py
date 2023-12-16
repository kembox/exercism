def is_paired(input_string):
    stack = []
    open = "({["
    close = ")}]"
    for c in input_string:
        if c in open:
            stack.append(c)
        if c in close:
            if len(stack) == 0 or open[close.index(c)] != stack.pop():
            #pop here is super cool
                return False
    return not stack