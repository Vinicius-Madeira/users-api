# Users API

This application tries to achieve a simple Web API using **Go** that fetches, create, update, or delete a user.

###### Disclaimer: Intended for study purposes only.

## About

Built according to the MVC architecture, making use of the [Gin](https://github.com/gin-gonic/gin?tab=readme-ov-file) framework to run the web server.

- **Title**: Users API
- **Version**: 1.0
- **Author**: Vinicius Madeira
- **Host**: localhost:8080

## Technical Features

- Integration with a MongoDB to store information about the users.
- JWT Token for validation of the incoming requests for most endpoints.
- Test suites for all layers of the application.
- Endpoint documentation on Swagger using the [Swag](https://github.com/swaggo/swag) library.
- Docker compose for easy setup.

## Prerequisites

Before getting started, make sure you have the following prerequisites installed on your system:

- [Go](https://golang.org/dl/): The Go programming language.
- [Docker](https://www.docker.com/get-started): Docker is required if you wish to run the application in a container.

## Installation

Follow the steps below to install the project in your development environment:

1. **Clone the repository:**

   ```
   git clone https://github.com/Vinicius-Madeira/users-api
   ```

2. **Navigate to the project directory:**

   ```
   cd users-api
   ```

3. **Build and run the application using Docker Compose:**

   ```
   docker compose up
   ```

## Running the Application

After installation, you can run the **Users API** with the following command (if you want to run it directly with Golang):

```
docker container run --name users-api-db -p 27017:27017 -d mongo

go run main.go
```

The application will be accessible at `http://localhost:8080`.

## Testing the Application

If you prefer, after running the project, visit: http://localhost:8080/swagger/index.html# to see and test all the route contracts.

The **Users API** offers REST endpoints for creating, listing, updating, and deleting users. You can use tools like [curl](https://curl.se/) or [Postman](https://www.postman.com/) to test the endpoints. Here are some `curl` command examples for testing the endpoints:

- **Create a user:**

  ```
  curl -X POST -H "Content-Type: application/json" -d '{"name": "John", "email": "john@example.com", "age": 30, "password": "password$#@$#323"}' http://localhost:8080/createUser
  ```

- **Update a user:**

  ```
  curl -X PUT -H "Content-Type: application/json" -d '{"name": "John Brighton"}' http://localhost:8080/updateUser/{userId}
  ```

- **Delete a user:**

  ```
  curl -X DELETE http://localhost:8080/deleteUser/{userID}
  ```

Remember to adjust the commands according to your needs and requirements.

## Data Models

### request.UserAuth
Structure containing the necessary fields for user login.

- `email` (string, required): The user's email (must be a valid email address).
- `password` (string, required): The user's password (must be at least 6 characters and contain at least one of the characters: !@#$%*).

### request.UserRequest
Structure containing the required fields for creating a new user.

- `age` (integer, required): The user's age (must be between 1 and 140).
- `email` (string, required): The user's email (must be a valid email address).
- `name` (string, required): The user's name (must be at least 4 characters and at most 100 characters).
- `password` (string, required): The user's password (must be at least 6 characters and contain at least one of the characters: !@#$%*).

### request.UserUpdateRequest
Structure containing fields to update user information.

- `age` (integer, required): The user's age (must be between 1 and 140).
- `name` (string, required): The user's name (must be at least 4 characters and at most 100 characters).

### response.UserResponse
Response structure containing user information.

- `age` (integer): The user's age.
- `email` (string): The user's email.
- `id` (string): The user's unique ID.
- `name` (string): The user's name.

### rest_err.Causes
Structure representing the causes of an error.

- `field` (string): The field associated with the error cause.
- `message` (string): Error message describing the cause.

### rest_err.RestErr
Structure describing why an error occurred.

- `causes` (array of rest_err.Causes): Error causes.
- `code` (integer): Error code.
- `error` (string): Error description.
- `message` (string): Error message.

## Endpoints

### Note

- For authentication, you should include the access token in the `Authorization` header as "Bearer <Insert access token here>" for protected endpoints.

The API offers the following endpoints:

1. **POST /createUser**
    - Description: Create a new user with the provided user information.
    - Parameters:
        - `userRequest` (body, required): User information for registration.
    - Responses:
        - 200: OK (User created successfully)
        - 400: Bad Request (Request error)
        - 500: Internal Server Error (Internal server error)

2. **DELETE /deleteUser/{userId}**
    - Description: Delete a user based on the provided ID parameter.
    - Parameters:
        - `userId` (path, required): ID of the user to be deleted.
    - Responses:
        - 200: OK (User deleted successfully)
        - 400: Bad Request (Request error)
        - 500: Internal Server Error (Internal server error)

3. **GET /getUserByEmail/{userEmail}**
    - Description: Retrieve user details based on the email provided as a parameter.
    - Parameters:
        - `userEmail` (path, required): Email of the user to be retrieved.
    - Responses:
        - 200: User information retrieved successfully
        - 400: Error: Invalid user ID
        - 404: User not found

4. **GET /getUserById/{userId}**
    - Description: Retrieve user details based on the user ID provided as a parameter.
    - Parameters:
        - `userId` (path, required): ID of the user to be retrieved.
    - Responses:
        - 200: User information retrieved successfully
        - 400: Error: Invalid user ID
        - 404: User not found

5. **POST /auth**
    - Description: Allow a user to log in and receive an authentication token.
    - Parameters:
        - `userLogin` (body, required): User login credentials.
    - Responses:
        - 200: Login successful, authentication token provided
        - 403: Error: Invalid login credentials

6. **PUT /updateUser/{userId}**
    - Description: Update user details based on the ID provided as a parameter.
    - Parameters:
        - `userId` (path, required): ID of the user to be updated.
        - `userRequest` (body, required): User information for update.
    - Responses:
        - 200: OK (User updated successfully)
        - 400: Bad Request (Request error)
        - 500: Internal Server Error (Internal server error)

We hope this Swagger documentation has been helpful in understanding and interacting with the API of the **Users API** project in Go. If you have any questions or need additional support, please don't hesitate to reach out. Enjoy using the API!