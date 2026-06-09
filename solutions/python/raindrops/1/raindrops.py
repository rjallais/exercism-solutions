def convert(number):
    result = ""
    hash_map = {3: "Pling", 5: "Plang", 7: "Plong"}
    for key in hash_map:
        if number % key == 0:
            result += hash_map[key]
    if not result:
        result = f"{number}"

    return result
