package mset

import (
	"fmt"
	"path/filepath"
)

type initc struct {
	name        string
	description string
	usage       string
}

var InitCMD = initc{
	name:        "init",
	description: fmt.Sprintf("Create an empty catalog %v to store the settings.xml files", getCatalogPath()),
	usage:       "mset init",
}

func (c initc) Name() string        { return c.name }
func (c initc) Description() string { return c.description }
func (c initc) Usage() string       { return c.usage }

func (c initc) Run([]string) error {

	// check if mset is already initialized
	if err := directoryExist(getCatalogPath()); err == nil {
		return fmt.Errorf("mset already initialized on: %v", getCatalogPath())
	}

	// check if exist .m2 directory
	if err := directoryExist(getMavenPath()); err != nil {
		return err
	}
	fmt.Println("checking .m2 directory OK")

	entryName := ""
	// ask for name of acual settings.xml if exist
	if err := fileExist(filepath.Join(getMavenPath(), mavenSettingsFileName)); err == nil {
		fmt.Println("info: exist a settings.xml file in .m2 directory. Provide a name to add it to the catalog")
		var op = ""
		fmt.Print("Enter a name (required) :")
		fmt.Scanln(&op)
		if !isValidName(op) {
			return fmt.Errorf("invalid name: (%v)", fileEntryRegexp)
		}
		entryName = op
	}

	// create .mset catalog
	err := createCatalogDirectory()
	if err != nil {
		return err
	}
	fmt.Println("creating catalog .mset OK")

	// add the first entry in the catalog
	if entryName != "" {
		err = AddCMD.Run([]string{
			entryName,
			filepath.Join(getMavenPath(), mavenSettingsFileName)},
		)
		if err != nil {
			return err
		}
	}

	fmt.Println("mset initialized")
	return nil
}
