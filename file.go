package mset

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	catalogBaseDir        = ".m2"
	catalogPathEnvVar     = "MSET_CATALOG_PATH" // replace default e.g ~/myDirectory
	mavenPathEnvVar       = "MSET_MAVEN_PATH"
	mavenSettingsFileName = "settings.xml"
	catalogDirectoryName  = ".mset"
	fileSuffix            = "-settings.xml"
	fileEntryRegexp       = `^[a-zA-Z0-9-]*$`
	configFileCurrent     = ".mcurrent"
)

// check if the file exist
func fileExist(path string) error {
	fi, err := pathExist(path)
	if err != nil {
		return err
	}

	if fi.IsDir() {
		return fmt.Errorf("path '%v' is a directory", path)
	}
	return nil
}

func directoryExist(path string) error {
	fi, err := pathExist(path)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return fmt.Errorf("path '%v' is a file", path)
	}
	return nil
}

func pathExist(path string) (os.FileInfo, error) {
	fi, err := os.Stat(path)

	if os.IsNotExist(err) {
		return nil, err
	}
	return fi, nil
}

//
func isValidName(name string) bool {
	match, _ := regexp.MatchString(fileEntryRegexp, name)
	return match
}

// validName+ 'settings.xml' e.g. danta-settings.xml
func setEntryName(name string) string {
	return name + fileSuffix
}

// check if all files in the catalog to check if the name exist
func isNameAlreadyTaken(name string) bool {
	m, _ := filesInCatalog()
	for k := range m {
		if k == name {
			return true
		}
	}
	return false
}

// if validName then can be saved in the catalog, checks for the EnvVar
func saveFileToCatalog(name, path string) error {
	destination := filepath.Join(getCatalogPath(), setEntryName(name))
	return copy(path, destination)
}

// a list of the names in the catalog, files that follow  xxxx-settings.xml
func filesInCatalog() (map[string]string, error) {
	catalogPath := getCatalogPath()
	allFiles, err := ioutil.ReadDir(catalogPath)
	files := map[string]string{}
	if err != nil {
		return files, err
	}

	for _, f := range allFiles {
		if !f.IsDir() {
			if strings.HasSuffix(f.Name(), fileSuffix) {
				files[strings.TrimSuffix(f.Name(), fileSuffix)] = filepath.Join(catalogPath, f.Name())
			}
		}
	}
	return files, nil
}

func findFilePathInCatalog(name string) (string, error) {
	files, err := filesInCatalog()
	if err != nil {
		return "", err
	}
	if len(files) == 0 {
		return "", fmt.Errorf("files not available in catalog")
	}
	for k, path := range files {
		if k == name {
			return path, nil
		}
	}
	return "", fmt.Errorf("file not found in catalog")
}

// save current file selected in a .mcurrent, the file is created if not present
func saveCurrentFile(name string) error {
	catalogPath := getCatalogPath()
	err := os.MkdirAll(catalogPath, 0666)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filepath.Join(catalogPath, name), os.O_RDONLY|os.O_CREATE, 0666)
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
func getMavenPath() string {
	if path := os.Getenv(mavenPathEnvVar); path != "" {
		return filepath.Join(path, catalogBaseDir)
	}
	return filepath.Join(getUserHomeDir(), catalogBaseDir)
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

func setFileToMavenPath(path string) error {
	mavenPath := getMavenPath()
	return copy(path, filepath.Join(mavenPath, mavenSettingsFileName))
}

func setNameToCurrent(name string) error {
	currentPath := filepath.Join(getCatalogPath(), configFileCurrent)
	ioutil.WriteFile(currentPath, []byte(name), 0644)
	return nil
}
func getCurrentName() (string, error) {
	currentPath := filepath.Join(getCatalogPath(), configFileCurrent)
	c, err := ioutil.ReadFile(currentPath)
	if err != nil {
		return "", err
	}
	return string(c), nil
}

func removeFile(path string) error {
	if err := fileExist(path); err != nil {
		return err
	}
	return os.Remove(path)
}
