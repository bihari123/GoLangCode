package matrix

import (
	"errors"
	"unicode"
)

// Define the Matrix type here.
type Matrix struct {
	col,row int
	arr [][]int
}
func (m *Matrix)fillMatrix(s string)error{
	m.col=0
	m.row=0
	for i:=0;i<len(s);i++{

		//fmt.Println("checking string")
		if s[i]!=' '{
			if s[i] !='\n'{
				var elem =0
				elem=(int(s[i])-int('0'))
				if  i<(len(s)-1) && s[i+1]!=' ' && s[i+1] != '\n' {


					for i<(len(s)-1) && s[i+1]!=' ' && s[i+1] != '\n' {
						if !unicode.IsDigit(rune(s[i])){
							return errors.New("not an integer")
						}
						elem=elem*10 +(int(s[i+1])-int('0'))
						//fmt.Println(i,elem)
						i++

					}



				}
				m.arr[m.row]=append(m.arr[m.row],elem)


			}else{
				m.row++
				m.arr=append(m.arr,[]int{})
			}
		}

	}
	return nil
}

func (m *Matrix)validate() (bool,error){
	col:=len(m.arr[0])
	for i:=range m.arr{
		if len(m.arr[i]) !=col{
			//fmt.Println(m)
			return false,errors.New("different number of columns")
		}
		for j:=range m.arr[i]{
			v:=m.arr[i][j]
			 if v<0 {
				return false, errors.New("negative value")
			}
		}
	}
	return true,nil
}

func New(s string) (*Matrix, error) {
	var m =Matrix{0,0,[][]int{{}}}
    err1:=m.fillMatrix(s)
	if err1 == nil{
		val,err:=m.validate()
		if !val{
			return &m,err
		}
	}else{
		return &m,err1
	}
	//fmt.Println(m)

	return &m,nil
	panic("Please implement the New function")
}

// Cols and Rows must return the results without affecting the matrix.

func (m *Matrix) Cols() [][]int {
	var m2 =make([][]int,len(m.arr[0]))
	for i:=range m2{
		m2[i]=append(m2[i],make([]int,len(m.arr))...)
	}
	for row,_:= range m.arr{
		for col,c:=range m.arr[row]{
			m2[col][row]=c
		}
	}
	return m2
	panic("Please implement the Cols function")
}

func (m *Matrix) Rows() [][]int {
	var m2 =make([][]int,len(m.arr))
	for row,_:= range m.arr{
		//m2[row]=append(m2[row],make([]int,len(m.arr[row]))...)
		for _,c:=range m.arr[row]{
			m2[row]=append(m2[row],c)
		}
	}
	return m2
	panic("Please implement the Rows function")
}

func (m *Matrix) Set(row, col, val int) bool {
	if row>=0 && row<len(m.arr) && col>=0 && col <len(m.arr[0]){
		m.arr[row][col]=val
	    return true
	}
	return false
	panic("Please implement the Set function")
}
