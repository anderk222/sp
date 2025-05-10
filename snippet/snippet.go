/*
Copyright © 2025 Anderson Macias <anderk222@gmail.com>
*/

package snippet

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Snippet struct {
	Content         string `json:"-"`
	Name            string
	Description     string
	FileContentName string
}

func SaveSnippets(fileName string, snippets map[string]Snippet) error {

	n, err := json.Marshal(snippets)
	/*
	   Copyright © 2025 NAME HERE <EMAIL ADDRESS>
	*/
	if err != nil {
		return err
	}

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(n)

	if err != nil {
		return err
	}

	return nil
}

func SaveSnippetContent(fileName string, snippet *Snippet) error {

	snippet.FileContentName = fileName

	// Crear el directorio si no existe
	dir := filepath.Dir(fileName)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	/*
	   Copyright © 2025 NAME HERE <EMAIL ADDRESS>
	*/
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := file.WriteString(snippet.Content); err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) (map[string]Snippet, error) {

	data, err := os.ReadFile(filename)

	if err != nil {
		return map[string]Snippet{}, err
	}

	var snippets map[string]Snippet

	if err = json.Unmarshal(data, &snippets); err != nil {
		return map[string]Snippet{}, err
	}

	return snippets, nil
}

func RetrieveSnippet(fileName, name string) (Snippet, error) {

	snippets, err := ReadItems(fileName)

	if err != nil {
		return Snippet{}, err
	}

	snippet, exists := snippets[name]
	if !exists {
		return Snippet{}, fmt.Errorf("snippet not found")
	}

	return snippet, nil
}

func InjectSnippet(fileSnippets, targetFile string, name string) error {
	snippet, err := RetrieveSnippet(fileSnippets, name)
	if err != nil {
		fmt.Println("Error retrieving snippet:", err)
		return err
	}

	snippetContent, err := RetrieveSnippetContent(snippet)

	if err != nil {

		return err

	}

	file, err := os.OpenFile(targetFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening or creating file:", err)
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(snippetContent); err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	return nil

}

func RetrieveSnippetContent(snippet Snippet) (string, error) {

	bytes, err := os.ReadFile(snippet.FileContentName)
	if err != nil {
		return "", err
	}

	return string(bytes), nil

}

func GenerateFilenameContentSnippet(filenameInfoSnippets string, snippet Snippet) string {

	fileName := fmt.Sprintf("%s%s%s",
		filenameInfoSnippets+"snippets",
		string(os.PathSeparator), snippet.Name)

	return fileName
}

func DeleteSnippet(fileName, name string) error {
	snippet, err := RetrieveSnippet(fileName, name)
	if err != nil {
		return err
	}

	err = os.Remove(snippet.FileContentName)

	if err != nil {
		// Si ocurre un error, lo reportamos
		return fmt.Errorf("error deleting snippet file content: %w", err)
	}

	snippets, err := ReadItems(fileName)

	if err != nil {
		return fmt.Errorf("error reading items: %w", err)
	}

	delete(snippets, name)

	err = SaveSnippets(fileName, snippets)
	if err != nil {
		return fmt.Errorf("error saving snippets after deletion: %w", err)
	}
	return nil
}
