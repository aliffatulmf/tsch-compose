# TSCH Compose

This program is designed to automate and simplify the process of creating
and managing task schedules (task scheduler). The main inspiration for this program
comes from Docker Compose, which allows users to define and manage containerized
applications easily through YAML configuration files. Following a similar approach,
this program provides ease in scheduling tasks of various complexities.

## Configuration File Structure

The tsch-compose.yaml file is expected to have the following structure:

```yaml
namespace: "user"

tasks:
  task1:
    name: "Task 1"
    schedule_type: "ONCE"
    run: "program.exe"
    arguments:
      - "arg 1"
      - "arg 2"
  task2:
    name: "Task 2"
    schedule_type: "DAILY"
    run: "C:\\program2.exe"
    arguments:
      - "arg 1"
```
Each task is expected to have a name, schedule_type, run, and arguments. 

Sure, let's break down the fields in the `Task` struct:

- `name`: This field is used to specify the name of the task. It's a required field and it's used to identify the task.

- `schedule_type`: This field is used to specify when the task should be run. It's also a required field. Each of these values represents a different schedule on which the task could be run.

- `run`: This field is used to specify the command or script that should be run for the task. This could be a shell command, a script file, or any other executable command.

- `arguments`: This field is used to specify any arguments that should be passed to the command or script specified in the `run` field. This is an array, so it can contain multiple arguments.

The schedule_type must be one of the following: 

- **MINUTE** - Specifies the number of minutes before the task should run.
- **HOURLY** - Specifies the number of hours before the task should run.
- **DAILY** - Specifies the number of days before the task should run. 
- **WEEKLY** - Specifies the number of weeks before the task should run.
- **MONTHLY** - Specifies the number of months before the task should run.
- **ONCE** - Specifies that that task runs once at a specified date and time.
- **ONSTART** - Specifies that the task runs every time the system starts. You can specify a start date, or run the task the next time the system starts.
- **ONLOGON** - Specifies that the task runs whenever a user (any user) logs on.You can specify a date, or run the task the next time the user logs on.
- **ONIDLE** -  Specifies that the task runs whenever the system is idle for a specified period of time. You can specify a date, or run the task the next time the system is idle.
