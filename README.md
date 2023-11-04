# Project Title
Command runner lets you run a sequence of commands in a container or host machine. It is useful for running a sequence of commands in a CI/CD pipeline or doing setup locally.

Replace steps.yaml with your own configuration file and provide the necessary flags as needed.


## Sample Files
Below are some sample YAML configuration files that you can use to get started:

sample-steps.yaml:

```yaml
name: microservice
description: Will support creating and deploying a microservice
steps:
  - name: setup
    commands: |
      echo 'Setup Dev Environment'
      #! Setup Virtual Environment
      python -m venv ./venv
      source ./venv/bin/activate
      pip install --upgrade pip
      pip install -r requirements.txt      
  - name: create
    image: ubuntu:18.04
    commands: |
      echo 'Creating service {name}'
      helm create {appname}
  - name: update
    commands: |
      echo 'Update service {name}'
      helm upgrade --install {appname}
```

To use the sample configuration file, run the following command:

```bash
cmd-runner cmd.yaml --name=demouser --appname=demoapp
```

## Installation and Build Instructions
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
