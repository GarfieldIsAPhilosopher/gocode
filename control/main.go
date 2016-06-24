package main

import "fmt"
func main(){
if true==false {
fmt.Println("true")
}else{

fmt.Println("false")

}

 for i:=0;i<2;i++ {
fmt.Println("Go!")
}

for{
fmt.Println("loop")
break
}

j := true

for j{
fmt.Print("second loop")
j=false

}
for k:=0;k<30;k++{
fmt.Print(k)

if k%10==9{
fmt.Println("")
}


}
}
