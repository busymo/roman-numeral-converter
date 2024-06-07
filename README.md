# Roman Numeral Converter API

This is a simple API that converts a range of numbers to Roman numerals.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.16 or higher
- Docker

### Building the Project

1. Clone the repository:

```bash
git clone https://github.com/username/roman-numeral-converter.git
cd roman-numeral-converter

2. Build the Docker image:

```bash
docker build -t roman-numeral-converter .

3. Run the Docker image:

```bash
docker run -p 8080:8080 roman-numeral-converter

The application will be available at http://localhost:8080.

Using the API
The API has a single endpoint, /convert, which accepts a range query parameter in the format 'from-to'. For example, to convert the numbers 1 to 10 to Roman numerals, you would send a GET request to http://localhost:8080/convert?range=1-10.

Running the Tests
To run the tests, use the go test command:

```bash
go test ./...