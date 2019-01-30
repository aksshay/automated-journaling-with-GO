package compilemarkdown

func title(s string) string {
	return "# " + s + "\n"
}

func header(s string) string {
	return "## " + s + "\n"
}

func subHeader(s string) string {
	return "### " + s + "\n"
}

func italic(s string) string {
	return "*" + s + "*"
}

func lineBreak() string {
	return "*** \n"
}

func link(text string, link string) string {
	return "[" + text + "]" + "(" + link + ")"
}

func toTop() string {
	return "[Top](#top)\n"
}

func tableHeaders(headers []string) string {
	var s string
	for range headers {
		s = s + "|--------------------|"
	}
	s = s + "/n"
	s = s + tableCell(headers)
	for range headers {
		s = s + "|--------------------|"
	}
	s = s + "/n"
	return s
}

func tableCell(cell []string) string {
	var s string
	for _, c := range cell {
		s = s + "|" + c + "|"
	}
	s = s + "\n"
	return s
}
