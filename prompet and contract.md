---
description: "ELKHAWAGA ERP — عقد صارم + برومبت كود شامل لبناء نظام ERP مدمج CRM (Go + Angular + PostgreSQL) مع دعم كامل للعربية والإنجليزية"
globs:
alwaysApply: true
---
---

# ملف «برومبت كود» رسمي لبناء نظام ERP مدمج CRM — مع **عَقْد صارِم**

> **ملاحظة مهمة للمُنفِّذ (نظام الذكاء الاصطناعي/الوكيل البرمجي مثل Cline أو Kroser AI):** هذا الملف هو **العقد والمرجعية الوحيدة** للتنفيذ. الالتزام **الحرفي** مطلوب. **يُمنع** القفز على أي بند أو تغيير التنسيق أو إهمال أي شرط. أي غموض يُحسم لصالح **الالتزام بالنص أدناه**.

---

## 0) بطاقة المشروع المختصرة

* **اسم المشروع:** ELKHAWAGA-ERP
* **الغرض:** نظام ERP متكامل مدمج معه CRM بجميع الموديولات الأساسية.
* **النواة:** تصميم **موديولات قابلة للتفعيل/التعطيل ديناميكيًا** بدون التأثير على بقية الموديولات، مع ترابط داخلي مضبوط (Events/Contracts).
* **الباك إند:** Go (Golang)
* **الواجهة:** Angular (حصراً)
* **قاعدة البيانات:** PostgreSQL (حصراً)
* **الـ UI:** SweetAlert2 (أو ما يعادله) للنوافذ المنبثقة، Select2 أو ما يعادله لحقل البحث المُوسّع.
* **مكتبات مُقترحة:**

  * Go: Fiber أو Gin + GORM/SQLC، Redis للجلسات/الكاش، Zap للّوجينغ.
  * Angular: Angular Material أو Tailwind + ngx-select أو ng-select (بديل Select2)، SweetAlert2.
  * بنية **Feature Flags** للتحكم بتفعيل/تعطيل الموديولات.
* **الأولوية في الإطلاق (Sprint-1):** شاشة تسجيل الدخول الإبداعية + **قسم العملاء** مكتمل الأساسيات. باقي الأقسام تُعرض برسالة «قيد التطوير».

---

## 1) العَقْد الصارم (غير قابل للتفاوض)

1. **عدم تجاوز الشروط الأساسية**: ممنوع تغيير التقنيات المحددة (Go/Angular/PostgreSQL) أو بدائل الواجهات (SweetAlert/Select2-بديله) دون نص صريح هنا.
2. **وحدة الهيدر/الفوتر/قائمة التطبيقات الجانبية**: قالب موحد لكل الموديولات. يُسمح فقط بتحديث محتوى المساحات الداخلية.
3. **ملف إعدادات موحد** (Backend + Frontend) مع دعم متغيّرات البيئة. لا تُنشأ ملفات إعدادات إضافية خارج الهيكل المبين لاحقًا.
4. **نظام الموديولات**: كل موديول **قابل للتفعيل/التعطيل** دون كسر الترابط. يجب أن يكون هناك جدول/خدمة FeatureFlags + حماية على مستوى الراوتر والواجهات.
5. **الأداء والاستعلامات**: التصميم يجب أن يدير جداول بـ ≥ 20 عمود و ≥ 20,000 صف بكفاءة. **ممنوع** استخدام جلب غير مُصفّى أو بدون Pagination/Indexing.
6. **الأمان والجلسات**: تسجيل الدخول **باسم مستخدم وكلمة مرور** فقط. **جلسة واحدة فعّالة لكل مستخدم**: إذا سجل الدخول من جهاز آخر، تُلغى الجلسة السابقة فورًا.
7. **RBAC ديناميكي**: أدوار وصلاحيات تتوسع تلقائيًا مع كل موديول جديد. ربط الصلاحيات بالدور تلقائيًا عند إنشاء الدور.
8. **تحكم برؤية الموديولات** حسب الدور/المستخدم.
9. **ثيمات UI (4 ثيمات على الأقل)** بينها ثيم ليلي، وباختلاف حقيقي لطريقة العرض/أماكن القوائم.
10. **شاشة رئيسية موحدة**: رسالة ترحيبية + لوغو الشركة (قابل للتغيير من الإعدادات) ويستخدم اللوغو ذاته في كل مكان يُطلب فيه لوغو.
11. **نظام طباعة فواتير** + مُنشئ قوالب احترافي (Preview + Edit + Save) مع إدارة قوالب متعددة.
12. **استيراد/تصدير Excel** + خيار **مزامنة فورية مع Google Sheets**، ولا تظهر الأزرار إلا عند تفعيلها من لوحة التحكم (Backend).
13. **منع التكرار**: يمنع إدراج أكثر من عميل بنفس رقم الهاتف. البحث يعتمد **رقم العميل** أولاً.
14. **ترقيم العميل**: الكود/ID = **رقم الفرع + رقم العميل**.
15. **التزام التنسيق**: يُمنع تغيير تنسيق هذا المستند أو تخطي الجداول/القوائم. أي مخرجات (كود/وثائق) يجب أن تتبع الهيكل المحدد.

