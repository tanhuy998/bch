package mongoFunc

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	func_param_explosion_delemiter_regex = regexp.MustCompile(`[\t\s]*,[\t\s]*`)
)

type (
	__parameters__ struct {
		//__proto__  string
		param_list []string
	}
)

func (this *__parameters__) init() {

}

func (this *__parameters__) detectParameters(input string) {

	switch input {
	case "":
		this.param_list = nil
		return
	default:
		this.param_list = func_param_explosion_delemiter_regex.Split(input, -1)
	}

	this.validateParams()
}

func (this __parameters__) validateParams() {

	if len(this.param_list) == 0 {

		return
	}

	for _, name := range this.param_list {

		if name == "" {

			panic(
				fmt.Sprintf(`invalid definition for function's parameters, detected "%s"`, strings.Join(this.param_list, ",")),
			)
		}
	}
}

func (this *__parameters__) OverrideParams(params ...string) {

	this.param_list = make([]string, len(params))

	for i, name := range params {

		this.__setParam(i, name)
	}
}

func (this *__parameters__) AddParam(name string) {

	this.param_list = append(this.param_list, "")

	this.__setParam(len(this.param_list), name)
}

func (this *__parameters__) __setParam(index int, name string) {

	if name == "" {
		panic("parameter's name must not be empty")
	}

	this.param_list[index] = name
}

func (this *__parameters__) GetParam(index int) string {

	switch len(this.param_list) {
	case 0:
		return ""
	default:
		return this.param_list[index]
	}
}
