package game

import "github.com/manifoldco/promptui"

var templates = &promptui.SelectTemplates{
	Label:    " ",
	Active:   `{{ "✓" | green }} {{ . }}`,
	Inactive: "  {{ . }}",
}
