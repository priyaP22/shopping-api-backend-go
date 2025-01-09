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
# Pending Improvements

## Kubernetes Configuration
* Implement proper Secrets and ConfigMaps management
* Replace direct environment variable declarations in deployment manifests


## Migrations

### 1. **Initial Setup**
Ensure that **Goose** is installed in your development environment.

```bash
go install github.com/pressly/goose/v3@latest

```
### 2. **Creating Migrations**

To create a new migration (e.g., for adding a table or modifying the schema):

Run the following command to create a migration file in the `migrations` folder:

```bash
goose create <migration_name> sql
```
This will generate a new migration file with a timestamp, such as 20250109112606_create_shopping_items_table.sql.

In the generated migration file, add your up and down SQL statements:


```bash
+goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS shopping_items (
    name TEXT PRIMARY KEY,
    amount INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shopping_items;
-- +goose StatementEnd
```



### 3. Applying Migrations
To apply the migrations to the database:

Set up your database connection and ensure it’s accessible (check the credentials in .env).

Run the migration command:

``` bash
goose postgres "host=<db_host> port=<db_port> user=<db_user> password=<db_password> dbname=<db_name> sslmode=disable" up -dir=migrations
```
This will apply any unapplied migrations to the database.

### 4. Checking Migration Status
To check the current status of the database migrations, including which migrations have been applied, use the following command:

``` bash

goose postgres "host=<db_host> port=<db_port> user=<db_user> password=<db_password> dbname=<db_name> sslmode=disable" status -dir=migrations


### 5. Rolling Back Migrations
If you need to roll back the last migration, run:

``` bash
goose postgres "host=<db_host> port=<db_port> user=<db_user> password=<db_password> dbname=<db_name> sslmode=disable" down -dir=migrations

```

To rollback to a specific version, use:

``` bash

goose postgres "host=<db_host> port=<db_port> user=<db_user> password=<db_password> dbname=<db_name> sslmode=disable" down-to <version> -dir=migrations
```

### 6. Versioning of Migrations
Goose automatically keeps track of which migrations have been applied by maintaining a table (default: goose_db_version) in your database. You can customize the table name with the -table flag.



