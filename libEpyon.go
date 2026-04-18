package epyonLib

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"os/exec"
	"path"
	"reflect"
	"runtime"
	"time"
)

type Node struct {
	data any
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

// Returns the data from that node of the linkedlist
func (n *Node) Data() any {
	return n.data
}

// Returns the next node in the linkedlist
func (n *Node) Next() *Node {
	return n.next
}

// Returns node of the head of the linkedlist
func (l *LinkedList) Head() *Node {
	return l.head
}

// Returns length of linkedlist
func (l *LinkedList) Length() int {
	return l.length
}

// Adds item to the begining of a linkedlist
func (l *LinkedList) Push(data any) {
	newNode := &Node{data: data, next: l.head}
	l.head = newNode
	l.length++
}

// Removes item from the begining of a linkedlist
func (l *LinkedList) Pop() (any, error) {
	var data any
	var err error

	if l.length > 0 {
		data = l.Head().Data()
		l.head = l.Head().Next()
		l.length--
	} else {
		err = fmt.Errorf("linkedlist is empty")
	}

	return data, err
}

// Returns the data at the head of a linkedlist without modification
func (l *LinkedList) Peek() any {
	return l.Head().Data()
}

// Returns if the linkedlist is empty or not
func (l *LinkedList) IsEmpty() bool {
	if l.Head() == nil {
		return true
	} else {
		return false
	}
}

// Returns string to display the value of a given node
func (l *LinkedList) Print(node *Node) string {
	if node != nil {
		return PrintValue(node.data)
	}

	return ""
}

// Returns a string to display the value of the entire linkedlist
func (l *LinkedList) PrintAll() string {
	var output string
	currentNode := l.Head()

	fmt.Println("LENGTH: ", l.Length()-1, " >= ", int(math.Pow(10, 1)))

	for i := 0; i < l.Length(); i++ {
		switch {
		case l.Length()-1 >= int(math.Pow(10, 10)):
			output = fmt.Sprintf("%s%-11d: %s\n", output, i, PrintValue(currentNode.data))
		case l.Length()-1 >= int(math.Pow(10, 9)):
			output = fmt.Sprintf("%s%-10d: %s\n", output, i, PrintValue(currentNode.data))
		case l.Length()-1 >= int(math.Pow(10, 8)):
			output = fmt.Sprintf("%s%-9d: %s\n", output, i, PrintValue(currentNode.data))
		case l.Length()-1 >= int(math.Pow(10, 7)):
			output = fmt.Sprintf("%s%-8d: %s\n", output, i, PrintValue(currentNode.data))
		case l.Length()-1 >= int(math.Pow(10, 6)):
			output = fmt.Sprintf("%s%-7d: %s\n", output, i, PrintValue(currentNode.data))
		case l.Length()-1 >= int(math.Pow(10, 5)):
			output = fmt.Sprintf("%s%-6d: %s\n", output, i, PrintValue(currentNode.data))
		case l.Length()-1 >= int(math.Pow(10, 4)):
			output = fmt.Sprintf("%s%-5d: %s\n", output, i, PrintValue(currentNode.data))
		case l.Length()-1 >= int(math.Pow(10, 3)):
			output = fmt.Sprintf("%s%-4d: %s\n", output, i, PrintValue(currentNode.data))
		case l.Length()-1 >= int(math.Pow(10, 2)):
			output = fmt.Sprintf("%s%-3d: %s\n", output, i, PrintValue(currentNode.data))
		case l.Length()-1 >= int(math.Pow(10, 1)):
			output = fmt.Sprintf("%s%-2d: %s\n", output, i, PrintValue(currentNode.data))
		default:
			output = fmt.Sprintf("%s%d: %s\n", output, i, PrintValue(currentNode.data))
		}

		currentNode = currentNode.next
	}

	return output
}

// Creates a new linkedlist
func MakeLinkedList() *LinkedList {
	return &LinkedList{head: nil, length: 0}
}

// Returns byte slice of data from the pipe
func PipeRead(bufSize int) ([]byte, error) {
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, bufSize)
	var data []byte

	for {
		n, err := r.Read(buf[:cap(buf)])
		if err != nil && err != io.EOF {
			return data, err
		}

		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				return data, nil
			}
			return data, err
		}

		data = append(data, buf...)
		if err != nil && err != io.EOF {
			return data, err
		}
	}
}

