package loggerServicePort

type (
	ErrorLogger interface {
		Error(message string)
	}
)

type (
	ILogger interface {
		Print(msgs ...interface{})
		Printf(pattern string, vars ...interface{})
		Println(msgs ...interface{})
		Fatal(msgs ...interface{})
		Fatalf(pattern string, vars ...interface{})
		Fatalln(msgs ...interface{})
		Panic(msgs ...interface{})
		Panicf(pattern string, vars ...interface{})
		Panicln(vars ...interface{})
	}
)
