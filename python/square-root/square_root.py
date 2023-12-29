def square_root(number):
    for i in range(1,number+1):
        if i*i == number:
            return i
        elif i*i > number:
            break


