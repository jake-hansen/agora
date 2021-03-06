package zoomdomain

const (
	TypeInstant              = 1
	TypeScheduled            = 2
	TypeRecurringNoFixedTime = 3
	TypeRecurringFixedTime   = 8

	TypeRecurrenceDaily   = 1
	TypeRecurrenceWeekly  = 2
	TypeRecurrenceMonthly = 3
)

type Meeting struct {
	ID          int         `json:"id"`
	Topic       string      `json:"topic"`
	Type        int         `json:"type"`
	StartTime   string      `json:"start_time"`
	Duration    int         `json:"duration"`
	ScheduleFor string      `json:"schedule_for,omitempty"`
	Timezone    string      `json:"timezone,omitempty"`
	Password    string      `json:"password,omitempty"`
	Agenda      string      `json:"agenda"`
	Recurrence  *Recurrence `json:"recurrence,omitempty"`
	Settings    *Settings   `json:"settings,omitempty"`
	JoinURL     string      `json:"join_url"`
	StartURL    string      `json:"start_url"`
}

type Recurrence struct {
	Type           int    `json:"type"`
	RepeatInterval int    `json:"repeat_interval"`
	WeeklyDays     string `json:"weekly_days"`
	MonthlyDay     int    `json:"monthly_day"`
	MonthlyWeekDay int    `json:"monthly_week_day"`
	EndTimes       int    `json:"end_times"`
	EndDateTime    string `json:"end_date_time"`
}

type Settings struct {
	HostVideo        bool   `json:"host_video"`
	ParticipantVideo bool   `json:"participant_video"`
	CNMeeting        bool   `json:"cn_meeting"`
	INMeeting        bool   `json:"in_meeting"`
	JoinBeforeHost   bool   `json:"join_before_host"`
	JBHTime          int    `json:"jbh_time"`
	MuteUponEntry    bool   `json:"mute_upon_entry"`
	Watermark        bool   `json:"watermark"`
	UsePMI           bool   `json:"use_pmi"`
	ApprovalType     int    `json:"approval_type"`
	RegistrationType int    `json:"registration_type"`
	Audio            string `json:"audio"`
	AutoRecording    string `json:"auto_recording"`
}

type MeetingList struct {
	PageCount    int        `json:"page_count"`
	PageNumber   int        `json:"page_number"`
	PageSize     int        `json:"page_size"`
	TotalRecords int        `json:"total_records"`
	Meetings     []*Meeting `json:"meetings"`
}
