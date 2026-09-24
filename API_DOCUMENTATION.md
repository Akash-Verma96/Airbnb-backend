# Airbnb Backend — API Reference Documentation

> **Repository:** [Akash-Verma96/Airbnb-backend](https://github.com/Akash-Verma96/Airbnb-backend)  
> **Architecture:** Microservices — Node.js · TypeScript · Go · MySQL · Redis · BullMQ  
> **Updated from:** September 2026 backend source code (routes, controllers, validators, DTOs, Prisma schema, Sequelize models, AuthInGo middleware)

---

## Table of Contents

1. [System Overview](#1-system-overview)
2. [Gateway & Global Conventions](#2-gateway--global-conventions)
3. [AuthInGo — Authentication & User APIs](#3-authingo--authentication--user-apis)
4. [HotelService](#4-hotelservice)
   - [Health](#41-health)
   - [Hotel APIs](#42-hotel-apis)
   - [Room APIs](#43-room-apis)
   - [Room Generation](#44-room-generation)
   - [Scheduler](#45-scheduler)
5. [BookingService](#5-bookingservice)
   - [Health](#51-health)
   - [Booking APIs](#52-booking-apis)
6. [NotificationService](#6-notificationservice)
7. [Inter-Service Communication](#7-inter-service-communication)
8. [Data Models](#8-data-models)
9. [Environment Variables](#9-environment-variables)
10. [Quick Reference](#10-quick-reference)

---

# 1. System Overview

```text
                         Client / API Docs
                                |
                                v
                    +------------------------+
                    |      AuthInGo Gateway   |
                    |   JWT + Reverse Proxy  |
                    +-----------+------------+
                                |
                 +--------------+--------------+
                 |                             |
                 v                             v
        +------------------+          +------------------+
        |   HotelService   |          | BookingService  |
        | Node + TypeScript|          | Node + TypeScript|
        | Sequelize + MySQL|          | Prisma + MySQL   |
        +--------+---------+          +--------+---------+
                 |                             |
                 v                             v
        +------------------+          +------------------+
        | Redis + BullMQ   |          | Redis + Redlock  |
        | Room generation  |          | Distributed lock|
        +--------+---------+          +------------------+
                 |
                 v
        +------------------+
        | Notification     |
        | Mailer Worker    |
        +------------------+
```

### Services

| Service | Responsibility | Default Port |
|---|---|---:|
| AuthInGo | Authentication, users, roles, reverse proxy | `8080` |
| HotelService | Hotels, rooms, room generation, scheduler | `3001` |
| BookingService | Booking creation and confirmation | `3001` |
| NotificationService | Async email worker | `3001` |

> Each Node.js service defaults to port `3001`; production deployments must provide separate `PORT` values.

### Production Gateway

```text
https://airbnb-auth.onrender.com
```

The gateway currently proxies:

```text
/HotelService/*   -> https://airbnb-hotel-0job.onrender.com
/BookingService/* -> https://airbnb-backend-glo7.onrender.com
```

The gateway CORS configuration currently allows:

```text
https://airbnb-api-docs-dg88.onrender.com
```

---

# 2. Gateway & Global Conventions

## Gateway URL

For browser/API-docs requests, use:

```text
https://airbnb-auth.onrender.com
```

### Service prefixes

Hotel APIs:

```text
https://airbnb-auth.onrender.com/HotelService/api/v1/...
```

Booking APIs:

```text
https://airbnb-auth.onrender.com/BookingService/api/v1/...
```

AuthInGo user/role APIs are registered directly on the gateway:

```text
https://airbnb-auth.onrender.com/...
```

## API Versions

Node.js services register:

```text
/api/v1
/api/v2
```

The current v2 routers expose ping endpoints only.

## Request Headers

For JSON requests:

```http
Content-Type: application/json
```

For protected AuthInGo endpoints:

```http
Authorization: Bearer <JWT_TOKEN>
```

Node.js services also attach/use an `x-correlation-id` for request tracing.

## Authentication

Authentication is implemented by **AuthInGo**.

Protected endpoints use JWT authentication. The JWT middleware expects:

```http
Authorization: Bearer <JWT_TOKEN>
```

HotelService and BookingService themselves do not contain JWT middleware in the current source.

## JSON Response Conventions

Node.js services generally use:

```json
{
  "message": "Operation successful",
  "data": {},
  "success": true
}
```

AuthInGo uses:

```json
{
  "status": "success",
  "message": "Operation successful",
  "data": {}
}
```

> AuthInGo error responses currently use the key `stauts` in the response helper. This is the spelling used by the current implementation.

---

# 3. AuthInGo — Authentication & User APIs

AuthInGo is the Go-based gateway and authentication service.

## 3.1 Health Check

### `GET /ping`

Returns:

```json
{
  "status": "success",
  "message": "Service Health OK! check Completed",
  "data": "NULL"
}
```

Authentication: None.

---

## 3.2 Signup

### `POST /signup`

Creates a user and returns a JWT token.

Authentication: None.

### Request Body

| Field | Type | Required | Constraint |
|---|---|---|---|
| `username` | `string` | Yes | Minimum 3 characters |
| `email` | `string` | Yes | Valid email |
| `password` | `string` | Yes | Minimum 8 characters |

```json
{
  "username": "akash",
  "email": "akash@example.com",
  "password": "password123"
}
```

### Success

HTTP `200 OK`.

Response uses the AuthInGo success envelope:

```json
{
  "status": "success",
  "message": "User Created Successfully",
  "data": {
    "user": {},
    "jwtToken": "<JWT_TOKEN>"
  }
}
```

---

## 3.3 Login

### `POST /login`

Authenticates an existing user.

Authentication: None.

### Request Body

```json
{
  "email": "akash@example.com",
  "password": "password123"
}
```

Constraints:

- `email` is required and must be valid.
- `password` is required and must contain at least 8 characters.
- Login is protected by the current rate limiter.

### Rate Limit

The current middleware allows up to **5 requests per minute** globally before returning HTTP `429`.

### Success

HTTP `200 OK`.

```json
{
  "status": "success",
  "message": "User Logged In succesfully!",
  "data": {
    "user": {},
    "jwtToken": "<JWT_TOKEN>"
  }
}
```

---

## 3.4 Get Profile

### `GET /profile`

Authentication:

```http
Authorization: Bearer <JWT_TOKEN>
```

The JWT middleware extracts the user ID and email from the token.

### Success

HTTP `200 OK`.

```json
{
  "status": "success",
  "message": "User Fetched Successful",
  "data": {}
}
```

---

## 3.5 Get All Users

### `GET /all`

Authentication:

```http
Authorization: Bearer <JWT_TOKEN>
```

### Success

HTTP `200 OK`.

```json
{
  "status": "success",
  "message": "All Profile Fetched Successfully!",
  "data": []
}
```

---

## 3.6 Delete User

### `DELETE /`

Authentication:

```http
Authorization: Bearer <JWT_TOKEN>
```

### Request Body

```json
{
  "id": 1
}
```

### Success

HTTP `200 OK`.

```json
{
  "status": "success",
  "message": "User Deleted Successfully!",
  "data": "nil"
}
```

---

## 3.7 Get Role

### `GET /roles/:id`

Authentication: None in the current router.

Example:

```text
GET /roles/1
```

---

## 3.8 Get All Roles

### `GET /roles`

Authentication: None.

---

## 3.9 Create Role

### `POST /createRole`

Authentication: None in the current router.

### Request Body

```json
{
  "name": "manager",
  "description": "Hotel management role"
}
```

Constraints:

- `name`: minimum 3 characters.
- `description`: required, maximum 50 characters.

---

## 3.10 Update Role

### `PATCH /roles/:id`

Authentication: None in the current router.

### Request Body

Both fields are optional:

```json
{
  "name": "manager",
  "description": "Updated description"
}
```

---

## 3.11 Get Role By Name

### `GET /roleByName`

The current controller reads the role name from a JSON request body despite the route using `GET`.

Expected body:

```json
{
  "name": "admin"
}
```

> Because GET request bodies are inconsistently handled by clients/proxies, this endpoint should be treated carefully until the backend route is changed to use a query parameter or POST.

---

## 3.12 Get Role Permissions

### `GET /role/:id/permissions`

Example:

```text
GET /role/1/permissions
```

---

## 3.13 Assign Permission To Role

### `POST /assignpermissionToRole`

### Request Body

```json
{
  "id": 1,
  "permissionId": 2
}
```

---

## 3.14 Remove Permission From Role

### `DELETE /removePermissionFromRole`

### Request Body

```json
{
  "id": 1,
  "permissionId": 2
}
```

---

## 3.15 Assign Role To User

### `POST /roles/:userId/assign/:roleId`

This endpoint requires:

```text
JWT authentication
admin role
```

Example:

```text
POST /roles/7/assign/2
```

No request body is required.

---

## 3.16 Delete Role

### `DELETE /roles/:id`

Example:

```text
DELETE /roles/2
```

---

# 4. HotelService

Base path:

```text
/api/v1
```

Through the production gateway:

```text
https://airbnb-auth.onrender.com/HotelService/api/v1
```

---

# 4.1 Health

### `GET /ping`

```text
GET /HotelService/api/v1/ping
```

The current HotelService controller returns:

```json
{
  "message": "pong"
}
```

---

# 4.2 Hotel APIs

## Create Hotel

### `POST /hotels`

### Request Body

The active Zod validator requires:

| Field | Type | Required |
|---|---|---|
| `name` | `string` | Yes |
| `address` | `string` | Yes |
| `location` | `string` | Yes |
| `price` | `number` | Yes |
| `roomType` | `string` | Yes |
| `rating` | `number` | No |
| `ratingCount` | `number` | No |

Example:

```json
{
  "name": "The Grand Horizon",
  "address": "221B Baker Street, London, UK",
  "location": "London",
  "price": 2500,
  "roomType": "DELUXE",
  "rating": 4.7,
  "ratingCount": 1240
}
```

### Success

HTTP `201 Created`.

```json
{
  "message": "Hotel created Successfully!",
  "data": {},
  "success": true
}
```

> `hostId` exists in the TypeScript DTO, but the active `hotelSchema` does not accept it. The request documentation therefore follows the active validator.

---

## Get All Hotels

### `GET /hotels/getAllHotels`

Returns non-deleted hotels.

### Success

HTTP `202 Accepted`.

```json
{
  "message": "Hotel Detail Found!",
  "data": [],
  "success": true
}
```

---

## Get Hotel By ID

### `GET /hotels/:id`

Example:

```text
GET /hotels/1
```

### Success

HTTP `200 OK`.

```json
{
  "message": "Hotel found Successfully!",
  "data": {},
  "success": true
}
```

---

## Delete Hotel

### `DELETE /hotels/:id`

Soft-deletes the hotel.

Example:

```text
DELETE /hotels/1
```

### Success

HTTP `200 OK`.

```json
{
  "message": "Hotel Deleted Successfully",
  "data": {},
  "success": true
}
```

---

## Update Hotel

### `PATCH /hotels/:id`

> **Updated route:** The current backend uses the hotel ID as a path parameter.

Example:

```text
PATCH /HotelService/api/v1/hotels/1
```

### Request Body

The current update DTO supports:

| Field | Type | Required |
|---|---|---|
| `name` | `string` | No |
| `address` | `string` | No |
| `location` | `string` | No |
| `price` | `number` | No |
| `room_type` | `RoomType` | No |

Example:

```json
{
  "name": "The Grand Horizon Renovated",
  "address": "221B Baker Street, London",
  "location": "London",
  "price": 2800,
  "room_type": "DELUXE"
}
```

### Success

HTTP `202 Accepted`.

```json
{
  "message": "Hotel Updated Successfully",
  "data": {},
  "success": true
}
```

> The update route currently has no Zod validation middleware. The accepted shape is therefore based on `updateHotelDTO` and the service/controller implementation.

---

# 4.3 Room APIs

## Get Available Rooms

### `GET /rooms/getAvailableRooms`

### Query Parameters

| Parameter | Type | Required |
|---|---|---|
| `roomCategoryId` | `string` | Yes |
| `checkInDate` | `string` | Yes |
| `checkOutDate` | `string` | Yes |

Example:

```text
GET /HotelService/api/v1/rooms/getAvailableRooms?roomCategoryId=3&checkInDate=2026-09-25&checkOutDate=2026-09-30
```

### Success

HTTP `201 Created`.

```json
{
  "message": "Rooms Fetched Successfully!",
  "data": [],
  "success": true
}
```

> The controller intentionally returns `201` for this GET endpoint.

---

## Update Room Availability

### `POST /rooms/update-rooms-id`

### Request Body

```json
{
  "bookingId": 42,
  "roomIds": [101, 102, 103]
}
```

### Success

HTTP `201 Created`.

```json
{
  "message": "Rooms updated Successfully!",
  "data": [3],
  "success": true
}
```

---

# 4.4 Room Generation

## Generate Rooms

### `POST /generateRooms`

This endpoint queues an asynchronous BullMQ room-generation job.

### Request Body

| Field | Type | Required | Default |
|---|---|---|---|
| `roomCategoryId` | positive number | Yes | — |
| `startDate` | ISO datetime | Yes | — |
| `endDate` | ISO datetime | Yes | — |
| `priceOverride` | positive number | No | Category price |
| `batchSize` | positive number | No | `100` |

Example:

```json
{
  "roomCategoryId": 3,
  "startDate": "2026-09-25T00:00:00.000Z",
  "endDate": "2026-09-30T23:59:59.000Z",
  "priceOverride": 1750,
  "batchSize": 50
}
```

### Success

HTTP `200 OK`.

```json
{
  "message": "Hotel Room updated Successfully",
  "data": {},
  "success": true
}
```

The response indicates that the job was accepted by the queue; it does not represent completion of room generation.

### Alias

The same handler is also registered at:

```text
POST /hotels/generateRooms
```

Both routes currently use `RoomGenerationJobSchema`.

> `scheduleType` and `scheduledAt` exist in `RoomGenerationRequestSchema`, but the registered HTTP route validates against `RoomGenerationJobSchema`. The active `/generateRooms` API therefore uses the fields documented above.

---

# 4.5 Scheduler

## Start Scheduler

### `POST /scheduler/start`

No request body.

### Success

```json
{
  "message": "Room availability extension scheduler started successfully",
  "success": true,
  "data": {
    "status": "started"
  }
}
```

---

## Stop Scheduler

### `POST /scheduler/stop`

No request body.

### Success

```json
{
  "message": "Room availability extension scheduler stopped successfully",
  "success": true,
  "data": {
    "status": "stopped"
  }
}
```

---

## Scheduler Status

### `GET /scheduler/status`

### Success

```json
{
  "message": "Scheduler status retrieved successfully",
  "success": true,
  "data": {
    "isRunning": true
  }
}
```

---

## Manual Availability Extension

### `POST /scheduler/extend`

No request body.

### Success

```json
{
  "message": "Manual room availability extension completed successfully",
  "success": true,
  "data": {
    "action": "manual_extension_completed"
  }
}
```

---

# 5. BookingService

Base path:

```text
/api/v1
```

Through the gateway:

```text
https://airbnb-auth.onrender.com/BookingService/api/v1
```

The BookingService uses Prisma and Redis Redlock.

---

# 5.1 Health

## `GET /ping`

```text
GET /BookingService/api/v1/ping
```

Expected:

```json
{
  "message": "pong"
}
```

---

# 5.2 Booking APIs

## Create Booking

### `POST /booking`

Creates a booking in `PENDING` state.

### Request Body

| Field | Type | Required |
|---|---|---|
| `userId` | `number` | Yes |
| `hotelId` | `number` | Yes |
| `roomCategoryId` | `number` | Yes |
| `totalGuests` | `number` | Yes, minimum 1 |
| `bookingAmount` | `number` | Yes, minimum 1 |
| `checkInDate` | `string` | Yes |
| `checkOutDate` | `string` | Yes |

Example:

```json
{
  "userId": 7,
  "hotelId": 1,
  "roomCategoryId": 3,
  "totalGuests": 2,
  "bookingAmount": 750,
  "checkInDate": "2026-09-25",
  "checkOutDate": "2026-09-30"
}
```

### Success

HTTP `200 OK`.

```json
{
  "bookingId": 42,
  "IdempotencyKey": "b3f2a1c0-4d5e-6f7a-8b9c-0d1e2f3a4b5c"
}
```

> Unlike most Node.js service responses, the current booking controller returns this object directly without `message`, `data`, or `success`.

---

## Confirm Booking

### `POST /booking/confirm/:idempotencyKey`

Example:

```text
POST /BookingService/api/v1/booking/confirm/b3f2a1c0-4d5e-6f7a-8b9c-0d1e2f3a4b5c
```

No request body.

### Success

HTTP `200 OK`.

```json
{
  "bookingId": 42,
  "status": "CONFIRMED"
}
```

---

## Get All Bookings For User

### `GET /booking/getAllBookings/:id`

This endpoint is present in the current backend and was added to the documentation.

Example:

```text
GET /BookingService/api/v1/booking/getAllBookings/7
```

Here `id` is passed to `getAllBookingsService()` as a number.

### Success

HTTP `200 OK`.

```json
{
  "message": "Booking fetched successfully!",
  "data": []
}
```

---

# 6. NotificationService

NotificationService is primarily a background worker.

It does not expose a public notification HTTP API.

## Health

The service registers:

```text
GET /api/v1/ping
GET /api/v2/ping
```

The v1 router currently imports the v2 ping router internally, so both registrations should be treated according to the current source rather than assumed to represent separate functionality.

## Email Queue Payload

The worker consumes a `NotificationDto`:

```typescript
interface NotificationDto {
  to: string;
  subject: string;
  templateId: string;
  params: Record<string, any>;
}
```

Example:

```json
{
  "to": "guest@example.com",
  "subject": "Your Booking Confirmation",
  "templateId": "welcome",
  "params": {
    "name": "Akash",
    "appName": "Booking App"
  }
}
```

The worker uses BullMQ/Redis and Nodemailer.

---

# 7. Inter-Service Communication

## BookingService → HotelService

BookingService communicates with HotelService using Axios.

### Get Available Rooms

```http
GET /api/v1/rooms/getAvailableRooms
```

Query:

```text
roomCategoryId
checkInDate
checkOutDate
```

### Update Room Availability

```http
POST /api/v1/rooms/update-rooms-id
```

Body:

```json
{
  "bookingId": 42,
  "roomIds": [101, 102]
}
```

## Booking Flow

```text
Client
  |
  | POST /booking
  v
BookingService
  |
  | acquire Redis Redlock
  |
  | GET available rooms
  v
HotelService
  |
  | return available rooms
  v
BookingService
  |
  | create PENDING booking
  | create idempotency key
  | update room booking IDs
  |
  v
Client
  |
  | bookingId + IdempotencyKey
  |
  | POST /booking/confirm/:idempotencyKey
  v
BookingService
  |
  | PENDING -> CONFIRMED
  v
Client
```

---

# 8. Data Models

## HotelService — Hotel

The active hotel creation validator requires:

```text
name
address
location
price
roomType
```

Optional:

```text
rating
ratingCount
```

## Room Category

```text
hotelId
roomType
roomNo
price
```

### RoomType

```text
SINGLE
DOUBLE
FAMILY
DELUXE
SUITE
```

## Room

```text
id
hotelId
roomCategoryId
roomNo
price
dateOfAvailability
bookingId
createdAt
updatedAt
deletedAt
```

## Booking

```text
id
userId
hotelId
checkInDate
checkOutDate
roomCategoryId
bookingAmount
totalGuests
status
createdAt
updatedAt
```

### BookingStatus

```text
PENDING
CONFIRMED
CANCELLED
```

## IdempotencyKey

```text
id
idemKey
bookingId
finalized
createdAt
updatedAt
```

---

# 9. Environment Variables

## HotelService

| Variable | Default |
|---|---|
| `PORT` | `3001` |
| `REDIS_PORT` | `6379` |
| `REDIS_HOST` | `localhost` |
| `ROOM_CRON` | `0 2 * * *` |
| `DB_HOST` | `localhost` |
| `DB_USER` | `root` |
| `DB_PASSWORD` | `root` |
| `DB_NAME` | `test_db` |
| `DB_PORT` | `10337` |

## BookingService

| Variable | Default |
|---|---|
| `PORT` | `3001` |
| `REDIS_SERVER_URL` | `redis://localhost:6379` |
| `LOCK_TTL` | `5000` |
| `HOTEL_SERVICE_URL` | `http://localhost:3002` |
| `DATABASE_URL` | — |

## NotificationService

| Variable | Default |
|---|---|
| `PORT` | `3001` |
| `REDIS_PORT` | `6379` |
| `REDIS_HOST` | `localhost` |
| `MAIL_USER` | empty |
| `MAIL_PASS` | empty |

## AuthInGo

The current Go service reads environment configuration including:

```text
PORT
JWT_SECRET
```

Database and other configuration values are loaded through the AuthInGo config packages.

---

# 10. Quick Reference

## AuthInGo

| Method | Path | Auth |
|---|---|---|
| GET | `/ping` | Public |
| POST | `/signup` | Public |
| POST | `/login` | Public |
| GET | `/profile` | JWT |
| GET | `/all` | JWT |
| DELETE | `/` | JWT |
| GET | `/roles/:id` | Public |
| GET | `/roles` | Public |
| POST | `/createRole` | Public |
| PATCH | `/roles/:id` | Public |
| GET | `/roleByName` | Public |
| GET | `/role/:id/permissions` | Public |
| POST | `/assignpermissionToRole` | Public |
| DELETE | `/removePermissionFromRole` | Public |
| POST | `/roles/:userId/assign/:roleId` | JWT + admin |
| DELETE | `/roles/:id` | Public |

## HotelService

| Method | Path |
|---|---|
| GET | `/api/v1/ping` |
| POST | `/api/v1/hotels` |
| GET | `/api/v1/hotels/getAllHotels` |
| GET | `/api/v1/hotels/:id` |
| DELETE | `/api/v1/hotels/:id` |
| PATCH | `/api/v1/hotels/:id` |
| POST | `/api/v1/generateRooms` |
| POST | `/api/v1/hotels/generateRooms` |
| GET | `/api/v1/rooms/getAvailableRooms` |
| POST | `/api/v1/rooms/update-rooms-id` |
| POST | `/api/v1/scheduler/start` |
| POST | `/api/v1/scheduler/stop` |
| GET | `/api/v1/scheduler/status` |
| POST | `/api/v1/scheduler/extend` |

## BookingService

| Method | Path |
|---|---|
| GET | `/api/v1/ping` |
| POST | `/api/v1/booking` |
| POST | `/api/v1/booking/confirm/:idempotencyKey` |
| GET | `/api/v1/booking/getAllBookings/:id` |

## NotificationService

| Method | Path |
|---|---|
| GET | `/api/v1/ping` |
| GET | `/api/v2/ping` |

---

> **Important:** This document reflects the backend source supplied with the current project snapshot. Where DTOs, validators, and controllers disagree, the documentation identifies the active validator/route behavior instead of silently treating the DTO as the API contract.

*Updated from the September 2026 Airbnb backend codebase.*
