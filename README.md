# Task Tracker

Task Tracker is a CLI based application to manage your tasks, built as part of [roadmap.sh](https://roadmap.sh/projects/task-tracker) challenge.

## How to run

Clone the repository and run the following command:
```bash
git clone https://github.com/architxkumar/Task-Tracker
cd Task-Tracker
```

Run the following command to build and run the project:
```bash
go build -o task-cli
# Adding a new task
./task-cli add "Buy groceries"
# Updating and deleting tasks
./task-cli update 1 "Buy groceries and cook dinner"
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