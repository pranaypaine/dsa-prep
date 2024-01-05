# For the given array of integers, count even and odd elements.

# Examples: 

# Input: 
# int arr[5] = {2, 3, 4, 5, 6}
# Output: 
# Number of even elements = 3    
# Number of odd elements = 2  

# Input:
# int arr[5] = {22, 32, 42, 52, 62}
# Output: 
# Number of even elements = 5  
# Number of odd elements = 0


def check_odd_even(input_array):
    even_count = 0
    odd_count = 0
    
    for item in input_array:
        if (item & 1 == 1):
            odd_count += 1
        else:
            even_count += 1
            
    print("Number of even elements = ", even_count) 
    print("Number of odd elements = ", odd_count)
    
    

arr = [2, 3, 4, 5, 6]
check_odd_even(arr)