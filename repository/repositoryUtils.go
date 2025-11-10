package repository

import (
	"fmt"
	"regexp"

	"github.com/Informasjonsforvaltning/fdk-harvest-admin/model"
)

func isValidID(id string) bool {
	// Define a regular expression for a valid ID (e.g., alphanumeric)
	var validID = regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	return validID.MatchString(id)
}

func isValidPublisherID(publisherID string) bool {
	// PublisherID should follow the same format as ID
	// Allow alphanumeric characters, hyphens, and potentially underscores
	var validPublisherID = regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
	return validPublisherID.MatchString(publisherID) && publisherID != ""
}

func validateDataSource(dataSource *model.DataSource) error {
	// Validate using the model's Validate method
	if err := dataSource.Validate(); err != nil {
		return fmt.Errorf("data source validation failed: %w", err)
	}

	// Validate required fields are not empty
	if dataSource.URL == "" {
		return fmt.Errorf("url is required")
	}

	if dataSource.PublisherID == "" {
		return fmt.Errorf("publisherId is required")
	}

	if dataSource.DataSourceType == "" {
		return fmt.Errorf("dataSourceType is required")
	}

	if dataSource.DataType == "" {
		return fmt.Errorf("dataType is required")
	}

	return nil
}
