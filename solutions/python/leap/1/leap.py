def leap_year(year: int) -> bool:
    """
    Determine whether a given year is a leap year.
    :param year: int - the year to check
    :return: bool - True if the year is a leap year, False otherwise
    """
    if year % 400 == 0:
        return True
    elif year % 100 == 0:
        return False
    elif year % 4 == 0:
        return True
    else:
        return False
