// Code generated by Kitex v0.3.2. DO NOT EDIT.

package studentservice

import (
	"context"
	server0 "courseselection/kitex_gen/Server"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return studentServiceServiceInfo
}

var studentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "StudentService"
	handlerType := (*server0.StudentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"ShowCourse":      kitex.NewMethodInfo(showCourseHandler, newStudentServiceShowCourseArgs, newStudentServiceShowCourseResult, false),
		"SelectCourse":    kitex.NewMethodInfo(selectCourseHandler, newStudentServiceSelectCourseArgs, newStudentServiceSelectCourseResult, false),
		"QueryScore":      kitex.NewMethodInfo(queryScoreHandler, newStudentServiceQueryScoreArgs, newStudentServiceQueryScoreResult, false),
		"QuerySelection":  kitex.NewMethodInfo(querySelectionHandler, newStudentServiceQuerySelectionArgs, newStudentServiceQuerySelectionResult, false),
		"EvaluateRequest": kitex.NewMethodInfo(evaluateRequestHandler, newStudentServiceEvaluateRequestArgs, newStudentServiceEvaluateRequestResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "server",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.3.2",
		Extra:           extra,
	}
	return svcInfo
}

func showCourseHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*server0.StudentServiceShowCourseArgs)
	realResult := result.(*server0.StudentServiceShowCourseResult)
	success, err := handler.(server0.StudentService).ShowCourse(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newStudentServiceShowCourseArgs() interface{} {
	return server0.NewStudentServiceShowCourseArgs()
}

func newStudentServiceShowCourseResult() interface{} {
	return server0.NewStudentServiceShowCourseResult()
}

func selectCourseHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*server0.StudentServiceSelectCourseArgs)
	realResult := result.(*server0.StudentServiceSelectCourseResult)
	success, err := handler.(server0.StudentService).SelectCourse(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newStudentServiceSelectCourseArgs() interface{} {
	return server0.NewStudentServiceSelectCourseArgs()
}

func newStudentServiceSelectCourseResult() interface{} {
	return server0.NewStudentServiceSelectCourseResult()
}

func queryScoreHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*server0.StudentServiceQueryScoreArgs)
	realResult := result.(*server0.StudentServiceQueryScoreResult)
	success, err := handler.(server0.StudentService).QueryScore(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newStudentServiceQueryScoreArgs() interface{} {
	return server0.NewStudentServiceQueryScoreArgs()
}

func newStudentServiceQueryScoreResult() interface{} {
	return server0.NewStudentServiceQueryScoreResult()
}

func querySelectionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*server0.StudentServiceQuerySelectionArgs)
	realResult := result.(*server0.StudentServiceQuerySelectionResult)
	success, err := handler.(server0.StudentService).QuerySelection(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newStudentServiceQuerySelectionArgs() interface{} {
	return server0.NewStudentServiceQuerySelectionArgs()
}

func newStudentServiceQuerySelectionResult() interface{} {
	return server0.NewStudentServiceQuerySelectionResult()
}

func evaluateRequestHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*server0.StudentServiceEvaluateRequestArgs)
	realResult := result.(*server0.StudentServiceEvaluateRequestResult)
	success, err := handler.(server0.StudentService).EvaluateRequest(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newStudentServiceEvaluateRequestArgs() interface{} {
	return server0.NewStudentServiceEvaluateRequestArgs()
}

func newStudentServiceEvaluateRequestResult() interface{} {
	return server0.NewStudentServiceEvaluateRequestResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ShowCourse(ctx context.Context, req *server0.StudentShowCourseRequest) (r *server0.StudentShowCourseReponse, err error) {
	var _args server0.StudentServiceShowCourseArgs
	_args.Req = req
	var _result server0.StudentServiceShowCourseResult
	if err = p.c.Call(ctx, "ShowCourse", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SelectCourse(ctx context.Context, req *server0.SelectCourseRequest) (r *server0.SelectCourseResponse, err error) {
	var _args server0.StudentServiceSelectCourseArgs
	_args.Req = req
	var _result server0.StudentServiceSelectCourseResult
	if err = p.c.Call(ctx, "SelectCourse", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryScore(ctx context.Context, req *server0.StudentQueryScoreRequest) (r *server0.StudentQueryScoreResponse, err error) {
	var _args server0.StudentServiceQueryScoreArgs
	_args.Req = req
	var _result server0.StudentServiceQueryScoreResult
	if err = p.c.Call(ctx, "QueryScore", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QuerySelection(ctx context.Context, req *server0.StudentQuerySelectionRequest) (r *server0.StudentQuerySelectionResponse, err error) {
	var _args server0.StudentServiceQuerySelectionArgs
	_args.Req = req
	var _result server0.StudentServiceQuerySelectionResult
	if err = p.c.Call(ctx, "QuerySelection", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EvaluateRequest(ctx context.Context, req *server0.StudentEvaluateRequest) (r *server0.StudentEvaluateResponse, err error) {
	var _args server0.StudentServiceEvaluateRequestArgs
	_args.Req = req
	var _result server0.StudentServiceEvaluateRequestResult
	if err = p.c.Call(ctx, "EvaluateRequest", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
