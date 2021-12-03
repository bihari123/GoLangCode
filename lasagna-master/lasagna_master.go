package lasagna

// TODO: define the 'PreparationTime()' function

   func PreparationTime(layers []string, avgTime int) int{
       if avgTime>0{
           return avgTime*len(layers)
       }
       return 2*len(layers)
   }

// TODO: define the 'Quantities()' function
   func Quantities(layers []string) (int,float64){
       noodles:=0
       sauce:=0.0
       for _,layer:=range layers{
           if layer == "noodles"{
               noodles+=50
           }else if layer == "sauce"{
               sauce+=0.2
           }
       }
       
       return noodles,sauce
   } 

// TODO: define the 'AddSecretIngredient()' function
   func AddSecretIngredient(friendList,myList []string) []string{
       finalRecipe:=append(myList,friendList[len(friendList)-1])
       return finalRecipe
   }

// TODO: define the 'ScaleRecipe()' function
   func ScaleRecipe(originalRecipe []float64, servings int) []float64{
       var oneServingRecipe []float64
       var totalServingRecipe []float64
       for index,_:=range originalRecipe{
           oneServingRecipe=append(oneServingRecipe,originalRecipe[index]/2)
       }
       
       for index,_:=range originalRecipe{
           totalServingRecipe=append(totalServingRecipe,oneServingRecipe[index]*float64(servings))
       }
       
       return totalServingRecipe
       
   }
