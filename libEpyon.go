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

// Data func to export struct field
func (n *Node) Data() any {
	return n.data
}

// Next func to export struct field
func (n *Node) Next() *Node {
	return n.next
}

// Head func to export struct field
func (l *LinkedList) Head() *Node {
	return l.head
}

// Length func to export struct field
func (l *LinkedList) Length() int {
	return l.length
}

// Push adds item to the begining of a linkedlist
func (l *LinkedList) Push(data any) {
	newNode := &Node{data: data, next: l.head}
	l.head = newNode
	l.length++
}

// Pop removes item from the begining of a linkedlist
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

// Peek returns the data at the head of a linkedlist without modification
func (l *LinkedList) Peek() any {
	return l.Head().Data()
}

// IsEmpty returns if the linkedlist is empty or not
func (l *LinkedList) IsEmpty() bool {
	if l.Head() == nil {
		return true
	} else {
		return false
	}
}

// Print returns string to display the value of a given node
func (l *LinkedList) Print(node *Node) string {
	if node != nil {
		return PrintValue(node.data)
	}

	return ""
}

// PrintAll returns a string to display the value of the entire linkedlist
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

// MakeLinkedList creates a new linkedlist
func MakeLinkedList() *LinkedList {
	return &LinkedList{head: nil, length: 0}
}

// PipeRead returns byte slice of data from the pipe
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

// PipeExists return a boolean based on if a named pipe into the program exists
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

// TruncString will truncate the given string to the given width
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

// ClearScreen will clear the terminal on Windows or Linux
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

// CursorHome moves cursor to home position (0,0)
func (ansi Ansi) CurosrHome() {
	fmt.Fprintf(os.Stdout, "\033[H")
}

// CursorMove moves cursor to line Ansi.Line, column Ansi.Column
func (ansi Ansi) CursorMove() {
	fmt.Fprintf(os.Stdout, "\033[%d;%dH", ansi.Line, ansi.Column)
}

// CursorUp moves cursor up Ansi.Count lines
func (ansi Ansi) CursorUp() {
	fmt.Fprintf(os.Stdout, "\033[%dA", ansi.Count)
}

// CursorDown moves cursor down Ansi.Count lines
func (ansi Ansi) CursorDown() {
	fmt.Fprintf(os.Stdout, "\033[%dB", ansi.Count)
}

// CursorRight moves cursor right Ansi.Count lines
func (ansi Ansi) CursorRight() {
	fmt.Fprintf(os.Stdout, "\033[%dC", ansi.Count)
}

// CursorLeft moves cursor left Ansi.Count lines
func (ansi Ansi) CursorLeft() {
	fmt.Fprintf(os.Stdout, "\033[%dD", ansi.Count)
}

// CursorBeginningDown moves cursor to beginning of next line, Ansi.Count lines down
func (ansi Ansi) CursorBeginningDown() {
	fmt.Fprintf(os.Stdout, "\033[%dE", ansi.Count)
}

// CursorBeginningUp moves cursor to beginning of previous line, Ansi.Count lines up
func (ansi Ansi) CursorBeginningUp() {
	fmt.Fprintf(os.Stdout, "\033[%dF", ansi.Count)
}

// CursorColumn moves cursor to column Ansi.Column
func (ansi Ansi) CursorColumn() {
	fmt.Fprintf(os.Stdout, "\033[%dG", ansi.Column)
}

// CursorReqPos request cursor position
func (ansi Ansi) CursorReqPos() {
	fmt.Fprintf(os.Stdout, "\033[6n")
}

// CursorUpOne moves cursor one line up, scrolling if needed
func (ansi Ansi) CursorUpOne() {
	fmt.Fprintf(os.Stdout, "\033 M")
}

// CursorSavePos save cursor position (DEC) (recommended)
func (ansi Ansi) CursorSavePos() {
	fmt.Fprintf(os.Stdout, "\033 7")
}

// CursorLoadPos restores the cursor the the last saved position (DEC)
func (ansi Ansi) CursorLoadPos() {
	fmt.Fprintf(os.Stdout, "\033 8")
}

// CursorSavePosSCO save cursor position (SCO)
func (ansi Ansi) CursorSavePosSCO() {
	fmt.Fprintf(os.Stdout, "\033[s")
}

// CursorLoadPosSCO restores the cursor to the last saved position (SCO)
func (ansi Ansi) CursorLoadPosSCO() {
	fmt.Fprintf(os.Stdout, "\033[u")
}

// CursorEraseScreenEnd erase from cursor until end of screen
func (ansi Ansi) CursorEraseScreenEnd() {
	fmt.Fprintf(os.Stdout, "\033[0J")
}

// CursorEraseScreenBeginning erase from cursor to beginning of screen
func (ansi Ansi) CursorEraseScreenBeginning() {
	fmt.Fprintf(os.Stdout, "\033[1J")
}

// CursorEraseScreenAll erase entire screen
func (ansi Ansi) CursorEraseScreenAll() {
	fmt.Fprintf(os.Stdout, "\033[2J")
}

// CursorEraseSavedLines erase saved lines
func (ansi Ansi) CursorEraseSavedLines() {
	fmt.Fprintf(os.Stdout, "\033[3J")
}

