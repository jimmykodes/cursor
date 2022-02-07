package cursor

import (
	"fmt"
	"io"
	"os"
)

const (
	esc = "\x1b["
	// movement
	up    = "%dA"
	down  = "%dB"
	right = "%dC"
	left  = "%dD"
	toCol = "%dG"

	// cursor visibility
	hide = "?25l"
	show = "?25h"

	// buffer changing
	altBuff  = "?1049h"
	origBuff = "?1049l"

	// formatting
	clear         = "0m"
	bold          = "1m"
	dim           = "2m"
	italic        = "3m"
	underline     = "4m"
	blinking      = "5m"
	inverse       = "7m"
	invisible     = "8m"
	strikethrough = "9m"

	// colors
	color     = "%dm"
	blackFG   = 30
	blackBG   = 40
	redFG     = 31
	redBG     = 41
	greenFG   = 32
	greenBG   = 42
	yellowFG  = 33
	yellowBG  = 43
	blueFG    = 34
	blueBG    = 44
	magentaFG = 35
	magentaBG = 45
	cyanFG    = 36
	cyanBG    = 46
	whiteFG   = 37
	whiteBG   = 47

	// erasing
	clearScreen          = "J"
	clearToEndOfScreen   = "0J"
	clearToStartOfScreen = "1J"
	clearEntireScreen    = "2J"
	clearCurrentLine     = "K"
	clearToEndOfLine     = "0K"
	clearToStartOfLine   = "1K"
	clearEntireLine      = "2K"
)

func New(out io.Writer) *Cursor {
	return &Cursor{out: out}
}

func NewDefault() *Cursor {
	return New(os.Stderr)
}

type Cursor struct {
	out io.Writer
}

func (c *Cursor) escapeAndWrite(s string) *Cursor {
	fmt.Fprintf(c.out, "%s%s", esc, s)
	return c
}

func (c *Cursor) color(code int) *Cursor {
	return c.escapeAndWrite(fmt.Sprintf(color, code))
}

// visibility

func (c *Cursor) Hide() *Cursor {
	return c.escapeAndWrite(hide)
}
func (c *Cursor) Show() *Cursor {
	return c.escapeAndWrite(show)
}

// buffer

func (c *Cursor) AltBuffer() *Cursor {
	return c.escapeAndWrite(altBuff)
}

func (c *Cursor) OriginalBuffer() *Cursor {
	return c.escapeAndWrite(origBuff)
}

// movement

func (c *Cursor) Up(lines int) *Cursor {
	return c.escapeAndWrite(fmt.Sprintf(up, lines))
}
func (c *Cursor) Down(lines int) *Cursor {
	return c.escapeAndWrite(fmt.Sprintf(down, lines))
}

func (c *Cursor) Right(cols int) *Cursor {
	return c.escapeAndWrite(fmt.Sprintf(right, cols))
}

func (c *Cursor) Left(cols int) *Cursor {
	return c.escapeAndWrite(fmt.Sprintf(left, cols))
}
func (c *Cursor) ToCol(col int) *Cursor {
	return c.escapeAndWrite(fmt.Sprintf(toCol, col))
}

// formatting

func (c *Cursor) Clear() *Cursor {
	return c.escapeAndWrite(clear)
}
func (c *Cursor) Bold() *Cursor {
	return c.escapeAndWrite(bold)
}
func (c *Cursor) Dim() *Cursor {
	return c.escapeAndWrite(dim)
}
func (c *Cursor) Italic() *Cursor {
	return c.escapeAndWrite(italic)
}
func (c *Cursor) Underline() *Cursor {
	return c.escapeAndWrite(underline)
}
func (c *Cursor) Blinking() *Cursor {
	return c.escapeAndWrite(blinking)
}
func (c *Cursor) Inverse() *Cursor {
	return c.escapeAndWrite(inverse)
}
func (c *Cursor) Invisible() *Cursor {
	return c.escapeAndWrite(invisible)
}
func (c *Cursor) Strikethrough() *Cursor {
	return c.escapeAndWrite(strikethrough)
}

// colors

func (c *Cursor) Black() *Cursor {
	return c.color(blackFG)
}
func (c *Cursor) BlackBG() *Cursor {
	return c.color(blackBG)
}
func (c *Cursor) Red() *Cursor {
	return c.color(redFG)
}
func (c *Cursor) RedBG() *Cursor {
	return c.color(redBG)
}
func (c *Cursor) Green() *Cursor {
	return c.color(greenFG)
}
func (c *Cursor) GreenBG() *Cursor {
	return c.color(greenBG)
}
func (c *Cursor) Yellow() *Cursor {
	return c.color(yellowFG)
}
func (c *Cursor) YellowBG() *Cursor {
	return c.color(yellowBG)
}
func (c *Cursor) Blue() *Cursor {
	return c.color(blueFG)
}
func (c *Cursor) BlueBG() *Cursor {
	return c.color(blueBG)
}
func (c *Cursor) Magenta() *Cursor {
	return c.color(magentaFG)
}
func (c *Cursor) MagentaBG() *Cursor {
	return c.color(magentaBG)
}
func (c *Cursor) Cyan() *Cursor {
	return c.color(cyanFG)
}
func (c *Cursor) CyanBG() *Cursor {
	return c.color(cyanBG)
}
func (c *Cursor) White() *Cursor {
	return c.color(whiteFG)
}
func (c *Cursor) WhiteBG() *Cursor {
	return c.color(whiteBG)
}

// erasing

func (c *Cursor) ClearScreen() *Cursor {
	return c.escapeAndWrite(clearScreen)
}
func (c *Cursor) ClearToEndOfScreen() *Cursor {
	return c.escapeAndWrite(clearToEndOfScreen)
}
func (c *Cursor) ClearToStartOfScreen() *Cursor {
	return c.escapeAndWrite(clearToStartOfScreen)
}
func (c *Cursor) ClearEntireScreen() *Cursor {
	return c.escapeAndWrite(clearEntireScreen)
}
func (c *Cursor) ClearCurrentLine() *Cursor {
	return c.escapeAndWrite(clearCurrentLine)
}
func (c *Cursor) ClearToEndOfLine() *Cursor {
	return c.escapeAndWrite(clearToEndOfLine)
}
func (c *Cursor) ClearToStartOfLine() *Cursor {
	return c.escapeAndWrite(clearToStartOfLine)
}
func (c *Cursor) ClearEntireLine() *Cursor {
	return c.escapeAndWrite(clearEntireLine)
}
