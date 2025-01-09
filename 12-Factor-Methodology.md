# **12-Factor App Methodology**

_This document describes the **12-factor app** methodology applied to the project, including explanations of each principle and how it has been implemented. The goal is to ensure that our application is maintainable, scalable, and follows best practices for cloud-native applications._

---

## 1. **Codebase** 

### Principle:
There should be a **single codebase** tracked in version control, with many deploys.

### Applied:
The project is maintained in a **single Git repository**. The code is versioned using **Git**.

---

## 2. **Dependencies**

### Principle:
**Explicitly declare** and isolate dependencies.

### Applied:
The project uses **Go Modules** (`go.mod` and `go.sum`) for dependency management, ensuring that all dependencies are clearly specified.

---

## 3. **Config**

### Principle:
Store configuration in the **environment** (as environment variables).

### Applied:
All sensitive configuration, such as database credentials and API keys, are stored in the **.env** file and read as environment variables at runtime.

```. env
# Example .env file

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=your_user
POSTGRES_PASSWORD=your_password
POSTGRES_DB=your_db
```

## 4. **Backing Services**

### **Principle:**
Treat backing services (databases, caches, etc.) as **attached resources**.

### **Applied:**
PostgreSQL is used as the backing **database service**, and the connection details are configured using **environment variables**.

---

## 5. **Build, Release, Run**

### **Principle:**
Strictly separate the **build**, **release**, and **run** stages.

### **Applied:**
The project is **containerized** using **Docker**. It uses a **Dockerfile** to build the application image and the **docker-compose.yml** file to manage the containerized services.

Implemented **CI/CD pipeline** to separate build, release, and run stages.

---

## 6. **Processes**

### **Principle:**
Execute the app as one or more **stateless processes**.

### **Applied:**
The app is **stateless** as it does not maintain any state between requests; data is stored in a **PostgreSQL database** instead.

---

## 7. **Port Binding**

### **Principle:**
Export services via **port binding**.

### **Applied:**
The app binds to **port 8080** using the **Dockerfile** and **docker-compose.yml** to make the service available externally.

---

## 8. **Concurrency**

### **Principle:**
Scale out via the **process model**.

### **Applied:**
The app is designed to be **horizontally scalable**. By using **Docker** and **Kubernetes** (in deployment), we can scale the number of replicas of the app.

---

## 9. **Disposability**

### **Principle:**
Maximize robustness with **fast startup** and **graceful shutdown**.

### **Applied:**
The app is designed to **shut down gracefully** and release resources (e.g., database connections) when it receives termination signals.

### ** Fix Needed:**
Have to ensure that graceful shutdown logic is implemented using **Go’s os.Signal** package (e.g., `SIGTERM`, `SIGINT`).
 
---

# Dev/Prod Parity

## Principle
Keep **development**, **staging**, and **production** environments as similar as possible.

## Applied
* **Docker** for local development, **Kubernetes** for production/staging.
* Use `.env` files for environment variables.
* **Docker Compose** loads environment files based on the `ENV` variable (defaults to `.env.development` if not set).

## How It Works

### Docker Compose
```yaml
env_file:
  - .env.${ENV:-development}
```
This keeps environments consistent and helps with smooth transitions between local and production setups.

### 11. Logs :ledger:

**Principle:** Treat logs as event streams.

**Applied:**  
The app logs events, especially errors and connection status messages, using the **Go log package**. Logs are crucial for debugging and monitoring. These logs are generated within the application, and we can inspect the logs by viewing the container logs.

---


### 12. Admin Processes

**Principle:** Run administrative/management tasks as one-off processes.

**Fix needed:** Administrative tasks such as database migrations or backups can be managed separately. Have to ensure to handle one-off tasks like migrations through Kubernetes jobs or Docker containers.

---
