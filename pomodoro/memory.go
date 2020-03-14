package pomodoro

import (
	"errors"
	"sync"
)
//TODO tests
type MemoryPomoBot struct {
	pomodori map[int]*Pomodoro
	//TODO in the future we should only lock for a single userId
	mutex *sync.Mutex
}
// storage all pomodori in memory not usable for production
func MemoryBot() *MemoryPomoBot{
	return &MemoryPomoBot {
		pomodori: make(map[int]*Pomodoro),
		mutex: &sync.Mutex{},
	}
}

func (bot *MemoryPomoBot) UpdateState() (map[int]Step, error) {
	bot.mutex.Lock()
	defer bot.mutex.Unlock()
	updates := make(map[int]Step)
	for _, p := range bot.pomodori {
		if p.NeedsUpdate() {
			step := p.NextStep()
			updates[p.UserId] = step
		}
	}
	return updates, nil
}

func (bot *MemoryPomoBot) AddPomodoro(userId int) error {
	bot.mutex.Lock()
	defer bot.mutex.Unlock()
	if bot.pomodori[userId] != nil {
		return errors.New("Cannot have multiple pomodori running or paused")
	}
	pomodoro := New(userId)
	bot.pomodori[userId] = pomodoro

	return nil
}

func (bot *MemoryPomoBot) PausePomodoro(userId int) error {
	bot.mutex.Lock()
	defer bot.mutex.Unlock()

	pomodoro := bot.pomodori[userId]
	if pomodoro == nil || pomodoro.State != RUNNING {
		return errors.New("No Running pomodoro")
	}

	pomodoro.Pause()

	return nil
}
func (bot *MemoryPomoBot) ResumePomodoro(userId int) error {
	bot.mutex.Lock()
	defer bot.mutex.Unlock()

	pomodoro := bot.pomodori[userId]
	if pomodoro == nil || pomodoro.State != PAUSED {
		return errors.New("No Paused pomodoro")
	}
	pomodoro.Resume()

	return nil
}

func (bot *MemoryPomoBot) StopPomodoro(userId int) error {
	bot.mutex.Lock()
	defer bot.mutex.Unlock()
	pomodoro := bot.pomodori[userId]
	if pomodoro == nil {
		return errors.New("There's no running or paused pomodoro")
	}
	bot.pomodori[userId] = nil
	return nil
}
