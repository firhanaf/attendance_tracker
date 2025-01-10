package service

import "attendance_app/features/attendances"

type AttendanceService struct {
	attendanceData attendances.AttendanceDataInterface
}

func (a AttendanceService) Add(input attendances.AttendanceCore) error {
	errInsert := a.attendanceData.Create(input)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

func (a AttendanceService) ReadAll() ([]attendances.AttendanceCore, error) {
	//TODO implement me
	panic("implement me")
}

func (a AttendanceService) Edit(isCheckout bool) error {
	//TODO implement me
	panic("implement me")
}

func New(repo attendances.AttendanceDataInterface) attendances.AttendanceServiceInterface {
	return &AttendanceService{
		attendanceData: repo,
	}
}
