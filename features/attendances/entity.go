package attendances

type AttendanceCore struct {
	ID         uint
	Date       string
	Checkin    string
	Checkout   string
	IsCheckout bool
}

type AttendanceDataInterface interface {
	Create(input AttendanceCore) error
	GetAll() ([]AttendanceCore, error)
	Update(isCheckout bool) error
}

type AttendanceServiceInterface interface {
	Add(input AttendanceCore) error
	ReadAll() ([]AttendanceCore, error)
	Edit(isCheckout bool) error
}
