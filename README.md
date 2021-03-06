# go-list

go-list is the backend API project. 

### Tech Used

This Project use some technologies like:

* [Go lang] - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
* [Glide] - Vendor Package Management for Golang
* [Echo] - High performance, extensible, minimalist Go web framework
* [Swagger] - Swagger open source and pro tools have helped millions of API developers, teams, and organizations deliver great APIs.
* [Docker] - BUILD, SHIP, RUN


### Instal Dependencies with glide
```sh
$ glide up -u -s
```

### Run Project development mode
```sh
$ go run server.go
```

### Build Docker Project
```sh
$ docker build -t [image_name]:[version/tag] .
```

### Run Project using docker
```sh
$ docker run -p 1323:1323 --name=[container_name] -d  [image_name]
```

### call Example 
```request
$ example:
/api/list?venture=ZAP&type=RENTAL&offSet=0&limit=50&orderByPrice=desc
```

### To do
 - Include FileBeat for logs
 - Write Tests
 - Documentation

License
----

@jeisonkunii

### Thanks
mom


[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)


   [Go lang]: <https://golang.org/>
   [Glide]: <https://github.com/Masterminds/glide>
   [Swagger]: <https://swagger.io/>
   [Echo]: <https://echo.labstack.com>
   [Docker]: <https://www.docker.com/>