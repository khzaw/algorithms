def maximum_sub_array_sum(l):
    """
    Find the contiguous subarrays of numbers which has the largest sum.
    l is the list of numbers.

    Bi = The maximum subarray sum ending at position i
    Ai = The element at position i

    Bi+1 = max(Ai+1, Ai+1 + Bi)

    max(B1, B2, B3, ... , Bi+1)
    """

    max_ending = max_so_far = l[0]
    for num in l[1:]:
        max_ending = max(num, num + max_ending)
        max_so_far = max(max_so_far, max_ending)
    return max_so_far
