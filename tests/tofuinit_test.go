// package cmd

// import (
// 	"io"
// 	"os"
// 	"path/filepath"
// 	"testing"

// 	"github.com/spf13/afero"
// )

// // FileSystem abstraction
// type FileSystem interface {
// 	Create(name string) (afero.File, error)
// 	Mkdir(name string, perm os.FileMode) error
// }

// type OSFileSystem struct{}

// func (OSFileSystem) Create(name string) (afero.File, error) {
// 	return os.Create(name)
// }

// func (OSFileSystem) Mkdir(name string, perm os.FileMode) error {
// 	return os.Mkdir(name, perm)
// }

// var MainFiles = []string{
// 	"main.tf",
// 	"variables.tf",
// 	"versions.tf",
// 	"outputs.tf",
// }

// var TestFiles = []string{
// 	"main.tftest.hcl",
// }

// type Data struct {
// 	Title       string
// 	Description string
// }

// func CreateFile(fs FileSystem, dir, filename, content string) error {
// 	path := filepath.Join(dir, filename)
// 	file, err := fs.Create(path)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	_, err = io.WriteString(file, content)
// 	return err
// }

// // ... (rest of your functions)

// // === Tests ===
// func TestCreateFile(t *testing.T) {
// 	fs := afero.NewMemMapFs()

// 	dir := "testDir"
// 	filename := "testFile.txt"
// 	content := "This is a test."

// 	err := CreateFile(fs, dir, filename, content)
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
	}

// 	filePath := filepath.Join(dir, filename)
// 	exists, err := afero.Exists(fs, filePath)
// 	if err != nil || !exists {
// 		t.Errorf("File was not created: %s", filePath)
// 	}
// }
