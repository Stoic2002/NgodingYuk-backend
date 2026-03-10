-- Script to truncate all tables in the NgodingYuk database and reset their primary key sequences
-- WARNING: This will delete ALL data in all tables permanently.

TRUNCATE TABLE 
    users,
    challenges,
    courses,
    modules,
    lessons,
    lesson_quizzes,
    user_challenge_progress,
    user_lesson_progress,
    user_course_enrollments,
    user_xp_history,
    certificates
RESTART IDENTITY CASCADE;

-- Note on CASCADE: 
-- This ensures that if there are foreign key constraints linking these tables, 
-- PostgreSQL will automatically truncate the dependent tables as well.
--
-- Note on RESTART IDENTITY:
-- This resets any auto-incrementing ID sequences back to 1.
