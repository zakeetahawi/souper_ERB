-- الهجرة الأولية لقاعدة بيانات ELKHAWAGA ERP
-- إنشاء الجداول الأساسية

-- 1. جدول الفروع
CREATE TABLE IF NOT EXISTS branches (
    id SERIAL PRIMARY KEY,
    code VARCHAR(8) UNIQUE NOT NULL,
    name_ar TEXT NOT NULL,
    name_en TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- 2. جدول المستخدمين
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(64) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    branch_id INT REFERENCES branches(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- 3. جدول جلسات المستخدمين
CREATE TABLE IF NOT EXISTS user_sessions (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    device_fingerprint TEXT NOT NULL,
    refresh_token TEXT UNIQUE NOT NULL,
    ip INET,
    user_agent TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    revoked_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- 4. جدول الأدوار
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- 5. جدول الصلاحيات
CREATE TABLE IF NOT EXISTS permissions (
    id SERIAL PRIMARY KEY,
    module_key VARCHAR(64) NOT NULL,
    action VARCHAR(64) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    UNIQUE(module_key, action)
);

-- 6. جدول علاقة المستخدمين بالأدوار
CREATE TABLE IF NOT EXISTS user_roles (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    role_id INT REFERENCES roles(id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, role_id)
);

-- 7. جدول علاقة الأدوار بالصلاحيات
CREATE TABLE IF NOT EXISTS role_permissions (
    role_id INT REFERENCES roles(id) ON DELETE CASCADE,
    permission_id INT REFERENCES permissions(id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);

-- 8. جدول أعلام الميزات
CREATE TABLE IF NOT EXISTS feature_flags (
    module_key VARCHAR(64) PRIMARY KEY,
    enabled BOOLEAN NOT NULL DEFAULT TRUE
);

-- 9. جدول أنواع العملاء المرجعية
CREATE TABLE IF NOT EXISTS ref_customer_types (
    id SERIAL PRIMARY KEY,
    name_ar TEXT UNIQUE NOT NULL,
    name_en TEXT
);

-- 10. جدول تصنيفات العملاء المرجعية
CREATE TABLE IF NOT EXISTS ref_customer_classifications (
    id SERIAL PRIMARY KEY,
    name_ar TEXT UNIQUE NOT NULL,
    name_en TEXT
);

-- 11. جدول العملاء
CREATE TABLE IF NOT EXISTS customers (
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
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

-- 12. جدول المحافظات
CREATE TABLE IF NOT EXISTS governorates (
    code VARCHAR(16) PRIMARY KEY,
    name_ar TEXT NOT NULL,
    name_en TEXT,
    districts TEXT[]
);

-- إنشاء الفهارس
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
CREATE INDEX IF NOT EXISTS idx_users_branch_id ON users(branch_id);
CREATE INDEX IF NOT EXISTS idx_user_sessions_user_id ON user_sessions(user_id);
CREATE INDEX IF NOT EXISTS idx_user_sessions_refresh_token ON user_sessions(refresh_token);
CREATE INDEX IF NOT EXISTS idx_customers_phone_primary ON customers(phone_primary);
CREATE INDEX IF NOT EXISTS idx_customers_branch_id_customer_code ON customers(branch_id, customer_code);
CREATE INDEX IF NOT EXISTS idx_customers_governorate_code ON customers(governorate_code);
CREATE INDEX IF NOT EXISTS idx_customers_interests ON customers USING GIN(interests);

-- إدخال البيانات الأولية

-- إدخال الفروع الافتراضية
INSERT INTO branches (code, name_ar, name_en) VALUES
('MAIN', 'الفرع الرئيسي', 'Main Branch'),
('CAIRO', 'فرع القاهرة', 'Cairo Branch'),
('GIZA', 'فرع الجيزة', 'Giza Branch'),
('ALEX', 'فرع الإسكندرية', 'Alexandria Branch'),
('ASWAN', 'فرع أسوان', 'Aswan Branch'),
('ASUUT', 'فرع أسيوط', 'Assiut Branch'),
('BNS', 'فرع بني سويف', 'Beni Suef Branch'),
('POR', 'فرع بورسعيد', 'Port Said Branch'),
('DAM', 'فرع دمياط', 'Damietta Branch'),
('SHG', 'فرع سوهاج', 'Sohag Branch'),
('SUZ', 'فرع السويس', 'Suez Branch'),
('SHM', 'فرع الشرقية', 'Sharqia Branch'),
('GHB', 'فرع الغربية', 'Gharbia Branch'),
('FAY', 'فرع الفيوم', 'Fayoum Branch'),
('QAL', 'فرع القليوبية', 'Qalyubia Branch')
ON CONFLICT (code) DO NOTHING;

-- إدخال الأدوار الافتراضية
INSERT INTO roles (name) VALUES
('مدير النظام'),
('مدير الفرع'),
('موظف'),
('محاسب'),
('مدير مبيعات')
ON CONFLICT (name) DO NOTHING;

-- إدخال أنواع العملاء الافتراضية
INSERT INTO ref_customer_types (name_ar, name_en) VALUES
('فرد', 'Individual'),
('شركة', 'Company'),
('جهة حكومية', 'Government Entity'),
('مكتب هندسي', 'Engineering Office'),
('مقاول', 'Contractor')
ON CONFLICT (name_ar) DO NOTHING;

-- إدخال تصنيفات العملاء الافتراضية
INSERT INTO ref_customer_classifications (name_ar, name_en) VALUES
('عميل عادي', 'Regular Customer'),
('عميل VIP', 'VIP Customer'),
('عميل ذهبي', 'Gold Customer'),
('عميل فضي', 'Silver Customer'),
('عميل برونزي', 'Bronze Customer')
ON CONFLICT (name_ar) DO NOTHING;

-- إدخال المحافظات المصرية
INSERT INTO governorates (code, name_ar, name_en, districts) VALUES
('C', 'القاهرة', 'Cairo', ARRAY['المعادي', 'حلوان', 'مدينة نصر', 'مصر الجديدة', 'شبرا', 'عين شمس', 'التجمع', 'الزمالك', 'جاردن سيتي', 'وسط البلد', 'العباسية', 'الزيتون', 'روض الفرج', 'بولاق', 'الوايلي', 'السيدة زينب', 'مصر القديمة', 'المقطم', 'طرة', '15 مايو']),
('G', 'الجيزة', 'Giza', ARRAY['الدقي', 'العجوزة', 'الهرم', '6 أكتوبر', 'الشيخ زايد', 'إمبابة', 'بولاق الدكرور', 'الوراق', 'أوسيم', 'كرداسة', 'البدرشين', 'الصف', 'أطفيح', 'الواحات البحرية', 'الفيوم', 'بني سويف', 'المنيا', 'أسيوط', 'سوهاج', 'قنا']),
('ALX', 'الإسكندرية', 'Alexandria', ARRAY['سيدي جابر', 'محرم بك', 'العجمي', 'المنتزه', 'المعمورة', 'بحري', 'وسط', 'جليم', 'سموحة', 'سابا باشا', 'ستانلي', 'سيدي بشر', 'ميامي', 'المندرة', 'رأس التين', 'الأنفوشي', 'العطارين', 'كرموز', 'محرم بك', 'المنتزه']),
('ASW', 'أسوان', 'Aswan', ARRAY['أسوان', 'كوم أمبو', 'دراو', 'إدفو', 'نصر النوبة', 'كلابشة', 'أبو سمبل', 'وادي العلاقي', 'الرديسية', 'البصيلية']),
('ASU', 'أسيوط', 'Assiut', ARRAY['أسيوط', 'ديروط', 'منفلوط', 'القوصية', 'أبنوب', 'أبو تيج', 'الغنايم', 'ساحل سليم', 'البداري', 'صدفا', 'الفتح', 'المنشأة', 'المراغة', 'ساحل سليم']),
('BNS', 'بني سويف', 'Beni Suef', ARRAY['بني سويف', 'الواسطي', 'ناصر', 'إهناسيا', 'ببا', 'سمسطا', 'الفيوم', 'طامية', 'سنورس', 'إطسا']),
('POR', 'بورسعيد', 'Port Said', ARRAY['بورسعيد', 'الضواحي', 'العرب', 'المناخ', 'الزهور', 'الشرق', 'الحي الثامن', 'الحي السابع', 'الحي السادس', 'الحي الخامس']),
('DAM', 'دمياط', 'Damietta', ARRAY['دمياط', 'فارسكور', 'الزرقا', 'كفر البطيخ', 'رأس البر', 'الروضة', 'السرو', 'الرهاوي', 'الزرقا', 'كفر سعد']),
('SHG', 'سوهاج', 'Sohag', ARRAY['سوهاج', 'أخميم', 'البلينا', 'المراغة', 'المنشأة', 'دار السلام', 'جرجا', 'طهطا', 'ساقلته', 'طما', 'جهينة', 'الغنايم']),
('SUZ', 'السويس', 'Suez', ARRAY['السويس', 'الأربعين', 'فايد', 'العتاقة', 'الجناين', 'القطاع', 'المنشية', 'الخمسين', 'السبعين', 'الحي العاشر']),
('SHM', 'الشرقية', 'Sharqia', ARRAY['الزقازيق', 'العاشر من رمضان', 'بلبيس', 'أبو حماد', 'ههيا', 'أبو كبير', 'فاقوس', 'القرين', 'كفر صقر', 'الإبراهيمية', 'ديرب نجم', 'كفر شكر', 'أولاد صقر', 'الحسينية', 'منيا القمح', 'مشتول السوق', 'البلينا']),
('GHB', 'الغربية', 'Gharbia', ARRAY['طنطا', 'المحلة الكبرى', 'كفر الزيات', 'زفتى', 'السنطة', 'قطور', 'بسيون', 'سمنود', 'شبين الكوم', 'المنوفية', 'أشمون', 'الباجور', 'قويسنا', 'بركة السبع', 'تلا', 'الشهداء', 'منوف', 'سرس الليان', 'أشمون']),
('FAY', 'الفيوم', 'Fayoum', ARRAY['الفيوم', 'طامية', 'سنورس', 'إطسا', 'يوسف الصديق', 'إبشواي', 'الجامعة', 'الواسطي', 'ناصر', 'إهناسيا']),
('QAL', 'القليوبية', 'Qalyubia', ARRAY['بنها', 'شبين القناطر', 'القناطر الخيرية', 'الخانكة', 'كفر شكر', 'أبو زعبل', 'الخصوص', 'شبرا الخيمة', 'العبور', 'الرحاب', 'مدينة السلام', 'المرج', 'المعادي', 'الزمالك']),
('QEN', 'قنا', 'Qena', ARRAY['قنا', 'قوص', 'نقادة', 'دشنا', 'الوقف', 'قفط', 'أبو تشت', 'فرشوط', 'نجع حمادي', 'دندرة', 'الرزيقات', 'الوقف']),
('KAF', 'كفر الشيخ', 'Kafr El Sheikh', ARRAY['كفر الشيخ', 'دسوق', 'فوه', 'مطوبس', 'البرلس', 'سيدي سالم', 'قلين', 'الحامول', 'الرياض', 'سيدي غازي', 'البرج', 'المراغة']),
('MTN', 'مطروح', 'Matruh', ARRAY['مرسى مطروح', 'الحمام', 'العلمين', 'الضبعة', 'النجيلة', 'سيدي براني', 'السلوم', 'سيوة', 'الوادي الجديد', 'الخارجة', 'الداخلة', 'باريس', 'الفرافرة']),
('MNF', 'المنوفية', 'Menoufia', ARRAY['شبين الكوم', 'أشمون', 'الباجور', 'قويسنا', 'بركة السبع', 'تلا', 'الشهداء', 'منوف', 'سرس الليان', 'أشمون', 'الباجور', 'قويسنا']),
('MIN', 'المنيا', 'Minya', ARRAY['المنيا', 'العدوة', 'مغاغة', 'بني مزار', 'مطاي', 'سمالوط', 'ملوي', 'دير مواس', 'أبو قرقاص', 'مطاي', 'العدوة', 'مغاغة']),
('WAD', 'الوادي الجديد', 'New Valley', ARRAY['الخارجة', 'الداخلة', 'باريس', 'الفرافرة', 'الواحات البحرية', 'الواحات الداخلة', 'الواحات الخارجة', 'الواحات الفرافرة']),
('NSR', 'شمال سيناء', 'North Sinai', ARRAY['العريش', 'الشيخ زايد', 'رفح', 'بئر العبد', 'الحسنة', 'نخل', 'الطور', 'سانت كاترين', 'أبو رديس', 'أبو زنيمة', 'رأس سدر', 'شرم الشيخ', 'دهب', 'نويبع', 'طابا']),
('JNS', 'جنوب سيناء', 'South Sinai', ARRAY['الطور', 'سانت كاترين', 'أبو رديس', 'أبو زنيمة', 'رأس سدر', 'شرم الشيخ', 'دهب', 'نويبع', 'طابا', 'نخل', 'الحسنة']),
('ISL', 'الإسماعيلية', 'Ismailia', ARRAY['الإسماعيلية', 'فايد', 'القنطرة شرق', 'القنطرة غرب', 'التل الكبير', 'أبو صوير', 'القصاصين', 'التل الكبير', 'فايد', 'القنطرة']),
('LUX', 'الأقصر', 'Luxor', ARRAY['الأقصر', 'إسنا', 'أرمنت', 'الطود', 'بياضة العرب', 'الزينية', 'البياضية', 'القرنة', 'الكرنك', 'الطود', 'إسنا', 'أرمنت']),
('RED', 'البحر الأحمر', 'Red Sea', ARRAY['الغردقة', 'رأس غارب', 'سفاجا', 'القصير', 'مرسى علم', 'شلاتين', 'حلايب', 'أبو رماد', 'الدهار', 'الشلاتين']),
('BEH', 'البحيرة', 'Beheira', ARRAY['دمنهور', 'رشيد', 'إدكو', 'أبو المطامير', 'أبو حمص', 'الدلنجات', 'المحمودية', 'الرحمانية', 'إيتاي البارود', 'حوش عيسى', 'شبراخيت', 'كوم حمادة', 'بدر', 'وادي النطرون', 'النوبارية', 'البيوم', 'السلوم']),
('DAK', 'الدقهلية', 'Dakahlia', ARRAY['المنصورة', 'طلخا', 'ميت غمر', 'دكرنس', 'أجا', 'منية النصر', 'السنبلاوين', 'المنزلة', 'الجمالية', 'شربين', 'المطرية', 'بلقاس', 'ميت سلسيل', 'جمصة', 'محلة دمنة', 'نبروه', 'السنبلاوين', 'المنزلة'])
ON CONFLICT (code) DO NOTHING;

-- إنشاء مستخدم افتراضي (كلمة المرور: admin123)
INSERT INTO users (username, password_hash, branch_id) VALUES
('admin', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 1)
ON CONFLICT (username) DO NOTHING;

-- تعيين دور مدير النظام للمستخدم الافتراضي
INSERT INTO user_roles (user_id, role_id) 
SELECT u.id, r.id 
FROM users u, roles r 
WHERE u.username = 'admin' AND r.name = 'مدير النظام'
ON CONFLICT DO NOTHING; 