package rendering

import (
  ui "github.com/gizak/termui/v3"
  "github.com/gizak/termui/v3/widgets"
  "image"
)


func Render() {
  // Variables
  dh := 100
  dw := 100
  // Check wether the UI can be initialized or not
  if err := ui.Init(); err != nil {
    panic(err)
  }
  defer ui.Close()

  // Elements
  // Title card
  //title := titleCard() 

  // Description
  //description := description()
  
  // Visualization
  canvas := canvas(dw, dh)

  // Render 
  ui.Render(canvas)

  // Quits if user presses a key
  uiEvents := ui.PollEvents()
  for {
    e := <-uiEvents
    if e.Type == ui.KeyboardEvent {
      break
    }
  }
}


// TITLE CARD
func titleCard() *widgets.Paragraph{
  p := widgets.NewParagraph()

  p.Text = "[GO ACO!](fg:yellow)"

  p.SetRect(0,0,20,3)
  return p
}

// DESCRIPTION
func description() *widgets.Paragraph{
  p := widgets.NewParagraph()
  
  p.Text = "[Here is my description and it is rad AF](fg:white)"

  p.SetRect(0,3,20,80)
  return p
}

// CANVAS

func canvas(width int, height int) *ui.Canvas{
  c := ui.NewCanvas()
  c.SetRect(0, 0, 50, 50)

  // Diagonal
  c.SetLine(image.Pt(0, 0), image.Pt(width,height), ui.ColorWhite)

  // Bottom
  c.SetLine(image.Pt(0, height), image.Pt(width,height), ui.ColorRed)

  // Top
  c.SetLine(image.Pt(0, 0), image.Pt(width,0), ui.ColorBlue)
  
  // Left
  c.SetLine(image.Pt(0, 0), image.Pt(1,height), ui.ColorYellow)

  // Right
  c.SetLine(image.Pt(width, 0), image.Pt(width-1,height), ui.ColorYellow)
  

  
  return c
}
