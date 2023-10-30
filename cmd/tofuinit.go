package cmd

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
)

// List of directories created
var Directories = []string{
	"tests",
	"docs",
	"examples",
}

// List of main files to be created
var MainFiles = []string{
	"main.tf",
	"variables.tf",
	"versions.tf",
	"outputs.tf",
}

// List of test files to be created
var TestFiles = []string{
	"main.tftest.hcl",
}

//go:embed templates/*
var tmplContent embed.FS

type TemplateData struct {
	DirName string
}

// Creates a file with the specified content.
func CreateFile(dir, filename, content string) error {
	path := filepath.Join(dir, filename)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, content)
	return err
}

// Handle errors, logging and exiting if there's a failure.
func CheckError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Function to create directory structure and necessary files
func CreateDirectoryAndFiles(dirName string) {

	err := os.Mkdir(dirName, 0755)
	CheckError(fmt.Sprintf("Failed to create directory %s", dirName), err)

	// IaC files for the main directory
	for _, file := range MainFiles {
		err = CreateFile(dirName, file, "# Generated by Tofuinit")
		CheckError(fmt.Sprintf("Failed to create file %s", file), err)
	}

	// Create all directories required
	for _, dir := range Directories {
		dirs := filepath.Join(dirName, dir)
		err := os.Mkdir(dirs, 0755)
		CheckError(fmt.Sprintf("Failed to create directory %s", dir), err)
	}
	// Create file in tests dir
	testDir := filepath.Join(dirName, "tests")
	for _, file := range TestFiles {
		err = CreateFile(testDir, file, "# Generated by Tofuinit")
		CheckError(fmt.Sprintf("Failed to create file %s in tests directory", file), err)
	}

	err = CreateMainReadmeFromTemplate(dirName)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = CreateDocsReadmeFromTemplate(dirName)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = CreateExamplesReadmeFromTemplate(dirName)
	if err != nil {
		fmt.Println("Error:", err)
	}

}

func CreateExamplesReadmeFromTemplate(dirName string) error {
	// Read content from embedded (assuming template.txt is a file inside templates directory)
	contentBytes, err := tmplContent.ReadFile("templates/examples_readme.md")
	if err != nil {
		return fmt.Errorf("Failed to read embedded template: %v", err)
	}
	content := string(contentBytes)

	// Create a new template and parse the content
	tmpl, err := template.New("readme").Parse(content)
	if err != nil {
		return fmt.Errorf("Failed to parse template: %v", err)
	}

	// Prepare the data
	data := TemplateData{
		DirName: dirName,
	}

	// Render the template
	subFolderName := "examples"
	subFolderPath := filepath.Join(dirName, subFolderName)
	path := filepath.Join(subFolderPath, "README.md")
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Failed to create README.md: %v", err)
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("Failed to render template: %v", err)
	}

	return nil
}

func CreateDocsReadmeFromTemplate(dirName string) error {
	// Read content from embedded (assuming template.txt is a file inside templates directory)
	contentBytes, err := tmplContent.ReadFile("templates/docs_readme.md")
	if err != nil {
		return fmt.Errorf("Failed to read embedded template: %v", err)
	}
	content := string(contentBytes)

	// Create a new template and parse the content
	tmpl, err := template.New("readme").Parse(content)
	if err != nil {
		return fmt.Errorf("Failed to parse template: %v", err)
	}

	// Prepare the data
	data := TemplateData{
		DirName: dirName,
	}

	// Render the template
	subFolderName := "docs"
	subFolderPath := filepath.Join(dirName, subFolderName)
	path := filepath.Join(subFolderPath, "README.md")
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Failed to create README.md: %v", err)
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("Failed to render template: %v", err)
	}

	return nil
}

func CreateMainReadmeFromTemplate(dirName string) error {
	// Read content from embedded (assuming template.txt is a file inside templates directory)
	contentBytes, err := tmplContent.ReadFile("templates/main_readme.md")
	if err != nil {
		return fmt.Errorf("Failed to read embedded template: %v", err)
	}
	content := string(contentBytes)

	// Create a new template and parse the content
	tmpl, err := template.New("readme").Parse(content)
	if err != nil {
		return fmt.Errorf("Failed to parse template: %v", err)
	}

	// Prepare the data
	data := TemplateData{
		DirName: dirName,
	}

	// Render the template
	path := filepath.Join(dirName, "README.md")
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Failed to create README.md: %v", err)
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("Failed to render template: %v", err)
	}

	return nil
}