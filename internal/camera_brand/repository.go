package camera_brand

import (
	"gorm.io/gorm"
)

type CameraBrandRepository interface {
	Create(*CameraBrand) error
	GetAll() ([]CameraBrand, error)
	GetByID(uint) (*CameraBrand, error)
	Update(*CameraBrand) error
	Delete(uint) error
}

type cameraBrandRepo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) CameraBrandRepository {
	return &cameraBrandRepo{db: db}
}

func (r *cameraBrandRepo) Create(brand *CameraBrand) error {
	return r.db.Create(brand).Error
}

func (r *cameraBrandRepo) GetAll() ([]CameraBrand, error) {
	var brands []CameraBrand
	err := r.db.Find(&brands).Error
	return brands, err
}

func (r *cameraBrandRepo) GetByID(id uint) (*CameraBrand, error) {
	var brand CameraBrand
	if err := r.db.First(&brand, id).Error; err != nil {
		return nil, err
	}
	return &brand, nil
}

func (r *cameraBrandRepo) Update(brand *CameraBrand) error {
	return r.db.Save(brand).Error
}

func (r *cameraBrandRepo) Delete(id uint) error {
	return r.db.Delete(&CameraBrand{}, id).Error
}
