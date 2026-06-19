# Distributed Authentication Architecture (Cross-App Auth Flow)

ระบบสถาปัตยกรรมการลงทะเบียนและเข้าสู่ระบบแบบกระจายส่วน (Distributed Authentication) ที่ออกแบบมาเพื่อแยกความรับผิดชอบ (Separation of Concerns) ทั้งในส่วนหน้าบ้าน (Frontend) และหลังบ้าน (Backend) โดยใช้เทคโนโลยีที่แตกต่างกันทำงานร่วมกันผ่านฐานข้อมูลเดียว (Shared Database)

---

## 🏗️ System Architecture & Port Mapping

การไหลของข้อมูลในระบบถูกออกแบบเป็น **Cross-App Redirect Pattern** เพื่อให้แต่ละโปรเจกต์ทำหน้าที่เฉพาะด้านของตัวเองอย่างเป็นอิสระ (Single Responsibility Principle)

```
                       ┌────────────────────────┐
                       │   Vue 3 (Register)     │ ◄──┐
                       │      Port: 5173        │    │
                       └───────────┬────────────┘    │
                                   │                 │
                                   │ POST /register  │ Redirect to
                                   ▼                 │ Register
                       ┌────────────────────────┐    │
                       │ C# ASP.NET Core 8 API  │    │
                       │      Port: 5001        │    │
                       └───────────┬────────────┘    │
                                   │                 │
                                   │ Writes hash     │
                                   ▼                 │
┌──────────────┐       ┌────────────────────────┐    │
│ SQLite DB    │ ◄───► │      shared/app.db     │    │
└──────────────┘       └────────────────────────┘    │
                                   ▲                 │
                                   │ Reads hash      │
                                   │                 │
                       ┌────────────────────────┐    │
                       │   Go (Gin) Auth API    │    │
                       │      Port: 5002        │    │
                       └───────────▲────────────┘    │
                                   │                 │
                                   │ POST /login     │
                                   │ (Issues JWT)    │
                                   │                 │
                       ┌───────────┴────────────┐    │
                       │  Angular 17 (Login)    │ ───┘
                       │      Port: 4200        │
                       └────────────────────────┘
```

| Service / App | Role | Technology Stack | Port |
| :--- | :--- | :--- | :--- |
| **frontend-vue** | หน้าจอสมัครสมาชิก (Register UI) | Vue 3, Vite, Pinia, Axios | `5173` |
| **backend-csharp** | จัดการบันทึกสมาชิกใหม่ (Register API) | ASP.NET Core 8, EF Core (SQLite) | `5001` |
| **frontend-angular** | หน้าจอเข้าสู่ระบบ & ยินดีต้อนรับ (Login & Welcome UI) | Angular 17 (Standalone Components) | `4200` |
| **backend-golang** | ตรวจสอบรหัสผ่าน & ออก JWT (Login API) | Go (Gin), golang-jwt/v5, go-sqlite3 | `5002` |
| **shared** | ฐานข้อมูลร่วม (Shared DB File) | SQLite (`app.db`) | *N/A* |

---

## 🛠️ Technical Design Decisions

### 1. Shared Database State (SQLite)
* เลือกใช้ SQLite เป็นตัวกลางเก็บข้อมูลบัญชีผู้ใช้ในรูปของไฟล์ [app.db](file:///Users/ho/interview-question-087/shared/app.db) ซึ่งทั้ง **C# Backend** (Entity Framework Core) และ **Go Backend** (Database/SQL Driver) เชื่อมโยงเข้าหาจุดเดียวกัน
* มีกลไกสร้างตารางอัตโนมัติ (Auto Schema Migration/Ensure Created) เมื่อเซิร์ฟเวอร์ใดเซิร์ฟเวอร์หนึ่งเริ่มทำงาน

### 2. Cryptographic Security (BCrypt)
* รหัสผ่านของผู้ใช้จะไม่ถูกบันทึกเป็น Plain Text
* **C# Backend** ทำหน้าที่ Hash รหัสผ่านด้วยอัลกอริทึม **BCrypt** ก่อนนำเข้าสู่ Database
* **Go Backend** ตรวจสอบรหัสผ่านโดยเทียบค่าความต่างของ Hash ด้วย BCrypt เพื่อความปลอดภัยระดับมาตรฐานอุตสาหกรรม

### 3. Stateless Authentication (JWT)
* เมื่อผ่านการตรวจสอบสิทธิ์ **Go Backend** จะสร้าง Token แบบสมมาตร (Symmetric Sign) โดยใช้ **HS256** และส่งมอบสิทธิ์ในรูปแบบ **JWT**
* **Angular Frontend** จะเก็บ JWT ไว้ใน `localStorage` และมีระบบ **Auth Guard (Route Protection)** ตรวจสอบความถูกต้องของ Token ก่อนให้ผู้ใช้เข้าถึงหน้า Dashboard/Welcome

---

## 🚀 Getting Started

เปิดโปรเจกต์ใน Visual Studio Code หรือ Terminal แล้วเปิดแท็บรันบริการต่าง ๆ ตามคำแนะนำด้านล่าง:

### 1. รันระบบฐานข้อมูลร่วมและสมัครสมาชิก (C# & Vue)
```bash
# 1.1 รัน Backend (C#)
cd backend-csharp
dotnet run

# 1.2 รัน Frontend (Vue)
cd frontend-vue
npm install
npm run dev
```

### 2. รันระบบเข้าสู่ระบบและตรวจสอบสิทธิ์ (Go & Angular)
```bash
# 2.1 รัน Backend (Go)
cd backend-golang
go run .

# 2.2 รัน Frontend (Angular)
cd frontend-angular
npm install
npx ng serve --port 4200
```

---

## 📊 Database Schema

ตารางถูกบันทึกไว้ใน [schema.sql](file:///Users/ho/interview-question-087/shared/schema.sql) โดยมีโครงสร้างดังนี้:

```sql
CREATE TABLE IF NOT EXISTS users (
  id         INTEGER  PRIMARY KEY AUTOINCREMENT,
  username   TEXT     NOT NULL UNIQUE,
  password   TEXT     NOT NULL,   -- เก็บรหัสผ่านในรูปแบบ BCrypt hash
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

---

## 🔍 How to Verify Password Encryption (การตรวจสอบการเข้ารหัสข้อมูล)

คุณสามารถเปิดดูข้อมูลในตารางเพื่อตรวจสอบได้ว่า Password ถูกเข้ารหัสด้วย BCrypt (ขึ้นต้นด้วย `$2a$...`) จริงตามข้อกำหนด โดยรันคำสั่งนี้ที่ Root Folder:

```bash
sqlite3 shared/app.db "SELECT id, username, password FROM users;"
```