---

## 2) هيكل المستودع (Monorepo)

```
repo-root/
├─ README.md
├─ LICENSE
├─ docs/
│  ├─ architecture.md
│  ├─ workflows.md
│  └─ api-contracts/
│     └─ customers.md
├─ infra/
│  ├─ docker/
│  │  ├─ Dockerfile.backend
│  │  ├─ Dockerfile.frontend
│  │  └─ docker-compose.yml
│  └─ migrations/
│     └─ 0001_init.sql
├─ backend/
│  ├─ cmd/
│  │  └─ server/main.go
│  ├─ internal/
│  │  ├─ config/
│  │  │  └─ config.go
│  │  ├─ db/
│  │  │  ├─ postgres.go
│  │  │  ├─ repo/ (SQLC أو GORM)
│  │  │  └─ seed/
│  │  │     └─ egypt_locations.json  ← (محافظات ومناطق مصر)
│  │  ├─ auth/
│  │  │  ├─ jwt.go
│  │  │  ├─ sessions.go (Single-Device Login)
│  │  │  └─ rbac.go
│  │  ├─ modules/
│  │  │  ├─ registry.go (Feature Flags)
│  │  │  ├─ customers/
│  │  │  │  ├─ api.go
│  │  │  │  ├─ service.go
│  │  │  │  └─ model.go
│  │  │  ├─ orders/
│  │  │  ├─ inventory/
│  │  │  ├─ field-survey/
│  │  │  ├─ factory/
│  │  │  ├─ installations/
│  │  │  ├─ data-sync/
│  │  │  ├─ maintenance/
│  │  │  ├─ development/
│  │  │  ├─ display-tuning/
│  │  │  ├─ sales/
│  │  │  ├─ marketing/
│  │  │  └─ projects/
│  │  ├─ billing/
│  │  │  ├─ templates/
│  │  │  └─ api.go (قوالب الفواتير)
│  │  ├─ settings/ (إعدادات الشركة + اللوغو)
│  │  ├─ export/
│  │  │  ├─ excel.go
│  │  │  └─ sheets.go (Google Sheets)
│  │  └─ ui/
│  │     └─ themes.go (تعريف الثيمات المتاحة)
│  └─ go.mod
└─ frontend/
   ├─ src/
   │  ├─ app/
   │  │  ├─ core/
   │  │  │  ├─ layout/ (Header/Footer/Sidebar موحّد)
   │  │  │  ├─ config.service.ts (Feature Flags)
   │  │  │  ├─ auth/ (Guards + Interceptors)
   │  │  │  └─ themes/ (4 ثيمات)
   │  │  ├─ shared/
   │  │  │  ├─ components/
   │  │  │  │  ├─ data-table/ (Virtual Scroll, Server Pagination)
   │  │  │  │  ├─ card-grid/
   │  │  │  │  ├─ select-advanced/ (ng-select أو ngx-select)
   │  │  │  │  └─ modals/ (SweetAlert2)
   │  │  │  └─ utils/
   │  │  ├─ pages/
   │  │  │  ├─ login/
   │  │  │  ├─ dashboard/
   │  │  │  ├─ customers/ (جدول/بطاقات + فلتر موسّع)
   │  │  │  └─ placeholders/ (شاشات «قيد التطوير» للموديولات الأخرى)
   │  │  └─ app-routing.module.ts
   │  └─ environments/
   │     ├─ environment.ts
   │     └─ environment.prod.ts
   └─ angular.json
```

