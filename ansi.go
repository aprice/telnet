package telnet

// Telnet Options
const (
	TeloptBINARY         = byte(0)   // 8-bit data path
	TeloptECHO           = byte(1)   // echo
	TeloptRCP            = byte(2)   // prepare to reconnect
	TeloptSGA            = byte(3)   // suppress go ahead
	TeloptNAMS           = byte(4)   // approximate message size
	TeloptSTATUS         = byte(5)   // give status
	TeloptTM             = byte(6)   // timing mark
	TeloptRCTE           = byte(7)   // remote controlled transmission and echo
	TeloptNAOL           = byte(8)   // negotiate about output line width
	TeloptNAOP           = byte(9)   // negotiate about output page size
	TeloptNAOCRD         = byte(10)  // negotiate about CR disposition
	TeloptNAOHTS         = byte(11)  // negotiate about horizontal tabstops
	TeloptNAOHTD         = byte(12)  // negotiate about horizontal tab disposition
	TeloptNAOFFD         = byte(13)  // negotiate about formfeed disposition
	TeloptNAOVTS         = byte(14)  // negotiate about vertical tab stops
	TeloptNAOVTD         = byte(15)  // negotiate about vertical tab disposition
	TeloptNAOLFD         = byte(16)  // negotiate about output LF disposition
	TeloptXASCII         = byte(17)  // extended ascii character set
	TeloptLOGOUT         = byte(18)  // force logout
	TeloptBM             = byte(19)  // byte macro
	TeloptDET            = byte(20)  // data entry terminal
	TeloptSUPDUP         = byte(21)  // supdup protocol
	TeloptSUPDUPOUTPUT   = byte(22)  // supdup output
	TeloptSNDLOC         = byte(23)  // send location
	TeloptTTYPE          = byte(24)  // terminal type
	TeloptEOR            = byte(25)  // end or record
	TeloptTUID           = byte(26)  // TACACS user identification
	TeloptOUTMRK         = byte(27)  // output marking
	TeloptTTYLOC         = byte(28)  // terminal location number
	Telopt3270REGIME     = byte(29)  // 3270 regime
	TeloptX3PAD          = byte(30)  // X.3 PAD
	TeloptNAWS           = byte(31)  // window size
	TeloptTSPEED         = byte(32)  // terminal speed
	TeloptLFLOW          = byte(33)  // remote flow control
	TeloptLINEMODE       = byte(34)  // Linemode option
	TeloptXDISPLOC       = byte(35)  // X Display Location
	TeloptOLDENVIRON     = byte(36)  // Old - Environment variables
	TeloptAUTHENTICATION = byte(37)  // Authenticate
	TeloptENCRYPT        = byte(38)  // Encryption option
	TeloptNEWENVIRON     = byte(39)  // New - Environment variables
	TeloptEXOPL          = byte(255) // extended-options-list
)

// TelOpts gives you a mapping of index to a textual representation of the telnet option.
var TelOpts = []string{"BINARY", "ECHO", "RCP", "SUPPRESS GO AHEAD", "NAME",
	"STATUS", "TIMING MARK", "RCTE", "NAOL", "NAOP",
	"NAOCRD", "NAOHTS", "NAOHTD", "NAOFFD", "NAOVTS",
	"NAOVTD", "NAOLFD", "EXTEND ASCII", "LOGOUT", "BYTE MACRO",
	"DATA ENTRY TERMINAL", "SUPDUP", "SUPDUP OUTPUT",
	"SEND LOCATION", "TERMINAL TYPE", "END OF RECORD",
	"TACACS UID", "OUTPUT MARKING", "TTYLOC",
	"3270 REGIME", "X.3 PAD", "NAWS", "TSPEED", "LFLOW",
	"LINEMODE", "XDISPLOC", "OLD-ENVIRON", "AUTHENTICATION",
	"ENCRYPT", "NEW-ENVIRON"}

// sub-option qualifiers
const (
	TelQualIS    = byte(0) // option is...
	TelQualSEND  = byte(1) // send option
	TelQualINFO  = byte(2) // ENVIRON: informational version of IS
	TelQualREPLY = byte(2) // AUTHENTICATION: client version of IS
	TelQualNAME  = byte(3) // AUTHENTICATION: client version of IS

	LFlowOFF        = byte(0) // Disable remote flow control
	LFlowON         = byte(1) // Enable remote flow control
	LFlowRESTARTANY = byte(2) // Restart output on any char
	LFlowRESTARTXON = byte(3) // Restart output only on XON
)

// ENCRYPTion suboptions
const (
	EncryptIS       = byte(0) // I pick encryption type ...
	EncryptSUPPORT  = byte(1) // I support encryption types ...
	EncryptREPLY    = byte(2) // Initial setup response
	EncryptSTART    = byte(3) // Am starting to send encrypted
	EncryptEND      = byte(4) // Am ending encrypted
	EncryptREQSTART = byte(5) // Request you start encrypting
	EncryptREQEND   = byte(6) // Request you send encrypting
	EncryptENCKEYID = byte(7)
	EncryptDECKEYID = byte(8)
	EncryptCNT      = byte(9)

	EncTypeANY      = byte(0)
	EncTypeDESCFB64 = byte(1)
	EncTypeDESOFB64 = byte(2)
	EncTypeCNT      = byte(3)
)

// EncryptNames is a mapping of bytes to strings for Encrypt constants.
var EncryptNames = []string{
	"IS", "SUPPORT", "REPLY", "START", "END",
	"REQUEST-START", "REQUEST-END", "ENC-KEYID", "DEC-KEYID"}

// EncTypeNames is a mapping of bytes to strings for EncType constants.
var EncTypeNames = []string{
	"ANY", "DES_CFB64", "DES_OFB64"}

// Telnet IAC constants
const (
	xEOF  = byte(236)
	SUSP  = byte(237)
	ABORT = byte(238)
	EOR   = byte(239)
	SE    = byte(240)
	NOP   = byte(241)
	DM    = byte(242)
	BRK   = byte(243)
	IP    = byte(244)
	AO    = byte(245)
	AYT   = byte(246)
	EC    = byte(247)
	EL    = byte(248)
	GA    = byte(249)
	SB    = byte(250)
	WILL  = byte(251)
	WONT  = byte(252)
	DO    = byte(253)
	DONT  = byte(254)
	IAC   = byte(255)
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
