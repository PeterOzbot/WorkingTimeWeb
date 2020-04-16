package workingday

//TODO:PTERO review this

import (
	"workingtimeweb/server/core"
)

type GitHubRepo struct {
	RepoID      int64  `json:"id"`
	RepoURL     string `json:"html_url"`
	RepoName    string `json:"full_name"`
	Language    string `json:"language"`
	Description string `json:"description"`
	Notes       string `json:"notes"`
}

type Service struct {
	userId string
	repo   core.Repository
}

func (s Service) GetWorkingDays() ([]*core.WorkingDay, error) {
	return s.repo.FindAll(map[string]interface{}{"userId": s.userId})
}

func (s Service) CreateWorkingDayFor(githubRepo GitHubRepo) (*core.WorkingDay, error) {
	workingDay := s.githubRepoToWorkingDay(githubRepo)
	err := s.repo.Create(workingDay)
	if err != nil {
		return nil, err
	}
	return workingDay, nil
}

func (s Service) UpdateWorkingDayWith(githubRepo GitHubRepo) (*core.WorkingDay, error) {
	workingDay := s.githubRepoToWorkingDay(githubRepo)
	err := s.repo.Create(workingDay)
	if err != nil {
		return nil, err
	}
	return workingDay, nil
}

func (s Service) RemoveWorkingDay(githubRepo GitHubRepo) (*core.WorkingDay, error) {
	workingDay := s.githubRepoToWorkingDay(githubRepo)
	err := s.repo.Delete(workingDay)
	if err != nil {
		return nil, err
	}
	return workingDay, nil
}

func (s Service) githubRepoToWorkingDay(githubRepo GitHubRepo) *core.WorkingDay {
	return &core.WorkingDay{
		/*	UserID:      s.userId,
			RepoID:      strconv.Itoa(int(githubRepo.RepoID)),
			RepoName:    githubRepo.RepoName,
			RepoURL:     githubRepo.RepoURL,
			Language:    githubRepo.Language,
			Description: githubRepo.Description,
			Notes:       githubRepo.Notes,*/
		//TODO:PTERO review this
	}
}

func NewService(repo core.Repository, userId string) Service {
	return Service{
		repo:   repo,
		userId: userId,
	}
}
