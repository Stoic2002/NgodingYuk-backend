package main

import (
	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"gorm.io/gorm"
)

func seedSQLBeginnerCourse(db *gorm.DB) {
	courseSQLBeg := domain.Course{
		ID:            CourseSQLBeg, // Using the same ID mapped previously for SQL Beginner
		Slug:          "sql-beginner",
		Language:      "sql",
		Level:         "beginner",
		TitleID:       "SQL untuk Pemula",
		TitleEN:       sp("SQL for Beginners"),
		DescriptionID: sp("Pelajari dasar-dasar SQL, bahasa standar untuk berinteraksi dengan database relasional. Mulai dari nol hingga mahir mengambil data!"),
		DescriptionEN: sp("Learn the basics of SQL, the standard language for interacting with relational databases. From zero to proficient in data retrieval!"),
		ThumbnailURL:  sp("https://storage.googleapis.com/ngodingyuk-assets/course/sql-beg.png"),
		OrderIndex:    10,
	}
	upsertCourse(db, &courseSQLBeg)

	mod1 := moduleUUID(4, 1)
	mod2 := moduleUUID(4, 2)
	mod3 := moduleUUID(4, 3)
	mod4 := moduleUUID(4, 4)

	modules := []domain.Module{
		{ID: mod1, CourseID: CourseSQLBeg, TitleID: "Fondasi & SELECT", TitleEN: sp("Foundations & SELECT"), OrderIndex: 1},
		{ID: mod2, CourseID: CourseSQLBeg, TitleID: "Filtering & Sorting", TitleEN: sp("Filtering & Sorting"), OrderIndex: 2},
		{ID: mod3, CourseID: CourseSQLBeg, TitleID: "Data Kosong & Metadata", TitleEN: sp("Null Data & Metadata"), OrderIndex: 3},
		{ID: mod4, CourseID: CourseSQLBeg, TitleID: "Agregasi & Grouping", TitleEN: sp("Aggregation & Grouping"), OrderIndex: 4},
	}

	for _, mod := range modules {
		upsertModule(db, &mod)
	}

	lessons := []domain.Lesson{
		{
			ID:         lessonUUID(4, 1),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod1,
			OrderIndex: 1,
			XPReward:   10,
			TitleID:    "1. Pengenalan SQL dan Database",
			TitleEN:    sp("1. Introduction to SQL and Databases"),
			ContentMarkdownID: `### Apa itu Database?
Database (basis data) adalah sekumpulan data yang diorganisasikan secara terstruktur, memfasilitasi penyimpanan, pencarian, dan pengelolaan informasi yang besar.

### Apa itu SQL?
SQL (Structured Query Language) adalah bahasa standar yang dirancang khusus untuk mengelola dan memanipulasi data dalam database relasional (Relational Database Management System - RDBMS) seperti PostgreSQL, MySQL, dan SQLite.

**Mengapa belajar SQL?**
- SQL menjadi bahasa utama di dunia data (Data Science, Data Engineering, Backend Development).
- Sangat mudah dipelajari karena sintaksnya menyerupai bahasa Inggris.

Mari kita mulai perjalanan belajar SQL!`,
			ContentMarkdownEN: sp(`### What is a Database?
A database is an organized collection of structured data, typically stored electronically, facilitating data management and retrieval.

### What is SQL?
SQL (Structured Query Language) is the standard language designed for managing and manipulating data in Relational Database Management Systems (RDBMS) like PostgreSQL, MySQL, and SQLite.

**Why learn SQL?**
- It's the primary language in the data world (Data Science, Data Engineering, Backend Development).
- It's highly readable and closely resembles plain English.

Let's start our SQL journey!`),
		},
		{
			ID:         lessonUUID(4, 2),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod1,
			OrderIndex: 2,
			XPReward:   10,
			TitleID:    "2. Struktur Tabel Data",
			TitleEN:    sp("2. Data Table Structure"),
			ContentMarkdownID: `### Bagaimana Data Disimpan?
Dalam database relasional, data disimpan dalam bentuk **Tabel** (seperti spreadsheet di Excel). 

Sebuah tabel terdiri dari:
1. **Kolom (Column / Field):** Bagian vertikal dari tabel yang menyimpan jenis data tertentu (misalnya Nama atau Umur).
2. **Baris (Row / Record):** Bagian horizontal yang merepresentasikan satu entitas data utuh (misalnya data Budi dengan semua informasinya).

Contoh tabel ` + "`" + `users` + "`" + `:
| id | name | age | country |
|---|---|---|---|
| 1 | Alice | 24 | USA |
| 2 | Budi | 20 | Indonesia |

Di pelajaran berikutnya, kita akan belajar cara mengambil data ini menggunakan SQL!`,
			ContentMarkdownEN: sp(`### How is Data Stored?
In relational databases, data is stored in **Tables** (similar to an Excel spreadsheet).

A table consists of:
1. **Columns (Fields):** The vertical part of the table representing specific data attributes (like Name or Age).
2. **Rows (Records):** The horizontal part representing a single complete entity (like all of Budi's data).

Example table ` + "`" + `users` + "`" + `:
| id | name | age | country |
|---|---|---|---|
| 1 | Alice | 24 | USA |
| 2 | Budi | 20 | Indonesia |

In the next lesson, we will learn how to retrieve this data using SQL!`),
		},
		{
			ID:         lessonUUID(4, 3),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod1,
			OrderIndex: 3,
			XPReward:   10,
			TitleID:    "3. Mengambil Data dengan SELECT",
			TitleEN:    sp("3. Retrieving Data with SELECT"),
			ContentMarkdownID: `### Perintah SELECT
Perintah terpenting dalam SQL adalah ` + "`" + `SELECT` + "`" + `. Perintah ini digunakan untuk membaca / mengambil data dari database.

**Sintaks dasar:**
` + "```" + `sql
SELECT nama_kolom
FROM nama_tabel;
` + "```" + `

### Menggunakan Astersik (*)
Jika kamu ingin mengambil **seluruh kolom** dalam sebuah tabel, gunakan tanda bintang ` + "`" + `*` + "`" + `:
` + "```" + `sql
SELECT * FROM users;
` + "```" + `
Perintah di atas akan menampilkan semua kolom dan semua baris yang ada di dalam tabel ` + "`" + `users` + "`" + `.`,
			ContentMarkdownEN: sp(`### The SELECT Statement
The most critical SQL command is ` + "`" + `SELECT` + "`" + `. It is used to read / retrieve data from the database.

**Basic Syntax:**
` + "```" + `sql
SELECT column_name
FROM table_name;
` + "```" + `

### Using the Asterisk (*)
If you want to fetch **all columns** in a table, use the asterisk ` + "`" + `*` + "`" + `:
` + "```" + `sql
SELECT * FROM users;
` + "```" + `
The command above will return all rows and all columns from the ` + "`" + `users` + "`" + ` table.`),
		},
		{
			ID:         lessonUUID(4, 4),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod1,
			OrderIndex: 4,
			XPReward:   10,
			TitleID:    "4. Memilih Kolom Spesifik",
			TitleEN:    sp("4. Selecting Specific Columns"),
			ContentMarkdownID: `### Variasi Kolom
Kadang kita tidak butuh semua data yang ada di tabel. Kita bisa menentukan secara spesifik kolom mana saja yang ingin ditampilkan dengan memisahkannya menggunakan tanda koma.

` + "```" + `sql
SELECT name, country 
FROM users;
` + "```" + `

Perintah di atas hanya akan menampilkan kolom **name** dan **country**, tanpa kolom lainnya seperti id atau age. Hal ini sangat baik untuk alasan performa database ketika datanya jutaan.`,
			ContentMarkdownEN: sp(`### Specific Columns
Often we don't need all data from a table. We can specifically declare which columns to display by separating them with commas.

` + "```" + `sql
SELECT name, country 
FROM users;
` + "```" + `

This query will only show the **name** and **country** columns. This practice is crucial for database optimization when dealing with millions of records.`),
		},
		{
			ID:         lessonUUID(4, 5),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod2,
			OrderIndex: 5,
			XPReward:   10,
			TitleID:    "5. Menyaring Data dengan WHERE",
			TitleEN:    sp("5. Filtering Data with WHERE"),
			ContentMarkdownID: `### Klausul WHERE
Bagaimana jika kita hanya ingin menampilkan pengguna yang berumur 20 tahun? Kita bisa menggunakan klausul ` + "`" + `WHERE` + "`" + ` untuk mem-filter data.

**Sintaks:**
` + "```" + `sql
SELECT nama_kolom
FROM nama_tabel
WHERE kondisi;
` + "```" + `

**Contoh:**
` + "```" + `sql
SELECT * FROM users
WHERE age = 20;
` + "```" + `
Hanya baris yang nilai kolom age-nya tepat 20 yang akan dimunculkan.`,
			ContentMarkdownEN: sp(`### The WHERE Clause
What if we only want to show users who are exactly 20 years old? We use the ` + "`" + `WHERE` + "`" + ` clause to filter rows.

**Syntax:**
` + "```" + `sql
SELECT column_name
FROM table_name
WHERE condition;
` + "```" + `

**Example:**
` + "```" + `sql
SELECT * FROM users
WHERE age = 20;
` + "```" + `
Only rows where the age column equals 20 will be displayed.`),
		},
		{
			ID:         lessonUUID(4, 6),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod2,
			OrderIndex: 6,
			XPReward:   10,
			TitleID:    "6. Operator Perbandingan",
			TitleEN:    sp("6. Comparison Operators"),
			ContentMarkdownID: `### Operator di Klausul WHERE
Selain menggunakan sama dengan (` + "`" + `=` + "`" + `), kita bisa menggunakan operator matematika lainnya:

- ` + "`" + `<` + "`" + ` : Kurang dari
- ` + "`" + `>` + "`" + ` : Lebih dari
- ` + "`" + `<=` + "`" + ` : Kurang dari atau sama dengan
- ` + "`" + `>=` + "`" + ` : Lebih dari atau sama dengan
- ` + "`" + `<>` + "`" + ` atau ` + "`" + `!=` + "`" + ` : Tidak sama dengan

**Contoh Kasus:**
Tampilkan pengguna yang berumur di atas 18 tahun:
` + "```" + `sql
SELECT id, name FROM users WHERE age > 18;
` + "```" + `
Tampilkan pekerja dengan gaji tidak sama dengan 5000:
` + "```" + `sql
SELECT * FROM employees WHERE salary <> 5000;
` + "```" + ``,
			ContentMarkdownEN: sp(`### Operators in WHERE Clause
Besides equals (` + "`" + `=` + "`" + `), we can use mathematical comparison operators:

- ` + "`" + `<` + "`" + ` : Less than
- ` + "`" + `>` + "`" + ` : Greater than
- ` + "`" + `<=` + "`" + ` : Less than or equal
- ` + "`" + `>=` + "`" + ` : Greater than or equal
- ` + "`" + `<>` + "`" + ` or ` + "`" + `!=` + "`" + ` : Not equal to

**Example:**
Show users older than 18:
` + "```" + `sql
SELECT id, name FROM users WHERE age > 18;
` + "```" + `
Show employees whose salary is not equal to 5000:
` + "```" + `sql
SELECT * FROM employees WHERE salary <> 5000;
` + "```" + ``),
		},
		{
			ID:         lessonUUID(4, 7),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod2,
			OrderIndex: 7,
			XPReward:   10,
			TitleID:    "7. Operator Logika (AND, OR, NOT)",
			TitleEN:    sp("7. Logical Operators (AND, OR, NOT)"),
			ContentMarkdownID: `### Menggabungkan Kondisi
Jika kamu punya lebih dari 1 syarat filter, kamu bisa menggabungkannya dengan operator logika.

- **AND**: Data akan tampil jika KEDUA kondisi benar.
` + "```" + `sql
SELECT * FROM users WHERE age >= 20 AND country = 'Indonesia';
` + "```" + `
- **OR**: Data akan tampil jika SALAH SATU kondisi benar.
` + "```" + `sql
SELECT * FROM users WHERE country = 'Indonesia' OR country = 'USA';
` + "```" + `
- **NOT**: Data akan dibalikkan (menjadi lawan kondisinya).
` + "```" + `sql
SELECT * FROM users WHERE NOT country = 'Malaysia';
` + "```" + ``,
			ContentMarkdownEN: sp(`### Combining Conditions
If you have multiple valid criteria, you can combine them using logical operators.

- **AND**: Returns the row if BOTH conditions are true.
` + "```" + `sql
SELECT * FROM users WHERE age >= 20 AND country = 'Indonesia';
` + "```" + `
- **OR**: Returns the row if AT LEAST ONE condition is true.
` + "```" + `sql
SELECT * FROM users WHERE country = 'Indonesia' OR country = 'USA';
` + "```" + `
- **NOT**: Reverses the true/false outcome.
` + "```" + `sql
SELECT * FROM users WHERE NOT country = 'Malaysia';
` + "```" + ``),
		},
		{
			ID:         lessonUUID(4, 8),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod2,
			OrderIndex: 8,
			XPReward:   10,
			TitleID:    "8. Menggunakan IN dan BETWEEN",
			TitleEN:    sp("8. Using IN and BETWEEN"),
			ContentMarkdownID: `### Operator IN
Operator ` + "`" + `IN` + "`" + ` sangat cocok sebagai pengganti OR yang banyak.
Daripada menulis: ` + "`" + `WHERE country = 'UK' OR country = 'USA' OR country = 'Canada'` + "`" + `,
Kamu bisa menulis:
` + "```" + `sql
SELECT * FROM users WHERE country IN ('UK', 'USA', 'Canada');
` + "```" + `

### Operator BETWEEN
` + "`" + `BETWEEN` + "`" + ` menyeleksi sebuah nilai yang berada dalam sebuah rentang.
Ambil data pengguna berumur antara 20 sampai 30 tahun (20 dan 30 ikut masuk):
` + "```" + `sql
SELECT * FROM users WHERE age BETWEEN 20 AND 30;
` + "```" + ``,
			ContentMarkdownEN: sp(`### IN Operator
The ` + "`" + `IN` + "`" + ` operator is a great shorthand for multiple OR conditions.
Instead of: ` + "`" + `WHERE country = 'UK' OR country = 'USA' OR country = 'Canada'` + "`" + `,
You write:
` + "```" + `sql
SELECT * FROM users WHERE country IN ('UK', 'USA', 'Canada');
` + "```" + `

### BETWEEN Operator
` + "`" + `BETWEEN` + "`" + ` selects values within a given range (inclusive).
Get users aged 20 to 30:
` + "```" + `sql
SELECT * FROM users WHERE age BETWEEN 20 AND 30;
` + "```" + ``),
		},
		{
			ID:         lessonUUID(4, 9),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod2,
			OrderIndex: 9,
			XPReward:   10,
			TitleID:    "9. Pencarian Teks dengan LIKE",
			TitleEN:    sp("9. Text Search with LIKE"),
			ContentMarkdownID: `### Pola Teks dengan LIKE
Jika kita tidak tahu teks spesifik dan ingin mencari polanya saja, gunakan klausul ` + "`" + `LIKE` + "`" + `.
LIKE dibantu dengan karakter wildcard ` + "`" + `%` + "`" + ` (mewakili jumlah huruf berapapun).

**Contoh:**
Mencari nama yang **berawalan** huruf 'A':
` + "```" + `sql
SELECT * FROM users WHERE name LIKE 'A%';
` + "```" + `
Mencari nama yang **berakhiran** huruf 'n':
` + "```" + `sql
SELECT * FROM users WHERE name LIKE '%n';
` + "```" + `
Mencari nama yang **mengandung** kata 'bud':
` + "```" + `sql
SELECT * FROM users WHERE name LIKE '%bud%';
` + "```" + ``,
			ContentMarkdownEN: sp(`### Text Patterns with LIKE
If we don't know the exact text and just want to search via patterns, we use the ` + "`" + `LIKE` + "`" + ` clause.
It's often paired with the wildcard character ` + "`" + `%` + "`" + ` (representing any sequence of characters).

**Examples:**
Names **starting** with 'A':
` + "```" + `sql
SELECT * FROM users WHERE name LIKE 'A%';
` + "```" + `
Names **ending** with 'n':
` + "```" + `sql
SELECT * FROM users WHERE name LIKE '%n';
` + "```" + `
Names **containing** the word 'bud':
` + "```" + `sql
SELECT * FROM users WHERE name LIKE '%bud%';
` + "```" + ``),
		},
		{
			ID:         lessonUUID(4, 10),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod2,
			OrderIndex: 10,
			XPReward:   10,
			TitleID:    "10. Mengurutkan Data dengan ORDER BY",
			TitleEN:    sp("10. Sorting Data with ORDER BY"),
			ContentMarkdownID: `### Klausul ORDER BY
Data dalam tabel tidak selalu berurutan sesuai keinginan kita. Untuk mengurutkannya, gunakan ` + "`" + `ORDER BY` + "`" + `.

Urutan Ascending (Menaik / Kecil ke Besar):
` + "```" + `sql
SELECT * FROM users ORDER BY age ASC;
` + "```" + `
*(Catatan: ASC adalah default, jadi boleh tidak ditulis).*

Urutan Descending (Menurun / Besar ke Kecil):
` + "```" + `sql
SELECT * FROM users ORDER BY age DESC;
` + "```" + ``,
			ContentMarkdownEN: sp(`### ORDER BY Clause
Database queries don't return data in a specific order by default. Use ` + "`" + `ORDER BY` + "`" + ` to sort result sets.

Ascending Order (Small to Large / A-Z):
` + "```" + `sql
SELECT * FROM users ORDER BY age ASC;
` + "```" + `
*(Note: ASC is the default behavior).*

Descending Order (Large to Small / Z-A):
` + "```" + `sql
SELECT * FROM users ORDER BY age DESC;
` + "```" + ``),
		},
		{
			ID:         lessonUUID(4, 11),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod2,
			OrderIndex: 11,
			XPReward:   10,
			TitleID:    "11. Membatasi Hasil dengan LIMIT",
			TitleEN:    sp("11. Limiting Results with LIMIT"),
			ContentMarkdownID: `### Klausul LIMIT
Terkadang kita hanya butuh 10 data teratas dari ribuan data yang ada. Gunakan ` + "`" + `LIMIT` + "`" + `.

**Contoh:** Ambil 5 user dengan umur termuda.
` + "```" + `sql
SELECT * FROM users
ORDER BY age ASC
LIMIT 5;
` + "```" + `
*Tips: ORDER BY sering dipakai bersamaan dengan LIMIT untuk mengambil "Top N" data.*`,
			ContentMarkdownEN: sp(`### LIMIT Clause
Sometimes you only need the top 10 results from a 1,000-row table. Use ` + "`" + `LIMIT` + "`" + ` (PostgreSQL/MySQL) to restrict the number of returned rows.

**Example:** Get the 5 youngest users.
` + "```" + `sql
SELECT * FROM users
ORDER BY age ASC
LIMIT 5;
` + "```" + `
*Tip: ORDER BY is frequently paired with LIMIT to fetch the "Top N" rows.*`),
		},
		{
			ID:         lessonUUID(4, 12),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod3,
			OrderIndex: 12,
			XPReward:   10,
			TitleID:    "12. Nilai NULL dan IS NULL",
			TitleEN:    sp("12. NULL Values and IS NULL"),
			ContentMarkdownID: `### Apa itu NULL?
` + "`" + `NULL` + "`" + ` adalah state/kemungkinan kondisi di mana sebuah data itu **tidak ada nilainya / kosong mutlak**. NULL **bukanlah** angka nol (0) dan bukan string kosong ('').

Untuk mengecek NULL, kita **TIDAK BOLEH** memakai tanda sama dengan ` + "`" + `=` + "`" + `. Kita menggunakan ` + "`" + `IS NULL` + "`" + ` atau ` + "`" + `IS NOT NULL` + "`" + `.

Ambil data pengguna yang nomor HP-nya belum diisi:
` + "```" + `sql
SELECT * FROM users WHERE phone_number IS NULL;
` + "```" + `
Ambil yang nomor HP-nya sudah ada:
` + "```" + `sql
SELECT * FROM users WHERE phone_number IS NOT NULL;
` + "```" + ``,
			ContentMarkdownEN: sp(`### What is NULL?
` + "`" + `NULL` + "`" + ` means a field has **no value / is absolutely empty**. It is **NOT** a zero value (0) nor an empty space ('').

To test for NULL values, we **CANNOT** use comparison operators like ` + "`" + `=` + "`" + `. We must use ` + "`" + `IS NULL` + "`" + ` or ` + "`" + `IS NOT NULL` + "`" + `.

Get users who haven't filled in their phone numbers:
` + "```" + `sql
SELECT * FROM users WHERE phone_number IS NULL;
` + "```" + `
Get users who have a phone number:
` + "```" + `sql
SELECT * FROM users WHERE phone_number IS NOT NULL;
` + "```" + ``),
		},
		{
			ID:         lessonUUID(4, 13),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod3,
			OrderIndex: 13,
			XPReward:   10,
			TitleID:    "13. Alias Kolom dengan AS",
			TitleEN:    sp("13. Column Aliases with AS"),
			ContentMarkdownID: `### Klausul AS
Dalam hasil query, nama kolom terkadang kurang jelas. Kita dapat menamainya sementara secara visual dengan ` + "`" + `AS` + "`" + ` (Alias).

` + "```" + `sql
SELECT 
    name AS nama_lengkap, 
    age AS umur_user
FROM users;
` + "```" + `

Hasil tabel yang keluar akan menggunakan header ` + "`" + `nama_lengkap` + "`" + ` dan ` + "`" + `umur_user` + "`" + ` alih-alih bawaan nama aslinya.`,
			ContentMarkdownEN: sp(`### AS Clause
Column names in your original tables might not always be readable. You can give them temporary visual aliases with ` + "`" + `AS` + "`" + `.

` + "```" + `sql
SELECT 
    name AS full_name, 
    age AS user_age
FROM users;
` + "```" + `

The resulting output will display the columns header as ` + "`" + `full_name` + "`" + ` and ` + "`" + `user_age` + "`" + `.`),
		},
		{
			ID:         lessonUUID(4, 14),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod4,
			OrderIndex: 14,
			XPReward:   10,
			TitleID:    "14. Pengenalan Fungsi Agregasi",
			TitleEN:    sp("14. Introduction to Aggregate Functions"),
			ContentMarkdownID: `### Apa itu Fungsi Agregat?
Sering kali bukan baris demi baris data yang kita inginkan, namun kalkulasi atau "rangkuman" dari data kolom tersebut.
Berapa jumlah pegawai? Berapa total gaji? Berapa nilai rata-rata?

SQL menyediakan Aggregate Functions yang sering dipakai:
- **COUNT()**
- **SUM()**
- **AVG()**
- **MIN()**
- **MAX()**

Di modul selanjutnya kita bahas tuntas satu per satu.`,
			ContentMarkdownEN: sp(`### What are Aggregate Functions?
Often, we don't need row-by-row data. Instead, we want an aggregated calculation mathematically representing the dataset.
How many workers? What's the total salaries? What's the average score?

SQL handles this natively with Aggregate Functions:
- **COUNT()**
- **SUM()**
- **AVG()**
- **MIN()**
- **MAX()**

Let's dive into each completely in the following lessons.`),
		},
		{
			ID:         lessonUUID(4, 15),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod4,
			OrderIndex: 15,
			XPReward:   10,
			TitleID:    "15. Menghitung Baris dengan COUNT",
			TitleEN:    sp("15. Counting Rows with COUNT"),
			ContentMarkdownID: `### Fungsi COUNT()
` + "`" + `COUNT()` + "`" + ` digunakan untuk menghitung jumlah baris / himpunan data yang cocok dengan kriteria yang diminta.

Menghitung total seluruh user:
` + "```" + `sql
SELECT COUNT(*) FROM users;
` + "```" + `

Menghitung total user yang dari 'USA':
` + "```" + `sql
SELECT COUNT(*) AS total_usa 
FROM users 
WHERE country = 'USA';
` + "```" + ``,
			ContentMarkdownEN: sp(`### COUNT() Function
` + "`" + `COUNT()` + "`" + ` determines the number of rows / records matching specified criteria.

Count all users:
` + "```" + `sql
SELECT COUNT(*) FROM users;
` + "```" + `

Count users from the 'USA' only:
` + "```" + `sql
SELECT COUNT(*) AS total_usa 
FROM users 
WHERE country = 'USA';
` + "```" + ``),
		},
		{
			ID:         lessonUUID(4, 16),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod4,
			OrderIndex: 16,
			XPReward:   10,
			TitleID:    "16. Fungsi SUM dan AVG",
			TitleEN:    sp("16. SUM and AVG Functions"),
			ContentMarkdownID: `### Fungsi SUM()
Gunakan nilai numerik di dalam ` + "`" + `SUM()` + "`" + ` untuk menotalkan jumlah kolom secara keseluruhan.
Contoh total gaji departemen IT:
` + "```" + `sql
SELECT SUM(salary) FROM employees WHERE department = 'IT';
` + "```" + `

### Fungsi AVG()
Gunakan ` + "`" + `AVG()` + "`" + ` (Average) untuk mencari rata-rata kolom numerik.
Berapa nilai rata-rata raport siswa?
` + "```" + `sql
SELECT AVG(score) FROM exam_results;
` + "```" + ``,
			ContentMarkdownEN: sp(`### SUM() Function
Use a numeric column within ` + "`" + `SUM()` + "`" + ` to add up the entire column's total value.
Example total IT department salary expense:
` + "```" + `sql
SELECT SUM(salary) FROM employees WHERE department = 'IT';
` + "```" + `

### AVG() Function
` + "`" + `AVG()` + "`" + ` calculates the arithmetic average of a numeric column.
What is the mean student examination exam score?
` + "```" + `sql
SELECT AVG(score) FROM exam_results;
` + "```" + ``),
		},
		{
			ID:         lessonUUID(4, 17),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod4,
			OrderIndex: 17,
			XPReward:   10,
			TitleID:    "17. Fungsi MIN dan MAX",
			TitleEN:    sp("17. MIN and MAX Functions"),
			ContentMarkdownID: `### Mencari Nilai Ekstrem
Sederhana, ` + "`" + `MIN()` + "`" + ` memberikan nilai paling kecil di dalam kolom.
` + "```" + `sql
SELECT MIN(price) FROM products;
` + "```" + `

Sedangkan ` + "`" + `MAX()` + "`" + ` memberikan nilai dengan besaran terbesar di kolom tersebut.
` + "```" + `sql
SELECT MAX(price) FROM products;
` + "```" + ``,
			ContentMarkdownEN: sp(`### Finding Extreme Extremities
Simply put, ` + "`" + `MIN()` + "`" + ` will return the smallest possible value in an entire column.
` + "```" + `sql
SELECT MIN(price) FROM products;
` + "```" + `

In contrast, ` + "`" + `MAX()` + "`" + ` returns the largest column value.
` + "```" + `sql
SELECT MAX(price) FROM products;
` + "```" + ``),
		},
		{
			ID:         lessonUUID(4, 18),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod4,
			OrderIndex: 18,
			XPReward:   10,
			TitleID:    "18. Mengelompokkan Data dengan GROUP BY",
			TitleEN:    sp("18. Grouping Data with GROUP BY"),
			ContentMarkdownID: `### Klausul GROUP BY
Group By biasanya selalu jalan beriringan dengan Fungsi Agregat (COUNT, SUM, AVG). Ini digunakan untuk mengelompokkan data bersarkan 1 atau lebih kolom, sehingga agregasi bekerja berdasarkan grup.

Berapa banyak jumlah pegawai per masing-masing divisi?
` + "```" + `sql
SELECT department, COUNT(*) AS jumlah_pegawai
FROM employees
GROUP BY department;
` + "```" + `
Sistem akan memecahnya misal: IT: 5, HR: 3, Sales: 4.`,
			ContentMarkdownEN: sp(`### GROUP BY Clause
Group By often partners directly with Aggregate functions. It groups rows sharing identical values so formulas are computed sequentially block-by-block.

How many workers are present per division department?
` + "```" + `sql
SELECT department, COUNT(*) AS total_employees
FROM employees
GROUP BY department;
` + "```" + `
The DB engines calculates this dynamically (e.g. IT: 5, HR: 3, Sales: 4).`),
		},
		{
			ID:         lessonUUID(4, 19),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod4,
			OrderIndex: 19,
			XPReward:   10,
			TitleID:    "19. Menyaring Grup dengan HAVING",
			TitleEN:    sp("19. Filtering Groups with HAVING"),
			ContentMarkdownID: `### Mengapa ada HAVING?
Klausul ` + "`" + `WHERE` + "`" + ` **tidak bisa** dipakai mengefilater fungsi agregasi! Kita tidak bisa berkata ` + "`" + `WHERE COUNT(*) > 5` + "`" + `.
Sebagai gantinya, gunakan ` + "`" + `HAVING` + "`" + ` tepat setelah ` + "`" + `GROUP BY` + "`" + `.

Tampilkan hanya departemen yang jumlah pegawainya lebih dari 5 orang:
` + "```" + `sql
SELECT department, COUNT(*)
FROM employees
GROUP BY department
HAVING COUNT(*) > 5;
` + "```" + ``,
			ContentMarkdownEN: sp(`### Why does HAVING exist?
The ` + "`" + `WHERE` + "`" + ` clause **cannot** filter an aggregate function directly. Say goodbye to logic like ` + "`" + `WHERE COUNT(*) > 5` + "`" + `.
Instead, integrate the ` + "`" + `HAVING` + "`" + ` keyword placed precisely immediately following your ` + "`" + `GROUP BY` + "`" + `.

Show only departments containing more than 5 members:
` + "```" + `sql
SELECT department, COUNT(*)
FROM employees
GROUP BY department
HAVING COUNT(*) > 5;
` + "```" + ``),
		},
		{
			ID:         lessonUUID(4, 20),
			CourseID:   CourseSQLBeg,
			ModuleID:   &mod4,
			OrderIndex: 20,
			XPReward:   10,
			TitleID:    "20. Ringkasan Sintaks SQL Dasar",
			TitleEN:    sp("20. Basic SQL Syntax Overview"),
			ContentMarkdownID: `### Struktur Utama Query SQL
Selamat! Kamu telah menguasai anatomi dasar SQL. Jika diurutkan, beginilah cara menyusun perintah SQL dari yang paling atas hingga paling bawah:

` + "```" + `sql
SELECT kolom1, kolom2, fungsi_agregasi()
FROM nama_tabel
WHERE filter_kondisi_biasa
GROUP BY kategori_kolom
HAVING filter_kondisi_agregasi
ORDER BY urutan_kolom ASC/DESC
LIMIT jumlah_batas_baris;
` + "```" + `
*(Catatan: Kamu jarang menulis sepanjang ini untuk setiap query, gunakan sesuaikan dengan kebutuhanmu saja).*

Selesaikan kuis untuk menamatkan Kelas Dasar SQL ini!`,
			ContentMarkdownEN: sp(`### The Ultimate Request Scaffold
Congratulations! You've mastered core SQL anatomy. If arranged syntactically, here's standard execution hierarchy precedence top-to-bottom:

` + "```" + `sql
SELECT col1, col2, aggregate_fn()
FROM table_name
WHERE row_condition_evaluations
GROUP BY aggregation_category
HAVING group_evaluation_clause
ORDER BY sorting_columns ASC/DESC
LIMIT returned_rows_count;
` + "```" + `
*(Note: Usually queries are much leaner, solely invoking what's critically required).*

Proceed to finish the exams unlocking your basic SQL certification!`),
		},
	}

	for _, lesson := range lessons {
		// Taking address into a new variable iteration properly scope
		l := lesson
		upsertLesson(db, &l)
	}
}