---

## 3) متغيّرات البيئة (ملف إعدادات موحّد)

### Backend (.env)

```
APP_NAME=zakeeERP
APP_ENV=production
APP_PORT=8080
APP_BASE_URL=https://example.com

DB_HOST=postgres
DB_PORT=5432
DB_NAME=zakee_erp
DB_USER=erp_user
DB_PASS=********

REDIS_HOST=redis
REDIS_PORT=6379

JWT_SECRET=change_me
JWT_EXPIRES=15m
REFRESH_EXPIRES=720h

SINGLE_DEVICE_LOGIN=true
DEVICE_FINGERPRINT_HEADER=X-Device-Fingerprint

FEATURE_IMPORT_EXPORT=true
FEATURE_GOOGLE_SHEETS_SYNC=true

FILES_STORAGE=local
UPLOADS_DIR=/data/uploads

COMPANY_LOGO_PATH=/data/uploads/company/logo.png
HEADER_LOGO_PATH=/data/uploads/company/header_logo.png
```

### Frontend (environment.ts)

```ts
export const environment = {
  production: false,
  apiBaseUrl: 'http://localhost:8080/api',
  singleDeviceLogin: true,
  features: {
    importExport: true,
    googleSheetsSync: true,
  },
  themes: ['light', 'dark', 'compact', 'panelled'],
  defaultTheme: 'light'
};
```

---

## 4) مواصفات **الموديولات المشتركة**

### 4.1 نظام الأدوار والصلاحيات (RBAC)

* جداول: `roles`, `permissions`, `role_permissions`, `user_roles`.
* كل موديول عند تسجيله في **registry** يصرّح بصلاحياته (CRUD + نشاطات خاصة)، ويُضاف تلقائيًا.
* واجهة إدارة: إنشاء دور ⇒ تحديد الموديولات ⇒ تُدرج الصلاحيات تلقائيًا مع إمكانية تخصيص دقيق.

### 4.2 تسجيل الدخول والجلسات (Single-Device)

* تدفق:

  1. `POST /auth/login` يعيد **JWT + RefreshToken** ويوثّق **device\_fingerprint** (من الهيدر) + IP/UA.
  2. يُسجّل في جدول `user_sessions` مع حالة **active**.
  3. عند محاولة تسجيل دخول جديد لنفس المستخدم ⇒ تُلغى الجلسة النشطة السابقة (soft invalidate) وتبطل RefreshToken.
  4. Interceptor أمام كل طلب يحقق من الجلسة.
* في الواجهة: إذا استُخدم الحساب في جهاز آخر ⇒ تنبيه SweetAlert2 ثم إعادة توجيه لشاشة الدخول.

### 4.3 الثيمات (4 ثيمات)

* `light`, `dark`, `compact` (جداول مضغوطة/Typography أصغر), `panelled` (قوائم جانبية بلوحات كبيرة).
* تخزين اختيار المستخدم في LocalStorage + سيرفر (تفضيلات المستخدم).

### 4.4 القالب الموحد (Header/Footer/Sidebar)

* Sidebar = قائمة الموديولات **الفعلية المفعّلة فقط** حسب الدور + FeatureFlags.
* أيقونات موحدة، و Badge حالة كل موديول (مفعّل/متوقف).

