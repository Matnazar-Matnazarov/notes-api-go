Zo‘r, endi **to‘liq konseptual tushuntirib beraman**, **kodsiz**, oddiy tilda. Bu joyni tushunsang — keyin kod o‘zi yozilib ketadi.

---

## LOYIHA NIMA QILADI? (Oddiy til bilan)

Bu loyiha — **Notes CRUD backend**.

Foydalanuvchi:

* eslatma (note) yozadi
* ko‘radi
* o‘zgartiradi
* o‘chiradi

Xullas, **oddiy daftar**, lekin web orqali.

---

## BACKEND’NING ASOSIY VAZIFASI

Backend — bu **miyaning o‘zi**.

U:

* frontenddan kelgan so‘rovlarni qabul qiladi
* ma’lumotni saqlaydi
* tekshiradi
* natijani qaytaradi

Frontend faqat:

* tugma
* forma
* ko‘rinish

---

## BACKEND ICHIDA NIMALAR BO‘LADI?

### 1️⃣ HTTP so‘rovlarni qabul qilish

Frontend:

* “mana bu note ni saqla”
* “hamma note’larni ber”
* “mana buni o‘chir”

deydi.

Backend:

* so‘rovni qabul qiladi
* tushunadi
* nima qilishini hal qiladi

---

### 2️⃣ Ma’lumotni tekshirish (validation)

Backend:

* title bo‘sh emasmi?
* content juda uzun emasmi?
* id mavjudmi?

Noto‘g‘ri bo‘lsa:

* xato qaytaradi
* frontend o‘sha xatoni ko‘rsatadi

---

### 3️⃣ Database bilan ishlash

Backend:

* yangi note ni DB ga yozadi
* DB dan note’larni olib keladi
* o‘zgartiradi
* o‘chiradi

Frontend **hech qachon** DB bilan gaplashmaydi.

---

### 4️⃣ Business logic

Bu yerda qoidalar turadi:

* qaysi ma’lumot saqlanadi
* qaysi holatda xato beriladi
* qanday javob qaytariladi

Hozircha qoidalar **juda oddiy**.

---

### 5️⃣ Response qaytarish

Backend:

* muvaffaqiyat bo‘lsa → ma’lumot
* xato bo‘lsa → xabar
* status code bilan

Masalan:

* 200 → hammasi joyida
* 400 → noto‘g‘ri so‘rov
* 404 → topilmadi

---

## FRONTEND BILAN ALOQA QANDAY?

Frontend:

1. forma to‘ldiradi
2. API ga so‘rov yuboradi
3. javobni oladi
4. ekranga chiqaradi

Frontend **hech qachon**:

* DB ni bilmaydi
* backend ichini ko‘rmaydi

---

## BU LOYIHA SENI NIMAGA O‘RGATADI?

Agar shu backendni qilsang, sen:

* REST API nima ekanini tushunasan
* CRUD qanday ishlashini bilasan
* Frontend bilan real aloqa qilasan
* Backend nima ish qilishini aniq tushunasan

---

## OCHIQ GAP

Bu loyiha:

* kichik
* sodda
* lekin **asosiy tushunchalarni to‘liq beradi**

---

## 📁 LOYIHA STRUKTURASI

Professional Go loyiha strukturasiga ega:

```
notes-api-go/
├── cmd/
│   └── api/
│       └── main.go              # Dastur kirish nuqtasi
├── internal/
│   ├── config/
│   │   └── config.go            # Konfiguratsiya boshqaruvi
│   ├── database/
│   │   └── database.go          # Database ulanishi va migratsiyalar
│   ├── handlers/
│   │   ├── note_handler.go      # HTTP handlerlar (CRUD operatsiyalar)
│   │   └── router.go            # API route'lar
│   ├── models/
│   │   └── note.go              # Database modellar
│   └── services/
│       └── note_service.go      # Business logic
├── .env.example                  # Environment variable'lar misoli
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── Makefile                      # Qulay buyruqlar
└── README.md
```

### Struktura tushuntirish:

