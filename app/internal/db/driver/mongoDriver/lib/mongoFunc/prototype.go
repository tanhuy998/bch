package mongoFunc

import (
	"fmt"
	"regexp"
	"strings"
)

type (
	MongoDBJSFuncPrototype string
)

var (
	func_detect_regex = regexp.MustCompile(`^function[\t\s]*\((?P<parameters>.*)\)[\t\s]*{(?P<body>.*)}$`)
)

const (
	param_regex_match_group_index int = 1
	body_regex_match_group_index  int = 2
)

const (
	mongo_func_default_prototype_pattern MongoDBJSFuncPrototype = `function (%s) {%s}`
)

type (
	__prototype__ struct {
		__arguments__
		body         string
		retStatement MongoFuncStatement
	}
)

func (this *__prototype__) clone() __prototype__ {

	ret := *this // copy "this" into ret

	ret.param_list = nil

	copy(ret.param_list, this.param_list)

	return ret
}

func (this *__prototype__) OverridePrototype(
	proto MongoDBJSFuncPrototype,
) {

	if v, exists := justed_raw_proto[proto]; exists {

		*this = v.clone()
		return
	}

	switch match := func_detect_regex.FindStringSubmatch(string(proto)); {
	case len(match) == 0:
		panic("invalid function raw prototype intended to be overriden")
	default:
		justed_raw_proto[proto] = this.clone()
		//this.__proto__ = proto
		this.detectParameters(match[param_regex_match_group_index])
		this.body = match[body_regex_match_group_index]
		return
	}
}

func (this *__prototype__) OverrideBody(body string) {

	this.body = body
}

func (this __prototype__) String() string {

	var (
		joined_params string
	)

	switch {
	case len(this.param_list) > 0:
		joined_params = strings.Join(this.param_list, ",")
		fallthrough
	default:
		return fmt.Sprintf(string(mongo_func_default_prototype_pattern), joined_params, this.body)
	}
}

func (this *__prototype__) Return(
	expr MongoFuncStatement,
) {

	pattern := `return $s`

	switch expr {
	case "":
		this.retStatement = fmt.Sprintf(pattern, `""`)
		return
	default:
		this.retStatement = fmt.Sprintf(pattern, expr)
	}
}
