package project

import (
	"errors"
	"todolist/entities"

	"github.com/labstack/gommon/log"

	"gorm.io/gorm"
)

type ProjectRepo struct {
	Db *gorm.DB
}

func NewRepoProject(db *gorm.DB) *ProjectRepo {
	return &ProjectRepo{
		Db: db,
	}
}

func (p *ProjectRepo) InsertProject(newProject entities.Project) (entities.Project, error) {
	if err := p.Db.Create(&newProject).Error; err != nil {
		log.Warn(err)
		return entities.Project{}, err
	}
	log.Info()
	return newProject, nil
}

func (p *ProjectRepo) GetAllProject(UserID uint) ([]entities.Project, error) {
	var GetAll []entities.Project

	if err := p.Db.Where("user_id = ?", UserID).Find(&GetAll).Error; err != nil {
		log.Warn(err)
		return GetAll, err
	}
	return GetAll, nil
}

func (p *ProjectRepo) GetProjectID(id uint, UserID uint) (entities.Project, error) {
	Project := entities.Project{}

	if err := p.Db.Where("id = ? AND user_id =?", id, UserID).First(&Project).Error; err != nil {
		log.Warn(err)
		return entities.Project{}, errors.New("Data Not Found")
	}

	log.Info()
	return Project, nil
}

func (p *ProjectRepo) UpdateProject(id uint, UserID uint, update entities.Project) (entities.Project, error) {
	var project entities.Project
	if err := p.Db.Where("id =? AND user_id =?", id, UserID).First(&project).Updates(&update).Find(&project).Error; err != nil {
		log.Warn(err)
		return project, errors.New("Data Not Found")
	}
	log.Info()
	return project, nil
}

func (p *ProjectRepo) DeleteProject(id uint, UserID uint) error {
	var delete entities.Project
	var deletetodo entities.Todo
	if err := p.Db.Where("id = ? AND user_id = ?", id, UserID).First(&delete).Delete(&delete).Error; err != nil {
		return err
	}
	if err := p.Db.Where("project_id=? AND user_id = ?", id, UserID).Delete(&deletetodo).Error; err != nil {
		log.Warn(err)
		return err
	}
	log.Info()
	return nil
}

func (p *ProjectRepo) AddTodoToProject(id uint, TodoID uint, UserID uint) (entities.Todo, error) {
	var TodoProject entities.Todo
	var project entities.Project
	if err := p.Db.Where("id =? AND user_id =?", id, UserID).First(&project).Error; err != nil {
		log.Warn(err)
		return TodoProject, err
	}
	if err := p.Db.Where("id = ? AND user_id = ?", TodoID, UserID).First(&TodoProject).Update("project_id", id).Error; err != nil {
		log.Warn(err)
		return TodoProject, err
	}

	log.Info()
	return TodoProject, nil
}

func (p *ProjectRepo) TodoProjectCompleted(id uint, TodoID uint, UserID uint) (entities.Todo, error) {
	var completed entities.Todo
	if err := p.Db.Where("id=? AND user_id=? AND project_id=?", TodoID, UserID, id).First(&completed).Update("completed", "yes").Error; err != nil {
		log.Warn(err)
		return completed, errors.New("Data Not Found")
	}
	return completed, nil
}

func (p *ProjectRepo) TodoProjectReopen(id uint, TodoID, UserID uint) (entities.Todo, error) {
	var uncompleted entities.Todo
	if err := p.Db.Where("id=? AND user_id=? AND project_id=?", TodoID, UserID, id).First(&uncompleted).Update("completed", "no").Error; err != nil {
		log.Warn(err)
		return uncompleted, errors.New("Data Not Found")
	}
	return uncompleted, nil
}

func (p *ProjectRepo) GetAllProjectCompleteTodo(id uint, UserID uint) ([]entities.Todo, error) {
	var GetAll []entities.Todo

	if err := p.Db.Where("user_id = ? AND project_id = ? AND completed = 'yes'", UserID, id).Find(&GetAll).Error; err != nil {
		log.Warn(err)
		return GetAll, err
	}
	return GetAll, nil
}

func (p *ProjectRepo) GetAllProjectUnCompleteTodo(id uint, UserID uint) ([]entities.Todo, error) {
	var GetAll []entities.Todo

	if err := p.Db.Where("user_id = ? AND project_id = ? AND completed='no'", UserID, id).Find(&GetAll).Error; err != nil {
		log.Warn(err)
		return GetAll, err
	}
	return GetAll, nil
}

func (p *ProjectRepo) MoveToHome(id uint, TodoID uint, UserID uint) error {
	var move entities.Todo
	if err := p.Db.Where("project_id=? AND todo_id AND user_id = ?", id, TodoID, UserID).First(&move).Update("project_id", 0).Error; err != nil {
		log.Warn(err)
		return err
	}
	return nil
}
