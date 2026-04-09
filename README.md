# Sociomile Technical Test

This project is a multi-tenant conversation and ticketing system.

## How to Run the Full Stack (Recommended)

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


## Default Credentials (Seeded)

- **Tenant 1 (Sociomile Enterprise)**
  - Admin: `admin@sociomile.com` / `password123`
  - Agent: `agent@sociomile.com` / `password123`
- **Tenant 2 (Kiki Tech Solutions)**
  - Admin: `admin2@tech.com` / `password123`
  - Agent: `agent@tech.com` / `password123`

---


## Not yet implemented

- Unit test
