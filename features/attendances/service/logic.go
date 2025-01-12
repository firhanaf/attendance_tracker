package service

import "attendance_app/features/attendances"

type AttendanceService struct {
	attendanceData attendances.AttendanceDataInterface
}

func (a AttendanceService) Download() ([]attendances.AttendanceCore, error) {
	result, err := a.attendanceData.Export()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a AttendanceService) Add(input attendances.AttendanceCore) error {
	errInsert := a.attendanceData.Create(input)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

func (a AttendanceService) ReadAll() ([]attendances.AttendanceCore, error) {
	result, err := a.attendanceData.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a AttendanceService) Edit(userid string, input attendances.AttendanceCore) error {
	err := a.attendanceData.Update(userid, input)
	if err != nil {
		return err
	}
	return nil
}

func New(repo attendances.AttendanceDataInterface) attendances.AttendanceServiceInterface {
	return &AttendanceService{
		attendanceData: repo,
	}
}
