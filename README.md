## Introduction

This application aims to consume from a Kafka topic using [kafka-go](https://github.com/segmentio/kafka-go). For each received message, a mail is send to the specified mailbox.

## Prerequisite

Fill up the environmment variables to configure Kafka and the SMTP server. A template is located under /config/env/default.yaml

## Quickstart

Run the application

```
go run main.go
```