### 4.5 النوافذ المنبثقة والبحث

* SweetAlert2 لجميع التأكيدات/التحذيرات.
* **Select2 بديل**: استخدم `ng-select` أو `ngx-select` مع **Server-side search** + **Virtual scroll**.

### 4.6 الاستيراد/التصدير والمزامنة

* أزرار Import/Export لا تظهر إلا إذا كانت `features.importExport === true`.
* زر "Google Sheets Sync" يظهر فقط إذا `features.googleSheetsSync === true`.
* Backend يوفر:

  * `POST /export/excel?entity=customers&filter=...`
  * `POST /import/excel?entity=customers` (يرد تقرير نتائج)
  * `POST /sync/google-sheets?entity=customers`

---

## 5) **قسم العملاء** (Sprint-1)

### 5.1 نموذج البيانات (Customers)

* **حقول أساسية (إلزامية):**

  * `customer_code` (مُولّد: **branch\_code + incremental\_number**)
  * `name` (اسم العميل)
  * `phone_primary` (رقم الهاتف الرئيسي، **فريد**)
  * `address_governorate` (من قائمة المحافظات)
  * `address_district` (من مناطق المحافظة المحددة)
  * `customer_type` (قابل للتهيئة من الباك إند)
  * `customer_classification` (قابل للتهيئة من الباك إند)
  * **إن كان التصنيف**: "شركة" أو "جهة حكومية" أو "مكتب هندسي" ⇒ إظهار حقل **المسؤول** + **هاتفه** مع زر `+` لإضافة أكثر من مسؤول.
  * `branch_id` (إذا كان المستخدم من الفرع الرئيسي ⇒ يختار من 15 فرعًا، وإلا يُثبت فرعه افتراضيًا).
* **حقول غير أساسية:**

  * `phone_secondary`
  * `interests` (وسوم/Tags)

### 5.2 القيود والقواعد

* **منع التكرار**: `UNIQUE(phone_primary)` + فحص على السيرفر.
* البحث يعتمد `customer_code` و/أو `phone_primary` **قبل الاسم**.
* فهارس PostgreSQL:

  * Index على `phone_primary`،
  * مركّب على (`branch_id`, `customer_code`),
  * GIN على `interests` إن خُزّنت كـ JSONB/Array.

### 5.3 واجهة العملاء (Angular)

* تبويب **جدول** (DataTable) + تبويب **بطاقات** (Card Grid).
* **فلتر موسّع** (Drawer أو Modal):

  * المحافظة/المنطقة (Cascading)
  * النوع، التصنيف
  * الفرع
  * مدى تاريخ الإنشاء
  * البحث بالهاتف/الكود
* الجدول يدعم: **Server-side pagination, sorting, filtering, virtual scroll**.
* البطاقات: عرض مختصر (الاسم، الهاتف، التصنيف، المحافظة/المنطقة، شارات).
* أزرار **استيراد/تصدير/مزامنة** تظهر فقط عند تفعيل الميزات.

### 5.4 REST API (عقود مختصرة)

```
GET   /api/customers          → list (query: page, size, sort, filters)
GET   /api/customers/{id}     → detail
POST  /api/customers          → create (يتحقق من uniqueness)
PUT   /api/customers/{id}     → update
DELETE/api/customers/{id}     → soft delete

GET   /api/geo/governorates   → [{code,name_ar,name_en}]
GET   /api/geo/districts?g=XX → [{code,name_ar,name_en}]

GET   /api/refs/types         → customer types
GET   /api/refs/classifications → customer classifications

POST  /api/export/excel?entity=customers
POST  /api/import/excel?entity=customers
POST  /api/sync/google-sheets?entity=customers
```

### 5.5 التحقق من الإدخال (Validation)

* الهاتف المصري: صيغة دولية/محلية مع تعويض الصفر/المقدمة تلقائيًا + Regex قوي.
* رفض إنشاء عميل جديد إذا كان `phone_primary` مستخدمًا.
* ربط المنطقة بالمحافظة (District must belong to governorate).

