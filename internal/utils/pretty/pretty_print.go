package pretty

import (
	"fmt"
	"io"
	"strings"

	"github.com/Ayobami0/phoenix/internal/config"
)

type PrettyPrint struct {
	conf config.PrintConfig
}

func New(conf config.PrintConfig) PrettyPrint {
	return PrettyPrint{conf: conf}
}

func (p PrettyPrint) ClearScreen() {
	if p.conf.Silent {
		return
	}
	fmt.Print(CLSCR)
}

func (p PrettyPrint) PrintLogo() {
	if p.conf.Silent {
		return
	}
	fmt.Println(CYAN + LOGO + RESET)
	fmt.Println(BOLD + "System Configuration Tool" + RESET)
	fmt.Println(YELLOW + strings.Repeat("▔", 50) + RESET)
	fmt.Println()
}

func (p PrettyPrint) PrintSilentCompletion(message string) {
	if !p.conf.Silent {
		return
	}
	fmt.Println(GREEN + "✓ " + RESET + message)
}

func (p PrettyPrint) PrintStep(step string) {
	if !p.conf.Silent {
		fmt.Println(BLUE + STEP0 + BOLD + step + RESET)
	}
}

func (p PrettyPrint) Println(s ...any) {
	if !p.conf.Silent {
		fmt.Println(s...)
	}
}

func (p PrettyPrint) PrintFormat(f string, args ...any) {
	if !p.conf.Silent {
		fmt.Printf(f, args...)
	}
}

func (p PrettyPrint) PrintSuccess(message string) {
	if !p.conf.Silent {
		fmt.Println("  " + GREEN + "✓ " + RESET + message)
	}
}

func (p PrettyPrint) PrintWarning(message string) {
	if !p.conf.Silent {
		fmt.Println("  " + YELLOW + "! " + RESET + message)
	}
}

func (p PrettyPrint) PrintError(message string) {
		fmt.Println("  " + RED + "✗ " + RESET + message)
}

func (p PrettyPrint) PrintFooter() {
	if !p.conf.Silent {
		fmt.Println(GREEN + BOLD + "╔" + strings.Repeat("═", 48) + "╗" + RESET)
		fmt.Println(GREEN + BOLD + "║" + strings.Repeat(" ", 48) + "║" + RESET)
		fmt.Println(GREEN + BOLD + "║" + strings.Repeat(" ", 14) + "Setup completed!" + strings.Repeat(" ", 18) + "║" + RESET)
		fmt.Println(GREEN + BOLD + "║" + strings.Repeat(" ", 48) + "║" + RESET)
		fmt.Println(GREEN + BOLD + "╚" + strings.Repeat("═", 48) + "╝" + RESET)
	}
}

func (p PrettyPrint) FprintComment(file io.Writer, v string) {
	if p.conf.Compress {
		return
	}
	fmt.Fprintln(file, "# ", v)
}

func (p PrettyPrint) Fprintf(file io.Writer, f string, v ...any) {
	fmt.Fprintf(file, f, v...)
}

func (p PrettyPrint) Fprintln(file io.Writer, v ...any) {
	fmt.Fprintln(file, v...)
}

func (p PrettyPrint) FprintEcho(file io.Writer, v ...string) {
	if p.conf.Compress {
		return
	}
	var builder strings.Builder

	for i, val := range v {
		builder.WriteString(fmt.Sprintf(`echo "%s"`, val))
		if i < len(v)-1 {
			builder.WriteRune('\n')
		}
	}
	fmt.Fprintln(file, builder.String())
}

func (p PrettyPrint) FprintNewLine(file io.Writer) {
	fmt.Fprintln(file)
}

// func Spinner(message string, duration time.Duration) {
// 	spinChars := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
// 	start := time.Now()
//
// 	for time.Since(start) < duration {
// 		for _, char := range spinChars {
// 			fmt.Printf("\r  %s %s%s%s", char, Cyan, message, Reset)
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 	}
// 	fmt.Println("\r  " + Green + "✓ " + Reset + message + "   ")
// }
