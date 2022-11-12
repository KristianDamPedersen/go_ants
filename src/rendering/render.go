package rendering

import (
  ui "github.com/gizak/termui/v3"
  "github.com/gizak/termui/v3/widgets"
  "image"
)

type pointList [][]int


func Render() {
  // Variables
  dh := 100 // Valid drawing area: dh - 1  
  dw := 100 // Valid drawing area: dw -1
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

  pointList := make([][]int, 6)
  
  for i := 0; i < 3; i++ {
    pointList[i] = make([]int, 2)
  }

  pointList[0] = []int{0, 100}
  pointList[1] = []int{100,0}
  pointList[2] = []int{30, 100}
  pointList[3] = []int{0, 0}
  pointList[4] = []int{50,50}
  pointList[5] = []int{50, 70}


  // Diagonal
  // c.SetLine(image.Pt(0, 0), image.Pt(width,height), ui.ColorWhite)

  // Bottom
  c.SetLine(image.Pt(0, height), image.Pt(width,height), ui.ColorRed)

  // Top
  c.SetLine(image.Pt(0, 0), image.Pt(width,0), ui.ColorBlue)
  
  // Left
  c.SetLine(image.Pt(0, 0), image.Pt(1,height), ui.ColorYellow)

  // Right
  c.SetLine(image.Pt(width, 0), image.Pt(width-1,height), ui.ColorYellow)

  // Lines
  for i := 0; i < len(pointList)-1; i++ {
    p1 := pointList[i]
    p2 := pointList[i+1]
    x1, y1 := p1[0], p1[1]
    x2, y2 := p2[0], p2[1]
    c.SetLine(image.Pt(x1, y1), image.Pt(x2, y2), ui.ColorGreen)
  }
  
  // Point
  for i := 0; i < len(pointList); i++ {
    point := pointList[i]
    x,y := point[0], point[1]
    makePoint(x, y, width, height, c)
  }

  return c
}

// Make a point that fits within the valid drawing area.
func makePoint(x int, y int, width int, height int, c *ui.Canvas) *ui.Canvas{
  if x == width && y == height {
    // Check if its the bottom right corner
    c.SetLine(image.Pt(x, y), image.Pt(x-2, y), ui.ColorWhite)
    c.SetLine(image.Pt(x, y-1), image.Pt(x-2, y-1), ui.ColorWhite)
  } else if x == width {
    // Check if its on the right
    c.SetLine(image.Pt(x, y), image.Pt(x-2, y), ui.ColorWhite)
    c.SetLine(image.Pt(x, y+1), image.Pt(x-2, y+1), ui.ColorWhite)
  } else if y == height {
    // Check if its on the bottom
    c.SetLine(image.Pt(x, y), image.Pt(x+2, y), ui.ColorWhite)
    c.SetLine(image.Pt(x, y-1), image.Pt(x+2, y-1), ui.ColorWhite)
  } else {
    c.SetLine(image.Pt(x, y), image.Pt(x+2, y), ui.ColorWhite)
    c.SetLine(image.Pt(x, y+1), image.Pt(x+2, y+1), ui.ColorWhite)
  }
  return c
}

