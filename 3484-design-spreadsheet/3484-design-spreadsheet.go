type Spreadsheet struct {
	sheet   []int
	rows    int
	columns int
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{
		sheet:   make([]int, rows*26),
		rows:    rows,
		columns: 26,
	}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
    this.sheet[this.flatIndex(cell)] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	this.sheet[this.flatIndex(cell)] = 0
}

func (this *Spreadsheet) GetValue(formula string) int {
	
    for i := 1; i < len(formula); i++ {
		if formula[i] == '+' {
			return this.processRef(formula[1:i]) + this.processRef(formula[i+1:])
		}
	}
    return 0
}

func (this *Spreadsheet) flatIndex(cell string) int {
	row := parseToInt(cell[1:]) - 1
	col := int(cell[0] - 'A')
    return row*this.columns + col
}

func (this *Spreadsheet) processRef(ref string) int {
    if ref[0] >= '0' && ref[0] <= '9' {
		return parseToInt(ref)
	}
	return this.sheet[this.flatIndex(ref)]
}

func parseToInt(ref string) int {
	results := 0
	for i := range ref {
		results = results * 10 + int(ref[i]-'0')
	}
	return results
}