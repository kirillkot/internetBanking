package web

import (
	"internetBanking/api/models"

	"github.com/jinzhu/gorm"
)

// ISimpleModel ...
type ISimpleModel interface {
	Name() string

	New() interface{}
	NewArray(len, cap int) interface{}
}

// IModel ...
type IModel interface {
	ISimpleModel

	Create(db *gorm.DB, user *models.User, object interface{}) (result interface{}, err error)
	Update(db *gorm.DB, user *models.User, id uint,
		updates map[string]interface{}) (object interface{}, err error)
	Get(db *gorm.DB, user *models.User, id uint) (object interface{}, err error)
	GetObjects(db *gorm.DB, user *models.User) (objects interface{}, err error)
	Delete(db *gorm.DB, user *models.User, id uint) (object interface{}, err error)
}

// BaseModel ...
type BaseModel struct {
	ISimpleModel
}

// NewBaseModel ...
func NewBaseModel(model ISimpleModel) BaseModel {
	return BaseModel{
		ISimpleModel: model,
	}
}

// Create ...
func (m BaseModel) Create(db *gorm.DB, user *models.User, model interface{}) (interface{}, error) {
	if err := db.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

// Update ...
func (m BaseModel) Update(db *gorm.DB, user *models.User, id uint,
	updates map[string]interface{}) (interface{}, error) {
	object := m.ISimpleModel.New()
	if err := db.Where("id = ?", id).Find(object).Error; err != nil {
		return nil, err
	}
	if err := db.Model(object).Updates(updates).Error; err != nil {
		return nil, err
	}
	return object, nil
}

// Get ...
func (m BaseModel) Get(db *gorm.DB, user *models.User, id uint) (interface{}, error) {
	object := m.ISimpleModel.New()
	if err := db.Where("id = ?", id).Find(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}

// GetObjects ...
func (m BaseModel) GetObjects(db *gorm.DB, user *models.User) (interface{}, error) {
	objects := m.ISimpleModel.NewArray(0, 32)
	if err := db.Find(objects).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return objects, nil
}

// Delete ...
func (m BaseModel) Delete(db *gorm.DB, user *models.User, id uint) (interface{}, error) {
	object := m.ISimpleModel.New()
	if err := db.Where("id = ?", id).Find(object).Error; err != nil {
		return nil, err
	}
	if err := db.Delete(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}
