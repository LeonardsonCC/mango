package spinner

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Spinner types.
var (
	Box1    = `⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏`
	Box2    = `⠋⠙⠚⠞⠖⠦⠴⠲⠳⠓`
	Box3    = `⠄⠆⠇⠋⠙⠸⠰⠠⠰⠸⠙⠋⠇⠆`
	Box4    = `⠋⠙⠚⠒⠂⠂⠒⠲⠴⠦⠖⠒⠐⠐⠒⠓⠋`
	Box5    = `⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠴⠲⠒⠂⠂⠒⠚⠙⠉⠁`
	Box6    = `⠈⠉⠋⠓⠒⠐⠐⠒⠖⠦⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈`
	Box7    = `⠁⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈⠈`
	Spin1   = `|/-\`
	Spin2   = `◴◷◶◵`
	Spin3   = `◰◳◲◱`
	Spin4   = `◐◓◑◒`
	Spin5   = `▉▊▋▌▍▎▏▎▍▌▋▊▉`
	Spin6   = `▌▄▐▀`
	Spin7   = `╫╪`
	Spin8   = `■□▪▫`
	Spin9   = `←↑→↓`
	Default = Box1
)

// Spinner is exactly what you think it is.
type Spinner struct {
	mu     sync.Mutex
	frames []rune
	length int
	pos    int
}

// New returns a spinner initialized with Default frames.
func New() *Spinner {
	s := &Spinner{}
	s.Set(Default)
	return s
}

// Set frames to the given string which must not use spaces.
func (s *Spinner) Set(frames string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.frames = []rune(frames)
	s.length = len(s.frames)
}

func Loading(loading <-chan struct{}, message string) {
	s := New()
	var strLen int

	isLoading := true
	for isLoading {
		select {
		case <-loading:
			isLoading = false
		default: // must have default for non-blocking op
		}

		str := fmt.Sprintf("\r%s %s", s.Next(), message)
		strLen = len(str)
		fmt.Print(str)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("\r%s\r", strings.Repeat(" ", strLen))
}

// Current returns the current rune in the sequence.
func (s *Spinner) Current() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	r := s.frames[s.pos%s.length]
	return string(r)
}

// Next returns the next rune in the sequence.
func (s *Spinner) Next() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	r := s.frames[s.pos%s.length]
	s.pos++
	return string(r)
}

// Reset the spinner to its initial frame.
func (s *Spinner) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.pos = 0
}
