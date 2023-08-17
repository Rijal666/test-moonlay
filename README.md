# Moonlay Technologies - Backend Test(Golang)
saya membuat handler backend untuk Todo List App menggunakan golang dengan framework echo
>**Todo List App**
Aplikasi yang umumnya digunakan untuk memelihara tugas sehari-hari atau membuat daftar semua yang harus dilakukan, dengan urutan prioritas tugas tertinggi hingga terendah. Sangat membantu dalam merencanakan jadwal harian.
## Getting Started

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/Rial666/test-moonlay.git
   ```
2. Install Go packages
   ```sh
   go mod tidy
   ```
3. create you're Database in PostgreSQL and setting DBurl in folder pkg/mysql/mysql.go
   ```sh
	DBurl := "host=localhost user=postgres password=your-password dbname=your-dbname sslmode=disable"
   ```

### Usage

1. Run the App with.

   ```sh
   go run main.go
   ```
