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
  title := titleCard() 
  ui.Render(title)

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

  p.SetRect(0,0,50,3)
  return p
}
