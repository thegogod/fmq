package logger

type Text string

func (self Text) PadLeft(n int) Text {
	padding := ""

	for i := 0; i < n; i++ {
		padding += " "
	}

	return Text(padding + string(self))
}

func (self Text) PadRight(n int) Text {
	padding := ""

	for i := 0; i < n; i++ {
		padding += " "
	}

	return Text(string(self) + padding)
}

func (self Text) String() string {
	return string(self)
}
