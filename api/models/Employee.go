package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Employee struct {
	ID           uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name         string    `gorm:"size:255; not null; unique" json:"name"`
	Sex          string    `gorm:"size:10; not null;" json:"sex"`
	Age          int       `gorm:"not null" json:"age"`
	Salary       int       `gorm:"not null" json:"salary"`
	DepartmentId int       `json:"department_id"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (e *Employee) Prepare() {
	e.ID = 0
	e.Name = html.EscapeString(strings.TrimSpace(e.Name))
	e.Sex = html.EscapeString(strings.TrimSpace(e.Sex))
	e.CreatedAt = time.Now()
	e.UpdatedAt = time.Now()
}

func (e *Employee) Validate() error {
	if e.Name == "" {
		return errors.New("Required Name")
	}
	if e.Sex == "" {
		return errors.New("Required Sex")
	}
	if e.Age < 1 {
		return errors.New("Required Age")
	}
	if e.Salary < 1 {
		return errors.New("Required Salary")
	}
	return nil
}

func (e *Employee) SaveEmployee(db *gorm.DB) (*Employee, error) {
	var err error
	err = db.Debug().Model(&Employee{}).Create(&e).Error
	if err != nil {
		return &Employee{}, err
	}
	return e, nil
}

func (e *Employee) FindAllEmployees(db *gorm.DB) (*[]Employee, error) {
	var err error
	employees := []Employee{}
	err = db.Debug().Model(&Employee{}).Limit(100).Find(&employees).Error
	if err != nil {
		return &[]Employee{}, err
	}
	return &employees, nil
}

func (e *Employee) FindEmployeeByID(db *gorm.DB, eid uint64) (*Employee, error) {
	var err error
	err = db.Debug().Model(&Employee{}).Where("id = ?", eid).Take(&e).Error
	if err != nil {
		return &Employee{}, err
	}
	return e, nil
}

func (e *Employee) UpdateEmployee(db *gorm.DB) (*Employee, error) {
	var err error
	err = db.Debug().Model(&Employee{}).Where("id = ?", e.ID).Updates(Employee{Name: e.Name, Sex: e.Sex, Age: e.Age, Salary: e.Salary, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Employee{}, err
	}
	return e, nil
}

func (e *Employee) DeleteEmployee(db *gorm.DB, eid uint64) (int64, error) {
	db = db.Debug().Model(&Employee{}).Where("id = ?", eid).Take(&Employee{}).Delete(&Employee{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Employee not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
