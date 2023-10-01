package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/olekukonko/ts"
)

const localVersion = "1.9.0"

var bold = color.New(color.Bold)
var boldBlue = color.New(color.Bold, color.FgBlue)
var boldViolet = color.New(color.Bold, color.FgMagenta)
var codeText = color.New(color.BgBlack, color.FgGreen, color.Bold)
var stopSpin = false

var programLoop = true
var chatId = ""
var cookies = ""
var configDir = ""
var userInput = ""
var executablePath = ""

func main() {
	execPath, err := os.Executable()
	if err == nil {
		executablePath = execPath
	}
	terminate := make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-terminate
		os.Exit(0)
	}()
	
	args := os.Args

	if len(args) > 1 && len(args[1]) > 1 {
		input := args[1]

		if input == "-v" || input == "--version" {
			fmt.Println("tgpt", localVersion)
		} else if input == "-cl" || input == "--changelog" {
			getVersionHistory()
		} else if input == "-w" || input == "--whole" {
			if len(args) > 2 && len(args[2]) > 1 {
				prompt := args[2]
				trimmedPrompt := strings.TrimSpace(prompt)
				if len(trimmedPrompt) < 1 {
					fmt.Println("You need to provide some text")
					fmt.Println(`Example: tgpt -w "What is encryption?"`)
					os.Exit(0)
				}
				getWholeText(trimmedPrompt, configDir+"/tgpt")
			} else {
				fmt.Println("You need to provide some text")
				fmt.Println(`Example: tgpt -w "What is encryption?"`)
				os.Exit(0)
			}
		} else if input == "-q" || input == "--quiet" {
			if len(args) > 2 && len(args[2]) > 1 {
				prompt := args[2]
				trimmedPrompt := strings.TrimSpace(prompt)
				if len(trimmedPrompt) < 1 {
					fmt.Println("You need to provide some text")
					fmt.Println(`Example: tgpt -q "What is encryption?"`)
					os.Exit(0)
				}
				getSilentText(trimmedPrompt, configDir+"/tgpt")
			} else {
				fmt.Println("You need to provide some text")
				fmt.Println(`Example: tgpt -q "What is encryption?"`)
				os.Exit(0)
			}
		} else if input == "-s" || input == "--shell" {
			if len(args) > 2 && len(args[2]) > 1 {
				prompt := args[2]
				go loading(&stopSpin)
				trimmedPrompt := strings.TrimSpace(prompt)
				if len(trimmedPrompt) < 1 {
					fmt.Println("You need to provide some text")
					fmt.Println(`Example: tgpt -s "How to update system"`)
					os.Exit(0)
				}
				shellCommand(trimmedPrompt)
			} else {
				fmt.Println("You need to provide some text")
				fmt.Println(`Example: tgpt -s "How to update system"`)
				os.Exit(0)
			}

		} else if input == "-c" || input == "--code" {
			if len(args) > 2 && len(args[2]) > 1 {
				prompt := args[2]
				trimmedPrompt := strings.TrimSpace(prompt)
				if len(trimmedPrompt) < 1 {
					fmt.Println("You need to provide some text")
					fmt.Println(`Example: tgpt -c "Hello world in Python"`)
					os.Exit(0)
				}
				codeGenerate(trimmedPrompt)
			} else {
				fmt.Println("You need to provide some text")
				fmt.Println(`Example: tgpt -c "Hello world in Python"`)
				os.Exit(0)
			}
		} else if input == "-u" || input == "--update" {
			update()
		} else if input == "-i" || input == "--interactive" {
			/////////////////////
			// Normal interactive
			/////////////////////

			reader := bufio.NewReader(os.Stdin)
			bold.Print("Interactive mode started. Press Ctrl + C or type exit to quit.\n\n")
			for {
				boldBlue.Println("╭─ You")
				boldBlue.Print("╰─> ")

				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Error reading input:", err)
					break
				}

				if len(input) > 1 {
					input = strings.TrimSpace(input)
					if len(input) > 1 {
						if input == "exit" {
							bold.Println("Exiting...")
							return
						}
						getData(input, configDir+"/tgpt", true)
					}

				}

			}

		} else if input == "-m" || input == "--multiline" {
			/////////////////////
			// Multiline interactive
			/////////////////////
			fmt.Print("\nPress Tab to submit and Ctrl + C to exit.\n")

			for programLoop {
				fmt.Print("\n")
				p := tea.NewProgram(initialModel())
				_, err := p.Run()

				if err != nil {
					fmt.Println(err)
					os.Exit(0)
				}
				if len(userInput) > 0 {
					getData(userInput, configDir+"/tgpt", true)
				}

			}

		} else if strings.HasPrefix(input, "-") {
			boldBlue.Println(`Usage: tgpt [Flag] [Prompt]`)

			boldBlue.Println("\nFlags:")
			fmt.Printf("%-50v Generate and Execute shell commands. (Experimental) \n", "-s, --shell")
			fmt.Printf("%-50v Generate Code. (Experimental)\n", "-c, --code")
			fmt.Printf("%-50v Gives response back without loading animation\n", "-q, --quiet")
			fmt.Printf("%-50v Gives response back as a whole text\n", "-w, --whole")

			boldBlue.Println("\nOptions:")
			fmt.Printf("%-50v Remove key file \n", "-f, --forget")
			fmt.Printf("%-50v Print version \n", "-v, --version")
			fmt.Printf("%-50v Print help message \n", "-h, --help")
			fmt.Printf("%-50v Start normal interactive mode \n", "-i, --interactive")
			fmt.Printf("%-50v Start multi-line interactive mode \n", "-m, --multiline")
			fmt.Printf("%-50v See changelog of versions \n", "-cl, --changelog")

			if runtime.GOOS != "windows" {
				fmt.Printf("%-50v Update program \n", "-u, --update")
			}

			boldBlue.Println("\nExamples:")
			fmt.Println(`tgpt "What is internet?"`)
			fmt.Println("tgpt -f")
			fmt.Println(`tgpt -m`)
			fmt.Println(`tgpt -s "How to update my system?"`)
		} else {
			go loading(&stopSpin)
			formattedInput := strings.TrimSpace(input)
			getData(formattedInput, configDir+"/tgpt", false)
		}

	} else {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		go loading(&stopSpin)
		formattedInput := strings.TrimSpace(input)
		getData(formattedInput, configDir+"/tgpt", false)
	}
}

// Multiline input

type errMsg error

type model struct {
	textarea textarea.Model
	err      error
}

func initialModel() model {
	size, _ := ts.GetSize()
	termWidth := size.Col()
	ti := textarea.New()
	ti.SetWidth(termWidth)
	ti.CharLimit = 200000
	ti.ShowLineNumbers = false
	ti.Placeholder = "Enter your prompt"
	ti.Focus()

	return model{
		textarea: ti,
		err:      nil,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case tea.KeyCtrlC:
			programLoop = false
			userInput = ""
			return m, tea.Quit

		case tea.KeyTab:
			userInput = m.textarea.Value()

			if len(userInput) > 1 {
				m.textarea.Blur()
				return m, tea.Quit
			}

		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return m.textarea.View()
}

//////////////////////////////
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}