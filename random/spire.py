def f(l, n):
    # the largest difference between the differences of the indices between the two same numbers within a subset of a given list.
    result = 0
    i = 0
    j = 0
    for i in range(n):
        for j in range(n):
            print "l[i], l[j]: ({li}, {lj})".format(li=l[i], lj=l[j])
            if(l[i] == l[j]):
                if(abs(i - j) > result):
                    result = abs(i - j)
    return result

def g(A, n):
    max_diff = 0
    for i in range(A):
        
