package gocaptcha

//testing包提供go包的自动化测试支持，通过go test命令(源码文件与测试文件要关联一起运行：go test Xxx.go  Xxx_test.go)
// 单元测试文件要以_test.go结尾
import "testing"

//每个测试用例可能并发执行，使用 testing.T 提供的日志输出可以保证日志跟随这个测试上下文一起打印输出
func TestRandom(t *testing.T){
	for i :=0; i < 100; i++{
		t.Log(Random(0, 1))
	}
}