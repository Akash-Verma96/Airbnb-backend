# 🏠 Airbnb Backend — Microservices Architecture

> **A production-grade, cloud-ready booking platform** built with a decoupled microservices architecture for authentication, hotel inventory, reservations, and asynchronous notifications.

[![Go](https://img.shields.io/badge/Go-1.x-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Node.js](https://img.shields.io/badge/Node.js-18%2B-339933?style=for-the-badge&logo=node.js&logoColor=white)](https://nodejs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.x-3178C6?style=for-the-badge&logo=typescript&logoColor=white)](https://www.typescriptlang.org/)
[![Express](https://img.shields.io/badge/Express.js-4.x-000000?style=for-the-badge&logo=express&logoColor=white)](https://expressjs.com/)
[![MySQL](https://img.shields.io/badge/MySQL-Latest-4479A1?style=for-the-badge&logo=mysql&logoColor=white)](https://www.mysql.com/)
[![Redis](https://img.shields.io/badge/Redis-Queue%20%26%20Locks-DC382D?style=for-the-badge&logo=redis&logoColor=white)](https://redis.io/)
[![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)
[![Prisma](https://img.shields.io/badge/Prisma-ORM-2D3748?style=for-the-badge&logo=prisma&logoColor=white)](https://www.prisma.io/)

---

## 📚 API Documentation

- **Live API Documentation:** https://airbnb-api-docs-dg88.onrender.com/
- **Full API Reference:** [API_DOCUMENTATION.md](./API_DOCUMENTATION.md)

The API documentation contains endpoint details, authentication requirements, request/response examples, validation rules, service routes, and environment configuration.

---

## 📖 Overview

This project is a **microservices-based backend system** for a property-rental platform inspired by Airbnb.

The backend is divided into independently deployable services:

- **AuthInGo** — Go-based authentication, authorization, JWT handling, role management, CORS, rate limiting, and API gateway/proxy.
- **HotelService** — Hotel inventory, room availability, room generation, and scheduler management.
- **BookingService** — Booking creation and confirmation with Redis-based distributed locking.
- **NotificationService** — Asynchronous email processing using Redis/BullMQ.

The architecture separates responsibilities so each service can evolve and scale independently.

---

## ✨ Key Features

- **Microservices Architecture** — Authentication, hotel management, booking, and notification responsibilities are separated into independent services.
- **API Gateway** — AuthInGo acts as the public gateway and proxies requests to HotelService and BookingService.
- **JWT Authentication** — Protected endpoints use `Authorization: Bearer <JWT_TOKEN>`.
- **Role-Based Authorization** — Administrative operations are protected using role-based middleware.
- **Rate Limiting** — Login requests are protected by a Redis-backed rate limiter.
- **Hotel & Room Management** — Supports hotel CRUD, room availability checks, room ID updates, and bulk room generation.
- **Booking Workflow** — Booking creation and confirmation are separated using an idempotency key.
- **Distributed Locking** — BookingService uses Redis/Redlock to reduce concurrent booking conflicts.
- **Asynchronous Notifications** — NotificationService processes email jobs independently of the main request flow.
- **Redis Integration** — Used for rate limiting, distributed locks, queues, and background processing.
- **Validation** — Request DTOs and validators protect API contracts.
- **Docker Support** — Services and infrastructure can be run locally using Docker Compose.
- **REST APIs** — Versioned APIs are exposed through `/api/v1` and `/api/v2` where applicable.

---

## 🏛️ System Architecture

```text
                         ┌──────────────────────────┐
                         │      Client / Frontend    │
                         └─────────────┬────────────┘
                                       │
                                       ▼
                         ┌──────────────────────────┐
                         │       AuthInGo Gateway    │
                         │       Go + Chi Router     │
                         │  JWT / CORS / Rate Limit │
                         └─────────────┬────────────┘
                                       │
                       ┌───────────────┴────────────────┐
                       │                                │
                       ▼                                ▼
             ┌───────────────────┐            ┌───────────────────┐
             │    HotelService   │            │   BookingService  │
             │ Node + TypeScript │◄───────────│ Node + TypeScript │
             │ Sequelize + MySQL │    HTTP    │ Prisma + MySQL    │
             └─────────┬─────────┘            └─────────┬─────────┘
                       │                                │
                       │                                │
                       ▼                                ▼
             ┌───────────────────┐            ┌───────────────────┐
             │ Redis + BullMQ    │            │ Redis + Redlock   │
             │ Room Generation   │            │ Distributed Locks │
             └─────────┬─────────┘            └─────────┬─────────┘
                       │                                │
                       └───────────────┬────────────────┘
                                       ▼
                         ┌──────────────────────────┐
                         │   NotificationService   │
                         │ Async Mail Worker       │
                         │ Redis + BullMQ + SMTP    │
                         └──────────────────────────┘
```

### Data Stores

| Component | Responsibility |
|---|---|
| MySQL | Hotel inventory and booking data |
| Redis | Rate limiting, distributed locks, queues, background jobs |
| Sequelize | HotelService database access |
| Prisma | BookingService database access |

### Request Flow

1. The client sends authentication or API requests to **AuthInGo**.
2. AuthInGo validates authentication/authorization requirements where applicable.
3. Hotel requests are proxied to **HotelService**.
4. Booking requests are proxied to **BookingService**.
5. BookingService communicates with HotelService to validate/update room availability.
6. Redis/Redlock protects critical booking operations from concurrent conflicts.
7. Notification jobs are processed asynchronously by **NotificationService**.

---

## 🌐 Production Services

| Service | Production URL |
|---|---|
| AuthInGo Gateway | `https://airbnb-auth.onrender.com` |
| HotelService | `https://airbnb-hotel-0job.onrender.com` |
| BookingService | `https://airbnb-backend-glo7.onrender.com` |
| API Documentation | `https://airbnb-api-docs-dg88.onrender.com` |

### Gateway Proxy Routes

```text
https://airbnb-auth.onrender.com/HotelService/*
        └──► HotelService

https://airbnb-auth.onrender.com/BookingService/*
        └──► BookingService
```

---

## 🛠️ Tech Stack

| Category | Technology |
|---|---|
| Authentication / Gateway | Go, Chi Router |
| Runtime | Node.js 18+ |
| Language | TypeScript 5.x |
| Backend Framework | Express.js |
| Hotel ORM | Sequelize |
| Booking ORM | Prisma |
| Primary Database | MySQL |
| Cache / Lock / Queue Store | Redis |
| Background Jobs | BullMQ |
| Distributed Lock | Redlock |
| Authentication | JWT |
| Validation | DTOs / Validators |
| Containerization | Docker & Docker Compose |
| Package Manager | npm |
| Development Tools | ts-node, nodemon, ESLint |

---

## 📁 Project Structure

```text
Airbnb-backend/
│
├── AuthInGo/
│   ├── controllers/
│   ├── dto/
│   ├── middlewares/
│   ├── router/
│   ├── services/
│   └── utils/
│
├── HotelService/
│   ├── src/
│   │   ├── controllers/
│   │   ├── routers/
│   │   ├── services/
│   │   ├── validators/
│   │   └── ...
│   └── ...
│
├── BookingService/
│   ├── src/
│   │   ├── controllers/
│   │   ├── routers/
│   │   ├── services/
│   │   ├── validators/
│   │   └── ...
│   ├── prisma/
│   └── ...
│
├── NotificationService/
│   ├── src/
│   └── ...
│
├── docker-compose.yml
├── API_DOCUMENTATION.md
└── README.md
```

---

## 🚀 Getting Started

### Prerequisites

Install:

- [Go](https://go.dev/)
- [Node.js](https://nodejs.org/) v18+
- [Docker](https://www.docker.com/) & Docker Compose
- [Git](https://git-scm.com/)
- MySQL
- Redis

### 1. Clone the Repository

```bash
git clone https://github.com/Akash-Verma96/Airbnb-backend.git
cd Airbnb-backend
```

### 2. Configure Environment Variables

Each service has its own configuration.

#### AuthInGo

```env
PORT=8080
JWT_SECRET=your_jwt_secret
```

Use the environment variables defined by the AuthInGo configuration for database and Redis connectivity.

#### HotelService

Typical local configuration:

```env
PORT=3001
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=root
DB_NAME=test_db
DB_PORT=10337

REDIS_HOST=localhost
REDIS_PORT=6379

ROOM_CRON=0 2 * * *
```

#### BookingService

```env
PORT=3001
DATABASE_URL="mysql://USER:PASSWORD@localhost:3306/DATABASE"

REDIS_SERVER_URL=redis://localhost:6379
LOCK_TTL=5000
HOTEL_SERVICE_URL=http://localhost:3002
```

#### NotificationService

```env
PORT=3001

REDIS_HOST=localhost
REDIS_PORT=6379

MAIL_USER=your_email@gmail.com
MAIL_PASS=your_app_password
```

> Use your actual deployment-specific values in production. Do not commit secrets to Git.

---

## 📦 Install Dependencies

```bash
cd HotelService
npm install

cd ../BookingService
npm install

cd ../NotificationService
npm install
```

For AuthInGo:

```bash
cd AuthInGo
go mod download
```

---

## ▶️ Run Locally

### AuthInGo

```bash
cd AuthInGo
go run .
```

### HotelService

```bash
cd HotelService
npm run dev
```

### BookingService

```bash
cd BookingService
npm run dev
```

### NotificationService

```bash
cd NotificationService
npm run dev
```

> Local ports are controlled by each service's environment/configuration. Do not assume production Render ports are the same as local ports.

---

## 🐳 Docker Compose

From the repository root:

```bash
docker compose up --build
```

To stop the stack:

```bash
docker compose down
```

---

## 📡 API Reference

The complete endpoint reference is maintained in:

**[→ API_DOCUMENTATION.md](./API_DOCUMENTATION.md)**

### AuthInGo

Authentication and gateway endpoints include:

```text
GET    /ping
POST   /signup
POST   /login
GET    /profile
GET    /all
DELETE /
```

Role management includes:

```text
GET    /roles/{id}
GET    /roles
POST   /createRole
PATCH  /roles/{id}
GET    /roleByName
GET    /role/{id}/permissions
POST   /assignpermissionToRole
DELETE  /removePermissionFromRole
POST   /roles/{userId}/assign/{roleId}
DELETE /roles/{id}
```

Protected routes require:

```http
Authorization: Bearer <JWT_TOKEN>
```

### HotelService

Main API areas:

```text
GET    /api/v1/ping

POST   /api/v1/hotels
GET    /api/v1/hotels/getAllHotels
GET    /api/v1/hotels/:id
PATCH  /api/v1/hotels/:id
DELETE /api/v1/hotels/:id

GET    /api/v1/rooms/getAvailableRooms
POST   /api/v1/rooms/update-rooms-id

POST   /api/v1/generateRooms
POST   /api/v1/hotels/generateRooms

POST   /api/v1/scheduler/start
POST   /api/v1/scheduler/stop
GET    /api/v1/scheduler/status
POST   /api/v1/scheduler/extend
```

### BookingService

```text
GET    /api/v1/ping

POST   /api/v1/booking
POST   /api/v1/booking/confirm/:idempotencyKey
GET    /api/v1/booking/getAllBookings/:id
```

### NotificationService

NotificationService primarily operates as an asynchronous mail worker. It does not expose a conventional public notification API for application workflows.

---

## 🔐 Authentication

AuthInGo uses JWT-based authentication.

Example:

```http
Authorization: Bearer <JWT_TOKEN>
```

Successful authentication returns user information together with a JWT token.

Example response shape:

```json
{
  "status": "success",
  "message": "Login successful",
  "data": {
    "user": {},
    "jwtToken": "..."
  }
}
```

---

## 🏨 Hotel & Room Management

HotelService handles:

- Hotel creation
- Hotel listing
- Hotel lookup by ID
- Hotel updates
- Hotel deletion
- Room availability
- Updating room IDs after booking operations
- Bulk room generation
- Availability scheduler operations

### Room Availability Query

```text
GET /api/v1/rooms/getAvailableRooms
```

Query parameters:

```text
roomCategoryId
checkInDate
checkOutDate
```

### Room Generation

```text
POST /api/v1/generateRooms
```

Request fields include:

```json
{
  "roomCategoryId": 1,
  "startDate": "2026-01-01",
  "endDate": "2026-01-10",
  "priceOverride": 2500,
  "batchSize": 100
}
```

---

## 🧾 Booking Workflow

BookingService uses a two-step flow:

### 1. Create Booking

```http
POST /api/v1/booking
```

Required fields include:

```json
{
  "userId": 1,
  "hotelId": 1,
  "totalGuests": 2,
  "bookingAmount": 5000,
  "checkInDate": "2026-01-10",
  "checkOutDate": "2026-01-12",
  "roomCategoryId": 1
}
```

The response contains a booking ID and idempotency key.

### 2. Confirm Booking

```http
POST /api/v1/booking/confirm/:idempotencyKey
```

The confirmation step performs the protected booking operation and returns the booking status.

### 3. Fetch User Bookings

```http
GET /api/v1/booking/getAllBookings/:id
```

---

## 🔒 Preventing Double Booking

BookingService uses **Redis + Redlock** to coordinate concurrent booking requests.

The lock is acquired before the critical availability/booking operation so concurrent requests do not independently reserve the same inventory.

The lock duration is configurable using:

```env
LOCK_TTL=5000
```

---

## 📬 Notification Processing

NotificationService handles asynchronous email processing using Redis/BullMQ and Nodemailer.

The notification worker is separated from the booking HTTP request so email processing does not need to be performed directly inside the main booking request.

---

## 🧠 Important Backend Design Patterns

### API Gateway

AuthInGo provides a single public entry point and proxies requests to internal services.

### JWT Authentication

Authentication is centralized in the Go gateway and protected routes use JWT claims.

### Role-Based Access Control

Administrative routes use role-based middleware to restrict access.

### Rate Limiting

Login requests use rate limiting to reduce repeated authentication attempts.

### Distributed Locking

Redlock prevents conflicting concurrent booking operations.

### Background Jobs

Redis/BullMQ is used for asynchronous work such as room generation and notification processing.

### Service Separation

Each service owns a focused business responsibility and can be developed/deployed independently.

---

## 🧪 Health Checks

Services expose ping/health-style endpoints.

Examples:

```http
GET /ping
GET /api/v1/ping
```

Use these endpoints to verify that the gateway or individual service is responding.

---

## 📝 API Documentation Notes

The API reference is generated from the **current backend implementation**, so route names and payloads should be treated as the source of truth.

For complete request/response schemas and endpoint details, see:

**[API_DOCUMENTATION.md](./API_DOCUMENTATION.md)**

---

## 📄 License

This project is licensed under the **MIT License** — see the [LICENSE](LICENSE) file for details.

---

## 👨‍💻 Author

**Akash Verma**

[![LinkedIn](https://img.shields.io/badge/LinkedIn-Connect-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/akash-verma-0675b2225/)

[![GitHub](https://img.shields.io/badge/GitHub-Follow-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/Akash-Verma96)

---

<p align="center">Made with ❤️ and Node.js · If this project helped you, consider giving it a ⭐</p>
