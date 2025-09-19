type Spreadsheet struct {
	sheet   map[string]int
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{
		sheet: make(map[string]int),
	}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
    this.sheet[cell] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	this.sheet[cell] = 0
}

func (this *Spreadsheet) GetValue(formula string) int {
    for i := 1; i < len(formula); i++ {
		if formula[i] == '+' {
			return this.processRef(formula[1:i]) + this.processRef(formula[i+1:])
		}
	}
    return 0
}

func (this *Spreadsheet) processRef(ref string) int {
    if ref[0] >= '0' && ref[0] <= '9' {
		return parseToInt(ref)
	}
	return this.sheet[ref]
}

func parseToInt(ref string) int {
	results := 0
	for i := range ref {
		results = results * 10 + int(ref[i]-'0')
	}
	return results
}