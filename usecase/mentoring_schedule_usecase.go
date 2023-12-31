package usecase

import (
	"errors"

	"github.com/alwinihza/talent-connect-be/delivery/api/request"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/repository"
)

type MentoringScheduleUsecase interface {
	FindAll() ([]model.MentoringSchedule, error)
	FindById(id string) (*model.MentoringSchedule, error)
	SaveData(*request.MentoringScheduleRequest) error
	DeleteData(id string) error
	FindScheduleByMentorId(id string) ([]model.MentoringSchedule, error)
	FindScheduleByMenteeId(id string) ([]model.MentoringSchedule, error)
	SaveFeedbackMentoring(request *model.MentorMenteeSchedule) error
}

type mentoringScheduleUsecase struct {
	repo         repository.MentoringScheduleRepo
	mentorMentee MentorMenteeUsecase
}

func (m *mentoringScheduleUsecase) FindAll() ([]model.MentoringSchedule, error) {
	return m.repo.List()
}

func (m *mentoringScheduleUsecase) FindById(id string) (*model.MentoringSchedule, error) {
	return m.repo.Get(id)
}

func (m *mentoringScheduleUsecase) FindScheduleByMentorId(id string) ([]model.MentoringSchedule, error) {
	return m.repo.FindByMentorId(id)
}

func (m *mentoringScheduleUsecase) FindScheduleByMenteeId(id string) ([]model.MentoringSchedule, error) {
	return m.repo.FindByMenteeId(id)
}

func (m *mentoringScheduleUsecase) SaveFeedbackMentoring(request *model.MentorMenteeSchedule) error {
	return m.repo.SaveFeedback(request)
}

func (m *mentoringScheduleUsecase) SaveData(payload *request.MentoringScheduleRequest) error {
	var mentorMenteeList []model.MentorMentee
	for _, mentorMenteeId := range payload.MentorMentees {
		mentorMentee, err := m.mentorMentee.FindById(mentorMenteeId)
		if err != nil {
			return err
		}
		mentorMenteeList = append(mentorMenteeList, *mentorMentee)
	}

	listOfMentorSchedule, err := m.repo.FindByMentorId(mentorMenteeList[0].MentorID)
	if err != nil {
		return err
	}

	for _, mentorSchedule := range listOfMentorSchedule {
		if mentorSchedule.StartDate.Equal(payload.StartDate) && payload.ID != mentorSchedule.ID {
			return errors.New("You already have schedule at this time")
		}
	}

	mentoringSchedule := model.MentoringSchedule{
		BaseModel:     model.BaseModel{},
		MentorMentees: mentorMenteeList,
		Name:          payload.Name,
		Link:          payload.Link,
		Description:   payload.Description,
		StartDate:     payload.StartDate,
		EndDate:       payload.EndDate,
	}

	if payload.ID != "" {
		mentoringSchedule.ID = payload.ID

		err := m.repo.RemoveAllMentorMentees(payload.ID)
		if err != nil {
			return err
		}
	}

	return m.repo.Save(&mentoringSchedule)
}

func (m *mentoringScheduleUsecase) DeleteData(id string) error {
	return m.repo.Delete(id)
}

func NewMentoringScheduleUsecase(
	repo repository.MentoringScheduleRepo,
	mentorMentee MentorMenteeUsecase,
) MentoringScheduleUsecase {
	return &mentoringScheduleUsecase{
		repo:         repo,
		mentorMentee: mentorMentee,
	}
}
