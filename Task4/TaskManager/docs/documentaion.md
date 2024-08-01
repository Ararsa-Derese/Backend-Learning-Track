# Task Manager API Documentation

## Overview
This document outlines the API endpoints provided by the Task Manager system, allowing for the management of tasks within a project management context. The system supports operations such as creating, updating, retrieving, and deleting tasks. You can find postman docs [here](https://documenter.getpostman.com/view/30253109/2sA3kdBHnG)

## API Endpoints

### Get All Tasks
- **Endpoint**: `GET /tasks`
- **Description**: Retrieves a list of all tasks.
- **Response**: An array of task objects.

### Get Task by ID
- **Endpoint**: `GET /tasks/:id`
- **Description**: Retrieves a task by its unique identifier.
- **Parameters**:
  - `id`: The unique identifier of the task.
- **Response**: A task object.

### Add a New Task
- **Endpoint**: `POST /tasks`
- **Description**: Adds a new task to the system.
- **Body**:
  - `title`: The title of the task.
  - `description`: The description of the task.
  - `status`: The current status of the task.
- **Response**: A message indicating success or failure.

### Update a Task
- **Endpoint**: `PUT /tasks/:id`
- **Description**: Updates an existing task.
- **Parameters**:
  - `id`: The unique identifier of the task to update.
- **Body**:
  - `title`: The new title of the task (optional).
  - `description`: The new description of the task (optional).
  - `status`: The new status of the task (optional).
- **Response**: A message indicating success or failure.

### Delete a Task
- **Endpoint**: `DELETE /tasks/:id`
- **Description**: Deletes a task from the system.
- **Parameters**:
  - `id`: The unique identifier of the task to delete.
- **Response**: A message indicating success or failure.

## Models
### Task
- **Fields**:
    - `ID`: Unique identifier for the task.
    - `Title`: Title of the task.
    - `Description`: Description of the task.
    - `DueDate`: Due date of the task.
    - `Status`: Current status of the task.

## Error Handling
All endpoints return appropriate HTTP status codes along with error messages in the case of failure.


## Error Handling
All endpoints return appropriate HTTP status codes along with error messages in the case of failure.

