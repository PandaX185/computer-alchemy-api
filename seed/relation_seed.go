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

func SeedRelations() error {
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
		MERGE (e2)-[:CREATES {with: $firstElement}]->(result)
	`

	for _, relation := range elementRelations {
		log.Printf("Seeding relation: %s + %s = %s", relation.FirstElement, relation.SecondElement, relation.ResultingElement)
		_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, query, map[string]any{
				"firstElement":     relation.FirstElement,
				"secondElement":    relation.SecondElement,
				"resultingElement": relation.ResultingElement,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to execute query: %w", err)
			}
			return result.Consume(ctx)
		})
		if err != nil {
			return fmt.Errorf("failed to seed relation %s + %s = %s: %w",
				relation.FirstElement, relation.SecondElement, relation.ResultingElement, err)
		}
	}

	return nil
}

var elementRelations = []models.Relation{
	{FirstElement: "Bit", SecondElement: "Bit", ResultingElement: "Byte"},
	{FirstElement: "Logic Gate", SecondElement: "Logic Gate", ResultingElement: "Transistor"},
	{FirstElement: "Logic Gate", SecondElement: "Logic Gate", ResultingElement: "Flip Flop"},
	{FirstElement: "Transistor", SecondElement: "Circuit", ResultingElement: "Chipset"},
	{FirstElement: "Chipset", SecondElement: "Circuit", ResultingElement: "Motherboard"},
	{FirstElement: "CPU", SecondElement: "Motherboard", ResultingElement: "PC"},
	{FirstElement: "RAM", SecondElement: "Storage", ResultingElement: "Memory"},
	{FirstElement: "Algorithm", SecondElement: "Data", ResultingElement: "Programming"},
	{FirstElement: "Programming", SecondElement: "Data", ResultingElement: "Software"},
	{FirstElement: "Software", SecondElement: "PC", ResultingElement: "Operating System"},
	{FirstElement: "Data", SecondElement: "Storage", ResultingElement: "Database"},
	{FirstElement: "Programming", SecondElement: "Web", ResultingElement: "API"},
	{FirstElement: "Server", SecondElement: "Cloud", ResultingElement: "Web"},
	{FirstElement: "Data", SecondElement: "Algorithm", ResultingElement: "Machine Learning"},
	{FirstElement: "Machine Learning", SecondElement: "Neural Network", ResultingElement: "Deep Learning"},
	{FirstElement: "Machine Learning", SecondElement: "Algorithm", ResultingElement: "AI"},
	{FirstElement: "Bit", SecondElement: "Bit", ResultingElement: "Binary Code"},
	{FirstElement: "Bit", SecondElement: "Circuit", ResultingElement: "Digital Signal"},
	{FirstElement: "Bit", SecondElement: "Register", ResultingElement: "Memory Buffer"},
	{FirstElement: "Bit", SecondElement: "Memory", ResultingElement: "Cache"},
	{FirstElement: "Logic Gate", SecondElement: "Circuit", ResultingElement: "ALU"},
	{FirstElement: "Logic Gate", SecondElement: "Register", ResultingElement: "Control Unit"},
	{FirstElement: "Logic Gate", SecondElement: "Binary Code", ResultingElement: "Instruction Set"},
	{FirstElement: "Logic Gate", SecondElement: "Data", ResultingElement: "Multiplexer"},
	{FirstElement: "Logic Gate", SecondElement: "Memory", ResultingElement: "Decoder"},
	{FirstElement: "Logic Gate", SecondElement: "ALU", ResultingElement: "CPU"},
}
