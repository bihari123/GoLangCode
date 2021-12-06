package hamming
import "errors"
func Distance(a, b string) (int, error) {
    hD:=0


    if len(a) == len(b){
        for i:=0; i<len(a);i++{
            if a[i] != b[i]{
                hD++
            }
        }       
        
        return hD,nil
    }

    return hD,errors.New("two strings are not equal")
	panic("Implement the Distance function")
}
