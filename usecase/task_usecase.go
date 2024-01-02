package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type ITaskUsecase interface {
	GetAllTasks(userId int) ([]model.TaskResponse, error)
	GetTaskById(userId int, taskId int) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId int, taskId int) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUseCase struct {
	tr repository.ITaskRepository
}

func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &taskUseCase{tr}
}

func (tu *taskUseCase) GetAllTasks(userId int) ([]model.TaskResponse, error) {
	tasks := []model.Task{}
	err := tu.tr.GetAllTasks(&tasks, userId)
	if err != nil {
		return nil, err
	}
	resTasks := []model.TaskResponse{}
	for _, task := range tasks {
		t := model.TaskResponse{
			ID: task.ID,
			Title: task.Title,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}

func (tu *taskUseCase) GetTaskById(userId int, taskId int) (model.TaskResponse, error) {
	task := model.Task{}
	err := tu.tr.GetTaskById(&task, userId, taskId)
	if err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID: task.ID,
		Title: task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUseCase) CreateTask(task model.Task) (model.TaskResponse, error) {
	err := tu.tr.CreateTask(&task)
	if err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID: task.ID,
		Title: task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUseCase) UpdateTask(task model.Task, userId int, taskId int) (model.TaskResponse, error) {
	err := tu.tr.UpdateTask(&task, userId, taskId)
	if err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID: task.ID,
		Title: task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUseCase) DeleteTask(userId uint, taskId uint) error {
	err := tu.tr.DeleteTask(userId, taskId)
	if err != nil {
		return err
	}
	return nil
}