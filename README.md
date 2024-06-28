# Sudoku API

This is a simple Sudoku API built with Go and Gin. It provides an endpoint to generate a new Sudoku game with a unique solution.

## Requirements

- Go 1.16 or higher
- Gin Web Framework

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/DinohRatiarisandy/Sudoku.git

2. Navigate to the project directory
3. Install the dependencies:
   ```sh
   go mod tidy

### Run the server:
go run main.go

## Endpoints
Generates a new valid Sudoku game with a unique solution.

URL: `/newsudoku`

Method: `GET`

### Query Parameters:

difficulty (optional): Specifies the difficulty level of the Sudoku game. Possible values are `easy`, `medium`, `hard`, `evil`. Default is `easy`.
</br>
e.g: `localhost:8080/newsudoku?difficulty=evil`
</br></br>
If you want to test this API, it is available online at: [elsombrero.pro:1010/newsudoku](http://elsombrero.pro:1010/newsudoku)


 

