# دليل إعداد وتشغيل نظام ELKHAWAGA ERP

## المتطلبات الأساسية

### للباك إند (Go)
- Go 1.21 أو أحدث
- PostgreSQL 15 أو أحدث
- Redis 7 أو أحدث

### للفرونت إند (Angular)
- Node.js 18 أو أحدث
- npm أو yarn

## خطوات الإعداد

### 1. إعداد قاعدة البيانات

```bash
# إنشاء قاعدة البيانات
sudo -u postgres psql -c "CREATE DATABASE zakee_erp;"
sudo -u postgres psql -c "CREATE USER erp_user WITH PASSWORD 'password';"
sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE zakee_erp TO erp_user;"

# تشغيل ملف الهجرة
psql -U erp_user -d zakee_erp -f infra/migrations/0001_init.sql
```

### 2. إعداد الباك إند

```bash
cd backend

# تثبيت التبعيات
go mod download

# إنشاء ملف .env
cat > .env << EOF
APP_NAME=zakeeERP
APP_ENV=development
APP_PORT=8080
APP_BASE_URL=http://localhost:8080

DB_HOST=localhost
DB_PORT=5432
DB_NAME=zakee_erp
DB_USER=erp_user
DB_PASS=password
DB_SSLMODE=disable

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_DB=0

JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
JWT_EXPIRES=15m
REFRESH_EXPIRES=720h

SINGLE_DEVICE_LOGIN=true
DEVICE_FINGERPRINT_HEADER=X-Device-Fingerprint

FEATURE_IMPORT_EXPORT=true
FEATURE_GOOGLE_SHEETS_SYNC=true

FILES_STORAGE=local
UPLOADS_DIR=./uploads

COMPANY_LOGO_PATH=./uploads/company/logo.png
HEADER_LOGO_PATH=./uploads/company/header_logo.png
EOF

# تشغيل الباك إند
go run cmd/server/main.go
```

### 3. إعداد الفرونت إند

```bash
cd frontend

# تثبيت التبعيات
npm install

# تشغيل الفرونت إند
npm start
```

## بيانات الدخول الافتراضية

- **اسم المستخدم:** admin
- **كلمة المرور:** admin123

## الوصول للتطبيق

- **الفرونت إند:** http://localhost:4200
- **الباك إند API:** http://localhost:8080
- **صفحة الصحة:** http://localhost:8080/health

## الميزات المتاحة

### ✅ مكتمل
- ✅ نظام المصادقة (JWT + Single Device Login)
- ✅ نظام الصلاحيات والأدوار (RBAC)
- ✅ إدارة العملاء (CRUD)
- ✅ نظام الموديولات (Feature Flags)
- ✅ واجهة مستخدم حديثة (Angular Material)
- ✅ دعم كامل للعربية (RTL)
- ✅ قاعدة بيانات PostgreSQL مع البيانات الأولية
- ✅ نظام الجلسات المتقدم

### 🔄 قيد التطوير
- 🔄 نظام الطلبات
- 🔄 إدارة المخزون
- 🔄 المسح الميداني
- 🔄 إدارة المصنع
- 🔄 نظام التركيبات
- 🔄 مزامنة البيانات
- 🔄 نظام الصيانة
- 🔄 إدارة التطوير
- 🔄 ضبط العرض
- 🔄 نظام المبيعات
- 🔄 التسويق
- 🔄 إدارة المشاريع

## هيكل المشروع

```
ERP_ELKHAWAGA/
├── backend/                 # الباك إند (Go)
│   ├── cmd/server/         # نقطة الدخول
│   ├── internal/           # الكود الداخلي
│   │   ├── auth/          # المصادقة والصلاحيات
│   │   ├── config/        # الإعدادات
│   │   ├── db/           # قاعدة البيانات
│   │   └── modules/      # الموديولات
│   └── go.mod
├── frontend/               # الفرونت إند (Angular)
│   ├── src/app/          # التطبيق
│   │   ├── core/         # الخدمات الأساسية
│   │   ├── pages/        # الصفحات
│   │   └── shared/       # المكونات المشتركة
│   └── package.json
├── infra/                 # البنية التحتية
│   ├── docker/           # ملفات Docker
│   └── migrations/       # ملفات الهجرة
└── docs/                 # التوثيق
```

## API Endpoints

### المصادقة
- `POST /api/auth/login` - تسجيل الدخول
- `POST /api/auth/refresh` - تجديد التوكن
- `POST /api/auth/logout` - تسجيل الخروج
- `GET /api/auth/me` - بيانات المستخدم الحالي

### العملاء
- `GET /api/customers` - قائمة العملاء
- `GET /api/customers/{id}` - تفاصيل العميل
- `POST /api/customers` - إنشاء عميل جديد
- `PUT /api/customers/{id}` - تحديث العميل
- `DELETE /api/customers/{id}` - حذف العميل

### البيانات المرجعية
- `GET /api/geo/governorates` - قائمة المحافظات
- `GET /api/geo/districts?g={code}` - قائمة المناطق
- `GET /api/refs/types` - أنواع العملاء
- `GET /api/refs/classifications` - تصنيفات العملاء

## استكشاف الأخطاء

### مشاكل قاعدة البيانات
```bash
# التحقق من حالة PostgreSQL
sudo systemctl status postgresql

# إعادة تشغيل PostgreSQL
sudo systemctl restart postgresql

# التحقق من الاتصال
psql -U erp_user -d zakee_erp -c "SELECT version();"
```

### مشاكل Redis
```bash
# التحقق من حالة Redis
sudo systemctl status redis

# إعادة تشغيل Redis
sudo systemctl restart redis

# اختبار الاتصال
redis-cli ping
```

### مشاكل الباك إند
```bash
# التحقق من التبعيات
go mod verify

# تنظيف الكاش
go clean -cache

# إعادة بناء
go build ./cmd/server
```

### مشاكل الفرونت إند
```bash
# تنظيف node_modules
rm -rf node_modules package-lock.json
npm install

# تنظيف كاش npm
npm cache clean --force
```

## المساهمة في التطوير

1. Fork المشروع
2. إنشاء فرع جديد للميزة
3. تطوير الميزة
4. إضافة الاختبارات
5. إنشاء Pull Request

## الترخيص

هذا المشروع مرخص تحت رخصة MIT. 