
# Dating Sim Go

**`dating-sim-go`** is a backend application built in Go, designed to power a dating simulation experience. It provides features like user authentication, swipe mechanics, profile management, and more, all built with a focus on performance and scalability.

---

## ✨ Technical Document

for [Technical Document](https://docs.google.com/document/d/e/2PACX-1vT8flrtJAYy6lN1G3s0UQIlfc--3XKrt-L29wp6naFUV5xUimjFeq0nwi_heou5B2OGCOQ7ANscP7oF/pub) please go to this link

---

## 🚀 Features

### 🔑 User Authentication
- Routes for user login, registration, and session management.

### 💕 Swipe Mechanics
- Allows users to interact with potential matches.
- Optimized using **Redis** for caching swipe data and enhancing performance.

### 👤 Profile Management
- Supports creating, updating, and retrieving user profiles.

### ⭐ Upgrade Services
- Handles premium features or in-app enhancements via dedicated routes.

### 🩺 Health Check
- A `/health` endpoint to monitor application readiness and status.

---

## 🛠️ Tech Stack

- **Go (Golang)**: The core programming language, chosen for its speed and scalability.
- **Redis**: Used for caching dynamic data like swipe interactions.
- **SQL Database** (PostgreSQL/MySQL): Persistent storage for user data and settings.
- **Gorilla Mux**: A router library for clean and organized API endpoint management.
- **Custom Middleware**: Handles additional request/response processes such as logging and authentication.
- **Environment Variables**: Managed through a custom `dotenv` package.

---

## 🏗️ Architecture

1. **Handler-Based Structure**:
   - Clean separation of concerns for each feature (e.g., `Swipe`, `Auth`, `Profile`).
   - Grouped and organized routes for modularity.

2. **Custom Redis Client**:
   - `redis.RedisClient` is initialized at application startup and shared across handlers.

3. **Database Initialization**:
   - A dedicated initialization process (`db.Conn()`) ensures reliable database connections.

4. **Graceful Shutdown**:
   - Resources like Redis connections are cleaned up during server shutdown.

---

## 📦 Local Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/dating-sim-go.git
   cd dating-sim-go
   ```

2. Set up environment variables in a `.env` or just copy `.env.example` file:
   ```plaintext
   REDIS_ADDR=localhost:6379
   REDIS_PASSWORD=
   REDIS_DB=0
   DB_HOST=your-database-host
   DB_USER=your-database-user
   DB_PASSWORD=your-database-password
   DB_NAME=your-database-name
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

---

## 📦 Installation using docker compose

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/dating-sim-go.git
   cd dating-sim-go
   ```

2. Set up environment variables in a `.env` or just copy `.env.example` file:
   ```plaintext
   REDIS_ADDR=localhost:6379
   REDIS_PASSWORD=
   REDIS_DB=0
   DB_HOST=your-database-host
   DB_USER=your-database-user
   DB_PASSWORD=your-database-password
   DB_NAME=your-database-name
   ```
3. Build and start the services using Docker Compose:
   ```bash
    docker-compose up --build
   ```
4. Access the application at http://localhost:5000.

## 🧪 Testing

All testing of endpoints included at `dating-sim-go.postman_collection.json` using postman and other alternatives

## 🛡️ Endpoints

### **Health Check**
- **GET** `/health`

### **Auth Routes**
- Login, Registration, and Logout functionality.

### **Swipe Routes**
- Handles interactions with potential matches.

### **Profile Routes**
- Manage user profile information.

### **Upgrade Routes**
- Access premium or enhanced features.

---

## 🛞 Future Plans

1. **Gamification**:
   - Badges, levels, and rewards to enhance user engagement.
2. **AI Matching**:
   - A recommendation system based on user preferences and behavior.
3. **Real-Time Communication**:
   - WebSocket integration for live chat and notifications.

---

## 👨‍💻 Contributors

- **frhan23**: Project Owner and Main Developer

Feel free to contribute to this project by submitting issues or pull requests. 😊

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---

Enjoy building your dating simulation app! 💕
