package controller

import "github.com/era/pomobot/pomodoro"

var pomobot pomodoro.PomoBot

func Init(bot pomodoro.PomoBot) {
	pomobot = bot
}

func Instructions() string {
	return "/begin: starts a new pomodoro\n" +
		   "/pause: pauses the pomodoro\n" +
		   "/resume: resumes the pomodoro\n" +
		   "/end: stops the pomodoro\n"
}

func CreateNew(userId int) string {
	err := pomobot.AddPomodoro(userId)
	if err != nil {
		return err.Error()
	} else {
		return "Pomodoro added"
	}
}

func End(userId int) string {
	err := pomobot.StopPomodoro(userId)
	if err != nil {
		return err.Error()
	} else {
		return "Pomodoro stoped"
	}
}

func Pause(userId int) string {
	err := pomobot.PausePomodoro(userId)
	if err != nil {
		return err.Error()
	} else {
		return "Pomodoro paused"
	}
}

func Resume(userId int) string {
	err := pomobot.ResumePomodoro(userId)
	if err != nil {
		return err.Error()
	} else {
		return "Pomodoro resumed"
	}
}