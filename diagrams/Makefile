DIR := ${CURDIR}

server-podman:
	podman --cgroup-manager cgroupfs run -it --rm -p 8080:8080 -v $(DIR):/usr/local/structurizr docker.io/structurizr/lite

server-docker:
	docker run -it --rm -p 8080:8080 -v $(DIR):/usr/local/structurizr docker.io/structurizr/lite
