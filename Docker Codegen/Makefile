WORKSPACE_DIRECTORY = "/go/src/workspace"
UTILITY_CONTAINER_TAG = "codegen-demo"
CONTAINER_MAKE_TARGET = "generator"

SWAGGER_SPEC = "swagger.json"
SWAGGER_OUTPUT_DIR = "./generated"

PROTOBUF_INPUT = "./example.proto"
PROTOBUF_OUTPUT_DIR = "./generated/messages"

# Ensure Docker Container uses the UID and GID of the current user; preventing
#  permissions issues
USER_ID := $(shell id -u)
GROUP_ID := $(shell id -g)

# Run our image with the current directory mounted to a workspace directory on the
#  container, and then execute our make target.
execute:
	@docker build -t ${UTILITY_CONTAINER_TAG} . > /dev/null
	@docker run -t -u ${USER_ID}:${GROUP_ID} \
		-v `pwd`:${WORKSPACE_DIRECTORY} \
		-w ${WORKSPACE_DIRECTORY} ${UTILITY_CONTAINER_TAG} ${CONTAINER_MAKE_TARGET}

generate-protobuf:
	@protoc --go_out=${PROTOBUF_OUTPUT_DIR} ${PROTOBUF_INPUT}

generate-swagger:
	@swagger generate model --spec=${SWAGGER_SPEC} -t ${SWAGGER_OUTPUT_DIR}

generator: generate-protobuf generate-swagger