package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userId int) error
	GetTaskById(task *model.Task, userId int, taskId int) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userId int, taskId int) error
	DeleteTask(task *model.Task, taskId int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId int) error {
	if err := tr.db.Where("user_id=?", userId).Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetTaskById(task *model.Task, userId int, taskId int) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(task, taskId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) UpdateTask(task *model.Task, userId int, taskId int) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("user_id=? AND id=?", userId, taskId).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}

	if (result.RowsAffected == 0) {
		return fmt.Errorf("record not found")
	}

	return nil

	
}

func (tr *taskRepository) DeleteTask(task *model.Task, taskId int) error {
	if err := tr.db.Where("id=?", taskId).Delete(task).Error; err != nil {
		return err
	}
	return nil
}