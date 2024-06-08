package app

import "fmt"

type App struct {
	Alg Algorithms
}

type Algorithms interface {
	Start() (string, error)
}

func New() (*App, error) {
	a := &App{}
	return a, nil
}

func (a *App) Run(algs Algorithms) error {
	result, err := algs.Start()
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", result)
	return nil
}
