# Library Management System Documentation

## Overview
This document provides an overview and documentation for the Library Management System developed as part of Task3. The system is designed to manage books and members within a library, facilitating operations such as book lending, returns, and member registration.

## System Architecture
The Library Management System is structured into several key components:

- **Controllers**: Handle incoming HTTP requests and respond with the appropriate actions to perform.
- **Models**: Represent the data structures used within the system, including `Book` and `Member`.
- **Services**: Contain the business logic of the application, orchestrating actions between the controllers and models.

## Setup and Installation
To set up the Library Management System, ensure you have Go installed on your system. Then, navigate to the `Task3/LibraryManagement` directory and run:

```sh
go mod tidy
go build
./LibraryManagement
```
