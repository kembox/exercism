def append(list1, list2):
    return list1 + list2

def concat(lists):
    return [x for list in lists for x in list ]


def filter(function, list):
    return [x for x in list if function(x)]

def length(list):
    count = 0
    for i in list:
        count += 1
    return count


def map(function, list):
    return [function(x) for x in list]


def foldl(function, list, initial):
    if not list:
        return initial

    list = [initial] + list
    while len(list) > 1:
        list = [function(list[0],list[1])] + list[2:]
    return list[0]

def foldr(function, list, initial):
    if not list:
        return initial

    list = list + [initial] 
    while len(list) > 1:
        list = list[:-2] + [function(list[-1],list[-2])]
    return list[0]


def reverse(list):
    return list[::-1]