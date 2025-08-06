def steps(number):
    """
    Calculate the number of steps it takes to reach 1 using the Collatz conjecture.
    :param number: int - the starting number
    :return: int - the number of steps it takes to reach 1
    """
    if number < 1:
        # example when argument is zero or a negative integer
        raise ValueError("Only positive integers are allowed")
    elif number == 1:
        return 0
    else:
        if number % 2 == 0:
            # example when argument is even
            return steps(number // 2) + 1
        else:
            # example when argument is odd
            return steps(3 * number + 1) + 1
