"""Python doesn't have the concept of pointers. 
Go does. But with Go you can send an array or a map into a function, 
have it modified there without being returned and it gets changed.

"""

def upone(mutable, index):
    mutable[index] = mutable[index].upper()

list_ = ["a", "b", "c"]
upone(list_, 1)

print(list_)
dict_ = {"a": "anders", "b":"bengt"}
upone(dict_, "b")

print(dict_)#