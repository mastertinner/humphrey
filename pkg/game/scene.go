package game

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

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
	printIndent(s.Image)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	typeWrite(s.Body)
	fmt.Println("")

	items := make([]string, 0, len(s.Actions))
	for a := range s.Actions {
		items = append(items, a)
	}

	templates := &promptui.SelectTemplates{
		Label:    " ",
		Active:   `{{ "✓" | green }} {{ . }}`,
		Inactive: "  {{ . }}",
	}
	prompt := promptui.Select{
		Items:     items,
		Templates: templates,
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
