package cmd

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestCreateFile(t *testing.T) {
	fs := afero.NewMemMapFs() // Create an in-memory file system

	content := "Test content"
	dir := "/testdir"
	filename := "testfile.txt"

	err := CreateFileWithFs(fs, dir, filename, content)
	assert.Nil(t, err)

	fileContent, readErr := afero.ReadFile(fs, dir+"/"+filename)
	assert.Nil(t, readErr)
	assert.Equal(t, content, string(fileContent))
}

func TestCreateDirectoryAndFiles(t *testing.T) {
	fs := afero.NewMemMapFs()

	dirName := "/testdir"
	CreateDirectoryAndFilesWithFs(fs, dirName)

	// Assert directories
	for _, dir := range Directories {
		exists, _ := afero.DirExists(fs, dirName+"/"+dir)
		assert.True(t, exists)
	}

	// Assert main files
	for _, file := range MainFiles {
		exists, _ := afero.Exists(fs, dirName+"/"+file)
		assert.True(t, exists)
	}

	// Assert test files
	testDir := dirName + "/tests"
	for _, file := range TestFiles {
		exists, _ := afero.Exists(fs, testDir+"/"+file)
		assert.True(t, exists)
	}
}

// CreateFileWithFs is a version of CreateFile that accepts an afero.Fs
func CreateFileWithFs(fs afero.Fs, dir, filename, content string) error {
	path := dir + "/" + filename
	file, err := fs.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// CreateDirectoryAndFilesWithFs is a version of CreateDirectoryAndFiles that accepts an afero.Fs
func CreateDirectoryAndFilesWithFs(fs afero.Fs, dirName string) {
	fs.Mkdir(dirName, 0755)

	// IaC files for the main directory
	for _, file := range MainFiles {
		CreateFileWithFs(fs, dirName, file, "# Generated by Tofuinit")
	}

	// Create all directories required
	for _, dir := range Directories {
		fs.Mkdir(dirName+"/"+dir, 0755)
	}

	// Create file in tests dir
	testDir := dirName + "/tests"
	for _, file := range TestFiles {
		CreateFileWithFs(fs, testDir, file, "# Generated by Tofuinit")
	}

}