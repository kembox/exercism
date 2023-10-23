def is_pangram(sentence):
    alphabet=set([chr(65 + x) for x in range(0,26)])
    input_set=set("".join(sentence.split()).upper())
    return input_set.issuperset(alphabet)
