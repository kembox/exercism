plain = "abcdefghijklmnopqrstuvwxyz"
cipher = "zyxwvutsrqponmlkjihgfedcba"
import re

def encode(plain_text):
    plain_text = re.sub(r'[^a-zA-Z0-9]',"",plain_text).lower()
    print(f"normalized {plain_text}")
    s = ""
    for i in range(len(plain_text)):
        if plain_text[i].isnumeric():
            s += plain_text[i]
        else:
            pos = plain.index(plain_text[i])
            s += cipher[pos]

    return " ".join([s[i:i+5] for i in range(0,len(s),5) if i < len(s)])


def decode(ciphered_text):
    no_space = ciphered_text.replace(" ","")
    s = ""
    for i in range(len(no_space)):
        if no_space[i].isnumeric():
            s += no_space[i]
        else:
            pos = cipher.index(no_space[i])
            s += plain[pos]
    return s