- **cmd/api/** - Dastur ishga tushiriladigan joy
- **internal/** - Ichki kod (tashqi paketlar uchun yopiq)
  - **config/** - Sozlamalar (database, server port, va hokazo)
  - **database/** - Database ulanishi va migratsiyalar
  - **handlers/** - HTTP so'rovlarni qabul qiluvchi kodlar
  - **models/** - Database jadval strukturalari
  - **services/** - Business logic (qoidalar va validatsiya)

---

## 🚀 QANDAY ISHGA TUSHIRISH?

### 1. Talablar

- Go 1.25+ o'rnatilgan bo'lishi kerak
- PostgreSQL database ishga tushirilgan bo'lishi kerak

### 2. Database tayyorlash

PostgreSQL'da yangi database yarating:

```sql
CREATE DATABASE notes_db;
```

### 3. Environment variable'lar

`.env.example` faylini `.env` ga ko'chiring va o'z qiymatlarizni kiriting:

```bash
cp .env.example .env
```

`.env` faylini tahrirlang:

```env
SERVER_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sizning_parolingiz
DB_NAME=notes_db
DB_SSLMODE=disable
```

### 4. Dasturni ishga tushirish

```bash
# Makefile orqali
make run

# Yoki to'g'ridan-to'g'ri
go run ./cmd/api
```

Server `http://localhost:8080` da ishga tushadi.

---

## 📡 API ENDPOINT'LAR

### Health Check
```
GET /health
```

### Notes CRUD

1. **Yangi note yaratish**
   ```
   POST /api/v1/notes
   Content-Type: application/json
   
   {
     "title": "Eslatma sarlavhasi",
     "content": "Eslatma matni"
   }
   ```

2. **Barcha note'larni olish**
   ```
   GET /api/v1/notes
   ```

3. **Bitta note olish (ID bo'yicha)**
   ```
   GET /api/v1/notes/:id
   ```

4. **Note yangilash**
   ```
   PUT /api/v1/notes/:id
   Content-Type: application/json
   
   {
     "title": "Yangi sarlavha",
     "content": "Yangi matn"
   }
   ```

5. **Note o'chirish**
   ```
   DELETE /api/v1/notes/:id
   ```

---

## 🛠️ FOYDALI BUYRUQLAR

```bash
# Dasturni ishga tushirish
make run

# Dasturni build qilish
make build

# Testlarni ishga tushirish
make test

# Build fayllarni tozalash
make clean

# Swagger dokumentatsiyasini yaratish
make swagger
```

---

## 📚 SWAGGER DOKUMENTATSIYASI

Loyiha **avtomatik Swagger dokumentatsiyasi** bilan jihozlangan!

### Swagger UI'ni ko'rish

1. Dasturni ishga tushiring:
   ```bash
   make run
   ```

2. Brauzerda oching:
   ```
   http://localhost:8080/swagger/index.html
   ```

Swagger UI'da siz:
- ✅ Barcha API endpoint'larni ko'rishingiz mumkin
- ✅ Har bir endpoint'ning parametrlarini ko'rishingiz mumkin
- ✅ To'g'ridan-to'g'ri API'ni test qilishingiz mumkin
- ✅ Request/Response misollarini ko'rishingiz mumkin

### Swagger dokumentatsiyasini yangilash

Kod o'zgarganda, Swagger dokumentatsiyasini yangilash:

```bash
make swagger
```

Bu buyruq:
- Kod ichidagi annotatsiyalarni o'qiydi
- `docs/` papkasida yangi dokumentatsiya yaratadi
- Avtomatik ravishda barcha endpoint'larni yangilaydi
- Agar `swag` tool o'rnatilmagan bo'lsa, avtomatik o'rnatadi

### Swagger annotatsiyalari

Har bir handler funksiyasida Swagger annotatsiyalari mavjud:

```go
// @Summary      Create a new note
// @Description  Create a new note with title and content
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        note  body      models.CreateNoteRequest  true  "Note data"
// @Success      201   {object}  map[string]interface{}     "Note created successfully"
// @Router       /notes [post]
```

Bu annotatsiyalar Swagger UI'da avtomatik ko'rinadi.

---

## 💡 KOD TUSHUNARLILIGI

Har bir fayl **to'liq comment'lar** bilan yozilgan:

- **Qanday ishlaydi?** - Har bir funksiya va struct'ning vazifasi
- **Nima uchun?** - Nima uchun shunday yozilgan
- **Qanday ishlatiladi?** - Misollar va foydalanish

Kod **professional standartlar** bo'yicha yozilgan:
- ✅ Clean Architecture (handlers → services → database)
- ✅ Separation of Concerns (har bir qatlam o'z vazifasini bajaradi)
- ✅ Error handling (barcha xatolar to'g'ri qayta ishlanadi)
- ✅ Validation (ma'lumotlar tekshiriladi)
- ✅ Soft delete (note'lar to'liq o'chirilmaydi)

---

## 📚 QO'SHIMCHA MA'LUMOT

Bu loyiha quyidagi texnologiyalardan foydalanadi:

- **Gin** - HTTP web framework
- **GORM** - ORM (Object-Relational Mapping)
- **PostgreSQL** - Database
- **godotenv** - Environment variable'lar boshqaruvi
- **Swagger/OpenAPI** - API dokumentatsiyasi (swaggo/swag)

---

## 🎯 KEYINGI QADAMLAR

Loyihani kengaytirish uchun:

1. Authentication/Authorization qo'shish
2. Pagination qo'shish (note'lar ro'yxatida)
3. Search va filter qo'shish
4. Unit testlar yozish
5. Docker containerization
6. CI/CD pipeline


