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

func SeedElements() error {
	driver := config.ConnectToDB()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	query := `
		MATCH ()-[r]-()
		DELETE r
	`
	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, query, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to delete relationships: %w", err)
		}
		return nil, nil
	})
	if err != nil {
		return fmt.Errorf("failed to delete relationships: %w", err)
	}

	query = `
		MATCH (e:Element)
		DELETE e
	`
	_, err = session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(ctx, query, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to delete nodes: %w", err)
		}
		return nil, nil
	})
	if err != nil {
		return fmt.Errorf("failed to delete nodes: %w", err)
	}

	query = `
		MERGE (e:Element {name: $name})
		SET e.image = $image,
			e.description = $description
	`

	for i, element := range elements {
		log.Printf("Seeding element %d: %s", i+1, element.Name)
		_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, query, map[string]any{
				"name":        element.Name,
				"image":       element.Image,
				"description": element.Description,
			})
			if err != nil {
				return nil, fmt.Errorf("failed to execute query: %w", err)
			}
			return result.Consume(ctx)
		})
		if err != nil {
			return fmt.Errorf("failed to seed element %s: %w", element.Name, err)
		}
	}

	return nil
}

var elements = []models.Element{
	{Name: "Bit", Image: "bit.png", Description: "A bit is the smallest unit of digital information, representing either 0 or 1."},
	{Name: "Byte", Image: "byte.png", Description: "A byte is a unit of digital information consisting of 8 bits, capable of representing 256 different values."},
	{Name: "Logic Gate", Image: "logic-gate.png", Description: "Logic gates are fundamental electronic components that perform basic boolean operations (AND, OR, NOT, etc.) on binary inputs."},
	{Name: "Flip Flop", Image: "flip-flop.png", Description: "A flip flop is a digital circuit that can store one bit of information and maintain its state until actively changed."},
	{Name: "Transistor", Image: "transistor.png", Description: "A transistor is a semiconductor device used to amplify or switch electronic signals, forming the basic building block of modern electronics."},
	{Name: "Data", Image: "data.png", Description: "Data is any information processed or stored by a computer, including numbers, text, images, and other forms of information."},
	{Name: "Circuit", Image: "circuit.png", Description: "A circuit is an interconnected network of electronic components that processes, stores, or transmits electrical signals."},
	{Name: "Chipset", Image: "chipset.png", Description: "A chipset is a group of integrated circuits that manage data flow between the processor, memory, and peripherals."},
	{Name: "Motherboard", Image: "motherboard.png", Description: "A motherboard is the primary circuit board that connects and facilitates communication between all computer components."},
	{Name: "Algorithm", Image: "algorithm.png", Description: "An algorithm is a step-by-step procedure or formula for solving a problem or accomplishing a specific task."},
	{Name: "PC", Image: "pc.png", Description: "A PC (Personal Computer) is a multi-purpose computer designed for individual use, capable of running various software applications."},
	{Name: "Programming", Image: "programming.png", Description: "Programming is the process of creating sets of instructions (code) that tell a computer how to perform specific tasks."},
	{Name: "Software", Image: "software.png", Description: "Software is a collection of programs, data, and instructions that tell a computer system how to function and perform specific tasks."},
	{Name: "Operating System", Image: "operating-system.png", Description: "An operating system is core software that manages hardware resources and provides services for computer programs."},
	{Name: "Memory", Image: "memory.png", Description: "Memory is hardware that stores data and instructions for immediate or long-term use by a computer system."},
	{Name: "CPU", Image: "cpu.png", Description: "A CPU (Central Processing Unit) is the primary processor that performs most calculations and controls other computer components."},
	{Name: "File", Image: "file.png", Description: "A file is a collection of data stored on a computer's hard drive or other storage media, organized in a specific format."},
	{Name: "Storage", Image: "storage.png", Description: "Storage refers to non-volatile devices and media that retain data even when power is removed, like hard drives and SSDs."},
}
