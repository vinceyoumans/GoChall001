# P0020 

1 + 2 = 3
2 x 3 + 1 = 7

================

2 1 + = 3

2. 2
   1 1
   +. 3
   =. 3


2 3 x 1 + = 7

2 2
3 3
x 6
1 1
+ 7
  last item   len(st)-1

3 2 / 1 + = 2.5
3 3
3 3
/. 3/2



//======================


func main(){
var st []int


//step 1
a = []bytes{"2 1 + = 3"}

aa := string.split(a, " ")
// aa []strings{"2", "3", "x", "1", "+", "="}



for _; aaa:= range aa{

isOp, iVal := isOpperand(aaa). //.  isOp := bool,   iopVal int number,

if  isOp{
// +. -, t

    x := len(st)
    
    if aaa =="+"{
      res:= st[x] + st[x-1]
      st = append(st, res)
    }else if aaa == "-"{
      res:= st[x] - st[x-1]
      st = append(st, res) 
    }elseif aaa =="x"{
      res:= st[x]+10 + st[x-1]
      st = append(st, res)
    }elseif aaa =="/"{
      res:= st[x]  st[x-1]
      st = append(st, res)
    }  
    
    
    }elseif aaa =="g"{
      res:= st[x]+10 + st[x-1]
      st = append(st, res)
    }    
}else{
st= append(st, ival)
}


}

func isOpperand(s)bool, float{

// bool
isop := false
switch s{
case "+":
isOp = true

    case"-":
    isoptrue = true
    
    case"/":
        isoptrue = true
    
    case"g":
        isoptrue = true
}


if !isOp{

    iOpVal := strcnv.AtoI(s)

}else{
iOpVal = 99999. // never used
}


return isop, iOpVal
}

