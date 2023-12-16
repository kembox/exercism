import re
def is_paired(input_string):
    s , ok = re.sub(r"[^{}\[\]\(\)]","",input_string),1
    while ok:
        s, ok = re.subn(r"{}|\[\]|\(\)","",s)
    return not s 