### 5.6 المحافظات والمناطق (Egypt)

* **مصدر البيانات**: ملف seed `internal/db/seed/egypt_locations.json`.
* **تنسيق مقترح:**

```json
{
  "governorates": [
    {"code": "C", "name_ar": "القاهرة", "name_en": "Cairo", "districts": ["المعادي", "حلوان", "مدينة نصر", "مصر الجديدة", "شبرا", "عين شمس", "التجمع" ]},
    {"code": "G", "name_ar": "الجيزة", "name_en": "Giza", "districts": ["الدقي", "العجوزة", "الهرم", "6 أكتوبر", "الشيخ زايد", "إمبابة"]},
    {"code": "ALX", "name_ar": "الإسكندرية", "name_en": "Alexandria", "districts": ["سيدي جابر", "محرم بك", "العجمي", "المنتزه", "المعمورة", "بحري"]}
    // ... إستكمال بقية الـ 27 محافظة بكافة المراكز/الأحياء
  ]
}
```

> **ملاحظة**: يجب تضمين **القائمة الكاملة** لكل محافظات مصر ومناطقها فعليًا داخل الملف `egypt_locations.json` أثناء التنفيذ، وليس في الواجهة مباشرةً. يُستهلك عبر API `geo/*` المذكور أعلاه.

---

## 6) شاشة تسجيل الدخول الإبداعية

* تصميم بملء الشاشة: لوحة فنية/Illustration بسيطة + حقلا **اسم المستخدم** و **كلمة المرور** + زر دخول و"إظهار/إخفاء كلمة المرور".
* مُعرّف الجهاز (Device Fingerprint) يُقرأ من الواجهة ويرسل في الهيدر `X-Device-Fingerprint`.
* حالات الخطأ (بيانات غير صحيحة/جلسة مستخدمة بجهاز آخر) تظهر بـ SweetAlert2.
* تذكرة: "نسيت كلمة المرور" (تدفق لاحق عبر البريد/OTP اختياري).

---

## 7) الأداء وتحسين الاستعلامات (قيود مُلزمة)

* **Pagination + Sorting + Filtering** دائمًا في السيرفر.
* **Partial text search** بالهاتف/الكود باستخدام فهارس مناسبة (BTREE/Trigram عند الحاجة).
* تفادي `SELECT *`، استخدام أعمدة محددة + DTOs.
* قياس الاستعلامات > 200ms ⇒ يُسجّل ويُحسّن.
* كاش اختياري لقوائم المراجع (أنواع/تصنيفات/محافظات) عبر Redis.

---

## 8) تفعيل/تعطيل الموديولات (Feature Flags)

* جدول `feature_flags`:

  * `module_key`, `enabled`, `visible_to_roles[]`.
* على الواجهة: خدمة `ConfigService` تُخفي/تُظهر الروابط وتمنع التوجيه لموديول معطّل.
* على الباك إند: Middleware يمنع الوصول لراوتر الموديول عند التعطيل.

---

## 9) قوالب الفواتير (Billing Templates)

* منشئ قوالب Drag\&Drop بسيط (حقول: شعار، بيانات شركة، جدول بنود، إجماليات، ختم، ملاحظات).
* **Preview حي** و **حفظ متعدد النسخ** + تعيين قالب افتراضي لكل فرع/دور.
* طباعة PDF/HTML.

---

## 10) وورك فلو (منظور التنفيذ)

```mermaid
flowchart TD
  A[تهيئة المستودع وهيكل المجلدات] --> B[إعداد Docker + Postgres + Redis]
  B --> C[بذور البيانات: المحافظات/المناطق + الفروع + الأنواع/التصنيفات]
  C --> D[تنفيذ Auth + Single-Device Sessions]
  D --> E[تفعيل RBAC + Feature Flags]
  E --> F[واجهة Angular: Layout موحّد + الثيمات]
  F --> G[صفحة Login الإبداعية]
  G --> H[موديول العملاء: API + UI (جدول/بطاقات/فلتر)]
  H --> I[استيراد/تصدير Excel + مزامنة Sheets]
  I --> J[شاشات بقية الموديولات Placeholder «قيد التطوير»]
  J --> K[اختبارات الأداء + التحسين]
  K --> L[توثيق العقود + نشر أولي]
```

