package log

import logPool "app/internal/internal/log"

type (
	LogBroker struct {
	}
)

func (this LogBroker) Print(p0 ...interface{}) {

	logPool.Range(_print(p0))
}

func (this LogBroker) Printf(p0 string, p1 ...interface{}) {

	logPool.Range(_printf(p0, p1))
}

func (this LogBroker) Println(p0 ...interface{}) {

	logPool.Range(_println(p0))
}

func (this LogBroker) Fatal(p0 ...interface{}) {

	logPool.Range(_fatal(p0))
}

func (this LogBroker) Fatalf(p0 string, p1 ...interface{}) {

	logPool.Range(_fatalf(p0, p1))
}

func (this LogBroker) Fatalln(p0 ...interface{}) {

	logPool.Range(_fatalln(p0))
}

func (this LogBroker) Panic(p0 ...interface{}) {

	logPool.Range(_panic(p0))
}

func (this LogBroker) Panicf(p0 string, p1 ...interface{}) {

	logPool.Range(_panicf(p0, p1))
}

func (this LogBroker) Panicln(p0 ...interface{}) {

	logPool.Range(_panicln(p0))
}
