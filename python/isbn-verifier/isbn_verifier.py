def is_valid(isbn):
    sum = 0
    no_dash_isbn = isbn.lower().replace("-","")
    if len(no_dash_isbn) != 10:
        return False
    for i,v in enumerate(no_dash_isbn):
        if v == "x" and i == 9:
            v= 10
        elif not v.isdigit():
            return False
        sum += int(v)*(10-i)
    return sum % 11 == 0


