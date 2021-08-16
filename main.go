package main

import (
	cmd "mail.notification.com/cmd"
)

func main() {
	// Consume Kafka message and send a notification mail
	cmd.SendMail()

	//Consume Kafka Topic using consumer group
	cmd.ConsumeGroup()
}
