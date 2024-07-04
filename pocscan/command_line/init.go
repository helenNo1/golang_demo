package command_line

type CommandLine struct {
	Src       string
	ThreadNum int
	Mode      string
	Proxy     string
}

func NewCommandLine() *CommandLine {
	return &CommandLine{}
}
