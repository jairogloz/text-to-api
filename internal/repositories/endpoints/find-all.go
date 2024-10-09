package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text-to-api/internal/domain"
)

// FindAll returns all the endpoints for a given user.
//
// Currently, it reads the endpoints from a JSON file.
// Todo: Implement the logic to read the endpoints from the database.
func (r repository) FindAll(ctx context.Context, userID string) ([]domain.Endpoint, error) {

	// read json file
	endpoints, err := readJSONFile("./examples/endpoint-definition-expenses.json")
	if err != nil {
		return nil, fmt.Errorf("could not read JSON file: %w", err)
	}

	return endpoints, nil
}

// Function to read and unmarshal the JSON file
func readJSONFile(filePath string) ([]domain.Endpoint, error) {
	// Open the JSON file
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer jsonFile.Close()

	// Read the file contents
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	// Unmarshal the byte data into the struct
	var fileContent struct {
		Endpoints []domain.Endpoint `json:"endpoints"`
	}
	err = json.Unmarshal(byteValue, &fileContent)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	return fileContent.Endpoints, nil
}
