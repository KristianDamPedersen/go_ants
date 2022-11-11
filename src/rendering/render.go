package rendering

import (
  "fmt"
  ui "github.com/gizak/termui/v3"
  "github.com/gizak/termui/v3/widgets"
)

func TestRendering() {
  fmt.Println("I am from rendering package")
}

func Render() {

  // Check wether the UI can be initialized or not
  if err := ui.Init(); err != nil {
    panic(err)
  }
  defer ui.Close()

  // Elements
  // Title card
  title := titleCard() 

  // Description
  description := description()
  
  // Visualization

  // Render 
  ui.Render(title, description)

  // Quits if user presses a key
  uiEvents := ui.PollEvents()
  for {
    e := <-uiEvents
    if e.Type == ui.KeyboardEvent {
      break
    }
  }
}

func titleCard() *widgets.Paragraph{
  p := widgets.NewParagraph()

  p.Text = "[GO ACO!](fg:yellow)"

  p.SetRect(0,0,20,3)
  return p
}

func description() *widgets.Paragraph{
  p := widgets.NewParagraph()
  
  p.Text = "[Here is my description and it is rad AF](fg:white)"

  p.SetRect(0,3,20,80)
  return p
}
