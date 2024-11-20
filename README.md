# ✨ ระบบจัดการงาน (Task Management System)

ระบบ **Task Management System** ถูกพัฒนาขึ้นโดยใช้ **Vue.js** สำหรับส่วนหน้า (Frontend), **Golang** สำหรับส่วนหลัง (Backend) และ **PostgreSQL** สำหรับฐานข้อมูล ระบบรองรับการสร้าง แก้ไข และลบงาน รวมถึงการติดตามสถานะของงาน พร้อมระบบการยืนยันตัวตนด้วย **JWT Authentication**

---

## **คุณสมบัติ (Features)**

### **Frontend** 🖥️
- พัฒนาด้วย **Vue.js** และ **Tailwind CSS**
- Task Board ที่สามารถ:
  - ➕ เพิ่ม, ✏️ แก้ไข, ❌ ลบงาน
  - 📋 อัปเดตรายละเอียดงาน เช่น ชื่อ, คำอธิบาย, ความสำคัญ, วันที่ครบกำหนด
  - 🔄 จัดการสถานะของงานระหว่าง `To Do`, `In Progress`, และ `Done`
- 📊 Dashboard พร้อมกราฟวงกลม (Pie Chart) แสดงภาพรวมของสถานะงาน
- 🔐 ระบบยืนยันตัวตนด้วย **JWT Authentication**

### **Backend** ⚙️
- พัฒนาด้วย **Golang** และ **Gin Framework**
- RESTful API:
  - 🔑 `POST /login`: ยืนยันตัวตนและสร้าง JWT Token
  - 📥 `GET /tasks`: ดึงข้อมูลงานทั้งหมด
  - ➕ `POST /tasks`: เพิ่มงานใหม่
  - ✏️ `PUT /tasks/:id`: อัปเดตงาน
  - ❌ `DELETE /tasks/:id`: ลบงาน
- การเชื่อมต่อกับฐานข้อมูล **PostgreSQL**

### **ฐานข้อมูล** 🗄️
- ใช้ **PostgreSQL** และตาราง `tasks` สำหรับจัดการข้อมูลงาน
- ข้อมูลเริ่มต้นถูกกำหนดในไฟล์ `init.sql`

### **Authentication** 🔒
- ระบบยืนยันตัวตนที่ปลอดภัยด้วย **JWT (JSON Web Token)**
- มี Middleware สำหรับปกป้อง API ที่ต้องการสิทธิ์การเข้าถึง

---

## **การติดตั้ง (Setup)**

### **สิ่งที่ต้องมี (Prerequisites)** 🛠️
- ติดตั้ง **Docker** และ **Docker Compose**
- ติดตั้ง **Node.js** และ **npm** (สำหรับพัฒนา Frontend)

---

### **ขั้นตอนการติดตั้ง**

#### 1. Clone โค้ดจาก Repository:
```bash
git clone https://github.com/your-repo/task-management-system.git
cd task-management-system

#### 2.เริ่มต้นระบบด้วย Docker: รวมถึง Backend
docker-compose up --build

#### 3. ติดตั้งและรัน Frontend:
cd frontend
npm install
npm run dev
#### 4. เข้าใช้งานระบบ:
🌐 Frontend: http://localhost:5173
⚙️ Backend: http://localhost:8080



