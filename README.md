# Shopping API Backend

This is a simple API to manage shopping items with PostgreSQL, built with Go and the Gin framework. It includes functionality to add, update, delete, and retrieve shopping items. This backend is Dockerized for easy development and deployment.

## Prerequisites

Before you start, make sure you have the following installed:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)

## Setup Instructions

### 1. Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/your-username/shopping-api-backend-go.git
cd shopping-api-backend-go
```

### 2. Configure the Environment Variables

The application uses environment variables for configuration. You need to create a `.env` file to store your local or Codespace environment settings.

#### Generate .env for Local Development:

Copy the contents of the `.env.sample` file to create your `.env` file:

```bash
cp .env.sample .env
```

Open the newly created `.env` file and update the values for the following variables:

- `POSTGRES_HOST`: The address of your PostgreSQL database (e.g., localhost or db in Docker)
- `POSTGRES_PORT`: The port your PostgreSQL database is running on (default is 5432)
- `POSTGRES_USER`: The PostgreSQL username (default is admin)
- `POSTGRES_PASSWORD`: The PostgreSQL password (set this to whatever you want)
- `POSTGRES_DB`: The PostgreSQL database name (e.g., shoppingdb)
- `CODESPACE_NAME`: This should be set if you are using GitHub Codespaces. This value is set automatically by GitHub, but you can define it manually if necessary
- `GITHUB_COSPACE_DOMAIN`: This is the domain part used when in GitHub Codespaces (usually app.github.dev)

#### Set Environment Variables for GitHub Codespaces:

If you are working inside GitHub Codespaces, GitHub will automatically set the `CODESPACE_NAME` and `GITHUB_COSPACE_DOMAIN` environment variables. You do not need to manually set these values unless you want to override them.

Example:

```bash
# Example for GitHub Codespaces (no need to change these variables usually)
CODESPACE_NAME=your-codespace-name
GITHUB_COSPACE_DOMAIN=app.github.dev
```

### 3. Run the Application with Docker Compose

Once your `.env` file is set up, you can start the application with Docker Compose. This will start all necessary services, including the backend, frontend, and PostgreSQL database.

```bash
docker-compose up --build
```

This will:
- Build the Docker images for the backend and frontend services
- Start the services and make them available on your local machine or GitHub Codespace

### 4. Access the Application

- Backend API will be available at `http://localhost:8080` (or via your GitHub Codespace URL in the format `https://<codespace-name>-8080.app.github.dev`)

- Swagger Documentation will be available at the following URL:  
  `{BASE_URL}/swagger/index.html`

  Replace `{BASE_URL}` with the appropriate base URL for your environment:
  - For local development, use `http://localhost:8080`.
  - For GitHub Codespaces, use the URL in the format `https://<codespace-name>-8080.app.github.dev`.

- Frontend will be available at `http://localhost:5000`

### 5. Testing the API

You can test the API using curl, Postman, or any HTTP client. Example:

```bash
curl -X 'GET' 'http://localhost:8080/api/shoppingItems' -H 'accept: application/json'
```

### Stopping the Application

```bash
docker-compose down
```

## Troubleshooting

- If you encounter issues connecting to the database, make sure the environment variables in `.env` are set correctly, especially the `POSTGRES_HOST` and `POSTGRES_PORT`
- Make sure your Docker Compose environment is correctly set up, and that all required services (backend, db, frontend) are running