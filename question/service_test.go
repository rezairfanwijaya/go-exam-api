package question_test

import (
	"errors"
	"fmt"
	"testing"

	q "github.com/rezairfanwijaya/go-exam-api.git/question"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mock repository
type MockIQuestionRepository struct {
	mock.Mock
}

func (m *MockIQuestionRepository) Save(question q.Question) (q.Question, error) {
	args := m.Called(question)
	return args.Get(0).(q.Question), args.Error(1)
}

func (m *MockIQuestionRepository) FindByID(id int) (q.Question, error) {
	args := m.Called(id)
	return args.Get(0).(q.Question), args.Error(1)
}

func TestNewService(t *testing.T) {
	mock := new(MockIQuestionRepository)
	service := q.NewService(mock)
	assert.NotNil(t, service)
}

func TestSave(t *testing.T) {
	testCases := []struct {
		Name        string
		Question    q.Question
		Input       q.QuestionCreateInput
		Expectation q.Question
		WantError   bool
	}{
		{
			Name: "success",
			Question: q.Question{
				Question: "sebutkan 3 nama hewan berkaki 4",
			},
			Input: q.QuestionCreateInput{
				Question: "sebutkan 3 nama hewan berkaki 4",
			},
			Expectation: q.Question{
				ID:       1,
				Question: "sebutkan 3 nama hewan berkaki 4",
			},
		}, {
			Name: "failed",
			Question: q.Question{
				Question: "",
			},
			Input: q.QuestionCreateInput{
				Question: "",
			},
			Expectation: q.Question{
				ID:       0,
				Question: "",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			mock := new(MockIQuestionRepository)
			questionService := q.QuestionService{RepoQuestion: mock}

			if testCase.WantError {
				mock.On("Save", testCase.Question).Return(testCase.Expectation, nil)
				actual, _ := questionService.Save(testCase.Input)
				assert.Nil(t, actual)
			} else {
				mock.On("Save", testCase.Question).Return(testCase.Expectation, nil)
				_, err := questionService.Save(testCase.Input)
				assert.Nil(t, err)
			}

		})
	}
}

func TestGetByID(t *testing.T) {
	testCases := []struct {
		Name        string
		ID          int
		Expectation q.Question
		WantError   bool
	}{
		{
			Name: "success",
			ID:   1,
			Expectation: q.Question{
				ID:       1,
				Question: "berapakah kaki kucing",
			},
			WantError: false,
		}, {
			Name:        "failed",
			ID:          90,
			Expectation: q.Question{},
			WantError:   true,
		}, {
			Name:        "id is smaller than 1",
			ID:          0,
			Expectation: q.Question{},
			WantError:   true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			mock := new(MockIQuestionRepository)
			questionService := q.QuestionService{RepoQuestion: mock}

			if testCase.WantError {
				errMsg := fmt.Sprintf("question dengan id %v tidak ditemukan", testCase.ID)
				mock.On("FindByID", testCase.ID).Return(testCase.Expectation, errors.New(errMsg))
				actual, err := questionService.GetByID(testCase.ID)
				assert.Equal(t, testCase.Expectation, actual)
				assert.NotNil(t, err)
			} else {

				mock.On("FindByID", testCase.ID).Return(testCase.Expectation, nil)
				actual, err := questionService.GetByID(testCase.ID)
				assert.Nil(t, err)
				assert.NotNil(t, actual)
			}

		})
	}
}
