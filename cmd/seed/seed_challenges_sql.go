package main

import (
	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"gorm.io/gorm"
)

func seedSQLChallenges(db *gorm.DB) {
	si := sqlSchemaInfo()
	siJoin := sqlSchemaInfoJoin()

	challenges := []domain.Challenge{
		sqlEasy1(si), sqlEasy2(si), sqlEasy3(si), sqlEasy4(si), sqlEasy5(si),
		sqlMedium1(siJoin), sqlMedium2(siJoin), sqlMedium3(siJoin), sqlMedium4(si), sqlMedium5(siJoin),
		sqlHard1(siJoin), sqlHard2(siJoin), sqlHard3(siJoin),
	}

	for i := range challenges {
		upsert(db, &challenges[i], challenges[i].Slug)
	}
}

func sqlSchemaInfo() []byte {
	return j(map[string]interface{}{
		"tables": []map[string]interface{}{
			{"name": "produk",
				"columns": []map[string]string{{"name": "produk_id", "type": "INTEGER"}, {"name": "nama_produk", "type": "VARCHAR"}, {"name": "kategori", "type": "VARCHAR"}, {"name": "harga", "type": "INTEGER"}, {"name": "stok", "type": "INTEGER"}},
				"rows": []map[string]interface{}{
					{"produk_id": 1, "nama_produk": "Batik Tulis Solo", "kategori": "Fashion", "harga": 450000, "stok": 25},
					{"produk_id": 2, "nama_produk": "Kopi Toraja", "kategori": "Food", "harga": 85000, "stok": 100},
					{"produk_id": 3, "nama_produk": "Tas Rotan Bali", "kategori": "Craft", "harga": 275000, "stok": 15},
					{"produk_id": 4, "nama_produk": "Sambal Bu Rudy", "kategori": "Food", "harga": 35000, "stok": 200},
					{"produk_id": 5, "nama_produk": "Sepatu Kulit", "kategori": "Fashion", "harga": 320000, "stok": 40},
				}},
		},
	})
}

func sqlSchemaInfoJoin() []byte {
	return j(map[string]interface{}{
		"tables": []map[string]interface{}{
			{"name": "produk",
				"columns": []map[string]string{{"name": "produk_id", "type": "INTEGER"}, {"name": "nama_produk", "type": "VARCHAR"}, {"name": "kategori", "type": "VARCHAR"}, {"name": "harga", "type": "INTEGER"}, {"name": "stok", "type": "INTEGER"}},
				"rows": []map[string]interface{}{
					{"produk_id": 1, "nama_produk": "Batik Tulis Solo", "kategori": "Fashion", "harga": 450000, "stok": 25},
					{"produk_id": 2, "nama_produk": "Kopi Toraja", "kategori": "Food", "harga": 85000, "stok": 100},
					{"produk_id": 3, "nama_produk": "Tas Rotan Bali", "kategori": "Craft", "harga": 275000, "stok": 15},
					{"produk_id": 4, "nama_produk": "Sambal Bu Rudy", "kategori": "Food", "harga": 35000, "stok": 200},
					{"produk_id": 5, "nama_produk": "Sepatu Kulit", "kategori": "Fashion", "harga": 320000, "stok": 40},
				}},
			{"name": "pelanggan",
				"columns": []map[string]string{{"name": "pelanggan_id", "type": "INTEGER"}, {"name": "nama_pelanggan", "type": "VARCHAR"}, {"name": "kota", "type": "VARCHAR"}},
				"rows": []map[string]interface{}{
					{"pelanggan_id": 1, "nama_pelanggan": "Rina Susanti", "kota": "Jakarta"},
					{"pelanggan_id": 2, "nama_pelanggan": "Budi Prakoso", "kota": "Surabaya"},
					{"pelanggan_id": 3, "nama_pelanggan": "Dewi Lestari", "kota": "Bandung"},
					{"pelanggan_id": 4, "nama_pelanggan": "Ahmad Fauzi", "kota": "Jakarta"},
				}},
			{"name": "transaksi",
				"columns": []map[string]string{{"name": "transaksi_id", "type": "INTEGER"}, {"name": "pelanggan_id", "type": "INTEGER"}, {"name": "produk_id", "type": "INTEGER"}, {"name": "jumlah", "type": "INTEGER"}, {"name": "tanggal", "type": "DATE"}},
				"rows": []map[string]interface{}{
					{"transaksi_id": 1, "pelanggan_id": 1, "produk_id": 1, "jumlah": 2, "tanggal": "2024-01-15"},
					{"transaksi_id": 2, "pelanggan_id": 2, "produk_id": 2, "jumlah": 5, "tanggal": "2024-01-16"},
					{"transaksi_id": 3, "pelanggan_id": 1, "produk_id": 4, "jumlah": 10, "tanggal": "2024-02-01"},
					{"transaksi_id": 4, "pelanggan_id": 3, "produk_id": 1, "jumlah": 1, "tanggal": "2024-02-10"},
					{"transaksi_id": 5, "pelanggan_id": 4, "produk_id": 5, "jumlah": 1, "tanggal": "2024-02-15"},
				}},
		},
	})
}

