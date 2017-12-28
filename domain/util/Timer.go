package util

import "time"

//CountdownTimer 倒计时
type CountdownTimer struct {
	TotalTime int //seconds
	LeftTime  int //seconds
	Tick      *time.Ticker
}

func NewCountdownTimer(totalTime int) *CountdownTimer {
	return &CountdownTimer{TotalTime: totalTime, LeftTime: totalTime, Tick: time.NewTicker(1 * time.Second)}
}
func (c *CountdownTimer) Countdown() {
	c.LeftTime--
}
func (c *CountdownTimer) Close() {
	c.Tick.Stop()
}
func (c *CountdownTimer) IsTimeOut() bool {
	return c.LeftTime <= 0
}
