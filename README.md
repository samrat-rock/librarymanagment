# 📚 Library Management System - Golang

A simple and efficient Library Management System built using Go (Golang). This project provides basic CRUD functionality for managing books, users, and book issuances in a library setting. It is designed as a RESTful API and can be used as the backend for a full-stack library system.

---

## 🚀 Features

- 📖 Add, update, delete, and list books
- 👥 Register and manage library users
- 🔁 Issue and return books
- ✅ Track book availability
- 🔒 Environment-based DB configuration
- ❗ Basic validation and error handling
- 🧱 Modular code structure using Go's standard packages

---

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Framework:** net/http (standard library)
- **Database:** PostgreSQL / MySQL / SQLite (configurable)
- **ORM:** GORM
- **API Type:** RESTful JSON-based
- **Deployment:** Docker (optional)

---

## 📁 Project Structure
library-management/
│
├── controllers/ # HTTP handler functions
├── models/ # Data models for books, users, and transactions
├── routes/ # API route definitions
├── database/ # DB connection and migrations
├── utils/ # Helper functions and utilities
├── .env # Environment variables
├── main.go # Entry point of the application
└── go.mod # Go module file


---

## 🧩 Environment Variables

Create a `.env` file in the root directory:

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=librarydb
DB_SSLMODE=disable
PORT=8080



Update your database logic to load these values (using something like `github.com/joho/godotenv`).

---

## 🧱 Sample Database Schema (PostgreSQL)

```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE books (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  author VARCHAR(100),
  available BOOLEAN DEFAULT TRUE
);

CREATE TABLE issuances (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id),
  book_id INTEGER REFERENCES books(id),
  issue_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  return_date TIMESTAMP
);

 Local Setup (Without Docker)
Clone the repository

bash
Copy
Edit
git clone https://github.com/yourusername/library-management.git
cd library-management


Install dependencies

bash
Copy
Edit
go mod tidy
go run main.go

🧪 Testing
Use the following tools to test the API:

Postman

curl

Thunder Client (VSCode Extension)
