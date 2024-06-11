package route53applicationrecoverycontrollercellreadinessstatuschange

type CellStatusChange struct {
    NewState State `json:"new-state"`
    PreviousState State `json:"previous-state"`
    CellName string `json:"cell-name"`
}

func (c *CellStatusChange) SetNewState(newState State) {
    c.NewState = newState
}

func (c *CellStatusChange) SetPreviousState(previousState State) {
    c.PreviousState = previousState
}

func (c *CellStatusChange) SetCellName(cellName string) {
    c.CellName = cellName
}
