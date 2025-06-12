package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"userapi/models"
	"userapi/repository"
)

type mockUserRepository struct {
	users map[int64]*models.User
}

func newMockUserRepository() *mockUserRepository {
	return &mockUserRepository{
		users: make(map[int64]*models.User),
	}
}

func (m *mockUserRepository) Create(ctx context.Context, user *models.User) error {
	user.ID = int64(len(m.users) + 1)
	m.users[user.ID] = user
	return nil
}

func (m *mockUserRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	if user, exists := m.users[id]; exists {
		return user, nil
	}
	return nil, nil
}

func (m *mockUserRepository) Update(ctx context.Context, user *models.User) error {
	if _, exists := m.users[user.ID]; !exists {
		return nil
	}
	m.users[user.ID] = user
	return nil
}

func (m *mockUserRepository) Delete(ctx context.Context, id int64) error {
	delete(m.users, id)
	return nil
}

func (m *mockUserRepository) List(ctx context.Context) ([]*models.User, error) {
	users := make([]*models.User, 0, len(m.users))
	for _, user := range m.users {
		users = append(users, user)
	}
	return users, nil
}

func TestUserHandler_Create(t *testing.T) {
	repo := newMockUserRepository()
	handler := NewUserHandler(repo)

	tests := []struct {
		name       string
		payload    models.User
		wantStatus int
	}{
		{
			name: "valid user",
			payload: models.User{
				Name:        "John Doe",
				Age:         30,
				PhoneNumber: "+1234567890",
				Email:       "john@example.com",
			},
			wantStatus: http.StatusCreated,
		},
		{
			name: "invalid email",
			payload: models.User{
				Name:        "John Doe",
				Age:         30,
				PhoneNumber: "+1234567890",
				Email:       "invalid-email",
			},
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.payload)
			req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			handler.Create(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("Create() status = %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
}

func TestUserHandler_GetByID(t *testing.T) {
	repo := newMockUserRepository()
	handler := NewUserHandler(repo)

	// Create a test user
	user := &models.User{
		Name:        "John Doe",
		Age:         30,
		PhoneNumber: "+1234567890",
		Email:       "john@example.com",
	}
	repo.Create(context.Background(), user)

	tests := []struct {
		name       string
		id         string
		wantStatus int
	}{
		{
			name:       "existing user",
			id:         "1",
			wantStatus: http.StatusOK,
		},
		{
			name:       "non-existing user",
			id:         "999",
			wantStatus: http.StatusNotFound,
		},
		{
			name:       "invalid id",
			id:         "invalid",
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/users/"+tt.id, nil)
			w := httptest.NewRecorder()

			handler.GetByID(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("GetByID() status = %v, want %v", w.Code, tt.wantStatus)
			}
		})
	}
} 