func tc(pairs ...map[string]interface{}) []byte { return j(pairs) }
func eo(v interface{}) []byte                   { return j(v) }

func sqlEasy1(si []byte) domain.Challenge {
	return domain.Challenge{ID: challengeUUID(1), Slug: "hari-pertama-di-tokonusantara", Language: "sql", Difficulty: "easy", XPReward: 10, OrderIndex: 1,
		TitleID: "Hari Pertama di TokoNusantara", TitleEN: sp("First Day at TokoNusantara"),
		StoryID: "Selamat datang di TokoNusantara! 🎉 Hari ini hari pertamamu sebagai Data Analyst Junior.\n\nPak Budi: \"Coba tampilkan semua data dari tabel produk — cara terbaik mulai mengenal bisnis kita.\"",
		StoryEN: sp("Welcome to TokoNusantara! 🎉 Today is your first day as Junior Data Analyst.\n\nBudi: \"Try displaying all data from the produk table — best way to get familiar with our business.\""),
		TaskID:  "Tampilkan semua data dari tabel `produk`.", TaskEN: sp("Display all data from the `produk` table."),
		HintID: sp("Gunakan SELECT * untuk semua kolom."), HintEN: sp("Use SELECT * for all columns."),
		SchemaInfo: si, StarterCode: sp("-- Tulis query SQL kamu di sini\nSELECT"), SolutionCode: sp("SELECT * FROM produk;"),
		TestCases:      tc(map[string]interface{}{"description": "5 baris", "type": "row_count", "expected": 5}),
		ExpectedOutput: eo(map[string]interface{}{"row_count": 5})}
}

func sqlEasy2(si []byte) domain.Challenge {
	return domain.Challenge{ID: challengeUUID(2), Slug: "cari-produk-fashion", Language: "sql", Difficulty: "easy", XPReward: 10, OrderIndex: 2,
		TitleID: "Temukan Produk Fashion", TitleEN: sp("Find Fashion Products"),
		StoryID: "Tim marketing butuh daftar semua produk kategori Fashion untuk kampanye khusus.",
		StoryEN: sp("The marketing team needs all Fashion category products for a special campaign."),
		TaskID:  "Tampilkan semua produk dengan kategori 'Fashion'.", TaskEN: sp("Display all products with category 'Fashion'."),
		HintID: sp("Gunakan WHERE kategori = 'Fashion'."), HintEN: sp("Use WHERE kategori = 'Fashion'."),
		SchemaInfo: si, StarterCode: sp("-- Filter produk kategori Fashion\nSELECT"), SolutionCode: sp("SELECT * FROM produk WHERE kategori = 'Fashion';"),
		TestCases:      tc(map[string]interface{}{"description": "2 baris Fashion", "type": "row_count", "expected": 2}),
		ExpectedOutput: eo(map[string]interface{}{"row_count": 2})}
}

