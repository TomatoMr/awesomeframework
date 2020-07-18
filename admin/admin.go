package admin

import "sync"

type PprofController struct {
	OnOff   chan struct{}
	Timeout int
	Host    string
	Used    chan struct{}
}

var pCh PprofController
var once = sync.Once{}

func (p PprofController) Push(s struct{}) {
	p.OnOff <- s
}

func (p PprofController) Pop() struct{} {
	return <-p.OnOff
}

func Get() PprofController {
	once.Do(func() {
		pCh = PprofController{
			OnOff:   make(chan struct{}),
			Timeout: 600,
			Host:    "127.0.0.1:8888",
			Used:    make(chan struct{}),
		}
	})
	return pCh
}

func (p PprofController) Close() {
	close(p.OnOff)
	close(p.Used)
}
