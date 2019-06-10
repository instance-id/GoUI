package utils

//
//import (
//	term "github.com/nsf/termbox-go"
//	"strings"
//	"sync"
//)
//
//type Window interface {
//	Init() error
//	Draw(x, y, x1, y1 int)
//	HandleEvent(term.Event) (bool, error)
//}
//
//type TopIssueWindow struct {
//	Org         string
//	Target      string
//	Sort        string
//	Filter      string
//	Status      string
//	Alert       string
//	Focus       Window
//	ContextMenu Window
//	SortAsc     bool
//	drawSync    sync.Mutex
//
//	// Sub-Windows
//	Help              Window
//	Header            Window
//	FilterLine        Window
//	SortLine          Window
//	List              Window
//	ListMenu          Window
//	ListMilestoneMenu Window
//	ListPriorityMenu  Window
//	ListTypeMenu      Window
//	AlertModal        Window
//	StatusLine        Window
//}
//
//type Subwindow struct {
//	*TopIssueWindow
//}
//
//type HelpWindow struct {
//	*Subwindow
//}
//
//func utilHelp() {
//	printString(2, 0, "XTerm 256 color palette chart")
//	printString(screen.Width-19, 0, "F1 Help  F10 Exit")
//	helpText := []string{
//		"Control keys:",
//		"",
//		"F1 - show/hide this help",
//		"Up/Down - change lightness",
//		"Left/Right - change approximation method",
//		"",
//		"CIE76 - fastest",
//		"CIE94 - middle",
//		"CIE2000 - slowest",
//		"",
//		"",
//	}
//}
//
//// Draw the help window
//func (w *HelpWindow) Draw(x, y, x1, y1 int) {
//	if w.Focus != w.Help {
//		return
//	}
//
//	width, height := term.Size()
//	buffer := term.CellBuffer()
//	// dim the background
//	for ix := 0; ix < width; ix++ {
//		for iy := 0; iy < height; iy++ {
//			cell := buffer[iy*width+ix]
//			term.SetCell(ix, iy, cell.Ch, 235, cell.Bg)
//		}
//	}
//
//	// our overlay
//	overlay := `
//         **********************************************************************
//            ******************            ↳the current github search query
//              ↳sort +/- by a column
//   ↙  ↙  ↙   ↙
//  *** ****  ***  *****
//
//  ↙this number represents your milestone (0 means unassigned)
//  *
//   ↙this number represents your priority
//   *
//    ↙this number represents your type
//    *
//  *** ←together they are a sortable index, showing you the most relevant issues
//`
//	lines := strings.Split(overlay, "\n")
//	lines = lines[1:]
//	for iy, line := range lines {
//		for ix, c := range line {
//			fg := term.Attribute(5)
//			bg := term.ColorDefault
//			if c == '*' {
//				fg = term.ColorDefault | term.AttrUnderline
//				cell := buffer[iy*width+ix]
//				c = cell.Ch
//			} else if c == ' ' {
//				continue
//			}
//			term.SetCell(ix, iy, c, fg, bg)
//		}
//	}
//}
