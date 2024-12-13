package downloader

import (
	"encoding/json"
	"os"
	"testing"
)



func TestSaveNamesToFile(t *testing.T) {
	var testImagesNamesFile = "testImageNames.json"

	names := map[string]Name{
		"person1": {FileName: "person1.jpg", ImgUrl: "http://example.com/person1"},
		"person2": {FileName: "person2.jpg", ImgUrl: "http://example.com/person2"},
	}

	defer os.Remove(testImagesNamesFile)

	err := StoreFileName(names, testImagesNamesFile)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	data, err := os.ReadFile(testImagesNamesFile)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var savedNames map[string]Name
	err = json.Unmarshal(data, &savedNames)
	if err != nil {
		t.Fatalf("expected no error during unmarshalling, got %v", err)
	}

	// Check if the saved data matches the original data
	for key, name := range names {
		savedName, exists := savedNames[key]
		if !exists {
			t.Errorf("expected key %s to be present in saved data", key)
		}
		if savedName != name {
			t.Errorf("expected %v, got %v", name, savedName)
		}
	}
}
