package components

type dbQuery interface {
	// do the Db query for this event
	Query()
	// get any errors that happened
	Error() error
	// wait for this query to complete
	Wait()
	// called when query is done
	Done()
}

type dbEvent struct {
	X *XormDB
	// completetion channel
	chnl chan bool
}

func (ev *dbEvent) Wait() {
	<-ev.chnl
	close(ev.chnl)
}

func (ev *dbEvent) Done() {
	ev.chnl <- true
}
