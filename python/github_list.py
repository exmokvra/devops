from github import Github

g = Github("a", "b")

user = g.get_user("exmokvra")
print(user.name)

for repo in g.get_user().get_repos():
    print(repo.name)