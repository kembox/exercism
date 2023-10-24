def is_isogram(string):
    letters_only = [s.lower() for s in string if s != " " and s != "-"]
    return len(set(letters_only)) == len(letters_only)