func sqlEasy3(si []byte) domain.Challenge {
	return domain.Challenge{ID: challengeUUID(3), Slug: "produk-termahal", Language: "sql", Difficulty: "easy", XPReward: 10, OrderIndex: 3,
		TitleID: "Produk Paling Mahal", TitleEN: sp("The Most Expensive Product"),
		StoryID: "Tim finance butuh tahu produk termahal di katalog kita.",
		StoryEN: sp("Finance team needs to know the most expensive product in our catalog."),
		TaskID:  "Tampilkan 1 produk dengan harga tertinggi.", TaskEN: sp("Display the 1 product with the highest price."),
		HintID: sp("ORDER BY harga DESC LIMIT 1."), HintEN: sp("ORDER BY harga DESC LIMIT 1."),
		SchemaInfo: si, StarterCode: sp("-- Produk termahal\nSELECT"), SolutionCode: sp("SELECT * FROM produk ORDER BY harga DESC LIMIT 1;"),
		TestCases:      tc(map[string]interface{}{"description": "1 baris", "type": "row_count", "expected": 1}, map[string]interface{}{"description": "harga 450000", "type": "cell_value", "row": 0, "col": "harga", "expected": 450000}),
		ExpectedOutput: eo(map[string]interface{}{"row_count": 1})}
}

func sqlEasy4(si []byte) domain.Challenge {
	return domain.Challenge{ID: challengeUUID(4), Slug: "stok-menipis", Language: "sql", Difficulty: "easy", XPReward: 10, OrderIndex: 4,
		TitleID: "Stok Hampir Habis!", TitleEN: sp("Stock Running Low!"),
		StoryID: "Gudang bilang beberapa produk stoknya sangat sedikit. Cari semua di bawah 30 unit!",
		StoryEN: sp("Warehouse says some products are low. Find all below 30 units!"),
		TaskID:  "Tampilkan produk dengan stok < 30.", TaskEN: sp("Display products with stock < 30."),
		HintID: sp("WHERE stok < 30."), HintEN: sp("WHERE stok < 30."),
		SchemaInfo: si, StarterCode: sp("-- Stok rendah\nSELECT"), SolutionCode: sp("SELECT * FROM produk WHERE stok < 30;"),
		TestCases:      tc(map[string]interface{}{"description": "2 baris", "type": "row_count", "expected": 2}),
		ExpectedOutput: eo(map[string]interface{}{"row_count": 2})}
}

func sqlEasy5(si []byte) domain.Challenge {
	return domain.Challenge{ID: challengeUUID(5), Slug: "rekap-per-kategori", Language: "sql", Difficulty: "easy", XPReward: 15, OrderIndex: 5,
		TitleID: "Rekap Produk per Kategori", TitleEN: sp("Product Summary by Category"),
		StoryID: "CEO ingin tahu berapa jumlah produk di setiap kategori.",
		StoryEN: sp("CEO wants to know how many products in each category."),
		TaskID:  "Tampilkan jumlah produk per kategori.", TaskEN: sp("Display product count per category."),
		HintID: sp("GROUP BY kategori, COUNT(*)."), HintEN: sp("GROUP BY kategori, COUNT(*)."),
		SchemaInfo: si, StarterCode: sp("-- Jumlah per kategori\nSELECT"), SolutionCode: sp("SELECT kategori, COUNT(*) AS jumlah_produk FROM produk GROUP BY kategori;"),
		TestCases:      tc(map[string]interface{}{"description": "has kategori", "type": "has_column", "expected": "kategori"}),
		ExpectedOutput: eo(map[string]interface{}{"row_count": 3})}
}

