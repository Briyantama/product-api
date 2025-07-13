# 🛠️ Backend API - Vendor & Product Management

## 📚 Tech Stack

- **Go** (Golang)
- **Gin** (HTTP Web Framework)
- **PostgreSQL** (Relational Database)
- **JWT** (JSON Web Token for Authentication)
- **GORM** (ORM for Golang)
- **Postman** (for API testing)

---

## ⚙️ Setup Instructions

1. **Clone the repository**
   ```bash
   git clone https://github.com/Briyantama/product-api.git
   cd product-api
   cp .env.example .env
   go mod tidy
   go run main.go


## 📫 Postman Collection

- The Postman collection file is available in the repository:
- 📁 `docs/postman_collection.json`

Import this collection into Postman to test the available API endpoints.

---

## 📌 API Endpoints

### 🔐 Auth

| Method | Endpoint         | Description         |
|--------|------------------|---------------------|
| POST   | `/auth/register` | Register new user   |
| POST   | `/auth/login`    | Login and get JWT   |

---

### 🧾 Vendor (JWT Protected)

| Method | Endpoint     | Description                   |
|--------|--------------|-------------------------------|
| POST   | `/vendors`   | Register new vendor (by user) |
| GET    | `/vendors`   | Get vendors by user ID        |

---

### 📦 Product (JWT Protected)

| Method | Endpoint              | Description                |
|--------|-----------------------|----------------------------|
| POST   | `/products`           | Create product             |
| GET    | `/products/user`      | Get products by user ID    |
| GET    | `/products/vendor`    | Get products by vendor ID  |
| PUT    | `/products/:id`       | Update product by ID       |
| DELETE | `/products/:id`       | Delete product by ID       |

---