import json

dict_obj = {
    "name": "Guilherme",
    "car": "focus",
    "city": "Tampa"
}

j = json.dumps(dict_obj)

print(j)