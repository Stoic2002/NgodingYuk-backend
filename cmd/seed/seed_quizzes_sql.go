package main

import (
	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"gorm.io/gorm"
)

func seedSQLQuizzes(db *gorm.DB) {
	quizzes := []domain.LessonQuiz{
		// sql-beginner lesson 2 (SELECT Dasar)
		{ID: quizUUID(4, 2, 1), LessonID: lessonUUID(4, 2), OrderIndex: 1, XPReward: 5, CorrectIndex: 0,
			QuestionID: "Perintah SQL apa yang digunakan untuk mengambil data?", QuestionEN: sp("Which SQL command is used to retrieve data?"),
			OptionsID:     j([]string{"SELECT", "FETCH", "GET", "READ"}),
			OptionsEN:     j([]string{"SELECT", "FETCH", "GET", "READ"}),
			ExplanationID: sp("SELECT adalah perintah DML untuk mengambil data dari tabel."),
			ExplanationEN: sp("SELECT is the DML command for retrieving data from a table.")},
		{ID: quizUUID(4, 2, 2), LessonID: lessonUUID(4, 2), OrderIndex: 2, XPReward: 5, CorrectIndex: 2,
			QuestionID: "Apa arti SELECT * FROM produk?", QuestionEN: sp("What does SELECT * FROM produk mean?"),
			OptionsID:     j([]string{"Hapus semua dari produk", "Hitung jumlah produk", "Ambil semua kolom dari tabel produk", "Buat tabel produk"}),
			OptionsEN:     j([]string{"Delete all from produk", "Count all products", "Get all columns from produk table", "Create produk table"}),
			ExplanationID: sp("Tanda * berarti semua kolom, sehingga SELECT * FROM produk mengambil seluruh data dari tabel produk."),
			ExplanationEN: sp("The * symbol means all columns, so SELECT * FROM produk retrieves all data from the produk table.")},
		{ID: quizUUID(4, 2, 3), LessonID: lessonUUID(4, 2), OrderIndex: 3, XPReward: 5, CorrectIndex: 1,
			QuestionID: "Keyword apa untuk menampilkan nilai unik saja?", QuestionEN: sp("Which keyword shows only unique values?"),
			OptionsID:     j([]string{"UNIQUE", "DISTINCT", "DIFFERENT", "SINGLE"}),
			OptionsEN:     j([]string{"UNIQUE", "DISTINCT", "DIFFERENT", "SINGLE"}),
			ExplanationID: sp("DISTINCT menghilangkan baris duplikat dari hasil query."),
			ExplanationEN: sp("DISTINCT removes duplicate rows from query results.")},

		// sql-beginner lesson 3 (WHERE)
		{ID: quizUUID(4, 3, 1), LessonID: lessonUUID(4, 3), OrderIndex: 1, XPReward: 5, CorrectIndex: 3,
			QuestionID: "Operator mana yang digunakan untuk logika 'DAN' di SQL?", QuestionEN: sp("Which operator is used for logical 'AND' in SQL?"),
			OptionsID:     j([]string{"&", "&&", "+", "AND"}),
			OptionsEN:     j([]string{"&", "&&", "+", "AND"}),
			ExplanationID: sp("SQL menggunakan keyword AND untuk operasi logika AND."),
			ExplanationEN: sp("SQL uses the AND keyword for logical AND operations.")},
		{ID: quizUUID(4, 3, 2), LessonID: lessonUUID(4, 3), OrderIndex: 2, XPReward: 5, CorrectIndex: 0,
			QuestionID: "Apa fungsi klausa WHERE?", QuestionEN: sp("What does the WHERE clause do?"),
			OptionsID:     j([]string{"Memfilter baris berdasarkan kondisi", "Mengurutkan data", "Mengelompokkan data", "Membatasi jumlah baris"}),
			OptionsEN:     j([]string{"Filter rows based on conditions", "Sort data", "Group data", "Limit row count"}),
			ExplanationID: sp("WHERE digunakan untuk memfilter baris yang memenuhi kondisi tertentu."),
			ExplanationEN: sp("WHERE is used to filter rows that meet certain conditions.")},

		// sql-beginner lesson 4 (ORDER BY, LIMIT)
		{ID: quizUUID(4, 4, 1), LessonID: lessonUUID(4, 4), OrderIndex: 1, XPReward: 5, CorrectIndex: 1,
			QuestionID: "Apa arti DESC pada ORDER BY?", QuestionEN: sp("What does DESC mean in ORDER BY?"),
			OptionsID:     j([]string{"Ascending (naik)", "Descending (turun)", "Description", "Default"}),
			OptionsEN:     j([]string{"Ascending", "Descending", "Description", "Default"}),
			ExplanationID: sp("DESC = descending, mengurutkan dari nilai tertinggi ke terendah."),
			ExplanationEN: sp("DESC = descending, sorts from highest to lowest value.")},
		{ID: quizUUID(4, 4, 2), LessonID: lessonUUID(4, 4), OrderIndex: 2, XPReward: 5, CorrectIndex: 2,
			QuestionID: "Kombinasi LIMIT + OFFSET berguna untuk apa?", QuestionEN: sp("What is LIMIT + OFFSET useful for?"),
			OptionsID:     j([]string{"Sorting", "Filtering", "Pagination", "Joining"}),
			OptionsEN:     j([]string{"Sorting", "Filtering", "Pagination", "Joining"}),
			ExplanationID: sp("LIMIT + OFFSET digunakan untuk implementasi pagination (halaman-halaman data)."),
			ExplanationEN: sp("LIMIT + OFFSET is used to implement pagination (pages of data).")},

		// sql-beginner lesson 8 (DDL)
		{ID: quizUUID(4, 8, 1), LessonID: lessonUUID(4, 8), OrderIndex: 1, XPReward: 5, CorrectIndex: 1,
			QuestionID: "Constraint mana yang memastikan tidak ada nilai duplikat?", QuestionEN: sp("Which constraint ensures no duplicate values?"),
			OptionsID:     j([]string{"NOT NULL", "UNIQUE", "PRIMARY KEY", "DEFAULT"}),
			OptionsEN:     j([]string{"NOT NULL", "UNIQUE", "PRIMARY KEY", "DEFAULT"}),
			ExplanationID: sp("UNIQUE memastikan semua nilai dalam kolom tersebut unik (tidak ada duplikat)."),
			ExplanationEN: sp("UNIQUE ensures all values in the column are unique (no duplicates).")},

		// sql-intermediate lesson 4 (GROUP BY)
		{ID: quizUUID(5, 4, 1), LessonID: lessonUUID(5, 4), OrderIndex: 1, XPReward: 5, CorrectIndex: 1,
			QuestionID: "Apa perbedaan WHERE dan HAVING?", QuestionEN: sp("What is the difference between WHERE and HAVING?"),
			OptionsID:     j([]string{"Tidak ada perbedaan", "WHERE filter sebelum GROUP BY, HAVING setelah", "HAVING lebih cepat", "WHERE hanya untuk angka"}),
			OptionsEN:     j([]string{"No difference", "WHERE filters before GROUP BY, HAVING after", "HAVING is faster", "WHERE is only for numbers"}),
			ExplanationID: sp("WHERE memfilter baris sebelum pengelompokan, HAVING memfilter group setelah pengelompokan."),
			ExplanationEN: sp("WHERE filters rows before grouping, HAVING filters groups after grouping.")},

		// sql-intermediate lesson 5 (INNER JOIN)
		{ID: quizUUID(5, 5, 1), LessonID: lessonUUID(5, 5), OrderIndex: 1, XPReward: 5, CorrectIndex: 2,
			QuestionID: "INNER JOIN mengembalikan baris yang...?", QuestionEN: sp("INNER JOIN returns rows that...?"),
			OptionsID:     j([]string{"Ada di tabel kiri saja", "Ada di tabel kanan saja", "Ada kecocokan di kedua tabel", "Semua baris dari kedua tabel"}),
			OptionsEN:     j([]string{"Exist in left table only", "Exist in right table only", "Have a match in both tables", "All rows from both tables"}),
			ExplanationID: sp("INNER JOIN hanya mengembalikan baris yang memiliki kecocokan di kedua tabel yang di-join."),
			ExplanationEN: sp("INNER JOIN only returns rows that have a match in both joined tables.")},

		// sql-intermediate lesson 6 (LEFT/RIGHT JOIN)
		{ID: quizUUID(5, 6, 1), LessonID: lessonUUID(5, 6), OrderIndex: 1, XPReward: 5, CorrectIndex: 0,
			QuestionID: "Apa yang terjadi jika baris dari tabel kiri tidak punya pasangan di LEFT JOIN?", QuestionEN: sp("What happens if a left table row has no match in a LEFT JOIN?"),
			OptionsID:     j([]string{"Baris tetap muncul, kolom kanan diisi NULL", "Baris dihapus", "Error", "Baris diabaikan"}),
			OptionsEN:     j([]string{"Row still appears, right columns filled with NULL", "Row is deleted", "Error", "Row is ignored"}),
			ExplanationID: sp("LEFT JOIN menampilkan semua baris dari tabel kiri, dengan NULL untuk kolom tabel kanan yang tidak cocok."),
			ExplanationEN: sp("LEFT JOIN shows all rows from the left table, with NULL for non-matching right table columns.")},

		// sql-advanced lesson 1 (Window Functions)
		{ID: quizUUID(6, 1, 1), LessonID: lessonUUID(6, 1), OrderIndex: 1, XPReward: 5, CorrectIndex: 3,
			QuestionID: "Apa perbedaan window function dengan GROUP BY?", QuestionEN: sp("What's the difference between window functions and GROUP BY?"),
			OptionsID:     j([]string{"Tidak ada perbedaan", "Window function lebih lambat", "GROUP BY menggunakan OVER", "Window function tidak mengurangi jumlah baris"}),
			OptionsEN:     j([]string{"No difference", "Window functions are slower", "GROUP BY uses OVER", "Window functions don't reduce row count"}),
			ExplanationID: sp("Perbedaan utama: GROUP BY mengurangi baris menjadi 1 per group, window function mempertahankan semua baris."),
			ExplanationEN: sp("Key difference: GROUP BY reduces rows to 1 per group, window functions keep all rows.")},

		// sql-advanced lesson 2 (ROW_NUMBER/RANK)
		{ID: quizUUID(6, 2, 1), LessonID: lessonUUID(6, 2), OrderIndex: 1, XPReward: 5, CorrectIndex: 1,
			QuestionID: "Apa perbedaan RANK dan DENSE_RANK?", QuestionEN: sp("What's the difference between RANK and DENSE_RANK?"),
			OptionsID:     j([]string{"Tidak ada", "RANK loncat nomor setelah tie, DENSE_RANK tidak", "DENSE_RANK lebih lambat", "RANK hanya untuk angka"}),
			OptionsEN:     j([]string{"None", "RANK skips numbers after ties, DENSE_RANK doesn't", "DENSE_RANK is slower", "RANK is only for numbers"}),
			ExplanationID: sp("RANK: 1,2,2,4 (loncat). DENSE_RANK: 1,2,2,3 (tidak loncat)."),
			ExplanationEN: sp("RANK: 1,2,2,4 (skips). DENSE_RANK: 1,2,2,3 (no gaps).")},

		// sql-advanced lesson 4 (CTE)
		{ID: quizUUID(6, 4, 1), LessonID: lessonUUID(6, 4), OrderIndex: 1, XPReward: 5, CorrectIndex: 0,
			QuestionID: "Keyword apa yang digunakan untuk membuat CTE?", QuestionEN: sp("Which keyword is used to create a CTE?"),
			OptionsID:     j([]string{"WITH", "AS", "CTE", "DEFINE"}),
			OptionsEN:     j([]string{"WITH", "AS", "CTE", "DEFINE"}),
			ExplanationID: sp("CTE dibuat dengan keyword WITH: WITH nama_cte AS (SELECT ...)"),
			ExplanationEN: sp("CTEs are created with the WITH keyword: WITH cte_name AS (SELECT ...)")},
	}

	for i := range quizzes {
		upsertQuiz(db, &quizzes[i])
	}
}
