task-management-system/ ├── backend/ # โค้ดฝั่ง Backend ที่พัฒนาโดย Golang │ ├── Dockerfile # Dockerfile สำหรับ Backend │ ├── go.mod # Dependencies ของ Golang │ ├── go.sum # Module Dependencies │ ├── main.go # Entry Point สำหรับ Backend ├── database/ # ไฟล์ตั้งค่าฐานข้อมูล │ └── init.sql # สคริปต์ SQL สำหรับสร้างและใส่ข้อมูลเริ่มต้น ├── frontend/ # โค้ดฝั่ง Frontend ที่พัฒนาโดย Vue.js │ ├── src/ │ │ ├── assets/ # ไฟล์ Static (รูปภาพ, ไอคอน) │ │ ├── components/ # ส่วนประกอบย่อย (Reusable Components) │ │ ├── router/ # การตั้งค่า Vue Router │ │ ├── views/ # หน้าหลักของแอปพลิเคชัน (Login, Task Board) │ │ ├── App.vue # Root Component ของ Vue │ │ ├── main.js # จุดเริ่มต้นของ Frontend │ └── public/ # ไฟล์ Static ที่เป็น Public ├── docker-compose.yml # การตั้งค่า Docker Compose ├── README.md # เอกสารแนะนำโปรเจค
task-management-system/ ├── backend/ # โค้ดฝั่ง Backend ที่พัฒนาโดย Golang │ ├── Dockerfile # Dockerfile สำหรับ Backend │ ├── go.mod # Dependencies ของ Golang │ ├── go.sum # Module Dependencies │ ├── main.go # Entry Point สำหรับ Backend ├── database/ # ไฟล์ตั้งค่าฐานข้อมูล │ └── init.sql # สคริปต์ SQL สำหรับสร้างและใส่ข้อมูลเริ่มต้น ├── frontend/ # โค้ดฝั่ง Frontend ที่พัฒนาโดย Vue.js │ ├── src/ │ │ ├── assets/ # ไฟล์ Static (รูปภาพ, ไอคอน) │ │ ├── components/ # ส่วนประกอบย่อย (Reusable Components) │ │ ├── router/ # การตั้งค่า Vue Router │ │ ├── views/ # หน้าหลักของแอปพลิเคชัน (Login, Task Board) │ │ ├── App.vue # Root Component ของ Vue │ │ ├── main.js # จุดเริ่มต้นของ Frontend │ └── public/ # ไฟล์ Static ที่เป็น Public ├── docker-compose.yml # การตั้งค่า Docker Compose ├── README.md # เอกสารแนะนำโปรเจค
