package pomodoro
import(
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	pomodoro := New(123)
	if pomodoro.UserId != 123 {
		t.Errorf("UserId is not set properly (expected %d, got %d)", 123, pomodoro.UserId)
	}
	if pomodoro.State != RUNNING {
		t.Errorf("Pomodoro was created in the wrong state %d", pomodoro.State)
	}
}

func TestNeedsUpdateAfterWorking(t *testing.T) {
	startedAt := time.Now().Add(time.Minute * -26)
	pomodoro := &Pomodoro {
		StartedAt: &startedAt,
		PausedAt:  nil,
		State:     RUNNING,
		Step:	   WORKING,
		UserId:    123,
	}
	if !pomodoro.NeedsUpdate() {
		t.Error("Pomodoro should need update after 25 minutes of working step")
	}
}

func TestNeedsUpdateAfterRlx(t *testing.T) {
	startedAt := time.Now().Add(time.Minute * -5)
	pomodoro := &Pomodoro {
		StartedAt: &startedAt,
		PausedAt:  nil,
		State:     RUNNING,
		Step:	   RELAXING,
		UserId:    123,
	}
	if !pomodoro.NeedsUpdate() {
		t.Error("Pomodoro should need update after 5 minutes of relaxing step")
	}
}

func TestNeedsUpdateReturnsFalse(t *testing.T) {
	startedAt := time.Now().Add(time.Minute * -3)
	pomodoro := &Pomodoro {
		StartedAt: &startedAt,
		PausedAt:  nil,
		State:     RUNNING,
		Step:	   RELAXING,
		UserId:    123,
	}
	if pomodoro.NeedsUpdate() {
		t.Error("Pomodoro should not need update before 5 minutes of relaxing step")
	}
}

func TestNextStep(t *testing.T) {
	pomodoro := New(123)
	step := pomodoro.NextStep()
	if step != RELAXING {
		t.Error("Did not update correctly the step from working to relaxing")
	}
	step = pomodoro.NextStep()
	if step != WORKING {
		t.Error("Did not update correctly the step from relaxing to working")
	}
}

func TestPause(t *testing.T) {
	pomodoro := New(123)
	pomodoro.Pause()
	if pomodoro.State != PAUSED {
		t.Error("Pomodoro should be in paused state")
	}
}

func TestResume(t *testing.T) {
	pomodoro := New(123)
	pomodoro.Pause()
	pomodoro.Resume()
	if pomodoro.State != RUNNING {
		t.Error("Pomodoro should be running")
	}
}

