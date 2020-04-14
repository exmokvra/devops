# Docker
Always look for images ready to use. Amazon and Google have many

## Useful Docker commands
* Use Ansible just when creating and provisioning the machine. After that, remove Ansible

<code>
DOCKER RUN redis (will launch a docker with redis image)
Access the Redis of initiated redis: redis-cli -h <ip found with "docker ps">
DOCKER PS - lists all running containers
DOCKER INSPECT <id> - brings information about the container (grab CONTAINER_ID through "docker ps"
DOCKER RUN -rm --name redis redis 
* rm parameter - kills the container image when it's cancelled
DOCKER STOP <id> - stops container execution
DOCKER START <id> - starts a container execution
DOCKER EXEC -it <id> BASH - SSH to the container
DOCKER RM <name> - << removes stopped container
DOCKER RUN --network=host --name jenkins -p 8282:8080 p- 50000:50000 -v /home/sesteheim/devops jenkins
* Parameter --network=host << Good for the network when having to use a VPN as example
* Parameter -p << maps container ports to the PC ports
* Parameter -v << maps container volumes to the PC volumes
DOCKER BUILD -t <image name: usually <company>/<image name> <path>
* Used examlple: DOCKER BUILD -t guilherme/java8 . (do not forget the last dot)
DOCKER STATS << shows the hardware consumption of each container
</code>

### Docker Compose:
<code>DOCKER COMPOSE UP -d</code>