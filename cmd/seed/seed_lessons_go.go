package main

import (
	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"gorm.io/gorm"
)

func seedGoLessons(db *gorm.DB) {
	lessons := []domain.Lesson{
		// ===== golang-beginner (10 lessons) =====
		{ID: lessonUUID(1, 1), CourseID: CourseGoBeg, OrderIndex: 1, XPReward: 10,
			TitleID: "Pengenalan Go & Kenapa Go?", TitleEN: sp("Introduction to Go & Why Go?"),
			ContentMarkdownID: `# Pengenalan Go

Go (atau Golang) adalah bahasa pemrograman yang dikembangkan oleh Google pada tahun 2009. Go dirancang untuk kesederhanaan, efisiensi, dan performa tinggi.

## Kenapa Belajar Go?

1. **Sederhana & Mudah Dipelajari** — Sintaks Go sangat minimalis dibanding bahasa lain
2. **Cepat** — Go dikompilasi langsung ke machine code, hampir secepat C
3. **Concurrency Built-in** — Goroutine membuat pemrograman paralel sangat mudah
4. **Dipakai Perusahaan Besar** — Google, Uber, Tokopedia, Gojek semua pakai Go

## Program Pertama

` + "```go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Halo, Dunia!\")\n}\n```" + `

Setiap program Go dimulai dengan ` + "`package main`" + ` dan fungsi ` + "`main()`" + ` sebagai titik masuk. Package ` + "`fmt`" + ` digunakan untuk mencetak output.

## Menjalankan Program Go

Untuk menjalankan, simpan file sebagai ` + "`main.go`" + ` lalu jalankan:
` + "```bash\ngo run main.go\n```" + `

Go juga bisa dikompilasi menjadi binary executable:
` + "```bash\ngo build -o program main.go\n./program\n```",
			ContentMarkdownEN: sp(`# Introduction to Go

Go (or Golang) is a programming language developed by Google in 2009. Go is designed for simplicity, efficiency, and high performance.

## Why Learn Go?

1. **Simple & Easy to Learn** — Go's syntax is very minimalist compared to other languages
2. **Fast** — Go compiles directly to machine code, almost as fast as C
3. **Built-in Concurrency** — Goroutines make parallel programming very easy
4. **Used by Major Companies** — Google, Uber, and many startups use Go

## First Program

` + "```go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, World!\")\n}\n```" + `

Every Go program starts with ` + "`package main`" + ` and a ` + "`main()`" + ` function as the entry point. The ` + "`fmt`" + ` package is used for printing output.

## Running a Go Program

To run, save your file as ` + "`main.go`" + ` then execute:
` + "```bash\ngo run main.go\n```" + `

Go can also be compiled into a binary executable:
` + "```bash\ngo build -o program main.go\n./program\n```"),
		},
		{ID: lessonUUID(1, 2), CourseID: CourseGoBeg, OrderIndex: 2, XPReward: 10,
			TitleID: "Variabel & Tipe Data", TitleEN: sp("Variables & Data Types"),
			ContentMarkdownID: `# Variabel & Tipe Data

## Deklarasi Variabel

Go punya beberapa cara mendeklarasikan variabel:

` + "```go\n// Cara 1: var dengan tipe eksplisit\nvar nama string = \"Budi\"\nvar umur int = 25\n\n// Cara 2: short declaration (paling umum)\nnama := \"Budi\"\numur := 25\n\n// Cara 3: deklarasi tanpa nilai (zero value)\nvar skor int      // 0\nvar aktif bool     // false\nvar pesan string   // \"\"\n```" + `

## Tipe Data Dasar

| Tipe | Contoh | Zero Value |
|------|--------|------------|
| int | 42 | 0 |
| float64 | 3.14 | 0.0 |
| string | "hello" | "" |
| bool | true | false |

## Konversi Tipe

Go tidak melakukan konversi implisit. Kamu harus eksplisit:

` + "```go\nvar x int = 10\nvar y float64 = float64(x) // konversi int ke float64\n```" + `

## Multiple Assignment

` + "```go\na, b, c := 1, \"hello\", true\nfmt.Println(a, b, c) // 1 hello true\n```",
			ContentMarkdownEN: sp(`# Variables & Data Types

## Variable Declaration

Go has several ways to declare variables:

` + "```go\n// Method 1: var with explicit type\nvar name string = \"Budi\"\nvar age int = 25\n\n// Method 2: short declaration (most common)\nname := \"Budi\"\nage := 25\n\n// Method 3: declaration without value (zero value)\nvar score int      // 0\nvar active bool    // false\nvar message string // \"\"\n```" + `

## Basic Data Types

| Type | Example | Zero Value |
|------|---------|------------|
| int | 42 | 0 |
| float64 | 3.14 | 0.0 |
| string | "hello" | "" |
| bool | true | false |

## Type Conversion

Go does not do implicit conversion. You must be explicit:

` + "```go\nvar x int = 10\nvar y float64 = float64(x) // convert int to float64\n```" + `

## Multiple Assignment

` + "```go\na, b, c := 1, \"hello\", true\nfmt.Println(a, b, c) // 1 hello true\n```"),
		},
		{ID: lessonUUID(1, 3), CourseID: CourseGoBeg, OrderIndex: 3, XPReward: 10,
			TitleID: "Konstanta & Iota", TitleEN: sp("Constants & Iota"),
			ContentMarkdownID: "# Konstanta & Iota\n\nKonstanta adalah nilai yang tidak bisa diubah setelah dideklarasikan.\n\n```go\nconst Pi = 3.14159\nconst AppName = \"NgodingYuk\"\n```\n\n## Iota — Auto-Increment\n\n`iota` adalah fitur unik Go untuk membuat konstanta berurutan:\n\n```go\nconst (\n    Easy = iota   // 0\n    Medium        // 1\n    Hard          // 2\n)\n```\n\nIota sangat berguna untuk membuat enum-like values di Go. Setiap kali `iota` muncul dalam blok `const`, nilainya otomatis bertambah 1.",
			ContentMarkdownEN: sp("# Constants & Iota\n\nConstants are values that cannot be changed after declaration.\n\n```go\nconst Pi = 3.14159\nconst AppName = \"NgodingYuk\"\n```\n\n## Iota — Auto-Increment\n\n`iota` is a unique Go feature for creating sequential constants:\n\n```go\nconst (\n    Easy = iota   // 0\n    Medium        // 1\n    Hard          // 2\n)\n```\n\nIota is very useful for creating enum-like values in Go. Each time `iota` appears in a `const` block, its value automatically increments by 1."),
		},
		{ID: lessonUUID(1, 4), CourseID: CourseGoBeg, OrderIndex: 4, XPReward: 10,
			TitleID: "Operator", TitleEN: sp("Operators"),
			ContentMarkdownID: "# Operator di Go\n\n## Operator Aritmatika\n```go\na := 10\nb := 3\nfmt.Println(a + b)  // 13\nfmt.Println(a - b)  // 7\nfmt.Println(a * b)  // 30\nfmt.Println(a / b)  // 3 (integer division)\nfmt.Println(a % b)  // 1 (modulo / sisa bagi)\n```\n\n## Operator Perbandingan\n```go\nfmt.Println(5 == 5)  // true\nfmt.Println(5 != 3)  // true\nfmt.Println(5 > 3)   // true\nfmt.Println(5 <= 3)  // false\n```\n\n## Operator Logika\n```go\nfmt.Println(true && false)  // false (AND)\nfmt.Println(true || false)  // true (OR)\nfmt.Println(!true)          // false (NOT)\n```\n\nPerhatikan bahwa Go tidak punya operator `++i` (prefix), hanya `i++` (postfix) dan ini bukan ekspresi melainkan statement.",
			ContentMarkdownEN: sp("# Operators in Go\n\n## Arithmetic Operators\n```go\na := 10\nb := 3\nfmt.Println(a + b)  // 13\nfmt.Println(a - b)  // 7\nfmt.Println(a * b)  // 30\nfmt.Println(a / b)  // 3 (integer division)\nfmt.Println(a % b)  // 1 (modulo / remainder)\n```\n\n## Comparison Operators\n```go\nfmt.Println(5 == 5)  // true\nfmt.Println(5 != 3)  // true\nfmt.Println(5 > 3)   // true\nfmt.Println(5 <= 3)  // false\n```\n\n## Logical Operators\n```go\nfmt.Println(true && false)  // false (AND)\nfmt.Println(true || false)  // true (OR)\nfmt.Println(!true)          // false (NOT)\n```\n\nNote that Go doesn't have a `++i` (prefix) operator, only `i++` (postfix), and it's a statement, not an expression."),
		},
		{ID: lessonUUID(1, 5), CourseID: CourseGoBeg, OrderIndex: 5, XPReward: 15,
			TitleID: "Kontrol Alur: if, switch", TitleEN: sp("Control Flow: if, switch"),
			ContentMarkdownID: "# Kontrol Alur\n\n## If Statement\n```go\nskor := 85\nif skor >= 90 {\n    fmt.Println(\"A\")\n} else if skor >= 80 {\n    fmt.Println(\"B\")\n} else {\n    fmt.Println(\"C\")\n}\n```\n\nGo unik — kamu bisa deklarasi variabel di dalam if:\n```go\nif nilai := hitungNilai(); nilai > 50 {\n    fmt.Println(\"Lulus\")\n}\n```\n\n## Switch Statement\n```go\nhari := \"Senin\"\nswitch hari {\ncase \"Senin\", \"Selasa\":\n    fmt.Println(\"Awal minggu\")\ncase \"Sabtu\", \"Minggu\":\n    fmt.Println(\"Weekend!\")\ndefault:\n    fmt.Println(\"Hari biasa\")\n}\n```\n\nSwitch di Go tidak perlu `break` — otomatis berhenti setelah case cocok.",
			ContentMarkdownEN: sp("# Control Flow\n\n## If Statement\n```go\nscore := 85\nif score >= 90 {\n    fmt.Println(\"A\")\n} else if score >= 80 {\n    fmt.Println(\"B\")\n} else {\n    fmt.Println(\"C\")\n}\n```\n\nGo is unique — you can declare a variable inside if:\n```go\nif val := calculate(); val > 50 {\n    fmt.Println(\"Passed\")\n}\n```\n\n## Switch Statement\n```go\nday := \"Monday\"\nswitch day {\ncase \"Monday\", \"Tuesday\":\n    fmt.Println(\"Start of week\")\ncase \"Saturday\", \"Sunday\":\n    fmt.Println(\"Weekend!\")\ndefault:\n    fmt.Println(\"Regular day\")\n}\n```\n\nSwitch in Go doesn't need `break` — it automatically stops after a case matches."),
		},
		{ID: lessonUUID(1, 6), CourseID: CourseGoBeg, OrderIndex: 6, XPReward: 15,
			TitleID: "Loop dengan for", TitleEN: sp("Loops with for"),
			ContentMarkdownID: "# Loop dengan for\n\nGo hanya punya satu keyword loop: `for`. Tapi sangat fleksibel!\n\n## Classic For\n```go\nfor i := 0; i < 5; i++ {\n    fmt.Println(i)\n}\n```\n\n## While-style\n```go\nn := 0\nfor n < 10 {\n    n++\n}\n```\n\n## Infinite Loop\n```go\nfor {\n    // berjalan selamanya sampai break\n    break\n}\n```\n\n## Range (untuk slice/array)\n```go\nbuah := []string{\"apel\", \"jeruk\", \"mangga\"}\nfor i, b := range buah {\n    fmt.Printf(\"%d: %s\\n\", i, b)\n}\n```\n\nGunakan `_` jika tidak butuh index: `for _, b := range buah`",
			ContentMarkdownEN: sp("# Loops with for\n\nGo only has one loop keyword: `for`. But it's very flexible!\n\n## Classic For\n```go\nfor i := 0; i < 5; i++ {\n    fmt.Println(i)\n}\n```\n\n## While-style\n```go\nn := 0\nfor n < 10 {\n    n++\n}\n```\n\n## Infinite Loop\n```go\nfor {\n    // runs forever until break\n    break\n}\n```\n\n## Range (for slice/array)\n```go\nfruits := []string{\"apple\", \"orange\", \"mango\"}\nfor i, f := range fruits {\n    fmt.Printf(\"%d: %s\\n\", i, f)\n}\n```\n\nUse `_` if you don't need the index: `for _, f := range fruits`"),
		},
		{ID: lessonUUID(1, 7), CourseID: CourseGoBeg, OrderIndex: 7, XPReward: 15,
			TitleID: "Fungsi & Multiple Return", TitleEN: sp("Functions & Multiple Return"),
			ContentMarkdownID: "# Fungsi di Go\n\n## Deklarasi Fungsi\n```go\nfunc sapa(nama string) string {\n    return \"Halo, \" + nama + \"!\"\n}\n```\n\n## Multiple Return Values\nIni fitur khas Go yang sangat berguna:\n```go\nfunc bagi(a, b float64) (float64, error) {\n    if b == 0 {\n        return 0, fmt.Errorf(\"tidak bisa bagi dengan nol\")\n    }\n    return a / b, nil\n}\n\nhasil, err := bagi(10, 3)\nif err != nil {\n    fmt.Println(\"Error:\", err)\n} else {\n    fmt.Println(\"Hasil:\", hasil)\n}\n```\n\n## Named Return Values\n```go\nfunc swap(a, b int) (x, y int) {\n    x = b\n    y = a\n    return // naked return\n}\n```\n\n## Variadic Function\n```go\nfunc jumlah(angka ...int) int {\n    total := 0\n    for _, n := range angka {\n        total += n\n    }\n    return total\n}\nfmt.Println(jumlah(1, 2, 3, 4)) // 10\n```",
			ContentMarkdownEN: sp("# Functions in Go\n\n## Function Declaration\n```go\nfunc greet(name string) string {\n    return \"Hello, \" + name + \"!\"\n}\n```\n\n## Multiple Return Values\nThis is a signature Go feature:\n```go\nfunc divide(a, b float64) (float64, error) {\n    if b == 0 {\n        return 0, fmt.Errorf(\"cannot divide by zero\")\n    }\n    return a / b, nil\n}\n\nresult, err := divide(10, 3)\nif err != nil {\n    fmt.Println(\"Error:\", err)\n} else {\n    fmt.Println(\"Result:\", result)\n}\n```\n\n## Named Return Values\n```go\nfunc swap(a, b int) (x, y int) {\n    x = b\n    y = a\n    return // naked return\n}\n```\n\n## Variadic Function\n```go\nfunc sum(nums ...int) int {\n    total := 0\n    for _, n := range nums {\n        total += n\n    }\n    return total\n}\nfmt.Println(sum(1, 2, 3, 4)) // 10\n```"),
		},
		{ID: lessonUUID(1, 8), CourseID: CourseGoBeg, OrderIndex: 8, XPReward: 15,
			TitleID: "Array & Slice", TitleEN: sp("Arrays & Slices"),
			ContentMarkdownID: "# Array & Slice\n\n## Array (ukuran tetap)\n```go\nvar angka [3]int = [3]int{1, 2, 3}\nfmt.Println(angka[0]) // 1\nfmt.Println(len(angka)) // 3\n```\n\n## Slice (ukuran dinamis) — lebih sering dipakai\n```go\nbuah := []string{\"apel\", \"jeruk\", \"mangga\"}\nbuah = append(buah, \"pisang\")\nfmt.Println(buah) // [apel jeruk mangga pisang]\n```\n\n## Slice Operations\n```go\ndata := []int{10, 20, 30, 40, 50}\nfmt.Println(data[1:3])  // [20 30]\nfmt.Println(data[:2])   // [10 20]\nfmt.Println(data[3:])   // [40 50]\n```\n\n## Make\n```go\ns := make([]int, 5)     // len=5, cap=5\ns2 := make([]int, 0, 10) // len=0, cap=10\n```",
			ContentMarkdownEN: sp("# Arrays & Slices\n\n## Array (fixed size)\n```go\nvar nums [3]int = [3]int{1, 2, 3}\nfmt.Println(nums[0]) // 1\nfmt.Println(len(nums)) // 3\n```\n\n## Slice (dynamic size) — used more often\n```go\nfruits := []string{\"apple\", \"orange\", \"mango\"}\nfruits = append(fruits, \"banana\")\nfmt.Println(fruits) // [apple orange mango banana]\n```\n\n## Slice Operations\n```go\ndata := []int{10, 20, 30, 40, 50}\nfmt.Println(data[1:3])  // [20 30]\nfmt.Println(data[:2])   // [10 20]\nfmt.Println(data[3:])   // [40 50]\n```\n\n## Make\n```go\ns := make([]int, 5)     // len=5, cap=5\ns2 := make([]int, 0, 10) // len=0, cap=10\n```"),
		},
		{ID: lessonUUID(1, 9), CourseID: CourseGoBeg, OrderIndex: 9, XPReward: 15,
			TitleID: "Map", TitleEN: sp("Maps"),
			ContentMarkdownID: "# Map\n\nMap adalah struktur data key-value, mirip dictionary di Python.\n\n## Membuat Map\n```go\nharga := map[string]int{\n    \"nasi goreng\": 15000,\n    \"mie ayam\":    12000,\n    \"es teh\":       5000,\n}\nfmt.Println(harga[\"nasi goreng\"]) // 15000\n```\n\n## Operasi Map\n```go\n// Tambah/update\nharga[\"sate\"] = 20000\n\n// Hapus\ndelete(harga, \"es teh\")\n\n// Cek keberadaan key\nnilai, ada := harga[\"bakso\"]\nif !ada {\n    fmt.Println(\"Bakso belum ada di menu\")\n}\n```\n\n## Iterasi Map\n```go\nfor menu, h := range harga {\n    fmt.Printf(\"%s: Rp%d\\n\", menu, h)\n}\n```\n\nPerlu diingat: urutan iterasi map di Go bersifat random.",
			ContentMarkdownEN: sp("# Maps\n\nMap is a key-value data structure, similar to dictionaries in Python.\n\n## Creating a Map\n```go\nprices := map[string]int{\n    \"fried rice\": 15000,\n    \"noodles\":    12000,\n    \"iced tea\":    5000,\n}\nfmt.Println(prices[\"fried rice\"]) // 15000\n```\n\n## Map Operations\n```go\n// Add/update\nprices[\"satay\"] = 20000\n\n// Delete\ndelete(prices, \"iced tea\")\n\n// Check if key exists\nval, exists := prices[\"meatball\"]\nif !exists {\n    fmt.Println(\"Meatball not on the menu\")\n}\n```\n\n## Iterating Maps\n```go\nfor item, p := range prices {\n    fmt.Printf(\"%s: Rp%d\\n\", item, p)\n}\n```\n\nNote: map iteration order in Go is random."),
		},
		{ID: lessonUUID(1, 10), CourseID: CourseGoBeg, OrderIndex: 10, XPReward: 20,
			TitleID: "Struct & Method", TitleEN: sp("Structs & Methods"),
			ContentMarkdownID: "# Struct & Method\n\n## Struct\nStruct adalah custom type yang mengelompokkan data:\n```go\ntype Mahasiswa struct {\n    Nama   string\n    NIM    string\n    IPK    float64\n}\n\nmhs := Mahasiswa{\n    Nama: \"Budi\",\n    NIM:  \"2024001\",\n    IPK:  3.75,\n}\nfmt.Println(mhs.Nama) // Budi\n```\n\n## Method\nMethod adalah fungsi yang terikat ke suatu type:\n```go\nfunc (m Mahasiswa) LulusKah() bool {\n    return m.IPK >= 2.0\n}\n\nfmt.Println(mhs.LulusKah()) // true\n```\n\n## Pointer Receiver\nUntuk mengubah nilai field:\n```go\nfunc (m *Mahasiswa) NaikkanIPK(delta float64) {\n    m.IPK += delta\n}\nmhs.NaikkanIPK(0.1)\n```\n\n## Struct Embedding\n```go\ntype Orang struct {\n    Nama string\n    Umur int\n}\ntype Karyawan struct {\n    Orang\n    Jabatan string\n}\nk := Karyawan{Orang: Orang{Nama: \"Ani\", Umur: 30}, Jabatan: \"Engineer\"}\nfmt.Println(k.Nama) // Ani (promoted field)\n```",
			ContentMarkdownEN: sp("# Structs & Methods\n\n## Struct\nA struct is a custom type that groups data:\n```go\ntype Student struct {\n    Name string\n    ID   string\n    GPA  float64\n}\n\nstudent := Student{\n    Name: \"Budi\",\n    ID:   \"2024001\",\n    GPA:  3.75,\n}\nfmt.Println(student.Name) // Budi\n```\n\n## Method\nA method is a function bound to a type:\n```go\nfunc (s Student) HasGraduated() bool {\n    return s.GPA >= 2.0\n}\n\nfmt.Println(student.HasGraduated()) // true\n```\n\n## Pointer Receiver\nTo modify field values:\n```go\nfunc (s *Student) BoostGPA(delta float64) {\n    s.GPA += delta\n}\nstudent.BoostGPA(0.1)\n```\n\n## Struct Embedding\n```go\ntype Person struct {\n    Name string\n    Age  int\n}\ntype Employee struct {\n    Person\n    Title string\n}\ne := Employee{Person: Person{Name: \"Ani\", Age: 30}, Title: \"Engineer\"}\nfmt.Println(e.Name) // Ani (promoted field)\n```"),
		},
		// ===== golang-intermediate (8 lessons) =====
		{ID: lessonUUID(2, 1), CourseID: CourseGoMid, OrderIndex: 1, XPReward: 20,
			TitleID: "Pointer", TitleEN: sp("Pointers"),
			ContentMarkdownID: "# Pointer\n\nPointer menyimpan alamat memori dari sebuah variabel.\n\n```go\nx := 42\np := &x        // p adalah pointer ke x\nfmt.Println(*p) // 42 (dereferencing)\n*p = 100\nfmt.Println(x)  // 100 (x berubah!)\n```\n\n## Kenapa Pointer Penting?\n1. **Efisiensi** — passing struct besar tanpa copy\n2. **Mutasi** — mengubah value dari fungsi lain\n3. **Nil safety** — pointer bisa nil, harus dicek\n\n```go\nfunc ubahNilai(n *int) {\n    *n = *n * 2\n}\nangka := 5\nubahNilai(&angka)\nfmt.Println(angka) // 10\n```",
			ContentMarkdownEN: sp("# Pointers\n\nA pointer stores the memory address of a variable.\n\n```go\nx := 42\np := &x        // p is a pointer to x\nfmt.Println(*p) // 42 (dereferencing)\n*p = 100\nfmt.Println(x)  // 100 (x changed!)\n```\n\n## Why Are Pointers Important?\n1. **Efficiency** — passing large structs without copying\n2. **Mutation** — changing values from other functions\n3. **Nil safety** — pointers can be nil, must be checked\n\n```go\nfunc doubleValue(n *int) {\n    *n = *n * 2\n}\nnum := 5\ndoubleValue(&num)\nfmt.Println(num) // 10\n```"),
		},
		{ID: lessonUUID(2, 2), CourseID: CourseGoMid, OrderIndex: 2, XPReward: 25,
			TitleID: "Interface", TitleEN: sp("Interfaces"),
			ContentMarkdownID: "# Interface\n\nInterface mendefinisikan kontrak behavior tanpa implementasi.\n\n```go\ntype Hewan interface {\n    Suara() string\n}\n\ntype Kucing struct{}\nfunc (k Kucing) Suara() string { return \"Meong\" }\n\ntype Anjing struct{}\nfunc (a Anjing) Suara() string { return \"Guk\" }\n\nfunc cetakSuara(h Hewan) {\n    fmt.Println(h.Suara())\n}\n\ncetakSuara(Kucing{}) // Meong\ncetakSuara(Anjing{}) // Guk\n```\n\n## Empty Interface\n`interface{}` menerima tipe apapun (mirip `any` di Go 1.18+):\n```go\nfunc cetak(v interface{}) {\n    fmt.Println(v)\n}\n```\n\n## Type Assertion\n```go\nvar i interface{} = \"hello\"\ns, ok := i.(string)\nif ok {\n    fmt.Println(s) // hello\n}\n```",
			ContentMarkdownEN: sp("# Interfaces\n\nInterfaces define behavioral contracts without implementation.\n\n```go\ntype Animal interface {\n    Sound() string\n}\n\ntype Cat struct{}\nfunc (c Cat) Sound() string { return \"Meow\" }\n\ntype Dog struct{}\nfunc (d Dog) Sound() string { return \"Woof\" }\n\nfunc printSound(a Animal) {\n    fmt.Println(a.Sound())\n}\n\nprintSound(Cat{}) // Meow\nprintSound(Dog{}) // Woof\n```\n\n## Empty Interface\n`interface{}` accepts any type (same as `any` in Go 1.18+):\n```go\nfunc print(v interface{}) {\n    fmt.Println(v)\n}\n```\n\n## Type Assertion\n```go\nvar i interface{} = \"hello\"\ns, ok := i.(string)\nif ok {\n    fmt.Println(s) // hello\n}\n```"),
		},
		{ID: lessonUUID(2, 3), CourseID: CourseGoMid, OrderIndex: 3, XPReward: 25, TitleID: "Error Handling", TitleEN: sp("Error Handling"),
			ContentMarkdownID: "# Error Handling\n\nGo menggunakan pattern return error, bukan exception.\n\n```go\nfunc bagi(a, b float64) (float64, error) {\n    if b == 0 {\n        return 0, errors.New(\"pembagian dengan nol\")\n    }\n    return a / b, nil\n}\n\nhasil, err := bagi(10, 0)\nif err != nil {\n    log.Fatal(err)\n}\n```\n\n## Custom Error\n```go\ntype ValidationError struct {\n    Field   string\n    Message string\n}\n\nfunc (e *ValidationError) Error() string {\n    return fmt.Sprintf(\"%s: %s\", e.Field, e.Message)\n}\n```\n\n## errors.Is dan errors.As\n```go\nif errors.Is(err, os.ErrNotExist) {\n    fmt.Println(\"File tidak ditemukan\")\n}\n```",
			ContentMarkdownEN: sp("# Error Handling\n\nGo uses a return error pattern, not exceptions.\n\n```go\nfunc divide(a, b float64) (float64, error) {\n    if b == 0 {\n        return 0, errors.New(\"division by zero\")\n    }\n    return a / b, nil\n}\n\nresult, err := divide(10, 0)\nif err != nil {\n    log.Fatal(err)\n}\n```\n\n## Custom Error\n```go\ntype ValidationError struct {\n    Field   string\n    Message string\n}\n\nfunc (e *ValidationError) Error() string {\n    return fmt.Sprintf(\"%s: %s\", e.Field, e.Message)\n}\n```\n\n## errors.Is and errors.As\n```go\nif errors.Is(err, os.ErrNotExist) {\n    fmt.Println(\"File not found\")\n}\n```"),
		},
		{ID: lessonUUID(2, 4), CourseID: CourseGoMid, OrderIndex: 4, XPReward: 30, TitleID: "Goroutine & WaitGroup", TitleEN: sp("Goroutines & WaitGroup"),
			ContentMarkdownID: "# Goroutine & WaitGroup\n\nGoroutine adalah lightweight thread yang dikelola Go runtime.\n\n```go\ngo func() {\n    fmt.Println(\"Berjalan paralel!\")\n}()\n```\n\n## WaitGroup\nUntuk menunggu goroutine selesai:\n```go\nvar wg sync.WaitGroup\nfor i := 0; i < 3; i++ {\n    wg.Add(1)\n    go func(n int) {\n        defer wg.Done()\n        fmt.Println(\"Worker\", n)\n    }(i)\n}\nwg.Wait()\n```",
			ContentMarkdownEN: sp("# Goroutines & WaitGroup\n\nGoroutines are lightweight threads managed by the Go runtime.\n\n```go\ngo func() {\n    fmt.Println(\"Running in parallel!\")\n}()\n```\n\n## WaitGroup\nTo wait for goroutines to complete:\n```go\nvar wg sync.WaitGroup\nfor i := 0; i < 3; i++ {\n    wg.Add(1)\n    go func(n int) {\n        defer wg.Done()\n        fmt.Println(\"Worker\", n)\n    }(i)\n}\nwg.Wait()\n```"),
		},
		{ID: lessonUUID(2, 5), CourseID: CourseGoMid, OrderIndex: 5, XPReward: 30, TitleID: "Channel & select", TitleEN: sp("Channels & select"),
			ContentMarkdownID: "# Channel & select\n\nChannel adalah cara komunikasi antar goroutine.\n\n```go\nch := make(chan string)\ngo func() { ch <- \"halo dari goroutine\" }()\nmsg := <-ch\nfmt.Println(msg)\n```\n\n## Buffered Channel\n```go\nch := make(chan int, 3) // buffer 3\nch <- 1\nch <- 2\nfmt.Println(<-ch) // 1\n```\n\n## Select\n```go\nselect {\ncase msg := <-ch1:\n    fmt.Println(\"dari ch1:\", msg)\ncase msg := <-ch2:\n    fmt.Println(\"dari ch2:\", msg)\ncase <-time.After(time.Second):\n    fmt.Println(\"timeout!\")\n}\n```",
			ContentMarkdownEN: sp("# Channels & select\n\nChannels are the way goroutines communicate.\n\n```go\nch := make(chan string)\ngo func() { ch <- \"hello from goroutine\" }()\nmsg := <-ch\nfmt.Println(msg)\n```\n\n## Buffered Channel\n```go\nch := make(chan int, 3) // buffer of 3\nch <- 1\nch <- 2\nfmt.Println(<-ch) // 1\n```\n\n## Select\n```go\nselect {\ncase msg := <-ch1:\n    fmt.Println(\"from ch1:\", msg)\ncase msg := <-ch2:\n    fmt.Println(\"from ch2:\", msg)\ncase <-time.After(time.Second):\n    fmt.Println(\"timeout!\")\n}\n```"),
		},
		{ID: lessonUUID(2, 6), CourseID: CourseGoMid, OrderIndex: 6, XPReward: 25, TitleID: "Context", TitleEN: sp("Context"),
			ContentMarkdownID: "# Context\n\nPackage `context` digunakan untuk mengatur deadline, cancellation, dan value sharing antar goroutine.\n\n```go\nctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)\ndefer cancel()\n\nselect {\ncase <-time.After(3 * time.Second):\n    fmt.Println(\"Selesai\")\ncase <-ctx.Done():\n    fmt.Println(\"Timeout:\", ctx.Err())\n}\n```\n\n## Context in HTTP\n```go\nfunc handler(w http.ResponseWriter, r *http.Request) {\n    ctx := r.Context()\n    select {\n    case <-time.After(5 * time.Second):\n        fmt.Fprintln(w, \"done\")\n    case <-ctx.Done():\n        return // client disconnected\n    }\n}\n```",
			ContentMarkdownEN: sp("# Context\n\nThe `context` package is used to manage deadlines, cancellation, and value sharing between goroutines.\n\n```go\nctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)\ndefer cancel()\n\nselect {\ncase <-time.After(3 * time.Second):\n    fmt.Println(\"Done\")\ncase <-ctx.Done():\n    fmt.Println(\"Timeout:\", ctx.Err())\n}\n```\n\n## Context in HTTP\n```go\nfunc handler(w http.ResponseWriter, r *http.Request) {\n    ctx := r.Context()\n    select {\n    case <-time.After(5 * time.Second):\n        fmt.Fprintln(w, \"done\")\n    case <-ctx.Done():\n        return // client disconnected\n    }\n}\n```"),
		},
		{ID: lessonUUID(2, 7), CourseID: CourseGoMid, OrderIndex: 7, XPReward: 20, TitleID: "Package & Module", TitleEN: sp("Packages & Modules"),
			ContentMarkdownID: "# Package & Module\n\n## Package\nSetiap folder di Go adalah satu package.\n```\nmyapp/\n  main.go        (package main)\n  utils/\n    helper.go    (package utils)\n```\n\n## Go Modules\nInisialisasi project:\n```bash\ngo mod init github.com/user/myapp\n```\n\nTambah dependency:\n```bash\ngo get github.com/gofiber/fiber/v2\n```\n\n## Exported vs Unexported\nHuruf kapital = exported (bisa diakses dari luar package):\n```go\n// exported\nfunc Hitung() int { ... }\n// unexported\nfunc hitungInternal() int { ... }\n```",
			ContentMarkdownEN: sp("# Packages & Modules\n\n## Package\nEvery folder in Go is one package.\n```\nmyapp/\n  main.go        (package main)\n  utils/\n    helper.go    (package utils)\n```\n\n## Go Modules\nInitialize a project:\n```bash\ngo mod init github.com/user/myapp\n```\n\nAdd a dependency:\n```bash\ngo get github.com/gofiber/fiber/v2\n```\n\n## Exported vs Unexported\nCapitalized = exported (accessible from outside package):\n```go\n// exported\nfunc Calculate() int { ... }\n// unexported\nfunc calculateInternal() int { ... }\n```"),
		},
		{ID: lessonUUID(2, 8), CourseID: CourseGoMid, OrderIndex: 8, XPReward: 25, TitleID: "Testing", TitleEN: sp("Testing"),
			ContentMarkdownID: "# Testing di Go\n\nGo punya testing framework built-in.\n\n## File Test\nFile test harus berakhiran `_test.go`:\n```go\n// math.go\nfunc Add(a, b int) int { return a + b }\n\n// math_test.go\nfunc TestAdd(t *testing.T) {\n    result := Add(2, 3)\n    if result != 5 {\n        t.Errorf(\"Expected 5, got %d\", result)\n    }\n}\n```\n\nJalankan: `go test ./...`\n\n## Table-Driven Tests\n```go\nfunc TestAdd(t *testing.T) {\n    tests := []struct{ a, b, want int }{\n        {1, 2, 3},\n        {0, 0, 0},\n        {-1, 1, 0},\n    }\n    for _, tc := range tests {\n        got := Add(tc.a, tc.b)\n        if got != tc.want {\n            t.Errorf(\"Add(%d,%d) = %d, want %d\", tc.a, tc.b, got, tc.want)\n        }\n    }\n}\n```",
			ContentMarkdownEN: sp("# Testing in Go\n\nGo has a built-in testing framework.\n\n## Test File\nTest files must end with `_test.go`:\n```go\n// math.go\nfunc Add(a, b int) int { return a + b }\n\n// math_test.go\nfunc TestAdd(t *testing.T) {\n    result := Add(2, 3)\n    if result != 5 {\n        t.Errorf(\"Expected 5, got %d\", result)\n    }\n}\n```\n\nRun: `go test ./...`\n\n## Table-Driven Tests\n```go\nfunc TestAdd(t *testing.T) {\n    tests := []struct{ a, b, want int }{\n        {1, 2, 3},\n        {0, 0, 0},\n        {-1, 1, 0},\n    }\n    for _, tc := range tests {\n        got := Add(tc.a, tc.b)\n        if got != tc.want {\n            t.Errorf(\"Add(%d,%d) = %d, want %d\", tc.a, tc.b, got, tc.want)\n        }\n    }\n}\n```"),
		},
		// ===== golang-advanced (8 lessons) — shorter content =====
		{ID: lessonUUID(3, 1), CourseID: CourseGoAdv, OrderIndex: 1, XPReward: 35, TitleID: "Generics", TitleEN: sp("Generics"),
			ContentMarkdownID: "# Generics (Go 1.18+)\n\nGenerics memungkinkan fungsi dan struct bekerja dengan berbagai tipe data.\n\n```go\nfunc Min[T int | float64 | string](a, b T) T {\n    if a < b { return a }\n    return b\n}\n\nfmt.Println(Min(3, 7))       // 3\nfmt.Println(Min(3.14, 2.71)) // 2.71\n```\n\n## Type Constraints\n```go\ntype Number interface {\n    int | int64 | float64\n}\n\nfunc Sum[T Number](items []T) T {\n    var total T\n    for _, v := range items {\n        total += v\n    }\n    return total\n}\n```",
			ContentMarkdownEN: sp("# Generics (Go 1.18+)\n\nGenerics allow functions and structs to work with multiple data types.\n\n```go\nfunc Min[T int | float64 | string](a, b T) T {\n    if a < b { return a }\n    return b\n}\n\nfmt.Println(Min(3, 7))       // 3\nfmt.Println(Min(3.14, 2.71)) // 2.71\n```\n\n## Type Constraints\n```go\ntype Number interface {\n    int | int64 | float64\n}\n\nfunc Sum[T Number](items []T) T {\n    var total T\n    for _, v := range items {\n        total += v\n    }\n    return total\n}\n```"),
		},
		{ID: lessonUUID(3, 2), CourseID: CourseGoAdv, OrderIndex: 2, XPReward: 35, TitleID: "Reflection", TitleEN: sp("Reflection"),
			ContentMarkdownID: "# Reflection\n\nPackage `reflect` memungkinkan inspeksi tipe dan nilai pada runtime.\n\n```go\nv := reflect.ValueOf(42)\nfmt.Println(v.Type())  // int\nfmt.Println(v.Kind())  // int\nfmt.Println(v.Int())   // 42\n```\n\n## Inspect Struct Fields\n```go\ntype User struct {\n    Name string `json:\"name\"`\n    Age  int    `json:\"age\"`\n}\nt := reflect.TypeOf(User{})\nfor i := 0; i < t.NumField(); i++ {\n    f := t.Field(i)\n    fmt.Printf(\"%s (%s) tag=%s\\n\", f.Name, f.Type, f.Tag.Get(\"json\"))\n}\n```\n\nReflection powerful tapi lambat — gunakan hanya jika perlu.",
			ContentMarkdownEN: sp("# Reflection\n\nThe `reflect` package allows runtime type and value inspection.\n\n```go\nv := reflect.ValueOf(42)\nfmt.Println(v.Type())  // int\nfmt.Println(v.Kind())  // int\nfmt.Println(v.Int())   // 42\n```\n\n## Inspect Struct Fields\n```go\ntype User struct {\n    Name string `json:\"name\"`\n    Age  int    `json:\"age\"`\n}\nt := reflect.TypeOf(User{})\nfor i := 0; i < t.NumField(); i++ {\n    f := t.Field(i)\n    fmt.Printf(\"%s (%s) tag=%s\\n\", f.Name, f.Type, f.Tag.Get(\"json\"))\n}\n```\n\nReflection is powerful but slow — use only when necessary."),
		},
		{ID: lessonUUID(3, 3), CourseID: CourseGoAdv, OrderIndex: 3, XPReward: 35, TitleID: "Design Patterns di Go", TitleEN: sp("Design Patterns in Go"),
			ContentMarkdownID: "# Design Patterns di Go\n\n## Singleton\n```go\nvar once sync.Once\nvar instance *Config\n\nfunc GetConfig() *Config {\n    once.Do(func() { instance = &Config{} })\n    return instance\n}\n```\n\n## Factory\n```go\nfunc NewPayment(method string) Payment {\n    switch method {\n    case \"credit\": return &CreditPayment{}\n    case \"bank\": return &BankPayment{}\n    default: return &CashPayment{}\n    }\n}\n```\n\n## Strategy\nBiasanya diimplementasi lewat interface di Go.\n\n## Repository Pattern\nPattern ini sering dipakai untuk data access layer:\n```go\ntype UserRepo interface {\n    FindByID(id uuid.UUID) (*User, error)\n    Create(user *User) error\n}\n```",
			ContentMarkdownEN: sp("# Design Patterns in Go\n\n## Singleton\n```go\nvar once sync.Once\nvar instance *Config\n\nfunc GetConfig() *Config {\n    once.Do(func() { instance = &Config{} })\n    return instance\n}\n```\n\n## Factory\n```go\nfunc NewPayment(method string) Payment {\n    switch method {\n    case \"credit\": return &CreditPayment{}\n    case \"bank\": return &BankPayment{}\n    default: return &CashPayment{}\n    }\n}\n```\n\n## Strategy\nTypically implemented via interfaces in Go.\n\n## Repository Pattern\nCommonly used for the data access layer:\n```go\ntype UserRepo interface {\n    FindByID(id uuid.UUID) (*User, error)\n    Create(user *User) error\n}\n```"),
		},
		{ID: lessonUUID(3, 4), CourseID: CourseGoAdv, OrderIndex: 4, XPReward: 40, TitleID: "HTTP & REST API dengan Fiber", TitleEN: sp("HTTP & REST API with Fiber"),
			ContentMarkdownID: "# HTTP & REST API dengan Fiber\n\nFiber adalah web framework Go yang terinspirasi Express.js.\n\n```go\npackage main\nimport \"github.com/gofiber/fiber/v2\"\n\nfunc main() {\n    app := fiber.New()\n    app.Get(\"/api/hello\", func(c *fiber.Ctx) error {\n        return c.JSON(fiber.Map{\"message\": \"Hello!\"})\n    })\n    app.Listen(\":3000\")\n}\n```\n\n## Route Parameters & Body Parsing\n```go\napp.Get(\"/users/:id\", func(c *fiber.Ctx) error {\n    id := c.Params(\"id\")\n    return c.JSON(fiber.Map{\"id\": id})\n})\n\napp.Post(\"/users\", func(c *fiber.Ctx) error {\n    var body CreateUserDTO\n    if err := c.BodyParser(&body); err != nil {\n        return c.Status(400).JSON(fiber.Map{\"error\": err.Error()})\n    }\n    return c.Status(201).JSON(body)\n})\n```",
			ContentMarkdownEN: sp("# HTTP & REST API with Fiber\n\nFiber is a Go web framework inspired by Express.js.\n\n```go\npackage main\nimport \"github.com/gofiber/fiber/v2\"\n\nfunc main() {\n    app := fiber.New()\n    app.Get(\"/api/hello\", func(c *fiber.Ctx) error {\n        return c.JSON(fiber.Map{\"message\": \"Hello!\"})\n    })\n    app.Listen(\":3000\")\n}\n```\n\n## Route Parameters & Body Parsing\n```go\napp.Get(\"/users/:id\", func(c *fiber.Ctx) error {\n    id := c.Params(\"id\")\n    return c.JSON(fiber.Map{\"id\": id})\n})\n\napp.Post(\"/users\", func(c *fiber.Ctx) error {\n    var body CreateUserDTO\n    if err := c.BodyParser(&body); err != nil {\n        return c.Status(400).JSON(fiber.Map{\"error\": err.Error()})\n    }\n    return c.Status(201).JSON(body)\n})\n```"),
		},
		{ID: lessonUUID(3, 5), CourseID: CourseGoAdv, OrderIndex: 5, XPReward: 40, TitleID: "Database dengan GORM", TitleEN: sp("Database with GORM"),
			ContentMarkdownID: "# Database dengan GORM\n\nGORM adalah ORM populer untuk Go.\n\n## Setup\n```go\nimport \"gorm.io/driver/postgres\"\nimport \"gorm.io/gorm\"\n\ndb, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})\ndb.AutoMigrate(&User{})\n```\n\n## CRUD\n```go\n// Create\ndb.Create(&User{Name: \"Budi\"})\n\n// Read\nvar user User\ndb.First(&user, \"name = ?\", \"Budi\")\n\n// Update\ndb.Model(&user).Update(\"name\", \"Budi Updated\")\n\n// Delete\ndb.Delete(&user)\n```\n\n## Relations & Preload\n```go\ndb.Preload(\"Orders\").Find(&users)\n```",
			ContentMarkdownEN: sp("# Database with GORM\n\nGORM is a popular ORM for Go.\n\n## Setup\n```go\nimport \"gorm.io/driver/postgres\"\nimport \"gorm.io/gorm\"\n\ndb, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})\ndb.AutoMigrate(&User{})\n```\n\n## CRUD\n```go\n// Create\ndb.Create(&User{Name: \"Budi\"})\n\n// Read\nvar user User\ndb.First(&user, \"name = ?\", \"Budi\")\n\n// Update\ndb.Model(&user).Update(\"name\", \"Budi Updated\")\n\n// Delete\ndb.Delete(&user)\n```\n\n## Relations & Preload\n```go\ndb.Preload(\"Orders\").Find(&users)\n```"),
		},
		{ID: lessonUUID(3, 6), CourseID: CourseGoAdv, OrderIndex: 6, XPReward: 30, TitleID: "CLI Application", TitleEN: sp("CLI Application"),
			ContentMarkdownID: "# CLI Application\n\nPackage `os` dan `flag` untuk membuat CLI.\n\n## os.Args\n```go\nargs := os.Args[1:]\nfmt.Println(\"Arguments:\", args)\n```\n\n## Flag Package\n```go\nname := flag.String(\"name\", \"World\", \"your name\")\nflag.Parse()\nfmt.Printf(\"Hello, %s!\\n\", *name)\n```\n\nJalankan: `go run main.go -name=Budi`\n\nUntuk CLI yang lebih kompleks, gunakan library seperti `cobra`.",
			ContentMarkdownEN: sp("# CLI Application\n\nThe `os` and `flag` packages for building CLIs.\n\n## os.Args\n```go\nargs := os.Args[1:]\nfmt.Println(\"Arguments:\", args)\n```\n\n## Flag Package\n```go\nname := flag.String(\"name\", \"World\", \"your name\")\nflag.Parse()\nfmt.Printf(\"Hello, %s!\\n\", *name)\n```\n\nRun: `go run main.go -name=Budi`\n\nFor more complex CLIs, use libraries like `cobra`."),
		},
		{ID: lessonUUID(3, 7), CourseID: CourseGoAdv, OrderIndex: 7, XPReward: 35, TitleID: "Performance & Profiling", TitleEN: sp("Performance & Profiling"),
			ContentMarkdownID: "# Performance & Profiling\n\n## Benchmarks\n```go\nfunc BenchmarkFibo(b *testing.B) {\n    for i := 0; i < b.N; i++ {\n        Fibonacci(20)\n    }\n}\n```\n\nJalankan: `go test -bench=. -benchmem`\n\n## pprof\n```go\nimport _ \"net/http/pprof\"\ngo func() { http.ListenAndServe(\":6060\", nil) }()\n```\n\nAkses profil: `go tool pprof http://localhost:6060/debug/pprof/heap`\n\n## Tips Performa\n1. Gunakan `sync.Pool` untuk reduce GC pressure\n2. Pre-allocate slice dengan `make([]T, 0, expectedCap)`\n3. Hindari string concatenation dalam loop — pakai `strings.Builder`",
			ContentMarkdownEN: sp("# Performance & Profiling\n\n## Benchmarks\n```go\nfunc BenchmarkFibo(b *testing.B) {\n    for i := 0; i < b.N; i++ {\n        Fibonacci(20)\n    }\n}\n```\n\nRun: `go test -bench=. -benchmem`\n\n## pprof\n```go\nimport _ \"net/http/pprof\"\ngo func() { http.ListenAndServe(\":6060\", nil) }()\n```\n\nAccess profile: `go tool pprof http://localhost:6060/debug/pprof/heap`\n\n## Performance Tips\n1. Use `sync.Pool` to reduce GC pressure\n2. Pre-allocate slices with `make([]T, 0, expectedCap)`\n3. Avoid string concatenation in loops — use `strings.Builder`"),
		},
		{ID: lessonUUID(3, 8), CourseID: CourseGoAdv, OrderIndex: 8, XPReward: 40, TitleID: "Deployment: Docker & CI/CD", TitleEN: sp("Deployment: Docker & CI/CD"),
			ContentMarkdownID: "# Deployment: Docker & CI/CD\n\n## Dockerfile Multi-stage\n```dockerfile\nFROM golang:1.22-alpine AS builder\nWORKDIR /app\nCOPY go.mod go.sum ./\nRUN go mod download\nCOPY . .\nRUN CGO_ENABLED=0 go build -o server cmd/server/main.go\n\nFROM alpine:3.19\nCOPY --from=builder /app/server /server\nEXPOSE 8080\nCMD [\"/server\"]\n```\n\n## Docker Compose\n```yaml\nservices:\n  api:\n    build: .\n    ports: [\"8080:8080\"]\n    env_file: .env\n  db:\n    image: postgres:16-alpine\n    environment:\n      POSTGRES_DB: app\n```\n\n## GitHub Actions CI\n```yaml\non: push\njobs:\n  test:\n    runs-on: ubuntu-latest\n    steps:\n      - uses: actions/checkout@v4\n      - uses: actions/setup-go@v5\n      - run: go test ./...\n```",
			ContentMarkdownEN: sp("# Deployment: Docker & CI/CD\n\n## Multi-stage Dockerfile\n```dockerfile\nFROM golang:1.22-alpine AS builder\nWORKDIR /app\nCOPY go.mod go.sum ./\nRUN go mod download\nCOPY . .\nRUN CGO_ENABLED=0 go build -o server cmd/server/main.go\n\nFROM alpine:3.19\nCOPY --from=builder /app/server /server\nEXPOSE 8080\nCMD [\"/server\"]\n```\n\n## Docker Compose\n```yaml\nservices:\n  api:\n    build: .\n    ports: [\"8080:8080\"]\n    env_file: .env\n  db:\n    image: postgres:16-alpine\n    environment:\n      POSTGRES_DB: app\n```\n\n## GitHub Actions CI\n```yaml\non: push\njobs:\n  test:\n    runs-on: ubuntu-latest\n    steps:\n      - uses: actions/checkout@v4\n      - uses: actions/setup-go@v5\n      - run: go test ./...\n```"),
		},
	}

	for i := range lessons {
		upsertLesson(db, &lessons[i])
	}
}
