def equilateral(sides):
    return sides[0] == sides[1] == sides[2] and sides[0] > 0

def isosceles(sides):
    return sides[0] in sides[1:] or sides[1] == sides[2]

def scalene(sides):
    return (not isosceles(sides)) and ( not equilateral(sides))