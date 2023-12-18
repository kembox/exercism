def translate(text):
    return " ".join([stupid_pig(x) for x in text.split()])

def stupid_pig(text)->str:

    #Special case
    if text[:2] in ["xr","yt"]:
        return text + "ay"
    if text[:2] == "qu":
        return text[2:] + "qu" + "ay"

    if text[0] not in "ueoai" and text[1:3] == "qu":
        return text[3:] + text[0:3] + "ay"

    #Find the first vowel position
    i=-1
    for c in text:
        if c in "ueoai":
            i = text.index(c)
            break

    if i == 0:
        return text + "ay"
    elif i < 0:
        if "y" in text and 0 < text.index("y"):
            return text[text.index("y"):] + text[:text.index("y")] + "ay"
    else:
        return text[i:] + text[:i] + "ay"



