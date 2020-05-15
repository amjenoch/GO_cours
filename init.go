package main

import "fmt"

var WhatIsThe = AnswerToLife()




func main() {
    if WhatIsThe == 0 {
        fmt.Println("It's all a lie.")
    }else {
		fmt.Println("05")
	}
}


func init() {
    WhatIsThe = 1
}




func AnswerToLife() int {
    return 42
}
