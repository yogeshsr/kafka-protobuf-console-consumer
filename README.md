
## About
kafka-protobuf-console-consumer reads protobuf encoded messages from a kafka topic and prints its decoded JSON to console.

This utility can be useful when you want to quickly look into the protobuf messages of a topic or dlq. It works over the proto source file and don't need you to compile it using protoc.

## Installation
Current version 2.1.0

For other than OS X, please clone the repo and follow Build from source section

For OS X
```
brew tap yogeshsr/homebrew-tap
brew install kafka-protobuf-console-consumer
```  

## Usage
``` sh
$ kafka-protobuf-console-consumer -t topic_name -b broker1:port broker2:port \
--proto-dir /proto-folder-path proto --file sample.proto \
--message sample_package.SampleMessage

$ kafka-protobuf-console-consumer --help
usage: kafka-protobuf-console-consumer [<flags>]

Flags:
  -v, --version                  Version
  -d, --debug                    Enable Sarama logs
  -b, --broker-list=localhost:9092 ...
                                 List of brokers to connect
  -c, --consumer-group=CONSUMER-GROUP
                                 Consumer group to use
  -t, --topic=TOPIC              Topic name
      --proto-dir=PROTO-DIR ...  /foo/dir1 /bar/dir2 (add all dirs used by imports)
      --file=FILE                will be baz/a.proto that's in /foo/dir1/baz/a.proto
      --message=MESSAGE          Proto message name
      --from-beginning           Read from beginning
      --pretty                   Format output
      --with-separator           Adds separator between messages. Useful with --pretty
```

## Build from source
``` sh
$ cd $GOPATH/src
$ git clone https://github.com/yogeshsr/kafka-protobuf-console-consumer.git
$ glide install
$ GO111MODULE=on go build -o ./kafka-protobuf-console-consumer main.go
```

---
Please add a star if you like this project.
