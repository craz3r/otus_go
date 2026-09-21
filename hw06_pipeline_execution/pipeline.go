package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	currentIn := in

	for _, stage := range stages {
		stageIn := make(Bi)

		go func(prevOut In, currentIn Bi) {
			defer close(currentIn)

			for val := range prevOut {
				select {
				case <-done:
					return
				default:
				}

				select {
				case <-done:
					return
				case currentIn <- val:
				}
			}
		}(currentIn, stageIn)

		currentIn = stage(stageIn)
	}

	return currentIn
}
