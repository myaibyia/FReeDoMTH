# 🏗️ Architecture - FReeDoMTH Services

---

## 📐 System Design

### Core Concept
**Omnichannel Payment System** - Single backend, multiple frontends, plugin-based extensibility.

```
┌─────────────────────────────────────────┐
│         Go Backend (Central)            │
│  - REST API                             │
│  - Business Logic                       │
│  - Payment Processing                   │
│  - Plugin Management                    │
└──────────┬──────────────────────────────┘
           │
    ┌──────┴──────┐
    ↓             ↓
PostgreSQL     Redis (Cache)
+ JSONB
    │
    └──────┬──────────────────────────┐
           │                          │
      ┌────▼────┐          ┌─────────▼──────┐
      │ Vue Web │          │ Plugin Layer   │
      └────┬────┘          │ (FB/Line/TG)   │
           │               └─────────┬──────┘
      ┌────▼───────────┬──────────────┴─────┐
      ↓                ↓                     ↓
  Webhook        API Gateway          WeChat SDK
```

---

## 🔄 Data Flow

### Payment Transaction
```
1. User selects channel (Web/FB/Line/TG/WeChat)
2. Plugin sends order → Backend API
3. Backend validates + creates payment
4. Payment Router checks:
   - Language detection
   - Amount calculation (decimal.Decimal)
   - Currency preference
5. Route to:
   - Thai Banking (Omise/2C2P)
   - Crypto (if applicable)
6. Webhook callback → Update DB
7. Send confirmation to user
```

---

## 📦 Backend Structure (Go)

```
backend/
├── main.go                 # Entry point
├── config/
│   ├── config.go           # Load .env
│   └── database.go         # DB connection
├── models/
│   ├── user.go             # User entity
│   ├── payment.go          # Payment entity
│   ├── order.go            # Order entity
│   └── transaction.go      # Transaction entity
├── handlers/
│   ├── auth.go             # Authentication
│   ├── payment.go          # Payment endpoints
│   ├── order.go            # Order endpoints
│   └── webhook.go          # Webhook handlers
├── services/
│   ├── payment_service.go  # Payment logic
│   ├── order_service.go    # Order logic
│   ├── payment_router.go   # Route to bank/crypto
│   └── notification.go     # Send notifications
├── plugins/
│   ├── facebook.go         # FB plugin
│   ├── line.go             # Line plugin
│   ├── telegram.go         # TG plugin
│   ├── wechat.go           # WeChat plugin
│   └── plugin_interface.go # Plugin contract
├── middleware/
│   ├── auth.go             # Auth middleware
│   ├── cors.go             # CORS
│   └── logging.go          # Request logging
├── utils/
│   ├── decimal.go          # Decimal math
│   └── crypto.go           # Encryption
└── go.mod                  # Dependencies
```

---

## 🎨 Frontend Structure (Vue)

```
frontend/
├── package.json
├── src/
│   ├── main.js             # Entry point
│   ├── App.vue
│   ├── components/
│   │   ├── PaymentForm.vue
│   │   ├── OrderList.vue
│   │   ├── UserProfile.vue
│   │   └── shared/         # Reusable
│   ├── pages/
│   │   ├── Dashboard.vue
│   │   ├── Payment.vue
│   │   └── Orders.vue
│   ├── plugins/
│   │   ├── web.js          # Web frontend
│   │   ├── facebook.js     # FB adapter
│   │   ├── line.js         # Line adapter
│   │   ├── telegram.js     # TG adapter
│   │   └── wechat.js       # WeChat adapter
│   ├── services/
│   │   ├── api.js          # HTTP client
│   │   ├── auth.js         # Auth service
│   │   └── payment.js      # Payment service
│   ├── store/              # State management
│   └── styles/
├── public/
└── Dockerfile
```

---

## 💾 Database Schema (PostgreSQL + JSONB)

### Users Table
```sql
CREATE TABLE users (
  id UUID PRIMARY KEY,
  email VARCHAR(255) UNIQUE,
  phone VARCHAR(20),
  password_hash TEXT,
  profile JSONB,  -- Flexible user data
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
```

### Orders Table
```sql
CREATE TABLE orders (
  id UUID PRIMARY KEY,
  user_id UUID REFERENCES users(id),
  amount NUMERIC(19,4),  -- decimal.Decimal
  currency VARCHAR(3),
  status VARCHAR(50),
  items JSONB,  -- Flexible order items
  metadata JSONB,  -- Channel info, etc.
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
```

### Payments Table
```sql
CREATE TABLE payments (
  id UUID PRIMARY KEY,
  order_id UUID REFERENCES orders(id),
  amount NUMERIC(19,4),
  currency VARCHAR(3),
  status VARCHAR(50),
  payment_method VARCHAR(50),  -- 'thai_bank', 'crypto'
  gateway_response JSONB,
  webhook_data JSONB,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
```

### Transactions Table
```sql
CREATE TABLE transactions (
  id UUID PRIMARY KEY,
  payment_id UUID REFERENCES payments(id),
  transaction_hash VARCHAR(255),
  status VARCHAR(50),
  details JSONB,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
```

---

## 🔌 Plugin System

### Plugin Interface (Go)
```go
type Platform interface {
  GetName() string
  SendMessage(userID, message string) error
  HandleWebhook(data []byte) error
  GetUserProfile(userID string) (Profile, error)
}
```

### Adding New Plugin
1. **Backend**: Implement `Platform` interface
2. **Frontend**: Create Vue adapter component
3. **Register**: Add to plugin manager
4. **Test**: Unit + integration tests

---

## 🛡️ Security

### Authentication
- JWT tokens
- Refresh token rotation
- Rate limiting per API endpoint

### Payment Security
- PCI compliance considerations
- Encrypted sensitive data
- Webhook signature verification
- HTTPS only

### Database
- SQL injection prevention (parameterized queries)
- Role-based access (RBAC)
- Audit logging

---

## 📊 API Endpoints

### Authentication
- `POST /api/auth/register` - Register user
- `POST /api/auth/login` - Login
- `POST /api/auth/refresh` - Refresh token

### Orders
- `GET /api/orders` - List orders
- `POST /api/orders` - Create order
- `GET /api/orders/:id` - Get order details
- `PUT /api/orders/:id` - Update order

### Payments
- `POST /api/payments` - Create payment
- `GET /api/payments/:id` - Get payment status
- `POST /api/webhooks/payment` - Payment callback

### Plugins
- `GET /api/plugins` - List active plugins
- `POST /api/plugins/:name/webhook` - Plugin webhook

---

## 🚀 Deployment

### Development
```bash
docker-compose up
```

### Production
- Use managed PostgreSQL
- Deploy Go backend on Cloud Run / ECS / K8s
- Deploy Vue frontend on Netlify / Vercel / S3 + CloudFront
- Redis on managed service
- Webhook ingress via API Gateway

---

## 📈 Scalability

### Current Design Supports
- ✅ 1000s concurrent users
- ✅ Multi-region deployment
- ✅ Horizontal scaling (stateless backend)
- ✅ Database read replicas

### Future Improvements
- Event streaming (Kafka)
- GraphQL layer
- Microservices split
- Service mesh (Istio)

---

## 🔐 Monitoring

- Application metrics (Prometheus)
- Logs (ELK Stack)
- Tracing (Jaeger)
- Alerting (PagerDuty)

---

**Architecture Version:** 1.0  
**Last Updated:** 2026-06-08
