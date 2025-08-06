def square(number: int) -> int:
    """
    Calculate how many grains were on a given square.
    :param number: int - number referencing the square
    :return: int - number of grains on the square
    """
    if number > 64 or number < 1:
        # when the square value is not in the acceptable range
        raise ValueError("square must be between 1 and 64")
    else:
        # when the square value is in the acceptable range
        return 2 ** (number - 1)


def total() -> int:
    """
    Calculate the total number of grains on the chessboard.
    :return: int - total number of grains on the chessboard
    """
    total = 0
    for i in range(1, 65):
        total += square(i)
    return total
