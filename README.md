
## About
kafka-protobuf-console-consumer reads protobuf encoded messages from kafka topic and prints its decoded JSON to console.

## Installation
For OS X download the compiled executable [kafka-protobuf-console-consumer](./kafka-protobuf-console-consumer)  

For other OS, please clone the repo and Build from source section
## Usage
```
kafka-protobuf-console-consumer --help
usage: kafka-protobuf-console-consumer [<flags>]

Flags:
      --help                     Show context-sensitive help (also try --help-long and --help-man).
  -b, --broker-list=localhost:9092 ...
                                 List of brokers to connect
  -t, --topic=TOPIC              Topic name
      --proto-dir=PROTO-DIR ...  foo/dir1 bar/dir2
      --file=FILE                Proto file name
      --message=MESSAGE          Proto message name
      --from-beginning           Read from beginning
      --pretty                   Intent output
      --with-separator           Adds separator between messages. Useful with --pretty
```

## Build from source
``` sh
cd $GOPATH/src
git clone 
glide install
GO111MODULE=on go build -o ./kafka-protobuf-console-consumer main.go
./kafka-protobuf-console-consumer -t topic_name -b broker1:port broker2:port \
--proto-dir /proto-folder-path proto --file sample.proto \
--message sample_package.SampleMessage
```
