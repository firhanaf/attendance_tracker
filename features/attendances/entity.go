package attendances

type AttendanceCore struct {
	ID         uint
	UserID     string
	Date       string
	Checkin    string
	Checkout   string
	IsCheckout bool
}

type AttendanceDataInterface interface {
	Create(input AttendanceCore) error
	GetAll() ([]AttendanceCore, error)
	Update(userid string, input AttendanceCore) error
	Export() ([]AttendanceCore, error)
}

type AttendanceServiceInterface interface {
	Add(input AttendanceCore) error
	ReadAll() ([]AttendanceCore, error)
	Edit(userid string, input AttendanceCore) error
	Download() ([]AttendanceCore, error)
}
