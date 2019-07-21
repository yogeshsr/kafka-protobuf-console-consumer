
build:
	GO111MODULE=on go build -ldflags "-X main.versionInfo=`cat version.txt` -X main.versionDate=`date -u +%Y-%m-%d.%H:%M:%S`"  -o ./kafka-protobuf-console-consumer main.go


package:
		tar -czf kafka-protobuf-console-consumer-`cat version.txt`.tar.gz kafka-protobuf-console-consumer

