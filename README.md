# Companies API

Welcome to the Companies API, a demonstration of modern backend development using Golang.

## Overview

This project showcases a RESTful API with essential CRUD operations for managing company data. It integrates seamlessly with MySQL for data storage and is containerized with Docker for easy deployment.

### Key Technologies

- **Golang 1.22:** Leveraging the power and efficiency of Golang for robust backend development.
- **MySQL:** A reliable relational database management system for persistent data storage.
- **🐳 Docker:** Simplifies deployment and ensures consistent environments across different platforms.

## Getting Started

### Running Locally

#### Using Docker Compose 🐳

Start both MySQL and the API service with a single command:

```bash
docker-compose up -d --build mysql api
```

#### Running Separately

Alternatively, build and run the API service and MySQL individually:

```bash
# Build the image
docker build -t c-api-image .

# Run the image with environment variables and network setup
docker run --env-file .env --network companies-stack -p 8080:8080 c-api-image
```

### Setting Up the API

1. **Configure MySQL Connection**

   Set the `MYSQL_DSN` environment variable to connect the API with MySQL. Example:

   ```bash
   export MYSQL_DSN="root:password@tcp(mysql:3306)/the_company_db?parseTime=true&sql_mode=NO_ZERO_DATE"
   ```

2. **Adjust Build Configuration**

   Modify `GOOS` and `GOARCH` in the `Makefile` according to your local machine architecture.

3. **Build and Run**

   Use the Makefile to build the API and execute the compiled executable:

   ```bash
   make build
   ./companies-api
   ```

### Additional Resources

Explore the API functionalities using the Postman collection available in the [docs](docs) folder.

### TODO:

- Add API documentation
- Add Unit Tests

## ☎️ Get in Touch

I'm always open to discussions, collaborations, and feedback. If you have any questions or just want to connect, feel free to reach out!

- **Email:** gbouramas@gmail.com
- **LinkedIn:** [Giannis Bouramas](https://www.linkedin.com/in/bouramas)
