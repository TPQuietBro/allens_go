package log

import (
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	// 打开文件系统, 设置写入模式, 0600 代表可以在多种操作系统运行
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

func Run(destination string) {
	// 定义自己的log, 其中fileLog实现了Write()方法默认就是io.Writer类型
	log = stlog.New(fileLog(destination), "[go]", stlog.LstdFlags)
}

// 日志服务handler
func RegisterHandlers() {
	http.HandleFunc("/log", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		// 针对post请求写入日志
		case http.MethodPost:
			msg, err := ioutil.ReadAll(request.Body)
			if err != nil || len(msg) == 0 {
				writer.WriteHeader(http.StatusBadRequest)
			}
			write(string(msg))
		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

// 调用自己的log 写入日志信息
func write(message string) {
	log.Printf("%v\n", message)
}
