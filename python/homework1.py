import os
from flask import Flask
app = Flask(__name__)

@app.route("/conf/env")
def healthchecker():
    f = os.popen('printenv')
    return f.read()

@app.route("/env/<name>/<value>")
def create_env_var(name, value):
    with open(os.path.expanduser("~/.bashrc"), "a") as outfile:
        outfile.write("export "+name+"="+value+"\n")
    return "OK"

if __name__ == '__main__':
    app.run(debug='true')