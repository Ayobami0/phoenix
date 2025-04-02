package pretty

const LOGO = `
  ____  _                      _
 |  _ \| |__   ___   ___ _ __ (_)_  __
 | |_) | '_ \ / _ \ / _ \ '_ \| \ \/ /
 |  __/| | | | (_) |  __/ | | | |>  <
 |_|   |_| |_|\___/ \___|_| |_|_/_/\_\
`

// ANSI Colors
const (
	RESET    = "\033[0m"
	BOLD     = "\033[1m"
	RED      = "\033[31m"
	GREEN    = "\033[32m"
	YELLOW   = "\033[33m"
	BLUE     = "\033[34m"
	MAGENTA  = "\033[35m"
	CYAN     = "\033[36m"
	WHITE    = "\033[37m"
	BGRED    = "\033[41m"
	BGGREEN  = "\033[42m"
	BGYELLOW = "\033[43m"
	BGBLUE   = "\033[44m"
	CLSCR    = "\033[H\033[2J"

	STEP0 = "⟹ "
	STEP1 = " ▶ "
	STEP2 = "  » "

	SUCCESS = "✓ "
	WARN    = "! "
	ERROR   = "✗ "

	FIRE = "🔥"

	TAB = "	"
)
