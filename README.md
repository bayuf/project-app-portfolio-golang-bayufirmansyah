<div align="center">

# 🚀 Portfolio Web App — Golang

**A full-stack personal portfolio web application built with pure Go, featuring a dynamic admin panel, PostgreSQL integration, structured logging, and clean layered architecture.**

[![Go Version](https://img.shields.io/badge/Go-1.25.1-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-pgx%20v5-336791?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Chi Router](https://img.shields.io/badge/Router-Chi%20v5-FF6C37?style=for-the-badge)](https://github.com/go-chi/chi)
[![Zap Logger](https://img.shields.io/badge/Logger-Zap-yellow?style=for-the-badge)](https://github.com/uber-go/zap)
[![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)](LICENSE)

| Platform | Link |
|:---:|:---:|
| <img src="https://icons.iconarchive.com/icons/marcus-roberto/google-play/512/Google-Drive-icon.png" alt="Google Drive" width="20" /> Google Drive | [Watch Demo](https://drive.google.com/file/d/17BAywpqEsXdYbMi-wp6M17b2vVFv2aEn/view?usp=sharing) |

</div>

---

## 📋 Table of Contents

- [Overview](#-overview)
- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Project Structure](#-project-structure)
- [Architecture](#-architecture)
- [Prerequisites](#-prerequisites)
- [Installation & Setup](#-installation--setup)
- [Environment Variables](#-environment-variables)
- [Running the App](#-running-the-app)
- [Running Tests](#-running-tests)
- [API Routes](#-api-routes)

---

## 🌟 Overview

This is a **production-ready personal portfolio web application** written entirely in **pure Go** — no frameworks, no boilerplate generators. The app allows a developer to showcase their profile, skills, projects, and services through a dynamic, database-driven website with a protected admin dashboard for content management.

The application demonstrates real-world backend engineering concepts including:
- Clean **layered architecture** (Handler → Service → Repository)
- **PostgreSQL** connection pooling via `pgx/v5`
- Cookie-based **authentication & protected routes**
- Structured, rotating **logging** with Uber Zap + Lumberjack
- **Server-side HTML rendering** using Go's `text/template`
- **Input validation** with `go-playground/validator`
- **Unit testing** with `pgxmock` for database mocking

---

## ✨ Features

### Public Pages
| Page | Description |
|------|-------------|
| 🏠 **Home** | Landing page with hero section and summary |
| 👤 **About** | Developer profile, experience, and background |
| 🛠️ **Services** | Service offerings with icons and descriptions |
| 🗂️ **Portfolio** | Showcase of projects with links |
| 📩 **Contact** | Contact form that saves messages to the database |

### Admin Panel (Protected)
| Feature | Description |
|---------|-------------|
| 🔐 **Login** | Secure login with cookie-based session |
| 📊 **Dashboard** | Admin control panel for content management |
| 📥 **Download** | File download capability |

### Backend Capabilities
- 🔄 **Layered architecture** — clean separation of concerns
- 🧪 **Unit-tested repository layer** with mock database
- 📦 **Connection pooling** — efficient PostgreSQL pool via `pgx/pgxpool`
- 📋 **Structured logging** — rotating daily log files via Zap + Lumberjack
- ✅ **Input validation** — form data validated with `go-playground/validator`
- 🍪 **Cookie authentication** — HTTP-only cookie middleware for admin protection

---

## 🛠 Tech Stack

| Category | Technology |
|----------|-----------|
| **Language** | Go 1.25.1 |
| **Router** | [Chi v5](https://github.com/go-chi/chi) |
| **Database** | PostgreSQL (via [pgx v5](https://github.com/jackc/pgx)) |
| **Templating** | Go `text/template` (Server-Side Rendering) |
| **Logging** | [Uber Zap](https://github.com/uber-go/zap) + [Lumberjack](https://github.com/natefinch/lumberjack) |
| **Validation** | [go-playground/validator v10](https://github.com/go-playground/validator) |
| **Config** | [godotenv](https://github.com/joho/godotenv) |
| **Testing** | [testify](https://github.com/stretchr/testify) + [pgxmock v4](https://github.com/pashagolub/pgxmock) |
| **Security** | [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto) |

---

## 📁 Project Structure

```
project-app-portfolio-golang-bayufirmansyah/
│
├── main.go                  # Entry point — wires all layers together
├── go.mod / go.sum          # Module definition and dependencies
├── .env                     # Environment variables (not committed)
├── .gitignore
│
├── db/
│   └── conn.go              # PostgreSQL connection pool setup (pgxpool)
│
├── model/                   # Domain models / entities
│   ├── profile.go           # Developer profile (name, headline, about, etc.)
│   ├── project.go           # Portfolio project entries
│   ├── skill.go             # Skills with proficiency level
│   ├── offers.go            # Service offerings
│   ├── feedback.go          # User feedback model
│   ├── message.go           # Contact form messages
│   ├── nav.go               # Navigation items
│   ├── address.go           # Contact address info
│   ├── auth.go              # Auth / user credentials model
│   └── model.go             # Shared base model (ID, timestamps)
│
├── dto/                     # Data Transfer Objects (input/form payloads)
│   ├── auth.go              # Login DTO
│   └── message.go           # Contact message DTO
│
├── repository/              # Data access layer (SQL queries)
│   ├── repository.go        # Repository aggregator
│   ├── view_repository.go   # Fetch profile, projects, skills, offers
│   ├── auth_repository.go   # Auth/login queries
│   ├── input_repository.go  # Insert contact messages
│   └── view_repository_test.go  # Unit tests with pgxmock
│
├── services/                # Business logic layer
│
├── handler/                 # HTTP handlers (Controller layer)
│   ├── handler.go           # Handler aggregator
│   ├── home_handler.go      # Home page
│   ├── about_handler.go     # About page
│   ├── services_handler.go  # Services page
│   ├── portofolio_handler.go# Portfolio page
│   ├── contact_handler.go   # Contact form POST handler
│   ├── auth_handler.go      # Login / logout
│   ├── dashboard_handler.go # Admin dashboard
│   ├── nav_handler.go       # Navigation data
│   └── download_handler.go  # File download
│
├── middleware/
│   └── middleware.go        # Cookie-based auth middleware
│
├── router/                  # Route definitions (Chi)
│
├── utils/                   # Shared utilities (logger, validator)
│
├── views/                   # HTML templates (Go text/template)
│   ├── layouts/             # Base/shared layouts
│   ├── pages/               # Public pages (home, about, portfolio, etc.)
│   ├── sections/            # Reusable page sections
│   └── admin/               # Admin panel views (login, dashboard)
│
├── public/                  # Static assets (CSS, JS, images)
│
└── logs/                    # Rotating log files (auto-generated)
```

---

## 🏗 Architecture

This application follows a **clean layered architecture** pattern:

```
HTTP Request
     │
     ▼
 [ Router ]  ──── (Chi v5 router with middleware)
     │
     ▼
 [ Middleware ]  ──── (Auth check via cookie)
     │
     ▼
 [ Handler ]  ──── (Parse request, validate input, call service)
     │
     ▼
 [ Service ]  ──── (Business logic / orchestration)
     │
     ▼
 [ Repository ]  ──── (SQL queries via pgx / pgxpool)
     │
     ▼
 [ PostgreSQL ]
```

Each layer only knows about the layer directly below it, ensuring **testability**, **separation of concerns**, and **easy maintainability**.

---

## ✅ Prerequisites

Make sure you have the following installed:

- [Go 1.21+](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)

---

## 🚀 Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/bayuf/project-app-portfolio-golang-bayufirmansyah.git
cd project-app-portfolio-golang-bayufirmansyah
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Configure Environment

Copy the example env file and fill in your values:

```bash
cp .env.example .env
```

Edit `.env` with your PostgreSQL credentials (see [Environment Variables](#-environment-variables) section).

### 4. Setup Database

Run your PostgreSQL database and create the required schema. Ensure the connection URL in `.env` is correct.

---

## ⚙️ Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
# PostgreSQL connection URL
DB_URL=postgres://username:password@localhost:5432/portfolio_db
```

| Variable | Description | Example |
|----------|-------------|---------|
| `DB_URL` | Full PostgreSQL connection string | `postgres://user:pass@localhost:5432/mydb` |

---

## ▶️ Running the App

```bash
go run main.go
```

The server will start on:

```
http://localhost:8080
```

| URL | Description |
|-----|-------------|
| `http://localhost:8080/` | Home page |
| `http://localhost:8080/about` | About page |
| `http://localhost:8080/services` | Services page |
| `http://localhost:8080/portfolio` | Portfolio page |
| `http://localhost:8080/contact` | Contact page |
| `http://localhost:8080/login` | Admin login |
| `http://localhost:8080/admin/dashboard` | Admin dashboard *(protected)* |

---

## 🧪 Running Tests

Unit tests are provided for the repository layer using `pgxmock` to mock PostgreSQL behavior without needing a live database.

```bash
go test ./...
```

To run with verbose output:

```bash
go test -v ./...
```

To run only the repository tests:

```bash
go test -v ./repository/...
```

---

## 🗺 API Routes

| Method | Path | Description | Auth Required |
|--------|------|-------------|:---:|
| `GET` | `/` | Home page | ❌ |
| `GET` | `/about` | About page | ❌ |
| `GET` | `/services` | Services page | ❌ |
| `GET` | `/portfolio` | Portfolio page | ❌ |
| `GET` | `/contact` | Contact page | ❌ |
| `POST` | `/contact` | Submit contact message | ❌ |
| `GET` | `/login` | Admin login page | ❌ |
| `POST` | `/login` | Process login | ❌ |
| `GET` | `/admin/dashboard` | Admin dashboard | ✅ |
| `GET` | `/public/*` | Static file server | ❌ |

---

<div align="center">

Made with ❤️ using pure Go · [Bayu Firmansyah](https://github.com/bayuf)

</div>
