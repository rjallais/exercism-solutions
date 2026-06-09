"""Translate words or phrases into Pig Latin."""

VOWELS = frozenset({"a", "e", "i", "o", "u"})
VOWEL_PREFIXES = frozenset({"xr", "yt"})


def translate_word(word: str) -> str:
    """Translate a single lowercase word into Pig Latin."""
    if word[:1] in VOWELS or word[:2] in VOWEL_PREFIXES:
        return word + "ay"

    consonant_cluster = ""

    for index, char in enumerate(word):
        if char not in VOWELS:
            if char == "y" and consonant_cluster:
                return word[index:] + consonant_cluster + "ay"
            consonant_cluster += char
        else:
            if word[index - 1] == "q" and char == "u":
                return word[index + 1:] + consonant_cluster + "uay"
            return word[index:] + consonant_cluster + "ay"

    return word + "ay"


def translate(text: str) -> str:
    """Translate a space-separated string into Pig Latin."""
    return " ".join(translate_word(word) for word in text.split())
