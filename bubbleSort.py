def bSort(a):
    for i in range(len(a)):
        for j in range(len(a)-1, i, -1):
            if a[j] < a[j-1]:
                a[j], a[j-1] = a[j-1], a[j]

    return a


array = [9,8,7,6,5,4,3,2,1]
bSort(array)
print(array)