# Gopportunities API

The **Gopportunities API** is a RESTful web service designed to manage job opportunities, resumes, and newsletters. Built with Go and the Gin framework, it integrates Swagger for API documentation and SQLite for data persistence.

### Features
- **Job Opportunities**: Create, update, fetch, and delete job openings.
- **Resumes Support**: Upload, manage, and link resumes to job applications. 
- **Newsletter**: Manage newsletter subscribers and send periodic updates.
- **Authentication**: Secure API endpoints with authentication middleware.
- **Swagger Documentation**: Comprehensive API documentation generated with `swag`.

### Prerequisites
- Go 1.22 or later
- Docker (optional, for containerized deployment)
- SQLite

### Installation
#### Clone repo
```
git clone https://github.com/pedropassos06/gopportunities.git
cd gopportunities
```

#### Install dependencies
```
go mod download
```

#### Set up your own env variables
Copy the example file into the actual `.env` file and fill it out
```
cp .env.example .env
```

### Usage
#### Run locally
```
make run
```

#### Run with swagger doc
```
make run-with-docs
```

Access the API at: http://localhost:8080

Swagger documentation available at: http://localhost:8080/swagger/index.html

### Docker
#### Build the image
```
docker build -t gopportunities-api .
```

#### Run the container
```
docker run -p 8080:8080 --env-file .env gopportunities-api
```

### Overall Structure
```
.
├── cmd/
│   └── main.go           # Main entry point
├── config/               # Configuration files and logger
├── openings/             # Job openings module
├── resume/               # Resume module
├── middleware/           # Auth middleware
├── newsletter/           # Newsletter module
├── schemas/              # Data schemas (models)
├── router/               # API routes
├── utils/                # Utility functions
├── docs/                 # Swagger documentation
├── Makefile              # Task automation
├── Dockerfile            # Docker configuration
├── .env.example          # Example environment variables
└── README.md             # Project documentation
```

### Want to contribute?
1. Fork the repository.
2. Create a new feature branch.
3. Commit your changes.
4. Open a pull request.

### License
This project is licensed under the MIT License. See the [LICENSE](https://opensource.org/license/mit) file for details.