// Return a boolean based on if a named pipe into the program exists
func PipeExists() (bool, error) {
	fi, err := os.Stdin.Stat()

	if err != nil {
		return false, err
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

// Truncate the given string to the given width
func TruncString(text string, width int) string {
	if width < 0 {
		return ""
	}

	runes := []rune(text)
	if len(runes) <= width {
		return text
	}
	return string(runes[:width])
}

// Clear the terminal on Windows or Linux
func ClearScreen() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	err := cmd.Run()
	return err
}

type Ansi struct {
	Line, Column, Count int
}

// Moves cursor to home position (0,0)
func (ansi Ansi) CurosrHome() {
	fmt.Fprintf(os.Stdout, "\033[H")
}

// Moves cursor to line Ansi.Line, column Ansi.Column
func (ansi Ansi) CursorMove() {
	fmt.Fprintf(os.Stdout, "\033[%d;%dH", ansi.Line, ansi.Column)
}

// Moves cursor up Ansi.Count lines
func (ansi Ansi) CursorUp() {
	fmt.Fprintf(os.Stdout, "\033[%dA", ansi.Count)
}

// Moves cursor down Ansi.Count lines
func (ansi Ansi) CursorDown() {
	fmt.Fprintf(os.Stdout, "\033[%dB", ansi.Count)
}

// Moves cursor right Ansi.Count lines
func (ansi Ansi) CursorRight() {
	fmt.Fprintf(os.Stdout, "\033[%dC", ansi.Count)
}

// Moves cursor left Ansi.Count lines
func (ansi Ansi) CursorLeft() {
	fmt.Fprintf(os.Stdout, "\033[%dD", ansi.Count)
}

// Moves cursor to beginning of next line, Ansi.Count lines down
func (ansi Ansi) CursorBeginningDown() {
	fmt.Fprintf(os.Stdout, "\033[%dE", ansi.Count)
}

// Moves cursor to beginning of previous line, Ansi.Count lines up
func (ansi Ansi) CursorBeginningUp() {
	fmt.Fprintf(os.Stdout, "\033[%dF", ansi.Count)
}

// Moves cursor to column Ansi.Column
func (ansi Ansi) CursorColumn() {
	fmt.Fprintf(os.Stdout, "\033[%dG", ansi.Column)
}

// Request cursor position
func (ansi Ansi) CursorReqPos() {
	fmt.Fprintf(os.Stdout, "\033[6n")
}

// Moves cursor one line up, scrolling if needed
func (ansi Ansi) CursorUpOne() {
	fmt.Fprintf(os.Stdout, "\033 M")
}

// Save cursor position (DEC) (recommended)
func (ansi Ansi) CursorSavePos() {
	fmt.Fprintf(os.Stdout, "\033 7")
}

// Restores the cursor the the last saved position (DEC)
func (ansi Ansi) CursorLoadPos() {
	fmt.Fprintf(os.Stdout, "\033 8")
}

// Save cursor position (SCO)
func (ansi Ansi) CursorSavePosSCO() {
	fmt.Fprintf(os.Stdout, "\033[s")
}

// Restores the cursor to the last saved position (SCO)
func (ansi Ansi) CursorLoadPosSCO() {
	fmt.Fprintf(os.Stdout, "\033[u")
}

// Erase from cursor until end of screen
func (ansi Ansi) CursorEraseScreenEnd() {
	fmt.Fprintf(os.Stdout, "\033[0J")
}

// Erase from cursor to beginning of screen
func (ansi Ansi) CursorEraseScreenBeginning() {
	fmt.Fprintf(os.Stdout, "\033[1J")
}

// Erase entire screen
func (ansi Ansi) CursorEraseScreenAll() {
	fmt.Fprintf(os.Stdout, "\033[2J")
}

// Erase saved lines
func (ansi Ansi) CursorEraseSavedLines() {
	fmt.Fprintf(os.Stdout, "\033[3J")
}

// Erase from cursor to end of line
func (ansi Ansi) CursorEraseLineEnd() {
	fmt.Fprintf(os.Stdout, "\033[0K")
}

// Erase start of line to the cursor
func (ansi Ansi) CursorEraseLineBeginning() {
	fmt.Fprintf(os.Stdout, "\033[1K")
}

// Erase the entire line
func (ansi Ansi) CursorEraseLineAll() {
	fmt.Fprintf(os.Stdout, "\033[2K")
}

// Make cursor invisible
func (ansi Ansi) CursorInvisible() {
	fmt.Fprintf(os.Stdout, "\033[?25l")
}

// Make cursor visible
func (ansi Ansi) CursorVisible() {
	fmt.Fprintf(os.Stdout, "\033[?25h")
}

// Restore screen
func (ansi Ansi) ScreenRestore() {
	fmt.Fprintf(os.Stdout, "\033[?47l")
}

// Save screen
func (ansi Ansi) ScreenSave() {
	fmt.Fprintf(os.Stdout, "\033[?47h")
}

// Enables the alternative buffer
func (ansi Ansi) EnableAltBuffer() {
	fmt.Fprintf(os.Stdout, "\033[?1049h")
}

// Disable the alternative buffer
func (ansi Ansi) DisableAltBuffer() {
	fmt.Fprintf(os.Stdout, "\033[?1049l")
}

// Reset all modes (styles and colors)
func (ansi Ansi) Reset() {
	fmt.Fprintf(os.Stdout, "\033[0m")
}

// Set bold mode
func (ansi Ansi) TextBold() {
	fmt.Fprintf(os.Stdout, "\033[1m")
}

// Reset bold mode (and dim/faint mode)
func (ansi Ansi) TextBoldReset() {
	fmt.Fprintf(os.Stdout, "\033[22m")
}

// Set dim/faint mode
func (ansi Ansi) TextDim() {
	fmt.Fprintf(os.Stdout, "\033[2m")
}

// Reset dim/faint mode (and bold mode)
func (ansi Ansi) TextDimReset() {
	ansi.TextBoldReset()
}

// Set dim/faint mode
func (ansi Ansi) TextFaint() {
	ansi.TextDim()
}

// Reset dim/faint mode (and bold mode)
func (ansi Ansi) TextFaintReset() {
	ansi.TextDimReset()
}

// Set italic mode
func (ansi Ansi) TextItalic() {
	fmt.Fprintf(os.Stdout, "\033[3m")
}

// Reset italic mode
func (ansi Ansi) TextItalicReset() {
	fmt.Fprintf(os.Stdout, "\033[23m")
}

// Set underline mode
func (ansi Ansi) TextUnderline() {
	fmt.Fprintf(os.Stdout, "\033[4m")
}

// Reset underline mode
func (ansi Ansi) TextUnderlineReset() {
	fmt.Fprintf(os.Stdout, "\033[24m")
}

// Set blinking mode
func (ansi Ansi) TextBlinking() {
	fmt.Fprintf(os.Stdout, "\033[5m")
}

// Reset blinking mode
func (ansi Ansi) TextBlinkingReset() {
	fmt.Fprintf(os.Stdout, "\033[25m")
}

// Set inverse/reverse mode
func (ansi Ansi) TextInverse() {
	fmt.Fprintf(os.Stdout, "\033[7m")
}

// Reset inverse/reverse mode
func (ansi Ansi) TextInverseReset() {
	fmt.Fprintf(os.Stdout, "\033[27m")
}

// Set inverse/reverse mode
func (ansi Ansi) TextReverse() {
	ansi.TextInverse()
}

// Reset inverse/reverse mode
func (ansi Ansi) TextReverseReset() {
	ansi.TextInverseReset()
}

// Set hidden/invisible mode
func (ansi Ansi) TextHidden() {
	fmt.Fprintf(os.Stdout, "\033[8m")
}

// Reset hidden/invisible mode
func (ansi Ansi) TextHiddenReset() {
	fmt.Fprintf(os.Stdout, "\033[28m")
}

// Set hidden/invisible mode
func (ansi Ansi) TextInvisible() {
	ansi.TextHidden()
}

// Reset hidden/invisible mode
func (ansi Ansi) TextInvisibleReset() {
	ansi.TextHiddenReset()
}

// Set strikethrough mode
func (ansi Ansi) TextStrikethrough() {
	fmt.Fprintf(os.Stdout, "\033[9m")
}

// Reset strikethrough mode
func (ansi Ansi) TextStrikethroughReset() {
	fmt.Fprintf(os.Stdout, "\033[29m")
}

// Set foreground color
func (ansi Ansi) ColorBlackFG() {
	fmt.Fprintf(os.Stdout, "\033[30m")
}

// Set background color
func (ansi Ansi) ColorBlackBG() {
	fmt.Fprintf(os.Stdout, "\033[40m")
}

// Set foreground color
func (ansi Ansi) ColorRedFG() {
	fmt.Fprintf(os.Stdout, "\033[31m")
}

// Set background color
func (ansi Ansi) ColorRedBG() {
	fmt.Fprintf(os.Stdout, "\033[41m")
}

// Set foreground color
func (ansi Ansi) ColorGreenFG() {
	fmt.Fprintf(os.Stdout, "\033[32m")
}

// Set background color
func (ansi Ansi) ColorGreenBG() {
	fmt.Fprintf(os.Stdout, "\033[42m")
}

// Set foreground color
func (ansi Ansi) ColorYellowFG() {
	fmt.Fprintf(os.Stdout, "\033[33m")
}

// Set background color
func (ansi Ansi) ColorYellowBG() {
	fmt.Fprintf(os.Stdout, "\033[43m")
}

// Set foreground color
func (ansi Ansi) ColorBlueFG() {
	fmt.Fprintf(os.Stdout, "\033[34m")
}

// Set background color
func (ansi Ansi) ColorBlueBG() {
	fmt.Fprintf(os.Stdout, "\033[44m")
}

// Set foreground color
func (ansi Ansi) ColorMagentaFG() {
	fmt.Fprintf(os.Stdout, "\033[35m")
}

// Set background color
func (ansi Ansi) ColorMagentaBG() {
	fmt.Fprintf(os.Stdout, "\033[45m")
}

// Set foreground color
func (ansi Ansi) ColorCyanFG() {
	fmt.Fprintf(os.Stdout, "\033[36m")
}

// Set background color
func (ansi Ansi) ColorCyanBG() {
	fmt.Fprintf(os.Stdout, "\033[46m")
}

// Set foreground color
func (ansi Ansi) ColorWhiteFG() {
	fmt.Fprintf(os.Stdout, "\033[37m")
}

// Set background color
func (ansi Ansi) ColorWhiteBG() {
	fmt.Fprintf(os.Stdout, "\033[47m")
}

// Set foreground color
func (ansi Ansi) ColorDefaultFG() {
	fmt.Fprintf(os.Stdout, "\033[39m")
}

// Set background color
func (ansi Ansi) ColorDefaultBG() {
	fmt.Fprintf(os.Stdout, "\033[49m")
}

// Set foreground color
func (ansi Ansi) ColorBrightBlackFG() {
	fmt.Fprintf(os.Stdout, "\033[90m")
}

// Set background color
func (ansi Ansi) ColorBrightBlackBG() {
	fmt.Fprintf(os.Stdout, "\033[100m")
}

// Set foreground color
func (ansi Ansi) ColorBrightRedFG() {
	fmt.Fprintf(os.Stdout, "\033[91m")
}

// Set background color
func (ansi Ansi) ColorBrightRedBG() {
	fmt.Fprintf(os.Stdout, "\033[101m")
}

// Set foreground color
func (ansi Ansi) ColorBrightGreenFG() {
	fmt.Fprintf(os.Stdout, "\033[92m")
}

// Set background color
func (anis Ansi) ColorBrightGreenBG() {
	fmt.Fprintf(os.Stdout, "\033[102m")
}

// Set foreground color
func (anis Ansi) ColorBrightYellowFG() {
	fmt.Fprintf(os.Stdout, "\033[93m")
}

// Set background color
func (ansi Ansi) ColorBrightYellowBG() {
	fmt.Fprintf(os.Stdout, "\033[103m")
}

// Set foreground color
func (ansi Ansi) ColorBrightBlueFG() {
	fmt.Fprintf(os.Stdout, "\033[94m")
}

// Set background color
func (ansi Ansi) ColorBrightBlueBG() {
	fmt.Fprintf(os.Stdout, "\033[104m")
}

// Set foreground color
func (ansi Ansi) ColorBrightMagentaFG() {
	fmt.Fprintf(os.Stdout, "\033[95m")
}

// Set background color
func (ansi Ansi) ColorBrightMagentaBG() {
	fmt.Fprintf(os.Stdout, "\033[105m")
}

// Set foreground color
func (ansi Ansi) ColorBrightCyanFG() {
	fmt.Fprintf(os.Stdout, "\033[96m")
}

// Set background color
func (ansi Ansi) ColorBrightCyanBG() {
	fmt.Fprintf(os.Stdout, "\033[106m")
}

// Set foreground color
func (ansi Ansi) ColorBrightWhiteFG() {
	fmt.Fprintf(os.Stdout, "\033[97m")
}

// Set background color
func (ansi Ansi) ColorBrightWhiteBG() {
	fmt.Fprintf(os.Stdout, "\033[107m")
}

// Set foreground 256-bit color of value
func (ansi Ansi) Color256FG(value byte) {
	fmt.Fprintf(os.Stdout, "\033[38;5;%dm", value)
}

// Set background 256-bit color of value
func (ansi Ansi) Color256BG(value byte) {
	fmt.Fprintf(os.Stdout, "\033[38;5;%dm", value)
}

// Set foreground RGB color
func (ansi Ansi) ColorRgbFG(red, green, blue byte) {
	fmt.Fprintf(os.Stdout, "\033[38;2;%d;%d;%dm", red, green, blue)
}

// Set background RGB color
func (ansi Ansi) ColorRgbBG(red, green, blue byte) {
	fmt.Fprintf(os.Stdout, "\033[48;2;%d;%d;%dm", red, green, blue)
}

/*
func (ansi Ansi) ScreenMode40x25Monochrome() {
	fmt.Fprintf(os.Stdout, "\033[=0h")
}

func (ansi Ansi) ScreenMode40x25Color() {
	fmt.Fprintf(os.Stdout, "\033[=1h")
}

func (ansi Ansi) ScreenMode80x25Monochrome() {
	fmt.Fprintf(os.Stdout, "\033[=2h")
}

func (ansi Ansi) ScreenMode80x25Color() {
	fmt.Fprintf(os.Stdout, "\033[=3h")
}

func (ansi Ansi) ScreenMode320x200Color() {
	fmt.Fprintf(os.Stdout, "\033[=4h")
}

func (ansi Ansi) ScreenMode320x200Monochrome() {
	fmt.Fprintf(os.Stdout, "\033[=5h")
}

func (ansi Ansi) ScreenMode640x200
*/

// Takes a string and formats and outputs to STDERR
func Verbose(text string, enabled bool) {
	now := time.Now()
	if enabled {
		pc, file, line, ok := runtime.Caller(1)

		if !ok {
			Er(fmt.Errorf("error getting caller function\n"))
		}

		msg := fmt.Sprintf("%02d:%02d:%02d.%04d | %v | %s:%d | VERBOSE: %v\n", now.Hour(), now.Minute(), now.Second(), now.Nanosecond()/1000000, runtime.FuncForPC(pc).Name(), path.Base(file), line, text)
		fmt.Fprintf(os.Stderr, "%s", msg)
	}
}

// Checks for error state, if found, prints the error to STDERR and exits
func Er(err error) {
	now := time.Now()
	pc, file, line, ok := runtime.Caller(1)

	if !ok {
		log.Fatal("error getting caller function\n")
		os.Exit(1)
	}

	if err != nil {
		msg := fmt.Sprintf("%02d:%02d:%02d.%04d | %v | %s:%d | ERROR: %v\n", now.Hour(), now.Minute(), now.Second(), now.Nanosecond()/1000000, runtime.FuncForPC(pc).Name(), path.Base(file), line, err)
		fmt.Fprintf(os.Stderr, "%s", msg)
		os.Exit(2)
	}
}

// Retruns a string of the human readabel form.
// i.e. 1024 retruns "1 KiB" or "1.03 KB" for SI Units
func HumanizeBytes(number int64, SIUnits bool) string {
	if SIUnits {
		switch {
		case number >= int64(math.Pow(10, 18)):
			return fmt.Sprintf("%.02f EB", float64(number)/math.Pow(10, 18))
		case number >= int64(math.Pow(10, 15)):
			return fmt.Sprintf("%.02f PB", float64(number)/math.Pow(10, 15))
		case number >= int64(math.Pow(10, 12)):
			return fmt.Sprintf("%.02f TB", float64(number)/math.Pow(10, 12))
		case number >= int64(math.Pow(10, 9)):
			return fmt.Sprintf("%.02f GB", float64(number)/math.Pow(10, 9))
		case number >= int64(math.Pow(10, 6)):
			return fmt.Sprintf("%.02f MB", float64(number)/math.Pow(10, 6))
		case number >= int64(math.Pow(10, 3)):
			return fmt.Sprintf("%.02f KB", float64(number)/math.Pow(10, 3))
		default:
			return fmt.Sprintf("%.02f B", float64(number))
		}
	} else {
		switch {
		case number >= int64(math.Pow(2, 60)):
			return fmt.Sprintf("%.02f EiB", float64(number)/math.Pow(2, 60))
		case number >= int64(math.Pow(2, 50)):
			return fmt.Sprintf("%.02f PiB", float64(number)/math.Pow(2, 50))
		case number >= int64(math.Pow(2, 40)):
			return fmt.Sprintf("%.02f TiB", float64(number)/math.Pow(2, 40))
		case number >= int64(math.Pow(2, 30)):
			return fmt.Sprintf("%.02f GiB", float64(number)/math.Pow(2, 30))
		case number >= int64(math.Pow(2, 20)):
			return fmt.Sprintf("%.02f MiB", float64(number)/math.Pow(2, 20))
		case number >= int64(math.Pow(2, 10)):
			return fmt.Sprintf("%.02f KiB", float64(number)/math.Pow(2, 10))
		default:
			return fmt.Sprintf("%.02f B", float64(number))
		}
	}
}

// Returns a string to display the value(s) and datatype(s)
func PrintValue(input any) string {
	var output string
	t := reflect.ValueOf(input).Kind()
	switch t {
	case reflect.Slice:
		output = fmt.Sprintf("%sslice[ ", output)
		for i := 0; i < reflect.ValueOf(input).Len(); i++ {
			output = fmt.Sprintf("%s%s ", output, PrintValue(reflect.ValueOf(input).Index(i).Interface()))
		}
		output = fmt.Sprintf("%s]", output)
	case reflect.Map:
		output = fmt.Sprintf("%smap[ ", output)
		for _, key := range reflect.ValueOf(input).MapKeys() {
			value := reflect.ValueOf(input).MapIndex(key).Interface()
			output = fmt.Sprintf("%s%s:%s ", output, PrintValue(key.Interface()), PrintValue(value))
		}
		output = fmt.Sprintf("%s]", output)
	case reflect.Struct:
		output = fmt.Sprintf("%sstruct%+v", output, input)
	default:
		output = fmt.Sprintf("%s%s(%v)", output, t, input)
	}
	return output
}
