"""Translate words or phrases into Pig Latin."""

def translate(text: str) -> str:
    """Translate a space-separated string into Pig Latin.

    Args:
        text: A string containing one or more lowercase words separated by spaces.

    Returns:
        The Pig Latin translation of the input text.
    """
    fields = text.split()
    vowels = "aeiou"
    result = ""
    for f in fields:
        if f != fields[0]:
            result += " "
        if f[:1] in vowels or f[:2] in ("xr", "yt"):
            result += f + "ay"
            continue

        consonants = ""
        for k, v in enumerate(f):
            if v not in vowels:
                if v == "y" and consonants:
                    result += f[k:] + consonants + "ay"
                    break
                consonants += v
            else:
                if f[k-1] == "q" and v == "u":
                    result += f[k+1:] + consonants + "u" + "ay"
                    break
                result += f[k:] + consonants + "ay"
                break

    return result
