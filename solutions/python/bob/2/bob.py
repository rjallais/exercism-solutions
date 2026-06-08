"""
This module implements Bob's response logic.

Bob responds to messages according to these rules:
- Silence input: "Fine. Be that way!"
- Yelled question (all uppercase + ?): "Calm down, I know what I'm doing!"
- Yelling (all uppercase): "Whoa, chill out!"
- Question (ends with ?): "Sure."
- Everything else: "Whatever."
"""

def response(hey_bob):
    """
    Respond to Bob according to the rules.

    Args:
        hey_bob (str): The message Bob receives

    Returns:
        str: Bob's response to the message
    """

    phrase = hey_bob.strip()
    if len(phrase) == 0:
        return "Fine. Be that way!"
    if phrase.isupper() and phrase[-1] == "?":
        return "Calm down, I know what I'm doing!"
    if phrase.isupper():
        return "Whoa, chill out!"
    if phrase[-1] == "?":
        return "Sure."
    return "Whatever."
