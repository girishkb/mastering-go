package pipeline


type Executor func(interface{}) (interface{},error)

type  Pipeline interface {
	Pipe(executor Executor) Pipeline
	Merge() <- chan interface{}
}

type pipeline struct {
	dataC chan interface{}
	errC chan error
	executors []Executor
}

func New(f func(chan interface{})) Pipeline {
	inc := make(chan interface{})
	go f(inc)
	return &pipeline{
		dataC:     inc,
		errC:      make(chan error),
		executors: []Executor{},
	}
}

func (p *pipeline) Pipe(executor Executor) Pipeline {
	p.executors = append(p.executors,executor)
	return p
}

func (p *pipeline) Merge() <-chan interface{} {
	for i := 0; i < len(p.executors); i++ {
		p.dataC,p.errC = run(p.dataC,p.executors[i])
	}
	return p.dataC
}

func run(inC <-chan interface{}, f Executor)  (chan interface{}, chan error){
	outC := make(chan interface{})
	errC := make(chan error)
	go func() {
		defer close(outC)
		for v := range inC {
			res, err := f(v)
			if err !=nil {
				errC <- err
				continue
			}
			outC <-res

		}
	}()
	return outC,errC
}