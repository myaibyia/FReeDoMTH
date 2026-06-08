# 📡 API Specification - FReeDoMTH Services

**API Version:** v1  
**Base URL:** `http://localhost:8080/api/v1` (dev) | `https://api.freedomth.com/api/v1` (prod)

---

## 🔐 Authentication

All endpoints except `/auth/*` and `/health` require JWT Bearer token.

```
Authorization: Bearer <jwt_token>
```

---

## 📋 Endpoints

### Health Check
```
GET /health
```
**Response:**
```json
{
  "status": "ok",
  "service": "freedomth-backend"
}
```

---

### Authentication

#### Register User
```
POST /auth/register
Content-Type: application/json
```

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "secure_password",
  "phone": "+66812345678",
  "name": "John Doe"
}
```

**Response (201):**
```json
{
  "id": "uuid",
  "email": "user@example.com",
  "phone": "+66812345678",
  "name": "John Doe",
  "created_at": "2026-06-08T10:00:00Z"
}
```

**Errors:**
- `400` - Invalid input
- `409` - Email already exists

---

#### Login
```
POST /auth/login
Content-Type: application/json
```

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "secure_password"
}
```

**Response (200):**
```json
{
  "access_token": "eyJhbGc...",
  "refresh_token": "eyJhbGc...",
  "expires_in": 3600,
  "user": {
    "id": "uuid",
    "email": "user@example.com"
  }
}
```

**Errors:**
- `401` - Invalid credentials
- `404` - User not found

---

### Orders

#### List Orders
```
GET /orders?page=1&limit=20&status=pending
Authorization: Bearer <token>
```

**Query Parameters:**
- `page` (optional) - Page number, default: 1
- `limit` (optional) - Items per page, default: 20
- `status` (optional) - Filter by status: pending, completed, cancelled

**Response (200):**
```json
{
  "total": 50,
  "page": 1,
  "limit": 20,
  "orders": [
    {
      "id": "uuid",
      "user_id": "uuid",
      "amount": "1500.00",
      "currency": "THB",
      "status": "pending",
      "items": [
        {
          "id": "item1",
          "name": "Product A",
          "quantity": 1,
          "price": "1500.00"
        }
      ],
      "metadata": {
        "channel": "web",
        "platform": "web"
      },
      "created_at": "2026-06-08T10:00:00Z",
      "updated_at": "2026-06-08T10:00:00Z"
    }
  ]
}
```

---

#### Create Order
```
POST /orders
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "items": [
    {
      "name": "Product A",
      "quantity": 1,
      "price": "1500.00"
    }
  ],
  "currency": "THB",
  "metadata": {
    "channel": "web"
  }
}
```

**Response (201):**
```json
{
  "id": "uuid",
  "user_id": "uuid",
  "amount": "1500.00",
  "currency": "THB",
  "status": "pending",
  "items": [...],
  "created_at": "2026-06-08T10:00:00Z"
}
```

**Errors:**
- `400` - Invalid input
- `401` - Unauthorized

---

#### Get Order
```
GET /orders/{id}
Authorization: Bearer <token>
```

**Response (200):**
```json
{
  "id": "uuid",
  "user_id": "uuid",
  "amount": "1500.00",
  "currency": "THB",
  "status": "pending",
  "items": [...],
  "created_at": "2026-06-08T10:00:00Z",
  "updated_at": "2026-06-08T10:00:00Z"
}
```

**Errors:**
- `404` - Order not found
- `401` - Unauthorized

---

### Payments

#### Create Payment
```
POST /payments
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "order_id": "uuid",
  "payment_method": "thai_bank",
  "currency": "THB",
  "metadata": {
    "channel": "web",
    "bank": "omise"
  }
}
```

**Response (201):**
```json
{
  "id": "uuid",
  "order_id": "uuid",
  "amount": "1500.00",
  "currency": "THB",
  "status": "pending",
  "payment_method": "thai_bank",
  "created_at": "2026-06-08T10:00:00Z"
}
```

**Errors:**
- `400` - Invalid input
- `404` - Order not found
- `401` - Unauthorized

