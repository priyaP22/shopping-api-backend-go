# Shopping API Backend

A simple PostgreSQL-based shopping items API built with Go and Gin framework. Features include CRUD operations for shopping items, with both Docker and Kubernetes deployment options.

## Related Documentation
- [12-Factor Application Principles Documentation](https://github.com/priyaP22/shopping-api-backend-go/blob/main/12-Factor-Methodology.md)

## Prerequisites

- [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/) for container deployment
- Minikube and kubectl for Kubernetes deployment

## Docker Setup

### 1. Clone and Configure

```bash
# Clone the repository and navigate to project directory
git clone https://github.com/your-username/shopping-api-backend-go.git
cd shopping-api-backend-go

# Create environment file from template
cp .env.sample .env.development
```

### 2. Environment Configuration

Update the following variables in your `.env.${ENV}` file:

- `POSTGRES_HOST`: Database address (localhost or db)
- `POSTGRES_PORT`: Database port (default: 5432)
- `POSTGRES_USER`: Database username (default: admin)
- `POSTGRES_PASSWORD`: Your chosen password
- `POSTGRES_DB`: Database name (e.g., shoppingdb)

For GitHub Codespaces, `CODESPACE_NAME` and `GITHUB_COSPACE_DOMAIN` are automatically set.

I'll convert this into proper Markdown format with clear headers, code blocks, and formatting.

# Launch Application

### Running the Application

```bash
# Set the environment and build/start all services
ENV=development docker-compose up --build
```

### Supported Environments

* `ENV=development` - Loads `.env.development`
* `ENV=staging` - Loads `.env.staging`
* `ENV=production` - Loads `.env.production`

### Default Behavior

If no `ENV` is specified, it defaults to `.env.development`.

## Kubernetes Setup

### 1. Initialize Cluster

```bash
# Start a local Kubernetes cluster
minikube start

# Switch Kubernetes context to minikube
kubectl config use-context minikube
```

### 2. Deploy and Configure

```bash
# Apply all Kubernetes configurations from k8s directory
kubectl apply -f k8s/

# Set environment variable for the deployment if you are using codespace
kubectl set env deployment/shopping-api-backend-go CODESPACE_NAME=$CODESPACE_NAME

# Forward the service port to local machine
kubectl port-forward svc/shopping-api-backend-go-service 8080:8080
```

## Accessing the Application

The API is available at:
- Local: `http://localhost:8080`
- Codespaces: `https://<codespace-name>-8080.app.github.dev`

Swagger documentation: `{BASE_URL}/swagger/index.html`
Frontend interface: `http://localhost:5000`

## Troubleshooting

- Verify environment variables in `.env.${ENV}`
- Check service status:
  ```bash
  # List all running Docker containers and their status
  docker-compose ps

  # List all Kubernetes pods and their status
  kubectl get pods
  ```
- View logs:
  ```bash
  # View logs from all Docker services
  docker-compose logs

  # View logs from Kubernetes deployment
  kubectl logs deployment/shopping-api-backend-go
  ```

## Shutdown

```bash
# Stop and remove Docker containers, networks, and volumes
docker-compose down

# Remove Kubernetes resources
kubectl delete -f k8s/

# Stop the Minikube cluster
minikube stop
```
