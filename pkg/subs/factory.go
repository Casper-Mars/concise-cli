package subs

type WorkerFactory interface {
	CreateWorker(file string, placeholder map[string]string) Worker
}

type defaultWorkerFactory struct{}

func (d defaultWorkerFactory) CreateWorker(file string, placeholder map[string]string) Worker {
	return Worker{
		file:        file,
		placeholder: placeholder,
	}
}

func NewDefaultWorkerFactory() WorkerFactory {
	return defaultWorkerFactory{}
}
