package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		type: 用于类型定义和类型别名

			1. 类型定义：type 类型名 Type
			2. 类型别名：type 类型名 = Type

	*/
	var i1 myint //main.myint类
	var i2 = 100
	i1 = 200
	fmt.Println(i1, i2)

	var name mystring //main.mystring类
	name = "小王"
	var s1 string
	s1 = "小李"
	fmt.Println(name, s1)

	fmt.Printf("%T,%T,%T,%T\n", i1, i2, name, s1) //main.myint,int,main.mystring,string

	fmt.Println("-----------------------")
	res1 := fun1()            // fun1()的返回值是一个函数。也就是res1是返回的函数func(int,int)string
	fmt.Println(res1(10, 11)) //调用函数res1(10,11)

	fmt.Println("-----------------------")
	var i3 myint2
	i3 = i2
	fmt.Printf("%T,%d\n", i3, i3) //int,100  说明myint2和int是一样的 myint2只是别名。

	var s Student
	//s.name = "小王" //ambiguous selector s.name 不明确的选择
	s.Person.name = "小王" //结构体中的结构体本来就应该这样赋值，但是由于可识别，一般就简写了，也就是“提升字节”。但是该情况，程序不可区分，所以就具体写明了赋值的对象。
	//s.fun1() //ambiguous selector s.fun1
	fmt.Printf("%T,%T\n", s.Person, s.People)

	s.Person.fun1()

	s.People.name = "小李"
	s.People.fun2()
	s.People.fun2()

	s.Person.fun2()

}

// 1. 定义一个新的类型
type myint int
type mystring string

// 2. 定义函数类型  当函数的参数或返回值是一个复杂的函数时，可以用定义函数类型，简化。
type myfun func(int, int) string

func fun1() myfun { //fun1函数的返回值，是一个myfun类型的函数。
	fun := func(a, b int) string {
		s := strconv.Itoa(a) + strconv.Itoa(b) //Itoa把 int 转化成 string
		return s
	}
	return fun
}

// 3. 定义 类型别名
type myint2 = int //不是重新定义新的数据类型，只是给int起了别名，和int可以换用。

// 4. 非本地类型不能定义方法
/*
package main
import(
	"time"
)
//定义time.Duration的别名为MyDuration
type MyDuration = time.Duration // 怎么解决这个问题呢？方法1：在原time.Duration内添加需要的方法（函数）。方法2：去掉等号 type MyDuration time.Duration
//为MyDuration添加一个函数
func (m MyDuration) EasySet(a string){
}

不能执行，因为这相当于给原time.Duration添加新方法。如果能执行，这将很危险。

func main(){
}


*/

// 5. 非本地类型不能定义方法
type Person struct {
	name string
}

type People = Person

func (p Person) fun1() { //People = Person 都可以用这个方法
	fmt.Println("Person --> ", p.name)
}

func (p People) fun2() { //函数名不能相同  //People = Person 都可以用这个方法
	fmt.Println("People --> ", p.name)
}

type Student struct {
	Person
	People
}
