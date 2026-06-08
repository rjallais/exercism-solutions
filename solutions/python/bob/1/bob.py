def response(hey_bob):
    phrase = hey_bob.strip()
    if len(phrase) == 0:
        return "Fine. Be that way!"
    elif phrase.isupper() and phrase[-1] is '?':
        return "Calm down, I know what I'm doing!"
    elif phrase.isupper():
        return "Whoa, chill out!"
    elif phrase[-1] is '?':
        return "Sure."
    else:
        return "Whatever."
