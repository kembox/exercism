# Possible sublist categories.
# Change the values as you see fit.
SUBLIST = 1 
SUPERLIST = 2 
EQUAL = 3 
UNEQUAL = 4 


def sublist(list_one, list_two):
    if list_one == list_two:
        return EQUAL
    if len(list_one) < len(list_two):
        if issub(list_one,list_two):
            return SUBLIST
        else:
            return UNEQUAL
    else:
        if issub(list_two,list_one):
            return SUPERLIST
        else:
            return UNEQUAL
        
def issub(l1,l2)->bool:
    is_sub = False
    for i in range(0,len(l2)-len(l1)+1):
        if l1 == l2[i:i+len(l1)]:
            is_sub = True
    return is_sub