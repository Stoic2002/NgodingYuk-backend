package service

import (
	"errors"

	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"github.com/arulkarim/ngodingyuk-server/pkg/hash"
	jwtpkg "github.com/arulkarim/ngodingyuk-server/pkg/jwt"
	"github.com/google/uuid"
)

type AuthService struct {
	userRepo UserRepository
}

func NewAuthService(userRepo UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// === Request/Response types (kept local to service for simplicity) ===

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type AuthResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	User         UserResponse `json:"user"`
}

type UserResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	XP          int64  `json:"xp"`
	Level       int64  `json:"level"`
	StreakCount int64  `json:"streak_count"`
	Locale      string `json:"locale"`
}

// === Service Methods ===

func (s *AuthService) Register(req RegisterRequest) (*AuthResponse, error) {
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return nil, errors.New("username, email, and password are required")
	}
	if len(req.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	// Check if email already exists
	existing, _ := s.userRepo.FindByEmail(req.Email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	// Check if username already exists
	existingUser, _ := s.userRepo.FindByUsername(req.Username)
	if existingUser != nil {
		return nil, errors.New("username already taken")
	}

	hashedPassword, err := hash.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &domain.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		XP:           0,
		Level:        1,
		StreakCount:  0,
		Locale:       "id",
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("failed to create user")
	}

	accessToken, err := jwtpkg.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshToken, err := jwtpkg.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         toUserResponse(user),
	}, nil
}

func (s *AuthService) Login(req LoginRequest) (*AuthResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}

	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !hash.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("invalid email or password")
	}

	accessToken, err := jwtpkg.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshToken, err := jwtpkg.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         toUserResponse(user),
	}, nil
}

func (s *AuthService) RefreshToken(req RefreshRequest) (*AuthResponse, error) {
	claims, err := jwtpkg.ValidateToken(req.RefreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	uid, err := uuid.Parse(claims.UserID)
	if err != nil {
		return nil, errors.New("invalid user ID in token")
	}

	user, err := s.userRepo.FindByID(uid)
	if err != nil {
		return nil, errors.New("user not found")
	}

	accessToken, err := jwtpkg.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshToken, err := jwtpkg.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         toUserResponse(user),
	}, nil
}

func (s *AuthService) GetProfile(userID uuid.UUID) (*UserResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	resp := toUserResponse(user)
	return &resp, nil
}

func (s *AuthService) GetUserByID(userID uuid.UUID) (*domain.User, error) {
	return s.userRepo.FindByID(userID)
}

func toUserResponse(user *domain.User) UserResponse {
	return UserResponse{
		ID:          user.ID.String(),
		Username:    user.Username,
		Email:       user.Email,
		XP:          user.XP,
		Level:       user.Level,
		StreakCount: user.StreakCount,
		Locale:      user.Locale,
	}
}
