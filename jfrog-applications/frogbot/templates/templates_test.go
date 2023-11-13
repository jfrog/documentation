package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v2"

	"github.com/jfrog/jfrog-client-go/utils"
	"github.com/jfrog/jfrog-client-go/utils/errorutils"
)

const (
	maxRetriesToDownloadSchema           = 5
	durationBetweenSchemaDownloadRetries = 10 * time.Second
)

func TestFrogbotSchema(t *testing.T) {
	// Download frogbot schema
	schemaLoader := downloadYamlSchema(t, "https://raw.githubusercontent.com/jfrog/frogbot/master/schema/frogbot-schema.json")

	// Validate config in the docs
	validateYamlSchema(t, schemaLoader, filepath.Join(".frogbot", "frogbot-config.yml"), "")
}

func TestGitHubActionsTemplates(t *testing.T) {
	// Download GitHub Actions schema
	schemaLoader := downloadYamlSchema(t, "https://json.schemastore.org/github-workflow.json")

	// Validate workflows in the "github-actions" directory
	validateYamlsInDirectory(t, "github-actions", schemaLoader)
}

// Download a Yaml schema.
// t      - Testing object
// schema - The schema file to download
func downloadYamlSchema(t *testing.T, url string) gojsonschema.JSONLoader {
	var response *http.Response
	var err error
	retryExecutor := utils.RetryExecutor{
		MaxRetries:               maxRetriesToDownloadSchema,
		RetriesIntervalMilliSecs: int(durationBetweenSchemaDownloadRetries.Milliseconds()),
		ErrorMessage:             "Failed to download schema",
		ExecutionHandler: func() (bool, error) {
			response, err = http.Get(url)
			if err != nil {
				return true, err
			}
			err = errorutils.CheckResponseStatus(response, http.StatusOK)
			return err != nil, err
		},
	}
	assert.NoError(t, retryExecutor.Execute())
	assert.Equal(t, http.StatusOK, response.StatusCode, response.Status)
	// Check server response and read schema bytes
	defer func() {
		assert.NoError(t, response.Body.Close())
	}()
	schemaBytes, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	return gojsonschema.NewBytesLoader(schemaBytes)
}

// Validate all yml files in the given directory against the input schema
// t            - Testing object
// schemaLoader - The schema to use in the validation
// path	        - Yaml directory path
func validateYamlsInDirectory(t *testing.T, path string, schemaLoader gojsonschema.JSONLoader) {
	err := filepath.Walk(path, func(schemaFilePath string, info os.FileInfo, err error) error {
		assert.NoError(t, err)
		if strings.HasSuffix(info.Name(), "yml") {
			validateYamlSchema(t, schemaLoader, schemaFilePath, "")
		}
		return nil
	})
	assert.NoError(t, err)
}

// Validate a Yaml file against the input Yaml schema
// t            - Testing object
// schemaLoader - The schema to use in the validation
// yamlFilePath - Yaml file path
// expectError  - Expected error or an empty string if error is not expected
func validateYamlSchema(t *testing.T, schemaLoader gojsonschema.JSONLoader, yamlFilePath, expectError string) {
	t.Run(filepath.Base(yamlFilePath), func(t *testing.T) {
		// Read yaml
		yamlFile, err := os.ReadFile(yamlFilePath)
		assert.NoError(t, err)

		// Unmarshal yaml to object
		var yamlObject interface{}
		err = yaml.Unmarshal(yamlFile, &yamlObject)
		assert.NoError(t, err)

		// Convert the Yaml config to JSON config to help the json parser validate it.
		// The reason we don't do the convert by as follows:
		// YAML -> Unmarshall -> Go Struct -> Marshal -> JSON
		// is because the config's struct includes only YAML annotations.
		jsonObject := convertYamlToJson(yamlObject)

		// Load and validate json with the schema
		documentLoader := gojsonschema.NewGoLoader(jsonObject)
		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		assert.NoError(t, err)
		if expectError != "" {
			assert.False(t, result.Valid())
			assert.Contains(t, result.Errors()[0].String(), expectError)
		} else {
			assert.True(t, result.Valid(), result.Errors())
		}
	})
}

// Recursively convert yaml interface to JSON interface
func convertYamlToJson(yamlValue interface{}) interface{} {
	switch yamlMapping := yamlValue.(type) {
	case map[interface{}]interface{}:
		jsonMapping := map[string]interface{}{}
		for key, value := range yamlMapping {
			if key == true {
				// "on" is considered a true value for the Yaml Unmarshaler. To work around it, we set the true to be "on".
				key = "on"
			}
			jsonMapping[fmt.Sprint(key)] = convertYamlToJson(value)
		}
		return jsonMapping
	case []interface{}:
		for i, value := range yamlMapping {
			yamlMapping[i] = convertYamlToJson(value)
		}
	}
	return yamlValue
}
