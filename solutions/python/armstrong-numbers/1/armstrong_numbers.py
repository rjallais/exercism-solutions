def is_armstrong_number(number: int) -> bool:
    """
    Calculate whether a number is an Armstrong number.
    :param number: int - number to check
    :return: bool - True if the number is an Armstrong number, False otherwise
    """
    # convert the number to a string
    number_str = str(number)
    # calculate the length of the number
    number_length = len(number_str)
    # initialize the sum
    armstrong_sum = 0
    # loop through the number's digits
    for digit in number_str:
        # square the digit
        square = int(digit) ** number_length
        # add the square to the sum
        armstrong_sum += square
    # check if the sum is equal to the original number
    return armstrong_sum == number
