package statsd_wrapper

const (
	//Default Ip Address and port for Statsd
	DefaultStatsdAddr = "127.0.0.1:8125"
	//Default Directory name for metrics in Graphite
	DefaultStatsdDir = "default"
)

type Config struct {
	Addr string
	Dir  string
}

//we are defining StatsdClient interface. This way we can use stubbed type in testing.
type StatsdClient interface {
	Inc(stat string, value int64, rate float32) error
	Dec(stat string, value int64, rate float32) error
	Gauge(stat string, value int64, rate float32) error
	Timing(stat string, value int64, rate float32) error
	Close() error
}
