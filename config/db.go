package config

import (
	"context"
	"log"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var driverInstance neo4j.DriverWithContext

func ConnectToDB() neo4j.DriverWithContext {
	if driverInstance == nil || driverInstance.VerifyConnectivity(context.Background()) != nil {
		driver, err := neo4j.NewDriverWithContext(
			os.Getenv("NEO4J_URI"),
			neo4j.BasicAuth(os.Getenv("NEO4J_USER"), os.Getenv("NEO4J_PASSWORD"), ""),
		)

		if err != nil {
			log.Fatal(err)
		}

		driverInstance = driver
	}

	return driverInstance
}
