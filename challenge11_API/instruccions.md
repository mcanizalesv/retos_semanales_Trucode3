[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/knDAa_Kb)
# RESTful API Task Manager

## Objective:

Create a web application using Go, Gorm, and Gin to manage data for Users, Projects, and Tasks. This exercise focuses on setting up routes, defining multiple models, and establishing relationships between them.

## Instructions:

### Setup:

Install Go, Gorm, and Gin on your machine.

Create a new Go project directory.

### Models:

Define three models:

- User with name and surname,
- Project with name and delivery date,
- Task with title, description, estimated hours,

with appropriate fields for each.

#### Establish relationships between the models:

Projects are the most important, they could have attached engineers and tasks.

Engineers can be attached to any Project, but they should belong ONLY to one project at a time.

A Project can have multiple Tasks.

Use Gorm to create the corresponding database tables and define associations.

### Router Setup:

Implement routes for the following actions using Gin:

- Create Project

- Get All Projects (Just showing id and name)

- Get Project by ID (Show all the available data of it, including engineers and tasks)

- Create User

- Get All Users (Just showing id and name)

- Get User by ID (Showing the current working project)

- Reasign a User to a different project

- Create Task

### Controller:

Create handler functions for each route of the corresponding CRUD operation.

Use Gorm to perform database operations (create, read, update, delete) for each model.

### Documentation:

Rename the README file to instructions.md, then create a new README file with the following description:

Provide simple documentation for your API, specifying the available endpoints and the expected request/response format.

Include any necessary setup instructions for users who want to run your application.

### Extra:

- Implement cascading deletion. When deleting a Project, ensure that associated Engineers and Tasks are also deleted.
