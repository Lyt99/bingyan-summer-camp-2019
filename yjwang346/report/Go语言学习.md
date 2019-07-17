

> **Go语言学习：**
>
> 一个相加的函数（例子）：
>
> > ```Go
> > package main
> > import "fmt"
> > func add(x int, y int) int {
> > 	return x + y
> > }
> > func main() {
> > 	fmt.Println(add(42, 13))
> > }
> > 类型只在变量名之后
> > ```
>
> ***var***语句用于声明一个变量列表，类型在最后
>
> ```go
> var i, j int = 1, 2
> var c, python, java = true, false, "no!"//初始值存在，可以省略类型
> ```
>
> ```go
> primes := [6]int{2, 3, 5, 7, 11, 13}
> ```
>
> ```go
> var s []int = primes[1:3]  //切片；更改切片的元素会修改其底层数组中对应的元素
> ```
>
> 切片部分以及映射等部分内容有些不明白，接着吸收*通过B站看视频明白了一些*
>
> 在Go中函数也是一种变量，我们可以通过`type`来定义它，它的类型就是所有拥有相同的参数，相同的返回值的一种类型，可以把这个类型的函数当做值来传递
>
> ```go
> ype testInt func(int) bool // 声明了一个函数类型
> /*前面部分函数声明省略了*/
> func filter(slice []int, f testInt) []int {
> var result []int
> for _, value := range slice {
>   if f(value) {
>       result = append(result, value)
>   }
> }
> return result
> }
> func main(){
> slice := []int {1, 2, 3, 4, 5, 7}
> fmt.Println("slice = ", slice)
> odd := filter(slice, isOdd)    // 函数当做值来传递了
> fmt.Println("Odd elements of slice are: ", odd)
> even := filter(slice, isEven)  // 函数当做值来传递了
> fmt.Println("Even elements of slice are: ", even)
> }//区别就是filter两次引用的函数发生了变化
> ```
>
> struct的匿名字段：能够实现字段的继承字段相同时最外层优先访问
>
> ***method***  是附属在一个给定的类型上的，他的语法和函数的声明语法几乎一样，只是在`func`后面增加了一个receiver(也就是method所依从的主体)。
>
> 可以定义在任何你自定义的类型、内置类型、struct等各种类型上面。
>
> method的语法如下：
>
> ```go
> func (r ReceiverType) funcName(parameters) (results)
> ```
>
> **method**的使用：(method area() 是属于Rectangle的)
>
> ```go
> type Rectangle struct {
> width, height float64
> }
> func (r Rectangle) area() float64 {
> return r.width*r.height
> }
> /*……*/
> fmt.Println("Area of r1 is: ", r1.area())
> ```
>
> **method的继承**：如果匿名字段实现了一个method，那么包含这个匿名字段的struct也能调用该								method
>
> **interface**：是一组method的组合
>
> golang对于不确定返回值可以用interface{}代替
>
> **并发**
>
> **channels**：可以通过channel接受或发送值 			***必须用创建channel***
>
> ```go
> c := make(chan int, 2)//缓存写满之后，代码将会阻塞，直到其他goroutine从channel 中读取							//一些元素，腾出空间
> fmt.Println(<-c)
> ```
>
> 
>
> 可以利用**select**来设置超时
>
> 

