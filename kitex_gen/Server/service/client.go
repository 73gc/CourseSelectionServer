// Code generated by Kitex v0.3.2. DO NOT EDIT.

package service

import (
	"context"
	server0 "courseselection/kitex_gen/Server"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Login(ctx context.Context, req *server0.LoginRequest, callOptions ...callopt.Option) (r *server0.LoginResponse, err error)
	ChangePassword(ctx context.Context, req *server0.ChangePasswordRequenst, callOptions ...callopt.Option) (r *server0.ChangePasswordReponse, err error)
	ShowCourse(ctx context.Context, req *server0.StudentShowCourseRequest, callOptions ...callopt.Option) (r *server0.StudentShowCourseReponse, err error)
	SelectCourse(ctx context.Context, req *server0.SelectCourseRequest, callOptions ...callopt.Option) (r *server0.SelectCourseResponse, err error)
	QueryScore(ctx context.Context, req *server0.StudentQueryScoreRequest, callOptions ...callopt.Option) (r *server0.StudentQueryScoreResponse, err error)
	QuerySelection(ctx context.Context, req *server0.StudentQuerySelectionRequest, callOptions ...callopt.Option) (r *server0.StudentQuerySelectionResponse, err error)
	EvaluateRequest(ctx context.Context, req *server0.StudentEvaluateRequest, callOptions ...callopt.Option) (r *server0.StudentEvaluateResponse, err error)
	ShowCourseToTeacher(ctx context.Context, req *server0.TeacherQueryCourseRequest, callOptions ...callopt.Option) (r *server0.TeacherQueryCourseResponse, err error)
	ShowStudentInfo(ctx context.Context, req *server0.ShowStudentInfoRequest, callOptions ...callopt.Option) (r *server0.ShowStudentInfoResponse, err error)
	InputScore(ctx context.Context, req *server0.TeacherInputScoreRequest, callOptions ...callopt.Option) (r *server0.TeacherInputScoreResponse, err error)
	ModifyShowCourse(ctx context.Context, req *server0.TeacherQueryCourseRequest, callOptions ...callopt.Option) (r *server0.TeacherQueryCourseResponse, err error)
	ModifyShowStudent(ctx context.Context, req *server0.ShowStudentInfoRequest, callOptions ...callopt.Option) (r *server0.ShowStudentInfoResponse, err error)
	ModifyScore(ctx context.Context, req *server0.TeacherModifyScoreRequest, callOptions ...callopt.Option) (r *server0.TeacherModifyScoreResponse, err error)
	QueryCourse(ctx context.Context, req *server0.TeacherQueryCourseRequest, callOptions ...callopt.Option) (r *server0.TeacherQueryCourseResponse, err error)
	ShowStudentScore(ctx context.Context, req *server0.ShowStudentInfoRequest, callOptions ...callopt.Option) (r *server0.ShowStudentInfoResponse, err error)
	ShowCourseSelection(ctx context.Context, req *server0.TeacherQueryCourseRequest, callOptions ...callopt.Option) (r *server0.TeacherQueryCourseResponse, err error)
	StudentCourseSelection(ctx context.Context, req *server0.ShowStudentInfoRequest, callOptions ...callopt.Option) (r *server0.ShowStudentInfoResponse, err error)
	QueryStudentInfo(ctx context.Context, callOptions ...callopt.Option) (r *server0.AdminQueryStudentInfoResponse, err error)
	QueryTeacherInfo(ctx context.Context, callOptions ...callopt.Option) (r *server0.AdminQueryTeacherInfoResponse, err error)
	QueryCourseInfo(ctx context.Context, callOptions ...callopt.Option) (r *server0.AdminQueryCourseInfoResponse, err error)
	AddStudent(ctx context.Context, req *server0.AdminAddStudentInfoRequest, callOptions ...callopt.Option) (r *server0.AdminAddStudentInfoResponse, err error)
	DeleteStudent(ctx context.Context, req *server0.AdminDeleteStudentInfoRequest, callOptions ...callopt.Option) (r *server0.AdminDeleteStudentInfoResponse, err error)
	AddTeacher(ctx context.Context, req *server0.AdminAddTeacherInfoRequest, callOptions ...callopt.Option) (r *server0.AdminAddTeacherInfoResponse, err error)
	DeleteTeacher(ctx context.Context, req *server0.AdminDeleteTeacherInfoRequest, callOptions ...callopt.Option) (r *server0.AdminDeleteStudentInfoResponse, err error)
	AddCourse(ctx context.Context, req *server0.AdminAddCourseInfoRequest, callOptions ...callopt.Option) (r *server0.AdminAddCourseInfoResponse, err error)
	DeleteCourse(ctx context.Context, req *server0.AdminDeleteCourseInfoRequest, callOptions ...callopt.Option) (r *server0.AdminDeleteCourseInfoResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kServiceClient struct {
	*kClient
}

func (p *kServiceClient) Login(ctx context.Context, req *server0.LoginRequest, callOptions ...callopt.Option) (r *server0.LoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, req)
}

func (p *kServiceClient) ChangePassword(ctx context.Context, req *server0.ChangePasswordRequenst, callOptions ...callopt.Option) (r *server0.ChangePasswordReponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChangePassword(ctx, req)
}

func (p *kServiceClient) ShowCourse(ctx context.Context, req *server0.StudentShowCourseRequest, callOptions ...callopt.Option) (r *server0.StudentShowCourseReponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ShowCourse(ctx, req)
}

func (p *kServiceClient) SelectCourse(ctx context.Context, req *server0.SelectCourseRequest, callOptions ...callopt.Option) (r *server0.SelectCourseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SelectCourse(ctx, req)
}

func (p *kServiceClient) QueryScore(ctx context.Context, req *server0.StudentQueryScoreRequest, callOptions ...callopt.Option) (r *server0.StudentQueryScoreResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryScore(ctx, req)
}

func (p *kServiceClient) QuerySelection(ctx context.Context, req *server0.StudentQuerySelectionRequest, callOptions ...callopt.Option) (r *server0.StudentQuerySelectionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QuerySelection(ctx, req)
}

func (p *kServiceClient) EvaluateRequest(ctx context.Context, req *server0.StudentEvaluateRequest, callOptions ...callopt.Option) (r *server0.StudentEvaluateResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EvaluateRequest(ctx, req)
}

func (p *kServiceClient) ShowCourseToTeacher(ctx context.Context, req *server0.TeacherQueryCourseRequest, callOptions ...callopt.Option) (r *server0.TeacherQueryCourseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ShowCourseToTeacher(ctx, req)
}

func (p *kServiceClient) ShowStudentInfo(ctx context.Context, req *server0.ShowStudentInfoRequest, callOptions ...callopt.Option) (r *server0.ShowStudentInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ShowStudentInfo(ctx, req)
}

func (p *kServiceClient) InputScore(ctx context.Context, req *server0.TeacherInputScoreRequest, callOptions ...callopt.Option) (r *server0.TeacherInputScoreResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.InputScore(ctx, req)
}

func (p *kServiceClient) ModifyShowCourse(ctx context.Context, req *server0.TeacherQueryCourseRequest, callOptions ...callopt.Option) (r *server0.TeacherQueryCourseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ModifyShowCourse(ctx, req)
}

func (p *kServiceClient) ModifyShowStudent(ctx context.Context, req *server0.ShowStudentInfoRequest, callOptions ...callopt.Option) (r *server0.ShowStudentInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ModifyShowStudent(ctx, req)
}

func (p *kServiceClient) ModifyScore(ctx context.Context, req *server0.TeacherModifyScoreRequest, callOptions ...callopt.Option) (r *server0.TeacherModifyScoreResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ModifyScore(ctx, req)
}

func (p *kServiceClient) QueryCourse(ctx context.Context, req *server0.TeacherQueryCourseRequest, callOptions ...callopt.Option) (r *server0.TeacherQueryCourseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryCourse(ctx, req)
}

func (p *kServiceClient) ShowStudentScore(ctx context.Context, req *server0.ShowStudentInfoRequest, callOptions ...callopt.Option) (r *server0.ShowStudentInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ShowStudentScore(ctx, req)
}

func (p *kServiceClient) ShowCourseSelection(ctx context.Context, req *server0.TeacherQueryCourseRequest, callOptions ...callopt.Option) (r *server0.TeacherQueryCourseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ShowCourseSelection(ctx, req)
}

func (p *kServiceClient) StudentCourseSelection(ctx context.Context, req *server0.ShowStudentInfoRequest, callOptions ...callopt.Option) (r *server0.ShowStudentInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StudentCourseSelection(ctx, req)
}

func (p *kServiceClient) QueryStudentInfo(ctx context.Context, callOptions ...callopt.Option) (r *server0.AdminQueryStudentInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryStudentInfo(ctx)
}

func (p *kServiceClient) QueryTeacherInfo(ctx context.Context, callOptions ...callopt.Option) (r *server0.AdminQueryTeacherInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryTeacherInfo(ctx)
}

func (p *kServiceClient) QueryCourseInfo(ctx context.Context, callOptions ...callopt.Option) (r *server0.AdminQueryCourseInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryCourseInfo(ctx)
}

func (p *kServiceClient) AddStudent(ctx context.Context, req *server0.AdminAddStudentInfoRequest, callOptions ...callopt.Option) (r *server0.AdminAddStudentInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddStudent(ctx, req)
}

func (p *kServiceClient) DeleteStudent(ctx context.Context, req *server0.AdminDeleteStudentInfoRequest, callOptions ...callopt.Option) (r *server0.AdminDeleteStudentInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteStudent(ctx, req)
}

func (p *kServiceClient) AddTeacher(ctx context.Context, req *server0.AdminAddTeacherInfoRequest, callOptions ...callopt.Option) (r *server0.AdminAddTeacherInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddTeacher(ctx, req)
}

func (p *kServiceClient) DeleteTeacher(ctx context.Context, req *server0.AdminDeleteTeacherInfoRequest, callOptions ...callopt.Option) (r *server0.AdminDeleteStudentInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteTeacher(ctx, req)
}

func (p *kServiceClient) AddCourse(ctx context.Context, req *server0.AdminAddCourseInfoRequest, callOptions ...callopt.Option) (r *server0.AdminAddCourseInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddCourse(ctx, req)
}

func (p *kServiceClient) DeleteCourse(ctx context.Context, req *server0.AdminDeleteCourseInfoRequest, callOptions ...callopt.Option) (r *server0.AdminDeleteCourseInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteCourse(ctx, req)
}
