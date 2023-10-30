def rotate(text, key):
    ret = ""
    for x in text:
        if x.isalpha():
            n = ord(x)+ (key%26)
            base = 65 # ord('A')

            #Detect if it's a uppercase character and set base range accordingly
            if x != x.upper():
                base += 32

            #Handle wrapped around input
            if n <= (base + 25):
                ret += chr(n) 
            else:
                ret += chr(n-26)
                #ret += chr(base -1 + n - (base+25))
                #If number is out side of the a-z/A-Z range, rotate it to the beginning
        else:
            ret += x

    return ret

