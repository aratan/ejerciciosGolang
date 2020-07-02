package main

import (
	"fmt"
)

func main(){
	//varables
	var dato = []int{0,1,2,3,4,5,6,7,8,9,10,11,12}
	var n = len(dato) //longitud
    var x = 1  // nº secreto
    var p = n/2 //la mitad

    for i:=0; i < n; i++{

		if x > dato[(p)]{
			fmt.Printf("Sí, X es mayor que %v\n",p)
			p = p*2
			
		}else{
			fmt.Printf("Sí, X es menor que %v\n",p)
		    p = p/2						
			
		}
		
		if x == dato[(p)]{
			fmt.Printf("X es %v\nIteraciones:%v",p, i)
			break						
		} 	
    } 
}