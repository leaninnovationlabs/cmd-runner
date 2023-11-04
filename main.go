package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

// CommandTask represents the overall task structure with a name, description, and steps.
type CommandTask struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Steps       []struct {
		Name     string `yaml:"name"`
		Image    string `yaml:"image,omitempty"`
		Commands string `yaml:"commands"`
	} `yaml:"steps"`
}

func ListSteps(filePath string) {
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %s", err)
	}

	var task CommandTask
	err = yaml.Unmarshal(yamlFile, &task)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s", err)
	}

	fmt.Println("Available Steps:")
	for _, step := range task.Steps {
		fmt.Printf("- %s\n", step.Name)
	}
}

func ExecuteCommands(stepCmd string) string {
	replacedCmd := stepCmd
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		placeholder := fmt.Sprintf("{%s}", pair[0])
		if !strings.Contains(replacedCmd, placeholder) {
			continue
		}
		replacedCmd = strings.ReplaceAll(replacedCmd, placeholder, pair[1])
	}
	return replacedCmd
}

// Helper function to process and combine commands into a single executable string.
func combineCommands(commands string) string {
	// Split the commands into lines and combine them with '&&' ensuring that
	// each individual command is trimmed of whitespace and not a comment.
	var combinedCommands []string
	commandLines := strings.Split(commands, "\n")
	for _, line := range commandLines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" && !strings.HasPrefix(trimmedLine, "#") {
			combinedCommands = append(combinedCommands, trimmedLine)
		}
	}
	return strings.Join(combinedCommands, " && ")
}

func runSteps(logger *log.Logger, yamlFilePath string, dynamicFlags map[string]string, stepName string) {

	// Read the YAML file.
	yamlFile, err := os.ReadFile(yamlFilePath)
	if err != nil {
		logger.Fatalf("Error reading YAML file: %s", err)
	}

	// Parse the YAML.
	var task CommandTask
	err = yaml.Unmarshal(yamlFile, &task)
	if err != nil {
		logger.Fatalf("Error parsing YAML file: %s", err)
	}

	logger.Printf("Running task: %s\n", task.Name)
	logger.Printf("Description: %s\n", task.Description)

	for _, step := range task.Steps {

		var cmd *exec.Cmd

		if stepName != "" && step.Name != stepName {
			continue
		}
		logger.Printf("Executing step: %s\n", step.Name)

		if step.Image != "" {
			logger.Printf("Using image: %s\n", step.Image)
			logger.Printf("Using Docker image: %s\n", step.Image)
			err := exec.Command("docker", "pull", step.Image).Run()
			if err != nil {
				logger.Fatalf("Failed to pull Docker image: %s", err)
			}

			dockerRunArgs := []string{"run", "--rm", "-v", fmt.Sprintf("%s:/%s", os.Getenv("PWD"), os.Getenv("PWD")), "-w", os.Getenv("PWD"), step.Image}

			step.Commands = combineCommands(step.Commands)

			// Execute the commands with dynamic flags replacement.
			for key, value := range dynamicFlags {
				placeholder := fmt.Sprintf("{%s}", key)
				step.Commands = strings.ReplaceAll(step.Commands, placeholder, value)
				step.Commands = ExecuteCommands(step.Commands)
			}

			commands := strings.Split(step.Commands, "\n")
			for _, cmdStr := range commands {
				cmdStr = strings.TrimSpace(cmdStr)
				if cmdStr == "" || strings.HasPrefix(cmdStr, "#") {
					continue
				}

				// Create a docker run command that mounts the current directory
				// Parse the commands to run them inside the docker container
				dockerRunArgs = append(dockerRunArgs, "bash", "-c", cmdStr)
				cmd = exec.Command("docker", dockerRunArgs...)
				cmd.Stdout = logger.Writer()
				cmd.Stderr = logger.Writer()
				err := cmd.Run()
				if err != nil {
					logger.Printf("Error executing command '%s': %s\n", cmdStr, err)
				}

				dockerRunArgs = dockerRunArgs[:8]
			}
		} else {
			logger.Printf("Using host machine to run commands\n")
			commands := strings.Split(step.Commands, "\n")

			for _, cmdStr := range commands {
				cmdStr = strings.TrimSpace(cmdStr)
				// Ignore empty lines and comments.
				if cmdStr == "" || strings.HasPrefix(cmdStr, "#") {
					continue
				}

				// Clean the arguments for the next command to avoid repetition.
				// Split the commands into lines and process them.
				// Ignore empty lines and comments.
				for key, value := range dynamicFlags {
					placeholder := fmt.Sprintf("{%s}", key)
					// Perform replacements in the command string.
					cmdStr = strings.ReplaceAll(cmdStr, placeholder, value)
					cmdStr = ExecuteCommands(cmdStr)
				}

				cmdArgs := strings.Fields(cmdStr)
				if len(cmdArgs) == 0 {
					continue
				}

				// Execute the commands with dynamic flags replacement.
				cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
				cmd.Stdout = logger.Writer()
				cmd.Stderr = logger.Writer()
				err := cmd.Run()
				if err != nil {
					logger.Printf("Error executing command '%s': %s\n", cmdStr, err)
				}
			}
		}

	}
}

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
