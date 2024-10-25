package logger

const (
	ANSI_CODE_RESET string = "\x1b[0m"

	ANSI_CODE_BOLD            string = "\x1b[1m"
	ANSI_CODE_BOLD_RESET      string = "\x1b[22m"
	ANSI_CODE_DIM             string = "\x1b[2m"
	ANSI_CODE_DIM_RESET       string = "\x1b[22m"
	ANSI_CODE_ITALIC          string = "\x1b[3m"
	ANSI_CODE_ITALIC_RESET    string = "\x1b[23m"
	ANSI_CODE_UNDERLINE       string = "\x1b[4m"
	ANSI_CODE_UNDERLINE_RESET string = "\x1b[24m"
	ANSI_CODE_BLINK           string = "\x1b[5m"
	ANSI_CODE_BLINK_RESET     string = "\x1b[25m"
	ANSI_CODE_REVERSE         string = "\x1b[7m"
	ANSI_CODE_REVERSE_RESET   string = "\x1b[27m"
	ANSI_CODE_HIDE            string = "\x1b[8m"
	ANSI_CODE_HIDE_RESET      string = "\x1b[28m"
	ANSI_CODE_STRIKE          string = "\x1b[9m"
	ANSI_CODE_STRIKE_RESET    string = "\x1b[29m"

	ANSI_CODE_FOREGROUND_RESET   string = "\x1b[0m"
	ANSI_CODE_BACKGROUND_RESET   string = "\x1b[0m"
	ANSI_CODE_FOREGROUND_BLACK   string = "\x1b[30m"
	ANSI_CODE_BACKGROUND_BLACK   string = "\x1b[40m"
	ANSI_CODE_FOREGROUND_RED     string = "\x1b[31m"
	ANSI_CODE_BACKGROUND_RED     string = "\x1b[41m"
	ANSI_CODE_FOREGROUND_GREEN   string = "\x1b[32m"
	ANSI_CODE_BACKGROUND_GREEN   string = "\x1b[42m"
	ANSI_CODE_FOREGROUND_YELLOW  string = "\x1b[33m"
	ANSI_CODE_BACKGROUND_YELLOW  string = "\x1b[43m"
	ANSI_CODE_FOREGROUND_BLUE    string = "\x1b[34m"
	ANSI_CODE_BACKGROUND_BLUE    string = "\x1b[44m"
	ANSI_CODE_FOREGROUND_MAGENTA string = "\x1b[35m"
	ANSI_CODE_BACKGROUND_MAGENTA string = "\x1b[45m"
	ANSI_CODE_FOREGROUND_CYAN    string = "\x1b[36m"
	ANSI_CODE_BACKGROUND_CYAN    string = "\x1b[46m"
	ANSI_CODE_FOREGROUND_WHITE   string = "\x1b[37m"
	ANSI_CODE_BACKGROUND_WHITE   string = "\x1b[47m"
	ANSI_CODE_FOREGROUND_GRAY    string = "\x1b[90m"
	ANSI_CODE_FOREGROUND_DEFAULT string = "\x1b[39m"
	ANSI_CODE_BACKGROUND_DEFAULT string = "\x1b[49m"

	ANSI_CODE_ERASE_SCREEN_END   string = "\x1b[0Jm"
	ANSI_CODE_ERASE_SCREEN_START string = "\x1b[1Jm"
	ANSI_CODE_ERASE_SCREEN       string = "\x1b[2Jm"
	ANSI_CODE_ERASE_LINE_END     string = "\x1b[0K"
	ANSI_CODE_ERASE_LINE_START   string = "\x1b[1K"
	ANSI_CODE_ERASE_LINE         string = "\x1b[2K"
)

func (self Text) Bold() Text {
	return Text(ANSI_CODE_BOLD + string(self) + ANSI_CODE_BOLD_RESET)
}

func (self Text) Dim() Text {
	return Text(ANSI_CODE_DIM + string(self) + ANSI_CODE_DIM_RESET)
}

func (self Text) Italic() Text {
	return Text(ANSI_CODE_ITALIC + string(self) + ANSI_CODE_ITALIC_RESET)
}

func (self Text) Underline() Text {
	return Text(ANSI_CODE_UNDERLINE + string(self) + ANSI_CODE_UNDERLINE_RESET)
}

func (self Text) Blink() Text {
	return Text(ANSI_CODE_BLINK + string(self) + ANSI_CODE_BLINK_RESET)
}

func (self Text) Reverse() Text {
	return Text(ANSI_CODE_REVERSE + string(self) + ANSI_CODE_REVERSE_RESET)
}

func (self Text) Hide() Text {
	return Text(ANSI_CODE_HIDE + string(self) + ANSI_CODE_HIDE_RESET)
}

func (self Text) Strike() Text {
	return Text(ANSI_CODE_STRIKE + string(self) + ANSI_CODE_STRIKE_RESET)
}

func (self Text) BlackForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_BLACK + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) BlackBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_BLACK + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) RedForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_RED + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) RedBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_RED + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) GreenForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_GREEN + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) GreenBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_GREEN + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) YellowForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_YELLOW + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) YellowBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_YELLOW + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) BlueForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_BLUE + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) BlueBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_BLUE + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) MagentaForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_MAGENTA + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) MagentaBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_MAGENTA + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) CyanForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_CYAN + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) CyanBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_CYAN + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) WhiteForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_WHITE + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) WhiteBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_WHITE + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) GrayForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_GRAY + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) DefaultForeground() Text {
	return Text(ANSI_CODE_FOREGROUND_DEFAULT + string(self) + ANSI_CODE_FOREGROUND_RESET)
}

func (self Text) DefaultBackground() Text {
	return Text(ANSI_CODE_BACKGROUND_DEFAULT + string(self) + ANSI_CODE_BACKGROUND_RESET)
}

func (self Text) EraseScreenEnd() Text {
	return Text(string(self) + ANSI_CODE_ERASE_SCREEN_END)
}

func (self Text) EraseScreenStart() Text {
	return Text(string(self) + ANSI_CODE_ERASE_SCREEN_END)
}

func (self Text) EraseScreen() Text {
	return Text(string(self) + ANSI_CODE_ERASE_SCREEN)
}

func (self Text) EraseLineEnd() Text {
	return Text(string(self) + ANSI_CODE_ERASE_LINE_END)
}

func (self Text) EraseLineStart() Text {
	return Text(string(self) + ANSI_CODE_ERASE_LINE_END)
}

func (self Text) EraseLine() Text {
	return Text(string(self) + ANSI_CODE_ERASE_LINE)
}
