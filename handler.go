package main

import (
	"context"
	server0 "courseselection/kitex_gen/Server"
)

// AdminServiceImpl implements the last service interface defined in the IDL.
type AdminServiceImpl struct{}

// QueryStudentInfo implements the AdminServiceImpl interface.
func (s *AdminServiceImpl) QueryStudentInfo(ctx context.Context) (resp *server0.AdminQueryStudentInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryTeacherInfo implements the AdminServiceImpl interface.
func (s *AdminServiceImpl) QueryTeacherInfo(ctx context.Context) (resp *server0.AdminQueryTeacherInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryCourseInfo implements the AdminServiceImpl interface.
func (s *AdminServiceImpl) QueryCourseInfo(ctx context.Context) (resp *server0.AdminQueryCourseInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// AddStudent implements the AdminServiceImpl interface.
func (s *AdminServiceImpl) AddStudent(ctx context.Context, req *server0.AdminAddStudentInfoRequest) (resp *server0.AdminAddStudentInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteStudent implements the AdminServiceImpl interface.
func (s *AdminServiceImpl) DeleteStudent(ctx context.Context, req *server0.AdminDeleteStudentInfoRequest) (resp *server0.AdminDeleteStudentInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// AddTeacher implements the AdminServiceImpl interface.
func (s *AdminServiceImpl) AddTeacher(ctx context.Context, req *server0.AdminAddTeacherInfoRequest) (resp *server0.AdminAddTeacherInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteTeacher implements the AdminServiceImpl interface.
func (s *AdminServiceImpl) DeleteTeacher(ctx context.Context, req *server0.AdminDeleteTeacherInfoRequest) (resp *server0.AdminDeleteStudentInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// AddCourse implements the AdminServiceImpl interface.
func (s *AdminServiceImpl) AddCourse(ctx context.Context, req *server0.AdminAddCourseInfoRequest) (resp *server0.AdminAddCourseInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteCourse implements the AdminServiceImpl interface.
func (s *AdminServiceImpl) DeleteCourse(ctx context.Context, req *server0.AdminDeleteCourseInfoRequest) (resp *server0.AdminDeleteCourseInfoResponse, err error) {
	// TODO: Your code here...
	return
}
