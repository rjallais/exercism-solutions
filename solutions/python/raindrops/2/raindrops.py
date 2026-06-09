"""Module providing a function that converts numbers into raindrop sounds."""

def convert(number):
    """Function to convert numbers into raindrop sounds."""
    result = ""
    hash_map = {3: "Pling", 5: "Plang", 7: "Plong"}
    for key, value in hash_map.items():
        if number % key == 0:
            result += value
    if not result:
        result = f"{number}"

    return result
