package pomodoro
import("time")

type State int32

const (
	FINISHED   State = 0
	RUNNING State = 1
	PAUSED State = 1
)

type Step int32

const (
	WORKING Step = 1
	RELAXING Step = 2
)

type Pomodoro struct {
	StartedAt *time.Time
	PausedAt  *time.Time
	State     State
	Step      Step
	UserId    int
}

type PomoBot interface {
	AddPomodoro(userId int) error
	StopPomodoro(userId int) error
	PausePomodoro(userId int) error
	ResumePomodoro(userId int) error
	UpdateState() (map[int]Step, error)
}

func New(userId int) *Pomodoro{
	now := time.Now()
	return &Pomodoro {
		StartedAt: &now,
		PausedAt:  nil,
		State:     RUNNING,
		Step:	   WORKING,
		UserId:    userId,
	}
}

func (pomodoro *Pomodoro) NeedsUpdate() bool {
	switch pomodoro.Step {
	case WORKING:
		delta := time.Minute * 25
		return pomodoro.StartedAt.Add(delta).Before(time.Now())
	case RELAXING:
		delta := time.Minute * 5
		return pomodoro.StartedAt.Add(delta).Before(time.Now())
	}
	return false
}

func (pomodoro *Pomodoro) NextStep() Step {
	now := time.Now()
	nextStep := (pomodoro.Step + 1) % 2
	pomodoro.Step = nextStep
	pomodoro.StartedAt = &now
	return pomodoro.Step
}

func (pomodoro *Pomodoro) Pause() {
	now := time.Now()
	pomodoro.PausedAt = &now
	pomodoro.State = PAUSED
}

func (pomodoro *Pomodoro) Resume() {
	delta := pomodoro.StartedAt.Sub(*pomodoro.PausedAt)
	newStart := time.Now().Add(delta)
	pomodoro.StartedAt = &newStart
	pomodoro.State = RUNNING
}
