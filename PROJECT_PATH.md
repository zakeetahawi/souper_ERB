# مسار المشروع - نظام ELKHAWAGA ERP

## 📍 المسار الحالي
```
/home/zakee/ERP_ELKHAWAGA/
```

## 🌐 الريبو الرسمي
```
https://github.com/zakeetahawi/souper_ERB.git
```

## 🚀 كيفية البدء

### 1. الانتقال للمجلد
```bash
# من الريبو الرسمي
git clone https://github.com/zakeetahawi/souper_ERB.git
cd souper_ERB

# أو من المسار المحلي
cd /home/zakee/ERP_ELKHAWAGA/
```

### 2. الإعداد السريع
```bash
# الإعداد الكامل
make setup

# تشغيل المشروع
make dev
```

### 3. الوصول للتطبيق
- **الفرونت إند**: http://localhost:4200
- **الباك إند API**: http://localhost:8080
- **بيانات الدخول**: admin / admin123

## 📁 هيكل المجلدات

```
/home/zakee/ERP_ELKHAWAGA/
├── 📄 README.md                    # الدليل الرئيسي
├── 📄 LICENSE                      # رخصة MIT
├── 📄 Makefile                     # أوامر التشغيل
├── 📄 SETUP.md                     # دليل الإعداد
├── 📄 ACHIEVEMENT_REPORT.md        # تقرير الإنجاز
├── 📄 PROJECT_PATH.md              # هذا الملف
├── 📄 .gitignore                   # ملفات Git المُستثناة
│
├── 📁 docs/                        # التوثيق
│   ├── 📄 architecture.md          # البنية المعمارية
│   ├── 📄 workflows.md             # سير العمل
│   └── 📁 api-contracts/           # عقود API
│       └── 📄 customers.md         # عقود موديول العملاء
│
├── 📁 backend/                     # الباك إند (Go)
│   ├── 📄 go.mod                   # تبعيات Go
│   ├── 📄 go.sum                   # checksums التبعيات
│   │
│   ├── 📁 cmd/server/              # نقطة الدخول
│   │   └── 📄 main.go              # الملف الرئيسي
│   │
│   └── 📁 internal/                # الكود الداخلي
│       ├── 📁 config/              # الإعدادات
│       │   └── 📄 config.go        # تحميل الإعدادات
│       │
│       ├── 📁 db/                  # قاعدة البيانات
│       │   ├── 📄 postgres.go      # اتصال PostgreSQL
│       │   ├── 📁 repo/            # النماذج
│       │   │   └── 📄 model.go     # نماذج GORM
│       │   └── 📁 seed/            # البيانات الأولية
│       │       └── 📄 egypt_locations.json
│       │
│       ├── 📁 auth/                # المصادقة والصلاحيات
│       │   ├── 📄 jwt.go           # JWT
│       │   ├── 📄 sessions.go      # إدارة الجلسات
│       │   ├── 📄 rbac.go          # نظام الصلاحيات
│       │   └── 📄 api.go           # API المصادقة
│       │
│       └── 📁 modules/             # الموديولات
│           ├── 📄 registry.go      # تسجيل الموديولات
│           └── 📁 customers/       # موديول العملاء
│               ├── 📄 model.go     # DTOs
│               ├── 📄 service.go   # منطق الأعمال
│               └── 📄 api.go       # API handlers
│
├── 📁 frontend/                    # الفرونت إند (Angular)
│   ├── 📄 package.json             # تبعيات Node.js
│   ├── 📄 angular.json             # إعدادات Angular
│   ├── 📄 tsconfig.json            # إعدادات TypeScript
│   │
│   └── 📁 src/                     # كود المصدر
│       ├── 📄 index.html           # الصفحة الرئيسية
│       ├── 📄 main.ts              # نقطة الدخول
│       ├── 📄 styles.scss          # الأنماط العامة
│       │
│       ├── 📁 environments/        # متغيرات البيئة
│       │   ├── 📄 environment.ts   # بيئة التطوير
│       │   └── 📄 environment.prod.ts
│       │
│       └── 📁 app/                 # التطبيق
│           ├── 📄 app.component.ts # المكون الرئيسي
│           ├── 📄 app.config.ts    # إعدادات التطبيق
│           ├── 📄 app-routing.module.ts
│           │
│           ├── 📁 core/            # الخدمات الأساسية
│           │   └── 📁 auth/        # المصادقة
│           │       ├── 📄 auth.service.ts
│           │       ├── 📄 auth.guard.ts
│           │       └── 📄 auth.interceptor.ts
│           │
│           └── 📁 pages/           # الصفحات
│               ├── 📁 login/       # صفحة تسجيل الدخول
│               ├── 📁 dashboard/   # لوحة التحكم
│               └── 📁 customers/   # إدارة العملاء
│
└── 📁 infra/                       # البنية التحتية
    ├── 📁 docker/                  # ملفات Docker
    │   ├── 📄 docker-compose.yml   # تكوين الخدمات
    │   ├── 📄 Dockerfile.backend   # صورة الباك إند
    │   ├── 📄 Dockerfile.frontend  # صورة الفرونت إند
    │   └── 📄 nginx.conf           # إعدادات Nginx
    │
    └── 📁 migrations/              # هجرات قاعدة البيانات
        └── 📄 0001_init.sql        # الهجرة الأولية
```

