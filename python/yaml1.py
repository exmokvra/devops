import yaml

with open("/home/sesterheim/devops/python/yaml_test.yaml", 'r') as stream:
    try:
        y = yaml.load(stream)
        print(y["tags"][0]["name"])
        print(y)
    except yaml.YAMLError as exc:
        print(exc)