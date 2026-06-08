# 🚀 Services By FReeDoMTH

**Service Unlock - Omnichannel Payment System**

---

## 📋 Tech Stack

| Layer | Technology | Notes |
|-------|-----------|-------|
| **Backend** | Go | Type-safe, Fast, Modular |
| **Database** | PostgreSQL + JSONB | ACID-safe, Flexible schema |
| **Frontend** | Vue.js | Lightweight, Reusable |
| **Plugin System** | Webhook + API | Extensible |
| **Payment** | Custom Integration | Thai Banks + Crypto |

---

## 🌍 Supported Channels

- 🌐 **Web** (Full)
- 📘 **Facebook** (Basic)
- 💬 **Line** (Basic)
- ✈️ **Telegram** (Basic)
- 📱 **WeChat** (Basic)
- 🔌 **Future Platforms** (Plugin-ready)

---

## 🏗️ Architecture

```
┌──────────────────────────────┐
│   Go Backend (Modular)       │
│  ✅ decimal.Decimal (Payment)│
│  ✅ DHRU API integration     │
│  ✅ Plugin system ready      │
└──────┬───────────────────────┘
       │ (REST API)
   PostgreSQL + JSONB
       │
    ┌──┴──┬──────┬─────┬─────┬────┐
    ↓     ↓      ↓     ↓     ↓    ↓
  Vue Plugins (Web, FB, Line, TG, WeChat)
```

---

## 📁 Project Structure

```
FReeDoMTH/
├── README.md (this file)
├── .gitignore
├── docker-compose.yml
├── backend/
│   ├── go.mod
│   ├── main.go
│   ├── config/
│   ├── handlers/
│   ├── models/
│   ├── services/
│   └── plugins/
├── frontend/
│   ├── package.json
│   ├── src/
│   │   ├── components/
│   │   ├── pages/
│   │   └── plugins/
│   └── public/
├── database/
│   └── migrations/
└── docs/
    ├── architecture.md
    └── api-spec.md
```

---

## 🚀 Quick Start

### Backend (Go)
```bash
cd backend
go mod download
go run main.go
```

### Frontend (Vue)
```bash
cd frontend
npm install
npm run dev
```

### Database
```bash
docker-compose up postgres
```

---

## 💳 Payment Features

- ✅ Thai Banking Integration
- ✅ Cryptocurrency Support
- ✅ Decimal Precision (No floating-point errors)
- ✅ Multi-currency Support

---

## 🔌 Plugin System

### Add New Platform
1. Create plugin in `backend/plugins/`
2. Implement Webhook handler
3. Register in API
4. Deploy frontend component in `frontend/src/plugins/`

---

## 📚 Documentation

- [Architecture](./docs/architecture.md)
- [API Specification](./docs/api-spec.md)

---

## 👤 Author

**FReeDoMTH**

---

## 📝 License

MIT

---

**Status:** 🚧 In Development

Last Updated: 2026-06-08
