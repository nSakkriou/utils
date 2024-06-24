package util

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func IsEmpty(str string) bool {
	return str == ""
}

func ArrayContains(array []string, elem string) bool {
	for _, arrayElem := range array {
		if elem == arrayElem {
			return true
		}
	}

	return false
}

func FileExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}

		return false
	}
	return true
}

func Prefix(s, char string) string {
	if len(char) != 1 {
		return s
	}
	if len(s) > 0 && s[0] == char[0] {
		return s
	}
	return char + s
}

// Replace all targetString occurence in file (file path), by newString
func ReplaceAndWrite(filePath, targetString, newString string) error {
	file, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Impossible de d'ouvrir le fichier " + filePath)
		return err
	}

	fileString := string(file[:])
	newFileString := strings.ReplaceAll(fileString, targetString, newString)
	file = []byte(newFileString)

	err = os.WriteFile(filePath, file, 0777)

	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier " + filePath)
	}

	return err
}

// Copy all file and folder from source dir to dest dir
func CopyAll(sourceDir, destDir string) error {
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == sourceDir {
			return nil
		}

		relativePath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destDir, relativePath)

		if info.IsDir() {
			err = os.MkdirAll(destPath, info.Mode())
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(path, destPath)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func CopyFile(sourcePath, destPath string) error {
	destDir := filepath.Dir(destPath)
	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return err
	}

	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	sourceInfo, err := os.Stat(sourcePath)
	if err != nil {
		return err
	}

	return os.Chmod(destPath, sourceInfo.Mode())
}
