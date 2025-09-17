package ui

import "TheVengefulBlade/internal/models"

func ClearScreen()

func RenderHeader(title string)

func RenderMenu(options []string) int

func RenderInventory(c *models.Character) 

func RenderStatBar(label string, current, max int)

func PromptEnter()