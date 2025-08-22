# عقود API - العملاء

## نظرة عامة

هذا المستند يحدد عقود API الخاصة بموديول العملاء في نظام ELKHAWAGA ERP.

## النقاط النهائية (Endpoints)

### 1. قائمة العملاء

#### GET /api/customers

**الوصف:** جلب قائمة العملاء مع دعم التصفية والترتيب والتصفح.

**المعاملات:**
- `page` (int, optional): رقم الصفحة (افتراضي: 1)
- `size` (int, optional): عدد العناصر في الصفحة (افتراضي: 20)
- `sort` (string, optional): حقل الترتيب (افتراضي: created_at)
- `order` (string, optional): اتجاه الترتيب (asc/desc, افتراضي: desc)
- `search` (string, optional): البحث في الاسم أو رقم الهاتف
- `governorate` (string, optional): تصفية حسب المحافظة
- `district` (string, optional): تصفية حسب المنطقة
- `type` (int, optional): تصفية حسب النوع
- `classification` (int, optional): تصفية حسب التصنيف
- `branch` (int, optional): تصفية حسب الفرع

**الاستجابة:**
```json
{
  "success": true,
  "data": {
    "customers": [
      {
        "id": 1,
        "customer_code": "C001",
        "name": "أحمد محمد",
        "phone_primary": "+201234567890",
        "phone_secondary": "+201234567891",
        "governorate_code": "C",
        "governorate_name": "القاهرة",
        "district_name": "المعادي",
        "customer_type": {
          "id": 1,
          "name_ar": "فرد",
          "name_en": "Individual"
        },
        "customer_classification": {
          "id": 1,
          "name_ar": "عميل عادي",
          "name_en": "Regular Customer"
        },
        "branch": {
          "id": 1,
          "code": "MAIN",
          "name_ar": "الفرع الرئيسي"
        },
        "interests": ["تقنية", "تصميم"],
        "created_at": "2024-01-01T10:00:00Z",
        "updated_at": "2024-01-01T10:00:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "size": 20,
      "total": 100,
      "total_pages": 5
    }
  }
}
```

### 2. تفاصيل العميل

#### GET /api/customers/{id}

**الوصف:** جلب تفاصيل عميل محدد.

**المعاملات:**
- `id` (int, required): معرف العميل

**الاستجابة:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "customer_code": "C001",
    "name": "أحمد محمد",
    "phone_primary": "+201234567890",
    "phone_secondary": "+201234567891",
    "governorate_code": "C",
    "governorate_name": "القاهرة",
    "district_name": "المعادي",
    "customer_type": {
      "id": 1,
      "name_ar": "فرد",
      "name_en": "Individual"
    },
    "customer_classification": {
      "id": 1,
      "name_ar": "عميل عادي",
      "name_en": "Regular Customer"
    },
    "branch": {
      "id": 1,
      "code": "MAIN",
      "name_ar": "الفرع الرئيسي"
    },
    "interests": ["تقنية", "تصميم"],
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z"
  }
}
```

### 3. إنشاء عميل جديد

#### POST /api/customers

**الوصف:** إنشاء عميل جديد.

**البيانات المطلوبة:**
```json
{
  "name": "أحمد محمد",
  "phone_primary": "+201234567890",
  "phone_secondary": "+201234567891",
  "governorate_code": "C",
  "district_name": "المعادي",
  "customer_type_id": 1,
  "customer_classification_id": 1,
  "branch_id": 1,
  "interests": ["تقنية", "تصميم"]
}
```

**الاستجابة:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "customer_code": "C001",
    "message": "تم إنشاء العميل بنجاح"
  }
}
```

### 4. تحديث العميل

#### PUT /api/customers/{id}

**الوصف:** تحديث بيانات عميل موجود.

**البيانات المطلوبة:**
```json
{
  "name": "أحمد محمد محمود",
  "phone_primary": "+201234567890",
  "phone_secondary": "+201234567891",
  "governorate_code": "C",
  "district_name": "المعادي",
  "customer_type_id": 1,
  "customer_classification_id": 1,
  "branch_id": 1,
  "interests": ["تقنية", "تصميم", "برمجة"]
}
```

**الاستجابة:**
```json
{
  "success": true,
  "data": {
    "message": "تم تحديث العميل بنجاح"
  }
}
```

### 5. حذف العميل

#### DELETE /api/customers/{id}

**الوصف:** حذف عميل (حذف ناعم).

**الاستجابة:**
```json
{
  "success": true,
  "data": {
    "message": "تم حذف العميل بنجاح"
  }
}
```

## النقاط النهائية المرجعية

### 1. المحافظات

#### GET /api/geo/governorates

**الاستجابة:**
```json
{
  "success": true,
  "data": [
    {
      "code": "C",
      "name_ar": "القاهرة",
      "name_en": "Cairo"
    },
    {
      "code": "G",
      "name_ar": "الجيزة",
      "name_en": "Giza"
    }
  ]
}
```

### 2. المناطق

#### GET /api/geo/districts?g={governorate_code}

**الاستجابة:**
```json
{
  "success": true,
  "data": [
    {
      "code": "MAADI",
      "name_ar": "المعادي",
      "name_en": "Maadi"
    },
    {
      "code": "HELWAN",
      "name_ar": "حلوان",
      "name_en": "Helwan"
    }
  ]
}
```

### 3. أنواع العملاء

#### GET /api/refs/types

**الاستجابة:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name_ar": "فرد",
      "name_en": "Individual"
    },
    {
      "id": 2,
      "name_ar": "شركة",
      "name_en": "Company"
    }
  ]
}
```

### 4. تصنيفات العملاء

#### GET /api/refs/classifications

**الاستجابة:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "name_ar": "عميل عادي",
      "name_en": "Regular Customer"
    },
    {
      "id": 2,
      "name_ar": "عميل VIP",
      "name_en": "VIP Customer"
    }
  ]
}
```

## رموز الأخطاء

| الكود | الوصف |
|-------|--------|
| 400 | بيانات غير صحيحة |
| 401 | غير مصرح |
| 403 | محظور |
| 404 | غير موجود |
| 409 | تكرار في البيانات |
| 422 | بيانات غير صالحة |
| 500 | خطأ في الخادم |

## أمثلة الأخطاء

```json
{
  "success": false,
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "رقم الهاتف مطلوب",
    "details": {
      "phone_primary": ["رقم الهاتف مطلوب"]
    }
  }
}
```

```json
{
  "success": false,
  "error": {
    "code": "DUPLICATE_PHONE",
    "message": "رقم الهاتف مستخدم بالفعل"
  }
}
``` 