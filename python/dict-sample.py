dict = {
    "brand": "Ford",
    "model": "Focus",
    "year": 2011
}

print(dict)

dict["year"] = 2020
print(dict["year"])

for k in dict:
    print("Key " + k)

for value in dict.values():
    print("Value " + str(value))