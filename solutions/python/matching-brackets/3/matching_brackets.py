"""Check whether brackets in a string are correctly balanced and paired."""

def is_paired(input_string):
    """Return whether all brackets in the input string are balanced.

    The function processes parentheses, square brackets, and curly braces
    using a stack. Non-bracket characters are ignored.

    Args:
        input_string: The string to examine.

    Returns:
        True if every opening bracket is closed in the correct order;
        otherwise, False.
    """

    pairing = []
    for char in input_string:
        if char in "([{":
            pairing.append(char)
        elif char == ")":
            if not (pairing and pairing.pop()) == "(":
                return False
        elif char == "]":
            if not (pairing and pairing.pop()) == "[":
                return False
        elif char == "}":
            if not (pairing and pairing.pop()) == "{":
                return False

    return len(pairing) == 0
