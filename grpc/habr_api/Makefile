HABR_PATH=./habrapi/habr_api.proto

all:
	protoc -I habrapi/ ${HABR_PATH} \
		--go_out=plugins=grpc:habrapi

.PHONY: all hello