package log

import (
	logManager "app/internal/internal/log"
	"log"
)

func _print(p0 []interface{}) logManager.LogPoolRangeFunc {

	return func(logger *log.Logger) {

		logger.Print(p0...)
	}
}

func _printf(p0 string, p1 []interface{}) logManager.LogPoolRangeFunc {

	return func(logger *log.Logger) {

		logger.Printf(p0, p1...)
	}
}

func _println(p0 []interface{}) logManager.LogPoolRangeFunc {

	return func(logger *log.Logger) {

		logger.Println(p0...)
	}
}

func _fatal(p0 []interface{}) logManager.LogPoolRangeFunc {

	return func(logger *log.Logger) {

		logger.Fatal(p0...)
	}
}

func _fatalf(p0 string, p1 []interface{}) logManager.LogPoolRangeFunc {

	return func(logger *log.Logger) {

		logger.Fatalf(p0, p1...)
	}
}

func _fatalln(p0 []interface{}) logManager.LogPoolRangeFunc {

	return func(logger *log.Logger) {

		logger.Fatalln(p0...)
	}
}

func _panic(p0 []interface{}) logManager.LogPoolRangeFunc {

	return func(logger *log.Logger) {

		logger.Panic(p0...)
	}
}

func _panicf(p0 string, p1 []interface{}) logManager.LogPoolRangeFunc {

	return func(logger *log.Logger) {

		logger.Panicf(p0, p1...)
	}
}

func _panicln(p0 []interface{}) logManager.LogPoolRangeFunc {

	return func(logger *log.Logger) {

		logger.Panicln(p0...)
	}
}
