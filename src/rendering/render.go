package rendering

import (
  ui "github.com/gizak/termui/v3"
  "github.com/gizak/termui/v3/widgets"
  "image"
  "math/rand"
  "time"
)

type pointList [][]int


func Render() {
  // Variables
  dh := 100 // Valid drawing area: dh - 1  
  dw := 100 // Valid drawing area: dw -1
  pl := generatePointList(20, dw, dh)

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
  canvas := canvas(dw, dh, pl)

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

func canvas(width int, height int, pl pointList) *ui.Canvas{
  c := ui.NewCanvas()
  c.SetRect(0, 0, 50, 50)

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
  for i := 0; i < len(pl)-1; i++ {
    p1 := pl[i]
    p2 := pl[i+1]
    x1, y1 := p1[0], p1[1]
    x2, y2 := p2[0], p2[1]
    c.SetLine(image.Pt(x1, y1), image.Pt(x2, y2), ui.ColorGreen)
  }
  
  // Point
  for i := 0; i < len(pl); i++ {
    point := pl[i]
    x,y := point[0], point[1]

    makePoint(x, y, width, height, c)
  }

  return c
}

// Generate a random pointlist 
func generatePointList(n int, width int, height int) pointList {
  s := rand.NewSource(time.Now().UnixNano())
  r := rand.New(s)

  var x int
  var y int
  
  pl := make(pointList, n)
  for i := 0; i < n; i++ {
    // We add 1 to h and w due to the possiblity of getting 0 (we subtract one after)
    randX := r.Intn(width) // Generatea a random int between [0,width+1)
    randY := r.Intn(height)// Generates a random int between [0,height+1)
    x = randX
    y = randY

    pl[i] = []int{x,y}
  }
  return pl
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

