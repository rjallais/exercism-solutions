"""Calculate a dart throw score from Cartesian coordinates."""

import math

def score(x, y):
    """Return the dart score for a point on the target.

    Args:
        x: The x-coordinate of the dart throw.
        y: The y-coordinate of the dart throw.

    Returns:
        The score based on the distance from the origin:
        10 for radius 1 or less, 5 for radius greater than 1 and up to 5,
        1 for radius greater than 5 and up to 10, and 0 otherwise.
    """

    radius = math.sqrt(x * x + y * y)

    if radius > 10:
        return 0
    if radius > 5:
        return 1
    if radius > 1:
        return 5
    return 10
