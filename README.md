# Sociomile Technical Test

This project is a multi-tenant conversation and ticketing system.

## 🚀 How to Run the Full Stack (Recommended)

The easiest way to run the entire application (Database, Backend, and Frontend) is using Docker Compose.

1.  **Ensure you have Docker and Docker Compose installed.**
2.  **Run the following command in the root directory:**
    ```bash
    docker-compose up --build
    ```
3.  **Access the application:**
    - **Frontend**: [http://localhost:3000](http://localhost:3000)
    - **Backend API**: [http://localhost:8080](http://localhost:8080)
    - **Swagger UI**: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## 🛠 Running Locally (Development)

If you prefer to run the components individually:

### 1. Database
Ensure MySQL is running on port `3306` with the credentials found in `backend/internal/config/db.go` or `docker-compose.yml`. Use `backend/migrations/init.sql` to seed the database.

### 2. Backend (Go)
```bash
cd backend
go run cmd/api/main.go
```

### 3. Frontend (React)
```bash
cd frontend
npm install
npm run dev
```
The frontend will be available at [http://localhost:3000](http://localhost:3000).

---

## 🔑 Default Credentials (Seeded)

- **Tenant 1 (Sociomile Enterprise)**
  - Admin: `admin@sociomile.com` / `password123`
  - Agent: `agent@sociomile.com` / `password123`
- **Tenant 2 (Kiki Tech Solutions)**
  - Admin: `admin2@tech.com` / `password123`
  - Agent: `agent@tech.com` / `password123`
