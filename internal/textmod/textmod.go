package textmod

func ModifyText(text string) string {
	text = ConvertBinToDecimal(text)
	text = ConvertHexToDecimal(text)
	text = ConvertAtoAn(text)
	text = ChangeCase(text)
	text = ModifyPunctuationAndQuotes(text)
	return text
}
