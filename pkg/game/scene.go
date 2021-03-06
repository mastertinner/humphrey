package game

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

var promptuiTemplates = &promptui.SelectTemplates{
	Label:    " ",
	Active:   `{{ "✓" | green }} {{ . }}`,
	Inactive: "  {{ . }}",
}

// Scene is a game scene.
type Scene struct {
	Image   string
	Body    string
	Actions map[string]Renderer
}

// Render prints the scene to the terminal.
func (s *Scene) Render() {
	clearScreen()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println(withIndentation(s.Image))
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	typeWrite(s.Body)
	fmt.Println("")

	items := make([]string, 0, len(s.Actions))
	for a := range s.Actions {
		items = append(items, a)
	}

	prompt := promptui.Select{
		Items:     items,
		Templates: promptuiTemplates,
	}

	_, choice, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	next, ok := s.Actions[choice]
	if !ok {
		log.Fatalf("%q is not a valid choice", choice)
	}
	next.Render()
}
