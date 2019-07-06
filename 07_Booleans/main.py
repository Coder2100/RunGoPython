"""Go doesn't have a quick way to evaluate if something is "truthy".
 In Python, for example, you can use an if statement on any type and
 most types have a way of automatically converting to True or False.
  For example you can do:
"""

x = 1
if x:
    print("Yes")
y = []
if y:
    print("This is not going to be outputed")
#Go does not have a similr approach
#Need to do explicitly for every data type.see main.go

print(not True)#False
print(False or True)#False
print(True or False)#True
print(True and True)#True
print(False and True)#False