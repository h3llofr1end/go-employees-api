package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Department struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255; not null; unique" json:"title"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (d *Department) Prepare() {
	d.ID = 0
	d.Title = html.EscapeString(strings.TrimSpace(d.Title))
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
}

func (d *Department) Validate() error {
	if d.Title == "" {
		return errors.New("Required Title")
	}
	return nil
}

func (d *Department) SaveDepartment(db *gorm.DB) (*Department, error) {
	var err error
	err = db.Debug().Model(&Department{}).Create(&d).Error
	if err != nil {
		return &Department{}, err
	}
	return d, nil
}

func (d *Department) FindAllDepartments(db *gorm.DB) (*[]Department, error) {
	var err error
	departments := []Department{}
	err = db.Debug().Model(&Department{}).Limit(100).Find(&departments).Error
	if err != nil {
		return &[]Department{}, err
	}
	return &departments, nil
}

func (d *Department) FindDepartmentByID(db *gorm.DB, did uint64) (*Department, error) {
	var err error
	err = db.Debug().Model(&Department{}).Where("id = ?", did).Take(&d).Error
	if err != nil {
		return &Department{}, err
	}
	return d, nil
}

func (d *Department) UpdateDepartment(db *gorm.DB) (*Department, error) {
	var err error
	err = db.Debug().Model(&Department{}).Where("id = ?", d.ID).Updates(Department{Title: d.Title, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Department{}, err
	}
	return d, nil
}

func (d *Department) DeleteDepartment(db *gorm.DB, did uint64) (int64, error) {
	db = db.Debug().Model(&Department{}).Where("id = ?", did).Take(&Department{}).Delete(&Department{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Department not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
