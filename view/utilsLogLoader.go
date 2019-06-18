package view

import (
	ui "github.com/instance-id/clui"
)

var LogChanView chan bool
var LogRunningView = false

func LoadLogs() error {
	LogChanView = make(chan bool)
	LogRunning = true
	defer Tails.Done()

	for {
		select {
		case line, ok := <-Tails.Lines:
			if !ok {
				_ = Tails.Wait()
				break
			}
			if []string{line.Text} != nil {
				LogViewer.Log.AddText([]string{line.Text})
				ui.PutEvent(ui.Event{Type: ui.EventRedraw})
			}
		case <-LogChanView:
			LogRunning = false
			return Tails.Stop()
		}
	}
}
