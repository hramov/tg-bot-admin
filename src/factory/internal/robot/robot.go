package robot

import (
	"github.com/go-vgo/robotgo"
	"log"
)

type Robot interface {
	Create(name, title, description string) (string, error)
}

func New() Robot {
	return &robot{}
}

type robot struct {
}

func (r *robot) Create(name, title, description string) (string, error) {

	telegramPID, err := robotgo.FindIds("Telegram")
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	if len(telegramPID) > 0 {
		err = robotgo.ActiveName("Telegram")
		if err != nil {
			log.Println(err.Error())
		}
		robotgo.Sleep(1)

		log.Println(robotgo.GetTitle())

		robotgo.TypeStr("/newbot")
		robotgo.KeyTap("enter")

		robotgo.Sleep(2)

		robotgo.TypeStr(title)
		robotgo.KeyTap("enter")

		robotgo.Sleep(2)

		robotgo.TypeStr(name + "_bot")
		robotgo.KeyTap("enter")

	}

	return "", nil
}
