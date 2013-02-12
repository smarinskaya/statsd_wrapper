package statsd_wrapper

//define StatsdStub type
type Stub struct {
	TimingFunc func(stat string, value int64, rate float32) error
	IncFunc    func(stat string, value int64, rate float32) error
	DecFunc    func(stat string, value int64, rate float32) error
	GaugeFunc  func(stat string, value int64, rate float32) error
	CloseFunc  func() error
}

func (sd *Stub) Timing(stat string, value int64, rate float32) error {
	if sd.TimingFunc != nil {
		return sd.TimingFunc(stat, value, rate)
	}
	return nil
}

func (sd *Stub) Inc(stat string, value int64, rate float32) error {
	if sd.IncFunc != nil {
		return sd.IncFunc(stat, value, rate)
	}
	return nil
}

func (sd *Stub) Dec(stat string, value int64, rate float32) error {
	if sd.DecFunc != nil {
		return sd.DecFunc(stat, value, rate)
	}
	return nil
}

func (sd *Stub) Gauge(stat string, value int64, rate float32) error {
	if sd.GaugeFunc != nil {
		return sd.GaugeFunc(stat, value, rate)
	}
	return nil
}

func (sd *Stub) Close() error {
	return nil
}
