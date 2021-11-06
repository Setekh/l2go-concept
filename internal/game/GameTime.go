package game

import (
	"time"
)

const (
	TICKS_PER_SECOND        = 10
	MILLIS_IN_TICK          = 1000 / TICKS_PER_SECOND
	IN_GAME_DAYS_PER_DAY    = 6
	MILLIS_PER_IN_GAME_DAY  = (3600000 * 24) / IN_GAME_DAYS_PER_DAY
	SECONDS_PER_IN_GAME_DAY = MILLIS_PER_IN_GAME_DAY / 1000
	TICKS_PER_IN_GAME_DAY   = SECONDS_PER_IN_GAME_DAY * TICKS_PER_SECOND
)

type TimeController interface {
	GetGameTicks() int
	GetGameTime() int
	GetGameHour() int
	Start()
}

type timeController struct {
	referenceTime int64
	isRunning     bool
}

func CreateTimeController() TimeController {
	controller := &timeController{
		referenceTime: time.Now().Unix(),
	}
	controller.Start()
	return controller
}

func (tc *timeController) GetGameTicks() int {
	return int((currentTimeMillis() - tc.referenceTime) / MILLIS_IN_TICK)
}

func (tc *timeController) GetGameTime() int {
	return (tc.GetGameTicks() % TICKS_PER_IN_GAME_DAY) / MILLIS_IN_TICK
}

func (tc *timeController) GetGameHour() int {
	return tc.GetGameTime() / 60
}

func (tc *timeController) Start() {
	tc.isRunning = true

	go func() {
		defer func() {
			tc.isRunning = false
		}()

		for tc.isRunning {
			nextTickTime := ((currentTimeMillis() / MILLIS_IN_TICK) * MILLIS_IN_TICK) + 100

			// TODO process here?

			sleepTime := nextTickTime - currentTimeMillis()
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		}
	}()
}

func currentTimeMillis() int64 {
	return time.Now().UnixMilli()
}
