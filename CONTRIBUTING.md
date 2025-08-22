# دليل المساهمة - نظام ELKHAWAGA ERP

شكراً لاهتمامك بالمساهمة في تطوير نظام ELKHAWAGA ERP! هذا الدليل سيساعدك على البدء في المساهمة.

## 📋 جدول المحتويات

- [كيفية المساهمة](#كيفية-المساهمة)
- [إعداد بيئة التطوير](#إعداد-بيئة-التطوير)
- [معايير الكود](#معايير-الكود)
- [عملية التطوير](#عملية-التطوير)
- [الاختبارات](#الاختبارات)
- [التوثيق](#التوثيق)
- [الإبلاغ عن الأخطاء](#الإبلاغ-عن-الأخطاء)
- [طلب الميزات](#طلب-الميزات)

## 🤝 كيفية المساهمة

### أنواع المساهمات المطلوبة

- 🐛 **إصلاح الأخطاء** - تحديد وإصلاح المشاكل
- ✨ **ميزات جديدة** - إضافة وظائف جديدة
- 📚 **تحسين التوثيق** - تحديث وتحسين الوثائق
- 🎨 **تحسينات الواجهة** - تحسين تجربة المستخدم
- ⚡ **تحسينات الأداء** - تحسين سرعة وكفاءة النظام
- 🔒 **تحسينات الأمان** - تعزيز أمان النظام
- 🧪 **اختبارات** - إضافة اختبارات جديدة

### خطوات المساهمة

1. **Fork المشروع**
   ```bash
   git clone https://github.com/your-username/ERP_ELKHAWAGA.git
   cd ERP_ELKHAWAGA
   ```

2. **إنشاء فرع جديد**
   ```bash
   git checkout -b feature/your-feature-name
   # أو
   git checkout -b fix/your-bug-fix
   ```

3. **إجراء التغييرات**
   - اتبع معايير الكود
   - أضف اختبارات للميزات الجديدة
   - حدث التوثيق عند الحاجة

4. **اختبار التغييرات**
   ```bash
   make test
   ```

5. **Commit التغييرات**
   ```bash
   git add .
   git commit -m "feat: add new customer search feature"
   ```

6. **Push الفرع**
   ```bash
   git push origin feature/your-feature-name
   ```

7. **إنشاء Pull Request**
   - املأ قالب Pull Request
   - وصف التغييرات بوضوح
   - أضف لقطات شاشة إذا لزم الأمر

## 🛠️ إعداد بيئة التطوير

### المتطلبات الأساسية

- Go 1.21+
- Node.js 18+
- PostgreSQL 15+
- Redis 7+
- Git

### إعداد المشروع

```bash
# استنساخ المشروع
git clone https://github.com/your-username/ERP_ELKHAWAGA.git
cd ERP_ELKHAWAGA

# إعداد قاعدة البيانات
make create-db
make migrate

# تثبيت التبعيات
make install

# إعداد البيئة
make env-backend

# تشغيل المشروع
make dev
```

### أدوات التطوير الموصى بها

- **IDE**: VS Code, GoLand, WebStorm
- **Database**: pgAdmin, DBeaver
- **API Testing**: Postman, Insomnia
- **Git**: GitKraken, SourceTree

## 📝 معايير الكود

### Go (Backend)

#### التنسيق
```go
// استخدام gofmt
go fmt ./...

// استخدام golint
golint ./...

// استخدام go vet
go vet ./...
```

#### معايير التسمية
```go
// المتغيرات والوظائف
var customerName string
func GetCustomerByID(id int) (*Customer, error)

// الثوابت
const MaxRetries = 3

// الأنواع
type CustomerService struct {
    repo CustomerRepository
}

// الواجهات
type CustomerRepository interface {
    FindByID(id int) (*Customer, error)
}
```

#### التعليقات
```go
// CustomerService provides business logic for customer operations
type CustomerService struct {
    repo CustomerRepository
}

// GetCustomerByID retrieves a customer by their ID
func (s *CustomerService) GetCustomerByID(id int) (*Customer, error) {
    // Implementation
}
```

### TypeScript/Angular (Frontend)

#### التنسيق
```bash
# استخدام Prettier
npx prettier --write src/

# استخدام ESLint
npx eslint src/
```

#### معايير التسمية
```typescript
// المتغيرات والوظائف
const customerName: string = 'John Doe';
function getCustomerById(id: number): Observable<Customer> {
    // Implementation
}

// الأنواع والواجهات
interface Customer {
    id: number;
    name: string;
    email: string;
}

// الخدمات
@Injectable({
    providedIn: 'root'
})
export class CustomerService {
    // Implementation
}
```

#### التعليقات
```typescript
/**
 * Service for managing customer operations
 */
@Injectable({
    providedIn: 'root'
})
export class CustomerService {
    /**
     * Retrieves a customer by their ID
     * @param id Customer ID
     * @returns Observable of customer data
     */
    getCustomerById(id: number): Observable<Customer> {
        // Implementation
    }
}
```

## 🔄 عملية التطوير

### دورة التطوير

1. **التخطيط**
   - فهم المتطلبات
   - تصميم الحل
   - تحديد الاختبارات

2. **التطوير**
   - كتابة الكود
   - اتباع معايير الكود
   - كتابة الاختبارات

3. **الاختبار**
   - تشغيل الاختبارات المحلية
   - اختبار التكامل
   - اختبار الأداء

4. **المراجعة**
   - مراجعة الكود الذاتية
   - طلب مراجعة من الفريق
   - معالجة التعليقات

5. **النشر**
   - دمج التغييرات
   - نشر في بيئة الاختبار
   - نشر في الإنتاج

### Git Workflow

```bash
# تحديث الفرع الرئيسي
git checkout main
git pull origin main

# إنشاء فرع جديد
git checkout -b feature/new-feature

# تطوير الميزة
# ... كتابة الكود ...

# إضافة التغييرات
git add .

# Commit مع رسالة واضحة
git commit -m "feat: add customer search functionality

- Add search by name and phone
- Add pagination support
- Add unit tests
- Update documentation"

# Push الفرع
git push origin feature/new-feature

# إنشاء Pull Request
```

### رسائل Commit

نتبع [Conventional Commits](https://www.conventionalcommits.org/):

```bash
# أنواع الرسائل
feat: add new feature
fix: fix a bug
docs: update documentation
style: format code
refactor: refactor code
test: add tests
chore: maintenance tasks

# أمثلة
feat: add customer search functionality
fix: resolve authentication token issue
docs: update API documentation
style: format customer service code
refactor: improve error handling
test: add customer service tests
chore: update dependencies
```

## 🧪 الاختبارات

### Backend Tests (Go)

```bash
# تشغيل جميع الاختبارات
go test ./...

# تشغيل اختبارات موديول معين
go test ./internal/modules/customers/...

# تشغيل اختبارات مع تغطية
go test -cover ./...

# تشغيل اختبارات الأداء
go test -bench=. ./...
```

#### كتابة الاختبارات
```go
func TestCustomerService_CreateCustomer(t *testing.T) {
    // Arrange
    service := NewCustomerService(mockRepo)
    customer := &Customer{
        Name: "John Doe",
        Phone: "0123456789",
    }

    // Act
    result, err := service.CreateCustomer(customer)

    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, customer.Name, result.Name)
}
```

### Frontend Tests (Angular)

```bash
# تشغيل اختبارات الوحدة
npm test

# تشغيل اختبارات E2E
npm run e2e

# تشغيل اختبارات مع تغطية
npm run test:coverage
```

#### كتابة الاختبارات
```typescript
describe('CustomerService', () => {
    let service: CustomerService;
    let httpMock: HttpTestingController;

    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [HttpClientTestingModule],
            providers: [CustomerService]
        });
        service = TestBed.inject(CustomerService);
        httpMock = TestBed.inject(HttpTestingController);
    });

    it('should create customer', () => {
        const customer = { name: 'John Doe', phone: '0123456789' };
        
        service.createCustomer(customer).subscribe(result => {
            expect(result.name).toBe(customer.name);
        });

        const req = httpMock.expectOne('/api/customers');
        expect(req.request.method).toBe('POST');
        req.flush(customer);
    });
});
```

## 📚 التوثيق

### تحديث التوثيق

عند إضافة ميزات جديدة، تأكد من تحديث:

1. **README.md** - إذا كانت الميزة مهمة للمستخدمين
2. **API Documentation** - إذا أضفت endpoints جديدة
3. **Inline Comments** - تعليقات في الكود
4. **CHANGELOG.md** - تسجيل التغييرات

### معايير التوثيق

- استخدم اللغة العربية للتوثيق الرئيسي
- استخدم اللغة الإنجليزية للتعليقات في الكود
- اكتب أمثلة واضحة
- أضف لقطات شاشة عند الحاجة

## 🐛 الإبلاغ عن الأخطاء

### قالب الإبلاغ عن الأخطاء

```markdown
## وصف المشكلة
وصف واضح ومختصر للمشكلة.

## خطوات إعادة الإنتاج
1. اذهب إلى '...'
2. انقر على '...'
3. انتقل إلى '...'
4. انظر إلى الخطأ

## السلوك المتوقع
وصف لما يجب أن يحدث.

## لقطات الشاشة
إذا كان ذلك مناسباً، أضف لقطات شاشة.

## معلومات النظام
- نظام التشغيل: [مثل Windows 10]
- المتصفح: [مثل Chrome 90]
- إصدار التطبيق: [مثل 1.0.0]

## معلومات إضافية
أي معلومات أخرى حول المشكلة.
```

## ✨ طلب الميزات

### قالب طلب الميزة

```markdown
## ملخص الميزة
وصف واضح ومختصر للميزة المطلوبة.

## المشكلة التي تحلها
وصف للمشكلة التي ستحلها هذه الميزة.

## الحل المقترح
وصف للحل المقترح.

## البدائل المدروسة
وصف للبدائل التي تم النظر فيها.

## معلومات إضافية
أي معلومات أخرى مفيدة.
```

## 🏷️ التصنيفات

### Labels للمسائل

- `bug` - خطأ في النظام
- `enhancement` - تحسين ميزة موجودة
- `feature` - ميزة جديدة
- `documentation` - تحسين التوثيق
- `good first issue` - مناسب للمبتدئين
- `help wanted` - يحتاج مساعدة
- `priority: high` - أولوية عالية
- `priority: medium` - أولوية متوسطة
- `priority: low` - أولوية منخفضة

### Labels للـ Pull Requests

- `ready for review` - جاهز للمراجعة
- `work in progress` - قيد العمل
- `needs testing` - يحتاج اختبار
- `breaking change` - تغيير غير متوافق
- `hotfix` - إصلاح عاجل

## 📞 التواصل

### قنوات التواصل

- **GitHub Issues** - للإبلاغ عن الأخطاء وطلب الميزات
- **GitHub Discussions** - للمناقشات العامة
- **Pull Requests** - للمراجعة والتعليقات

### إرشادات التواصل

- كن محترماً ومهذباً
- استخدم اللغة العربية أو الإنجليزية
- قدم معلومات واضحة ومفصلة
- استجب للتعليقات بسرعة

## 🏆 الاعتراف

سنعترف بمساهماتك في:

- ملف CONTRIBUTORS.md
- إصدارات GitHub
- التوثيق
- الاجتماعات والمناسبات

---

**شكراً لمساهمتك في تطوير نظام ELKHAWAGA ERP!** 🚀 