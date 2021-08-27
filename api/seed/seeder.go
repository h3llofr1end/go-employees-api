package seed

import (
	"log"

	"github.com/h3llofr1end/go-employees-api/api/models"
	"github.com/jinzhu/gorm"
)

var departments = []models.Department{
	models.Department{
		Title: "Маркетинг",
	},
	models.Department{
		Title: "Бухгалтерия",
	},
	models.Department{
		Title: "IT",
	},
}

var employees = []models.Employee{
	models.Employee{
		Name:   "Вася",
		Sex:    "m",
		Age:    27,
		Salary: 100000,
	},
	models.Employee{
		Name:   "Лена",
		Sex:    "f",
		Age:    23,
		Salary: 60000,
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Department{}, &models.Employee{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Department{}, &models.Employee{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}
	err = db.Debug().Model(&models.Employee{}).AddForeignKey("department_id", "departments(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error :%v", err)
	}

	for i, _ := range departments {
		err = db.Debug().Model(&models.Department{}).Create(&departments[i]).Error
		if err != nil {
			log.Fatalf("cannot seed departments table: %v", err)
		}
	}

	for i, _ := range employees {
		err = db.Debug().Model(&models.Employee{}).Create(&employees[i]).Error
		if err != nil {
			log.Fatalf("cannot seed employees table: %v", err)
		}
	}
}
