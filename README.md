

<p align="center">
  <a href="https://go-eqn-solver.azurewebsites.net">
    <img src="http://cl.ly/a4c860f3a890/calculator.svg" alt="Eqn solver logo" width=72 height=72>
  </a>

  <h3 align="center">Golang Equation Solver</h3>

  <p align="center">
    A simple microservice to solve linear equations, built in Golang.
    <br>
    <a href="https://go-eqn-solver.azurewebsites.net/docs"><strong>Explore go-eqn-solver docs »</strong></a>
    <br>
    <a href="https://go-eqn-solver.azurewebsites.net"><strong>Demo it now »</strong></a>
  </p>

</p>

<br>

### Setup

1. Install `docker`
2. Clone the repository and `cd` into the directory
3. Build the image through  `docker build -t go-eqn-solver .`
4. Run using  `docker run -d -p 8080:8080 go-eqn-solver`

### Tests

1.  `cd` into the project directory
2. Run tests using `go test ./...`

### Usage

##### Webpage

1. Browse to http://localhost:8080 to access the front-end equation solver

**RESTFul API**

###### HTTP Request

`POST https://go-eqn-solver.azurewebsites.net/api/solve`

###### JSON Request

Key | Description | Example
--------- | ----------- | --------
eqn | A linear equation with one unknown *x* | `2x=4`


###### JSON Response
Given a valid request, here are the possible responses. The response is returned as JSON with a single `result` field.

Result | Description
--------- | -----------
x={number} | The equation has a single solution, `number`, rounded to 8 decimal places.
Infinite solutions | The equation has infinite solutions
No solution | The equation is unsolvable