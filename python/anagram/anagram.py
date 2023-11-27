def find_anagrams(word, candidates):
    w_l = word.lower()
    L=[]
    for candidate in candidates:
        c_l = candidate.lower()
        if w_l == c_l:
            continue
        elif sorted(w_l) == sorted(c_l):
            L.append(candidate)
    return L

