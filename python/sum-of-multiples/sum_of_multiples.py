def sum_of_multiples(limit, multiples):
    scores = set()
    for i in multiples:
        if i:
            for j in range(i,limit,i):
                scores.add(j)
    return sum(scores)
