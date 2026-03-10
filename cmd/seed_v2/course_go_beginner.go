package main

import (
	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"gorm.io/gorm"
)

func seedGoBeginnerCourse(db *gorm.DB) {
	courseGoBeg := domain.Course{
		ID:            CourseGoBeg,
		Slug:          "go-beginner",
		Language:      "go",
		Level:         "beginner",
		TitleID:       "Golang untuk Pemula",
		TitleEN:       sp("Golang for Beginners"),
		DescriptionID: sp("Pelajari fundamental bahasa pemrograman Go (Golang) ciptaan Google. Simpel, sangat cepat, dan kuat untuk menangani backend modern skala besar. Mulai dari sintaks sintaks dasar hingga konkurensi goroutine!"),
		DescriptionEN: sp("Learn the fundamentals of Go (Golang) created by Google. Simple, blazing fast, and powerful for handling modern large-scale backends. From basic syntax to goroutine concurrency!"),
		ThumbnailURL:  sp("https://storage.googleapis.com/ngodingyuk-assets/course/go-beg.png"),
		OrderIndex:    20,
	}
	upsertCourse(db, &courseGoBeg)

	mod1 := moduleUUID(5, 1)
	mod2 := moduleUUID(5, 2)
	mod3 := moduleUUID(5, 3)
	mod4 := moduleUUID(5, 4)
	mod5 := moduleUUID(5, 5)

	modules := []domain.Module{
		{ID: mod1, CourseID: CourseGoBeg, TitleID: "Pengenalan & Dasar Sintaks", TitleEN: sp("Introduction & Basic Syntax"), OrderIndex: 1},
		{ID: mod2, CourseID: CourseGoBeg, TitleID: "Struktur Data", TitleEN: sp("Data Structures"), OrderIndex: 2},
		{ID: mod3, CourseID: CourseGoBeg, TitleID: "Fungsi & Pointer", TitleEN: sp("Functions & Pointers"), OrderIndex: 3},
		{ID: mod4, CourseID: CourseGoBeg, TitleID: "Struct & Interface", TitleEN: sp("Structs & Interfaces"), OrderIndex: 4},
		{ID: mod5, CourseID: CourseGoBeg, TitleID: "Topik Lanjutan & Ekosistem", TitleEN: sp("Advanced Topics & Ecosystem"), OrderIndex: 5},
	}

	for _, mod := range modules {
		upsertModule(db, &mod)
	}

	lessons := []domain.Lesson{
		{
			ID:         lessonUUID(5, 1),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod1,
			OrderIndex: 1,
			XPReward:   10,
			TitleID:    "1. Pengenalan Go (Golang)",
			TitleEN:    sp("1. Introduction to Go (Golang)"),
			ContentMarkdownID: `### Apa itu Go / Golang?
Go adalah bahasa pemrograman *open source* statis dan dikompilasi (compiled) yang diciptakan di Google oleh Robert Griesemer, Rob Pike, dan Ken Thompson pada tahun 2009.

**Mengapa belajar Go?**
- **Sangat Cepat**: Dikompilasi langsung menjadi kode mesin yang dieksekusi CPU tanpa "virtual machine" perantara berat.
- **Sederhana**: Sintaks minimalis, terinspirasi dari C namun sangat mudah dibaca layaknya Python.
- **Konkurensi Mutakhir (Goroutine)**: Secara bawaan dirancang untuk menahan beban jutaan proses lalu lintas jaringan bersamaan (*Goroutines* & *Channels*).
- **Static Typing**: Keamanan tipe data terjaga layaknya Java dan C++.

Banyak perusahaan besar seperti Uber, Twitch, Grab, Tokopedia, dan Google menggunakannya untuk menopang *high-performance backend api*.`,
			ContentMarkdownEN: sp(`### What is Go / Golang?
Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson back in 2009.

**Why learn Go?**
- **Blazing Fast**: Compiled directly to machine code executed by CPU devoid of heavy intermediate virtual machines.
- **Simplicity**: Minimalistic syntax, derived from C but highly readable acting similar to Python.
- **Modern Concurrency**: Natively designed to sustain millions of parallel concurrent network traffic loads (*Goroutines* & *Channels*).
- **Static Typing**: Data type safety is well maintained analogous to Java and C++.

Tech conglomerates like Uber, Twitch, Grab, Tokopedia, and Google utilize it globally to propel *high-performance backend API*.`),
		},
		{
			ID:         lessonUUID(5, 2),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod1,
			OrderIndex: 2,
			XPReward:   10,
			TitleID:    "2. Struktur Hello World",
			TitleEN:    sp("2. Hello World Structure"),
			ContentMarkdownID: `### Anatomi Program Go
Mari kita lihat program eksekusi Go paling mendasar:

` + "```" + `go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
` + "```" + `

**Penjelasan:**
1. ` + "`" + `package main` + "`" + `: Mendeklarasikan paket wajib (` + "`" + `main` + "`" + `) yang mendefinisikan bahwa program ini adalah aplikasi yang dapat dieksekusi (executable).
2. ` + "`" + `import "fmt"` + "`" + `: Memasukkan *package* bawaan ` + "`" + `fmt` + "`" + ` (Format) yang menyimpan fungsi pencetakan teks.
3. ` + "`" + `func main() {}` + "`" + `: Titik masuk eksekusi utama (entry point) dari aplikasi Go manapun. Program dimulai dan diakhiri dari sini.
4. ` + "`" + `fmt.Println(...)` + "`" + `: Memerintahkan komputer mem-print pesan di Terminal.`,
			ContentMarkdownEN: sp(`### Go Program Anatomy
Let's see the most fundamental Go execution layout:

` + "```" + `go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
` + "```" + `

**Explanations:**
1. ` + "`" + `package main` + "`" + `: Declares the compulsory package name. Utilizing ` + "`" + `main` + "`" + ` dictates that this script equates to a standalone executable bin.
2. ` + "`" + `import "fmt"` + "`" + `: Injects the built-in standard formatter package (` + "`" + `fmt` + "`" + `) providing console textual print functions.
3. ` + "`" + `func main() {}` + "`" + `: Central execution entry point for any conventional Go application. Operations boot up and wind down intrinsically here.
4. ` + "`" + `fmt.Println(...)` + "`" + `: Commands the computer device to log/print messages upon the console.`),
		},
		{
			ID:         lessonUUID(5, 3),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod1,
			OrderIndex: 3,
			XPReward:   10,
			TitleID:    "3. Variabel dan Konstanta",
			TitleEN:    sp("3. Variables and Constants"),
			ContentMarkdownID: `### Membuat Variabel
Variabel digunakan untuk menyimpan kepingan data. Di Go, kamu bisa menetapkannya dengan cara formal atau singkat (short variable declaration).

` + "```" + `go
// Deklarasi formal (Menuliskan tipe data)
var heroName string = "Gatotkaca"
var hp int = 150

// Deklarasi Singkat (Tipe data ditebak otomatis oleh Go)
power := 9000
mana := 50.5
` + "```" + `
*Catatan: Deklarasi singkat bertanda ` + "`" + `:=` + "`" + ` hanya dapat digunakan di DALAM body function.*

### Konstanta (Constant)
Digunakan untuk nilai mutlak yang pantang berubah atau tidak akan diolah ulang semenjak kodingan dikompilasi.
` + "```" + `go
const pi = 3.14
const appName = "NgodingYuk"
` + "```" + ``,
			ContentMarkdownEN: sp(`### Creating Variables
Variables allocate storage container states. In Go, you dictate them utilizing long formalities or shorthand notations.

` + "```" + `go
// Explicit Formal Declarations (Denoting explicit types)
var heroName string = "Gatotkaca"
var hp int = 150

// Short Declarations (The Go compiler implicitly delegates semantic datatypes)
power := 9000
mana := 50.5
` + "```" + `
*Note: Shorthand ` + "`" + `:=` + "`" + ` syntax relies solely INSIDE inner functions boundaries.*

### Constants
Represent invariant absolutes intended never to shift their states throughout runtime lifecycle spans.
` + "```" + `go
const pi = 3.14
const appName = "NgodingYuk"
` + "```" + ``),
		},
		{
			ID:         lessonUUID(5, 4),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod1,
			OrderIndex: 4,
			XPReward:   10,
			TitleID:    "4. Tipe Data Dasar",
			TitleEN:    sp("4. Basic Data Types"),
			ContentMarkdownID: `### Kelompok Tipe Data
Keamanan komputasi di Go diawasi super ketat. Data harus konsisten.

1. **Integer (Bilangan Bulat)**
   Mencakup angka minus/positif utuh. (Contoh: ` + "`" + `int` + "`" + `, ` + "`" + `int8` + "`" + `, ` + "`" + `int64` + "`" + `). Semakin tinggi angkanya, cakupan nilai maksimalnya semakin lebar.
2. **Float (Desimal Pecahan)**
   Mencakup angka diakhiri titik (` + "`" + `float32` + "`" + `, ` + "`" + `float64` + "`" + `).
3. **String (Teks)**
   Harus dikurung dalam kutip ganda ` + "`" + `"Teks"` + "`" + ` atau backtick ` + "`" + `` + "`" + `Teks untuk banyak baris` + "`" + `` + "`" + `.
4. **Bool (Boolean)**
   Benar (` + "`" + `true` + "`" + `) bernilai 1, Salah (` + "`" + `false` + "`" + `) bernilai 0 mutlak di logika perbandingan.

` + "```" + `go
var umur int = 25
var gaji float64 = 4500.50
var nama string = "Arul"
var kerja bool = true
` + "```" + ``,
			ContentMarkdownEN: sp(`### Primitive Datatypes Mapping
Go dictates highly rigid computational safety. Data persistence must align tightly.

1. **Integers (Whole Numbers)**
   Sustains literal negatives/positives mathematically. (e.g. ` + "`" + `int` + "`" + `, ` + "`" + `int8` + "`" + `, ` + "`" + `int64` + "`" + `). The digit expands bit capacity sizes.
2. **Floats (Decimal Fractions)**
   Floating decimal coordinates bounded internally (` + "`" + `float32` + "`" + `, ` + "`" + `float64` + "`" + `).
3. **Strings (Characters)**
   Must be encapsulated thoroughly using double-quotes ` + "`" + `"Lexicon"` + "`" + ` or multi-lined backticks.
4. **Bool (Booleans)**
   Rightful conditionals evaluating binary dimensions. Exact true/false parameters exclusively.

` + "```" + `go
var age int = 25
var salary float64 = 4500.50
var name string = "Arul"
var employment bool = true
` + "```" + ``),
		},
		{
			ID:         lessonUUID(5, 5),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod1,
			OrderIndex: 5,
			XPReward:   10,
			TitleID:    "5. Operator & Logika",
			TitleEN:    sp("5. Operators & Logics"),
			ContentMarkdownID: `### Operator Aritmatika
Operator matematika dasar layaknya kalkulator wajar:
- Penjumlahan (` + "`" + `+` + "`" + `)
- Pengurangan (` + "`" + `-` + "`" + `)
- Perkalian (` + "`" + `*` + "`" + `)
- Pembagian (` + "`" + `/` + "`" + `) (Khusus ` + "`" + `int` + "`" + `, sisa pecahan akan *dihapus habis*/dibulatkan ke bawah)
- Modulo (` + "`" + `%` + "`" + `) (Mengambil sisa bagi)

Secara perbandingan kondisional boolean, kalian mempunyai ` + "`" + `==` + "`" + ` (sama), ` + "`" + `!=` + "`" + ` (berbeda), ` + "`" + `>` + "`" + ` (lebih besar).

### Operator Logika AND/OR
Di Golang, ` + "`" + `&&` + "`" + ` adalah (AND / Dan). Semuanya harus TRUE.
Sedangkan ` + "`" + `||` + "`" + ` adalah (OR / Atau). Cukup salah satu kondisi yang TRUE maka total logikanya TRUE.`,
			ContentMarkdownEN: sp(`### Arithmetic Operators
Core literal logical evaluation similar to generic calculator grids:
- Addition (` + "`" + `+` + "`" + `)
- Subtraction (` + "`" + `-` + "`" + `)
- Multiplication (` + "`" + `*` + "`" + `)
- Division (` + "`" + `/` + "`" + `) (Strictly for ` + "`" + `int` + "`" + `, fractions trailing comma yields empty void truncated).
- Modulo (` + "`" + `%` + "`" + `) (Extrapolating remainder residues).

For boolean comparative analysis, usage commands rely onto ` + "`" + `==` + "`" + ` (equals), ` + "`" + `!=` + "`" + ` (unequals), ` + "`" + `>` + "`" + ` (magnitudes).

### Logical Operator Chaining
In Go mappings, ` + "`" + `&&` + "`" + ` portrays exact (AND). Entire grids assert to TRUE.
Subsequently ` + "`" + `||` + "`" + ` relays contextual (OR) parameters. Single isolated truth validates totality unconditionally.`),
		},
		{
			ID:         lessonUUID(5, 6),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod1,
			OrderIndex: 6,
			XPReward:   10,
			TitleID:    "6. Percabangan If-Else",
			TitleEN:    sp("6. Conditional If-Else"),
			ContentMarkdownID: `### Menyalurkan Arah Logika Kode
Go menolak penggunaan kurung bulat pada syarat percabangan (seperti Java). Kode jadi lebih mulus.

` + "```" + `go
nilai := 80

if nilai >= 90 {
    fmt.Println("Grade A")
} else if nilai >= 80 {
    fmt.Println("Grade B")
} else {
    fmt.Println("Drop out")
}
` + "```" + `

**If Bersarang (Nested):**
Kalian juga bisa mendeklarasikan variabel kecil yang usianya khusus terisolasi cuma hidup di dalam kurung kurawal statement if tesebut.

` + "```" + `go
// x disisipkan ke blok awal if secara aman
if x := 10; x > 5 {
    fmt.Println(x) // Berhasil, nyala dan berfungsi
}
// fmt.Println(x) // Error bila dipanggil sekarang!
` + "```" + ``,
			ContentMarkdownEN: sp(`### Routing Code Directions
Go rejects parenthetical groupings upon relational checks unlike Java counterparts ensuring pure streamlined readability.

` + "```" + `go
score := 80

if score >= 90 {
    fmt.Println("Grade A")
} else if score >= 80 {
    fmt.Println("Grade B")
} else {
    fmt.Println("Drop out")
}
` + "```" + `

**Nested Initializations Scope:**
Developers harness isolated inline initialization variables breathing solely bounded exclusively across current block perimeters.

` + "```" + `go
// x inserted dynamically
if x := 10; x > 5 {
    fmt.Println(x) // Triggerable properly
}
// fmt.Println(x) // Crash failures triggering outside boundaries
` + "```" + ``),
		},
		{
			ID:         lessonUUID(5, 7),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod2,
			OrderIndex: 7,
			XPReward:   10,
			TitleID:    "7. Perulangan For",
			TitleEN:    sp("7. For-Loops"),
			ContentMarkdownID: `### Satu-Satunya Keyword Perulangan!
Golang itu anti-ribet, ia tak mengenal bahasa *While* atau *Do-While*. Semua skenario iterasi cukup ditangani murni lewat 1 kata: ` + "`" + `for` + "`" + `!

**Gaya Numerik Klasik:**
` + "```" + `go
for i := 1; i <= 5; i++ {
    fmt.Println("Cetak ke-", i)
}
` + "```" + `

**Gaya Meniru *While*:**
` + "```" + `go
hitung := 1
for hitung < 10 {
    fmt.Println("Mengejar 10...")
    hitung++
}
` + "```" + `

**Gaya Infinity (Tak terbatas):**
` + "```" + `go
for {
    fmt.Println("Siklus tanpa akhir...")
    break // Untuk keluar manual dari neraka loop
}
` + "```" + ``,
			ContentMarkdownEN: sp(`### The Sovereign Loop Keyword!
Go eliminates clutter. It aggressively strips down *While* or *Do-While* concepts natively handling entire recursive iterational landscapes unified singularly by: ` + "`" + `for` + "`" + `!

**Classical Iteration Paradigm:**
` + "```" + `go
for i := 1; i <= 5; i++ {
    fmt.Println("Print log-", i)
}
` + "```" + `

**While-Mimickery Emulation:**
` + "```" + `go
count := 1
for count < 10 {
    fmt.Println("Chasing 10...")
    count++
}
` + "```" + `

**Abyssal Infinity Cycles:**
` + "```" + `go
for {
    fmt.Println("Unending abyss...")
    break // Force shatter the loop logic exit mechanisms
}
` + "```" + ``),
		},
		{
			ID:         lessonUUID(5, 8),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod2,
			OrderIndex: 8,
			XPReward:   10,
			TitleID:    "8. Array & Slice (Antrean Data)",
			TitleEN:    sp("8. Array & Slices (Data Queues)"),
			ContentMarkdownID: `### Array (Data Kaku Bertotal Statis)
Array adalah keranjang statis berisi tipe tipe serupa. Syarat Array Go: ukurannya absolut wajib didefinisikan semenjak kode dijalankan dan takkan bisa dikembangkan lagi kelak (Fixed Length).

` + "```" + `go
var names [3]string // Kapasitas 3 Data
names[0] = "A"
names[1] = "B"
names[2] = "C"
// names[3] = "D" // FATAL ERROR: Melebihi batas blok kaku!
` + "```" + `

### Slices (Data Cair Super Luwes!)
Dalam praktiknya, Programmer Go 95% memihak menggunakan Slices, yakni adaptasi melonggarkan Array yang dapat membesar (*append*) sesuka memori RAM kamu tanpa batas kaku deklarasi awal. Slice tidak diberi nomor blok total dari awal.

` + "```" + `go
// Perhatikan, kurung sikunya KOSONG [ ] tanpa angka
kawan := []string{"Budi", "Siti"}

// Tambah data baru sepuasnya secara dinamis!
kawan = append(kawan, "Anto", "Sarah")
fmt.Println(kawan) // [Budi, Siti, Anto, Sarah]
` + "```" + ``,
			ContentMarkdownEN: sp(`### Array (Rigid Fixed Cap Constraints)
Arrays act structurally forming homogeneous primitive blocks. A paramount mandate limits arrays internally demanding rigid length allocations preemptively before executions scale up.

` + "```" + `go
var names [3]string // Maximum bound sets up to 3 indices arrays
names[0] = "A"
names[1] = "B"
names[2] = "C"
// names[3] = "D" // FATAL PANIC: Matrix out of structural thresholds
` + "```" + `

### Slices (Limitless Fluid Expansion!)
Empirically, native Go engineering pivots 95% workflows utilizing Slices—dynamic wrappers surrounding base arrays projecting volatile capacity allocations capable of mutating boundlessly accommodating expanding application RAM heaps linearly via standard ` + "`" + `append()` + "`" + `. Notice inner brackets void of digits initializations.

` + "```" + `go
// Emtpy internal capacity declarations [ ] dictates slice generations
peers := []string{"Budi", "Siti"}

// Expand boundaries recursively efficiently!
peers = append(peers, "Anto", "Sarah")
fmt.Println(peers) // [Budi, Siti, Anto, Sarah]
` + "```" + ``),
		},
		{
			ID:         lessonUUID(5, 9),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod2,
			OrderIndex: 9,
			XPReward:   10,
			TitleID:    "9. Map (Peta Pasangan Data)",
			TitleEN:    sp("9. Map (Dictionary Pairings)"),
			ContentMarkdownID: `### Memetakan Key-Value Layaknya Kamus
Bila *Slice* memanggil anggotanya memakai index angka matematis ` + "`" + `data[0]` + "`" + `, struktur struktur **Map** bekerja melacak nilai layaknya kamus/absen mencari elemen menggunakan kode/kalimat *Key*.

Kalian mendirikan pondasi menggunakan syntax ` + "`" + `map[TIPE_KEY]TIPE_VALUE` + "`" + `.
Contoh kita ingin menyimpan Absen umur Siswa. Nama Teks (string) sebagai gemboknya, dan Umur Angka (int) sebagai nilainya.

` + "```" + `go
umurSiswa := make(map[string]int) // Pendirian objek pondasi map resmi

umurSiswa["Raffi"] = 28
umurSiswa["Nagita"] = 26

// Memanggil isi map Raffi
fmt.Println("Umurnya adalah", umurSiswa["Raffi"]) // 28

// Menghapus kunci entri spesifik seutuhnya
delete(umurSiswa, "Nagita") 
` + "```" + ``,
			ContentMarkdownEN: sp(`### Mapping Key-Value Schematics analogous to Dictionaries
Alternatively contrasted by *Slices* deploying numerical offsets grids fetching nodes traversing internally ` + "`" + `data[0]` + "`" + `, associative **Map** allocations navigate lookup endpoints strictly binding mapping identity semantic sequences acting identically evaluating as hash keys.

Architectures structure foundational maps composing ` + "`" + `map[KEY_TYPES]VALUE_TYPES` + "`" + ` matrices.
We forge lookup demographics linking student identifiers string signatures encapsulating target mathematical numeric payload nodes.

` + "```" + `go
studentAges := make(map[string]int) // Bootstraps associative map initializations 

studentAges["Raffi"] = 28
studentAges["Nagita"] = 26

// Extraction endpoint node retrievals
fmt.Println("Age evaluates to", studentAges["Raffi"]) // 28

// Demolishes dictionary target branches
delete(studentAges, "Nagita") 
` + "```" + ``),
		},
		{
			ID:         lessonUUID(5, 10),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod3,
			OrderIndex: 10,
			XPReward:   10,
			TitleID:    "10. Dasar Menulis Fungsi (Function)",
			TitleEN:    sp("10. Writing Fundamentals (Functions)"),
			ContentMarkdownID: `### Menyatukan Logika ke dalam Pemanggil
Bayangkan kalian menulis script pengecekan ganjil-genap rumit sebanyak 20 baris. Apakah jika dibutuhkan 3x di beda tempat harus *copy-paste* kode kotor itu 3x? Tidak! Bungkus kode panjang itu di dalam **Fungsi**.

**Konstruktor fungsi diawali keyword ` + "`" + `func` + "`" + `.**

` + "```" + `go
// Pembuatan Blok Fungsi Modular
// Nama fungsinya 'SayHello', menerima syarat 1 variabel parameter masuk 'namaLengkap'
func SayHello(namaLengkap string) {
    fmt.Println("Selamat bergabung di markas militer utama, Jendral", namaLengkap)
}

func main() {
    // Memanggil ulang fungsinya tanpa harus capek rekayasa sistem lagi
    SayHello("Antigravity")
    SayHello("Maximus")
}
` + "```" + ``,
			ContentMarkdownEN: sp(`### Aggregating Modular Operational Logics
Envision typing complex even-odd matrix parsers scaling 20 lines heavily. Should sequential duplicate usages duplicate manual paste footprints continuously violating strict *Don't Repeat Yourself (DRY)* tenets? No! Enclose vast operational pipelines securely forming robust **Functions**.

**Functional declarations explicitly prefix implementations via standard ` + "`" + `func` + "`" + ` namespaces.**

` + "```" + `go
// Modular Entity Fabrication Procedures
// Assigning functional names 'SayHello', capturing exactly 1 incoming parametric pipeline argument 'fullname' string semantics
func SayHello(fullName string) {
    fmt.Println("Welcome traversing the main strategic headquarters, General", fullName)
}

func main() {
    // Calling recurrent logical blueprints iteratively preventing systemic engineering overloads 
    SayHello("Antigravity")
    SayHello("Maximus")
}
` + "```" + ``),
		},
		{
			ID:         lessonUUID(5, 11),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod3,
			OrderIndex: 11,
			XPReward:   10,
			TitleID:    "11. Multiple Return Values (Kembalian Banyak)",
			TitleEN:    sp("11. Multiple Return Values"),
			ContentMarkdownID: `### Super Power Eksklusif Go
Satu kelebihan utama Go yang membuat Developer Python/C sakit hati iri adalah fleksibilitas **satu buah fungsi Go mampu membuahkan multi balasan kembalian secara terstruktur** tanpa mendirikan array objek keranjang kompleks. 

` + "```" + `go
// Menerima parameter sisi persegi panjang. 
// AJaibnya, dia memberikan BALASAN 2 data integer sekaligus (Luas & Keliling)!
func hitungPersegi(panjang int, lebar int) (int, int) {
    luas := panjang * lebar
    keliling := 2 * (panjang + lebar)
    
    // Melempar multi nilai
    return luas, keliling
}

func main() {
    // Menangkap lemparan ke dalam dua panitia penampung berbeda langsung otomatis
    L, K := hitungPersegi(10, 5)
    fmt.Printf("Ketangkap! Luas = %d, Keliling = %d\n", L, K)
    
    // Catatan: Jika kamu cuma butuh Luasnya dan tak peduli soal keliling, 
    // Go melarang keras adanya variabel yang tidak dipakai. Oleh karenanya gunakan garis bawah penolakan (Blank Identifier) "_"
    LL, _ := hitungPersegi(20, 5)
}
` + "```" + ``,
			ContentMarkdownEN: sp(`### Exclusive Go Functional Superpowers
A pivotal superiority standard Go environments exhibit causing rival framework developers monumental grievances resides encompassing flexible dynamic **single origin function capabilities emitting structural multiple native callback payloads independently** skipping traditional heavyweight wrapper instantiations.

` + "```" + `go
// Consuming grid geometric rect parameter matrices.
// Dynamically, yielding explicitly concurrent parallel payload exports collectively (Area & Perimeter outputs)!
func calculateRectangle(long int, wide int) (int, int) {
    area := long * wide
    circumference := 2 * (long + wide)
    
    // Dual node ejection syntax
    return area, circumference
}

func main() {
    // Direct capture mechanisms harvesting tuple matrices seamlessly automatically
    A, C := calculateRectangle(10, 5)
    fmt.Printf("Payload secured! Area = %d, Circumference = %d\n", A, C)
    
    // Note: Extrapolating target specific isolation omitting rest necessitates strict variable compliance formats.
    // Go compiler heavily punishes unused orphaned code paths. Employ underscores (Blank Identifiers) '_' bypassing errors securely.
    AA, _ := calculateRectangle(20, 5)
}
` + "```" + ``),
		},
		{
			ID:         lessonUUID(5, 12),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod3,
			OrderIndex: 12,
			XPReward:   10,
			TitleID:    "12. Pengenalan Tipografi Pointer & Alamat RAM",
			TitleEN:    sp("12. Typography Introductions Pointer & RAM Address"),
			ContentMarkdownID: `### Menunjuk Alamat Memori Asli Tembus Mesin (Pointer)
Bayangkan skenario ini: Kamu mengkopi *(copy-paste)* file dokumen ms-world laporan rekanmu menjadi "laporan-b.docx" di flashdisk. Kamu merubah total teksnya menjadi font Alien merah di flashdisk mu. Pertanyaannya: Apakah file laporan asli teman di laptop nya asalnya ikut berubah? Tentu tidak. Hal tersebut dinamakan **Pass By Value** (Di-kopi ganda ke bilik memori RAM palsu baru).

Kode Go secara standar mengkopi duplikat murni. Namun dengan Pointers (berlogo Asterisk ` + "`" + `*` + "`" + ` dan Ambersand ` + "`" + `&` + "`" + `), kamu dapat menembak langsung *Pass By Reference* mengubah file aslinya meloncati batas pemisah memori RAM.

- ` + "`" + `&var` + "`" + ` (Ampersand): Memberitahu alamat kode rahasia tempat memori bersemayam (Mendapatkan Remote Kontrol).
- ` + "`" + `*var` + "`" + ` (Asterisk): Menunjuk penampung ke nilai asli dari alamat Remote Kontrol agar bisa diretas atau dibaca.

` + "```" + `go
func ubahAsli(teksTarget *string) {
    // Kita menembak derefencenya (isi kotak suratnya) karena kita diutus pegang remotenya (*string), ubah dari luar batas!
    *teksTarget = "HEKED" 
}

func main() {
    dataAsli := "KunciEmas"
    fmt.Println(dataAsli) // KunciEmas
    
    // Alih-alih melempar teks aslinya, kita melempar remot alamat RAM aslinya pakai lambang &
    ubahAsli(&dataAsli)
    
    fmt.Println(dataAsli) // HEKED !! Alangkah saktinya, file asalnya berhasil tercorang langsung tanpa kembali.
}
` + "```" + ``,
			ContentMarkdownEN: sp(`### Traversing Direct Memory RAM References Deep Engine Specs (Pointers)
Theorize environmental scenarios conceptually: Extrapolating copies *(copy-paste)* duplication word documentation frameworks relaying "doc-b.docx" clones upon external isolated hardware partitions. Modifications mutating contextual font structures rendering alien crimson iterations affect isolated duplicates heavily. Question: Do principal origin laptop source files endure mutations symmetrically? Assuredly false. Terminology designates strictly classifying **Pass By Value** operations (Symmetric duplication memory heap forging fake RAM clones implicitly).

Standard behavioral native Go instantiates pure duplications linearly. Implementing Pointers integration overrides this via explicit (Asterisk ` + "`" + `*` + "`" + ` and Ampersand ` + "`" + `&` + "`" + `) syntactic logic enabling *Pass By Reference* overriding memory boundaries hacking localized structural source variables bridging dimensional bounds explicitly.

- ` + "`" + `&var` + "`" + ` (Ampersand): Locates encrypted dimensional hexadecimal system memory allocation keys (Procuring universal operational Remote Remotes).
- ` + "`" + `*var` + "`" + ` (Asterisk): Decrypting pointing variables retrieving actual payload values behind encrypted Remote addresses triggering remote edits correctly.

` + "```" + `go
func overrideSources(targetTexts *string) {
    // Bypassing bounds directly mutating dereferenced mailbox internals exploiting external remote accesses (*string) contexts bypassing security perimeters!
    *targetTexts = "HACKED" 
}

func main() {
    originPayload := "GoldKey"
    fmt.Println(originPayload) // GoldKey
    
    // Refrains dumping payload variables strings literally, instead transmitting literal hexadecimal pointer referencing ampersand addresses.
    overrideSources(&originPayload)
    
    fmt.Println(originPayload) // HACKED !! Magical occurrences overriding localized parent source constructs irreversibly rendering state anomalies securely.
}
` + "```" + ``),
		},
		{
			ID:         lessonUUID(5, 13),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod4,
			OrderIndex: 13,
			XPReward:   10,
			TitleID:    "13. Cetakan Struct (Object-Oriented di Go)",
			TitleEN:    sp("13. Struct Blueprints (Object-Oriented Go)"),
			ContentMarkdownID: `### Membuat Tipe Data Custom Sendiri!
Go bukalah bahasa pure OOP konvensional layaknya Java/C# (Ia tidak mengenal kata kerja ` + "`" + `class` + "`" + ` atau *Inheritance* kaku). Sebaliknya, Go mengantongi fitur *Struct* yang mewakilkan kombinasi dari variabel variabel primitif menjadi satu entitas.

Contoh kita mendirikan cetak biru (Blueprint) entitas karakter bernama ` + "`" + `Pegawai` + "`" + `.

` + "```" + `go
// Deklarasi template wujud cetak biru pegawai
type Pegawai struct {
    ID     int
    Nama   string
    Aktif  bool
}

func main() {
    // Mulai mencetak kloning orangnya layaknya tuhan
    var manager Pegawai
    manager.ID = 100
    manager.Nama = "Arul"
    manager.Aktif = true
    
    // Atau memakai cara penulisan elit kilat 1 detik (Inline Construct)
    cepu := Pegawai{ID: 200, Nama: "Bejox", Aktif: false}
    
    fmt.Println("CEO kita adalah bapak", manager.Nama)
}
` + "```" + ``,
			ContentMarkdownEN: sp(`### Forging Custom Architectural Datatype Entities!
Fundamentally, Go inherently neglects conventional rigid Object-Oriented generic implementations heavily mirroring standard Java/C# specifications (Deprecating verbs encompassing absolute ` + "`" + `classes` + "`" + ` inheritance trees strictly). Reversely compensating flexibility leveraging robust encapsulated *Struct* blueprints grouping associative disjoint primitive arrays fabricating singular modular abstractions flawlessly.

Constructing conceptual blueprint character configurations modeling systemic administrative ` + "`" + `Employee` + "`" + ` hierarchies.

` + "```" + `go
// Architectonic template molding blueprint explicit declarative formats structures
type Employee struct {
    ID       int
    FullName string
    IsActive bool
}

func main() {
    // Activating divine entity cloning orchestrations programmatically
    var manager Employee
    manager.ID = 100
    manager.FullName = "Arul"
    manager.IsActive = true
    
    // Swift elite inline syntax abbreviations deploying instantaneous instantiation configurations.
    informant := Employee{ID: 200, FullName: "Bejox", IsActive: false}
    
    fmt.Println("Enterprise top dog CEO correlates matching", manager.FullName)
}
` + "```" + ``),
		},
		{
			ID:         lessonUUID(5, 14),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod4,
			OrderIndex: 14,
			XPReward:   10,
			TitleID:    "14. Menempelkan Fungsi ke Struct (Method)",
			TitleEN:    sp("14. Attaching Functions to Structs (Methods)"),
			ContentMarkdownID: `### Memberikan Kehidupan Kepada Objek Kaku (Method)
Di bab sebelumnya, struct ` + "`" + `Pegawai` + "`" + ` itu layaknya patung mati (cuma tempat menyimpan data Nama dan ID belaka). Bagaimana kalau kita ingin membuatnya bisa punya "Tindakan" (*behavior*) layaknya game RPG di mana si pegawai ini bisa "Menendang pintu" atau "Berkenalan"?

Itulah yang dia sebut Method. Sebuah fungsi khusus yang kita ikat silang secara harfiah ke leher Struct nya! (Perhatikan kurung kecil yang menyusup di depan tulisan *func*).

` + "```" + `go
type Robot struct {
    Nama string
    Tipe string
}

// Ini bukan fungsi biasa! Ini METHOD. 
// Perhatikan sisipan ganjil (r Robot) di belakang tulisan func? Itu namanya (Receiver / Penerima)
// Mulai detik ini, fungsi Sapaan ini adalah BUDAK MUTLAK miliknya struct Robot.
func (r Robot) Sapaan() {
    fmt.Println("Bip bop. Namaku adalah", r.Nama, "dengan sistem modulatis tipe", r.Tipe)
}

func main() {
    wallE := Robot{Nama: "Wall-E", Tipe: "Scavenger"}
    
    // Cara memanggilnya persis memecahkan class OOP modern dengan syntax TITIK
    wallE.Sapaan()
}
` + "```" + ``,
			ContentMarkdownEN: sp(`### Sparking Behavioral Animation towards Rigid Entities (Methods)
Evidencing preliminary chapters, basic fundamental ` + "`" + `Employee` + "`" + ` structures exhibit inert rigid mannequin attributes (Storing exclusively inert variable node assignments). Postulate orchestrating active RPG game interactions animating functional operations directing designated components initiating physical engagements (Kicking boundaries/Emitting conversational sequences)?

Taxonomies categorizing bounded functions manifest inherently defined as Methods. Specialized distinct operational functions bound explicitly intertwining strictly surrounding fundamental parent Struct perimeters! (Identify arbitrary syntactical brackets interceding explicitly prepending *func* lexicons).

` + "```" + `go
type Droid struct {
    Designation string
    ModelType   string
}

// Negates standard function isolation taxonomy completely! Validating METHOD categorizations.
// Dissect anomalous invasive (r Droid) insertion nodes following func logic? Receiver binding mechanics trigger!
// Operating concurrently, specific function domains restrictively shackle operational executions bounded totally serving Droid architecture templates exclusively globally.
func (r Droid) Greetings() {
    fmt.Println("Bip bop. Identifying designation sequence equals", r.Designation, "incorporating core modularity specs typing", r.ModelType)
}

func main() {
    wallE := Droid{Designation: "Wall-E", ModelType: "Scavenger"}
    
    // Triggering sequential pipeline executions simulating modern OOP class interactions utilizing explicit literal DOT chaining syntax.
    wallE.Greetings()
}
` + "```" + ``),
		},
		{
			ID:         lessonUUID(5, 15),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod4,
			OrderIndex: 15,
			XPReward:   10,
			TitleID:    "15. Mewariskan Kewajiban (Interface)",
			TitleEN:    sp("15. Inheriting Obligatory Contracts (Interface)"),
			ContentMarkdownID: `### Apa Itu Kontrak Sosial Interface?
Go tak punya "Inheritance" / Kelas Anak Bapak beneran dari fitur OOP. Go cuma menggunakan sistem **Interface** yang berperan sebagai "Daftar Syarat & Ketentuan Bebas".

Bayangkan tipe data bernama ` + "`" + `Bangun Datar` + "`" + `. Nah, lingkaran dan Segitiga sangat berbeda rumusnya total sedari lahir. Namun mereka bernaung di alam yang sama, yaitu mereka *SAMA-SAMA HARUS WAJIB PUNYA NILAI LUAS*.

` + "```" + `go
// Kita buat papan kontrak (Hukum absolut interface semesta alam): "Siapa pun lu! Benda apapun lu! Kalo mu nyemplung dan lulus audisi jadi golongan gue ` + "`" + `BangunDatar` + "`" + `, lu WAJIB ngabarin gue rumus Luas lo tu apa return nya double float."
type BangunDatar interface {
    HitungLuas() float64
}

// Tersangka 1: Lingkaran
type Lingkaran struct {
    Jari float64
}
// Diam-diam si lingkaran menyanggupi syarat kontrak (Menambahkan fungsi dengan nama yang PLES PLEK SAMA persis HitungLuas). 
// Alhasil, Golang otomatis secara ajaib meresmikan Lingkaran adalah Anak Angkat sah dari kelompok BangunDatar.
func (l Lingkaran) HitungLuas() float64 {
    return 3.14 * l.Jari * l.Jari
}

// Tersangka 2: Persegi Panjang
type Persegi struct {
    Sisi float64
}
func (p Persegi) HitungLuas() float64 {
    return p.Sisi * p.Sisi
}
` + "```" + ``,
			ContentMarkdownEN: sp(`### What Defines Interface Social Paradigms Structurally?
Categorically denouncing hierarchical traditional parent-child class Inheritance architectures. Go supplants paradigms adopting unified implicit **Interface** systems establishing flexible "Obligatory Regulatory Contract Terms".

Theorize formulating universal datatypes encapsulated explicitly acting universally bounding ` + "`" + `FlatShapes` + "`" + `. Mathematical derivatives intersecting logic dictating Triangles comparing distinct structural parameters opposing Spherical circles natively inherently disjoint functionally. Contrastingly interconnected establishing single governing physical absolutes. Both instances conceptually demand bounding universal mathematical *Total Calculated Areas Output Variables*.

` + "```" + `go
// Dictating universal absolute cosmic structural contractual agreements manifesting explicit generic abstraction tiers: "Entities adhering arbitrary topologies! Submitting applications bypassing validation integrating core taxonomy classes registering ` + "`" + `FlatShapes` + "`" + ` demographics, MUST mandatorily satisfy implementations returning strictly absolute generic Area formulas yielding designated double floats format."
type FlatShape interface {
    CalculateArea() float64
}

// Subject Hypothesis 1: Circular entities
type Circle struct {
    Radius float64
}
// Implicitly implicitly registering adherence satisfying boundary contract parameters smoothly (Injecting identical matching function namespaces accurately mapping abstract CalculateArea schemas). 
// Consequently Go compiler magically intrinsically inaugurates Circular architectural branches declaring official structural adoption confirming categorical compliance integrating FlatShape clusters silently flexibly.
func (l Circle) CalculateArea() float64 {
    return 3.14 * l.Radius * l.Radius
}

// Subject Hypothesis 2: Equilateral Squares 
type Box struct {
    Side float64
}
func (p Box) CalculateArea() float64 {
    return p.Side * p.Side
}
` + "```" + ``),
		},
		{
			ID:         lessonUUID(5, 16),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod5,
			OrderIndex: 16,
			XPReward:   10,
			TitleID:    "16. Menghadapi Error dengan Kesatria (Error Handling)",
			TitleEN:    sp("16. Facing Errors Valiantly (Error Handling)"),
			ContentMarkdownID: `### Seni Pengecekan Error di Golang
Dunia programming itu tidak indah. Terkadang file yang kamu suruh minta dicari di komputer nggak ada dan terhapus orang tak dikenal. Golang paling benci (dan bahkan membuang sama sekali) sistem tradisional seperti ` + "`" + `try-catch` + "`" + ` yang suka *ngeredam* log sistem error diem dim dim hingga aplikasinya tahu-tahu mental *Crash*.

Golang secara transparan mewajibkan developer memegang 2 buah kembalian di seluruh sistem di semesta kerjanya: ` + "`" + `Hasil_Asli` + "`" + ` dan ` + "`" + `Status_Error` + "`" + `.

` + "```" + `go
import (
    "errors"
    "fmt"
)

// Contoh fungsi bagi. Kalau membagi nol gawat kan? Aplikasi kasir lu bisa Crash.
func Bagi(a, b int) (int, error) {
    if b == 0 {
        // Karena ada bahaya nol menumpuk, lemparkan Tanda Error Merah menggunakan pkg errors.
        return 0, errors.New("Tidak bisa membagi nilai absolut denga angka nol brother!")
    }
    // Jika lewat dan aman lalui validasi, maka return hasil normalnya namun error dibuang jadi "nil" alias (Nihil/Aman sentosa)
    return a / b, nil
}

func main() {
    hasil, balasan_error := Bagi(10, 0)
    
    // Setiap ngejalanin program apapun di Golang, kamu BAKAL dan SELALU WAJIB ngecek percabangan error != nil. Ini signature khas Go sedunia. Cuma Developer Go aslik yang ngetik snippet legendaris ini.
    if balasan_error != nil {
        fmt.Println("Waduh ketangkep basah log server ancur:", balasan_error.Error())
        return // Keluar pelarian jangan di lanjot masbrow!
    }
    
    fmt.Println("Semuanya mulus, ini hadiahnya kalkulator:", hasil)
}
` + "```" + ``,
			ContentMarkdownEN: sp(`### Artisanal Go Error Paradigms
Programming ecosystems lack romanticized stability parameters dynamically navigating external variables crashing constantly (Missing local server IO configurations manipulated by external vectors). Go rejects natively implicitly orchestrating structural components circumventing traditional ` + "`" + `try-catch` + "`" + ` block methodologies aggressively muffling systemic execution warnings silently forcing spontaneous application crashes randomly bypassing monitoring diagnostics.

Go mandates absolute explicitly enforced transparency compelling developers integrating binary asynchronous system tuple callbacks dictating mandatory return paths consistently exporting parallel variables globally traversing structural nodes: ` + "`" + `Actual_Data_Loads` + "`" + ` intersecting respective ` + "`" + `Error_Status_Checks` + "`" + `.

` + "```" + `go
import (
    "errors"
    "fmt"
)

// Simulated fractional division function implementations. Dividing absolute zero endpoints risks crashing explicit terminal Point of Sale POS executions globally.
func Divide(a, b int) (int, error) {
    if b == 0 {
        // Triggering explicit danger protocols detecting absolute zero denominator stack overflows projecting rigid crimson Error signals executing core 'errors' pkg formatting.
        return 0, errors.New("Evaluating absolute numerical division traversing categorical zero denominators denied inherently brother!")
    }
    // Validating safe mathematical fractions explicitly bypassing internal security checks validating structural boundaries correctly returning operational computations throwing out target error variables translating implicitly nullifying "nil" signals strictly indicating (Voided/Benign Operational Health Contexts).
    return a / b, nil
}

func main() {
    results, payload_error := Divide(10, 0)
    
    // Engaging structural runtime sequences traversing generic Go orchestrations inherently mandates rigorous manual procedural validation sequences evaluating explicit 'if err != nil' checks ubiquitously mapping canonical Go engineering globally. Solely elite Go engineers document these exact legendary snippets.
    if payload_error != nil {
        fmt.Println("Critical infrastructure structural log interception failures breaching parameters:", payload_error.Error())
        return // Abandoning manual operational thread execution lines terminating process urgently exiting sequences!
    }
    
    fmt.Println("Operational executions smooth flawlessly, yielding calculated terminal payloads:", results)
}
` + "```" + ``),
		},
		{
			ID:         lessonUUID(5, 17),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod5,
			OrderIndex: 17,
			XPReward:   10,
			TitleID:    "17. Paket Bawaan Wajib (fmt, strings, math)",
			TitleEN:    sp("17. Mandatory Standard Libraries (fmt, strings, math)"),
			ContentMarkdownID: `### Menyewa Peralatan Tempur Koding Gratisan (StdLib Go)
Sistem library bawaan Go sering disebut sebagai "Standard Library" (Stdlib) terkuat sedunia programming. Kamu sangat mandiri dalam menulis segala hal, dari membangun HTTP Web Server hingga enkripsi crypto *Hash 256* canggih, **semunya bisa dilakukaN murni tanpa perlu men-download package library Pihak Ketiga (seperti npm, pip)**.

6. Beberapa senapan (Package) yang sangat familiar sering diimport di awal kode ` + "`" + `.go` + "`" + ` kamu:

1. ` + "`" + `fmt` + "`" + ` = Memformat output terminal (Membentuk *Println*, *Printf* angka desimal dll).
2. ` + "`" + `strings` + "`" + ` = Mengolah operasi teks manipulasi kata rumit (Misal memisahkan kata dengan spasi, mengecilkan hurup besar kapital ` + "`" + `strings.ToLower()` + "`" + `, hingga mencocokkan pattern kalimat pencarian teks substring!).
3. ` + "`" + `math` + "`" + ` = Melakukan fungsional kalkulator yang lebih edan dari tambah/kurang/kali/bagi (Misal membumbuhkan fungsi trigonometri Sin/Cos, menancapkan rumus pembulatan eksklusif akar kuadrat, dll)
4. ` + "`" + `strconv` + "`" + ` = Konvertor sihir memanjakan logika pergantian jenis wujud benda. Berfungsi memaksakan integer angka 90 menjadi string Teks "90" (*String to Converter*). Sangat vital saat komunikasi log front-end HTML. 
`,
			ContentMarkdownEN: sp(`### Leasing Extensively Armed Free Toolchains (Go Standard Libraries)
Globally renowned ecosystem architectures establish innate Go package systems heavily credited formulating programming world's pinnacle "Standard Library" (Stdlib) hierarchies autonomously robust implicitly out-of-the-box. Facilitating extreme localized self-sufficient environments executing intricate structural developments ranging building full-fledged embedded HTTP Web Server hosts validating advanced cryptographic encryption sequences *Hash 256* integrations perfectly natively **implementing complex execution layers avoiding mandatory Third-Party dependency installations strictly maintaining isolation cleanly (Discarding generic npm, pip packages).**

Essential fundamental core components generically integrated mapping basic routine executable initializations configuring primitive ` + "`" + `.go` + "`" + ` scripts natively typically include:

1. ` + "`" + `fmt` + "`" + ` = Terminal output structural text sequence manipulation algorithms (Constructing literal console *Println*, abstract hexadecimal binary formatting sequences *Printf* fractional configurations, etc).
2. ` + "`" + `strings` + "`" + ` = Synthesizing explicit character operations parsing complicated string matrix logic (Slicing delimited sequence spaces iteratively implicitly enforcing programmatic uppercase downshifts utilizing robust ` + "`" + `strings.ToLower()` + "`" + ` procedures intercepting textual search algorithm pattern indexing mappings matching exact dynamic query substring operations!).
3. ` + "`" + `math` + "`" + ` = Empowering advanced systematic mathematical floating matrix operational computations vastly scaling generic mathematical addition/subtraction/modulo calculations (Projecting abstract geometric angular trigonometry derivations rendering Cos/Sin, executing strict rigid rounding logic computing specific quadratic square root executions extensively).
4. ` + "`" + `strconv` + "`" + ` = Literal systemic object mutation magical data-type convertors ensuring dimensional shifts evaluating physical form manipulations dynamically. Explicitly forcing categorical variable integers formatting digits exactly mapping literal payload text equivalents "90" (*String to Converter* sequences). Vital necessity bridging asynchronous translation APIs frontend endpoints bridging logical text payload processing HTML inputs securely.`),
		},
		{
			ID:         lessonUUID(5, 18),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod5,
			OrderIndex: 18,
			XPReward:   10,
			TitleID:    "18. Berkenalan Bersama Goroutine Mutakhir",
			TitleEN:    sp("18. Meeting the state of the art Goroutines"),
			ContentMarkdownID: `### Ilmu Sihir Multitasking Sejati Milik Google
Tahukah kamu *Thread* paralel di bahasa Java memakan RAM hingga 1 MB per kepala, sehingga laptop canggih dengan RAM 8 GB hanya mentok me-*spawn* sekitar 8000 operasi bersamaan sebelum ia ngos-ngosan Hang.

Di Golang, sang dewa arsitektur komputer Google memalsukan hardware RAM itu menjadi Thread ringan teringan di planet bumi. Sebuah "Goroutine" di Go hanya memakan 2 KB memori per buahnya. Alhasil, server abal-abal 1 GB aja kuat menjalankan hingga SETENGAH JUTA goroutine tugas barengan tanpa Crash sama sekali!

Cara pakainya sungguh bikin ketawa karena terlalu gampang abis. Cukup tambahkan kata sakti: ` + "`" + `go` + "`" + ` sebelum memanggil fungsi fungsi tertentu.

` + "```" + `go
import (
    "fmt"
    "time"
)

func cetakSiksaMesin(pesan string) {
    for i := 0; i < 5; i++ {
        fmt.Println(pesan)
        time.Sleep(100 * time.Millisecond) // jeda tidur mesin
    }
}

func main() {
    // Kalau kamu manggil 2 di bawah secara biasa kek gini, proses kode bakalan nungguin blok fungsi pertama selesai sampai mati baru lanjut blok baris kedua (Synchronous).
    // cetakSiksaMesin("Lambat banget 1")
    // cetakSiksaMesin("Lambat banget 2")
    
    // NAH Kalau kita kasih kata mutiara "go" depannya. Dia bakal dilempar ke Alam Paralel baru jalan di background system diam-diam barengan eksekusi membelah core prosessor asli secara serempak.
    go cetakSiksaMesin("Ngedrift Multitasking 1")
    go cetakSiksaMesin("Ngedrift Multitasking 2")
    
    // Tapi karena goroutine ini jalan di alam gaib background, kadang main func program utamanya udah selesai kelewat duluan padahal si 2 kuli di atas baru nguli kerja semen detik pertama.
    // Akibat program selesai, semua fungsi kuli alam paralel mati terbunuh dipaksa henti. 
    // Jadi sementara khusus materi pemula untuk menahan laju pintu keluar program kita sumbat kancing pake penundaan 1 detik agar kuli 100 ms nya kelar kerja aman semua sebelum exit aplikasi utuh.
    time.Sleep(1 * time.Second) 
}
` + "```" + `
Selamat, kamu baru saja menjadi master asynchronus!`,
			ContentMarkdownEN: sp(`### Google's Ultimate True Parallel Mutability Magic Spells
Conceiving contextual systemic Thread parallel operational loads integrating Java structural environments intrinsically demand exponential 1 MB physical RAM sector overhead distributions per execution node concurrently, dictating standard 8GB advanced laptops strictly bottleneck reaching peaking thresholds around exactly 8000 parallel node limits before enduring critical system freeze locks. 

Mapping abstract computing architecture paradigms Go developers formulated engineered explicit hardware RAM virtualization paradigms forging absolute lightest computing system Threads known across Earth's limits. Standard "Goroutines" allocate merely 2 KB fractional payload segments explicitly mapping specific execution logic overheads autonomously. Consequently enabling heavily restricted cheap 1 GB limited proxy instances securely scaling massive parallel asynchronous operations orchestrating literally HALF A MILLION concurrent background operational sequences lacking fatal server kernel panic crashes seamlessly!

Activating execution configurations demands extremely simplified parameters inducing laughter observing ridiculous implementation simplicity models. Solely prefix the legendary absolute power keyword: ` + "`" + `go` + "`" + ` preemptively initiating target function executions concurrently!

` + "```" + `go
import (
    "fmt"
    "time"
)

func executeMachineEngineTortureLogics(designationPayloadMessage string) {
    for i := 0; i < 5; i++ {
        fmt.Println(designationPayloadMessage)
        time.Sleep(100 * time.Millisecond) // initiating mechanical structural delay pauses
    }
}

func main() {
    // Standard execution bindings trigger sequential linear processing locking external operational flow preventing consecutive step integrations blocking thread flows entirely until initial executions shatter logic parameters completing iterations linearly (Synchronous operations).
    // executeMachineEngineTortureLogics("Lethargically sluggish block 1")
    // executeMachineEngineTortureLogics("Lethargically sluggish block 2")
    
    // CONVERSELY Injecting explicit legendary semantic prefix keyword validations invoking "go" syntax logic orchestrations. Forcing functions deploying instances across fragmented parallel abstract universe layers activating explicit background background stealth operations processing concurrently shattering authentic hardware multi-core CPU physical bounds instantly processing synchronously symmetrically concurrently.
    go executeMachineEngineTortureLogics("Asynchronous concurrent drifting mechanisms 1")
    go executeMachineEngineTortureLogics("Asynchronous concurrent drifting mechanisms 2")
    
    // Consequentially resulting explicit parallel environment integrations operating concurrently within independent abstract layers, core initial thread operational blocks running main func pipelines bypass logic sequences resolving validations drastically earlier bypassing active background slave deployments actively executing sequential operations generating microsecond delays securely trailing validations initially.
    // Forcefully main routine resolving operations terminating program scopes inevitably murders background working goroutine slaves explicitly terminating active workloads permanently.
    // Hence uniquely mapping beginner structural architectures temporarily suspending exit portals bypassing exit validation procedures using explicit synthetic 1 second timeout variables ensuring internal 100 ms slave loops complete iterations thoroughly traversing operations executing correctly terminating completely prior ultimate explicit environment exit instances.
    time.Sleep(1 * time.Second) 
}
` + "```" + `
Congratulations mastering asynchronous advanced systemic execution workflows!`),
		},
		{
			ID:         lessonUUID(5, 19),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod5,
			OrderIndex: 19,
			XPReward:   10,
			TitleID:    "19. Pipa Goroutine Sakti (Channel)",
			TitleEN:    sp("19. Mystic Goroutine Pipelines (Channels)"),
			ContentMarkdownID: `### Telepati Asinkron Antar Alam
Bicara soal trik rahasia sebelumnya. Sangat haram hukumnya di level produksi (skala rill gojek/tokped) untuk menahan fungsi program selesai di baris terakhir exit *Main Function* menggunakan sandi kotor ` + "`" + `Time.Sleep` + "`" + ` menunda waktu acak gaje gitu. (Bagaimana jika program kuli background gagal, terhenti lebih lama misal butuh 5 detik API fetch data macet, lalu timer time sleep 1 detik udah abis? Terputus lah kulinya mati mendadak wadoooh!).

Untuk mengatasi problem di atas di mana dua alam *Goroutine Asynchronous Parallel* dan *Kondisi Rutin Sinkronis Main App* itu berbeda waktu jalan (sebut saja alam manusia dan alam jin), maka Go menyediakan **Channel**, sebatang pipa beton kuat magis supranatural lintas ruang waktu dimensi. 

Lewat pipa ini, alam Jin (` + "`" + `go fun()` + "`" + ` goroutine) melempar serpihan info data integer dari dalam pipa masuk. Lalu di sisi ujung keluarnya, si alam Manusia secara setia dan diam membeku MENGHALANGI pintu program exit demi terhipnotis di depan ujung pipa, sembari siaga menunggui benda itu kelempar keluar di tangkap di baris var nya. Secara magis sinkronasi aman!. 

` + "```" + `go
package main

import "fmt"

func kuliDiAlamGaib(pipaPenyambung chan string) {
    fmt.Println("Jin bekerja memahat batu 10 tahun (0.01 detik)...")
    // Setelah batu beres, jin ngirim pesannya MASUK ke dalam ujung pipa alam sana pake operator '<-' PANAH MASUK KE PIPA (chanel)
    pipaPenyambung <- "Batu Candi Prambanan Selesai Bos jin!"
}

func main() { // Alam manusia
    // 1. Manusia Membangun Pipa Ajaib Khusus Angkut Tipe Data String memakai func make(chan type_data)
    pipaLintasDimensi := make(chan string)

    // 2. Kirim utusan kerja di alam gaib jalan selaras parelal (go)
    go kuliDiAlamGaib(pipaLintasDimensi)

    // 3. Main program santai (gak perlu sleep-sleep gaje). Kita suruh alam manusia buat diam nangkring jongkok nungguin ujung pipa ngeluarin barang (Biar gampang lihat PANAH KELUAR DARI SAMPING PIPANYA MASUK KE VAR). 
    // Baris ini bakalan nge-BEKU-IN APP alias mem-BLOCKED jalan eksekusi selanjutnya hingga kapanpun paket stringnya tiba di lempar. Aman!
    hasilTangkapan := <-pipaLintasDimensi

    fmt.Println("Laporan Jin masuk telinga:", hasilTangkapan)
}
` + "```" + ``,
			ContentMarkdownEN: sp(`### Asynchronous Inter-Dimensional Telepathy
Reflecting upon previously conceptualized hidden manipulation matrix variables manipulating main functional terminations implementing primitive ` + "`" + `Time.Sleep` + "`" + ` randomization execution structural blockages remains vehemently catastrophic configuring robust scalable commercial backend system environments explicitly. (Considering empirical latency variances projecting dynamic 5-second asynchronous explicit data fetch delays triggering failures halting external pipeline nodes arbitrarily terminating executing worker components entirely traversing ungraceful shutdown instances triggering forced core network closures anomalously).

Countering specific structural asynchronous latency inconsistencies operating inherently diverging *Goroutine Parallel execution speeds* conflicting explicitly *Synchronous Main Routing loops* mapping inherently (Analogy depicting separate Human conscious universe operations traversing parallel stealth execution Jinn matrices explicitly), Go dynamically integrates magical robust concrete bridging pipelines navigating inter-dimensional boundary limits designating **Channels**.

Empowering explicit channel endpoints, isolated executing functions (Goroutines instance mappings) projects textual literal payload messages injecting binary streams intersecting inbound pipe channels natively. Conversely executing external termination nodes, native original sequences remain structurally frozen indefinitely blocking exact termination execution paths unconditionally waiting anticipating specific sequential structural variable extractions popping implicitly magically mitigating sync errors.

` + "```" + `go
package main

import "fmt"

func parallelBackgroundWorkerNodes(bridgingExecutionPipeline chan string) {
    fmt.Println("Worker traversing abstract environments allocating computing vectors (0.01s bounds)...")
    // Concluding operation bounds, executing node logically transmits target operational payload states INBOUND intersecting mapped pipeline parameters implementing dynamic specific '<-' LEFT INBOUND ARROW operator sequences
    bridgingExecutionPipeline <- "Execution payload operation validations explicitly resolved successfully Administrator node!"
}

func main() { // Original universe sequences
    // 1. System explicitly manufactures magical target structural payload string data routing generic pipeline arrays instantiating explicit func make(chan data_type) instances natively.
    crossDimensionalRoutingPipes := make(chan string)

    // 2. Initiating decoupled abstract parallel concurrent operations utilizing explicit execution (go) parameters.
    go parallelBackgroundWorkerNodes(crossDimensionalRoutingPipes)

    // 3. Negating archaic latency sleep approximations completely! Executing operations command structural blockages isolating execution path bounds pausing entire thread lifecycle waiting endpoints anticipating explicitly specific payload message data ejections returning specific structural endpoints successfully. (Observe dynamic outbound arrow structures targeting target variables).
    // Executing target node inherently FREEZES execution lifecycle states establishing blocking execution routing preventing sequential main system operations prematurely entirely circumventing timing bugs accurately reliably.
    capturedPayloads := <-crossDimensionalRoutingPipes

    fmt.Println("Received asynchronous operation execution payload terminal updates:", capturedPayloads)
}
` + "```" + ``),
		},
		{
			ID:         lessonUUID(5, 20),
			CourseID:   CourseGoBeg,
			ModuleID:   &mod5,
			OrderIndex: 20,
			XPReward:   25,
			TitleID:    "20. Kelulusan (Masa Depan Gophers)",
			TitleEN:    sp("20. Graduation (Futures for Gophers)"),
			ContentMarkdownID: `### Selamat Lulus Gopher Muda! 🏆

*(Gopher: Sebutan imut maskot tikus tanah lucu programmer bahasa Go sedunia).*
Kamu berhasil menguasai seluruh fundamental keras dari kasta tertinggi tata bahasa pemrograman dewa yang merajai Microservices zaman sekarang. Jika kamu perhatikan teliti, Go itu simpel kan? Ga muter muter banyak *keyword* aneh warisan dewa java purba yang membingungkan. 

**Mau ngapain lagi nih gw bang arul setelah ini?**
1. **Latih Logika Terus**: Masuk menu Chalenge! Silakan bertarung menaklukkan 50 ++ level test case coding sandbox murni langsung di browser app NgodingYuks. Asah jam terbang Go kalian dari level easy - expert!.
2. **Kembangkan Portfolio Proyek**: Saatnya kamu keluar goha. Cobain instal library framework seperti **Echo** (Framework Web mirip express JS node) atau pake **Fiber**. Kemudian satukan nyawanya dan tarik pelatuknya ke database aseli pake framework **GORM**. Percayalah rasanya bakal muluuuuus banget bikin website/api pakai Go.
3. **Pahami Clean Architecture**: Gada gunanya ngoding super elit pake channel gorutine kalo foldering project Go lu ancur (semuanya dimasukin ke func main wkwk). Coba pelajari desain *Clean Architecture MVC* untuk level advance gopher kelak.

Sekian dan terima kasih kawan kawan. Minta doanya ya bang kawan moga *NgodingYuk* makin jaya sukses melang-lang buana hehe, See you and Keep Coding!`,
			ContentMarkdownEN: sp(`### Monumental Congratulations Junior Gopher Graduates! 🏆

*(Gopher: Universal global affectionate nomenclature defining enthusiastic Go programmer domains representing iconic adorable mascot equivalents).*
Successfully conquering overarching rigorous architectural foundational basics integrating supreme top-tier contemporary microservices execution structural syntax abstractions defining modern computing operations flawlessly! Observing empirically, Go configurations intuitively synthesize simplicity properly negating bloated confusing legacy archaic class hierarchies plaguing alternative convoluted enterprise ecosystems inherently perfectly.

**Navigating Future Trajectories Post-Graduation Directives?**
1. **Continuously Expand Abstract Logic Horizons**: Engage explicit execution logic traversing NgodingYuks interactive browser Challenge platforms encompassing rigorous interactive coding sandbox tests manipulating 50+ multidimensional scenarios directly integrating logic implementations iterating dynamic complexity schemas systematically enhancing elite execution mechanics!
2. **Expanding Open-Source Portfolio Repository Architectures**: Exploring comprehensive structural advanced components implementing explicit scalable operational framework APIs adopting modern routing engines matching **Echo** (Correlating Node.js Express analogues natively) adopting high-speed variant alternatives leveraging **Fiber** parameters strictly formatting execution boundaries structurally hooking persistent structural ORM data schemas utilizing **GORM** executing robust scalable structural implementations deploying functional architectural backend Web/API domains concurrently robustly.
3. **Adopting Modular Architectural Domain Standards**: Negating advanced explicit parallel computing parameters natively executed perfectly strictly lacks robust utility assuming structural folder operational implementations aggregate chaotically mapping messy spaghetti code paradigms grouping unstructured functions indiscriminately. Mastering explicit structured *Clean Architecture MVC* separations guarantees massive scalability integrations systematically expanding Gopher backend environment boundaries expertly.

Concluding explicitly extending immense gratitude participants navigating exhaustive learning sequences continuously interacting comprehensively. Explicitly seeking continuous communal blessing parameters elevating overarching *NgodingYuk* organizational expansion paradigms sequentially systematically achieving sustained success globally dynamically! Continuously coding perpetually executing consistently mastering syntax integrations flawlessly! `),
		},
	}

	for _, lesson := range lessons {
		upsertLesson(db, &lesson)
	}
}
