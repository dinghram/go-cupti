package cupti

import (
	tr "github.com/rai-project/tracer"
	context "golang.org/x/net/context"
)

type Options struct {
	ctx            context.Context
	tracer         tr.Tracer
	samplingPeriod int
	activities     []string
	domains        []string
	callbacks      []string
}

type Option func(o *Options)

func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.ctx = ctx
	}
}

func Tracer(ctx context.Context) Option {
	return func(o *Options) {
		o.tracer = tracer
	}
}

func Activities(activities []string) Option {
	return func(o *Options) {
		o.activities = activities
	}
}

func Domains(domains []string) Option {
	return func(o *Options) {
		o.domains = domains
	}
}

func Callbacks(callbacks []string) Option {
	return func(o *Options) {
		o.callbacks = callbacks
	}
}

func SamplingPeriod(s int) Option {
	return func(o *Options) {
		o.samplingPeriod = s
	}
}

func NewOptions(opts ...Option) *Options {
	Config.Wait()

	options := &Options{
		ctx:            context.Background(),
		tracer:         tracer,
		samplingPeriod: Config.SamplingPeriod,
		activities:     Config.Activities,
		domains:        Config.Domains,
		callbacks:      Config.Callbacks,
	}

	for _, o := range opts {
		o(options)
	}

	return options
}
