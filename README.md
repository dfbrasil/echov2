# echov2
Api Rest in GoLang - Echo Framework

# Echo API

This is a simple application that allows you to perform CRUD operations on a list of users.

## Usage

To run the application, you should have Go installed on your machine. Then, follow these steps:

1. Clone the repository to your machine:
   `git clone https://github.com/ruangustavo/echo-api.git`

2. Navigate to the repository directory: `cd echo-api`

3. Run the application: `go run main.go`

4. Use your favorite API client (such as Postman or cURL) to send requests to the API endpoints listed above.

Enjoy! If you have any questions or issues, please feel free to reach out to us.

## Endpoints

| Endpoint                | Method | Description                                                                                                                 |
| ----------------------- | ------ | --------------------------------------------------------------------------------------------------------------------------- |
| `/api/users`            | GET    | Returns a list of all users in the system.                                                                                  |
| `/api/users/:id`        | GET    | Returns the user with the specified ID.                                                                                     |
| `/api/users/create`     | POST   | Allows you to create a new user. You should include a JSON payload with the user's details in the request body.             |
| `/api/users/:id/update` | PUT    | Allows you to update an existing user. You should include a JSON payload with the updated user details in the request body. |
| `/api/users/:id/delete` | DELETE | Allows you to delete a user with the specified ID.                                                                          |

## Future Updates

Here are some potential updates we may make to the application in the future:

- [ ] Add support for user authentication and authorization
- [ ] Implement pagination for the `/api/users` endpoint to limit the number of results returned
- [ ] Add support for searching and filtering users by various criteria, such as name or email address
- [ ] Improve error handling and provide more informative error messages to clients
- [ ] Implement unit and integration tests to ensure the application works as expected

https://academicoifrnedu-my.sharepoint.com/:p:/g/personal/d_brasil_academico_ifrn_edu_br/EXzmuc0zKHJFtaYhhuDdackBwTePxcmB3z4m7gM9Oc3XWg?e=fDU6XO
