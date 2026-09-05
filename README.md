# 🏠 Airbnb Backend — Microservices Architecture

> **A production-grade, cloud-ready booking platform** built with a decoupled microservices pattern — handling hotel inventory, reservations, and asynchronous email notifications at scale.

[![Node.js](https://img.shields.io/badge/Node.js-18%2B-339933?style=for-the-badge&logo=node.js&logoColor=white)](https://nodejs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5.x-3178C6?style=for-the-badge&logo=typescript&logoColor=white)](https://www.typescriptlang.org/)
[![Express](https://img.shields.io/badge/Express.js-4.x-000000?style=for-the-badge&logo=express&logoColor=white)](https://expressjs.com/)
[![MySQL](https://img.shields.io/badge/MySQL-Latest-4479A1?style=for-the-badge&logo=mysql&logoColor=white)](https://www.mysql.com/)
[![Redis](https://img.shields.io/badge/Redis-Queue%20%26%20Cache-DC382D?style=for-the-badge&logo=redis&logoColor=white)](https://redis.io/)
[![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)
[![Prisma](https://img.shields.io/badge/Prisma-ORM-2D3748?style=for-the-badge&logo=prisma&logoColor=white)](https://www.prisma.io/)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](LICENSE)

---

## API Documentation
See [API_DOCUMENTATION.md](./API_DOCUMENTATION.md) for the full reference.

## 📖 Overview

This project is a **scalable backend system** for a property-rental platform inspired by Airbnb, architected as independent microservices rather than a monolith. It is designed for backend engineers and hiring teams who want to see real-world patterns: service decomposition, asynchronous inter-service communication via message queues, and type-safe API design with TypeScript. The platform solves the core challenge of keeping booking reliability, hotel data management, and user notifications loosely coupled and independently deployable.

---

## ✨ Key Features

- **Microservices Architecture** — Three independently deployable services (`HotelService`, `BookingService`, `NotificationService`), each with its own responsibility and data store.
- **Asynchronous Job Queue** — Booking confirmations trigger email jobs via a **Bull Queue** (backed by Redis), ensuring the HTTP response is never blocked by notification delivery.
- **Type-Safe API Layer** — Full **TypeScript** across all services eliminates a class of runtime bugs and enables safe refactoring at scale.
- **Prisma ORM with MYSQL** — Strongly-typed database access with auto-generated migrations, eliminating raw SQL injection risk and accelerating schema evolution.
- **Decoupled Notification Service** — The `NotificationService` consumes queue jobs independently, making it trivially swappable (e.g., swap email for SMS) with zero impact on booking logic.
- **Docker Compose Orchestration** — Single-command local setup spins up all services, databases, and Redis with correct networking and environment injection.
- **Room Availability Management** — Stateful reservation tracking prevents double-booking by validating room availability before confirming reservations.
- **RESTful API Design** — Clean, versioned REST endpoints with consistent request/response contracts across all services.

---

## 🛠️ Tech Stack

| Category | Technology |
|---|---|
| **Runtime** | Node.js 18+ |
| **Language** | TypeScript 5.x |
| **Framework** | Express.js |
| **ORM** | Prisma 7 |
| **Primary Database** | MYSQL |
| **Cache / Queue Store** | Redis |
| **Job Queue** | Bull (BullMQ) |
| **Containerization** | Docker & Docker Compose |
| **Package Manager** | npm |
| **Dev Tooling** | ts-node, nodemon, ESLint |

---

## 🏛️ System Architecture

The system follows an **event-driven microservices** pattern. Each service owns its domain completely and communicates asynchronously through a shared Redis-backed Bull Queue.

```
┌─────────────────────────────────────────────────────────────┐
│                       Client Application                    │
└────────────────────────────┬────────────────────────────────┘
                             │
              ┌──────────────┴──────────────┐
              │                             │
   ┌──────────▼──────────┐      ┌──────────▼──────────┐
   │    HotelService     │      │   BookingService    │
   │    Port: 3001       │◄─────│   Port: 3002       │
   │  (Sequelize + MySQL)│ HTTP │ Prisma + MySQL/Maria│
   └─────────────────────┘      └──────────┬──────────┘
              │                            │
   ┌──────────▼──────────┐                 │
   │   Redis + BullMQ    │◄────────────────┘
   │  (Room Gen Queue)   │
   └──────────┬──────────┘
              │ Job Processing
   ┌──────────▼──────────┐      ┌─────────────────────┐
   │  RoomGeneration     │      │  NotificationService│
   │     Worker          │      │  (Mailer Worker)    │
   └─────────────────────┘      └─────────────────────┘
```

**Data Stores**
- **MySQL** — Hotel inventory (Sequelize) · Booking records (Prisma)
- **Redis** — BullMQ queue state · Redlock distributed locks

**Flow summary:**
1. The client queries `HotelService` to browse available properties.
2. On booking, `BookingService` validates room availability, persists the reservation in MYSQL, then enqueues a notification job in Redis.
3. `NotificationService` picks up the job asynchronously and dispatches the confirmation email — completely decoupled from the HTTP response cycle.

---

## 🚀 Getting Started

### Prerequisites

Ensure the following are installed on your machine:

- [Node.js](https://nodejs.org/) v18+
- [Docker](https://www.docker.com/) & Docker Compose
- [Git](https://git-scm.com/)

### 1. Clone the Repository

```bash
git clone https://github.com/Akash-Verma96/Airbnb-backend.git
cd Airbnb-backend
```

### 2. Configure Environment Variables

Each service requires its own `.env` file. Use the templates below.

**`HotelService/.env`**
```env
PORT=3001
DB_HOST=localhost
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=airbnb_hotels
REDIS_HOST=localhost
REDIS_PORT=6379
ROOM_CRON=0 2 * * *
```

**`BookingService/.env`**
```env
PORT=3002
DATABASE_URL="mysql://USER:PASSWORD@localhost:3306/airbnb_bookings"
REDIS_SERVER_URL=redis://localhost:6379
LOCK_TTL=5000
HOTEL_SERVICE_URL=http://localhost:3001
```

**`NotificationService/.env`**
```env
PORT=3003
REDIS_HOST=localhost
REDIS_PORT=6379
MAIL_USER=your_email@gmail.com
MAIL_PASS=your_app_password
```

### 3. Install Dependencies (per service)

```bash
# Install for each service
cd BookingService && npm install && cd ..
cd HotelService && npm install && cd ..
cd NotificationService && npm install && cd ..
```

### 4. Run Database Migrations

```bash
# Run Prisma migrations for services that use MYSQL
cd BookingService
npx prisma migrate dev --name init
cd ../HotelService
npx prisma migrate dev --name init
```

### 5. Start with Docker Compose (Recommended)

```bash
# From the repo root — spins up all services, MYSQL, and Redis
docker-compose up --build
```

### 6. Start Services Individually (Development)

```bash
# Terminal 1 — Hotel Service
cd HotelService && npm run dev

# Terminal 2 — Booking Service
cd BookingService && npm run dev

# Terminal 3 — Notification Service
cd NotificationService && npm run dev
```

Services will be available at:
- HotelService → `http://localhost:3001`
- BookingService → `http://localhost:3002`
- NotificationService → `http://localhost:3003`

---


**Typical usage flow:**
1. `POST /hotels` — Register a new hotel listing.
2. `GET /hotels?city=Mumbai` — Query available hotels by city.
3. `POST /bookings` — Reserve a room; triggers async email notification.
4. `DELETE /bookings/:id` — Cancel a reservation and release the room.

---

## 📡 API Reference

Full endpoint reference — request bodies, response schemas, query parameters, 
error codes, and data models — is documented in:

**[→ API_DOCUMENTATION.md](./API_DOCUMENTATION.md)**

Quick summary:
- **HotelService** — 14 endpoints covering hotel CRUD, room availability, 
  room generation jobs, and the availability scheduler
- **BookingService** — 2 endpoints implementing a two-phase booking flow 
  with distributed locking
- **NotificationService** — background worker only, no public HTTP endpoints

---

## 🧠 Lessons Learned / Technical Challenges

### Challenge 1: Preventing Double-Booking Under Concurrent Requests

**The Problem:** When two users simultaneously attempted to book the last available room, both requests would read "1 room available," pass the validation check, and both get confirmed — resulting in an over-committed booking.

**The Solution:** A **Redlock distributed lock** (via Redis) is acquired on 
`hotel:<hotelId>` before the availability check runs. This ensures only one 
booking request can proceed at a time per hotel, eliminating the race condition. 
The lock is scoped to a configurable TTL (`LOCK_TTL` env var, default 5000ms).

---

### Challenge 2: Decoupling Notification Delivery from Booking Latency

**The Problem:** Sending confirmation emails synchronously during the booking flow introduced ~800ms of latency per request (SMTP round-trip), degrading user experience and coupling the booking transaction's success to email provider uptime.

**The Solution:** Email delivery was extracted entirely into `NotificationService` and placed behind a **Bull Queue**. The `BookingService` simply enqueues a lightweight job payload (booking ID, user email, template key) and returns the HTTP response immediately. `NotificationService` processes jobs asynchronously with retry logic, backoff, and dead-letter handling — so a temporary SMTP failure never surfaces to the end user.

---

### Challenge 3: Maintaining Type Safety Across Service Boundaries

**The Problem:** With three separate Node.js services, shared data contracts (DTOs, enums, error shapes) were initially duplicated in each codebase — creating silent drift when one service changed a field name.

**The Solution:** Shared types were extracted into a common types package and imported by each service. Combined with **TypeScript strict mode**, any contract mismatch surfaces at compile time rather than in production.

---

## 📄 License & Contact

This project is licensed under the **MIT License** — see the [LICENSE](LICENSE) file for details.

---

**Akash Verma**


[![LinkedIn](https://img.shields.io/badge/LinkedIn-Connect-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/akash-verma-0675b2225/)
[![GitHub](https://img.shields.io/badge/GitHub-Follow-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/Akash-Verma96)


---

<p align="center">Made with ❤️ and Node.js · If this project helped you, consider giving it a ⭐</p>
