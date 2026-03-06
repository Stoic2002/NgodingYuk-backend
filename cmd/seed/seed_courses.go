package main

import (
	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"gorm.io/gorm"
)

func seedCourses(db *gorm.DB) {
	courses := []domain.Course{
		{
			ID: CourseGoBeg, Slug: "golang-beginner", Language: "golang", Level: "beginner",
			TitleID: "Golang untuk Pemula", TitleEN: sp("Golang for Beginners"),
			DescriptionID: sp("Pelajari dasar-dasar bahasa Go dari variabel, tipe data, hingga struct dan method."),
			DescriptionEN: sp("Learn Go fundamentals from variables, data types, to structs and methods."),
			OrderIndex:    1,
		},
		{
			ID: CourseGoMid, Slug: "golang-intermediate", Language: "golang", Level: "intermediate",
			TitleID: "Golang Tingkat Menengah", TitleEN: sp("Golang Intermediate"),
			DescriptionID: sp("Kuasai pointer, interface, goroutine, channel, dan testing di Go."),
			DescriptionEN: sp("Master pointers, interfaces, goroutines, channels, and testing in Go."),
			OrderIndex:    2,
		},
		{
			ID: CourseGoAdv, Slug: "golang-advanced", Language: "golang", Level: "advanced",
			TitleID: "Golang Tingkat Lanjut", TitleEN: sp("Golang Advanced"),
			DescriptionID: sp("Pelajari generics, reflection, design patterns, REST API, dan deployment."),
			DescriptionEN: sp("Learn generics, reflection, design patterns, REST APIs, and deployment."),
			OrderIndex:    3,
		},
		{
			ID: CourseSQLBeg, Slug: "sql-beginner", Language: "sql", Level: "beginner",
			TitleID: "Belajar SQL dari Nol", TitleEN: sp("Learn SQL from Scratch"),
			DescriptionID: sp("Mulai dari SELECT dasar hingga DDL — fondasi untuk bekerja dengan database."),
			DescriptionEN: sp("Start from basic SELECT to DDL — the foundation for working with databases."),
			OrderIndex:    4,
		},
		{
			ID: CourseSQLMid, Slug: "sql-intermediate", Language: "sql", Level: "intermediate",
			TitleID: "SQL Tingkat Menengah", TitleEN: sp("SQL Intermediate"),
			DescriptionID: sp("Kuasai JOIN, subquery, aggregate function, dan filtering lanjutan."),
			DescriptionEN: sp("Master JOIN, subqueries, aggregate functions, and advanced filtering."),
			OrderIndex:    5,
		},
		{
			ID: CourseSQLAdv, Slug: "sql-advanced", Language: "sql", Level: "advanced",
			TitleID: "SQL Tingkat Lanjut", TitleEN: sp("SQL Advanced"),
			DescriptionID: sp("Pelajari window function, CTE, indexing, transaksi, dan desain database."),
			DescriptionEN: sp("Learn window functions, CTEs, indexing, transactions, and database design."),
			OrderIndex:    6,
		},
	}

	for i := range courses {
		upsert(db, &courses[i], courses[i].Slug)
	}
}
