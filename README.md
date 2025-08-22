# نظام ELKHAWAGA ERP (Souper ERB)

نظام إدارة موارد المؤسسات (ERP) شامل مبني بـ Go (الباك إند) و Angular (الفرونت إند)، مع دعم كامل للعربية والإنجليزية.

## 🚀 المميزات الرئيسية

- **موديولات قابلة للتفعيل/التعطيل ديناميكيًا** بدون التأثير على بقية الموديولات
- **نظام صلاحيات متقدم (RBAC)** مع دعم الأدوار الديناميكية
- **جلسة واحدة فعالة لكل مستخدم** (Single-Device Login)
- **4 ثيمات مختلفة** (light, dark, compact, panelled)
- **دعم كامل للعربية والإنجليزية** مع RTL
- **نظام طباعة فواتير متقدم** مع قوالب قابلة للتخصيص
- **استيراد/تصدير Excel** ومزامنة مع Google Sheets

## 🛠️ التقنيات المستخدمة

- **Backend:** Go (Golang) + Fiber + GORM + PostgreSQL
- **Frontend:** Angular + Angular Material + Tailwind CSS
- **Database:** PostgreSQL
- **Cache:** Redis
- **UI Components:** SweetAlert2, ng-select

## 📋 الموديولات المتاحة

- ✅ **العملاء** (Customers) - مكتمل
- 🔄 **الطلبات** (Orders) - قيد التطوير
- 🔄 **المخزون** (Inventory) - قيد التطوير
- 🔄 **المسح الميداني** (Field Survey) - قيد التطوير
- 🔄 **المصنع** (Factory) - قيد التطوير
- 🔄 **التركيبات** (Installations) - قيد التطوير
- 🔄 **مزامنة البيانات** (Data Sync) - قيد التطوير
- 🔄 **الصيانة** (Maintenance) - قيد التطوير
- 🔄 **التطوير** (Development) - قيد التطوير
- 🔄 **ضبط العرض** (Display Tuning) - قيد التطوير
- 🔄 **المبيعات** (Sales) - قيد التطوير
- 🔄 **التسويق** (Marketing) - قيد التطوير
- 🔄 **المشاريع** (Projects) - قيد التطوير

## ⚡ التثبيت السريع

### باستخدام Makefile (مُوصى به)

```bash
# استنساخ المشروع
git clone https://github.com/zakeetahawi/souper_ERB.git
cd souper_ERB

# الإعداد الكامل
make setup

# تشغيل المشروع
make dev
```

### التثبيت اليدوي

```bash
# استنساخ المشروع
git clone https://github.com/zakeetahawi/souper_ERB.git
cd souper_ERB

# إعداد قاعدة البيانات
make create-db
make migrate

# تثبيت التبعيات
make install

# إعداد البيئة
make env-backend

# تشغيل التطبيق
make dev
```

## 📋 المتطلبات

- Go 1.21+
- Node.js 18+
- PostgreSQL 15+
- Redis 7+

## 🌐 الاستخدام

- **الفرونت إند**: http://localhost:4200
- **الباك إند API**: http://localhost:8080
- **بيانات الدخول الافتراضية**: admin / admin123

## 🔧 أوامر Makefile المتاحة

```bash
make help          # عرض جميع الأوامر
make install       # تثبيت التبعيات
make build         # بناء المشروع
make dev           # تشغيل وضع التطوير
make test          # تشغيل الاختبارات
make clean         # تنظيف الملفات
make docker-build  # بناء Docker
make docker-run    # تشغيل Docker
make setup         # الإعداد الكامل
make status        # فحص حالة الخدمات
```

## 📁 هيكل المشروع

```
souper_ERB/
├── backend/                 # الباك إند (Go)
│   ├── cmd/server/         # نقطة دخول التطبيق
│   ├── internal/           # الحزم الداخلية
│   │   ├── auth/          # المصادقة والتفويض
│   │   ├── config/        # الإعدادات
│   │   ├── db/           # نماذج قاعدة البيانات والهجرات
│   │   └── modules/      # موديولات الأعمال
│   └── go.mod
├── frontend/               # الفرونت إند (Angular)
│   ├── src/app/          # كود التطبيق
│   │   ├── core/         # الخدمات الأساسية
│   │   ├── pages/        # مكونات الصفحات
│   │   └── shared/       # المكونات المشتركة
│   └── package.json
├── infra/                 # البنية التحتية
│   ├── docker/           # إعدادات Docker
│   └── migrations/       # هجرات قاعدة البيانات
├── docs/                 # التوثيق
├── Makefile              # أوامر التشغيل
└── SETUP.md              # دليل الإعداد المفصل
```

## 📚 وثائق API

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

## 🐳 النشر بـ Docker

```bash
# البناء والتشغيل بـ Docker Compose
make docker-build
make docker-run
```

## 🔍 استكشاف الأخطاء

راجع ملف `SETUP.md` للحصول على دليل مفصل لاستكشاف الأخطاء وحل المشاكل الشائعة.

## 🤝 المساهمة

1. Fork المشروع
2. إنشاء فرع للميزة الجديدة
3. إجراء التغييرات
4. إضافة الاختبارات
5. إنشاء Pull Request

## 📄 الترخيص

هذا المشروع مرخص تحت رخصة MIT.

## 📞 الدعم

للمساعدة والدعم، راجع:
- `README_EN.md` - الدليل بالإنجليزية
- `SETUP.md` - دليل الإعداد المفصل
- `ACHIEVEMENT_REPORT.md` - تقرير الإنجاز
- `PROJECT_PATH.md` - دليل مسار المشروع
- `CONTRIBUTING.md` - دليل المساهمة
- `PROJECT_STATUS.md` - حالة المشروع

---

**آخر تحديث**: ديسمبر 2024  
**حالة المشروع**: ✅ مكتمل (Sprint 1)  
**الإصدار**: 1.0.0
