package refine

func AllChecks (text string) string {
	text = HexToDecimal(text)
	text = BinToDecimal(text)
	text = Uppercase(text)
	text = Lowercase(text)
	text = Capitalize(text)
	text = Punctuations(text)
	
	return text
}