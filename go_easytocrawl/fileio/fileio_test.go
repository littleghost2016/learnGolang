package fileio

import (
	"testing"
	"os"
	"fmt"
)

var (
	filepath string
)

func TestMain(m *testing.M) {
	clearFile()
	m.Run()
	clearFile()
}

func TestWorkFlow(t *testing.T) {
	t.Run("WriteFileWithIoutil", testWriteFileWithIoutil)
	t.Run("ReadFileWithIoutil", testReadFileWithIoutil)
	t.Run("WriteFileWithOS", testWriteFileWithOS)
	t.Run("ReadFileWithOS", testReadFileWithOS)
}

func testWriteFileWithIoutil(t *testing.T) {
	content := "a\n"
	err := WriteFileWithIoutil(filepath, content)
	if err != nil {
		t.Errorf("Error of WriteFileWithIoutil: %v", err)
	}
}

func testReadFileWithIoutil(t *testing.T) {
	content, err := ReadFileWithIoutil(filepath)
	fmt.Println(content)
	if err != nil || content != "a" {
		t.Errorf("Error of ReadFileWithIoutil: %v", err)
	}
}

func testWriteFileWithOS(t *testing.T) {
	content := "b"
	err := WriteFileWithOS(filepath, content)
	if err != nil {
		t.Errorf("Error of WriteFileWithOS: %v", err)
	}
}

func testReadFileWithOS(t *testing.T) {
	content, err := ReadFileWithOS(filepath)
	if err != nil && content != "a\nb" {
		t.Errorf("Error of ReadFileWithOS: %v", err)
	}
}

func clearFile() {
	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println("clearFile error in os.Getwd: ", err)
		os.Exit(0)
	}

	filepath = currentPath + "/test.txt"

	_, err = os.Stat(filepath)
	if err != nil || os.IsNotExist(err) {
		return
	}

	err = os.Remove(filepath)
	if err != nil {
		fmt.Println("clearFile error: ", err)
		os.Exit(0)
	}
}