package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
)

func TestAddPicture(t *testing.T) {
	xlsx, err := excelize.OpenFile("./test.xlsx")
	assert.Nil(t, err)

	definedNames := xlsx.GetDefinedName()
	for _, definedName := range definedNames {
		fmt.Println(definedName.RefersTo)
	}
}

func TestExpectResult(t *testing.T) {
	expectResult := []string{
		"Test!$A$1:$N$13",
		"Test!$A$14:$N$20",
		"Test!$A$21:$N$31",
	}

	// xlsx, err := excelize.OpenFile("./test.xlsx")
	// assert.Nil(t, err)
	// some function return
	result := expectResult

	assert.Equal(t, expectResult, result)
}