func sqlMedium1(si []byte) domain.Challenge {
	return domain.Challenge{ID: challengeUUID(6), Slug: "laporan-penjualan-bergabung", Language: "sql", Difficulty: "medium", XPReward: 25, OrderIndex: 6,
		TitleID: "Laporan Penjualan — Siapa Beli Apa?", TitleEN: sp("Sales Report — Who Bought What?"),
		StoryID: "Gabungkan tabel pelanggan, transaksi, dan produk untuk laporan lengkap.",
		StoryEN: sp("Join pelanggan, transaksi, and produk tables for a full report."),
		TaskID:  "Tampilkan nama_pelanggan, nama_produk, jumlah.", TaskEN: sp("Display nama_pelanggan, nama_produk, jumlah."),
		HintID: sp("INNER JOIN dua kali."), HintEN: sp("Use INNER JOIN twice."),
		SchemaInfo: si, StarterCode: sp("-- JOIN tiga tabel\nSELECT\nFROM transaksi t"),
		SolutionCode:   sp("SELECT p.nama_pelanggan, pr.nama_produk, t.jumlah\nFROM transaksi t\nINNER JOIN pelanggan p ON t.pelanggan_id = p.pelanggan_id\nINNER JOIN produk pr ON t.produk_id = pr.produk_id;"),
		TestCases:      tc(map[string]interface{}{"description": "has nama_pelanggan", "type": "has_column", "expected": "nama_pelanggan"}),
		ExpectedOutput: eo(map[string]interface{}{"row_count": 5})}
}

func sqlMedium2(si []byte) domain.Challenge {
	return domain.Challenge{ID: challengeUUID(7), Slug: "pelanggan-terbaik", Language: "sql", Difficulty: "medium", XPReward: 25, OrderIndex: 7,
		TitleID: "Siapa Pelanggan Terbaik?", TitleEN: sp("Who Are Our Best Customers?"),
		StoryID: "Cari 5 pelanggan dengan total belanja (harga × jumlah) tertinggi.",
		StoryEN: sp("Find the 5 customers with highest total spending (price × quantity)."),
		TaskID:  "Top 5 pelanggan total belanja terbesar.", TaskEN: sp("Top 5 customers by total spending."),
		HintID: sp("SUM(harga * jumlah), GROUP BY, ORDER BY DESC LIMIT 5."), HintEN: sp("SUM(harga * jumlah), GROUP BY, ORDER BY DESC LIMIT 5."),
		SchemaInfo: si, StarterCode: sp("-- Top 5 pelanggan\nSELECT"),
		SolutionCode:   sp("SELECT p.nama_pelanggan, SUM(pr.harga * t.jumlah) AS total_belanja\nFROM transaksi t\nINNER JOIN pelanggan p ON t.pelanggan_id = p.pelanggan_id\nINNER JOIN produk pr ON t.produk_id = pr.produk_id\nGROUP BY p.nama_pelanggan\nORDER BY total_belanja DESC LIMIT 5;"),
		TestCases:      tc(map[string]interface{}{"description": "has total_belanja", "type": "has_column", "expected": "total_belanja"}),
		ExpectedOutput: eo(map[string]interface{}{"row_count": 4})}
}

func sqlMedium3(si []byte) domain.Challenge {
	return domain.Challenge{ID: challengeUUID(8), Slug: "produk-belum-terjual", Language: "sql", Difficulty: "medium", XPReward: 25, OrderIndex: 8,
		TitleID: "Produk Belum Pernah Terjual", TitleEN: sp("Products Never Sold"),
		StoryID: "Cari produk yang tidak pernah ada di tabel transaksi.",
		StoryEN: sp("Find products that never appear in the transaksi table."),
		TaskID:  "Tampilkan produk tanpa transaksi.", TaskEN: sp("Display products with no transactions."),
		HintID: sp("LEFT JOIN lalu WHERE IS NULL."), HintEN: sp("LEFT JOIN then WHERE IS NULL."),
		SchemaInfo: si, StarterCode: sp("-- Produk tanpa transaksi\nSELECT"),
		SolutionCode:   sp("SELECT pr.* FROM produk pr LEFT JOIN transaksi t ON pr.produk_id = t.produk_id WHERE t.produk_id IS NULL;"),
		TestCases:      tc(map[string]interface{}{"description": "2 baris", "type": "row_count", "expected": 2}),
		ExpectedOutput: eo(map[string]interface{}{"row_count": 2})}
}

