import re
def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError("Strands must be of equal length.")
    elif len(strand_a) == 0:
        return 0
    if not re.match(r"^[A-Z]+$",strand_a+strand_b):
        raise ValueError("strands must contain upper case characters only")

    count = 0
    for i in range(len(strand_a)):
        if strand_a[i] != strand_b[i]:
                count += 1
    return count
