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
  if err := ui.Init(); err != nil {
    panic(err)
  }
  defer ui.Close()

  title := titleCard() 
  ui.Render(title)

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
  p.Text = "Go ants!"
  p.SetRect(0,0,25,5)
  return p
}
