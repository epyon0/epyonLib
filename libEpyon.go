package libEpyon

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
		data = l.head.data
		l.head = l.head.next
		l.length--
	} else {
		err = fmt.Errorf("linkedlist is empty")
	}

	return data, err
}

// Peek returns the data at the head of a linkedlist without modification
func (l *LinkedList) Peek() any {
	return l.head.data
}

// IsEmpty returns if the linkedlist is empty or not
func (l *LinkedList) IsEmpty() bool {
	if l.head == nil {
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
	currentNode := l.head

	for i := 0; i < l.length; i++ {
		output = fmt.Sprintf("%s%d: %s\n", output, i, PrintValue(currentNode.data))
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
