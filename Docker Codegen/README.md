# Docker Utility Image for Go Code Generation (Swagger and Protocol Buffers) ![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/fergusinlondon/swagger-protobuf-go-generation.svg)

This Docker image contains all the tools required to generate code for Go from Swagger API and Protocol Buffer definitions. Based upon an `alpine` base image, this image includes:

- `go`
- `make`
- `go-swagger`
- `protoc`
- `protoc-gen-go`

## Demonstration

This repository was originally intended for a [blog post](https://fergus.london/code/automating-code-generation-with-docker/) I wrote to illustrate the usefulness of Docker utility images, and their application for running code generators. 

So for a demonstration of this image, simply clone the [github repository](https://github.com/FergusInLondon/docker-codegen-demonstration) and run `make`. Usage should be clear from the associated `Makefile` contained in the repository, which also contains user/group ID flags for execution of the container.

## Improvements

As this grew out of a quick blog post, there's a few corners that I cut - considering I have a few applications for this image myself, I'll likely straighten any issues out when I get time. But notably:

- Build times could be decreased, the base image installs *and then subsequently removes* things that we actually need.
- There's no version control for `go-swagger`; this could be fixed via a simple `ARG` in the `Dockerfile`.

In short, it makes more sense to rework the `Dockerfile` to use an official Go alpine image as the base, and install `go-swagger` and `protoc` as static binaries.

### Example (w/ included `Makefile`

```
➜  mkdir demo; cd demo
➜  git clone git@github.com:FergusInLondon/docker-codegen-demonstration.git
➜  make
2019/04/11 05:47:43 validating spec /go/src/workspace/swagger.json
2019/04/11 05:47:44 preprocessing spec with option:  minimal flattening
2019/04/11 05:47:44 building a plan for generation
2019/04/11 05:47:44 planning definitions
2019/04/11 05:47:44 planning operations
2019/04/11 05:47:44 grouping operations into packages
2019/04/11 05:47:44 planning meta data and facades
2019/04/11 05:47:44 rendering 6 models
2019/04/11 05:47:44 rendering 1 templates for model ApiResponse
2019/04/11 05:47:44 name field ApiResponse
2019/04/11 05:47:44 package field models
2019/04/11 05:47:44 creating generated file "api_response.go" in "generated/models" as definition
2019/04/11 05:47:44 executed template asset:model
2019/04/11 05:47:44 rendering 1 templates for model Category
2019/04/11 05:47:44 name field Category
2019/04/11 05:47:44 package field models
2019/04/11 05:47:44 creating generated file "category.go" in "generated/models" as definition
2019/04/11 05:47:44 executed template asset:model
2019/04/11 05:47:44 rendering 1 templates for model Order
2019/04/11 05:47:44 name field Order
2019/04/11 05:47:44 package field models
2019/04/11 05:47:44 creating generated file "order.go" in "generated/models" as definition
2019/04/11 05:47:44 executed template asset:model
2019/04/11 05:47:44 rendering 1 templates for model Pet
2019/04/11 05:47:44 name field Pet
2019/04/11 05:47:44 package field models
2019/04/11 05:47:44 creating generated file "pet.go" in "generated/models" as definition
2019/04/11 05:47:44 executed template asset:model
2019/04/11 05:47:44 rendering 1 templates for model Tag
2019/04/11 05:47:44 name field Tag
2019/04/11 05:47:44 package field models
2019/04/11 05:47:44 creating generated file "tag.go" in "generated/models" as definition
2019/04/11 05:47:44 executed template asset:model
2019/04/11 05:47:44 rendering 1 templates for model User
2019/04/11 05:47:44 name field User
2019/04/11 05:47:44 package field models
2019/04/11 05:47:44 creating generated file "user.go" in "generated/models" as definition
2019/04/11 05:47:44 executed template asset:model
2019/04/11 05:47:44 Generation completed!

For this generation to compile you need to have some packages in your GOPATH:

	* github.com/go-openapi/runtime
	* github.com/jessevdk/go-flags

You can get these now with: go get -u -f generated/...
➜  ls generated/messages 
example.pb.go
➜  ls generated/models 
api_response.go  category.go  order.go  pet.go  tag.go  user.go
➜ 
```
