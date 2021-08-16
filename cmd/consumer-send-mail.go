package cmd

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/smtp"

	kafka "github.com/segmentio/kafka-go"
	env "mail.notification.com/pkg"
)

type Mail struct {
	from     string
	password string
	to       []string
	smtpHost string
	smtpPort string
	message  []byte
}

func SendMail() {
	env := env.GetEnv()
	consumer([]string{env.Kafka.Url}, env.Kafka.Topic)
}

func consumer(broker []string, topic string) {
	k := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     broker,
		Topic:       topic,
		GroupID:     "3",
		StartOffset: kafka.LastOffset,
	})

	for {
		msg, err := k.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Kafka reader has failed, err: %v", err)
		}
		send(msg.Value)
	}
}

func send(body []byte) {

	env := env.GetEnv()
	var buffer bytes.Buffer
	buffer.Write([]byte(fmt.Sprintf("Card ID %v have been created", body)))
	mail := &Mail{
		from:     env.Mail.From,
		password: env.Mail.Password,
		to:       env.Mail.To,
		smtpHost: env.Mail.Smtphost,
		smtpPort: env.Mail.Smtpport,
		message:  buffer.Bytes(),
	}

	auth := smtp.PlainAuth("", mail.from, mail.password, mail.smtpHost)
	err := smtp.SendMail(mail.smtpHost+":"+mail.smtpPort, auth, mail.from, mail.to, mail.message)
	if err != nil {
		fmt.Println(err)
		return
	}
}
