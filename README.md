# echov2
Api Rest in GoLang - Echo Framework

# Echo API

This is a simple application that allows you to perform CRUD operations on a list of books.

## Usage

To run the application, you should have Go installed on your machine. Then, follow these steps:

1. Clone the repository to your machine:
   `git clone https://github.com/dfbrasil/echov2.git`

2. Navigate to the repository directory: `cd echov2`

3. Run the application: `go run main.go`

4. Use your favorite API client (such as Postman or cURL) to send requests to the API endpoints listed above.

Enjoy! If you have any questions or issues, please feel free to reach out to us.

## Endpoints

| Endpoint                | Method | Description                                                                                                                 |
| ----------------------- | ------ | --------------------------------------------------------------------------------------------------------------------------- |
| `/users`            | GET    | Returns a list of all users in the system.                                                                                  |
| `/users/:id`        | GET    | Returns the user with the specified ID.                                                                                     |
| `/users`     | POST   | Allows you to create a new user. You should include a JSON payload with the user's details in the request body.             |
| `/users/:id` | PUT    | Allows you to update an existing user. You should include a JSON payload with the updated user details in the request body. |
| `/api/:id` | DELETE | Allows you to delete a user with the specified ID.                                                                          |

## PowerPoint Apresentation about Echo/Best Practices (PT-BR)

https://academicoifrnedu-my.sharepoint.com/:p:/g/personal/d_brasil_academico_ifrn_edu_br/EXzmuc0zKHJFtaYhhuDdackBwTePxcmB3z4m7gM9Oc3XWg?e=fDU6XO
