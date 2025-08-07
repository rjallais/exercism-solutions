def isTriangle(sides):
    a, b, c = sorted(sides)
    return a > 0 and a + b > c


def equilateral(sides):
    if not isTriangle(sides):
        return False
    return len(set(sides)) == 1


def isosceles(sides):
    if not isTriangle(sides):
        return False
    return len(set(sides)) < 3


def scalene(sides):
    if not isTriangle(sides):
        return False
    return len(set(sides)) == 3
