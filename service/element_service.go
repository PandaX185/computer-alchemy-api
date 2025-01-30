package service

import (
	"context"
	"log"
	"time"

	"github.com/PandaX185/computer-alchemy-api/config"
	"github.com/PandaX185/computer-alchemy-api/models"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetAllElements() []models.Element {
	driver := config.ConnectToDB()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := `
		MATCH (e:Element)
		RETURN e.name, e.image, e.description
	`

	var result []models.Element
	_, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		records, err := tx.Run(ctx, query, nil)
		if err != nil {
			return nil, err
		}
		for records.Next(ctx) {
			var element models.Element
			element.Name = records.Record().Values[0].(string)
			element.Image = records.Record().Values[1].(string)
			element.Description = records.Record().Values[2].(string)
			result = append(result, element)
		}
		return result, nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return result
}

func GetElementByName(name string) models.Element {
	driver := config.ConnectToDB()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := `
		MATCH (e:Element)
		WHERE toLower(e.name) = toLower($name)
		RETURN e.name, e.image, e.description
	`

	var result models.Element
	_, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		records, err := tx.Run(ctx, query, map[string]any{"name": name})
		if err != nil {
			return nil, err
		}
		for records.Next(ctx) {
			result.Name = records.Record().Values[0].(string)
			result.Image = records.Record().Values[1].(string)
			result.Description = records.Record().Values[2].(string)
		}
		return result, nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return result
}

func GetCombinationResult(firstElement, secondElement string) []models.Element {
	driver := config.ConnectToDB()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := `
		MATCH (e1:Element {name: $firstElement})-[r:CREATES]->(result:Element)
		MATCH (e2:Element {name: $secondElement})-[r:CREATES]->(result:Element)
		WHERE r.with = $secondElement
		RETURN result.name, result.image, result.description
	`

	var result []models.Element
	_, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		records, err := tx.Run(ctx, query, map[string]any{
			"firstElement":  firstElement,
			"secondElement": secondElement,
		})
		if err != nil {
			return nil, err
		}
		for records.Next(ctx) {
			var element models.Element
			element.Name = records.Record().Values[0].(string)
			element.Image = records.Record().Values[1].(string)
			element.Description = records.Record().Values[2].(string)
			result = append(result, element)
		}
		return result, nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return result
}
