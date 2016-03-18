def quicksort(alist):
    return quicksort_helper(alist, 0, len(alist) - 1)

def quicksort_helper(alist, low, high):
    if low < high:
        pivot = partition(alist, low, high)
        quicksort_helper(alist, low, pivot-1)
        quicksort_helper(alist, pivot+1, high)
    return alist

def partition(alist, low, high):
    pivot = alist[low]

    leftmark = low + 1
    rightmark = high


    done = False
    while not done:
        while leftmark <= rightmark and alist[leftmark] <= pivot:
            leftmark = leftmark + 1

        while alist[rightmark]  >= pivot and rightmark >= leftmark:
            rightmark = rightmark - 1

        if rightmark < leftmark:
            done = True
        else:
            temp = alist[leftmark]
            alist[leftmark] = alist[rightmark]
            alist[rightmark] = temp

    temp = alist[low]
    alist[low] = alist[rightmark]
    alist[rightmark] = temp

    return rightmark


assert quicksort([3,4,2,1,5]) == [1,2,3,4,5]