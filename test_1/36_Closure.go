package main


func closureTest() {
    var str string = "hello World"
    f = func(){
    	str = "hello" + str
    }

    sunqi := f()
    tom := f()
    jerry := f()
}