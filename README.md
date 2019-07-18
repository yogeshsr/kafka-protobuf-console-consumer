
## About
kafka-protobuf-console-consumer reads protobuf encoded messages from a kafka topic and prints its decoded JSON to console.

This utility can be useful when you want to quickly look into the protobuf messages of a topic or dlq. It works over the proto source file and don't need you to compile it using protoc.

## Installation
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
      --help                     Show context-sensitive help (also try --help-long and --help-man).
  -v, --version                  Version
  -d, --debug                    Enable Sarama logs
  -b, --broker-list=localhost:9092 ...
                                 List of brokers to connect
  -t, --topic=TOPIC              Topic name
      --proto-dir=PROTO-DIR ...  /path_to_sample_proto
      --file=FILE                Proto file name (sample.proto)
      --message=MESSAGE          Proto message name (sample_package.SampleMessage)
      --from-beginning           Read from beginning
      --pretty                   Intent output
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
