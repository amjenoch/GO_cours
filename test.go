package main 

import "fmt"


	var y int 		= 42
	var z string	= "Salut,je suis un être heureux"
	var a string	= `James dit,"Salut,je suis un être heureux"`
	var c int 		= 1



	const (
		d = iota 
		b = iota 
		x = iota 
	)

	const (
		dd = iota
		bb
		xx 
		cc
		vv
		tt
	)
func main() {


		zz := 4
		zc := zz << 1

	fmt.Println(y)
	fmt.Printf("%T\n",y)
	fmt.Printf("%b\n",y)
	fmt.Printf("%x\n",y)
	fmt.Printf("%#x\n",y)



	fmt.Println(z)
	fmt.Printf("%T\n",z)
	

	fmt.Println(a)
	fmt.Printf("%T\n",a)

	fmt.Println(c<<10)

	fmt.Println(d)
	fmt.Println(b)
	fmt.Println(x)

	fmt.Println(dd)
	fmt.Println(bb)
	fmt.Println(xx)

	fmt.Println(cc)
	fmt.Println(vv)
	fmt.Println(tt)
	fmt.Printf("%d\t\t%b\n",zz,zz)
	fmt.Printf("%d\t\t%b\n",zc,zc)


}