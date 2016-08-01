package telerus

import (
	"log"

	"github.com/Sirupsen/logrus"
	"gopkg.in/telegram-bot-api.v4"
)

// AllLevels is
var AllLevels = []logrus.Level{
	logrus.DebugLevel,
	logrus.InfoLevel,
	logrus.WarnLevel,
	logrus.ErrorLevel,
	logrus.FatalLevel,
	logrus.PanicLevel,
}

// LevelThreshold Returns every logging level above and including the given parameter.
func LevelThreshold(l logrus.Level) []logrus.Level {
	for i := range AllLevels {
		if AllLevels[i] == l {
			return AllLevels[i:]
		}
	}
	return []logrus.Level{}
}

// TelerusHook is
type TelerusHook struct {
	AuthToken      string
	ChatID         int64
	Bot            *tgbotapi.BotAPI
	AcceptedLevels []logrus.Level
}

func (th *TelerusHook) initBot() error {
	bot, err := tgbotapi.NewBotAPI(th.AuthToken)
	if err != nil {
		log.Panic(err)
	}
	th.Bot = bot
	return nil
}

// Levels is
func (th *TelerusHook) Levels() []logrus.Level {
	if th.AcceptedLevels == nil {
		return AllLevels
	}
	return th.AcceptedLevels
}

// Fire is
func (th *TelerusHook) Fire(e *logrus.Entry) error {
	if th.Bot == nil {
		if err := th.initBot(); err != nil {
			return err
		}
	}

	msg := tgbotapi.NewMessage(th.ChatID, e.Message)
	msg.ParseMode = "markdown"

	_, err := th.Bot.Send(msg)

	return err

}
