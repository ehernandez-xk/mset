package mset

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
)

const (
	catalogBaseDir       = ".m2"
	catalogPathEnvVar    = "MSET_CATALOG_PATH" // replace default e.g ~/myDirectory
	catalogDirectoryName = ".mset"
	fileSuffix           = "-settings.xml"
	fileEntryRegexp      = `^[a-zA-Z0-9-]*$`
	configFileCurrent    = ".mcurrent"
)

// check if the file exist
func fileExist(path string) error {
	fi, err := os.Stat(path)

	if os.IsNotExist(err) {
		return err
	}
	if fi.IsDir() {
		return fmt.Errorf("path '%v' is a directory", path)
	}
	return nil
}

//
func isValidName(name string) bool {
	e := regexp.MustCompile(fileEntryRegexp)
	return e.MatchString(name)
}

// validName+ 'settings.xml' e.g. danta-settings.xml
func setEntryName(name string) string {
	return name + fileSuffix
}

// check if all files in the catalog to check if the name exist
func isNameAlreadyTaken(name string) bool {
	return false
}

// if validName then can be saved in the catalog, checks for the EnvVar
func saveFileToCatalog(name, path string) error {
	destination := filepath.Join(getCatalogPath(), setEntryName(name))
	return copy(path, destination)
}

// a list of the names in the catalog, files that follow  xxxx-settings.xml
func filesInCatalog() []string {
	return []string{}
}

// save current file selected in a .mcurrent, the file is created if not present
func saveCurrentFile(name string) error {
	catalogPath := getCatalogPath()
	err := os.MkdirAll(catalogPath, 0666)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(catalogPath+"/"+name, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	fmt.Println(f.Name())
	return nil
}

// checks if the Envar is set otherwise returns default catalog path
func getCatalogPath() string {
	if path := os.Getenv(catalogPathEnvVar); path != "" {
		return filepath.Join(path, catalogDirectoryName)
	}
	return filepath.Join(getUserHomeDir(), catalogBaseDir, catalogDirectoryName)
}

func getUserHomeDir() string {
	u, _ := user.Current()
	return u.HomeDir
}

// creates the directory .mset in the path
func createCatalogDirectory() error {
	return os.MkdirAll(getCatalogPath(), 0755)
}

// copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
