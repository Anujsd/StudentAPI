# CRUD REST API Made Using Go

This REST API is built using Go and provides endpoints for basic CRUD operations on student records.

## Run locally

```
docker compose up
```

## Endpoints

| Endpoint          | Method | Description                                                                                                                           |
| ----------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------- |
| `/v1/students`    | GET    | Retrieves a list of all students.                                                                                                     |
| `/v1/student/:id` | GET    | Retrieves information about a specific student identified by ID.                                                                      |
| `/v1/student`     | POST   | Adds a new student to the database. Requires providing the student's details in the request body.                                     |
| `/v1/student/:id` | PUT    | Updates information about an existing student identified by ID. Requires providing the updated student's details in the request body. |
| `/v1/student/:id` | DELETE | Deletes a student record from the database based on ID.                                                                               |

## Usage

To use this API, make HTTP requests to the provided endpoints using appropriate HTTP methods (GET, POST, PUT, DELETE).

For example:

- To retrieve all students, send a GET request to `/v1/students`.
- To add a new student, send a POST request to `/v1/student` with the student's details in the request body.

## Setup Without Docker

1. Clone this repository to your local machine.
2. If you haven't already, install Go.
3. Use `go get .` to install the necessary dependencies.
4. Obtain the PostgreSQL connection string and create an environment variable named `CONNECTION_STRING`.
5. Build and run the application.

## Dependencies

- [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) - Go web framework used for building the REST API.

## License

This project is licensed under the [MIT License](LICENSE).
