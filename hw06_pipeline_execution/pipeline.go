package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if len(stages) == 0 {
		return in
	}

	currentIn := in

	for _, stage := range stages {
		stageOut := stage(currentIn)
		nextIn := make(Bi)

		go func(src Out, dst Bi) {
			defer close(dst)
			for {
				select {
				case <-done:
					return
				case val, ok := <-src:
					if !ok {
						return
					}
					dst <- val
				}
			}
		}(stageOut, nextIn)

		if done != nil {
			go func(src Out) {
				<-done
				for range src {
					// Вычитываем данные для предотвращения дедлока
				}
			}(stageOut)
		}

		currentIn = nextIn
	}

	return currentIn
}
