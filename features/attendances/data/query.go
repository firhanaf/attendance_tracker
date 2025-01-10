package data

import (
	"attendance_app/features/attendances"
	"errors"
	"gorm.io/gorm"
)

type AttendanceQuery struct {
	db *gorm.DB
}

func (repo *AttendanceQuery) GetAll() ([]attendances.AttendanceCore, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *AttendanceQuery) Update(isCheckout bool) error {
	//TODO implement me
	panic("implement me")
}

func (repo *AttendanceQuery) Create(input attendances.AttendanceCore) error {
	attendanceModel := AttendanceCoreToModel(input)
	tx := repo.db.Create(&attendanceModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("No rows were affected")
	}
	return nil
}

func New(db *gorm.DB) *AttendanceQuery {
	return &AttendanceQuery{
		db: db,
	}
}
