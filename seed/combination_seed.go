package seed

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/PandaX185/computer-alchemy-api/config"
	"github.com/PandaX185/computer-alchemy-api/models"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func SeedCombinations() error {
	driver := config.ConnectToDB()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	query := `
		MATCH (e1:Element {name: $firstElement})
		MATCH (e2:Element {name: $secondElement})
		MATCH (result:Element {name: $resultingElement})
		MERGE (e1)-[:COMBINES_WITH]->(e2)
		MERGE (e1)<-[:COMBINES_WITH]-(e2)
		MERGE (e1)-[:CREATES {with: $secondElement}]->(result)
	`

	for _, combination := range elementCombinations {
		log.Printf("Seeding combination: %s + %s = %s", combination.FirstElement, combination.SecondElement, combination.ResultingElement)
		_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, query, map[string]any{
				"firstElement":     combination.FirstElement,
				"secondElement":    combination.SecondElement,
				"resultingElement": combination.ResultingElement,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to execute query: %w", err)
			}
			return result.Consume(ctx)
		})
		if err != nil {
			return fmt.Errorf("failed to seed combination %s + %s = %s: %w",
				combination.FirstElement, combination.SecondElement, combination.ResultingElement, err)
		}
	}

	return nil
}

var elementCombinations = []models.Combination{
	{FirstElement: "Bit", SecondElement: "Bit", ResultingElement: "Byte"},
	{FirstElement: "Byte", SecondElement: "Byte", ResultingElement: "Data"},
	{FirstElement: "Logic Gate", SecondElement: "Logic Gate", ResultingElement: "Transistor"},
	{FirstElement: "Logic Gate", SecondElement: "Logic Gate", ResultingElement: "Flip Flop"},
	{FirstElement: "Logic Gate", SecondElement: "Transistor", ResultingElement: "Circuit"},
	{FirstElement: "Transistor", SecondElement: "Transistor", ResultingElement: "Chipset"},
	{FirstElement: "Flip Flop", SecondElement: "Flip Flop", ResultingElement: "Memory"},
	{FirstElement: "Chipset", SecondElement: "Chipset", ResultingElement: "CPU"},
	{FirstElement: "Chipset", SecondElement: "Memory", ResultingElement: "Motherboard"},
	{FirstElement: "CPU", SecondElement: "Motherboard", ResultingElement: "PC"},
	{FirstElement: "Data", SecondElement: "Data", ResultingElement: "File"},
	{FirstElement: "File", SecondElement: "File", ResultingElement: "Storage"},
	{FirstElement: "Data", SecondElement: "Circuit", ResultingElement: "Algorithm"},
	{FirstElement: "Algorithm", SecondElement: "Data", ResultingElement: "Programming"},
	{FirstElement: "Programming", SecondElement: "Data", ResultingElement: "Software"},
	{FirstElement: "Software", SecondElement: "PC", ResultingElement: "Operating System"},
}
