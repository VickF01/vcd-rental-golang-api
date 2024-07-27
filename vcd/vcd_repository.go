package vcd

import "gorm.io/gorm"

type Repository interface {
	Create(vcd VCD) (VCD, error)
	GetAllVCD() ([]VCD, error)
	GetOneVCD(ID int) (VCD, error)
	UpdateVCD(vcd VCD) (VCD, error)
	DeleteVCD(vcd VCD) (VCD, error)
}

type repository struct{ db *gorm.DB }

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(vcd VCD) (VCD, error) {
	err := r.db.Create(&vcd).Error
	return vcd, err
}

func (r *repository) GetAllVCD() ([]VCD, error) {
	var vcds []VCD
	err := r.db.Find(&vcds).Error
	return vcds, err
}

func (r *repository) GetOneVCD(ID int) (VCD, error) {
	var vcd VCD
	err := r.db.Find(&vcd, ID).Error
	return vcd, err
}

func (r *repository) UpdateVCD(vcd VCD) (VCD, error) {
	err := r.db.Save(&vcd).Error
	return vcd, err
}

func (r *repository) DeleteVCD(vcd VCD) (VCD, error) {
	err := r.db.Delete(&vcd).Error
	return vcd, err
}