---

## 11) قاعدة بيانات (مقتطف مخطط جداول)

```sql
-- الفروع (15 فرع كبداية)
CREATE TABLE branches (
  id SERIAL PRIMARY KEY,
  code VARCHAR(8) UNIQUE NOT NULL,
  name_ar TEXT NOT NULL
);

-- المستخدمون
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(64) UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  branch_id INT REFERENCES branches(id),
  is_active BOOLEAN DEFAULT TRUE
);

-- الجلسات (Single-Device)
CREATE TABLE user_sessions (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id) ON DELETE CASCADE,
  device_fingerprint TEXT NOT NULL,
  refresh_token TEXT UNIQUE NOT NULL,
  ip INET,
  user_agent TEXT,
  is_active BOOLEAN DEFAULT TRUE,
  created_at TIMESTAMPTZ DEFAULT now(),
  revoked_at TIMESTAMPTZ
);
CREATE INDEX ON user_sessions(user_id);

-- RBAC
CREATE TABLE roles (
  id SERIAL PRIMARY KEY,
  name VARCHAR(64) UNIQUE NOT NULL
);
CREATE TABLE permissions (
  id SERIAL PRIMARY KEY,
  module_key VARCHAR(64) NOT NULL,
  action VARCHAR(64) NOT NULL,
  UNIQUE(module_key, action)
);
CREATE TABLE role_permissions (
  role_id INT REFERENCES roles(id) ON DELETE CASCADE,
  permission_id INT REFERENCES permissions(id) ON DELETE CASCADE,
  PRIMARY KEY (role_id, permission_id)
);
CREATE TABLE user_roles (
  user_id INT REFERENCES users(id) ON DELETE CASCADE,
  role_id INT REFERENCES roles(id) ON DELETE CASCADE,
  PRIMARY KEY (user_id, role_id)
);

-- Feature Flags
CREATE TABLE feature_flags (
  module_key VARCHAR(64) PRIMARY KEY,
  enabled BOOLEAN NOT NULL DEFAULT TRUE
);

-- المراجع: الأنواع/التصنيفات
CREATE TABLE ref_customer_types (
  id SERIAL PRIMARY KEY,
  name_ar TEXT UNIQUE NOT NULL
);
CREATE TABLE ref_customer_classifications (
  id SERIAL PRIMARY KEY,
  name_ar TEXT UNIQUE NOT NULL
);

-- العملاء
CREATE TABLE customers (
  id BIGSERIAL PRIMARY KEY,
  customer_code VARCHAR(32) UNIQUE NOT NULL,
  name TEXT NOT NULL,
  phone_primary VARCHAR(20) UNIQUE NOT NULL,
  phone_secondary VARCHAR(20),
  governorate_code VARCHAR(16) NOT NULL,
  district_name TEXT NOT NULL,
  customer_type_id INT REFERENCES ref_customer_types(id),
  customer_classification_id INT REFERENCES ref_customer_classifications(id),
  branch_id INT REFERENCES branches(id) NOT NULL,
  interests TEXT[] DEFAULT '{}',
  created_at TIMESTAMPTZ DEFAULT now(),
  updated_at TIMESTAMPTZ DEFAULT now()
);
CREATE INDEX ON customers(phone_primary);
CREATE INDEX ON customers(branch_id, customer_code);
```



- كل إضافة جديدة يجب أن تحترم نفس الهيكل.  
- النظام ثنائي اللغة (AR/EN) إلزامياً في كل الموديولات.  
- الهيدر/الفوتر والقائمة الجانبية موحدة لجميع الموديولات.  

---
