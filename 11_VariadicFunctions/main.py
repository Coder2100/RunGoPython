"""
In Python you can accept varying types with somefunction(*args)
 but this is not possible with Go. You can however,
 make the type an interface thus being able to get much more rich type structs.
"""

def average(*numbers):
    return sum(numbers)/ len(numbers)


print(average(1,3,4,6,7))