---

#### Get Payment Status
```
GET /payments/{id}
Authorization: Bearer <token>
```

**Response (200):**
```json
{
  "id": "uuid",
  "order_id": "uuid",
  "amount": "1500.00",
  "currency": "THB",
  "status": "completed",
  "payment_method": "thai_bank",
  "gateway_response": {
    "transaction_id": "txn_123456",
    "timestamp": "2026-06-08T10:05:00Z"
  },
  "created_at": "2026-06-08T10:00:00Z",
  "updated_at": "2026-06-08T10:05:00Z"
}
```

**Errors:**
- `404` - Payment not found
- `401` - Unauthorized

---

### Webhooks

#### Payment Webhook
```
POST /webhooks/payment
Content-Type: application/json
X-Webhook-Signature: <hmac-sha256>
```

**Request Body (from payment gateway):**
```json
{
  "event": "payment.completed",
  "payment_id": "uuid",
  "transaction_id": "txn_123456",
  "status": "completed",
  "amount": "1500.00",
  "currency": "THB",
  "timestamp": "2026-06-08T10:05:00Z"
}
```

**Response (200):**
```json
{
  "status": "received",
  "message": "Webhook processed successfully"
}
```

**Errors:**
- `400` - Invalid webhook data
- `401` - Invalid signature

---

### Plugins

#### List Active Plugins
```
GET /plugins
Authorization: Bearer <token>
```

**Response (200):**
```json
{
  "plugins": [
    {
      "name": "web",
      "status": "active",
      "version": "1.0.0",
      "endpoints": ["POST /api/v1/plugins/web/webhook"]
    },
    {
      "name": "facebook",
      "status": "active",
      "version": "1.0.0",
      "endpoints": ["POST /api/v1/plugins/facebook/webhook"]
    },
    {
      "name": "line",
      "status": "active",
      "version": "1.0.0",
      "endpoints": ["POST /api/v1/plugins/line/webhook"]
    },
    {
      "name": "telegram",
      "status": "active",
      "version": "1.0.0",
      "endpoints": ["POST /api/v1/plugins/telegram/webhook"]
    },
    {
      "name": "wechat",
      "status": "active",
      "version": "1.0.0",
      "endpoints": ["POST /api/v1/plugins/wechat/webhook"]
    }
  ]
}
```

---

#### Plugin Webhook (Generic)
```
POST /plugins/{plugin_name}/webhook
Content-Type: application/json
```

**Request Body (plugin-specific):**
```json
{
  "event": "message",
  "user_id": "fb_user_123",
  "platform_user_id": "123456789",
  "message": "I want to make a payment",
  "metadata": {}
}
```

**Response (200):**
```json
{
  "status": "received",
  "message": "Plugin webhook processed"
}
```

---

## 🔄 Payment Flow

### Standard Flow

```
1. User creates order
   POST /orders → Order ID

2. User initiates payment
   POST /payments → Payment ID (status: pending)

3. Payment gateway processes
   (User redirected to bank/crypto)

4. Gateway sends webhook
   POST /webhooks/payment
   (Payment status: completed/failed)

5. Frontend polls or uses SSE
   GET /payments/{id} → Updated status

6. Order status updated
   (Order status: paid/failed)
```

---

## 🔒 Error Codes

| Code | Meaning |
|------|---------|
| 200 | OK |
| 201 | Created |
| 400 | Bad Request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not Found |
| 409 | Conflict |
| 500 | Internal Server Error |
| 503 | Service Unavailable |

---

## 📝 Pagination

List endpoints support pagination:

```
Query parameters:
- page: Page number (default: 1)
- limit: Items per page (default: 20, max: 100)

Response includes:
- total: Total items count
- page: Current page
- limit: Items per page
```

---

## 🔑 Rate Limiting

```
X-RateLimit-Limit: 1000
X-RateLimit-Remaining: 999
X-RateLimit-Reset: 1623000000
```

Default: 1000 requests per hour per user

---

**API Specification Version:** 1.0  
**Last Updated:** 2026-06-08
