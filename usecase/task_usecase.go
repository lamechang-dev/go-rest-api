package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type ITaskUseCase interface {
	GetAllTasks(userId int) ([]model.TaskResponse, error)
	GetTaskById(userId int, taskId int) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId int, taskId int) (model.TaskResponse, error)
	DeleteTask(userId int, taskId int) error
}

type taskUseCase struct {
	tr repository.ITaskRepository
}

func NewTaskUseCase(tr repository.ITaskRepository) ITaskUseCase {
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