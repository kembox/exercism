def find(search_list, value):
    sorted_list = sorted(search_list)
    if not sorted_list:
        raise ValueError("value not in array")

    low = 0
    high= len(sorted_list) - 1
    while low <= high:
        mid =low + (high-low) // 2 #floor divison
        print(mid)
        if sorted_list[mid] == value:
            return mid
        elif value < sorted_list[mid]:
            high = mid - 1
        else:
            low = mid + 1
    raise ValueError("value not in array")