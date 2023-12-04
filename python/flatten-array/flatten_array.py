def flatten(iterable):
    L=[]
    for i in iterable:
        if is_iterable(i):
            L.extend(flatten(i))
        else:
            L.append(i)
    return [x for x in L if x is not None]

def is_iterable(myobject):
    try:
        _ = iter(myobject)
    except TypeError:
        return False
    return True