package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/arulkarim/ngodingyuk-server/internal/config"
	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Helper function to create a string pointer
func sp(s string) *string { return &s }

func j(v interface{}) json.RawMessage {
	if s, ok := v.(string); ok && len(s) > 0 && (s[0] == '{' || s[0] == '[') {
		return json.RawMessage(s)
	}
	b, err := json.Marshal(v)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v, error: %v", v, err)
	}
	return b
}

// Hardcoded UUID for the SQL Beginner course
var (
	CourseSQLBeg = uuid.MustParse("a0000000-0000-0000-0000-000000000004")
)

// Lesson UUID helper — generates deterministic UUID from course + order
func lessonUUID(courseIdx, lessonIdx int) uuid.UUID {
	return uuid.MustParse(fmt.Sprintf("e%07d-0000-0000-0000-%012d", courseIdx, lessonIdx))
}

// Quiz UUID helper
func quizUUID(courseIdx, lessonIdx, quizIdx int) uuid.UUID {
	return uuid.MustParse(fmt.Sprintf("f%07d-%04d-0000-0000-%012d", courseIdx, lessonIdx, quizIdx))
}

func main() {
	cfg := config.Load()
	db := config.InitDB(cfg.DatabaseURL)

	log.Println("Seeding SQL Beginner Course (v2)...")
	seedSQLBeginnerCourse(db)

	log.Println("Seeding SQL Beginner Quizzes (v2)...")
	seedSQLBeginnerQuizzes(db)

	log.Println("Seeding SQL Challenges (v2)...")
	seedSQLChallenges(db)

	log.Println("✅ Seeding v2 complete!")
}

func upsertCourse(db *gorm.DB, course *domain.Course) {
	var count int64
	db.Model(course).Where("id = ?", course.ID).Count(&count)
	if count == 0 {
		if err := db.Create(course).Error; err != nil {
			log.Printf("  ⚠ Failed to seed course (id=%s): %v\n", course.ID, err)
		} else {
			log.Printf("  ✓ Seeded course: %s\n", course.TitleID)
		}
	} else {
		// Update existing record
		if err := db.Save(course).Error; err != nil {
			log.Printf("  ⚠ Failed to update course (id=%s): %v\n", course.ID, err)
		} else {
			log.Printf("  ✓ Updated course: %s\n", course.TitleID)
		}
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
		if err := db.Save(lesson).Error; err != nil {
			log.Printf("  ⚠ Failed to update lesson (id=%s): %v\n", lesson.ID, err)
		} else {
			log.Printf("  ✓ Updated lesson: %s\n", lesson.TitleID)
		}
	}
}

func upsertQuiz(db *gorm.DB, quiz *domain.LessonQuiz) {
	var count int64
	db.Model(quiz).Where("id = ?", quiz.ID).Count(&count)
	if count == 0 {
		if err := db.Create(quiz).Error; err != nil {
			log.Printf("  ⚠ Failed to seed quiz (id=%s): %v\n", quiz.ID, err)
		}
	} else {
		if err := db.Save(quiz).Error; err != nil {
			log.Printf("  ⚠ Failed to update quiz (id=%s): %v\n", quiz.ID, err)
		}
	}
}

func upsertChallenge(db *gorm.DB, challenge *domain.Challenge) {
	var count int64
	db.Model(challenge).Where("id = ?", challenge.ID).Count(&count)
	if count == 0 {
		if err := db.Create(challenge).Error; err != nil {
			log.Printf("  ⚠ Failed to seed challenge (slug=%s): %v\n", challenge.Slug, err)
		} else {
			log.Printf("  ✓ Seeded challenge: %s\n", challenge.Slug)
		}
	} else {
		if err := db.Save(challenge).Error; err != nil {
			log.Printf("  ⚠ Failed to update challenge (slug=%s): %v\n", challenge.Slug, err)
		}
	}
}
