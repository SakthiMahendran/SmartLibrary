package controller

import (
	"github.com/SakthiMahendran/SmartLibrary/dbdriver"
	"github.com/SakthiMahendran/SmartLibrary/dbdriver/models"
	"golang.org/x/crypto/bcrypt"
)

type AdminController struct {
	dbDriver dbdriver.DBDriver
}

func NewAdminController(dbDriver dbdriver.DBDriver) *AdminController {
	return &AdminController{
		dbDriver: dbDriver,
	}
}

func (ac *AdminController) CreateAdmin(name, id, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin := &models.Admin{
		Name:     name,
		ID:       id,
		Password: string(hashedPassword),
	}

	if err := ac.dbDriver.InsertAdmin(admin); err != nil {
		return err
	}

	return nil
}

func (ac *AdminController) AuthenticateAdmin(id, password string) (bool, error) {
	admin, err := ac.dbDriver.FindAdminByID(id)
	if err != nil {
		return false, err
	}

	if admin == nil {
		return false, nil // Admin not found
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return false, nil
	}

	return true, nil
}
