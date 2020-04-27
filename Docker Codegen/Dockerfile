FROM nanoservice/protobuf

MAINTAINER Fergus In London <fergus@fergus.london>
ARG GITHUB_URL="https://api.github.com/repos/go-swagger/go-swagger/releases/latest"

# Install make, the go tooling, and any build dependencies
RUN apk add --update make curl jq bash git make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

# protoc actually requires language specific generators; so we're
#  opting for `protoc-gen-go`.
RUN go get -u github.com/golang/protobuf/protoc-gen-go

# Download Swagger Binary and configure for execution in the $PATH
RUN curl -o swagger -L `curl -s $GITHUB_URL | jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url'`
RUN mv swagger /usr/local/bin/swagger
RUN chmod +x /usr/local/bin/swagger

# Remove build specific packages
RUN apk del curl jq
ENTRYPOINT ["make"]
