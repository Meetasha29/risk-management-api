# Risky Plumber API

## Overview
This is a simple Risk Management API that allows users to:
- Create a new risk (`POST /v1/risks`)
- Retrieve all risks (`GET /v1/risks`)
- Retrieve a specific risk by ID (`GET /v1/risks/{id}`)

The API is built using Go and uses an in-memory data store for simplicity.

---

## 🚀 How to Run

### **1️⃣ Prerequisites**
- Install [Go](https://golang.org/dl/)
- Clone the repository:
  ```sh
  git clone <repository-url>
  cd risky_plumber
  ```
- Install dependencies:
  ```sh
  go mod tidy
  ```

### **2️⃣ Run the Server**
```sh
  go run main.go
```
The server will start on **`http://localhost:8080`**.

---

## 📌 API Endpoints

| **Method** | **Endpoint** | **Description** |
|-----------|-------------|----------------|
| `POST` | `/v1/risks` | Create a new risk |
| `GET` | `/v1/risks` | Get all risks |
| `GET` | `/v1/risks/{id}` | Get a risk by ID |

---

## 🛠️ Postman Collection
1. Create Risk (`POST /v1/risks`) - 
curl --location 'http://localhost:8080/v1/risks' \
--header 'Content-Type: application/json' \
--data '{
    "state": "OPEN",
    "title": "risk3",
    "desc": "risk1-2"
}'
2. Get All Risks (`GET /v1/risks`)
curl --location 'http://localhost:8080/v1/risks'
3. Get Risk by ID (`GET /v1/risks/{id}`)
curl --location 'http://localhost:8080/v1/risk/447e18f8-b8e1-4c20-a5c3-bbb4adc1cb7d'

---

## 📝 For Reviewers

1. **Risk Uniqueness**: Risks are **not unique on any parameter** other than the auto-generated UUID, as per the requirements.
2. **Mutex Usage**: Since UUIDs are unique, **no `sync.Mutex` lock was needed** for concurrent writes.
3. **Decoupled Interfaces**: The code is structured with **interfaces to allow extensibility and separation of concerns**.
4. **SQL Implementation**: Instead of the in-memory implementation, it would be easy to add a SQL-based implementation by simply replacing the repository layer with a database-backed version.

---

## 🔥 Future Enhancements
- Add persistence using PostgreSQL.
- Implement API authentication.
- Extend risk models with additional metadata.

---

### 🎯 **Contact & Contributions**
Feel free to open an issue or contribute via pull requests!
