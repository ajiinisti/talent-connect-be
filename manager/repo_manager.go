package manager

import "github.com/alwinihza/talent-connect-be/repository"

type RepoManager interface {
	RoleRepo() repository.RoleRepo
	UserRepo() repository.UserRepo
	MentoringScheduleRepo() repository.MentoringScheduleRepo
	MentorMenteeRepo() repository.MentorMenteeRepo
	ProgramRepo() repository.ProgramRepo
	ActivityRepo() repository.ActivityRepo
	ParticipantRepo() repository.ParticipantRepo
	QuestionRepo() repository.QuestionRepo
	QuestionCategoryRepo() repository.QuestionCategoryRepo
	QuestionAnswerRepo() repository.QuestionAnswerRepo
	EvaluationCategoryRepo() repository.EvaluationCategoryRepo
	EvaluationRepo() repository.EvaluationRepo
	AnswerRepo() repository.AnswerRepo
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) RoleRepo() repository.RoleRepo {
	return repository.NewRoleRepo(r.infra.Conn())
}

func (r *repoManager) UserRepo() repository.UserRepo {
	return repository.NewUserRepo(r.infra.Conn())
}

func (r *repoManager) MentoringScheduleRepo() repository.MentoringScheduleRepo {
	return repository.NewMentoringScheduleRepo(r.infra.Conn())
}

func (r *repoManager) MentorMenteeRepo() repository.MentorMenteeRepo {
	return repository.NewMentorMenteeRepo(r.infra.Conn())
}

func (r *repoManager) ProgramRepo() repository.ProgramRepo {
	return repository.NewProgramRepo(r.infra.Conn())
}

func (r *repoManager) ActivityRepo() repository.ActivityRepo {
	return repository.NewActivityRepo(r.infra.Conn())
}

func (r *repoManager) ParticipantRepo() repository.ParticipantRepo {
	return repository.NewParticipantRepo(r.infra.Conn())
}

func (r *repoManager) QuestionRepo() repository.QuestionRepo {
	return repository.NewQuestionRepo(r.infra.Conn())
}

func (r *repoManager) QuestionCategoryRepo() repository.QuestionCategoryRepo {
	return repository.NewQuestionCategoryRepo(r.infra.Conn())
}

func (r *repoManager) QuestionAnswerRepo() repository.QuestionAnswerRepo {
	return repository.NewQuestionAnswerRepo(r.infra.Conn())
}

func (r *repoManager) EvaluationCategoryRepo() repository.EvaluationCategoryRepo {
	return repository.NewEvaluationCategoryRepo(r.infra.Conn())
}

func (r *repoManager) EvaluationRepo() repository.EvaluationRepo {
	return repository.NewEvaluationRepo(r.infra.Conn())
}

func (r *repoManager) AnswerRepo() repository.AnswerRepo {
	return repository.NewAnswerRepo(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infra: infra,
	}
}
