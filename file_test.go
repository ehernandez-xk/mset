package mset

import (
	"testing"
)

func TestFileExist(T *testing.T) {

}

func TestIsValidName(t *testing.T) {
	table := []struct {
		name  string
		valid bool
	}{
		{name: "danta", valid: true},
		{name: "myProy", valid: true},
		{name: "my-proy", valid: true},
		{name: "danta1", valid: true},
		{name: "1danta", valid: true},
		{name: "danta.xml", valid: false},
		{name: "danta*", valid: false},
		{name: "danta>", valid: false},
	}
	for _, row := range table {
		got := isValidName(row.name)
		if got != row.valid {
			t.Errorf("name '%v' got valid: %v instead of valid: %v", row.name, got, row.valid)
		}
	}

}

func TestSetEntryName(t *testing.T) {
	input := "danta"
	output := "danta-settings.xml"
	got := setEntryName(input)
	if got != output {
		t.Errorf("got: %v instead of %v", got, output)
	}
}

func TestSaveCurrentFile(t *testing.T) {
	err := saveCurrentFile("danta")
	if err != nil {
		t.Error(err.Error())
	}
}
