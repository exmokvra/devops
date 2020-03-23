import sys
from ruamel.yaml import YAML

string_yaml = """\
#example
name:
    # details
    family: Sesterheim
    given: Guilherme
"""

yaml = YAML()
code = yaml.load(string_yaml)
#code['name']['given'] = 'GUILHERME'

yaml.dump(code, sys.stdout)