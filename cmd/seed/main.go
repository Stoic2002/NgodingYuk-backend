package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/arulkarim/ngodingyuk-server/internal/config"
	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Helper function to create a string pointer
func sp(s string) *string { return &s }

// Helper to marshal JSON
func j(v interface{}) json.RawMessage {
	b, _ := json.Marshal(v)
	return b
}

// Hardcoded UUIDs for courses (stable references)
var (
	CourseGoBeg  = uuid.MustParse("a0000000-0000-0000-0000-000000000001")
	CourseGoMid  = uuid.MustParse("a0000000-0000-0000-0000-000000000002")
	CourseGoAdv  = uuid.MustParse("a0000000-0000-0000-0000-000000000003")
	CourseSQLBeg = uuid.MustParse("a0000000-0000-0000-0000-000000000004")
	CourseSQLMid = uuid.MustParse("a0000000-0000-0000-0000-000000000005")
	CourseSQLAdv = uuid.MustParse("a0000000-0000-0000-0000-000000000006")
)

// Lesson UUID helper — generates deterministic UUID from course + order
func lessonUUID(courseIdx, lessonIdx int) uuid.UUID {
	return uuid.MustParse(fmt.Sprintf("b%07d-0000-0000-0000-%012d", courseIdx, lessonIdx))
}

// Quiz UUID helper
func quizUUID(courseIdx, lessonIdx, quizIdx int) uuid.UUID {
	return uuid.MustParse(fmt.Sprintf("c%07d-%04d-0000-0000-%012d", courseIdx, lessonIdx, quizIdx))
}

// Challenge UUID helper
func challengeUUID(idx int) uuid.UUID {
	return uuid.MustParse(fmt.Sprintf("d0000000-0000-0000-0000-%012d", idx))
}

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Seeding courses...")
	seedCourses(db)

	log.Println("Seeding Go lessons...")
	seedGoLessons(db)

	log.Println("Seeding SQL lessons...")
	seedSQLLessons(db)

	log.Println("Seeding Go quizzes...")
	seedGoQuizzes(db)

	log.Println("Seeding SQL quizzes...")
	seedSQLQuizzes(db)

	log.Println("Seeding SQL challenges...")
	seedSQLChallenges(db)

	log.Println("Seeding Go challenges...")
	seedGoChallenges(db)
	seedGoChallengesExp(db)

	log.Println("Seeding Expanded SQL challenges...")
	seedSQLChallengesExp(db)

	log.Println("Seeding Expanded Go Quizzes...")
	seedGoQuizzesExp(db)

	log.Println("Seeding Expanded SQL Quizzes...")
	seedSQLQuizzesExp(db)

	log.Println("✅ Seeding complete!")
}

func upsert(db *gorm.DB, model interface{}, slug string) {
	var count int64
	db.Model(model).Where("slug = ?", slug).Count(&count)
	if count == 0 {
		if err := db.Create(model).Error; err != nil {
			log.Printf("  ⚠ Failed to seed (slug=%s): %v\n", slug, err)
		} else {
			log.Printf("  ✓ Seeded: %s\n", slug)
		}
	} else {
		log.Printf("  · Skipped (exists): %s\n", slug)
	}
}

func upsertLesson(db *gorm.DB, lesson *domain.Lesson) {
	var count int64
	db.Model(lesson).Where("id = ?", lesson.ID).Count(&count)
	if count == 0 {
		if err := db.Create(lesson).Error; err != nil {
			log.Printf("  ⚠ Failed to seed lesson (id=%s): %v\n", lesson.ID, err)
		} else {
			log.Printf("  ✓ Seeded lesson: %s\n", lesson.TitleID)
		}
	} else {
		log.Printf("  · Skipped lesson (exists): %s\n", lesson.TitleID)
	}
}

func upsertQuiz(db *gorm.DB, quiz *domain.LessonQuiz) {
	var count int64
	db.Model(quiz).Where("id = ?", quiz.ID).Count(&count)
	if count == 0 {
		if err := db.Create(quiz).Error; err != nil {
			log.Printf("  ⚠ Failed to seed quiz (id=%s): %v\n", quiz.ID, err)
		}
	}
}
