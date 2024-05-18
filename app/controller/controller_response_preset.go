package controller

import "github.com/kataras/iris/v12/mvc"

func Response() mvc.Response {

	return mvc.Response{}
}

func BadRequest(err error) mvc.Response {

	return mvc.Response{
		Err: err,
	}
}

func Redirect(link string) mvc.Response {

	return mvc.Response{
		Code: 301,
		Path: link,
	}
}

func NotFound() mvc.Response {

	return mvc.Response{
		Code: 404,
	}
}

func Created() mvc.Response {

	return mvc.Response{
		Code: 201,
	}
}

func Ok() mvc.Response {

	return mvc.Response{
		Code: 200,
	}
}
