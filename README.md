# Project Title

A brief description of what this project does and who it's for.

## Installation Instructions

Before you begin, ensure you have met the following requirements:
* You have installed the latest version of [Go](https://golang.org/dl/).
* You have a Windows/Linux/macOS machine.

To install `cmd-runner`, follow these steps:

## Linux and macOS:
```bash
git clone https://github.com/leaninnovationlabs/cmd-runner.git
cd cmd-runner
```

## How to Build and Run
To build and run cmd-runner, follow these steps:


``` bash
go build -o cmd-runner
./cmd-runner steps.yaml --name=demouser --appname=demoapp
```

Replace steps.yaml with your own configuration file and provide the necessary flags as needed.


## Sample Files
Below are some sample YAML configuration files that you can use to get started:

sample-steps.yaml:

```yaml
name: microservice
description: Will support creating and deploying a microservice
steps:
  - name: create
    image: ubuntu:18.04
    commands: |
      echo 'Creating workload {name}'
  - name: update
    commands: |
      echo 'Hello world {appname}'
```

To use the sample configuration file, run the following command:

```bash
cmd-runner cmd.yaml --name=demouser --appname=demoapp
```

