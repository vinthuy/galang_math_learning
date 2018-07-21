package offer_singleton

//题目：设计一个类，我们只能生成该类的一个实例
type Singleton struct {
	Name string
}

var SingletonObj *Singleton

func init() {
	SingletonObj = &Singleton{"test"}
}
