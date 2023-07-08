# Golang CRUD with MongoDB

This is a simple CRUD (Create, Read, Update, Delete) application written in Golang using MongoDB as the database. It demonstrates how to perform basic database operations using the official MongoDB driver for Golang.

## Prerequisites

Before running this application, ensure that you have the following prerequisites installed:

- Go programming language (version 1.16 or later)
- MongoDB (running locally or accessible via connection string)

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/your-username/golang-mongodb-crud.git
Change into the project directory:

shell
Copy code
cd golang-mongodb-crud
Install the dependencies:

shell
Copy code
go mod download
Update the MongoDB connection URI:

Open the main.go file and update the uri variable inside the main function with your MongoDB connection string.

Usage
To run the application, use the following command:

shell
Copy code
go run main.go
The application will connect to the MongoDB server using the provided connection string and perform CRUD operations on the "persons" collection in the "mydb" database.

License
This project is licensed under the MIT License.

Contributing
Contributions are welcome! If you find any issues or want to extend the functionality of this project, feel free to open a pull request.

