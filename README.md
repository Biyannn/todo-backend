# 📝 Todo List Fullstack App

A simple fullstack Todo List application built with **Go** for the backend and **React** for the frontend.  
The backend uses **JSON file storage** and is deployed on **Koyeb**, while the frontend is deployed on **Vercel**.

---

## 🚀 Live Demo

- **Frontend**: https://todo-frontend.vercel.app
- **Backend API**: https://remaining-idaline-ayyubian-37a73d46.koyeb.app

---

## 🧩 Features

- Create todo
- Update todo (toggle completed)
- Delete single todo
- Delete all todos
- Filter by status & priority
- REST API with JSON response

---

## 🛠 Tech Stack

### Backend
- Go (net/http)
- JSON file storage
- Docker
- Koyeb

### Frontend
- React (Vite)
- Axios
- Tailwind CSS
- Vercel

---

## 📡 API Endpoints

| Method | Endpoint | Description |
|------|--------|------------|
| GET | `/todos` | Get all todos |
| POST | `/todos` | Create new todo |
| PUT | `/todos/:id` | Toggle todo |
| DELETE | `/todos/:id` | Delete todo |
| DELETE | `/todos` | Delete all todos |

---

## ⚠️ Notes

- JSON storage is **not persistent** on cloud restart.
- This project is intended for **learning & portfolio purposes**.

---

## 👨‍💻 Author

**Ayyubian Ar Raufan Kamal**  
Junior Fullstack Developer
