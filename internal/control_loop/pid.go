package control_loop

import (
	"github.com/markusressel/fan2go/internal/util"
	"math"
)

// PidControlLoop is a PidLoop based control loop implementation.
type PidControlLoop struct {
	pidLoop *util.PidLoop
}

// NewPidControlLoop creates a PidControlLoop, which is a very simple control that Pidly applies the given
// target pwm. It can also be used to gracefully approach the target by
// utilizing the "maxPwmChangePerCycle" property.
func NewPidControlLoop(
	p float64,
	i float64,
	d float64,
) *PidControlLoop {
	// TODO: somehow incorporate default values
	//pidLoop = util.NewPidLoop(
	//	0.03,
	//	0.002,
	//	0.0005,
	//)

	return &PidControlLoop{
		pidLoop: util.NewPidLoop(p, i, d),
	}
}

func (l *PidControlLoop) Cycle(target int, lastSetPwm int) int {
	result := l.pidLoop.Loop(float64(target), float64(lastSetPwm))

	// ensure we are within sane bounds
	coerced := util.Coerce(float64(lastSetPwm)+result, 0, 255)
	stepTarget := int(math.Round(coerced))

	return stepTarget
}
