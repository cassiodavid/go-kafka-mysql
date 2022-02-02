package main

import (
	"database/sql"
	"encoding/json"
	kafika "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gokafkamysql/infra/kafka"
	"github.com/gokafkamysql/infra/repository"
	usercase2 "github.com/gokafkamysql/usercase"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/exemplodb")
	if err != nil {
		log.Fatalln("sql erro conecção -- > ", err)
		panic(err.Error())
	}

	productRepository := repository.ProductMySqlRepository{Db: db}
	usercase := usercase2.CreateProduct{Repository: productRepository}

	var msgChan = make(chan *kafika.Message)
	configMapConsumer := &kafika.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "appgo",
	}
	topics := []string{"product"}
	consumer := kafka.NewConsumer2(configMapConsumer, topics)

	go consumer.Consume(msgChan)

	log.Println("Esperando Mensagens do Kafka ")
	for msg := range msgChan {
		var input usercase2.CreateProductInputDto
		err := json.Unmarshal(msg.Value, &input)
		if err != nil {
			log.Println(err)
		}
		output, err := usercase.Execute(input)
		if err != nil {
			log.Println("Error no Execute", err)
		} else {
			log.Println(output)
		}
	}
}

//{"name":"farinha", "quantidade":"4"}
