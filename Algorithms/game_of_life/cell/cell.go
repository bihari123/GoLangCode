package cell

type Cell struct {
	status string
}

func New() Cell{
	c:=Cell{"dead"}
	return c
}
func (c *Cell) SetDead() {
	c.status="dead"
}
func (c *Cell) SetAlive(){
	c.status="alive"
}
func(c Cell) IsAlive() bool{
	if c.status == "alive"{
		return true
	}else{
		return false
	}
}
func (c Cell) GetStatus() string{
	var stat string
	if c.status == "alive"{
		stat = "O"
	}else {
		stat="*"
	}
	return stat
}

func (c Cell) TestGetStatus() string{

	return c.status
}