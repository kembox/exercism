color_codes = {
        "black":  0,
        "brown":  1,
        "red":    2,
        "orange": 3,
        "yellow": 4,
        "green":  5,
        "blue":   6,
        "violet": 7,
        "grey":   8,
        "white":  9,
}

units = {"giga":1000000000,"mega":1000000,"kilo":1000}

def label(colors):
    digits = ""
    for x in colors[:2]:
        digits += str(color_codes[x])

    value = str(int(digits)) + color_codes[colors[2]] * "0"

    if int(value) > 1000000000:
        unit = "giga"
    elif int(value) >= 1000000:
        unit = "mega"
    elif int(value) >= 1000:
        unit = "kilo" 
    else:
        unit = ""


    try:
       value = str(int(int(value) / units[unit]))
    except KeyError:
        pass

    return str(value) + " " + str(unit) + "ohms"