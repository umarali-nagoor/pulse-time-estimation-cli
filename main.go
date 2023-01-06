package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
	"github.com/IBM-Cloud/pulse-time-estimation-cli/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "pulse"
	app.Usage = "A command line tool to estimate time of Terraform resources"
	app.Action = func(c *cli.Context) error {
		fmt.Println(" Execute 'pulse-cli -h' to get valid commands supported by this tool.")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "predict",
			Usage: "Predicts estimated time for given plan file",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "plan", Required: true, Usage: "input plan file path in json format"},
			},
			Action: func(c *cli.Context) error {
				planFilePath := c.String("plan")
				if planFilePath == "" || !strings.HasSuffix(planFilePath, ".json") {
					return fmt.Errorf("[ERROR] Please provide valid path to plan file")
				}
				err := commands.Predict(planFilePath)
				if err != nil {
					return fmt.Errorf("%+v", err)
				}

				return nil
			},
		},
		{
			Name:  "predict-get",
			Usage: "Get predicted time for a given job id",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "id", Required: true, Usage: "job id"},
			},
			Action: func(c *cli.Context) error {
				jobID := c.String("id")
				if jobID == "" {
					return fmt.Errorf("[ERROR] Please provide valid id")
				}
				err := commands.GetPrediction(jobID)
				if err != nil {
					return fmt.Errorf("%+v", err)
				}

				return nil
			},
		},
		{
			Name:  "predict-delete",
			Usage: "Delete predict job id",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "id", Required: true, Usage: "job id"},
			},
			Action: func(c *cli.Context) error {
				jobID := c.String("id")
				if jobID == "" {
					return fmt.Errorf("[ERROR] Please provide valid id")
				}
				err := commands.DeletePrediction(jobID)
				if err != nil {
					return fmt.Errorf("%+v", err)
				}

				return nil
			},
		},
		{
			Name:  "predict-get-status",
			Usage: "Get predict job status",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "id", Required: true, Usage: "job id"},
			},
			Action: func(c *cli.Context) error {
				jobID := c.String("id")
				if jobID == "" {
					return fmt.Errorf("[ERROR] Please provide valid id")
				}
				err := commands.GetPredictionStatus(jobID)
				if err != nil {
					return fmt.Errorf("%+v", err)
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