func sqlMedium4(si []byte) domain.Challenge {
	return domain.Challenge{ID: challengeUUID(9), Slug: "rata-rata-harga-per-kategori", Language: "sql", Difficulty: "medium", XPReward: 20, OrderIndex: 9,
		TitleID: "Produk di Atas Rata-rata", TitleEN: sp("Products Above Category Average"),
		StoryID: "Cari produk yang harganya di atas rata-rata kategorinya.",
		StoryEN: sp("Find products priced above their category average."),
		TaskID:  "Produk dengan harga > AVG(harga) per kategori.", TaskEN: sp("Products with price > AVG(price) per category."),
		HintID: sp("Correlated subquery + AVG."), HintEN: sp("Correlated subquery + AVG."),
		SchemaInfo: si, StarterCode: sp("-- Di atas rata-rata\nSELECT"),
		SolutionCode:   sp("SELECT p.* FROM produk p WHERE p.harga > (SELECT AVG(p2.harga) FROM produk p2 WHERE p2.kategori = p.kategori);"),
		TestCases:      tc(map[string]interface{}{"description": "ada hasil", "type": "row_count", "expected": 3}),
		ExpectedOutput: eo(map[string]interface{}{"row_count": 3})}
}

func sqlMedium5(si []byte) domain.Challenge {
	return domain.Challenge{ID: challengeUUID(10), Slug: "transaksi-bulan-ini", Language: "sql", Difficulty: "medium", XPReward: 25, OrderIndex: 10,
		TitleID: "Transaksi per Kota", TitleEN: sp("Transactions by City"),
		StoryID: "Ringkasan transaksi per kota, urutkan dari total penjualan tertinggi.",
		StoryEN: sp("Transaction summary per city, ordered by highest total sales."),
		TaskID:  "Kota, jumlah transaksi, total penjualan.", TaskEN: sp("City, transaction count, total sales."),
		HintID: sp("JOIN 3 tabel, GROUP BY kota."), HintEN: sp("JOIN 3 tables, GROUP BY city."),
		SchemaInfo: si, StarterCode: sp("-- Per kota\nSELECT"),
		SolutionCode:   sp("SELECT p.kota, COUNT(t.transaksi_id) AS jumlah_transaksi, SUM(pr.harga * t.jumlah) AS total_penjualan\nFROM transaksi t INNER JOIN pelanggan p ON t.pelanggan_id = p.pelanggan_id\nINNER JOIN produk pr ON t.produk_id = pr.produk_id\nGROUP BY p.kota ORDER BY total_penjualan DESC;"),
		TestCases:      tc(map[string]interface{}{"description": "has kota", "type": "has_column", "expected": "kota"}),
		ExpectedOutput: eo(map[string]interface{}{"row_count": 3})}
}

func sqlHard1(si []byte) domain.Challenge {
	return domain.Challenge{ID: challengeUUID(11), Slug: "ranking-produk-per-kategori", Language: "sql", Difficulty: "hard", XPReward: 50, OrderIndex: 11,
		TitleID: "Ranking Terlaris per Kategori", TitleEN: sp("Top-Selling Ranking per Category"),
		StoryID: "Tampilkan top 3 produk terlaris per kategori menggunakan window function.",
		StoryEN: sp("Show top 3 best-selling products per category using window functions."),
		TaskID:  "Top 3 per kategori dengan ROW_NUMBER.", TaskEN: sp("Top 3 per category with ROW_NUMBER."),
		HintID: sp("ROW_NUMBER() OVER (PARTITION BY ...) lalu WHERE rn <= 3."), HintEN: sp("ROW_NUMBER() OVER (PARTITION BY ...) then WHERE rn <= 3."),
		SchemaInfo: si, StarterCode: sp("-- Top 3 per kategori\nWITH penjualan AS ("),
		SolutionCode:   sp("WITH penjualan AS (\n  SELECT pr.nama_produk, pr.kategori, COALESCE(SUM(t.jumlah),0) AS total_terjual\n  FROM produk pr LEFT JOIN transaksi t ON pr.produk_id = t.produk_id\n  GROUP BY pr.produk_id, pr.nama_produk, pr.kategori\n), ranked AS (\n  SELECT *, ROW_NUMBER() OVER (PARTITION BY kategori ORDER BY total_terjual DESC) AS rn FROM penjualan\n)\nSELECT nama_produk, kategori, total_terjual, rn AS ranking FROM ranked WHERE rn <= 3 ORDER BY kategori, rn;"),
		TestCases:      tc(map[string]interface{}{"description": "has ranking", "type": "has_column", "expected": "ranking"}),
		ExpectedOutput: eo(map[string]interface{}{"row_count": 5})}
}

