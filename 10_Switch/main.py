def userInput():
    return int(input())

number = userInput()
if number == 1:
    print("Sabelo")
elif number ==2:
    print("Ntobeko")
elif number == 3:
    print("Monwabisi")
else:
    print("I have no idea what is %d" %number)

    #python3 main.py

# Alternative solution
number2 = int(input())
db = {1:"Sabelo", 2:"Ntobeko",3:"Monwabisi",4:"Zukile"}
print(db.get(number2,"I have no idea what is %d" %number))

#python3 main.py