# Project Title

Golang CRUD with MongoDB

## Description

This is a simple CRUD (Create, Read, Update, Delete) application written in Golang using MongoDB as the database. It demonstrates how to perform basic database operations using the official MongoDB driver for Golang.

## Prerequisites

Before running this application, ensure that you have the following prerequisites installed:

- Go programming language (version 1.16 or later)
- MongoDB (running locally or accessible via connection string)

## Installation

1. Clone the repository:

git clone https://github.com/your-username/golang-mongodb-crud.git


2. Change into the project directory:

cd golang-mongodb-crud


3. Install the dependencies:

go mod download


4. Update the MongoDB connection URI:

Open the `main.go` file and update the `uri` variable inside the `main` function with your MongoDB connection string.

## Usage

To run the application, use the following command:

go run main.go


The application will connect to the MongoDB server using the provided connection string and perform CRUD operations on the "persons" collection in the "mydb" database.

## License

This project is licensed under the [MIT License](LICENSE).

## Contributing

Contributions are welcome! If you find any issues or want to extend the functionality of this project, feel free to open a pull request.

## Acknowledgments

- The official [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) documentation was referenced while creating this project.
