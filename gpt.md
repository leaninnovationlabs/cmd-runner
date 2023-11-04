# Steps


- Step1: Giving basic structure for what needs to be done

```
Help me to write a program in go, that will take list of command line statements to execute and run them. Will provide the statements to execute in yaml format, with name of the step and commands to execute

- create
  commands: |
    #!/bin/bash
    helm upgrade flowscript-ml-api stable-myapp/microservice --values ./dev-service-def.yaml --namespace kube-system

- update
  commands: |
       echo 'hello world'

```


Output:

1. Gave me how to setup the go packages, in this case it wanted to use gopkg.in/yaml.v2
2. Gave an example of main.go where it create a struct for parsing the yaml and store it , provided a way to parse each step and run them one by one
3. Gave instructions on how to run the go file
4. Gave me a sample based on what provided, but it fixed few things like the yaml format, and the commands to run. Gave a key value syntax for the commands to run, and also gave a way to run the commands in parallel

What I needed to fix:
1. Package seem to be old, so fixed it to use os

- Step2: Now I want to pass some params and what those to be injected in the commands to run. Gave the following prompt

```
can pass some parameters while running this code and the params need to be evaluated before running the commands. So commands can look something like the following 

- name: update
  commands: |
    echo 'hello world {name} and {message}'

- name: update
  commands: |
    echo 'hello world {message}'

_____

We can invoke from go by passing hte name and message parameters,

 go run main.go steps.yaml  --name=Sam  --message=test

```

Output:
1. Told me to use a flag library to parse the command-line arguments.
2. Gave some sample implementation with a simple template substitution mechanism.
3. Regenerated the code again with the new changes, where it reads the command line arguments and pass it to the template substitution mechanism

Able to run this file

Problem: It gave the code where it was exactly looking for the keys name and message, I should I have made it clear but I understood that it was looking for the keys name and message


- Step3: Now I made it more explicit that the parametes need to be dynamic. Here is the prompt

```
change the logic above so that the flags are dynamic, can be any key value pair that user can pass
```

Output:
1. Able to pass the params as key value pairs nicely from the command line, worked like a charm

```
go run main.go steps.yaml --name=Sam --message=test --anotherKey=anotherValue
```


Step4: Want to add couple of attributes name and description to the steps, and want to print them out. Here is the prompt

```
Need to change the format of the yaml so that we have multiple steps under each 

name: microservice
description: Will support creating and deploying a microserve
steps:
  - name: create
    image: ubuntu:18.04
    commands: |
      #!/bin/bash
      echo 'create workload {name}'

  - name: update
    commands: |
      echo 'hello world {appname}'

```

Output:
1. Gave me a way to parse the yaml and store it in a struct, and also gave me a way to print the steps. Dint sweat at all, it was a breeze. Got the new struct and everything worked fine


Step5: Thought time to do some documentation, why to write when I can command to do it for me. Here is the prompt

```
generate a readme file with following instructions in markup format

1. Installation instruction
2. How to Build and run 
3. Give some sample files
```

Output:

1. Gave a nice format with bunch of placeholders
2. Since I am using the chat interface, looks like the output is somewhat messed up. Parts of the markup gave me in nice code blocks but some of them as html. Had to copy and format. Now I am spoiled and felts that way too much work


Step6: See all the output going into to console, which is good but sometimes we need to save it to a log file so here is my prompt

```
can you add support to output log to a file if users provides by specifying a parameter --out
```

Output:
1. Worked like a charm, now I can pass an output parameter and it will write the output to the file. I am happy


Step7: All but, but I am still not 100% happy. I want to run all the commands in a docker container so that we can run them in a nicely isolated environment. Here is my prompt

```
If the user specified image value in the steps section, need to add instructions to download the docker image specified and run the commands within that specific docker image. Modify the command execution based on this
```

Output:
1. Gave me the instructions on how to add code in go that will do docker pull if image is provided and run the commands in the docker container. Thats amazin. Tested it and it worked, oohoo
2. This time, it did not generate the whole file, just gave me the parts that will do what I need to do and provided me instructions on where to add. Getting smarter


Step7: This time I tested by passing the logs file under logs folder and it failed. So I prompted it to fix it

```
If the folder structure does not exists when log file is specified, create the required folder structure
```

Output:
1. Gave me exact code that will check if the dir exist and if not create it. Copy paste and I am done



Step8: Wanted to add a function that will list out all the available steps

```
give me a function to list out all the available steps
```

Output:
1. This is easy, gave me the function and how to hook it in



Step9: Now I started to get ambitious. I know lot of times I want to pass things like keys and other info via an env file rather than params. So here is my prompt


```
some of the parameters can be derived by reading a .env file and set the parameters while running the commands
```

Output:

1. Gave me to include github.com/joho/godotenv library
2. Gave me instructions to load at the beginign of the main function
3. Gave me a function that will take a string and replace based on env variables. Exactly what I wanted
4. Provided me instructions on where to add in the big code



Step10: Realized that if I have a multi line command, when running it from docker its failing. I wasnt sure why, or got too lazy to even debug. So asked GPT with the following prompt

``` 
while passing the commands and run it from docker, its giving an error if there are multiple commands per line. Can you see how I can fix that
```


Output:
1. Gave me a command that will split the multiline command and append with && so that when its passing to docker run command so that it will work. Not optimal but it works


-----------------

Final prompts for a web page

Step11: Now I want to generate a html page with following sections

```
So this program essentially take a build file with multiple steps. Users can give their instructions using cli commands and the program will generate them and generate the output in log files. Users can even run the the commands in a docker environment of their choice. Allows users to define multiple steps. This makes it super easy to provide build steps and make it as part of source code and run it locally on developer box without any complex setup or relying on cloud runners. Can you generate a html page with following sections

1. Hero section, with a one line pitch
2. Features section, with 3 advantages of the product
3. Example section
```

Output:
1. Gave me the html that I just added to docs folder



Step12: I can make some text tweaks but not happy with basic html. Want it to use some styling using tailwind. So here is my prompt

```
can you generate using tailwindui
```

Output:
1. Nice tailwind page, which is good enough to go with
