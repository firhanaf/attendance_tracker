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
	var result []Attendances
	tx := repo.db.Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var resultCore []attendances.AttendanceCore
	for _, v := range result {
		resultCore = append(resultCore, attendances.AttendanceCore{
			ID:         v.ID,
			UserID:     v.UserID,
			Date:       v.Date,
			Checkin:    v.Checkin,
			Checkout:   v.Checkout,
			IsCheckout: v.IsCheckout,
		})
	}
	return resultCore, nil
}

func (repo *AttendanceQuery) Update(userid string, input attendances.AttendanceCore) error {
	attendanceModel := AttendanceCoreToModel(input)
	tx := repo.db.Where("user_id = ?", userid).Updates(&attendanceModel)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (repo *AttendanceQuery) Create(input attendances.AttendanceCore) error {
	attendanceModel := AttendanceCoreToModel(input)
	tx := repo.db.Create(&attendanceModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no rows were affected")
	}
	return nil
}

func (repo *AttendanceQuery) Export() ([]attendances.AttendanceCore, error) {
	var result []Attendances
	tx := repo.db.Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var resultCore []attendances.AttendanceCore
	for _, v := range result {
		resultCore = append(resultCore, attendances.AttendanceCore{
			ID:         v.ID,
			UserID:     v.UserID,
			Date:       v.Date,
			Checkin:    v.Checkin,
			Checkout:   v.Checkout,
			IsCheckout: v.IsCheckout,
		})
	}
	return resultCore, nil
}

func New(db *gorm.DB) *AttendanceQuery {
	return &AttendanceQuery{
		db: db,
	}
}
