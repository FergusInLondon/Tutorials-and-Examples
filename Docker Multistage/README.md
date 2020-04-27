# The **simplest** "*Docker Multistage Build*" example ever.

An incredibly simple example of Docker's *Multistage Builds*.. in part because they are incredibly simple themselves.

Check out the [Docco documentation](https://fergusinlondon.github.io/Super-Simple-Docker-Multi-Stage/) of what's *actually* going on, and [the accompanying blog post](https://fergus.london/docker-multi-stage-builds-are-awesome/) which covers *why* this is probably something you should be doing, as well as a few of the pitfalls and their mitigations.

## Example

Simply `git clone ...` this repository and run `docker build .`: **bonus points if you don't have any of the Go toolkit installed**.

```
$ docker build .
Sending build context to Docker daemon  588.8kB
Step 1/10 : FROM golang:alpine
 ---> 52d894fca6d4
Step 2/10 : WORKDIR /usr/src/app
Removing intermediate container a01877a39214
 ---> 35a6f54f8056
Step 3/10 : COPY . .
 ---> 1d509aff8f31
Step 4/10 : RUN go build -v -o MultiStageExample
 ---> Running in 04b213adbb5c
_/usr/src/app
Removing intermediate container 04b213adbb5c
 ---> a36ad17a2bb2
Step 5/10 : FROM alpine
latest: Pulling from library/alpine
ff3a5c916c92: Already exists 
Digest: sha256:7df6db5aa61ae9480f52f0b3a06a140ab98d427f86d8d5de0bedab9b8df6b1c0
Status: Downloaded newer image for alpine:latest
 ---> 3fd9065eaf02
Step 6/10 : MAINTAINER Fergus In London <fergus@fergus.london>
 ---> Running in 6d48c333f627
Removing intermediate container 6d48c333f627
 ---> bdc26c2454d6
Step 7/10 : WORKDIR /var/app/example
Removing intermediate container daf9eac0d6ed
 ---> a33e047f8f90
Step 8/10 : COPY --from=0 /usr/src/app/MultiStageExample .
 ---> 04f49b7b8f00
Step 9/10 : RUN chmod +x MultiStageExample
 ---> Running in bf73cbeeefbf
Removing intermediate container bf73cbeeefbf
 ---> 7e2a32eb402d
Step 10/10 : ENTRYPOINT ["/var/app/example/MultiStageExample"]
 ---> Running in 574da5e1462b
Removing intermediate container 574da5e1462b
 ---> 86b2413a096c
Successfully built 86b2413a096c
$ docker run 86b2413a096c
 Hello World!
$
```

