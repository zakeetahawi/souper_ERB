# ELKHAWAGA ERP System

A comprehensive Enterprise Resource Planning (ERP) system built with Go (Backend) and Angular (Frontend), featuring integrated Customer Relationship Management (CRM) with full Arabic and English support.

## 🚀 Key Features

- **Dynamic Module System** - Activate/deactivate modules without affecting others
- **Advanced RBAC** - Role-based access control with dynamic permissions
- **Single Device Login** - One active session per user
- **4 UI Themes** - light, dark, compact, panelled
- **Full Bilingual Support** - Arabic and English with RTL support
- **Advanced Invoice System** - Customizable templates with drag & drop
- **Excel Import/Export** - With Google Sheets synchronization

## 🛠️ Technologies Used

- **Backend:** Go (Golang) + Fiber + GORM + PostgreSQL
- **Frontend:** Angular + Angular Material + Tailwind CSS
- **Database:** PostgreSQL
- **Cache:** Redis
- **UI Components:** SweetAlert2, ng-select

## 📋 Available Modules

- ✅ **Customers** - Complete customer management system
- 🔄 **Orders** - Order management (in development)
- 🔄 **Inventory** - Inventory management (in development)
- 🔄 **Field Survey** - Field survey management (in development)
- 🔄 **Factory** - Factory management (in development)
- 🔄 **Installations** - Installation management (in development)
- 🔄 **Data Sync** - Data synchronization (in development)
- 🔄 **Maintenance** - Maintenance management (in development)
- 🔄 **Development** - Development management (in development)
- 🔄 **Display Tuning** - Display tuning management (in development)
- 🔄 **Sales** - Sales management (in development)
- 🔄 **Marketing** - Marketing management (in development)
- 🔄 **Projects** - Project management (in development)

## ⚡ Quick Setup

### Using Makefile (Recommended)

```bash
# Clone the repository
git clone https://github.com/zakeetahawi/souper_ERB.git
cd souper_ERB

# Complete setup
make setup

# Run the project
make dev
```

### Manual Installation

```bash
# Clone the repository
git clone https://github.com/zakeetahawi/souper_ERB.git
cd souper_ERB

# Setup database
make create-db
make migrate

# Install dependencies
make install

# Setup environment
make env-backend

# Run the application
make dev
```

## 📋 Prerequisites

- Go 1.21+
- Node.js 18+
- PostgreSQL 15+
- Redis 7+

## 🌐 Access

- **Frontend**: http://localhost:4200
- **Backend API**: http://localhost:8080
- **Default Credentials**: admin / admin123

## 🔧 Available Makefile Commands

```bash
make help          # Show all available commands
make install       # Install dependencies
make build         # Build the project
make dev           # Run development mode
make test          # Run tests
make clean         # Clean temporary files
make docker-build  # Build Docker images
make docker-run    # Run with Docker
make setup         # Complete setup
make status        # Check service status
```

## 📁 Project Structure

```
souper_ERB/
├── backend/                 # Backend (Go)
│   ├── cmd/server/         # Application entry point
│   ├── internal/           # Internal packages
│   │   ├── auth/          # Authentication & authorization
│   │   ├── config/        # Configuration
│   │   ├── db/           # Database models & migrations
│   │   └── modules/      # Business modules
│   └── go.mod
├── frontend/               # Frontend (Angular)
│   ├── src/app/          # Application code
│   │   ├── core/         # Core services
│   │   ├── pages/        # Page components
│   │   └── shared/       # Shared components
│   └── package.json
├── infra/                 # Infrastructure
│   ├── docker/           # Docker configuration
│   └── migrations/       # Database migrations
├── docs/                 # Documentation
├── Makefile              # Build commands
└── SETUP.md              # Detailed setup guide
```

## 📚 API Documentation

### Authentication
- `POST /api/auth/login` - User login
- `POST /api/auth/refresh` - Refresh token
- `POST /api/auth/logout` - User logout
- `GET /api/auth/me` - Get current user

### Customers
- `GET /api/customers` - List customers
- `GET /api/customers/{id}` - Get customer details
- `POST /api/customers` - Create customer
- `PUT /api/customers/{id}` - Update customer
- `DELETE /api/customers/{id}` - Delete customer

### Reference Data
- `GET /api/geo/governorates` - List governorates
- `GET /api/geo/districts?g={code}` - List districts
- `GET /api/refs/types` - Customer types
- `GET /api/refs/classifications` - Customer classifications

## 🐳 Docker Deployment

```bash
# Build and run with Docker Compose
make docker-build
make docker-run
```

## 🔍 Troubleshooting

Refer to `SETUP.md` for detailed troubleshooting guide and common problem solutions.

## 🤝 Contributing

1. Fork the project
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 📞 Support

For help and support, check:
- `README.md` - Main guide (Arabic)
- `SETUP.md` - Detailed setup guide
- `ACHIEVEMENT_REPORT.md` - Achievement report
- `PROJECT_PATH.md` - Project path guide

---

**Last Updated**: December 2024  
**Project Status**: ✅ Complete (Sprint 1) 