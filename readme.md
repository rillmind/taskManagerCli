# Task Manager CLI

A quick command line task manager written in Golang. This application allows you to add, update
delete, list, and mark the status of tasks. Tasks are stored in a JSON file for persistence.

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
go build -o task-cli ./cmd
sudo cp task-cli /usr/bin/
```

Windows: 

```sh
git clone https://github.com/rillmind/taskManagerCli.git
cd taskManagerCli
go build -o task-cli.exe .\cmd\
```

Note: If you are on Windows, you need to add the executable binary (.exe) to the environment
variables manually. On Linux, just enter your super user password. If you want to use it without
add to environment variables, in Linux: just delete the last line and run './task-cli', in
Windows: just run normally with './task-cli' without adding or removing command lines.

## Run example

```sh
# Adding a new task
./task-cli add "Buy contonete"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
./task-cli update 1 "Buy contonete and snacks"
./task-cli delete 1

# Marking a task as in progress or done
./task-cli mark-in-progress 1
./task-cli mark-done 1

# Listing all tasks
./task-cli list

# Listing tasks by status
./task-cli list done
./task-cli list todo
./task-cli list in-progress
```

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
