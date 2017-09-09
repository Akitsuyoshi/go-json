# go-json

This simple api is written in golang.

### Prerequisites
1. You've already installed Docker in your local or VM
2. the knowledge of some docker to use


### the settting to use

1. Install this repository into your local

 `$ git clone https://github.com/Akitsuyoshi/go-json`

2. Build the Docker image

 `$ docker build -t go-json .`

3. Run the image

 `$ docker run -p 80:8081 go-json`

And visit localhost:80 in your browser, you can see the application run.


If you'd like to stop, follow the commands
```
// it shows the container ID, which is running
$ docker ps

// To stop the process with ID
$ docker stop CONTAINER_ID
```