func sqlHard2(si []byte) domain.Challenge {
	return domain.Challenge{ID: challengeUUID(12), Slug: "analisis-pertumbuhan-bulanan", Language: "sql", Difficulty: "hard", XPReward: 50, OrderIndex: 12,
		TitleID: "Pertumbuhan Penjualan Bulanan", TitleEN: sp("Monthly Sales Growth"),
		StoryID: "Total penjualan per bulan + perbandingan bulan sebelumnya + persentase pertumbuhan.",
		StoryEN: sp("Monthly total sales + comparison with previous month + growth percentage."),
		TaskID:  "Total per bulan, prev month, diff, persen.", TaskEN: sp("Monthly total, prev, diff, percentage."),
		HintID: sp("DATE_TRUNC + LAG() window function."), HintEN: sp("DATE_TRUNC + LAG() window function."),
		SchemaInfo: si, StarterCode: sp("-- Pertumbuhan\nWITH monthly AS ("),
		SolutionCode:   sp("WITH monthly AS (\n  SELECT DATE_TRUNC('month', t.tanggal) AS bulan, SUM(pr.harga * t.jumlah) AS total\n  FROM transaksi t INNER JOIN produk pr ON t.produk_id = pr.produk_id GROUP BY 1\n)\nSELECT bulan, total, LAG(total) OVER (ORDER BY bulan) AS prev,\n  total - LAG(total) OVER (ORDER BY bulan) AS diff\nFROM monthly ORDER BY bulan;"),
		TestCases:      tc(map[string]interface{}{"description": "has bulan", "type": "has_column", "expected": "bulan"}),
		ExpectedOutput: eo(map[string]interface{}{"row_count": 2})}
}

func sqlHard3(si []byte) domain.Challenge {
	return domain.Challenge{ID: challengeUUID(13), Slug: "cohort-pelanggan-baru", Language: "sql", Difficulty: "hard", XPReward: 60, OrderIndex: 13,
		TitleID: "Cohort Analysis Pelanggan", TitleEN: sp("Customer Cohort Analysis"),
		StoryID: "Kelompokkan pelanggan berdasarkan bulan pertama kali mereka transaksi (cohort).",
		StoryEN: sp("Group customers by the month of their first transaction (cohort)."),
		TaskID:  "Cohort bulan pertama + jumlah pelanggan baru.", TaskEN: sp("First month cohort + new customer count."),
		HintID: sp("MIN(tanggal) per pelanggan, DATE_TRUNC, COUNT."), HintEN: sp("MIN(tanggal) per customer, DATE_TRUNC, COUNT."),
		SchemaInfo: si, StarterCode: sp("-- Cohort\nWITH first_purchase AS ("),
		SolutionCode:   sp("WITH first_purchase AS (\n  SELECT pelanggan_id, DATE_TRUNC('month', MIN(tanggal)) AS cohort_bulan\n  FROM transaksi GROUP BY pelanggan_id\n)\nSELECT cohort_bulan, COUNT(pelanggan_id) AS jumlah_pelanggan_baru\nFROM first_purchase GROUP BY cohort_bulan ORDER BY cohort_bulan;"),
		TestCases:      tc(map[string]interface{}{"description": "has cohort_bulan", "type": "has_column", "expected": "cohort_bulan"}),
		ExpectedOutput: eo(map[string]interface{}{"row_count": 2})}
}
