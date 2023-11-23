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

color_tolerances = {
    "grey" : 0.05,
    "violet" : 0.1,
    "blue" : 0.25,
    "green" : 0.5,
    "brown" : 1,
    "red" : 2,
    "gold" : 5,
    "silver" : 10,
}

units = {"giga":1000000000,"mega":1000000,"kilo":1000}

def resistor_label(colors):
    if len(colors) == 1:
        return str(color_codes[colors[0]]) + " ohms"
    digits = ""
    *numbers,multiplier,tolerance = colors
    for x in numbers:
        digits += str(color_codes[x])

    myval = int(str(int(digits)) + color_codes[multiplier] * "0")

    if myval > 1000000000:
        unit = "giga"
    elif myval >= 1000000:
        unit = "mega"
    elif myval >= 1000:
        unit = "kilo" 
    else:
        unit = ""


    try:
       myval = int(myval) / units[unit]
    except KeyError:
        pass


    return f'{myval:g}' + " " + str(unit) + "ohms ±" + str(color_tolerances[tolerance]) +"%"
