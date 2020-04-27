# Multi-Stage Docker Builds for Golang
# ====================================

# Docker is awesome, and once you get your head around it - it's also pretty simple. You know what else is awesome though? Docker multi-stage builds. And guess what? They're pretty simple themselves..

# Essentially, you just write two `Dockerfile`s in one actual `Dockerfile`. Sound a bit odd? Here's a worked example..

# Stage 1: The Build Container
# ----------------------------
# The first stage is quite obvious: we need to define a container that is capable of actually building an executable from our source code.

# As we're using Go, the natural choice is to use the the official golang image - specifically the lightweight `alpine` based one which contains all the  Go build tools.
FROM golang:alpine

# The `$GOPATH` for this image is set to `/usr/src`, so we'll set our working directory to `/usr/src/app` and then copy our source code here. 
WORKDIR /usr/src/app
COPY . .

# Next up we'll build our executable, this is as simple as running `go build` - just as you would normally. We'll set our output executable to be called `MultiStageExample`: we'll need this later. 
RUN go build -v -o MultiStageExample


# Stage 2: The Distributable Container
#-------------------------------------
# Once we've built our executable, we need to create the actual image that we can distribute and deploy.

# We're opting to use `alpine` as a base image again, due to it's small size. This is more important for this stage, as this is the image we'll actually deploy/push.
FROM alpine
MAINTAINER Fergus In London <fergus@fergus.london>

# Set our working directory and... **bang, here's the magic**. By utilising `--from=0 [directory/file]` with the `COPY` directive, we're able to instruct Docker that we want to copy a file *from the initial (i.e 0) container*. This is genuinely all the magic there is to a multi-stage build using Docker! Pretty neat, huh? 
WORKDIR /var/app/example
COPY --from=0 /usr/src/app/MultiStageExample .

# Lastly - and just like any other Docker image - we ensure that our target file is actually executable (`chmod +x ..`), and we set it as the image entrypoint. Easy.
RUN chmod +x MultiStageExample
ENTRYPOINT ["/var/app/example/MultiStageExample"]
