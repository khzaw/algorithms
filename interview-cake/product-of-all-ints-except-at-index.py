from functools import reduce
import operator


def prod(iterable):
    return reduce(operator.mul, iterable, 1)


def get_products_of_all_ints_except_at_index_naive(l):
    # Do not use division in your solution
    # O(n^2) time and O(n^2) space
    result = []
    for i in l:
        x = l[:]
        x.remove(i)
        result.append(prod(x))
    return result


def get_products_of_all_ints_except_at_index(l):
    products_of_ints_except_index = [None] * len(l)

    product = 1
    for i in range(len(l)):
        products_of_ints_except_index[i] = product
        product *= l[i]

    product = 1
    for i in range(len(l) - 1, -1, -1):
        products_of_ints_except_index[i] *= product
        product *= l[i]

    return products_of_ints_except_index

assert(get_products_of_all_ints_except_at_index([1, 7, 3, 4]) == [84, 12, 28, 21])
assert(get_products_of_all_ints_except_at_index([0, 7, 3, 4]) == [84, 0, 0, 0])
assert(get_products_of_all_ints_except_at_index([1, 2, 6, 5, 9]) == [540, 270, 90, 108, 60])
