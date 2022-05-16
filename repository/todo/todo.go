package todo

import (
	"errors"
	"todolist/entities"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type TodoRepo struct {
	Db *gorm.DB
}

func NewRepoTodo(db *gorm.DB) *TodoRepo {
	return &TodoRepo{
		Db: db,
	}
}

func (t *TodoRepo) InsertTodo(newTodo entities.Todo) (entities.Todo, error) {
	if err := t.Db.Create(&newTodo).Error; err != nil {
		log.Warn(err)
		return entities.Todo{}, err
	}

	log.Info()
	return newTodo, nil
}

func (t *TodoRepo) GetAllCompleteTodo(UserID uint) ([]entities.Todo, error) {
	var GetAll []entities.Todo

	if err := t.Db.Where("user_id = ? AND project_id = ? AND completed = 'yes'", UserID, 0).Find(&GetAll).Error; err != nil {
		log.Warn(err)
		return GetAll, err
	}
	return GetAll, nil
}

func (t *TodoRepo) GetAllUnCompleteTodo(UserID uint) ([]entities.Todo, error) {
	var GetAll []entities.Todo

	if err := t.Db.Where("user_id = ? AND project_id = ? AND completed='no'", UserID, 0).Find(&GetAll).Error; err != nil {
		log.Warn(err)
		return GetAll, err
	}
	return GetAll, nil
}

func (t *TodoRepo) GetTodoID(id uint, UserID uint) (entities.Todo, error) {
	Todo := entities.Todo{}

	if err := t.Db.Where("id = ? AND user_id =?", id, UserID).First(&Todo).Error; err != nil {
		log.Warn(err)
		return entities.Todo{}, errors.New("Data Not Found")
	}

	log.Info()
	return Todo, nil
}

func (t *TodoRepo) UpdateTodo(id uint, UserID uint, update entities.Todo) (entities.Todo, error) {
	var todo entities.Todo
	if err := t.Db.Where("id =? AND user_id =?", id, UserID).First(&todo).Updates(&update).Find(&todo).Error; err != nil {
		log.Warn(err)
		return todo, errors.New("Data Not Found")
	}
	log.Info()
	return todo, nil
}

func (t *TodoRepo) DeleteTodo(id uint, UserID uint) error {
	var delete entities.Todo

	if err := t.Db.Where("id = ? AND user_id = ?", id, UserID).First(&delete).Delete(&delete).Error; err != nil {
		return err
	}
	log.Info()
	return nil

}

func (t *TodoRepo) Completed(id uint, UserID uint) (entities.Todo, error) {
	var completed entities.Todo
	if err := t.Db.Where("id=? AND user_id=?", id, UserID).First(&completed).Update("completed", "yes").Error; err != nil {
		log.Warn(err)
		return completed, errors.New("Data Not Found")
	}
	return completed, nil
}

func (t *TodoRepo) Reopen(id uint, UserID uint) (entities.Todo, error) {
	var uncompleted entities.Todo
	if err := t.Db.Where("id=? AND user_id=?", id, UserID).First(&uncompleted).Update("completed", "no").Error; err != nil {
		log.Warn(err)
		return uncompleted, errors.New("Data Not Found")
	}
	return uncompleted, nil
}
