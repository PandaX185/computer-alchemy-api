package service

import (
	"context"
	"errors"
	"time"

	"github.com/PandaX185/computer-alchemy-api/config"
	"github.com/PandaX185/computer-alchemy-api/dto"
	"github.com/PandaX185/computer-alchemy-api/models"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CombineElements(firstElement, secondElement string) ([]*models.Element, error) {
	driver := config.ConnectToDB()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := `
		MATCH (e1:Element)
		WHERE toLower(e1.name) = toLower($firstElement)
		MATCH (e2:Element)
		WHERE toLower(e2.name) = toLower($secondElement)
		MATCH (e1)-[r1:CREATES]->(result:Element)
		MATCH (e2)-[r2:CREATES]->(result:Element)
		WHERE toLower(r1.with) = toLower($secondElement)
		RETURN result.name, result.image, result.description
	`

	result := []*models.Element{}
	_, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		records, err := tx.Run(ctx, query, map[string]any{
			"firstElement":  firstElement,
			"secondElement": secondElement,
		})
		if err != nil {
			return nil, err
		}
		for records.Next(ctx) {
			element := &models.Element{
				Name:        records.Record().Values[0].(string),
				Image:       records.Record().Values[1].(string),
				Description: records.Record().Values[2].(string),
			}
			result = append(result, element)
		}

		return nil, nil
	})

	if len(result) == 0 {
		err = errors.New("invalid combination")
	}

	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetAllCombinations() []*dto.CombinationResponse {
	combinations := []*dto.CombinationResponse{}

	driver := config.ConnectToDB()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := `
		MATCH (e1:Element)-[c:CREATES]->(result:Element)
		WHERE c.with IS NOT NULL
		MATCH (e2:Element {name: c.with})
		RETURN e1.name, e1.image, e1.description, e2.name, e2.image, e2.description, result.name, result.image, result.description
	`

	session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		records, err := tx.Run(ctx, query, map[string]any{})
		if err != nil {
			return nil, err
		}
		for records.Next(ctx) {
			vals := records.Record().Values
			combination := &dto.CombinationResponse{
				FirstElement:     &models.Element{Name: vals[0].(string), Image: vals[1].(string), Description: vals[2].(string)},
				SecondElement:    &models.Element{Name: vals[3].(string), Image: vals[4].(string), Description: vals[5].(string)},
				ResultingElement: &models.Element{Name: vals[6].(string), Image: vals[7].(string), Description: vals[8].(string)},
			}
			combinations = append(combinations, combination)
		}
		return nil, nil
	})

	return combinations
}

func GetAllElementCombinations(element string) []*dto.CombinationResponse {
	combinations := []*dto.CombinationResponse{}

	driver := config.ConnectToDB()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := `
		MATCH (e1:Element)-[c:CREATES]->(result:Element)
		WHERE c.with IS NOT NULL
		MATCH (e2:Element {name: c.with})
		WHERE toLower(e1.name) = toLower($element) OR toLower(e2.name) = toLower($element)
		RETURN e1.name, e1.image, e1.description, e2.name, e2.image, e2.description, result.name, result.image, result.description
	`

	session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		records, err := tx.Run(ctx, query, map[string]any{
			"element": element,
		})
		if err != nil {
			return nil, err
		}
		for records.Next(ctx) {
			vals := records.Record().Values
			combination := &dto.CombinationResponse{
				FirstElement:     &models.Element{Name: vals[0].(string), Image: vals[1].(string), Description: vals[2].(string)},
				SecondElement:    &models.Element{Name: vals[3].(string), Image: vals[4].(string), Description: vals[5].(string)},
				ResultingElement: &models.Element{Name: vals[6].(string), Image: vals[7].(string), Description: vals[8].(string)},
			}
			combinations = append(combinations, combination)
		}
		return nil, nil
	})

	return combinations
}

func GetAllResultCombinations(resultingElement string) []*dto.CombinationResponse {
	combinations := []*dto.CombinationResponse{}

	driver := config.ConnectToDB()
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	query := `
		MATCH (e1:Element)-[c:CREATES]->(result:Element)
		WHERE c.with IS NOT NULL
		MATCH (e2:Element {name: c.with})
		WHERE toLower(result.name) = toLower($resultingElement)
		RETURN e1.name, e1.image, e1.description, e2.name, e2.image, e2.description, result.name, result.image, result.description
	`

	session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		records, err := tx.Run(ctx, query, map[string]any{
			"resultingElement": resultingElement,
		})
		if err != nil {
			return nil, err
		}
		for records.Next(ctx) {
			vals := records.Record().Values
			combination := &dto.CombinationResponse{
				FirstElement:     &models.Element{Name: vals[0].(string), Image: vals[1].(string), Description: vals[2].(string)},
				SecondElement:    &models.Element{Name: vals[3].(string), Image: vals[4].(string), Description: vals[5].(string)},
				ResultingElement: &models.Element{Name: vals[6].(string), Image: vals[7].(string), Description: vals[8].(string)},
			}
			combinations = append(combinations, combination)
		}
		return nil, nil
	})

	return combinations
}
