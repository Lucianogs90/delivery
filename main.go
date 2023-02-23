package main

import (
	/* "fmt" */
	"log"

	/* "github.com/codeedu/simulator/application/route" */
	"github.com/codeedu/simulator/infra/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	/* route := route.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()

	stringJson, _ := route.ExportJsonPositions()

	fmt.Println(stringJson[0]) */

  producer := kafka.NewKafkaProducer()
  kafka.Publish("Testando", "readtest", producer)

  for{
    _ = 1
  }
}
