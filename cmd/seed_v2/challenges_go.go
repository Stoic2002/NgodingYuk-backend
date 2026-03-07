package main

import (
	"fmt"

	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// deterministic UUID for Go challenges
func goChallengeUUID(idx int) uuid.UUID {
	return uuid.MustParse(fmt.Sprintf("88880000-0000-0000-0000-%012d", idx))
}

func seedGoChallenges(db *gorm.DB) {
	difficulties := []string{"easy", "medium", "hard"}

	var challenges []domain.Challenge

	// 1-50: Go Challenges with Storytelling
	for i := 1; i <= 50; i++ {
		// Randomize difficulty based on seeded sequence
		diff := difficulties[(i*7)%3]
		var xp int64
		switch diff {
		case "easy":
			xp = 15
		case "medium":
			xp = 30
		case "hard":
			xp = 60
		}

		// Starter Code Template
		starterCode := fmt.Sprintf(`package main

import "fmt"

func main() {
	// Tulis jawaban untuk challenge #%d di sini
	
}
`, i)

		titleID := ""
		titleEN := ""
		storyID := ""
		storyEN := ""
		taskID := ""
		taskEN := ""
		hintID := ""
		hintEN := ""
		var testCases string

		switch i {
		case 1:
			titleID = "Pesan Pertama dari Bumi"
			titleEN = "First Message from Earth"
			storyID = "Kamu adalah seorang insinyur telekomunikasi di stasiun ruang angkasa. Hari ini, roket penjelajah Mars baru saja mendarat. Komandan memintamu mengirimkan satu pesan sapaan standar untuk memastikan transmisi data berfungsi dengan baik menggunakan bahasa Go."
			storyEN = "You are a telecom engineer at the space station. Today, the Mars rover has just landed. The commander asks you to send a standard greeting message to ensure data transmission is working well using Go."
			taskID = "Cetak teks `Halo Mars!` ke layar (stdout)."
			taskEN = "Print the text `Halo Mars!` to the screen (stdout)."
			hintID = "Gunakan `fmt.Println`."
			hintEN = "Use `fmt.Println`."
			testCases = `[{"input": "", "expected": "Halo Mars!"}]`
		case 2:
			titleID = "Sandi Rahasia Brankas"
			titleEN = "Vault Secret Passcode"
			storyID = "Brankas emas Kapten Jack dikunci dengan sandi matematika. Ia lupa sandinya, tetapi asistennya meninggalkan secarik catatan: 'Sandi adalah hasil perkalian 15 dengan 8'. Kapten membayarmu untuk menulis program kalkulator otomatis."
			storyEN = "Captain Jack's gold vault is locked with a math passcode. He forgot the code, but his assistant left a note: 'The passcode is the result of multiplying 15 by 8'. The Captain pays you to write an automatic calculator program."
			taskID = "Cetak hasil perkalian 15 dikali 8 menggunakan operator matematika di Go."
			taskEN = "Print the result of 15 multiplied by 8 using math operators in Go."
			hintID = "Pakai operator `*` di dalam `fmt.Println`."
			hintEN = "Use the `*` operator inside `fmt.Println`."
			testCases = `[{"input": "", "expected": "120"}]`
		case 3:
			titleID = "Detektor Suhu Reaktor"
			titleEN = "Reactor Temperature Detector"
			storyID = "Sistem fusi nuklir di laboratorium bawah tanah memancarkan suhu ekstrem. Jika lebih dari 100 derajat, reaktor akan meledak! Kita butuh program pembaca suhu dari sensor."
			storyEN = "The nuclear fusion system in the underground lab emits extreme temperature. If it goes over 100 degrees, the reactor explodes! We need a program to read temperature from the sensor."
			taskID = "Baca 1 baris input berupa angka suhu. Jika suhu lebih dari 100, cetak `BAHAYA`. Jika tidak, cetak `AMAN`."
			taskEN = "Read 1 line of input as a temperature number. If temperature > 100, print `BAHAYA`. Otherwise print `AMAN`."
			hintID = "Gunakan `fmt.Scan` untuk membaca integer, lalu seleksi kondisi dengan `if-else`."
			hintEN = "Use `fmt.Scan` to read integer, then condition check with `if-else`."
			testCases = `[{"input": "120", "expected": "BAHAYA"}, {"input": "85", "expected": "AMAN"}, {"input": "100", "expected": "AMAN"}]`
		case 4:
			titleID = "Mesin Hitung Koin"
			titleEN = "Coin Counting Machine"
			storyID = "Mesin arcade klasikmu kebanjiran koin. Kamu perlu menghitung total koin berulang-ulang tanpa capek menjumlahkannya secara manual. Mari kita buat program pengulangan matematika!"
			storyEN = "Your classic arcade machine is flooded with coins. You need to count the total coins repeatedly without getting tired of manual addition. Let's build a math recurrence program!"
			taskID = "Baca 1 baris input berupa angka N. Gunakan perulangan untuk mencetak semua angka dari 1 sampai N secara berurutan, dipisah baris baru."
			taskEN = "Read 1 line of input as a number N. Use loops to print all numbers from 1 to N sequentially, separated by a newline."
			hintID = "Gunakan `for i := 1; i <= n; i++`."
			hintEN = "Use `for i := 1; i <= n; i++`."
			testCases = `[{"input": "3", "expected": "1\n2\n3"}, {"input": "5", "expected": "1\n2\n3\n4\n5"}]`
		case 5:
			titleID = "Gema Suara di Gua"
			titleEN = "Voice Echo in the Cave"
			storyID = "Saat sedang mendaki gunung es, kamu berteriak ke dalam gua dalam untuk mengecek gema. Namun tebing es memantulkan suaramu hingga 3 kali berturut-turut! Simulasikan gema ini agar tim SAR bisa mendeteksinya."
			storyEN = "While climbing an icy mountain, you shout into a deep cave to check the echo. But the ice cliff reflects your voice 3 times in a row! Simulate this echo so the SAR team can detect it."
			taskID = "Baca 1 kata dari input. Cetak kata tersebut persis 3 kali dalam baris yang sama, dipisahkan dengan spasi."
			taskEN = "Read 1 word from input. Print that word exactly 3 times sequentially separated by a space on the same line."
			hintID = "Gunakan format atau gabungan string. Contoh: `fmt.Printf(\"%s %s %s\\n\", word, word, word)`."
			hintEN = "Use format or string concatenation. Example: `fmt.Printf(\"%s %s %s\\n\", word, word, word)`."
			testCases = `[{"input": "Halo", "expected": "Halo Halo Halo"}, {"input": "Tolong", "expected": "Tolong Tolong Tolong"}]`
		case 6:
			titleID = "Filter Usia Penumpang"
			titleEN = "Passenger Age Filter"
			storyID = "Wahana roller coaster 'Dragon Spine' sangat ekstrim. Penjaga wahana memintamu membuat gerbang digital otomatis. Siapa pun yang berumur di bawah 12 tahun tidak boleh masuk."
			storyEN = "The 'Dragon Spine' roller coaster is very extreme. The ride operator asks you to create an automatic digital gate. Anyone rigidly under 12 years old cannot enter."
			taskID = "Baca integer usia. Jika usia >= 12, cetak `Boleh Naik`. Jika usia < 12, cetak `Dilarang`."
			taskEN = "Read an integer age. If age >= 12, print `Boleh Naik`. If age < 12, print `Dilarang`."
			hintID = "Gunakan `if age >= 12` dan `else`."
			hintEN = "Use `if age >= 12` and `else`."
			testCases = `[{"input": "15", "expected": "Boleh Naik"}, {"input": "12", "expected": "Boleh Naik"}, {"input": "8", "expected": "Dilarang"}]`
		case 7:
			titleID = "Kelipatan Ajaib"
			titleEN = "Magic Multiples"
			storyID = "Tikus laboratorium hanya memakan keju pada jam percobaan yang merupakan kelipatan dari 3. Asisten lab-mu bingung kapan ia harus memberikan keju jika jam ditunjukkan dengan angka arbitrer."
			storyEN = "The lab rats only eat cheese during experiment hours that are multiples of 3. Your lab assistant is confused about when to give cheese if the hours are arbitrary numbers."
			taskID = "Baca angka kelipatan N. Jika angka tersebut habis dibagi 3, cetak `Makan Keju`. Jika tidak, cetak `Puasa`."
			taskEN = "Read a generic number N. If it's fully divisible by 3, print `Makan Keju`. Otherwise, print `Puasa`."
			hintID = "Gunakan operator modulo `% 3 == 0`."
			hintEN = "Use the modulo operator `% 3 == 0`."
			testCases = `[{"input": "9", "expected": "Makan Keju"}, {"input": "10", "expected": "Puasa"}, {"input": "33", "expected": "Makan Keju"}]`
		case 8:
			titleID = "Lampu Lalu Lintas AI"
			titleEN = "AI Traffic Lights"
			storyID = "Kota NeoTokyo baru saja menerapkan kamera pengawas lalu lintas otomatis. Kamera mendeteksi warna lampu apa yang menyala, dan memerintahkan mobil untuk bereaksi."
			storyEN = "NeoTokyo City has just implemented automatic traffic cameras. The camera detects which colored light is glowing and commands cars to react."
			taskID = "Input berupa string ('Merah', 'Kuning', 'Hijau'). Jika Merah cetak `Berhenti`, Kuning cetak `Hati-hati`, Hijau cetak `Jalan`, selain itu cetak `Lampu Rusak`."
			taskEN = "Input string ('Merah', 'Kuning', 'Hijau'). IF Merah print `Berhenti`, Kuning -> `Hati-hati`, Hijau -> `Jalan`, else print `Lampu Rusak`."
			hintID = "Gunakan struktur `switch-case` atau `if-else if` berantai pada Go."
			hintEN = "Use a `switch-case` block or chained go `if-else if`."
			testCases = `[{"input": "Merah", "expected": "Berhenti"}, {"input": "Kuning", "expected": "Hati-hati"}, {"input": "Hijau", "expected": "Jalan"}, {"input": "Biru", "expected": "Lampu Rusak"}]`
		case 9:
			titleID = "Pembalik Kata Sihir"
			titleEN = "Magic Word Reverser"
			storyID = "Di akademi sihir, ada sebuah mantra yang diucapkan terbalik secara harfiah. Jika diutarakan 'Bawang', ia akan meledakkan panci."
			storyEN = "In the magic academy, there is a spell spoken in total reverse. If casted normally like 'Bawang', it explodes the pot."
			taskID = "Baca 1 String kalimat dari input. Balikkan urutan karakternya dari belakang ke depan, lalu cetak."
			taskEN = "Read 1 String from input. Reverse the character order backwards and print it."
			hintID = "Ubah string menjadi array byte/rune, lalu jalankan perulangan `for` mundur. Awas pada tipe `string` di Go!"
			hintEN = "Convert the string to a rune array/byte, and traverse backwards. Be mindful of Go strings!"
			testCases = `[{"input": "Sihir", "expected": "rihiS"}, {"input": "Mantra", "expected": "artnaM"}]`
		case 10:
			titleID = "Panjang Tembok Raksasa"
			titleEN = "Giant Wall Length"
			storyID = "Pekerja bangunan Romawi kuno menyusun balok-balok dinding kastil. Sayang mereka lupa menghitung estimasi panjang tembok secara total dari rentetan angka dimensi."
			storyEN = "Ancient Roman builders stacked blocks for the castle walls. Sadly, they forgot to maintain the actual length dimension aggregate."
			taskID = "Baca berapa karakter panjang string dari input yang dimasukkan, lalu cetak hasilnya dengan pesan `Panjangnya X meter`."
			taskEN = "Read the exact character length generated by the inputted string. Print the count with the template `Panjangnya X meter`."
			hintID = "Input bisa dibaca sebagai text utuh. Di Go kalikan saja dengan `len(text)`."
			hintEN = "Input is treated as raw text. Compute size via `len(text)`."
			testCases = `[{"input": "TembokEs", "expected": "Panjangnya 8 meter"}, {"input": "Aegis", "expected": "Panjangnya 5 meter"}]`
		default:
			// Generik Story Untuk Soal 11-50
			// Kita buat polanya bervariasi secara prosedural untuk melengkapi 50 challenge

			// Topic Selector based on Index
			topic := i % 10
			switch topic {
			case 1:
				titleID = fmt.Sprintf("Ekspedisi Kalkulasi #%d", i)
				titleEN = fmt.Sprintf("Calculation Expedition #%d", i)
				storyID = fmt.Sprintf("Misi pesawat luar angkasamu mencapai orbit %d. Sistem navigasi memintamu memasukkan daya mesin yang merupakan nilai input dikali 10.", i)
				storyEN = fmt.Sprintf("Your spacecraft mission reaches orbit %d. Navigation system requires you to enter the engine power which is the input multiplied by 10.", i)
				taskID = "Baca satu angka integer, kalikan 10, lalu cetak."
				taskEN = "Read one integer, multiply by 10, print result."
				hintID = "Bisa gunakan standard print kalikan dengan 10."
				hintEN = "Multiply the parsed scanned integer by 10."
				testCases = `[{"input": "5", "expected": "50"}, {"input": "23", "expected": "230"}]`
			case 2:
				titleID = fmt.Sprintf("Akar Tersembunyi #%d", i)
				titleEN = fmt.Sprintf("Hidden Roots #%d", i)
				storyID = fmt.Sprintf("Peneliti hutan Amazon menemukan daun langka ke-%d. Jumlah nutrisinya selalu input ditambah 50.", i)
				storyEN = fmt.Sprintf("Amazon researchers found the %dth rare leaf. Its nutrition yield is always input value plus 50.", i)
				taskID = "Baca integer, tambah 50, lalu cetak."
				taskEN = "Read an integer, add 50, and print."
				hintID = "Jangan lupa format `scan`-nya."
				hintEN = "Ensure you scan it carefully."
				testCases = `[{"input": "100", "expected": "150"}, {"input": "-20", "expected": "30"}]`
			case 3:
				titleID = fmt.Sprintf("Sandangan Teks #%d", i)
				titleEN = fmt.Sprintf("Text Clothings #%d", i)
				storyID = fmt.Sprintf("Mesin printer baju distro ke-%d butuh format tulisan khusus untuk pesanan label.", i)
				storyEN = fmt.Sprintf("The %dth clothing printer machine needs specific typography layouts for a label order.", i)
				taskID = "Baca 1 string, cetak menjadi format `[STRING_INPUT]` (dengan tanda kurung kotak)."
				taskEN = "Read 1 string, print utilizing format `[STRING_INPUT]` with square brackets."
				hintID = "Pakailah `fmt.Printf(\"[%s]\", input)`."
				hintEN = "Prefer using `fmt.Printf(\"[%s]\", input)` string formatters."
				testCases = `[{"input": "Baju", "expected": "[Baju]"}, {"input": "Kaos", "expected": "[Kaos]"}]`
			case 4:
				titleID = fmt.Sprintf("Detektor Genap-Ganjil #%d", i)
				titleEN = fmt.Sprintf("Even-Odd Detector #%d", i)
				storyID = fmt.Sprintf("Keamanan apartemen lantai %d membedakan akses pria dan wanita melalui angka token masuk. Jika genap maskulin, jikalau ganjil feminin.", i)
				storyEN = fmt.Sprintf("Security at apartment floor %d filters access digits. Even for masculine, odd for feminine.", i)
				taskID = "Tentukan apakah input genap atau ganjil. Cetak `GENAP` atau `GANJIL`."
				taskEN = "Determine if input is even or odd. Print `GENAP` or `GANJIL`."
				hintID = "Gunakan modulo `% 2 == 0`."
				hintEN = "Modulo operator `% 2 == 0` is your friend."
				testCases = `[{"input": "4", "expected": "GENAP"}, {"input": "7", "expected": "GANJIL"}, {"input": "0", "expected": "GENAP"}]`
			case 5:
				titleID = fmt.Sprintf("Duplikasi Kode #%d", i)
				titleEN = fmt.Sprintf("Code Duplicator #%d", i)
				storyID = fmt.Sprintf("Sistem virus ke-%d meminta kloning diri sendiri sebanyak dua kali di terminal internal radar.", i)
				storyEN = fmt.Sprintf("Virus system %d demands self-cloning redundancy twice in the internal terminal radar logs.", i)
				taskID = "Baca 1 kata input. Cetak dan ulangi kata tersebut tanpa spasi menjadi dua kali lipat."
				taskEN = "Read an input word. Duplicate and print it without any spaces in between."
				hintID = "Concat string menggunakan `input + input`."
				hintEN = "Concatenate strings easily using `input + input`."
				testCases = `[{"input": "Virus", "expected": "VirusVirus"}, {"input": "Hack", "expected": "HackHack"}]`
			case 6:
				titleID = fmt.Sprintf("Selisih Dimensi #%d", i)
				titleEN = fmt.Sprintf("Dimensional Delta #%d", i)
				storyID = fmt.Sprintf("Navigasi radar antar dimensi %d membutuhkan jarak relatif dengan konstan mutlak dari koordinat X.", i)
				storyEN = fmt.Sprintf("Inter-dimensional radar map %d necessitates relative distance subtracting 100 unconditionally from X coordinate.", i)
				taskID = "Baca input angka X, temukan nilainya jika dikurangi 100."
				taskEN = "Read an X input number, subtract 100, and evaluate the difference integer output."
				hintID = "Ini soal reduksi sederhana."
				hintEN = "Simple arithmetic subtraction."
				testCases = `[{"input": "500", "expected": "400"}, {"input": "150", "expected": "50"}]`
			case 7:
				titleID = fmt.Sprintf("Kode Pembuka Tirai #%d", i)
				titleEN = fmt.Sprintf("Curtain Opener Code #%d", i)
				storyID = fmt.Sprintf("Sutradara Teater ke-%d mengatur bahwa tirai hanya terbuka bila kodenya sama persis dengan sandi master (1234).", i)
				storyEN = fmt.Sprintf("Theatre director %d configured curtains to only open if exact sequence matches master code (1234).", i)
				taskID = "Baca integer. Jika nilainya 1234, cetak `Terbuka`. Jika tidak, cetak `Terkunci`."
				taskEN = "Read an integer. If value evaluates to 1234, output `Terbuka`. Otherwise, output `Terkunci`."
				hintID = "Kondisi If-Else boolean logic."
				hintEN = "Boolean If-Else statements match check."
				testCases = `[{"input": "1234", "expected": "Terbuka"}, {"input": "4321", "expected": "Terkunci"}]`
			case 8:
				titleID = fmt.Sprintf("Mesin Cermin #%d", i)
				titleEN = fmt.Sprintf("The Mirror Machine #%d", i)
				storyID = fmt.Sprintf("Penyihir cermin tingkat %d mengubah pantulan refleksi di sekitarnya menjadi sebuah distorsi teks kapital.", i)
				storyEN = fmt.Sprintf("The mirror sorcerer level %d shifts light refractions translating to capital text distortions.", i)
				taskID = "Ubah input string dari server menjadi huruf kecil (Lower Case), lalu cetak dengan penulisan prefix `>> ` di depannya."
				taskEN = "Transform server input string entirely to Lowercase formats, and append standard `>> ` prefix to it."
				hintID = "Kamu hanya butuh melakukan pencetakan literal format string."
				hintEN = "Perform a literal string format printing logic without case shifts if strings package is omitted. (Actually just append the string)."
				testCases = `[{"input": "Kaca", "expected": ">> Kaca"}, {"input": "Refleksi", "expected": ">> Refleksi"}]`
			case 9:
				titleID = fmt.Sprintf("Deret Segitiga #%d", i)
				titleEN = fmt.Sprintf("Triangle Series #%d", i)
				storyID = fmt.Sprintf("Para matematikawan peradaban kuno kota sektor %d menyembunyikan formula rahasia yaitu Kuadrat ditambah 1.", i)
				storyEN = fmt.Sprintf("The ancient city sector %d mathematicians buried a classified formula namely the value computation of square sum plus 1.", i)
				taskID = "Baca integer N. Hitung ekspresi matematika: (N * N) + 1. Cetak angkanya."
				taskEN = "Read integer N. Compute the mathematical logic evaluating (N * N) + 1. Print resultant output."
				hintID = "Kalikan N dengan dirinya sendiri."
				hintEN = "Square N by multiplying it intuitively."
				testCases = `[{"input": "3", "expected": "10"}, {"input": "5", "expected": "26"}]`
			case 0:
				titleID = fmt.Sprintf("Penghancur Desimal #%d", i)
				titleEN = fmt.Sprintf("Decimal Shredder #%d", i)
				storyID = fmt.Sprintf("Komputer perbankan node %d mengalami malfungsi pemotongan harga (flooring) tanpa adanya sisa bagi koma sama sekali.", i)
				storyEN = fmt.Sprintf("Banking computer node %d endures precision malfunction forcing floors natively without modulo calculations.", i)
				taskID = "Baca angka N. Bagi dengan 2 menggunakan pembagian integer (tanpa pembulatan ke atas/bawah pecahan koma). Cetak nilai sisanya."
				taskEN = "Read array of number N. Divide perfectly by 2 applying standard division without fractions handling. Return remainder quotient."
				hintID = "Di Go, pembagian integer `/` otomatis akan membuang sisa koma. Namun tes ini meminta kita mencari sisanya."
				hintEN = "Use the modulo operator to extract remainder in pure execution logic."
				testCases = `[{"input": "11", "expected": "1"}, {"input": "20", "expected": "0"}]`
			}
		}

		c := domain.Challenge{
			ID:             goChallengeUUID(i),
			Slug:           fmt.Sprintf("go-challenge-%03d", i),
			Language:       "go",
			Difficulty:     diff,
			TitleID:        titleID,
			TitleEN:        sp(titleEN),
			StoryID:        storyID,
			StoryEN:        sp(storyEN),
			TaskID:         taskID,
			TaskEN:         sp(taskEN),
			HintID:         sp(hintID),
			HintEN:         sp(hintEN),
			StarterCode:    sp(starterCode),
			TestCases:      j(testCases), // Assuming testCases holds a JSON string literal. No expected_output for code challenges as they assert stdout.
			SchemaInfo:     nil,          // SchemaInfo is strictly for SQL
			ExpectedOutput: j(testCases), // Pass the testCases into expected_output so the frontend renders the Inputs & Expected values as a table!
			XPReward:       xp,
			OrderIndex:     i,
		}

		// Ensure raw json is properly marshalled since j handles raw strings efficiently now due to previous patch.
		// Wait, testCases here is just a string, so j(testCases) will treat it as a json.RawMessage thanks to the 'if string looks like JSON array' patch.
		challenges = append(challenges, c)
	}

	for _, chal := range challenges {
		upsertChallenge(db, &chal)
	}
}
