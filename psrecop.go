package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	sourceDir      = "AppData\\Roaming\\Microsoft\\Windows\\PowerShell\\PSReadLine"
	destinationDir = "C:\\windows\\temp\\PSHISTORY"
	fileName       = "ConsoleHost_history.txt"
)

const (
	colorGreen  = "\033[1;32m"
	colorYellow = "\033[1;33m"
	colorPurple = "\033[1;35m"
	colorRed    = "\033[1;31m"
	colorReset  = "\033[0m"
)

func main() {
	// get directories
	userDirs, err := getUsersDirectories()
	if err != nil {
		fmt.Printf("%sError al obtener los directorios de usuarios: %v%s\n", colorYellow, err, colorReset)
		return
	}
	// create dest directory
	if err := os.MkdirAll(destinationDir, os.ModePerm); err != nil {
		fmt.Printf("%sError creating destination directory: %v%s\n", colorYellow, err, colorReset)
		return
	}

	// get history from each user
	for _, userDir := range userDirs {
		sourcePath := filepath.Join("C:\\Users", userDir, sourceDir, fileName)
		destinationPath := filepath.Join(destinationDir, fmt.Sprintf("%s_%s", userDir, fileName))

		if err := copyFile(sourcePath, destinationPath); err != nil {
			fmt.Printf("%sPowershell history not found for user '%s' %v\n", colorRed, userDir, colorReset)
		} else {
			fmt.Printf("%sPowershell history of user '%s' succesfully saved.%s\n", colorGreen, userDir, colorReset)
		}
	}
	fmt.Printf("%sFiles saved in C:\\Windows\\Temp\\PSHISTORY directory%s\n", colorPurple, colorReset)
}

// get users directories
func getUsersDirectories() ([]string, error) {
	var userDirs []string

	entries, err := os.ReadDir("C:\\Users")
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			userDirs = append(userDirs, entry.Name())
		}
	}

	return userDirs, nil
}

// copy from source to dest
func copyFile(sourcePath, destinationPath string) error {
	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
