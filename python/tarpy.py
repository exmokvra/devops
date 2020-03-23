import tarfile
import os

files = os.listdir("/home/sesterheim/devops/python")

tar = tarfile.open("source.tar.gz", "w:gz")
for name in files:
    tar.add(name)
tar.close()