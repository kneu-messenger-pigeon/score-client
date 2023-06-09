package score

import (
	"github.com/h2non/gock"
	scoreApi "github.com/kneu-messenger-pigeon/score-api"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var host = "https://localhost:8080"
var studentId = uint32(999)
var disciplineId = 123
var lessonId = 245

func TestClient_GetStudentDisciplines(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedResult := scoreApi.DisciplineScoreResults{
			{
				Discipline: scoreApi.Discipline{
					Id:   100,
					Name: "Капітал!",
				},
				ScoreRating: scoreApi.ScoreRating{
					Total:         17,
					StudentsCount: 25,
					Rating:        8,
					MinTotal:      10,
					MaxTotal:      20,
				},
			},
			{
				Discipline: scoreApi.Discipline{
					Id:   110,
					Name: "Гроші та лихварство",
				},
				ScoreRating: scoreApi.ScoreRating{
					Total:         12,
					StudentsCount: 25,
					Rating:        12,
					MinTotal:      7,
					MaxTotal:      17,
				},
			},
		}

		gock.New(host).
			Get("/v1/students/999/disciplines").
			Reply(200).
			JSON(expectedResult)

		client := Client{
			Host: host,
		}

		actualResult, err := client.GetStudentDisciplines(studentId)
		assert.Equal(t, expectedResult, actualResult)
		assert.NoError(t, err)
		assert.False(t, gock.HasUnmatchedRequest())
	})

	t.Run("error api", func(t *testing.T) {
		gock.New(host).
			Get("/v1/students/999/disciplines").
			Reply(500).
			JSON(`{
				"error": "Test error description"
			}`)

		client := Client{
			Host: host,
		}

		actualResult, err := client.GetStudentDisciplines(studentId)

		assert.Error(t, err)
		assert.Equal(t, "API error: Test error description", err.Error())
		assert.Empty(t, actualResult)
	})

	t.Run("error http", func(t *testing.T) {
		gock.New(host).
			Get("/v1/students/999/disciplines").
			Reply(500)

		client := Client{
			Host: host,
		}

		actualResult, err := client.GetStudentDisciplines(studentId)

		assert.Error(t, err)
		assert.Equal(t, "Receive http code: 500", err.Error())
		assert.Empty(t, actualResult)
	})
}

func TestClient_GetStudentDiscipline(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedResult := scoreApi.DisciplineScoreResult{
			Discipline: scoreApi.Discipline{
				Id:   199,
				Name: "Капітал!",
			},
			ScoreRating: scoreApi.ScoreRating{
				Total:         17,
				StudentsCount: 25,
				Rating:        8,
				MinTotal:      10,
				MaxTotal:      20,
			},
			Scores: []scoreApi.Score{
				{
					Lesson: scoreApi.Lesson{
						Id:   lessonId,
						Date: time.Date(2023, time.Month(2), 12, 0, 0, 0, 0, time.Local),
						Type: scoreApi.LessonType{
							Id:        5,
							ShortName: "МК",
							LongName:  "Модульний контроль.",
						},
					},
					FirstScore:  floatPointer(4.5),
					SecondScore: nil,
					IsAbsent:    true,
				},
			},
		}

		gock.New(host).
			Get("/v1/students/999/disciplines/123").
			Reply(200).
			JSON(expectedResult)

		client := Client{
			Host: host,
		}

		actualResult, err := client.GetStudentDiscipline(studentId, disciplineId)
		assert.Equal(t, expectedResult, actualResult)
		assert.NoError(t, err)
		assert.False(t, gock.HasUnmatchedRequest())
	})

	t.Run("error api", func(t *testing.T) {
		gock.New(host).
			Get("/v1/students/999/disciplines/123").
			Reply(500).
			JSON(`{
				"error": "Test error description"
			}`)

		client := Client{
			Host: host,
		}

		actualResult, err := client.GetStudentDiscipline(studentId, disciplineId)

		assert.Error(t, err)
		assert.Equal(t, "API error: Test error description", err.Error())
		assert.Empty(t, actualResult)
	})

	t.Run("error http", func(t *testing.T) {
		gock.New(host).
			Get("/v1/students/999/disciplines/123").
			Reply(500)

		client := Client{
			Host: host,
		}

		actualResult, err := client.GetStudentDiscipline(studentId, disciplineId)

		assert.Error(t, err)
		assert.Equal(t, "Receive http code: 500", err.Error())
		assert.Empty(t, actualResult)
	})
}

func TestClient_GetStudentScore(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedResult := scoreApi.DisciplineScore{
			Discipline: scoreApi.Discipline{
				Id:   disciplineId,
				Name: "Капітал!",
			},
			Score: scoreApi.Score{
				Lesson: scoreApi.Lesson{
					Id:   lessonId,
					Date: time.Date(2023, time.Month(2), 12, 0, 0, 0, 0, time.Local),
					Type: scoreApi.LessonType{
						Id:        5,
						ShortName: "МК",
						LongName:  "Модульний контроль.",
					},
				},
				FirstScore:  floatPointer(4.5),
				SecondScore: nil,
				IsAbsent:    true,
			},
		}

		gock.New(host).
			Get("/v1/students/999/disciplines/123/scores/245").
			Reply(200).
			JSON(expectedResult)

		client := Client{
			Host: host,
		}

		actualResult, err := client.GetStudentScore(studentId, disciplineId, lessonId)
		assert.Equal(t, expectedResult, actualResult)
		assert.NoError(t, err)
		assert.False(t, gock.HasUnmatchedRequest())
	})

	t.Run("error api", func(t *testing.T) {
		gock.New(host).
			Get("/v1/students/999/disciplines/123/scores/245").
			Reply(500).
			JSON(`{
				"error": "Test error description"
			}`)

		client := Client{
			Host: host,
		}

		actualResult, err := client.GetStudentScore(studentId, disciplineId, lessonId)

		assert.Error(t, err)
		assert.Equal(t, "API error: Test error description", err.Error())
		assert.Empty(t, actualResult)
	})

	t.Run("error http", func(t *testing.T) {
		gock.New(host).
			Get("/v1/students/999/disciplines/123/scores/245").
			Reply(500)

		client := Client{
			Host: host,
		}

		actualResult, err := client.GetStudentScore(studentId, disciplineId, lessonId)

		assert.Error(t, err)
		assert.Equal(t, "Receive http code: 500", err.Error())
		assert.Empty(t, actualResult)
	})
}

func floatPointer(value float32) *float32 {
	return &value
}
