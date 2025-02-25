package bank



var deposits = make(chan int) // 发送存款余额
var ballances = make(chan int) // 接收余额


func Deposits(amount int) {
	deposits <- amount
}

func Balances() int{
	return <- ballances
}


func teller(){
	var ballance int // ballance被限制在teller goroutine中
	for {
		select{
		case amount := <- deposits:
			ballance += amount
		case ballances <- ballance:
		}
	}
}


func init(){
	go teller() // 启动监控goroutine
}