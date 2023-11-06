import re
def answer(question):
    to_translate = {
        "What is" : "",
        "minus": "-",
        "plus" : "+",
        "multiplied by" : "*",
        "divided by" : "/",
        "?" : ""
    }

    for k,v in to_translate.items():
        question = question.replace(k,v)
    q = re.sub(' {2,}',' ',question)
    #dedup spaces

    if q.replace(" ","").isdigit():
        return int(q.replace(" ",""))
    #Just a number

    p1= re.compile(r"[a-z]+")
    #Unknown operation

    p2 = re.compile(r"(-)?\d+\s(\+|\-|\*|\/)\s(\-)?\d+")
    #1 correct expression 

    p3= re.compile(r"\d\s\d")
    #2 digits in a row
    #To cover multiple expressions

    if p1.search(q):
        raise ValueError("unknown operation")

    if p3.search(q):
        raise ValueError("syntax error")

    if p2.search(q):
        #Handle requirement about expression order
        while p2.search(q):
            left_most_exp = p2.search(q).group()
            #Match the 1st expression

            val=str(int(eval(left_most_exp)))
            #Calculate value

            q=q.replace(left_most_exp,val)
            #Replace expression by value to continue

        return eval(q)
    else:
        raise ValueError("syntax error")



