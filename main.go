package main

import (
	"fmt"
	"reflect"
	"strconv"
	"study/reptile"
	"time"
)

type FibonacciFunc func(int) int

const Max = 50

var fibs [Max]int

func fibonacci(a int) int {
	if a == 1 {
		return 0
	}
	if a == 2 {
		return 1
	}
	index := a - 1
	if fibs[index] != 0 {
		return fibs[index]
	}
	num := fibonacci(a-1) + fibonacci(a-2)
	fibs[index] = num

	return num
}

func fibonacciTail(n, first, second int) int {
	if n < 2 {
		return first
	}

	return fibonacciTail(n-1, second, first+second)
}

func fibonacci1(n int) int {
	return fibonacciTail(n, 0, 1)
}

func fibonacciExecTime(plyFunc FibonacciFunc) FibonacciFunc {
	return func(a int) int {
		start := time.Now()
		c := plyFunc(a)
		end := time.Since(start)
		fmt.Printf("耗时:%v\n", end)
		return c
	}
}

func ageSum(users []map[string]string) int {
	var sum int
	for _, user := range users {
		num, _ := strconv.Atoi(user["age"])
		sum += num
	}

	return sum
}

func mapToString(items []map[string]string, f func(map[string]string) string) []string {
	newSile := make([]string, len(items))
	for _, item := range items {
		newSile = append(newSile, f(item))
	}

	return newSile
}

func fieldSum(items []string, f func(string) int) int {
	var sum int
	for _, item := range items {
		sum += f(item)
	}

	return sum
}

type A interface {
	Foo()
}
type B interface {
	A
	Bar()
}

type T struct{}

func (t T) Foo() {
	fmt.Println("call Foo function from interface A.")
}
func (t T) Bar() {
	fmt.Println("call Bar function from interface B.")
}

type Container struct {
	s reflect.Value
}

// 通过传入存储元素类型和容量来初始化容器
func NewContainer(t reflect.Type, size int) *Container {
	if size <= 0 {
		size = 64
	}
	// 基于切片类型实现这个容器，这里通过反射动态初始化这个底层切片
	return &Container{
		s: reflect.MakeSlice(reflect.SliceOf(t), 0, size),
	}
}

// 添加元素到容器，通过空接口声明传递的元素类型，表明支持任何类型
func (c *Container) Put(val interface{}) error {
	// 通过反射对实际传递进来的元素类型进行运行时检查，
	// 如果与容器初始化时设置的元素类型不同，则返回错误信息
	// c.s.Type() 对应的是切片类型，c.s.Type().Elem() 应的才是切片元素类型
	if reflect.ValueOf(val).Type() != c.s.Type().Elem() {
		return fmt.Errorf("put error: cannot put a %T into a slice of %s",
			val, c.s.Type().Elem())
	}
	// 如果类型检查通过则将其添加到容器中
	c.s = reflect.Append(c.s, reflect.ValueOf(val))
	return nil
}

// 从容器中读取元素，将返回结果赋值给 val，同样通过空接口指定元素类型
func (c *Container) Get(val interface{}) error {
	// 还是通过反射对元素类型进行检查，如果不通过则返回错误信息
	// Kind 与 Type 相比范围更大，表示类别，如指针，而 Type 则对应具体类型，如 *int
	// 由于 val 是指针类型，所以需要通过 reflect.ValueOf(val).Elem() 获取指针指向的类型
	if reflect.ValueOf(val).Kind() != reflect.Ptr ||
		reflect.ValueOf(val).Elem().Type() != c.s.Type().Elem() {
		return fmt.Errorf("get error: needs *%s but got %T", c.s.Type().Elem(), val)
	}
	// 将容器第一个索引位置值赋值给 val 指针
	reflect.ValueOf(val).Elem().Set(c.s.Index(0))
	// 然后删除容器第一个索引位置值
	c.s = c.s.Slice(1, c.s.Len())
	return nil
}

//func divide() {
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Println("Runtime panic caught: %v\\n", err)
//		}
//	}()
//
//	var i = 1
//	var j = 0
//	k := i / j
//	fmt.Printf("%d / %d = %d\n", i, j, k)
//}

func main() {

	//斐波那契
	//f := 50
	//fibonacci := fibonacciExecTime(fibonacci)
	//r := fibonacci(f)
	//fmt.Printf("The %dth number of fibonacci sequence is %d\n\n", f, r)
	//
	//fibonacci1 := fibonacciExecTime(fibonacci1)
	//s := fibonacci1(f)
	//fmt.Printf("The %dth number of fibonacci1 sequence is %d\n\n", f, s)

	//求和

	//users := []map[string]string{
	//	{
	//		"name": "张三",
	//		"age":  "18",
	//	},
	//	{
	//		"name": "李四",
	//		"age":  "22",
	//	},
	//	{
	//		"name": "王五",
	//		"age":  "20",
	//	},
	//}

	//fmt.Printf("用户年龄累加结果: %d\n", ageSum(users))

	//ageSile := mapToString(users, func(user map[string]string) string {
	//	return user["age"]
	//})
	//
	//sum := fieldSum(ageSile, func(age string) int {
	//	intAge, _ := strconv.Atoi(age)
	//	return intAge
	//})
	//
	//fmt.Printf("用户年龄累加结果: %d\n", sum)
	//
	//t := T{}
	//type1 := reflect.TypeOf(t).Kind()
	//type2 := reflect.ValueOf(users).Type()
	//
	//fmt.Println(type1, type2)

	//nums := []int{1, 2, 3, 4, 5}
	//// 初始化容器，元素类型和 nums 中的元素类型相同
	//c := NewContainer(reflect.TypeOf(nums[0]), 16)
	//// 添加元素到容器
	//for _, n := range nums {
	//	if err := c.Put(n); err != nil {
	//		panic(err)
	//	}
	//}
	//// 从容器读取元素，将返回结果初始化为 0
	//num := 0
	//if err := c.Get(&num); err != nil {
	//	panic(err)
	//}
	//// 打印返回结果值
	//fmt.Printf("%v (%T)\n", num, num)
	//divide()
	//fmt.Println("divide 方法调用完毕，回到 main 函数")
	//tcp.Client()
	//tcp.Server()
	//udp.Server()
	//udp.Client()
	//http.Client()
	//http.Server()
	//webSocket.Server()
	//goroutine.MainGoroutine()
	//goroutine.Atomic()
	//mysql.Commit()
	//redis.Set()
	// 1.抽取的爬邮箱
	//reptile.GetEmail()
	// 2.抽取的爬邮箱
	//reptile.GetEmail2("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	// 3.爬链接
	//reptile.GetLink("http://www.baidu.com/s?wd=%E8%B4%B4%E5%90%A7%20%E7%95%99%E4%B8%8B%E9%82%AE%E7%AE%B1&rsv_spt=1&rsv_iqid=0x98ace53400003985&issp=1&f=8&rsv_bp=1&rsv_idx=2&ie=utf-8&tn=baiduhome_pg&rsv_enter=1&rsv_dl=ib&rsv_sug2=0&inputT=5197&rsv_sug4=6345")
	// 4.爬手机号
	//reptile.GetPhone("https://www.zhaohaowang.com/")
	// 5.爬身份证号
	//reptile.GetIdCard("https://henan.qq.com/a/20171107/069413.htm")
	// 6.爬图片
	//reptile.GetImg("http://image.baidu.com/search/index?tn=baiduimage&ps=1&ct=201326592&lm=-1&cl=2&nc=1&ie=utf-8&word=%E7%BE%8E%E5%A5%B3")
	reptile.GetPicture()
}
