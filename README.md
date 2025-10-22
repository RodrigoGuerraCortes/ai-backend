# AI Backend
[![Coverage Status](https://coveralls.io/repos/github/RodrigoGuerraCortes/ai-backend/badge.svg)](https://coveralls.io/github/RodrigoGuerraCortes/ai-backend)
[![Go CI](https://github.com/RodrigoGuerraCortes/ai-backend/actions/workflows/test.yml/badge.svg)](https://github.com/RodrigoGuerraCortes/ai-backend/actions/workflows/test.yml)


# 🧠 AI Backend (Clean Architecture · Go + Gemini)

A clean, testable, and production-ready backend built in **Go** that integrates with **Google Gemini** for AI-powered chat capabilities.  
Follows modern **Clean Architecture** principles with isolated layers for AI, services, HTTP, and tests.

---

## 🚀 Features

- 🧩 **Clean Architecture** — modular, layered, and maintainable design  
- 🤖 **Google Gemini Integration** — real AI chat using Gemini API  
- 🧠 **Mockable AI Interface** — unit tests run fast without real API calls  
- 🧪 **Unit + Integration Tests** — full coverage, real and mocked modes  
- 🧾 **Structured Logging** — request/response middlewares  
- ⚙️ **Request Validation & DTOs** — consistent and secure input handling  
- 📘 **Swagger Ready** — API documentation generation support  

---

## 🗂️ Project Structure

```

ai-backend/
├── cmd/api/                 # App entry point
├── config/                  # Environment & configuration loader
├── internal/
│   ├── ai/                  # Gemini client + AI interface + mocks
│   ├── service/             # Business logic (ChatService)
│   ├── http/
│   │   ├── handler/         # HTTP handlers (ChatHandler)
│   │   ├── middleware/      # Logging & error middlewares
│   │   └── router/          # API routes
│   └── dto/                 # Request/response DTOs
├── pkg/                     # Shared reusable helpers (optional)
├── tests/
│   └── integration/         # Real integration tests (Gemini API)
├── Makefile                 # Run, test & tidy commands
├── .env                     # Local environment variables
└── README.md

````

---

## ⚙️ Installation

### 1️⃣ Clone the repository
```bash
git clone https://github.com/RodrigoGuerraCortes/ai-backend.git
cd ai-backend
````

### 2️⃣ Install dependencies

```bash
go mod tidy
```

### 3️⃣ Configure environment

Create your `.env` file (or copy from `.env.example`):

```bash
GEMINI_API_KEY=your_google_api_key_here
```

Get your free API key from: [https://aistudio.google.com/app/apikey](https://aistudio.google.com/app/apikey)

### 4️⃣ Run the API

```bash
make run
```

Server starts at: **[http://localhost:8080](http://localhost:8080)**

---

## 🧪 Testing

### 🧩 Unit Tests (fast, mock AI)

```bash
make tests
```

### 🔗 Integration Tests (real Gemini API)

```bash
make integration-test
```

> ⚠️ Requires a valid `GEMINI_API_KEY` and an active Gemini API quota.

---

## 📡 API Endpoints

### **POST** `/api/v1/chat`

#### Request

```json
{
  "message": "Hablame del nacionalismo chileno"
}
```

#### Response

```json
{
  "reply": "El nacionalismo chileno es una ideología compleja..."
}
```

---

## 🧱 Clean Architecture Layers

| Layer             | Folder              | Responsibility                        |
| ----------------- | ------------------- | ------------------------------------- |
| **AI Layer**      | `/internal/ai`      | Gemini API client & AI interface      |
| **Service Layer** | `/internal/service` | Business rules & use-cases            |
| **HTTP Layer**    | `/internal/http`    | Routes, handlers & middlewares        |
| **DTO Layer**     | `/internal/dto`     | Data contracts for requests/responses |
| **Tests Layer**   | `/tests`            | Unit & integration testing setup      |

---

## 🧰 Makefile Commands

| Command                 | Description                |
| ----------------------- | -------------------------- |
| `make run`              | Run API locally            |
| `make tests`            | Run all unit tests         |
| `make integration-test` | Run real integration tests |
| `make tidy`             | Clean and sync Go modules  |

---

## 🧾 Roadmap

 Goal                                         |
 -------------------------------------------- |
| Project setup + Gemini client connection     |
| Clean Architecture, service layer, and tests |
| Logging, validation, and error middleware    |
| Swagger & Docker setup                       |
| CI/CD & GitHub Actions integration           |

---

## 🧰 Tech Stack

* **Language:** Go 1.22+
* **Framework:** Gin (HTTP API)
* **AI Provider:** Google Gemini
* **Testing:** Go `testing` + mocks
* **Docs:** Swagger / OpenAPI (coming next)
* **Build Tools:** Makefile, go-mod
* **Architecture:** Clean Architecture + Dependency Injection

---

## 👨‍💻 Author

**Rodrigo Guerra Cortés**
Full Stack Developer | Chile 🇨🇱
📧 [[rguerracortes@gmail.com](mailto:rguerracortes@gmail.com)]
🌐 [LinkedIn](https://linkedin.com/in/rodrigoguerracortes)

---

## 🪪 License

This project is licensed under the **MIT License**.
Feel free to use, fork, and improve it.

---
