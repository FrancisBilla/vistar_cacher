package downloader

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"vistar_cacher/namer"
)

type Name struct {
	FileName string `json:"fileName"`
	ImgUrl   string `json:"imgUrl"`
}

var imagesFolder = "./store/imageStore"
var imagesNamesFile = "./store/imageNames.json"

func ImageDownload(imgUrl string) error {
	//Get the response bytes from the url
	response, err := http.Get(imgUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	fileName := namer.GenerateName()

	err = imageSaver(fileName, response.Body)
	if err != nil {
		return fmt.Errorf("Error in saving image %v ", err)
	}

	names := make(map[string]Name)
	names[imgUrl] = Name{FileName: fileName, ImgUrl: imgUrl}

	err = StoreFileName(names, imagesNamesFile)
	if err != nil {
		return fmt.Errorf("Error in saving name: Func - StoreFileName %v ", err)
	}

	return nil
}

func imageSaver(fileName string, responseBody io.ReadCloser) error {
	if _, err := os.Stat(imagesFolder); os.IsNotExist(err) {
		err := os.MkdirAll(imagesFolder, os.ModePerm)
		if err != nil {
			return err
		}
	}
	filePath := filepath.Join(imagesFolder, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Copy the image content to the file
	_, err = io.Copy(file, responseBody)
	if err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}

	fmt.Printf("Image successfully saved to %s\n", filePath)
	return nil
}

func StoreFileName(names map[string]Name, filePath string) error {
	// Convert the map to JSON
	data, err := json.MarshalIndent(names, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}

	// Write JSON to file
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}
