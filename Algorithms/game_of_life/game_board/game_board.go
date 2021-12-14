package game_board

import (
	"fmt"
	"game_of_life/cell"
	"math/rand"
	"time"
)

type GameBoard struct{
	row int
	col int
	grid [][] cell.Cell
	gen_num int
}

func New(row int, col int) GameBoard{
	my_grid := make([][]cell.Cell,row)
	for i := range my_grid {
		my_grid[i] = make([]cell.Cell, col)
		for j:= range my_grid[i]{
			my_grid[i][j]=cell.New()
		}
	}
	g:= GameBoard{row,col,my_grid,0}
	g.Generate()
	return g
}

func (g GameBoard)GetGenNum() int{
	return g.gen_num
}

func (g *GameBoard) Generate(){
	for i:= range(g.grid){
		for j:=range(g.grid[i]){
			source := rand.NewSource(time.Now().UnixNano())
			r := rand.New(source)
			num:=r.Intn(3)
			if num % 3 == 0 {
				g.grid[i][j].SetAlive()
			}
		}
	}
}

func (g *GameBoard)PrintGrid(){
	for _,row:=range(g.grid){
		for _,elem:=range(row){
			fmt.Printf(" %s ",elem.GetStatus())
			time.Sleep(100000000)
		}
		fmt.Printf("\n")
	}
}
func (g *GameBoard)TestPrintGrid(){
	for _,row:=range(g.grid){
		for _,elem:=range(row){
			fmt.Printf(" %s ",elem.TestGetStatus())
		}
		fmt.Printf("\n")
	}
}

func (g *GameBoard)FindNeighbour(x int,y int) []cell.Cell{
	var neighbours []cell.Cell

	if x>0 && y> 0{
		neighbours= append(neighbours, g.grid[x-1][y-1])
	}
	if y>0{
		neighbours= append(neighbours, g.grid[x][y-1])
	}
	if x>0{
		neighbours= append(neighbours, g.grid[x-1][y])
	}
	if x>0 && y < (g.col-1){
		neighbours= append(neighbours, g.grid[x-1][y+1])
	}
	if x < (g.row-1) && y>0{
		neighbours= append(neighbours, g.grid[x+1][y-1])
	}
	if x<(g.row -1){
		neighbours= append(neighbours, g.grid[x+1][y])
	}
	if y<(g.col -1){
		neighbours= append(neighbours, g.grid[x][y+1])
	}
	if x<(g.row -1) && y<(g.col -1){
		neighbours= append(neighbours, g.grid[x+1][y+1])
	}
	return neighbours
}

func (g GameBoard) countLiving(i int,j int) int{
	var neighbours []cell.Cell
	neighbours=g.FindNeighbour(i,j)
	count:=0
	for _,elem:=range neighbours{
		if elem.IsAlive(){
			count++
		}
	}
	return count
}

func (g *GameBoard) Update(){
	for i:=range g.grid{
		for j:=range g.grid[i] {
			if g.grid[i][j].IsAlive(){
				if g.countLiving(i,j) >3 || g.countLiving(i,j)<2{
					g.grid[i][j].SetDead()
				}
			}else{
				if g.countLiving(i,j) == 3{
					g.grid[i][j].SetAlive()
				}
			}
		}
	}
}