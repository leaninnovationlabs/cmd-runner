# Project Title
Command runner lets you run a sequence of commands in a container or host machine. Useful to run a sequence of commands in a CI/CD pipeline or doing setup locally.

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

## Installation
Download the latest release from bin folder and add it to your path.

## Usage
You can run the command runner by passing the steps.yaml file and the parameters to be replaced in the commands.

```bash
cmd-runner cmd.yaml --name=demouser --appname=demoapp
```

## Installation and Build Instructions
Before you begin, ensure you have met the following requirements:
* You have installed the latest version of [Go](https://golang.org/dl/).
* You have a Windows/Linux/macOS machine.

To build from source follow these steps:

```bash
git clone https://github.com/leaninnovationlabs/cmd-runner.git
cd cmd-runner
```

## Build and Run
To build and run cmd-runner, follow these steps:

``` bash
go build -o cmd-runner
./cmd-runner steps.yaml --name=demouser --appname=demoapp
```
