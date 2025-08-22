# Makefile لنظام ELKHAWAGA ERP

.PHONY: help install build run test clean docker-build docker-run setup-db

# المتغيرات
APP_NAME=elkhawaga-erp
BACKEND_DIR=backend
FRONTEND_DIR=frontend
DOCKER_COMPOSE_FILE=infra/docker/docker-compose.yml

# المساعدة
help:
	@echo "أوامر متاحة:"
	@echo "  install      - تثبيت جميع التبعيات"
	@echo "  build        - بناء المشروع"
	@echo "  run          - تشغيل المشروع محلياً"
	@echo "  test         - تشغيل الاختبارات"
	@echo "  clean        - تنظيف الملفات المؤقتة"
	@echo "  docker-build - بناء صور Docker"
	@echo "  docker-run   - تشغيل المشروع بـ Docker"
	@echo "  setup-db     - إعداد قاعدة البيانات"
	@echo "  dev          - تشغيل وضع التطوير"

# تثبيت التبعيات
install:
	@echo "تثبيت تبعيات الباك إند..."
	cd $(BACKEND_DIR) && go mod download
	@echo "تثبيت تبعيات الفرونت إند..."
	cd $(FRONTEND_DIR) && npm install
	@echo "تم تثبيت جميع التبعيات بنجاح!"

# بناء المشروع
build:
	@echo "بناء الباك إند..."
	cd $(BACKEND_DIR) && go build -o bin/server cmd/server/main.go
	@echo "بناء الفرونت إند..."
	cd $(FRONTEND_DIR) && npm run build
	@echo "تم بناء المشروع بنجاح!"

# تشغيل المشروع محلياً
run:
	@echo "تشغيل الباك إند..."
	cd $(BACKEND_DIR) && go run cmd/server/main.go &
	@echo "تشغيل الفرونت إند..."
	cd $(FRONTEND_DIR) && npm start &
	@echo "المشروع يعمل على:"
	@echo "  - الفرونت إند: http://localhost:4200"
	@echo "  - الباك إند: http://localhost:8080"

# تشغيل وضع التطوير
dev:
	@echo "تشغيل وضع التطوير..."
	@echo "الباك إند يعمل على http://localhost:8080"
	@echo "الفرونت إند يعمل على http://localhost:4200"
	@echo "اضغط Ctrl+C لإيقاف التشغيل"
	@trap 'kill %1 %2' SIGINT; \
	cd $(BACKEND_DIR) && go run cmd/server/main.go & \
	cd $(FRONTEND_DIR) && npm start & \
	wait

# تشغيل الاختبارات
test:
	@echo "تشغيل اختبارات الباك إند..."
	cd $(BACKEND_DIR) && go test ./...
	@echo "تشغيل اختبارات الفرونت إند..."
	cd $(FRONTEND_DIR) && npm test

# تنظيف الملفات المؤقتة
clean:
	@echo "تنظيف ملفات الباك إند..."
	cd $(BACKEND_DIR) && go clean -cache -modcache
	rm -rf $(BACKEND_DIR)/bin
	@echo "تنظيف ملفات الفرونت إند..."
	cd $(FRONTEND_DIR) && npm run clean
	rm -rf $(FRONTEND_DIR)/dist
	@echo "تم التنظيف بنجاح!"

# بناء صور Docker
docker-build:
	@echo "بناء صور Docker..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) build
	@echo "تم بناء الصور بنجاح!"

# تشغيل المشروع بـ Docker
docker-run:
	@echo "تشغيل المشروع بـ Docker..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d
	@echo "المشروع يعمل على:"
	@echo "  - الفرونت إند: http://localhost:80"
	@echo "  - الباك إند: http://localhost:8080"

# إيقاف Docker
docker-stop:
	@echo "إيقاف المشروع..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down
	@echo "تم إيقاف المشروع!"

# إعداد قاعدة البيانات
setup-db:
	@echo "إعداد قاعدة البيانات..."
	@echo "يرجى التأكد من تشغيل PostgreSQL و Redis"
	@echo "ثم قم بتشغيل:"
	@echo "  make install"
	@echo "  cd backend && go run cmd/server/main.go"

# فحص حالة الخدمات
status:
	@echo "فحص حالة الخدمات..."
	@echo "PostgreSQL:"
	@systemctl is-active postgresql || echo "  غير نشط"
	@echo "Redis:"
	@systemctl is-active redis || echo "  غير نشط"
	@echo "Docker:"
	@docker --version || echo "  غير مثبت"

# إنشاء ملف .env للباك إند
env-backend:
	@echo "إنشاء ملف .env للباك إند..."
	@cat > $(BACKEND_DIR)/.env << EOF
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
	@echo "تم إنشاء ملف .env!"

# إنشاء قاعدة البيانات
create-db:
	@echo "إنشاء قاعدة البيانات..."
	sudo -u postgres psql -c "CREATE DATABASE zakee_erp;" || echo "قاعدة البيانات موجودة بالفعل"
	sudo -u postgres psql -c "CREATE USER erp_user WITH PASSWORD 'password';" || echo "المستخدم موجود بالفعل"
	sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE zakee_erp TO erp_user;"
	@echo "تم إنشاء قاعدة البيانات!"

# تشغيل ملف الهجرة
migrate:
	@echo "تشغيل ملف الهجرة..."
	psql -U erp_user -d zakee_erp -f infra/migrations/0001_init.sql
	@echo "تم تشغيل ملف الهجرة!"

# إعداد كامل
setup: create-db migrate env-backend install
	@echo "تم الإعداد الكامل للمشروع!"
	@echo "يمكنك الآن تشغيل: make dev" 