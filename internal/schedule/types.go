package schedule

type ST string

const (
	Minute  ST = "MINUTE"
	Hourly  ST = "HOURLY"
	Daily   ST = "DAILY"
	Weekly  ST = "WEEKLY"
	Monthly ST = "MONTHLY"
	Once    ST = "ONCE"
	OnStart ST = "ONSTART"
	OnLogon ST = "ONLOGON"
	OnIdle  ST = "ONIDLE"
)

var ScheduleTypes = []ST{Minute, Hourly, Daily, Weekly, Monthly, Once, OnStart, OnLogon, OnIdle}
