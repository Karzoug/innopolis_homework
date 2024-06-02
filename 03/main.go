package main

type Formatter interface {
	Format(string) string
}

type plainText struct{}

func (p plainText) Format(s string) string {
	return s
}

type bold struct{}

func (b bold) Format(s string) string {
	return "**" + s + "**"
}

type italic struct{}

func (i italic) Format(s string) string {
	return "_" + s + "_"
}

type code struct{}

func (c code) Format(s string) string {
	return "`" + s + "`"
}

type chainFormatter []Formatter

func (c chainFormatter) Format(s string) string {
	for _, f := range c {
		s = f.Format(s)
	}
	return s
}

func (c *chainFormatter) AddFormatter(f Formatter) {
	*c = append(*c, f)
}

func main() {
}
