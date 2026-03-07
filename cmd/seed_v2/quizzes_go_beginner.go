package main

import (
	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"gorm.io/gorm"
)

func seedGoBeginnerQuizzes(db *gorm.DB) {
	quizzes := []domain.LessonQuiz{
		// Lesson 1: Pengenalan Golang (1 quiz)
		{
			ID:           quizUUID(5, 1, 1),
			LessonID:     lessonUUID(5, 1),
			QuestionID:   "Siapa perusahaan di balik penciptaan bahasa pemrograman Go?",
			QuestionEN:   strPtr("Which company is behind the creation of the Go programming language?"),
			OptionsID:    j(`["Microsoft", "Google", "Facebook", "Apple"]`),
			OptionsEN:    j(`["Microsoft", "Google", "Facebook", "Apple"]`),
			CorrectIndex: 1, // Google
			XPReward:     5,
		},

		// Lesson 2: Hello World (2 quizzes)
		{
			ID:           quizUUID(5, 2, 1),
			LessonID:     lessonUUID(5, 2),
			QuestionID:   "Apa nama package wajib bagi sebuah program Go agar bisa dieksekusi memunculkan file bin (executable)?",
			QuestionEN:   strPtr("What is the mandatory package name for a Go program to be compiled as an executable binary?"),
			OptionsID:    j(`["package main", "package execute", "package program", "package app"]`),
			OptionsEN:    j(`["package main", "package execute", "package program", "package app"]`),
			CorrectIndex: 0,
			XPReward:     5,
		},
		{
			ID:           quizUUID(5, 2, 2),
			LessonID:     lessonUUID(5, 2),
			QuestionID:   "Package bawaan apa yang dipakai untuk mencetak teks ke layar terminal?",
			QuestionEN:   strPtr("Which built-in package is used to print text to the terminal screen?"),
			OptionsID:    j(`["log", "print", "fmt", "io"]`),
			OptionsEN:    j(`["log", "print", "fmt", "io"]`),
			CorrectIndex: 2, // fmt
			XPReward:     5,
		},

		// Lesson 3: Variabel dan Konstanta (2 quizzes)
		{
			ID:           quizUUID(5, 3, 1),
			LessonID:     lessonUUID(5, 3),
			QuestionID:   "Simbol apa yang digunakan untuk \"Deklarasi Singkat\" (Shorthand Declaration) variabel tanpa menuliskan var dan tipe datanya?",
			QuestionEN:   strPtr("Which symbol is used for shorthand variable declarations without writing `var` and the data type?"),
			OptionsID:    j(`["=", "==", ":=", "=>"]`),
			OptionsEN:    j(`["=", "==", ":=", "=>"]`),
			CorrectIndex: 2,
			XPReward:     5,
		},
		{
			ID:           quizUUID(5, 3, 2),
			LessonID:     lessonUUID(5, 3),
			QuestionID:   "Keyword apa yang dipakai untuk menyimpan nilai harga mati yang tak bisa diubah lagi semenjak dibuat?",
			QuestionEN:   strPtr("What keyword is used to store absolute absolute fixed values that cannot be modified after compilation?"),
			OptionsID:    j(`["var", "final", "const", "static"]`),
			OptionsEN:    j(`["var", "final", "const", "static"]`),
			CorrectIndex: 2, // const
			XPReward:     5,
		},

		// Lesson 4: Tipe Data Dasar (2 quizzes)
		{
			ID:           quizUUID(5, 4, 1),
			LessonID:     lessonUUID(5, 4),
			QuestionID:   "Tipe data manakah di Go yang digunakan untuk menampung angka koma (desimal)?",
			QuestionEN:   strPtr("Which data type in Go is used to store fractional (decimal) numbers?"),
			OptionsID:    j(`["float64", "decimal", "int", "boolean"]`),
			OptionsEN:    j(`["float64", "decimal", "int", "boolean"]`),
			CorrectIndex: 0,
			XPReward:     5,
		},
		{
			ID:           quizUUID(5, 4, 2),
			LessonID:     lessonUUID(5, 4),
			QuestionID:   "Nilai manakah yang PALING BENAR untuk dimasukkan ke dalam tipe data `bool`?",
			QuestionEN:   strPtr("Which value is the MOST CORRECT to be passed into a `bool` data type?"),
			OptionsID:    j(`["1", "\"true\"", "false", "0"]`),
			OptionsEN:    j(`["1", "\"true\"", "false", "0"]`),
			CorrectIndex: 2,
			XPReward:     5,
		},

		// Lesson 5: Operator Logika (2 quizzes)
		{
			ID:           quizUUID(5, 5, 1),
			LessonID:     lessonUUID(5, 5),
			QuestionID:   "Apa hasil dari operasi `10 % 3` (Modulo)?",
			QuestionEN:   strPtr("What is the result of the `10 % 3` (Modulo) operation?"),
			OptionsID:    j(`["3.3", "3", "1", "0"]`),
			OptionsEN:    j(`["3.3", "3", "1", "0"]`),
			CorrectIndex: 2, // 1 (remainder)
			XPReward:     5,
		},
		{
			ID:           quizUUID(5, 5, 2),
			LessonID:     lessonUUID(5, 5),
			QuestionID:   "Operator logika mana yang mengharuskan SEMUA kondisi bernilai True?",
			QuestionEN:   strPtr("Which logical operator requires ALL conditions to be True?"),
			OptionsID:    j(`["|| (OR)", "&& (AND)", "!= (NOT EQUAL)", "== (EQUAL)"]`),
			OptionsEN:    j(`["|| (OR)", "&& (AND)", "!= (NOT EQUAL)", "== (EQUAL)"]`),
			CorrectIndex: 1, // &&
			XPReward:     5,
		},

		// Lesson 6: Percabangan If-Else (2 quizzes)
		{
			ID:           quizUUID(5, 6, 1),
			LessonID:     lessonUUID(5, 6),
			QuestionID:   "Apakah kita WAJIB meletakkan kondisi eksekusi di dalam tanda kurung bulat `(...)` pada sintaks If di Golang?",
			QuestionEN:   strPtr("Are we MANDATORY to enclose execution logic conditions inside round brackets `(...)` regarding If syntax structures within Golang?"),
			OptionsID:    j(`["Wajib", "Tidak Wajib (Opsional)", "Dilarang Format Standar", "Bebas"]`),
			OptionsEN:    j(`["Mandatory", "Not Mandatory (Optional)", "Forbidden by Standard Formats", "Free"]`),
			CorrectIndex: 2, // Go actually formats it without parentheses
			XPReward:     5,
		},
		{
			ID:           quizUUID(5, 6, 2),
			LessonID:     lessonUUID(5, 6),
			QuestionID:   "Apa yang terjadi jika variabel dideklarasikan di sebaris sesaat bersama *If statement*? `if x := 10...`",
			QuestionEN:   strPtr("What happens when variables are declared inline simultaneously inside an *If statement block*? `if x := 10...`"),
			OptionsID:    j(`["Error kompilasi", "Variabel bisa diakses dari ujung manapun", "Umur variabel hanya bertahan hidup secara lokal di dalam cakupan kurang kurawal If tersebut", "Nilainya langsung berubah nol"]`),
			OptionsEN:    j(`["Compilation Error", "Variables accessible from any endpoint universally", "Variables life-cycle solely survives locally constrained within the respective If curly bracket blocks", "Its magnitude resets spontaneously addressing zeroes"]`),
			CorrectIndex: 2,
			XPReward:     5,
		},

		// Lesson 7: For Loops (2 quizzes)
		{
			ID:           quizUUID(5, 7, 1),
			LessonID:     lessonUUID(5, 7),
			QuestionID:   "Keyword manakah yang digunakan untuk membuat perulangan murni (Loop) di Go?",
			QuestionEN:   strPtr("Which explicit keyword handles pure explicit loop iterations in Go?"),
			OptionsID:    j(`["while", "for", "do-while", "foreach"]`),
			OptionsEN:    j(`["while", "for", "do-while", "foreach"]`),
			CorrectIndex: 1, // "for" is the only looping keyword in Go
			XPReward:     5,
		},
		{
			ID:           quizUUID(5, 7, 2),
			LessonID:     lessonUUID(5, 7),
			QuestionID:   "Keyword sakti apa yang dipakai untuk memberhentikan / membobol paksa siklus Loop dari dalam secara instan?",
			QuestionEN:   strPtr("Which magical keyword strictly shatters / halts Loop cycle execution scopes instantly escaping bounds dynamically from internal depths?"),
			OptionsID:    j(`["continue", "return", "break", "stop"]`),
			OptionsEN:    j(`["continue", "return", "break", "stop"]`),
			CorrectIndex: 2,
			XPReward:     5,
		},

		// Lesson 8: Array & Slices (2 quizzes)
		{
			ID:           quizUUID(5, 8, 1),
			LessonID:     lessonUUID(5, 8),
			QuestionID:   "Apa ciri fisik paling mencolok yang membedakan penulisan inisialisasi dasar *Slice* dibandingkan tipe Array konvensional purba?",
			QuestionEN:   strPtr("What paramount physical trait visibly distinguishes the initialization matrix formulating fundamental *Slices* isolating conventionally native archaic Arrays?"),
			OptionsID:    j(`["Slice pakai kurung buka biasa ()", "Slice pakai tanda kurung siku kosong [] / tanpa disinggung alokasi nomor dimensi kuadrat maksimum", "Slice ditulis pakai map{}", "Slice tidak bisa di println()"]`),
			OptionsEN:    j(`["Slices utilize regular opening parentheses ()", "Slices deploy vacant empty bracket enclosures [] / omitting maximum numeric multidimensional ceiling thresholds", "Slices written encapsulating explicit map{} syntaxes", "Slices inherently cannot log into println() parameters"]`),
			CorrectIndex: 1,
			XPReward:     5,
		},
		{
			ID:           quizUUID(5, 8, 2),
			LessonID:     lessonUUID(5, 8),
			QuestionID:   "Bagaimana instruksi valid teraman dan termudah untuk menyuntikkan data payload mentah ke dalam Slice (Ekspansi Dinamis)?",
			QuestionEN:   strPtr("What designates valid safest simplistic programmatic operation appending raw payload injection vectors seamlessly expanding internal Slice arrays globally?"),
			OptionsID:    j(`["Slice = Slice + Data_Baru", "Slice.insert(Data_Baru)", "Slice = append(Slice, Data_Baru)", "Slice.push(Data_Baru)"]`),
			OptionsEN:    j(`["Slice = Slice + New_Data", "Slice.insert(New_Data)", "Slice = append(Slice, New_Data)", "Slice.push(New_Data)"]`),
			CorrectIndex: 2,
			XPReward:     5,
		},

		// Lesson 9: Map (2 quizzes)
		{
			ID:           quizUUID(5, 9, 1),
			LessonID:     lessonUUID(5, 9),
			QuestionID:   "Peran absolut krusial apa yang ditawarkan format Map dibandingkan deretan array numerik Slice pasif standar?",
			QuestionEN:   strPtr("What intrinsic universally crucial role implements dictation leveraging Maps drastically bypassing standard passive numerically array string sequence Slices dynamically?"),
			OptionsID:    j(`["Menyimpan indeks Key menggunakan custom objek Teks / Kata sandi acak", "Lebih hemat RAM", "Hanya bisa menampung Float", "Map tidak memakan resource"]`),
			OptionsEN:    j(`["Dictates associating designated key indexes exploiting random encrypted custom Password / Text strings objects", "Highly converses RAM hardware limits naturally", "Ascribes strictly isolated Floating variables universally", "Maps neglects systemic load executions inherently zero resources"]`),
			CorrectIndex: 0,
			XPReward:     5,
		},
		{
			ID:           quizUUID(5, 9, 2),
			LessonID:     lessonUUID(5, 9),
			QuestionID:   "Bagaimana kita bisa dengan rapi menghapus suatu gembok/key (beserta nilainya) dari pundi pundi kamus Map tersebut?",
			QuestionEN:   strPtr("How exactly can one systemically securely demolish dictionary padlock indices / keys seamlessly abandoning their values globally across memory Map architectures?"),
			OptionsID:    j(`["hapus(kamusArea, \"Kunci\")", "delete(kamusArea, \"Kunci\")", "clear(kamusArea)", "kamusArea -= \"Kunci\""]`),
			OptionsEN:    j(`["hapus(domainMap, \"Keys\")", "delete(domainMap, \"Keys\")", "clear(domainMap)", "domainMap -= \"Keys\""]`),
			CorrectIndex: 1,
			XPReward:     5,
		},

		// Lesson 10 & 11: Functions (2 quizzes)
		{
			ID:           quizUUID(5, 10, 1),
			LessonID:     lessonUUID(5, 10),
			QuestionID:   "Dengan keyword apa sebuah deklarasi pondasi func terdefinisikan utuh di sistem arsitektur?",
			QuestionEN:   strPtr("Utilizing which isolated keyword specifies complete fundamental functional component implementations across standard runtime environments systemically?"),
			OptionsID:    j(`["function", "fn", "func", "def"]`),
			OptionsEN:    j(`["function", "fn", "func", "def"]`),
			CorrectIndex: 2, // func
			XPReward:     5,
		},
		{
			ID:           quizUUID(5, 11, 1),
			LessonID:     lessonUUID(5, 11),
			QuestionID:   "Jika sebuah fungsi memuntahkan (Return) 2 tipe data berbeda sekaligus, dan pemrogram hanya ingin peduli / mengambil 1 isi datanya. Syntax ajaib / sandi abai apa yang ia pakai menghindari teriak kompilator Go?",
			QuestionEN:   strPtr("Assuming designated execution loops explicitly spawn 2 simultaneous parallel payload yields returning structurally concurrently. Programming operator elects retrieving exclusively isolated 1 data load payloads natively. Which implicit magic syntax parameter universally ignores uncaptured variables bypassing Go compilation scream validations cleanly?"),
			OptionsID:    j(`["tanda tanya (?)", "titik-titik (...)", "Garis bawah (_) Alias Blank identifier", "Garis strip (-)"]`),
			OptionsEN:    j(`["Question marks (?)", "Ellipses (...)", "Underscores (_) As canonical Blank Identifier logic", "Dashes (-)"]`),
			CorrectIndex: 2,
			XPReward:     5,
		},

		// Lesson 12: Pointers (2 quizzes)
		{
			ID:           quizUUID(5, 12, 1),
			LessonID:     lessonUUID(5, 12),
			QuestionID:   "Apa fungsi dari logo magis ampersand `&` apabila ditempelkan berdampingan persis di kiri nama variabel? (Contoh `&nomorRute`)",
			QuestionEN:   strPtr("Decipher precise functionality mapping mystical abstract ampersand logos `&` prefixed identically flanking arbitrary runtime variable signatures implicitly left alignments conceptually? (Example `&routeNumbers`)"),
			OptionsID:    j(`["Maknanya itu variabel AND", "Untuk mendapatkan letak koordinat alamat RAM aslinya (Pointer Reference Address)", "Untuk nge-print huruf & di layar", "Untuk memanggil func panic()"]`),
			OptionsEN:    j(`["Denotes implicitly strict algebraic operational AND nodes", "Retrieves exact hexadecimal coordinates isolating origin authentic hardware RAM pointer dimension architectures (Memory Addresses)", "Enforces structural explicit & log terminal printings inherently", "Constructing function executions traversing explicit panic() commands systemically"]`),
			CorrectIndex: 1,
			XPReward:     5,
		},
		{
			ID:           quizUUID(5, 12, 2),
			LessonID:     lessonUUID(5, 12),
			QuestionID:   "Lalu apa kebalikan sihirnya dari Asterisk `*` terhadap remot letak RAM di atas tersebut?",
			QuestionEN:   strPtr("Infer explicit reciprocal mystical counterpart properties encompassing isolated asterisks `*` navigating dimensional bounds parsing target address coordinates referenced contextually comprehensively dynamically?"),
			OptionsID:    j(`["Berfungsi mensimulasikan rumus Matematika perkalian pangkat 3", "Mendepak (Dereferencing) pelacak koordinat tadi agar membelah dimensi mencabut muatan murni (Membongkar/Membaca target isian aslinya)", "Mematikan server", "Mendeklarasikan array secara massal"]`),
			OptionsEN:    j(`["Systematically mimicking strict arithmetic exponential algebraic scaling structures", "Evicting procedural memory index parameters slicing dimensional mappings decoupling explicit generic pointer address coordinates extracting origin fundamental payload values natively (Dereferencing structural core internal data)", "Fatal server thread termination executions unilaterally", "Mass explicit dynamic matrix initialization constructs linearly"]`),
			CorrectIndex: 1,
			XPReward:     5,
		},

		// Lesson 13 & 14: Structs and Methods (3 quizzes)
		{
			ID:           quizUUID(5, 13, 1),
			LessonID:     lessonUUID(5, 13),
			QuestionID:   "Sebagaimana yang kita pahami bahwa Go membenci perbudakan hirarki cetakan kaku 'Class Inheritance' pada bahasa konvensional, fitur absolut krusial apa yang disisipkan Golang memadukan kombinasi tipe dasar dalam wujud tunggal identitas OOP?",
			QuestionEN:   strPtr("Ascribing explicit Go paradigms fundamentally detesting systemic hierarchy structures mapping archaic strict explicit 'Class Inheritances' conventionally globally natively... Which structural monolithic parameters explicitly defines implementations binding core disjoint abstraction primitives mapping unified collective logic abstractions replicating robust flexible OOP behaviors conceptually universally?"),
			OptionsID:    j(`["Struct", "Enum", "Map", "Type-Casting"]`),
			OptionsEN:    j(`["Structs", "Enums", "Maps", "Type-Castings"]`),
			CorrectIndex: 0,
			XPReward:     5,
		},
		{
			ID:           quizUUID(5, 14, 1),
			LessonID:     lessonUUID(5, 14),
			QuestionID:   "Ketika programmer jenius mau melekatkan / mengaitkan fungsionalitas murni menempel di sisi pondasi raga jasad entitas Struct, aksi deklarasi itu disebut?",
			QuestionEN:   strPtr("Assuming elite engineers demand anchoring intrinsic core logical explicit functional boundaries fusing natively embedding logic seamlessly enveloping underlying foundational target schema entity Struct matrix templates specifically—How correctly represents contextual terminology dictating exact methodology instantiations dynamically inherently globally?"),
			OptionsID:    j(`["Inheritance Object", "Pointers Polymorph", "Method Receiver Functions", "Global Variable Casting"]`),
			OptionsEN:    j(`["Inheritance Objects", "Pointers Polymorph", "Method Receiver Functions", "Global Variable Castings"]`),
			CorrectIndex: 2,
			XPReward:     5,
		},
		{
			ID:           quizUUID(5, 14, 2),
			LessonID:     lessonUUID(5, 14),
			QuestionID:   "Dimanakah titik penulisan nama jasad penerima (Receiver target Struct) ini di definisikan pada sintaks deklaratif fungsinya secara utuh?",
			QuestionEN:   strPtr("Precisely where locates explicit positional index alignments dictating isolated literal receiver payload architectures mapping boundaries dictating (Struct receiver definitions) integrated natively evaluating function implementations fully?"),
			OptionsID:    j(`["Menyelip manis TEPAT setelah kata func tapi SEBELUM nama fungsinya. Contoh: func (p Kucing) Makan()", "Di baris paling bawah akhir aplikasi", "Terselip dibracket return parameter akhir: func Lari() (b Bebek)", "Sama dengan parameter normal: func Tembak(s Senapan)"]`),
			OptionsEN:    j(`["Succeeding smoothly PRECISELY post preceding 'func' generic invocations inherently occurring IMMEDIATELY PRECEDING function specific nomenclatures explicitly. Example schema: func (p Cat) Eat()", "Concluding universal runtime thread tail bounds absolutely universally linearly", "Nested parsing bounded inside trailing return sequence generic brackets: func Sprint() (b Ducks)", "Simulating explicit symmetrical standalone parametric syntax instances mapping logically accurately: func Fire(s Rifle)"]`),
			CorrectIndex: 0,
			XPReward:     5,
		},

		// Lesson 15: Interfaces (2 quizzes)
		{
			ID:           quizUUID(5, 15, 1),
			LessonID:     lessonUUID(5, 15),
			QuestionID:   "Fitur magis manakah yang memerankan peranan vital mumpuni seolah olah kontrak/perjanjian keharusan bagi objek-objek liar Struct di penjuru map demi lolos seleksi memegang peranan/jabatan yang digariskan tanpa peduli wujud asalnya apa?",
			QuestionEN:   strPtr("Dissect magical intrinsic parametric constructs orchestrating paramount execution limits conceptually mapping exact contractual structural adherence terms orchestrating diverse disjoint wild isolated functional Struct object properties successfully satisfying mandatory implementation validation screening phases assuming specialized generic roles explicitly unconditionally universally detached identifying original abstract root configurations?"),
			OptionsID:    j(`["Interface", "Dependency Injection", "Nested Struct", "Overloading"]`),
			OptionsEN:    j(`["Interfaces", "Dependency Injections", "Nested Structs", "Overloadings"]`),
			CorrectIndex: 0,
			XPReward:     5,
		},
		{
			ID:           quizUUID(5, 15, 2),
			LessonID:     lessonUUID(5, 15),
			QuestionID:   "Apakah golang menuntut kalimat pemaksaan secara tertulis (seperti `implements InterfaceIni`) pada ujung nama Deklarasi Struct demi validasi kontrak di atas?",
			QuestionEN:   strPtr("Does explicit contextual native Go operations demand exact literal syntactical implementation directives explicitly (manifesting archaic `implements ThisInterfaceEntity`) appending explicit declaration suffixes wrapping boundaries validating identical mapping terms implicitly above?"),
			OptionsID:    j(`["Ya, wajib dong. Kan kodingan java gitu", "Tergantung framework Go Fiber atau Go Gin yang jalan", "Iya tapi pakai huruf besar", "TIDAK! Go menganut asas implisit. Asalkan si Struct secara rahasia memoles fitur func dengan nama persis di papan kontrak, Go langsung mengadopsi sah anak itu otomatis"]`),
			OptionsEN:    j(`["Yes, inherently explicitly. Reflecting native generic Java architectures inherently structurally securely", "Relaying context mapped Go Fiber/Gin network framework domains exclusively depending heavily dependencies natively exclusively securely functionally dynamically specifically explicitly implicitly generally strictly purely unconditionally globally locally.", "Indeed albeit mapped entirely uppercase uppercase matrix", "FALSE! Go paradigms dictate implicit adoptions universally fundamentally dynamically effectively efficiently successfully seamlessly securely inherently optimally robustly reliably unconditionally natively purely structurally. Providing internal isolated local specific explicit native target instances mirror functions equivalently mapping exact contractual parameter schema domains identically structurally, compiler infrastructures integrate adoption automatically magically seamlessly"]`),
			CorrectIndex: 3,
			XPReward:     5,
		},

		// Lesson 16: Error Handling (2 quizzes)
		{
			ID:           quizUUID(5, 16, 1),
			LessonID:     lessonUUID(5, 16),
			QuestionID:   "Mekanisme abadi sakti apokaliptik yang menjadi ciri khas dan pembeda Developer Gopher Sejati dengan bahasa lain, di mana mereka menolak menahan sunyi kecelakaan kode dengan try-catch ga jelas?",
			QuestionEN:   strPtr("Evaluating apocalyptic ultimate immortal magical system mechanisms dictating native structural characteristics drastically severing authentic Developer Gopher profiles globally contrasting alternative primitive legacy systems explicitly rejecting silent anomaly suppressions universally integrating cryptic obscure disjoint try-catch anomalies?"),
			OptionsID:    j(`["Pengecekan Explicit Return Error (!= nil) terus menerus berlapis-lapis", "Try and Exception loop tak henti", "Throws exception message ke console.log javascript murahan", "Bakar CPU laptop"]`),
			OptionsEN:    j(`["Continuous multi-layered Explicit Validations checking (err != nil) natively relentlessly universally infinitely dynamically organically implicitly magically", "Endless Try/Exception infinite loops cascading unconditionally locally heavily mapping bounds continuously dynamically randomly specifically securely optimally.", "Throwing exception matrices routing target JavaScript console payload arrays mapping outputs", "Igniting target CPU laptop environments dynamically explicitly locally organically intrinsically inherently"]`),
			CorrectIndex: 0,
			XPReward:     5,
		},
		{
			ID:           quizUUID(5, 16, 2),
			LessonID:     lessonUUID(5, 16),
			QuestionID:   "Jika validasi sukses melewati seleksi blokade tanpa error cacat sama sekali. Benda ghaib apa (Tipe data Void Go) yang selayaknya kita umpan lemparkan balik ke baris pemanggil demi mendinginkan/menentramkan compiler?",
			QuestionEN:   strPtr("Postulating exact validations sequence loops mapping flawless explicit execution clearances systematically bypassing defective matrix bounds cleanly uncorrupted intrinsically—What exact explicit void/null target generic primitive Go configurations structurally dictate mandatory inbound payloads appeasing/calming strict explicit compiler parsing environments inherently globally?"),
			OptionsID:    j(`["Return error(0)", "Melempar Null() kayak Java", "Nil (Kosong mutlak nihil)", "Undefined Error Code X"]`),
			OptionsEN:    j(`["Return error(0)", "Dumping Java-esque explicit Null() formats natively", "Nil (Absolute pure void nothingness structurally flawlessly naturally dynamically)", "Undefined Anomalies Array Error Code X"]`),
			CorrectIndex: 2,
			XPReward:     5,
		},

		// Lesson 17, 18, 19: Standard Libs & Goroutines (4 quizzes)
		{
			ID:           quizUUID(5, 17, 1),
			LessonID:     lessonUUID(5, 17),
			QuestionID:   "Package ajaib apa yang biasa orang impor di atas kodingan khusus jika butuh memformat print/log output string keren di console terminal?",
			QuestionEN:   strPtr("Identifies explicit magical standard execution structural logic package components generally universally mapped prepending execution script configurations traversing specific textual format terminal print logic console arrays mapping target explicit domains explicitly functionally?"),
			OptionsID:    j(`["logz", "crypto/md5", "console_app_web_api", "fmt"]`),
			OptionsEN:    j(`["logz", "crypto/md5", "console_app_web_api", "fmt"]`),
			CorrectIndex: 3,
			XPReward:     5,
		},
		{
			ID:           quizUUID(5, 18, 1),
			LessonID:     lessonUUID(5, 18),
			QuestionID:   "Mesin pemecah belah alam dimensi paralel milik Go yang kekuatannya membuat thread OS laptop terasa layaknya mainan. Hanya mengkonsumsi 2 KB data dan dikontrol murni hanya dengan menaruh 1 keyword prefix sederhana saktinya di depan func... Apakah itu?",
			QuestionEN:   strPtr("Identifying Go's native inter-dimensional parallel matrix fracturing machines scaling functional infrastructure architectures obliterating archaic legacy OS hardware thread environments explicitly parsing payload limits scaling specifically isolated bounds averaging 2 KB parameters securely manipulating components solely incorporating 1 localized magical explicit prefix syntax explicitly natively universally?"),
			OptionsID:    j(`["ParallelMode()", "ThreadStart.func()", "Keyword ajaib awalan ` + "`" + `go func()` + "`" + `", "import JavaMultithreading library"]`),
			OptionsEN:    j(`["ParallelMode()", "ThreadStart.func()", "Magical explicit initialization execution prefixes appending locally dynamically ` + "`" + `go func()` + "`" + ` naturally flawlessly explicitly globally optimally functionally systemically", "import JavaMultithreading ecosystem matrix library abstractions arrays payloads domains structurally internally natively implicitly manually unconditionally fundamentally safely exclusively universally seamlessly securely safely purely accurately"]`),
			CorrectIndex: 2,
			XPReward:     5,
		},
		{
			ID:           quizUUID(5, 19, 1),
			LessonID:     lessonUUID(5, 19),
			QuestionID:   "Agar alam Goroutine sanggup berkomunikasi nyambung dengan alam Utamanya dan bertukar kabar data aman, Go menyodorkan saluran khusus. Apakah sebutannya?",
			QuestionEN:   strPtr("To ensure Goroutines can communicate safely with the main execution stream and transfer data reliably, Go provides a specific pipeline structure. What is it called?"),
			OptionsID:    j(`["Database SQL", "Memcached Redis", "Tipe Data Channel (chan)", "NPM Library Node.js"]`),
			OptionsEN:    j(`["SQL Database", "Memcached Redis", "Channel Data Type (chan)", "Node.js NPM Library"]`),
			CorrectIndex: 2,
			XPReward:     5,
		},
		{
			ID:           quizUUID(5, 19, 2),
			LessonID:     lessonUUID(5, 19),
			QuestionID:   "Bagaimana cara menumpahkan cairan data MASUK KE DALAM LUBANG pipa channel tersebut?",
			QuestionEN:   strPtr("How do we push payload data INSIDE the target channel pipeline?"),
			OptionsID:    j(`["Pakai func Pipa.Push(Data)", "Pakai Arrow Operator PANAH KIRI mengarah ke nama var Pipa: ` + "`" + `PipaDimensi <- DataRahasiaHartaKarun` + "`" + `", "Disiram Pake gayung biasa", "Pakai String format"]`),
			OptionsEN:    j(`["Use Pipe.Push(Data)", "Use LEFT INBOUND arrow referencing the channel variable: ` + "`" + `Pipeline <- Payload` + "`" + `", "Dump primitive bucket", "Use string format"]`),
			CorrectIndex: 1,
			XPReward:     5,
		},
	}

	for _, quiz := range quizzes {
		upsertQuiz(db, &quiz)
	}
}

func strPtr(s string) *string {
	return &s
}
