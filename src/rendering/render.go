package rendering

import (
  "fmt"
  ui "github.com/gizak/termui/v3"
  "github.com/gizak/termui/v3/widgets"
  "image"
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
  //title := titleCard() 

  // Description
  //description := description()
  
  // Visualization
  canvas := canvas()

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
func canvas() *ui.Canvas{
  c := ui.NewCanvas()
  c.SetRect(0, 0, 50, 50)

  // Diagonal
  c.SetLine(image.Pt(0, 0), image.Pt(50,50), ui.ColorWhite)

  // Bottom
  c.SetLine(image.Pt(0, 50), image.Pt(50,50), ui.ColorRed)

  // Top
  c.SetLine(image.Pt(0, 0), image.Pt(50,0), ui.ColorBlue)
  
  // Left
  c.SetLine(image.Pt(0, 0), image.Pt(1,50), ui.ColorYellow)

  // Right
  c.SetLine(image.Pt(50, 0), image.Pt(49,50), ui.ColorYellow)
  

  
  return c
}
