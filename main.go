package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path_to_yaml> --key=value [--out=log.txt]")
		os.Exit(1)
	}

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, continuing without it")
	}

	// Extract the YAML file argument.
	yamlFilePath := os.Args[1]

	// Initialize a logger to write to os.Stderr by default.
	logger := log.New(os.Stderr, "", log.LstdFlags)

	dynamicFlags := make(map[string]string)
	var outputFilePath string
	var stepName string

	// Process the dynamic flags and check for an output file.
	for _, arg := range os.Args[2:] {
		if strings.HasPrefix(arg, "--") {
			keyValue := strings.SplitN(arg[2:], "=", 2)
			if len(keyValue) != 2 {
				logger.Printf("Invalid flag format: %s\n", arg)
				os.Exit(1)
			}
			if keyValue[0] == "out" {
				outputFilePath = keyValue[1]
			} else {
				dynamicFlags[keyValue[0]] = keyValue[1]
			}

			if keyValue[0] == "step" {
				logger.Printf("Step to execute: %s\n", keyValue[1])
				stepName = keyValue[1]
			}
		}
	}

	// If an output file is specified, change the logger to write to the given file.
	if outputFilePath != "" {
		dir := filepath.Dir(outputFilePath)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			logger.Fatalf("Failed to create directory structure for log file: %s", err)
		}

		// Ensure the directory structure exists.
		outFile, err := os.Create(outputFilePath)
		if err != nil {
			logger.Fatalf("Error creating log file: %s", err)
		}
		defer outFile.Close()
		logger.SetOutput(outFile)
	}

	runSteps(logger, yamlFilePath, dynamicFlags, stepName)
}
