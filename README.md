# Roman Numeral Converter API

This is a simple API that converts a range of decimal numbers to Roman numerals. It returns a list of ordered Romand numbers from the smallest to the biggest (ascending).

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.16 or higher
- Docker

### Building the Project

1. Clone the repository:

```bash
git clone https://github.com/busymo/roman-numeral-converter.git
```

2. Go to the cloned repository
```bash
cd roman-numeral-converter
```

3. Build the Docker image:

```bash
docker build -t roman-numeral-converter .
```

4. Run the Docker image:

```bash
docker run -p 8080:8080 roman-numeral-converter
```

The API will be available at http://localhost:8080/convert. 

Swagger is available at: http://localhost:8080/swagger/index.html

5. Docker compose

Alternatively, one can use docker compose to build and run the project, as following:

```bash
docker-compose up
```

## Using the API
The API has a single endpoint, /convert, which accepts a range query parameter in the format 'from-to'. 

For example, to convert the numbers 1 to 10 to Roman numerals, you would send a GET request to:

http://localhost:8080/convert?range=1-10.

### Running the Tests
To run the tests, use the go test command:

```bash
go test -v ./...
```

## Future Work

### Improvements:
- **Error Handling:** Provide more detailed and user-friendly error messages.
- **Input Validation**: Implement more robust input validation.
- **Logging**: Add structured logging for better traceability.
- **Rate Limiting**: Prevent abuse by implementing rate limiting.

### Production (ready):
- **Security**: Add authentication and authorization mechanisms.
- **Scalability**: Set up load balancing and auto-scaling.
- **CI/CD**: Integrate with a CI/CD pipeline for automated testing and deployment.
- **Monitoring**: Implement monitoring and alerting systems.