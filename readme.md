# Task Manager CLI

A quick command line task manager written in Golang.

Project URL: https://roadmap.sh/projects/task-tracker

My solution: https://roadmap.sh/projects/task-tracker/solutions?u=6500425a5ce9f4ca58b8b017

## Requirements

- Go 1.24.8 or higher
- git

## Installation

Linux:

```sh
git clone https://github.com/rillmind/taskManagerCli.git
cd taskManagerCli
go build -o task-cli ./*/*.go
sudo cp task-cli /usr/bin/
```

Windows: 

```sh
git clone https://github.com/rillmind/taskManagerCli.git
cd taskManagerCli
go build -o task-cli ./*/*.go
```

Note: If you are on Windows, you need to add the executable binary (.exe) to the environment
variables manually. On Linux, just enter your super user password. If you want to use it without
add to environment variables, in Linux: just delete the last line and run './task-cli', in
Windows: just run normally with './task-cli' without adding or removing command lines.

## Code information

### Attributes

- id: int
- description: string
- status: enum
- createdAt: time.Now()
- updatedAt: time.Now()

### Functions

- **CreateTask :** Create a task and generate a json file to store the task.
- **UpdateTask :** Search task by id and change the 'title' or 'description' on '.json'.
- **DeleteTask :** Search task by id and delete from '.json'.
- **ListTasks :** Search all tasks in '.json' and show them on the screen.
- **ListTasksByStatus :** Search all tasks in '.json' by status and show them on the screen.
- **Mark :** Search task by id and change the 'status' on '.json'.
- **createJsonFile :** Creates the '.json' file with the tasks.
- **readJsonFile :** Reads the '.json' file with the tasks.