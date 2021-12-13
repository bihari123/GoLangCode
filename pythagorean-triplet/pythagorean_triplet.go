package pythagorean



type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	a, b, c := 0, 0, 0
	var arr []Triplet
	m := 2
	for c <= max {
		for n := 1; n < m; n++ {
			a = m*m - n*n
			b = 2 * m * n
			c = m*m + n*n
			if c > max {
				break
			}
			if a>b{
				b,a=a,b
			}
			if a>=min && b>=min && c>=min{
				arr = append(arr, Triplet{a, b, c})
			}

		}
		m++

	}
	return arr
	panic("Please implement the Range function")
}
func contains(s Triplet,t []Triplet)bool{
	for r:=range t{
		for c:=range t[r]{
			if t[r][c]==s[c]  {
				return true
			}
		}
	}
	return false
}
// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {

	var arr []Triplet
	for i:=1;i<p;i++{
		for j:=1;j<p;j++{
			a:=i
			b:=j
			c:=p-a-b
			if (a*a + b*b) == c*c {
				if a>b {
					a,b=b,a
				}
				if !contains(Triplet{a,b,c},arr){
					arr=append(arr,Triplet{a,b,c})
				}

			}
		}
	}
	return arr

	panic("Please implement the Sum function")
}
