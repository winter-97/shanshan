package main
import(
"bufio"
"fmt"
"strings"
"os"
)
func check(e error){
if e != nil{
panic(e)
}
}
func main(){
f,err := os.Open("/home/go/file")
check(err)

b1 := bufio.NewScanner(f)
check(err)
m := ""
for n := 1; n <= 2 ; b1.Scan(){
m = b1.Text()
n = n + 1
}

s := strings.Split(m,"|")
for _,s := range s{
fmt.Printf("%s\n",s)
}
))

f.Close()
}
