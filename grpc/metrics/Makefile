METRICS_PATH=./metrics/metrics.proto

all:
	protoc -I metrics/ ${METRICS_PATH} \
		--go_out=plugins=grpc:metrics

.PHONY: all