## 🔧 أوامر التشغيل

### أوامر Makefile المتاحة
```bash
# في مجلد المشروع
cd /home/zakee/ERP_ELKHAWAGA/

# عرض جميع الأوامر
make help

# الإعداد الكامل
make setup

# تثبيت التبعيات
make install

# تشغيل وضع التطوير
make dev

# بناء المشروع
make build

# تشغيل الاختبارات
make test

# تنظيف الملفات
make clean

# بناء Docker
make docker-build

# تشغيل Docker
make docker-run

# فحص حالة الخدمات
make status
```

## 🗄️ إعداد قاعدة البيانات

### إنشاء قاعدة البيانات
```bash
# إنشاء قاعدة البيانات والمستخدم
sudo -u postgres psql -c "CREATE DATABASE zakee_erp;"
sudo -u postgres psql -c "CREATE USER erp_user WITH PASSWORD 'password';"
sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE zakee_erp TO erp_user;"
```

### تشغيل ملف الهجرة
```bash
# تشغيل ملف الهجرة
psql -U erp_user -d zakee_erp -f infra/migrations/0001_init.sql
```

## 🔐 إعداد ملف البيئة

### إنشاء ملف .env للباك إند
```bash
# في مجلد backend
cd backend

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
```

## 🚀 تشغيل المشروع

### الطريقة الأولى: باستخدام Makefile
```bash
# الإعداد والتشغيل الكامل
make setup
make dev
```

### الطريقة الثانية: تشغيل منفصل
```bash
# تشغيل الباك إند
cd backend
go run cmd/server/main.go

# تشغيل الفرونت إند (في terminal آخر)
cd frontend
npm start
```

### الطريقة الثالثة: باستخدام Docker
```bash
# بناء وتشغيل بـ Docker
make docker-build
make docker-run
```

## 📊 حالة المشروع

### ✅ مكتمل
- ✅ البنية الأساسية
- ✅ نظام المصادقة والأمان
- ✅ موديول العملاء (API)
- ✅ التوثيق الشامل
- ✅ أوامر التشغيل

### 🔄 قيد التطوير
- 🔄 واجهة موديول العملاء (UI)
- 🔄 الموديولات الإضافية
- 🔄 نظام الفواتير
- 🔄 الاستيراد/التصدير

## 🆘 استكشاف الأخطاء

### مشاكل قاعدة البيانات
```bash
# التحقق من حالة PostgreSQL
sudo systemctl status postgresql

# إعادة تشغيل PostgreSQL
sudo systemctl restart postgresql
```

### مشاكل Redis
```bash
# التحقق من حالة Redis
sudo systemctl status redis

# إعادة تشغيل Redis
sudo systemctl restart redis
```

### مشاكل التبعيات
```bash
# تنظيف وإعادة تثبيت تبعيات Go
cd backend
go clean -cache -modcache
go mod download

# تنظيف وإعادة تثبيت تبعيات Node.js
cd frontend
rm -rf node_modules package-lock.json
npm install
```

## 📞 الدعم

للمساعدة والدعم، راجع:
- `README.md` - الدليل الرئيسي
- `SETUP.md` - دليل الإعداد المفصل
- `ACHIEVEMENT_REPORT.md` - تقرير الإنجاز

---

**آخر تحديث**: ديسمبر 2024  
**حالة المشروع**: ✅ مكتمل (Sprint 1) 