package http

//TODO:PTERO review this

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"workingtimeweb/server/core"
	workingDay "workingtimeweb/server/workingday"

	"github.com/julienschmidt/httprouter"
)

// Service : defines repository service
type Service struct {
	repo   core.Repository
	Router http.Handler
}

// New : creates new instance of service
func New(repo core.Repository) Service {
	service := Service{
		repo: repo,
	}

	router := httprouter.New()
	router.GET("/workingDays", service.Index)
	router.POST("/workingDays", service.Create)
	router.DELETE("/workingDays/:id", service.Delete)
	router.PUT("/workingDays/:id", service.Update)

	service.Router = UseMiddlewares(router)

	return service
}

// Index : loads all workingDays
func (s Service) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	service := workingDay.NewService(s.repo, r.Context().Value("userId").(string))
	workingDays, err := service.GetWorkingDays()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(workingDays)
}

func (s Service) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	service := workingDay.NewService(s.repo, r.Context().Value("userId").(string))
	payload, _ := ioutil.ReadAll(r.Body)

	githubRepo := workingDay.GitHubRepo{}
	json.Unmarshal(payload, &githubRepo)

	workingDay, err := service.CreateWorkingDayFor(githubRepo)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(workingDay)
}

func (s Service) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	service := workingDay.NewService(s.repo, r.Context().Value("userId").(string))

	repoID, _ := strconv.Atoi(params.ByName("id"))
	githubRepo := workingDay.GitHubRepo{RepoID: int64(repoID)}

	_, err := service.RemoveWorkingDay(githubRepo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s Service) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	service := workingDay.NewService(s.repo, r.Context().Value("userId").(string))
	payload, _ := ioutil.ReadAll(r.Body)

	githubRepo := workingDay.GitHubRepo{}
	json.Unmarshal(payload, &githubRepo)

	workingDay, err := service.UpdateWorkingDayWith(githubRepo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(workingDay)
}
