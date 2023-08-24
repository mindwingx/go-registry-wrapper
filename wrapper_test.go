package registrywrapper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInitRegistry(t *testing.T) {
	fileName := "config.yml"
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, fileName)
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	// instantiate
	v := New()

	// config file exists
	err = v.InitRegistry("yml", filePath)
	assert.Nil(t, err)

	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	// config not file exists

	err = v.InitRegistry("json", "nonexistent.json")
	assert.NotNil(t, err)
}

func TestParse(t *testing.T) {
	fileName := "config.yml"
	data := "test:\n  v1: k1\n  v2: k2"
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, fileName)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// instantiate
	v := New()

	// config file exists
	err = v.InitRegistry("yml", filePath)

	// Test with valid config file
	var parsedData map[string]interface{}
	err = v.Parse(&parsedData)
	assert.Nil(t, err)
	assert.Equal(t, "k1", parsedData["test"].(map[string]interface{})["v1"].(string))
	assert.Equal(t, "k2", parsedData["test"].(map[string]interface{})["v2"].(string))

	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
}

func TestValueOfSpecificKeyValues(t *testing.T) {
	fileName := "config.yml"
	data := "test:\n  v1: k1\n  v2: k2\nnew:\n  item: value\n  next: one"
	currentDir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", currentDir, fileName)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// instantiate
	v := New()

	// config file exists
	err = v.InitRegistry("yml", filePath)

	// Test with valid config file
	var parsedData map[string]interface{}

	val := v.ValueOf("new")
	err = val.Parse(&parsedData)
	assert.Nil(t, err)
	assert.Equal(t, "value", parsedData["item"].(string))
	assert.Equal(t, "one", parsedData["next"].(string))

	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
}
