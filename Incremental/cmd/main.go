package main

import (
	"clase-02/internal/application"
	"fmt"
)

/*
type Application struct {
	ConfigApplication
}

type ConfigApplication struct {
	Host   string
	Port   int
	DbFile string
}

func NewApplication(cfg ConfigApplication) *Application {
	return &Application{
		ConfigApplication: cfg,
	}
}

func (app *Application) Run() error {
	fmt.Println("Port: ", app.Port)
	fmt.Println("DbFile: ", app.DbFile)
	fmt.Println("Host: ", app.Host)

	return nil
}
*/

func main() {
	/*
		port, err := strconv.Atoi(os.Getenv("ENV_PORT"))
		if err != nil {
			fmt.Println(err)
			return
		}
		cfg := ConfigApplication{
			Host:   os.Getenv("ENV_HOST"),
			DbFile: os.Getenv("ENV_PATH_DB_FILE"),
			Port:   port,
		}

		app := NewApplication(cfg)

		app.Run()
	*/
	app := application.NewDefaultHTTP(":8080")

	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}

}
