import json

j = json.loads('{"sao_leopoldo" : "1234", "porto_alegre": "1769"}')

print("Was founded in "+j['sao_leopoldo'])