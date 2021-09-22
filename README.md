# graphql_demo
 A simple demo for using gqlgen https://github.com/99designs/gqlgen lib with golang
 
 ## Setup of the demo:

Must have golang installed version >= 12.0.0

make file consists of 3 steps: generate, build, run you can run all of them

```make all```

Start the http server on port 9090:

```make run```


# RUN BY DOCKER

```
 docker build -t go_graphql .
 
 docker run -p 9090:9090 go_graphql
 
 ```
