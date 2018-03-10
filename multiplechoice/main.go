package main

import (
	"fmt"
	"github.com/tockins/interact"
	"github.com/AlecAivazis/survey"
)

func main() {
	surv()
}

func surv() (err error) {
	answers := struct {
		Profile string
	}{}

	var qs = []*survey.Question{
		{
			Name: "profile",
			Prompt: &survey.Select{
				Message: "Choose a color:",
				Options: []string{"red", "blue", "green"},
				Default: "red",
			},
		},
	}

	// perform the questions
	err = survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

func inter() {
	interact.Run(&interact.Interact{
		Questions: []*interact.Question{
			{
				Quest: interact.Quest{
					Msg: "how much for a teacup?",
					Choices: interact.Choices{
						Alternatives: []interact.Choice{
							{
								Text:     "Gyokuro teapcup",
								Response: "20",
							},
							{
								Text:     "Sencha teacup",
								Response: -10,
							},
							{
								Text:     "Matcha teacup",
								Response: 15.50,
							},
						},
					},
				},
				Action: func(c interact.Context) interface{} {
					val, _ := c.Ans().Int()
					fmt.Println(val)
					return nil
				},
			},
		},
	})
}