// CursorEraseLineEnd erase from cursor to end of line
func (ansi Ansi) CursorEraseLineEnd() {
	fmt.Fprintf(os.Stdout, "\033[0K")
}

// CursorEraseLineBeginning erase start of line to the cursor
func (ansi Ansi) CursorEraseLineBeginning() {
	fmt.Fprintf(os.Stdout, "\033[1K")
}

// CursorEraseLineAll erase the entire line
func (ansi Ansi) CursorEraseLineAll() {
	fmt.Fprintf(os.Stdout, "\033[2K")
}

// Reset reset all modes (styles and colors)
func (ansi Ansi) Reset() {
	fmt.Fprintf(os.Stdout, "\033[0m")
}

// TextBold set bold mode
func (ansi Ansi) TextBold() {
	fmt.Fprintf(os.Stdout, "\033[1m")
}

// TextBoldReset reset bold mode (and dim/faint mode)
func (ansi Ansi) TextBoldReset() {
	fmt.Fprintf(os.Stdout, "\033[22m")
}

// TextDim set dim/faint mode
func (ansi Ansi) TextDim() {
	fmt.Fprintf(os.Stdout, "\033[2m")
}

// TextDimReset reset dim/faint mode (and bold mode)
func (ansi Ansi) TextDimReset() {
	ansi.TextBoldReset()
}

// TextFaint set dim/faint mode
func (ansi Ansi) TextFaint() {
	ansi.TextDim()
}

// TextFaintReset reset dim/faint mode (and bold mode)
func (ansi Ansi) TextFaintReset() {
	ansi.TextDimReset()
}

// TextItalic set italic mode
func (ansi Ansi) TextItalic() {
	fmt.Fprintf(os.Stdout, "\033[3m")
}

// TextItalicReset reset italic mode
func (ansi Ansi) TextItalicReset() {
	fmt.Fprintf(os.Stdout, "\033[23m")
}

// TextUnderline set underline mode
func (ansi Ansi) TextUnderline() {
	fmt.Fprintf(os.Stdout, "\033[4m")
}

// TextUnderlineReset reset underline mode
func (ansi Ansi) TextUnderlineReset() {
	fmt.Fprintf(os.Stdout, "\033[24m")
}

// TextBlinking set blinking mode
func (ansi Ansi) TextBlinking() {
	fmt.Fprintf(os.Stdout, "\033[5m")
}

// TextBlinkingReset reset blinking mode
func (ansi Ansi) TextBlinkingReset() {
	fmt.Fprintf(os.Stdout, "\033[25m")
}

// TextInverse set inverse/reverse mode
func (ansi Ansi) TextInverse() {
	fmt.Fprintf(os.Stdout, "\033[7m")
}

// TextInverseReset reset inverse/reverse mode
func (ansi Ansi) TextInverseReset() {
	fmt.Fprintf(os.Stdout, "\033[27m")
}

// TextReverse set inverse/reverse mode
func (ansi Ansi) TextReverse() {
	ansi.TextInverse()
}

// TextReverseReset reset inverse/reverse mode
func (ansi Ansi) TextReverseReset() {
	ansi.TextInverseReset()
}

// TextHidden set hidden/invisible mode
func (ansi Ansi) TextHidden() {
	fmt.Fprintf(os.Stdout, "\033[8m")
}

// TextHiddenReset reset hidden/invisible mode
func (ansi Ansi) TextHiddenReset() {
	fmt.Fprintf(os.Stdout, "\033[28m")
}

// TextInvisible set hidden/invisible mode
func (ansi Ansi) TextInvisible() {
	ansi.TextHidden()
}

// TextInvisibleReset reset hidden/invisible mode
func (ansi Ansi) TextInvisibleReset() {
	ansi.TextHiddenReset()
}

// TextStrikethrough set strikethrough mode
func (ansi Ansi) TextStrikethrough() {
	fmt.Fprintf(os.Stdout, "\033[9m")
}

// TextStrikthroughReset reset strikethrough mode
func (ansi Ansi) TextStrikethroughReset() {
	fmt.Fprintf(os.Stdout, "\033[29m")
}

/*
	Color.Black.Foreground = 30
	Color.Red.Foreground = 31
	Color.Green.Foreground = 32
	Color.Yellow.Foreground = 33
	Color.Blue.Foreground = 34
	Color.Magenta.Foreground = 35
	Color.Cyan.Foreground = 36
	Color.White.Foreground = 37

	Color.Black.Background = 40
	Color.Red.Background = 41
	Color.Green.Background = 42
	Color.Yellow.Background = 43
	Color.Blue.Background = 44
	Color.Magenta.Background = 45
	Color.Cyan.Background = 46
	Color.White.Background = 47
*/

func (ansi Ansi) ColorCyanFG() {
	fmt.Fprintf(os.Stdout, "\033[36m")
}

func (ansi Ansi) ColorCyanBG() {
	fmt.Fprintf(os.Stdout, "\033[46m")
}

// Verbose takes a string and formats and outputs to STDERR
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

// Er checks for error state, if found, prints the error to STDERR and exits
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

// HumanizeBytes retruns a string of the human readabel form.
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

// PrintValue returns a string to display the value(s) and datatype(s)
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
