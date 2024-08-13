package usecase

import (
	"cleantaskmanager/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
}
func NewTaskUsecase(taskRepository domain.TaskRepository) domain.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
	}
}

func (tu *taskUsecase) GetTasks(c context.Context, claims *domain.Claims) ([]domain.Task, error) {
	return tu.taskRepository.GetTasks(c, claims)
}

func (tu *taskUsecase) GetTask(c context.Context, claims *domain.Claims, id primitive.ObjectID) (*domain.Task, error) {
	return tu.taskRepository.GetTask(c, claims, id)
}

func (tu *taskUsecase) AddTask(c context.Context, claims *domain.Claims ,task *domain.Task) error {
	return tu.taskRepository.AddTask(c,claims, task)
}

func (tu *taskUsecase) UpdateTask(c context.Context, claims *domain.Claims, id primitive.ObjectID, task *domain.UpdateTask) error {
	return tu.taskRepository.UpdateTask(c, claims, id, task)
}

func (tu *taskUsecase) DeleteTask(c context.Context, claims *domain.Claims, id primitive.ObjectID) error {
	return tu.taskRepository.DeleteTask(c, claims, id)
}

