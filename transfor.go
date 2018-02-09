package main
import(
"bufio"
"fmt"
//"io"
//"io/ioutil"
"strings"
"os"
)
func check(e error){
if e != nil{
panic(e)
}
}
func main(){
//dat,err := ioutil.ReadFile("/home/go/file")
//check(err)
//fmt.Print(string(dat))

f,err := os.Open("/home/go/file")
check(err)

b1 := bufio.NewScanner(f)
//b1 := make([]byte,5)
//n1,err := f.Read(b1)
check(err)
m := ""
for n := 1; n <= 2 ; b1.Scan(){
//fmt.Println(b1.Text())
m = b1.Text()
n = n + 1
}

s := strings.Split(m,"|")
for _,s := range s{
fmt.Printf("%s\n",s)
}

//fmt.Printf("%d bytes: %s\n",n1,string(b1))

//o2,err := f.Seek(6,0)
//check(err)
//b2 := make([]byte,2)
//n2,err := f.Read(b2)
//check(err)
//fmt.Print("%d bytes @ %d: %s\n",n2,o2,string(b2))

//o3,err := f.Seek(6,0)
//check(err)
//b3 := make([]byte,2)
//n3,err := io.ReadAtLeast(f,b3,2)
//check(err)
//fmt.Printf("%d bytes @ %d: %s\n",n3,o3,string(b3))

//_,err = f.Seek(0,0)
//check(err)
//r4 := bufio.NewReader(f)
//b4,err := r4.Peek(5)
//check(err)
//fmt.Printf("5 bytes: %s\n",string(b4))

f.Close()
}
