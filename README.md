# 🚀 DockerTask4 - Simple REST API with Go Fiber & MySQL
Project ini merupakan implementasi **Simple REST API** menggunakan:
- Golang (Fiber Framework)
- MySQL Database
- Docker & Docker Compose

Fitur utama:
- ✅ POST (Create Note)
- ✅ GET (All Notes & By ID)
- ✅ PUT (Update Note)
- ✅ DELETE (Delete Note)

---

# 🧰 Tech Stack
- Golang (Fiber v3)
- GORM (ORM)
- MySQL
- Docker
- Docker Compose

---

# 📁 Struktur Project
DockerTask4/
│
├── config/
│└── env.go
│
├── controllers/
│└── note.go
│
├── databases/
│└── database.go
│
├── models/
│└── note.go
│
├── routes/
│└── note_routes.go
│
├── .env.example
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go # Entry point aplikasi
└── README.md

---

# ⚙️ Persiapan (WAJIB)
Pastikan sudah install:

### 1. Docker
Download:
https://www.docker.com/products/docker-desktop/

Cek di powershell:
```bash
docker --version
docker-compose --version
```

---

# 🚀 Cara Menjalankan Project
- Clone Repository di IDE yang kamu gunakan
```bash
git clone https://github.com/mauinGit/DockerTask4.git
cd DockerTask4
```

- Setup Environment
```bash
DB_USER=your_username
DB_PASSWORD=your_password
DB_HOST=your_merek_database
DB_PORT=your_port_number
DB_NAME=your_database_name
APP_PORT=3000
```

- note ports
pada file docker-compose.yml pada bagian
```bash
    ports:
       - "${YOUR_DB_PORT}:3306"
```
ubah ${YOUR_DB_PORT} menjadi angka ports yang biasanya kamu gunakan

- Jalankan Docker Compose
```bash
docker-compose up -d
```

- Jika kamu ingin melakukan perubahan pada code lakukan perubahan pada file, jalankan kembali docker compose
```bash
docker-compose up -d --build
```

---

# 📌 API Endpoint
Lakukan uji coba endpoint dengan menggunakan postman dengan menggunakan Body (form-data):
- POST (Create Note): `http://127.0.0.1:3000/note`
- GET (All Notes & By ID): `http://127.0.0.1:3000/note` & `http://127.0.0.1:3000/note/id`
- PUT (Update Note): `http://127.0.0.1:3000/note/id`
- DELETE (Delete Note): `http://127.0.0.1:3000/note/id`

---

# 👤 Author
Maulana Adiatma - Maret 2026