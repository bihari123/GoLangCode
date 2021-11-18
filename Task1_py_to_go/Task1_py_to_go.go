package main
import ("fmt"
        "bufio"
        "os"
        "strings"
       ) 

func JoinString(x,y string) string{
    return strings.TrimRight(x,"\n")+" "+strings.TrimRight(y,"\n")
}

func PrintAll(list []string){
    for index,element:=range list{
        if len(element)>0{
            fmt.Printf("%d %s \n", index,element)
        }
     }    
}

func main(){
    list:= []string{}
    var input,str1,str2 string
    reader:=bufio.NewReader(os.Stdin)
    fmt.Printf("Welcome to your first program in golang. Follow the instructions below to proceed:")
    for true{
           fmt.Printf("Enter \n i for joining two strings \n u to display the contents of the list  \n s to exit\n")
           input,_=reader.ReadString('\n')
           if strings.TrimRight(input,"\n")=="i"{
               fmt.Printf("Enter the first string: ")
               str1,_=reader.ReadString('\n')
               fmt.Printf("Enter the second string: ")
               str2,_=reader.ReadString('\n')
               list=append(list,JoinString(str1,str2))
           }else if strings.TrimRight(input,"\n")=="u"{
               PrintAll(list)
           }else if strings.TrimRight(input,"\n") == "s"{
               fmt.Printf("You chose to exit the program: ")
               break
           }else{
               fmt.Printf("Please enter a valid value \n")
           }
    }
}    
