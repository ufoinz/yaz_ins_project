# ToDo List Application

## Yazılım İnşası Dersi Projesi

Full-stack ToDo List application with user authentication and task management.

### Features
- User registration and login (JWT authentication)
- Create, edit, delete tasks with title, description and due date
- View all tasks for the current user
- Clean and responsive English UI
- REST API for all operations

### Screenshots
<img width="1919" height="1079" alt="image" src="https://github.com/user-attachments/assets/f48c5b08-a9ec-4706-a7ab-c3c373a7d4da" />
<img width="1920" height="1080" alt="image" src="https://github.com/user-attachments/assets/d4d85aa0-560b-4384-8378-45298799fbba" />


### Technologies
- **Backend**: Go (Gin framework), GORM, PostgreSQL
- **Frontend**: Angular 18
- **Authentication**: JWT
- **Database**: PostgreSQL

### Project Structure
```
yaz_ins_project/
├─ todo-backend/     # Go backend with Gin + GORM
├─ todo-frontend/    # Angular frontend
├─ .gitignore
└─ README.md
```

### How to Run

1. Clone the repository:
   ```bash
   git clone https://github.com/ufoinz/yaz_ins_project.git
   cd yaz_ins_project
   ```

2. **Backend:**
   ```bash
   cd todo-backend
   # Create .env file with your PostgreSQL credentials
   "DB_DSN=host=localhost user=postgres password=your_password dbname=todo_db port=5432 sslmode=disable
   PORT=8080
   JWT_SECRET=super_secret_key_2025"
   go mod tidy
   air
   ```
  Server will start at http://localhost:8080

3. **Frontend**:
   ```bash
   cd ../todo-frontend
   npm install
   npm start
   ```
   Application available at http://localhost:4200

   ### Creators
   **Beibarys Galymzhan, Alisho Odzhiev**
