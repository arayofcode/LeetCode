type Spreadsheet struct {
    sheet [][]int
}


func Constructor(rows int) Spreadsheet {
    cells := make([]int, rows*26)
    sheet := make([][]int, rows)
    for i := range sheet {
        sheet[i] = cells[i * 26: (i+1) * 26]
    }

    return Spreadsheet{
        sheet: sheet,
    }
}


func (this *Spreadsheet) SetCell(cell string, value int)  {
    row, column := this.indexes(cell)
    this.sheet[row][column] = value
}


func (this *Spreadsheet) ResetCell(cell string)  {
    row, column := this.indexes(cell)
    this.sheet[row][column] = 0
}


func (this *Spreadsheet) GetValue(formula string) int {
    // formula always starts with an =
    refs := strings.Split(formula[1:], "+")
    return this.computeResults(refs)
}

func (this *Spreadsheet) indexes(cell string) (int, int) {
    row, _ := strconv.Atoi(cell[1:])
    return row - 1, int(cell[0] - 'A')
}

func (this *Spreadsheet) computeResults(refs []string) int {
    results := 0
    for _, ref := range refs {
        val := this.refVal(ref)
        results += val
    }
    return results
}

func (this *Spreadsheet) refVal(ref string) int {
    if ref[0] < 'A' {
        val, _ := strconv.Atoi(ref)
        return val
    }
    row, column := this.indexes(ref)
    return this.sheet[row][column]
}