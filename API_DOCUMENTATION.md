# Airbnb Backend — API Reference Documentation

> **Repository:** [Akash-Verma96/Airbnb-backend](https://github.com/Akash-Verma96/Airbnb-backend)  
> **Architecture:** Microservices — Node.js · TypeScript · MySQL · Redis · BullMQ  
> **Generated from:** Live codebase analysis (routes, controllers, validators, DTOs, Prisma schema, Sequelize models)

---

## Table of Contents

1. [System Overview](#1-system-overview)
2. [Global Conventions](#2-global-conventions)
3. [Error Format](#3-error-format)
4. [HotelService (Port 3001)](#4-hotelservice-port-3001)
   - [Health Check](#41-health-check)
   - [Hotel Endpoints](#42-hotel-endpoints)
   - [Room Endpoints](#43-room-endpoints)
   - [Room Generation Endpoints](#44-room-generation-endpoints)
   - [Scheduler Endpoints](#45-scheduler-endpoints)
5. [BookingService (Port 3001 — configurable)](#5-bookingservice)
   - [Health Check](#51-health-check)
   - [Booking Endpoints](#52-booking-endpoints)
6. [NotificationService (async — no public HTTP endpoints)](#6-notificationservice)
7. [Inter-Service Communication](#7-inter-service-communication)
8. [Data Models](#8-data-models)
9. [Environment Variables](#9-environment-variables)

---

## 1. System Overview

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

> \* Each service defaults to port `3001` (overridden via `PORT` env variable).  
> In a real deployment, separate port values must be set per service.

---

**Data Stores**
- **MySQL** — Hotel inventory (Sequelize) · Booking records (Prisma)
- **Redis** — BullMQ queue state · Redlock distributed locks


## 2. Global Conventions

### Base URL Pattern

```
http://<host>:<port>/api/v1
```

All currently implemented routes live under `/api/v1`. The `/api/v2` prefix is registered in all services but only exposes a ping health-check endpoint.

### Request Headers

| Header | Value | Required |
|---|---|---|
| `Content-Type` | `application/json` | For all `POST` / `PATCH` requests with a body |
| `x-correlation-id` | `<uuid-string>` | Automatically injected by the `correlationMiddleware`; clients may supply it for tracing |

### Authentication

> **Note:** The current codebase does **not** implement JWT or session-based authentication middleware. All endpoints are publicly accessible. No `Authorization` header is required at this time.

### Response Envelope

All success responses share a consistent envelope:

```json
{
  "message": "<human-readable description>",
  "data": { ... },
  "success": true
}
```

### Pagination

No pagination is currently implemented; `GET` collection endpoints return all records.

---

## 3. Error Format

All errors are returned by the `genericErrorHandler` middleware with the following shape:

```json
{
  "success": false,
  "error": "<error message string>"
}
```

### Validation Error (400) — from Zod middleware

```json
{
  "message": "Invalid request body",
  "success": false,
  "error": {
    "issues": [
      {
        "code": "invalid_type",
        "expected": "number",
        "received": "undefined",
        "path": ["userId"],
        "message": "User ID must be present"
      }
    ]
  }
}
```

### Common HTTP Status Codes

| Code | Meaning | Triggered By |
|---|---|---|
| `200` | OK | Successful GET / POST / PATCH operations |
| `201` | Created | Hotel creation, room fetching (uses `CREATED` code in controller) |
| `202` | Accepted | Get all hotels |
| `400` | Bad Request | Zod validation failure; no available rooms; already-finalized idempotency key |
| `404` | Not Found | Hotel/booking/idempotency key not found |
| `409` | Conflict | Distributed lock already held (concurrent booking attempt) |
| `500` | Internal Server Error | Unhandled service/database errors |

---

## 4. HotelService (Port 3001)

Entry file: `HotelService/src/server.ts`  
Router tree:

```
/api/v1
  GET   /ping
  POST  /hotels
  GET   /hotels/getAllHotels
  GET   /hotels/:id
  DELETE /hotels/:id
  PATCH /hotels
  POST  /hotels/generateRooms       ← (also at /generateRooms)
  POST  /generateRooms
  GET   /rooms/getAvailableRooms
  POST  /rooms/update-rooms-id
  POST  /scheduler/start
  POST  /scheduler/stop
  GET   /scheduler/status
  POST  /scheduler/extend

/api/v2
  GET   /ping
```

---

### 4.1 Health Check

#### `GET /api/v1/ping`

Simple liveness probe.

**Authentication:** None

**Request Headers:** None required

**Request Body:** None

**Success Response — `200 OK`**

```json
{
  "message": "pong",
  "success": true
}
```

---

### 4.2 Hotel Endpoints

---

#### `POST /api/v1/hotels`

Creates a new hotel record.

**Authentication:** None  
**Validation Middleware:** `validateRequestBody(hotelSchema)`

**Request Headers:**

```
Content-Type: application/json
```

**Request Body:**

| Field | Type | Required | Description |
|---|---|---|---|
| `name` | `string` (min 1) | ✅ | Hotel name |
| `address` | `string` (min 1) | ✅ | Physical address |
| `location` | `string` (min 1) | ✅ | City / region string |
| `rating` | `number` (float) | ❌ | Average star rating (0–5) |
| `ratingCount` | `number` (integer) | ❌ | Number of ratings received |

```json
{
  "name": "The Grand Horizon",
  "address": "221B Baker Street, London, UK",
  "location": "London",
  "rating": 4.7,
  "ratingCount": 1240
}
```

**Success Response — `201 Created`**

```json
{
  "message": "Hotel created Successfully!",
  "data": {
    "id": 1,
    "name": "The Grand Horizon",
    "address": "221B Baker Street, London, UK",
    "location": "London",
    "rating": 4.7,
    "ratingCount": 1240,
    "deletedAt": null,
    "createdAt": "2025-07-15T10:00:00.000Z",
    "updatedAt": "2025-07-15T10:00:00.000Z"
  },
  "success": true
}
```

**Error Responses:**

| Code | Scenario | Body |
|---|---|---|
| `400` | Missing required field or wrong type | `{ "message": "Invalid request body", "success": false, "error": { ... } }` |
| `500` | Database error | `{ "success": false, "error": "Internal server error" }` |

---

#### `GET /api/v1/hotels/getAllHotels`

Returns all hotels that have not been soft-deleted (`deletedAt IS NULL`).

**Authentication:** None

**Request Headers:** None required

**Request Body:** None

**Success Response — `202 Accepted`**

```json
{
  "message": "Hotel Detail Found!",
  "data": [
    {
      "id": 1,
      "name": "The Grand Horizon",
      "address": "221B Baker Street, London, UK",
      "location": "London",
      "rating": 4.7,
      "ratingCount": 1240,
      "deletedAt": null,
      "createdAt": "2025-07-15T10:00:00.000Z",
      "updatedAt": "2025-07-15T10:00:00.000Z"
    },
    {
      "id": 2,
      "name": "Seaside Retreat",
      "address": "42 Ocean Drive, Miami, USA",
      "location": "Miami",
      "rating": null,
      "ratingCount": null,
      "deletedAt": null,
      "createdAt": "2025-07-16T08:30:00.000Z",
      "updatedAt": "2025-07-16T08:30:00.000Z"
    }
  ],
  "success": true
}
```

**Error Responses:**

| Code | Scenario | Body |
|---|---|---|
| `404` | No hotels found in DB | `{ "success": false, "error": "No Hotels found!" }` |
| `500` | Database error | `{ "success": false, "error": "Internal server error" }` |

---

#### `GET /api/v1/hotels/:id`

Retrieves a single hotel by its primary key.

**Authentication:** None

**Path Parameters:**

| Parameter | Type | Required | Description |
|---|---|---|---|
| `id` | `number` (integer) | ✅ | Hotel's auto-increment primary key |

**Request Body:** None

**Success Response — `200 OK`**

```json
{
  "message": "Hotel found Successfully!",
  "data": {
    "id": 1,
    "name": "The Grand Horizon",
    "address": "221B Baker Street, London, UK",
    "location": "London",
    "rating": 4.7,
    "ratingCount": 1240,
    "deletedAt": null,
    "createdAt": "2025-07-15T10:00:00.000Z",
    "updatedAt": "2025-07-15T10:00:00.000Z"
  },
  "success": true
}
```

**Error Responses:**

| Code | Scenario | Body |
|---|---|---|
| `404` | No hotel with the given ID | `{ "success": false, "error": "Hotel not found" }` |
| `500` | Database error | `{ "success": false, "error": "Internal server error" }` |

---

#### `DELETE /api/v1/hotels/:id`

Soft-deletes a hotel by setting its `deletedAt` timestamp. The record is NOT removed from the database.

**Authentication:** None

**Path Parameters:**

| Parameter | Type | Required | Description |
|---|---|---|---|
| `id` | `number` (integer) | ✅ | Hotel's primary key |

**Request Body:** None

**Success Response — `200 OK`**

```json
{
  "message": "Hotel Deleted Successfully",
  "data": {
    "id": 1,
    "name": "The Grand Horizon",
    "address": "221B Baker Street, London, UK",
    "location": "London",
    "rating": 4.7,
    "ratingCount": 1240,
    "deletedAt": "2025-07-20T14:22:00.000Z",
    "createdAt": "2025-07-15T10:00:00.000Z",
    "updatedAt": "2025-07-20T14:22:00.000Z"
  },
  "success": true
}
```

**Error Responses:**

| Code | Scenario | Body |
|---|---|---|
| `404` | Hotel with given ID does not exist | `{ "success": false, "error": "No hotel found!" }` |
| `500` | Database error | `{ "success": false, "error": "Internal server error" }` |

---

#### `PATCH /api/v1/hotels`

Updates the `name` of a hotel, identified by its `id` in the request body.

**Authentication:** None

**Request Headers:**

```
Content-Type: application/json
```

**Request Body:**

| Field | Type | Required | Description |
|---|---|---|---|
| `id` | `number` (integer) | ✅ | Hotel's primary key |
| `name` | `string` | ✅ | New name to assign to the hotel |

```json
{
  "id": 1,
  "name": "The Grand Horizon — Renovated"
}
```

**Success Response — `202 Accepted`**

```json
{
  "message": "Hotel Updated Successfully",
  "data": [1],
  "success": true
}
```

> The `data` field is the Sequelize `update()` return value: an array containing the number of affected rows.

**Error Responses:**

| Code | Scenario | Body |
|---|---|---|
| `400` | Missing `id` or `name` | `{ "success": false, "error": "<validation message>" }` |
| `500` | Database error | `{ "success": false, "error": "Internal server error" }` |

---

### 4.3 Room Endpoints

---

#### `GET /api/v1/rooms/getAvailableRooms`

Returns all rooms of a given category that are available (i.e., `bookingId IS NULL`) within the specified date range.

**Authentication:** None  
**Validation Middleware:** `validateQueryParam(roomSchema)`

**Query Parameters:**

| Parameter | Type | Required | Description |
|---|---|---|---|
| `roomCategoryId` | `string` → cast to `number` | ✅ | The room category to search within |
| `checkInDate` | `string` (ISO 8601 date, e.g. `YYYY-MM-DD`) | ✅ | Start of the stay |
| `checkOutDate` | `string` (ISO 8601 date, e.g. `YYYY-MM-DD`) | ✅ | End of the stay |

**Example Request:**

```
GET /api/v1/rooms/getAvailableRooms?roomCategoryId=3&checkInDate=2025-08-01&checkOutDate=2025-08-05
```

**Success Response — `201 Created`**

> Note: The controller uses `StatusCodes.CREATED` (201) even for this GET response (implementation quirk).

```json
{
  "message": "Rooms Fetched Successfully!",
  "data": [
    {
      "id": 101,
      "hotelId": 1,
      "roomCategoryId": 3,
      "roomNo": 201,
      "price": 150.00,
      "dateOfAvailability": "2025-08-01T00:00:00.000Z",
      "bookingId": null,
      "deletedAt": null,
      "createdAt": "2025-07-01T00:00:00.000Z",
      "updatedAt": "2025-07-01T00:00:00.000Z"
    },
    {
      "id": 102,
      "hotelId": 1,
      "roomCategoryId": 3,
      "roomNo": 201,
      "price": 150.00,
      "dateOfAvailability": "2025-08-02T00:00:00.000Z",
      "bookingId": null,
      "deletedAt": null,
      "createdAt": "2025-07-01T00:00:00.000Z",
      "updatedAt": "2025-07-01T00:00:00.000Z"
    }
  ],
  "success": true
}
```

**Error Responses:**

| Code | Scenario | Body |
|---|---|---|
| `400` | Missing or wrong-type query parameters | `{ "message": "Invalid request query", "success": false, "error": { ... } }` |
| `500` | Database error | `{ "success": false, "error": "Internal server error" }` |

---

#### `POST /api/v1/rooms/update-rooms-id`

Updates a batch of room records to associate them with a confirmed booking ID. Called internally by BookingService but also accessible directly.

**Authentication:** None  
**Validation Middleware:** `validateRequestBody(updateRoomAvailabilitySchema)`

**Request Headers:**

```
Content-Type: application/json
```

**Request Body:**

| Field | Type | Required | Description |
|---|---|---|---|
| `bookingId` | `number` (integer) | ✅ | The booking to link the rooms to |
| `roomIds` | `number[]` | ✅ | Array of room `id`s to mark as booked |

```json
{
  "bookingId": 42,
  "roomIds": [101, 102, 103]
}
```

**Success Response — `201 Created`**

```json
{
  "message": "Rooms updated Successfully!",
  "data": [3],
  "success": true
}
```

> `data` is the Sequelize `update()` return value (count of affected rows).

**Error Responses:**

| Code | Scenario | Body |
|---|---|---|
| `400` | Missing `bookingId` or `roomIds` | `{ "message": "Invalid request body", "success": false, "error": { ... } }` |
| `500` | Database error | `{ "success": false, "error": "Internal server error" }` |

---

### 4.4 Room Generation Endpoints

---

#### `POST /api/v1/generateRooms`

Enqueues an asynchronous BullMQ job to generate `Room` records for a specific `RoomCategory` across a date range. Processing happens in a background worker (`roomGeneration.processor.ts`).

**Authentication:** None  
**Validation Middleware:** `validateRequestBody(RoomGenerationJobSchema)`

**Request Headers:**

```
Content-Type: application/json
```

**Request Body:**

| Field | Type | Required | Default | Description |
|---|---|---|---|---|
| `roomCategoryId` | `number` (positive) | ✅ | — | Target room category ID |
| `startDate` | `string` (ISO 8601 datetime) | ✅ | — | First date to generate availability for |
| `endDate` | `string` (ISO 8601 datetime) | ✅ | — | Last date to generate availability for |
| `priceOverride` | `number` (positive) | ❌ | Category price | Override price per night |
| `batchSize` | `number` (positive) | ❌ | `100` | Number of records to insert per DB batch |

```json
{
  "roomCategoryId": 3,
  "startDate": "2025-09-01T00:00:00.000Z",
  "endDate": "2025-09-30T23:59:59.000Z",
  "priceOverride": 175.00,
  "batchSize": 50
}
```

**Success Response — `200 OK`**

> The job is enqueued asynchronously; this response confirms the job was accepted, not that rooms have been created.

```json
{
  "message": "Hotel Room updated Successfully",
  "data": {},
  "success": true
}
```

**Error Responses:**

| Code | Scenario | Body |
|---|---|---|
| `400` | Validation failure (negative ID, invalid datetime, etc.) | `{ "message": "Invalid request body", "success": false, "error": { ... } }` |
| `500` | Redis/BullMQ connection error | `{ "success": false, "error": "Internal server error" }` |

---

#### `POST /api/v1/hotels/generateRooms`

Alias of `POST /api/v1/generateRooms` registered under the hotel router. Accepts the same body and returns the same response.

> **Note:** This duplicate route exists because `generateRoomHandler` is imported in both `hotel.router.ts` and `roomGeneration.router.ts`. Prefer `/api/v1/generateRooms` for semantic clarity.

---

### 4.5 Scheduler Endpoints

The HotelService runs a `node-cron` scheduler that automatically extends room availability by one day for every room category (default: `0 2 * * *` — 2:00 AM UTC daily). These endpoints manage that scheduler at runtime.

---

#### `POST /api/v1/scheduler/start`

Starts the cron-based room availability extension scheduler. If already running, the request is a no-op (no error thrown).

**Authentication:** None

**Request Body:** None

**Success Response — `200 OK`**

```json
{
  "message": "Room availability extension scheduler started successfully",
  "success": true,
  "data": {
    "status": "started"
  }
}
```

**Error Responses:**

| Code | Scenario | Body |
|---|---|---|
| `500` | Unexpected error starting scheduler | `{ "message": "Failed to start room availability extension scheduler", "success": false, "error": "<message>" }` |

---

#### `POST /api/v1/scheduler/stop`

Stops the running cron scheduler and clears the internal task reference.

**Authentication:** None

**Request Body:** None

**Success Response — `200 OK`**

```json
{
  "message": "Room availability extension scheduler stopped successfully",
  "success": true,
  "data": {
    "status": "stopped"
  }
}
```

**Error Responses:**

| Code | Scenario | Body |
|---|---|---|
| `500` | Unexpected error stopping scheduler | `{ "message": "Failed to stop room availability extension scheduler", "success": false, "error": "<message>" }` |

---

#### `GET /api/v1/scheduler/status`

Returns whether the cron scheduler is currently active.

**Authentication:** None

**Request Body:** None

**Success Response — `200 OK`**

```json
{
  "message": "Scheduler status retrieved successfully",
  "success": true,
  "data": {
    "isRunning": true
  }
}
```

| Field | Type | Description |
|---|---|---|
| `isRunning` | `boolean` | `true` if the cron task exists and its status is `"scheduled"` |

**Error Responses:**

| Code | Scenario | Body |
|---|---|---|
| `500` | Error reading scheduler state | `{ "message": "Failed to get scheduler status", "success": false, "error": "<message>" }` |

---

#### `POST /api/v1/scheduler/extend`

Manually triggers the same room availability extension logic that the cron scheduler runs automatically. Useful for backfilling or testing.

**Authentication:** None

**Request Body:** None

**Success Response — `200 OK`**

```json
{
  "message": "Manual room availability extension completed successfully",
  "success": true,
  "data": {
    "action": "manual_extension_completed"
  }
}
```

**Error Responses:**

| Code | Scenario | Body |
|---|---|---|
| `500` | DB error during extension | `{ "message": "Failed to perform manual room availability extension", "success": false, "error": "<message>" }` |

---

## 5. BookingService

Entry file: `BookingService/src/server.ts`  
Database ORM: **Prisma** (MySQL / MariaDB)  
Distributed locking: **Redlock** over Redis

Router tree:

```
/api/v1
  GET   /ping
  POST  /booking
  POST  /booking/confirm/:idempotencyKey

/api/v2
  GET   /ping
```

---

### 5.1 Health Check

#### `GET /api/v1/ping`

**Authentication:** None  
**Request Body:** None

**Success Response — `200 OK`**

```json
{
  "message": "pong",
  "success": true
}
```

---

### 5.2 Booking Endpoints

#### Two-Phase Booking Flow

The BookingService implements a **two-phase commit** pattern to prevent double-booking and ensure data consistency under concurrent load:

1. **Phase 1 — `POST /booking`:** Acquires a distributed Redlock on the hotel resource, checks availability, creates a `PENDING` booking, generates an idempotency key (UUID v4), and links available rooms to the booking. Returns the `bookingId` and `idempotencyKey`.
2. **Phase 2 — `POST /booking/confirm/:idempotencyKey`:** Finalises the booking by updating its status from `PENDING` to `CONFIRMED` inside a Prisma DB transaction with a `SELECT ... FOR UPDATE` row lock on the idempotency key.

---

#### `POST /api/v1/booking`

Creates a new booking in `PENDING` state, acquires room slots for the requested date range, and returns an idempotency key to use for confirmation.

**Authentication:** None  
**Validation Middleware:** `validateRequestBody(createBookingSchema)`  
**Concurrency Control:** Redlock distributed lock on `hotel:<hotelId>` (TTL configured via `LOCK_TTL` env var, default `5000ms`)

**Request Headers:**

```
Content-Type: application/json
```

**Request Body:**

| Field | Type | Required | Constraint | Description |
|---|---|---|---|---|
| `userId` | `number` | ✅ | Integer | The user creating the booking |
| `hotelId` | `number` | ✅ | Integer | The hotel to book at |
| `roomCategoryId` | `number` | ✅ | Integer | The category of room to book |
| `totalGuests` | `number` | ✅ | ≥ 1 | Number of guests |
| `bookingAmount` | `number` | ✅ | > 1 | Total cost in base currency |
| `checkInDate` | `string` | ✅ | ISO 8601 date | Start date of the stay |
| `checkOutDate` | `string` | ✅ | ISO 8601 date | End date of the stay |

```json
{
  "userId": 7,
  "hotelId": 1,
  "roomCategoryId": 3,
  "totalGuests": 2,
  "bookingAmount": 750,
  "checkInDate": "2025-09-10",
  "checkOutDate": "2025-09-15"
}
```

**Success Response — `200 OK`**

```json
{
  "bookingId": 42,
  "IdempotencyKey": "b3f2a1c0-4d5e-6f7a-8b9c-0d1e2f3a4b5c"
}
```

| Field | Type | Description |
|---|---|---|
| `bookingId` | `number` | The newly-created booking's primary key |
| `IdempotencyKey` | `string` (UUID v4) | Must be used within the TTL window to confirm the booking |

**Error Responses:**

| Code | Scenario | Body |
|---|---|---|
| `400` | Validation failure (missing field, wrong type) | `{ "message": "Invalid request body", "success": false, "error": { ... } }` |
| `400` | No available rooms for the date range | `{ "success": false, "error": "No rooms available for given dates" }` |
| `400` | Total nights > available room-day slots | `{ "success": false, "error": "No rooms available for given dates" }` |
| `500` | Redlock already held (another concurrent booking in progress) | `{ "success": false, "error": "Error already lock acquired by other person." }` |
| `500` | HotelService unreachable | `{ "success": false, "error": "<axios error message>" }` |

---

#### `POST /api/v1/booking/confirm/:idempotencyKey`

Transitions a booking from `PENDING` → `CONFIRMED`. Uses a Prisma interactive transaction with a `SELECT ... FOR UPDATE` row lock to prevent race conditions from duplicate confirmation requests.

**Authentication:** None

**Path Parameters:**

| Parameter | Type | Required | Description |
|---|---|---|---|
| `idempotencyKey` | `string` (UUID v4) | ✅ | The key returned from `POST /api/v1/booking` |

**Request Body:** None

**Success Response — `200 OK`**

```json
{
  "bookingId": 42,
  "status": "CONFIRMED"
}
```

| Field | Type | Description |
|---|---|---|
| `bookingId` | `number` | Confirmed booking's primary key |
| `status` | `"CONFIRMED"` | Updated booking status |

**Error Responses:**

| Code | Scenario | Body |
|---|---|---|
| `400` | Invalid UUID format for `idempotencyKey` | `{ "success": false, "error": "Invalid idempotency key format" }` |
| `400` | Idempotency key was already finalized (duplicate confirm call) | `{ "success": false, "error": "Idempotency key already finalized" }` |
| `404` | Idempotency key not found in database | `{ "success": false, "error": "Idempotency key not found" }` |
| `500` | Database transaction error | `{ "success": false, "error": "Internal server error" }` |

---

## 6. NotificationService

> **The NotificationService does not expose public-facing HTTP API endpoints.**

It is a **background worker service** that consumes email jobs from a BullMQ queue (backed by Redis) and sends transactional emails using Nodemailer + Handlebars templates.

### Architecture

```
BookingService ──► Redis BullMQ Queue (mailer-queue)
                                │
                   NotificationService Worker
                                │
                       renderMailTemplate()   ← Handlebars (.hbs)
                                │
                         Nodemailer sendMail()
                                │
                        SMTP Server (Gmail / custom)
```

### Queue Job Schema (`NotificationDto`)

Jobs are added to the queue using `addEmailToQueue()`. The payload conforms to the following interface:

```typescript
interface NotificationDto {
  to: string;          // Recipient email address
  subject: string;     // Email subject line
  templateId: string;  // Handlebars template identifier (e.g., "welcome")
  params: Record<string, any>;  // Template variable substitutions
}
```

**Example Job Payload:**

```json
{
  "to": "guest@example.com",
  "subject": "Your Booking Confirmation — The Grand Horizon",
  "templateId": "welcome",
  "params": {
    "name": "Alex",
    "appName": "Booking App"
  }
}
```

### Worker Behaviour

| Event | Action |
|---|---|
| Job received with name `"payload:mail"` | Render Handlebars template → send via Nodemailer |
| Job name mismatch | Throw `Error("Invalid job name")` → job marked as failed |
| Nodemailer error | Throw `InternalServerError` → job retried per BullMQ policy |
| `completed` | Log success |
| `failed` | Log failure: `"Email processing failed!"` |

### Health Check

The v1 and v2 ping routes are registered (same as the other services):

```
GET /api/v1/ping   →  { "message": "pong", "success": true }
GET /api/v2/ping   →  { "message": "pong", "success": true }
```

---

## 7. Inter-Service Communication

BookingService calls HotelService synchronously over HTTP using **Axios**:

| Call | Method | HotelService Endpoint | Purpose |
|---|---|---|---|
| `getAvailableRooms()` | `GET` | `/api/v1/rooms/getAvailableRooms` | Check room availability before creating a booking |
| `updateRoomAvailability()` | `POST` | `/api/v1/rooms/update-rooms-id` | Associate booked room IDs with the new booking ID |

BookingService base URL for HotelService is configured via `HOTEL_SERVICE_URL` env variable (default: `http://localhost:3002`).

**Sequence Diagram — Create Booking:**

```
Client          BookingService           HotelService         Redis (Redlock)
  │                   │                       │                      │
  │  POST /booking    │                       │                      │
  │──────────────────►│                       │                      │
  │                   │  acquire lock on      │                      │
  │                   │  hotel:<hotelId>      │                      │
  │                   │─────────────────────────────────────────────►│
  │                   │  lock acquired        │                      │
  │                   │◄─────────────────────────────────────────────│
  │                   │  GET /rooms/          │                      │
  │                   │  getAvailableRooms    │                      │
  │                   │──────────────────────►│                      │
  │                   │  [room list]          │                      │
  │                   │◄──────────────────────│                      │
  │                   │  createBooking (DB)   │                      │
  │                   │  createIdempotency    │                      │
  │                   │  POST /rooms/         │                      │
  │                   │  update-rooms-id      │                      │
  │                   │──────────────────────►│                      │
  │                   │  rooms updated        │                      │
  │                   │◄──────────────────────│                      │
  │  { bookingId,     │                       │                      │
  │    idempotencyKey}│                       │                      │
  │◄──────────────────│                       │                      │
```

---

## 8. Data Models

### HotelService — Sequelize / MySQL

#### `hotels` Table

| Column | SQL Type | Nullable | Default | Notes |
|---|---|---|---|---|
| `id` | `INTEGER` | No | auto-increment | Primary key |
| `name` | `VARCHAR` | No | — | Hotel display name |
| `address` | `VARCHAR` | No | — | Physical address |
| `location` | `VARCHAR` | No | — | City or region |
| `rating` | `FLOAT` | Yes | `NULL` | Average guest rating |
| `rating_count` | `INTEGER` | Yes | `NULL` | Total ratings count |
| `created_at` | `DATETIME` | Yes | `NOW()` | Auto-managed |
| `updated_at` | `DATETIME` | Yes | `NOW()` | Auto-managed |
| `deleted_at` | `DATETIME` | Yes | `NULL` | Soft-delete timestamp |

#### `room_categories` Table

| Column | SQL Type | Nullable | Default | Notes |
|---|---|---|---|---|
| `id` | `INTEGER` | No | auto-increment | Primary key |
| `hotel_id` | `INTEGER` | No | — | FK → `hotels.id` |
| `room_type` | `ENUM` | Yes | — | `SINGLE`, `DOUBLE`, `FAMILY`, `DELUXE`, `SUITE` |
| `room_no` | `INTEGER` | No | — | Room number |
| `price` | `INTEGER` | No | — | Base price per night |
| `created_at` | `DATETIME` | Yes | — | Auto-managed |
| `updated_at` | `DATETIME` | Yes | — | Auto-managed |
| `deleted_at` | `DATETIME` | Yes | `NULL` | Soft-delete timestamp |

#### `rooms` Table

| Column | SQL Type | Nullable | Default | Notes |
|---|---|---|---|---|
| `id` | `INTEGER` | No | auto-increment | Primary key |
| `hotel_id` | `INTEGER` | No | — | FK → `hotels.id` (CASCADE) |
| `room_category_id` | `INTEGER` | No | — | FK → `room_categories.id` |
| `room_no` | `INTEGER` | No | — | Room number |
| `price` | `FLOAT` | No | — | Price for this specific slot |
| `date_of_availability` | `DATETIME` | No | — | The calendar date this slot represents |
| `booking_id` | `INTEGER` | Yes | `NULL` | FK → Booking (set when booked) |
| `created_at` | `DATETIME` | Yes | — | Auto-managed |
| `updated_at` | `DATETIME` | Yes | — | Auto-managed |
| `deleted_at` | `DATETIME` | Yes | `NULL` | Soft-delete timestamp |

---

### BookingService — Prisma / MySQL

#### `Booking` Table (Prisma Model)

| Column | Type | Nullable | Default | Notes |
|---|---|---|---|---|
| `id` | `Int` | No | auto-increment | Primary key |
| `userId` | `Int` | No | — | Application user identifier |
| `hotelId` | `Int` | No | — | References HotelService hotel |
| `checkInDate` | `DateTime` | No | — | |
| `checkOutDate` | `DateTime` | No | — | |
| `roomCategoryId` | `Int` | No | — | References HotelService room category |
| `bookingAmount` | `Int` | No | — | Total cost |
| `totalGuests` | `Int` | No | — | Guest count |
| `status` | `BookingStatus` | No | `PENDING` | `PENDING` → `CONFIRMED` → `CANCELLED` |
| `createdAt` | `DateTime` | No | `now()` | |
| `updatedAt` | `DateTime` | No | auto-update | |

**BookingStatus enum values:** `PENDING` · `CONFIRMED` · `CANCELLED`

#### `IdempotencyKey` Table (Prisma Model)

| Column | Type | Nullable | Default | Notes |
|---|---|---|---|---|
| `id` | `Int` | No | auto-increment | Primary key |
| `idemKey` | `String` | No | — | UUID v4, unique |
| `bookingId` | `Int` | No | — | FK → `Booking.id`, unique (1:1) |
| `finalized` | `Boolean` | No | `false` | Set to `true` after confirm |
| `createdAt` | `DateTime` | No | `now()` | |
| `updatedAt` | `DateTime` | No | auto-update | |

---

## 9. Environment Variables

### HotelService

| Variable | Default | Description |
|---|---|---|
| `PORT` | `3001` | HTTP server port |
| `REDIS_PORT` | `6379` | Redis port |
| `REDIS_HOST` | `localhost` | Redis hostname |
| `ROOM_CRON` | `0 2 * * *` | Cron expression for the room availability extension job |
| `DB_HOST` | `localhost` | MySQL host |
| `DB_USER` | `root` | MySQL username |
| `DB_PASSWORD` | `root` | MySQL password |
| `DB_NAME` | `test_db` | MySQL database name |

### BookingService

| Variable | Default | Description |
|---|---|---|
| `PORT` | `3001` | HTTP server port |
| `REDIS_SERVER_URL` | `redis://localhost:6379` | Full Redis connection URL (used by Redlock) |
| `LOCK_TTL` | `5000` | Distributed lock time-to-live in milliseconds |
| `HOTEL_SERVICE_URL` | `http://localhost:3002` | Base URL for HotelService HTTP calls |
| `DATABASE_URL` | — | Prisma connection string (MySQL/MariaDB format) |

### NotificationService

| Variable | Default | Description |
|---|---|---|
| `PORT` | `3001` | HTTP server port |
| `REDIS_PORT` | `6379` | Redis port |
| `REDIS_HOST` | `localhost` | Redis hostname |
| `MAIL_USER` | `""` | SMTP sender email address |
| `MAIL_PASS` | `""` | SMTP sender password / app password |

---

## Appendix A — RoomType Enum Values

| Value | Description |
|---|---|
| `SINGLE` | Single occupancy room |
| `DOUBLE` | Double occupancy room |
| `FAMILY` | Family room |
| `DELUXE` | Deluxe room |
| `SUITE` | Suite |

---

## Appendix B — BookingStatus State Machine

```
  ┌──────────┐   POST /booking/confirm/:key   ┌───────────┐
  │ PENDING  │──────────────────────────────►│ CONFIRMED │
  └──────────┘                               └───────────┘
       │                                          │
       │  (manual cancellation — not yet          │
       │   exposed via HTTP endpoint)             │
       ▼                                          ▼
  ┌───────────┐                            ┌───────────┐
  │ CANCELLED │                            │ CANCELLED │
  └───────────┘                            └───────────┘
```

> **Note:** The `cancelBooking` function is implemented in `booking.repository.ts` but no HTTP route for cancellation is currently registered. It is a candidate for a future endpoint.

---

## Appendix C — Quick Reference

### HotelService Endpoints

| Method | Path | Description |
|---|---|---|
| `GET` | `/api/v1/ping` | Health check |
| `POST` | `/api/v1/hotels` | Create a hotel |
| `GET` | `/api/v1/hotels/getAllHotels` | List all non-deleted hotels |
| `GET` | `/api/v1/hotels/:id` | Get hotel by ID |
| `DELETE` | `/api/v1/hotels/:id` | Soft-delete hotel |
| `PATCH` | `/api/v1/hotels` | Update hotel name |
| `POST` | `/api/v1/generateRooms` | Enqueue room generation job |
| `POST` | `/api/v1/hotels/generateRooms` | Alias of above |
| `GET` | `/api/v1/rooms/getAvailableRooms` | Query available rooms by date range |
| `POST` | `/api/v1/rooms/update-rooms-id` | Link rooms to a booking |
| `POST` | `/api/v1/scheduler/start` | Start availability cron |
| `POST` | `/api/v1/scheduler/stop` | Stop availability cron |
| `GET` | `/api/v1/scheduler/status` | Get cron status |
| `POST` | `/api/v1/scheduler/extend` | Manually extend room availability |

### BookingService Endpoints

| Method | Path | Description |
|---|---|---|
| `GET` | `/api/v1/ping` | Health check |
| `POST` | `/api/v1/booking` | Create a pending booking |
| `POST` | `/api/v1/booking/confirm/:idempotencyKey` | Confirm a pending booking |

---

*Documentation generated by codebase analysis — September 2026*
