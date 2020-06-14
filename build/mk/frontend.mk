frontend-docker-run:		## Run the docker container
	sudo docker run -it -p 80:80 --name ui saferwall/ui

ui-build:		## Build the UI in docker
	sudo make docker-build IMG=ui DOCKER_FILE=ui/Dockerfile DOCKER_DIR=ui/

ui-release:		## build and release UI.
	sudo make docker-release IMG=ui VERSION=0.0.2 DOCKER_FILE=ui/Dockerfile DOCKER_DIR=ui/
