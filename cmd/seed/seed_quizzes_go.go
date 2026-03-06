package main

import (
	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"gorm.io/gorm"
)

func seedGoQuizzes(db *gorm.DB) {
	quizzes := []domain.LessonQuiz{
		// golang-beginner lesson 1 (Pengenalan Go)
		{ID: quizUUID(1, 1, 1), LessonID: lessonUUID(1, 1), OrderIndex: 1, XPReward: 5, CorrectIndex: 1,
			QuestionID: "Siapa yang mengembangkan bahasa Go?", QuestionEN: sp("Who developed the Go programming language?"),
			OptionsID:     j([]string{"Microsoft", "Google", "Facebook", "Apple"}),
			OptionsEN:     j([]string{"Microsoft", "Google", "Facebook", "Apple"}),
			ExplanationID: sp("Go dikembangkan oleh Google pada tahun 2009 oleh Robert Griesemer, Rob Pike, dan Ken Thompson."),
			ExplanationEN: sp("Go was developed by Google in 2009 by Robert Griesemer, Rob Pike, and Ken Thompson.")},
		{ID: quizUUID(1, 1, 2), LessonID: lessonUUID(1, 1), OrderIndex: 2, XPReward: 5, CorrectIndex: 2,
			QuestionID: "Apa fungsi dari `fmt.Println()`?", QuestionEN: sp("What does `fmt.Println()` do?"),
			OptionsID:     j([]string{"Membaca input", "Menghapus file", "Mencetak output ke terminal", "Membuat variabel"}),
			OptionsEN:     j([]string{"Read input", "Delete a file", "Print output to terminal", "Create a variable"}),
			ExplanationID: sp("fmt.Println() mencetak output ke terminal dan menambahkan newline di akhir."),
			ExplanationEN: sp("fmt.Println() prints output to the terminal and adds a newline at the end.")},
		{ID: quizUUID(1, 1, 3), LessonID: lessonUUID(1, 1), OrderIndex: 3, XPReward: 5, CorrectIndex: 0,
			QuestionID: "Apa nama file yang biasa digunakan sebagai entry point program Go?", QuestionEN: sp("What file is typically used as the entry point for a Go program?"),
			OptionsID:     j([]string{"main.go", "index.go", "app.go", "start.go"}),
			OptionsEN:     j([]string{"main.go", "index.go", "app.go", "start.go"}),
			ExplanationID: sp("Konvensinya file utama diberi nama main.go dengan package main dan fungsi main()."),
			ExplanationEN: sp("By convention, the main file is named main.go with package main and a main() function.")},

		// golang-beginner lesson 2 (Variabel & Tipe Data)
		{ID: quizUUID(1, 2, 1), LessonID: lessonUUID(1, 2), OrderIndex: 1, XPReward: 5, CorrectIndex: 2,
			QuestionID: "Apa zero value dari tipe data string di Go?", QuestionEN: sp("What is the zero value of the string type in Go?"),
			OptionsID:     j([]string{"nil", "null", `""`, "0"}),
			OptionsEN:     j([]string{"nil", "null", `""`, "0"}),
			ExplanationID: sp(`Zero value dari string di Go adalah string kosong "". Setiap tipe punya zero value default.`),
			ExplanationEN: sp(`The zero value of string in Go is an empty string "". Every type has a default zero value.`)},
		{ID: quizUUID(1, 2, 2), LessonID: lessonUUID(1, 2), OrderIndex: 2, XPReward: 5, CorrectIndex: 1,
			QuestionID: "Manakah cara short declaration variabel di Go?", QuestionEN: sp("Which is the short declaration syntax in Go?"),
			OptionsID:     j([]string{"var x = 10", "x := 10", "let x = 10", "int x = 10"}),
			OptionsEN:     j([]string{"var x = 10", "x := 10", "let x = 10", "int x = 10"}),
			ExplanationID: sp("Operator := adalah short declaration yang paling umum dipakai di Go."),
			ExplanationEN: sp("The := operator is the most common short declaration in Go.")},
		{ID: quizUUID(1, 2, 3), LessonID: lessonUUID(1, 2), OrderIndex: 3, XPReward: 5, CorrectIndex: 3,
			QuestionID: "Apa zero value dari tipe bool?", QuestionEN: sp("What is the zero value of the bool type?"),
			OptionsID:     j([]string{"nil", "0", "true", "false"}),
			OptionsEN:     j([]string{"nil", "0", "true", "false"}),
			ExplanationID: sp("Zero value bool adalah false."),
			ExplanationEN: sp("The zero value of bool is false.")},

		// golang-beginner lesson 5 (Control Flow)
		{ID: quizUUID(1, 5, 1), LessonID: lessonUUID(1, 5), OrderIndex: 1, XPReward: 5, CorrectIndex: 0,
			QuestionID: "Apakah switch di Go perlu break?", QuestionEN: sp("Does switch in Go need a break statement?"),
			OptionsID:     j([]string{"Tidak, otomatis berhenti", "Ya, wajib", "Tergantung versi Go", "Hanya untuk default"}),
			OptionsEN:     j([]string{"No, it stops automatically", "Yes, it's required", "Depends on Go version", "Only for default"}),
			ExplanationID: sp("Switch di Go otomatis berhenti setelah case yang cocok dieksekusi, tanpa perlu break."),
			ExplanationEN: sp("Switch in Go automatically stops after the matching case executes, no break needed.")},
		{ID: quizUUID(1, 5, 2), LessonID: lessonUUID(1, 5), OrderIndex: 2, XPReward: 5, CorrectIndex: 2,
			QuestionID: "Fitur unik apa yang dimiliki if di Go?", QuestionEN: sp("What unique feature does if have in Go?"),
			OptionsID:     j([]string{"Bisa tanpa kurung kurawal", "Tidak butuh kondisi", "Bisa deklarasi variabel di dalamnya", "Bisa return otomatis"}),
			OptionsEN:     j([]string{"Can work without braces", "Doesn't need conditions", "Can declare variables inside it", "Auto-returns"}),
			ExplanationID: sp("Go memungkinkan deklarasi variabel di dalam statement if, contoh: if val := calc(); val > 0 { ... }"),
			ExplanationEN: sp("Go allows variable declaration inside if statements, e.g.: if val := calc(); val > 0 { ... }")},

		// golang-beginner lesson 7 (Fungsi)
		{ID: quizUUID(1, 7, 1), LessonID: lessonUUID(1, 7), OrderIndex: 1, XPReward: 5, CorrectIndex: 1,
			QuestionID: "Apa keunikan fungsi di Go dibanding bahasa lain?", QuestionEN: sp("What makes Go functions unique compared to other languages?"),
			OptionsID:     j([]string{"Tidak butuh return", "Bisa return multiple values", "Harus selalu punya parameter", "Bersifat private by default"}),
			OptionsEN:     j([]string{"No return needed", "Can return multiple values", "Must always have parameters", "Private by default"}),
			ExplanationID: sp("Go mendukung multiple return values — fitur yang sangat berguna untuk error handling."),
			ExplanationEN: sp("Go supports multiple return values — very useful for error handling.")},
		{ID: quizUUID(1, 7, 2), LessonID: lessonUUID(1, 7), OrderIndex: 2, XPReward: 5, CorrectIndex: 3,
			QuestionID: "Apa itu variadic function?", QuestionEN: sp("What is a variadic function?"),
			OptionsID:     j([]string{"Fungsi tanpa return", "Fungsi rekursif", "Fungsi anonim", "Fungsi dengan jumlah parameter dinamis"}),
			OptionsEN:     j([]string{"Function without return", "Recursive function", "Anonymous function", "Function with variable number of parameters"}),
			ExplanationID: sp("Variadic function menerima jumlah parameter tak terbatas menggunakan ... (contoh: func sum(nums ...int))"),
			ExplanationEN: sp("Variadic functions accept unlimited parameters using ... (e.g.: func sum(nums ...int))")},

		// golang-beginner lesson 10 (Struct & Method)
		{ID: quizUUID(1, 10, 1), LessonID: lessonUUID(1, 10), OrderIndex: 1, XPReward: 5, CorrectIndex: 0,
			QuestionID: "Apa itu pointer receiver pada method?", QuestionEN: sp("What is a pointer receiver on a method?"),
			OptionsID:     j([]string{"Receiver yang bisa mengubah nilai field", "Receiver yang lebih cepat", "Receiver untuk interface", "Receiver tanpa parameter"}),
			OptionsEN:     j([]string{"A receiver that can modify field values", "A faster receiver", "A receiver for interfaces", "A receiver without parameters"}),
			ExplanationID: sp("Pointer receiver (*T) memungkinkan method mengubah nilai field dari struct."),
			ExplanationEN: sp("Pointer receiver (*T) allows the method to modify the struct's field values.")},

		// golang-intermediate lesson 2 (Interface)
		{ID: quizUUID(2, 2, 1), LessonID: lessonUUID(2, 2), OrderIndex: 1, XPReward: 5, CorrectIndex: 3,
			QuestionID: "Apa yang dimaksud empty interface di Go?", QuestionEN: sp("What is an empty interface in Go?"),
			OptionsID:     j([]string{"Interface yang tidak bisa diimplementasi", "Interface tanpa method", "Interface hanya untuk struct", "Interface yang menerima semua tipe (interface{} / any)"}),
			OptionsEN:     j([]string{"An interface that can't be implemented", "An interface without methods", "An interface only for structs", "An interface that accepts all types (interface{} / any)"}),
			ExplanationID: sp("Empty interface (interface{} atau any) tidak punya method sehingga semua tipe otomatis mengimplementasinya."),
			ExplanationEN: sp("Empty interface (interface{} or any) has no methods so all types automatically implement it.")},

		// golang-intermediate lesson 4 (Goroutine)
		{ID: quizUUID(2, 4, 1), LessonID: lessonUUID(2, 4), OrderIndex: 1, XPReward: 5, CorrectIndex: 2,
			QuestionID: "Apa fungsi WaitGroup di Go?", QuestionEN: sp("What is WaitGroup used for in Go?"),
			OptionsID:     j([]string{"Membuat goroutine baru", "Membatalkan goroutine", "Menunggu goroutine selesai", "Mengunci data"}),
			OptionsEN:     j([]string{"Create new goroutines", "Cancel goroutines", "Wait for goroutines to complete", "Lock data"}),
			ExplanationID: sp("sync.WaitGroup digunakan untuk menunggu sekelompok goroutine selesai sebelum melanjutkan eksekusi."),
			ExplanationEN: sp("sync.WaitGroup is used to wait for a group of goroutines to finish before continuing execution.")},

		// golang-intermediate lesson 5 (Channel)
		{ID: quizUUID(2, 5, 1), LessonID: lessonUUID(2, 5), OrderIndex: 1, XPReward: 5, CorrectIndex: 1,
			QuestionID: "Apa perbedaan buffered dan unbuffered channel?", QuestionEN: sp("What is the difference between buffered and unbuffered channels?"),
			OptionsID:     j([]string{"Tidak ada perbedaan", "Buffered channel punya kapasitas, unbuffered langsung blocking", "Unbuffered lebih cepat", "Buffered hanya untuk string"}),
			OptionsEN:     j([]string{"No difference", "Buffered has capacity, unbuffered blocks immediately", "Unbuffered is faster", "Buffered is only for strings"}),
			ExplanationID: sp("Unbuffered channel langsung blocking saat send/receive, buffered channel bisa menampung N item sebelum blocking."),
			ExplanationEN: sp("Unbuffered channels block immediately on send/receive, buffered channels can hold N items before blocking.")},
	}

	for i := range quizzes {
		upsertQuiz(db, &quizzes[i])
	}
}
