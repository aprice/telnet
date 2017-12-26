package telnet

// Telnet IAC constants
const (
	SE   = byte(240)
	NOP  = byte(241)
	BRK  = byte(243)
	IP   = byte(244)
	AO   = byte(245)
	AYT  = byte(246)
	EC   = byte(247)
	EL   = byte(248)
	GA   = byte(249)
	SB   = byte(250)
	WILL = byte(251)
	WONT = byte(252)
	DO   = byte(253)
	DONT = byte(254)
	IAC  = byte(255)
)

const Escape = byte('\033')

// ANSI control sequences
var (
	// styles
	Reset        = []byte("\033[0m")
	Bold         = []byte("\033[1m")
	Underline    = []byte("\033[4m")
	Conceal      = []byte("\033[8m")
	NormalWeight = []byte("\033[22m")
	NoUnderline  = []byte("\033[24m")
	Reveal       = []byte("\033[28m")
	// colors - foreground
	FGBlack   = []byte("\033[30m")
	FGRed     = []byte("\033[31m")
	FGGreen   = []byte("\033[32m")
	FGYellow  = []byte("\033[33m")
	FGBlue    = []byte("\033[34m")
	FGMagenta = []byte("\033[35m")
	FGCyan    = []byte("\033[36m")
	FGWhite   = []byte("\033[37m")
	FGDefault = []byte("\033[39m")
	// background
	BGBlack   = []byte("\033[40m")
	BGRed     = []byte("\033[41m")
	BGGreen   = []byte("\033[42m")
	BGYellow  = []byte("\033[43m")
	BGBlue    = []byte("\033[44m")
	BGMagenta = []byte("\033[45m")
	BGCyan    = []byte("\033[46m")
	BGWhite   = []byte("\033[47m")
	BGDefault = []byte("\033[49m")
	// xterm
	TitleBarFmt = "\033]0;%s\a"
)
