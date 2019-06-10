package utils

import (
	ui "github.com/instance-id/clui"
	"time"
)

type AnimateData struct {
	ui.BaseControl
	Button      *ui.ButtonNoShadow
	StartPOSX   int
	StartPOSY   int
	EndPOSX     int
	EndPOSY     int
	CurrentPOSX int
	CurrentPOSY int
	Sleep       int
	Duration    time.Duration
	Ticker      time.Ticker
}

func (a *AnimateData) AnimateBtnNShdw(durationMS time.Duration, tickSpeed time.Duration, distanceX int, distanceY int) {

	a.EndPOSX, a.EndPOSY = a.Button.Pos()
	a.StartPOSX = a.EndPOSX + distanceX
	a.StartPOSY = a.EndPOSY + distanceY
	a.CurrentPOSX = a.StartPOSX
	a.CurrentPOSY = a.StartPOSY
	a.Button.SetPos(a.StartPOSX, a.StartPOSY)
	//ui.PutEvent(ui.Event{Type: ui.EventRedraw})

	start := time.Duration(1 * time.Millisecond)
	end := start + durationMS*time.Millisecond

	for tick := start; tick >= end; tick++ {
		a.CurrentPOSX, a.CurrentPOSY = a.Button.Pos()
		//pos := a.Animate(tick)
		//a.CurrentPOSX--
		a.Button.SetPos(a.CurrentPOSX-1, a.CurrentPOSY)
		ui.PutEvent(ui.Event{Type: ui.EventRedraw})
		time.Sleep(durationMS / 500)
	}

}

//func (a *AnimateData) Animate(tick time.Duration) int {
//	ticker := time.NewTicker(tick)
//	defer ticker.Stop()
//	done := make(chan bool)
//	sleep := 100 * time.Millisecond
//	go func() {
//		time.Sleep(sleep)
//		done <- true
//	}()
//	ticks := 0
//	for {
//		select {
//		case <-done:
//			return a.CurrentPOSX - 1
//		case <-ticker.C:
//			ticks++
//		}
//	}
//}
