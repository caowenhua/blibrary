package controller

import "fmt"

func Menu(choice *int) int {
	fmt.Println("请选择功能：")
	fmt.Println("1.新建书籍")
	fmt.Println("2.查询书籍(按ISDN)")
	fmt.Println("3.查询书籍(按文件名)")
	fmt.Println("4.查询书籍(按价格)")
	fmt.Println("5.查询书籍(按日期)")
	fmt.Println("6.查询书籍(按作者)")
	fmt.Println("7.删除书籍(按ISDN)")

	fmt.Scanf("%d", &(*choice))
	//	fmt.Println(choice)
	return *choice
}

func Function(choice int) {
	switch choice {
	case 1:
		fmt.Println("function1.新建书籍")
		newBook()
	case 2:
		fmt.Println("function2.查询书籍(按ISDN)")
		findBookById()
	case 3:
		fmt.Println("function3.查询书籍(按文件名)")
		findBookByName()
	case 4:
		fmt.Println("function4.查询书籍(按价格)")
		findBookByPrice()
	case 5:
		fmt.Println("function5.查询书籍(按日期)")
		findBookByTime()
	case 6:
		fmt.Println("function6.查询书籍(按作者)")
		findBookByAuthor()
	case 7:
		fmt.Println("function7.删除书籍(按ISDN)")
		deleteBookById()
	default:
		fmt.Println("您好，请输入上述选项数字，谢谢！")
	}
}

func newBook() {
	var name, isbn, author string
	var price float64
	var time int64
	fmt.Println("请输入书籍名称")
	fmt.Scan(&name)
	fmt.Println("请输入书籍ISBN号")
	fmt.Scan(&isbn)
	fmt.Println("请输入书籍作者")
	fmt.Scan(&author)
	fmt.Println("请输入价格")
	fmt.Scan(&price)
	fmt.Println("请输入出版时间")
	fmt.Scan(&time)
}

func findBookById() {
	var isbn string
	fmt.Println("请输入书籍ISBN号")
	fmt.Scan(&isbn)
}

func findBookByName() {
	var name string
	fmt.Println("请输入书籍名称")
	fmt.Scan(&name)
}
func findBookByAuthor() {
	var author string
	fmt.Println("请输入书籍作者")
	fmt.Scan(&author)
}
func findBookByPrice() {
	var price float64
	fmt.Println("请输入书籍价格")
	fmt.Scan(&price)
}
func findBookByTime() {
	var time int64
	fmt.Println("请输入书籍出版时间")
	fmt.Scan(&time)
}
func deleteBookById() {
	var isbn string
	fmt.Println("请输入书籍ISBN号")
	fmt.Scan(&isbn)
}
