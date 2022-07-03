package main

import (
	"context"
	server0 "courseselection/kitex_gen/Server"
	"courseselection/sqlcontroller"
	"fmt"
	"log"
)

// ServiceImpl implements the last service interface defined in the IDL.
type ServiceImpl struct{}

// Login implements the ServiceImpl interface.
func (s *ServiceImpl) Login(ctx context.Context, req *server0.LoginRequest) (resp *server0.LoginResponse, err error) {
	// TODO: Your code here...
	query := fmt.Sprintf("SELECT passwd FROM userinfo WHERE userid=%s", req.Password)
	queryResp, err := sqlcontroller.Db.Query(query)
	resp = server0.NewLoginResponse()
	if err != nil {
		resp.Message = "出错了"
		log.Println(err.Error())
		return
	} else {
		var passwd string
		exist := false
		for queryResp.Next() {
			exist = true
			queryResp.Scan(&passwd)
		}
		if !exist {
			resp.Message = "用户不存在"
		} else if passwd != req.Password {
			resp.Message = "密码错误"
		} else {
			resp.Message = "登录成功"
			query = fmt.Sprintf("SELECT adminid FROM admininfo WHERE adminid=%s", req.Username)
			queryResp, err = sqlcontroller.Db.Query(query)
			exist = false
			for queryResp.Next() {
				exist = true
			}
			if exist {
				resp.Authority = new(int32)
				*resp.Authority = 1
				return
			}
			query = fmt.Sprintf("SELECT teacherid FROM teacherinfo WHERE teacherid=%s", req.Username)
			queryResp, err = sqlcontroller.Db.Query(query)
			exist = false
			for queryResp.Next() {
				exist = true
			}
			if exist {
				resp.Authority = new(int32)
				*resp.Authority = 2
				return
			}
			query = fmt.Sprintf("SELECT studentid FROM studentinfo WHERE studentid=%s", req.Username)
			queryResp, err = sqlcontroller.Db.Query(query)
			exist = false
			for queryResp.Next() {
				exist = true
			}
			if exist {
				resp.Authority = new(int32)
				*resp.Authority = 3
				return
			}
		}
	}
	return
}

// ChangePassword implements the ServiceImpl interface.
func (s *ServiceImpl) ChangePassword(ctx context.Context, req *server0.ChangePasswordRequenst) (resp *server0.ChangePasswordReponse, err error) {
	// TODO: Your code here...
	return
}

// ShowCourse implements the ServiceImpl interface.
func (s *ServiceImpl) ShowCourse(ctx context.Context, req *server0.StudentShowCourseRequest) (resp *server0.StudentShowCourseReponse, err error) {
	// TODO: Your code here...
	return
}

// SelectCourse implements the ServiceImpl interface.
func (s *ServiceImpl) SelectCourse(ctx context.Context, req *server0.SelectCourseRequest) (resp *server0.SelectCourseResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryScore implements the ServiceImpl interface.
func (s *ServiceImpl) QueryScore(ctx context.Context, req *server0.StudentQueryScoreRequest) (resp *server0.StudentQueryScoreResponse, err error) {
	// TODO: Your code here...
	return
}

// QuerySelection implements the ServiceImpl interface.
func (s *ServiceImpl) QuerySelection(ctx context.Context, req *server0.StudentQuerySelectionRequest) (resp *server0.StudentQuerySelectionResponse, err error) {
	// TODO: Your code here...
	return
}

// EvaluateRequest implements the ServiceImpl interface.
func (s *ServiceImpl) EvaluateRequest(ctx context.Context, req *server0.StudentEvaluateRequest) (resp *server0.StudentEvaluateResponse, err error) {
	// TODO: Your code here...
	return
}

// ShowCourseToTeacher implements the ServiceImpl interface.
func (s *ServiceImpl) ShowCourseToTeacher(ctx context.Context, req *server0.TeacherQueryCourseRequest) (resp *server0.TeacherQueryCourseResponse, err error) {
	// TODO: Your code here...
	return
}

// ShowStudentInfo implements the ServiceImpl interface.
func (s *ServiceImpl) ShowStudentInfo(ctx context.Context, req *server0.ShowStudentInfoRequest) (resp *server0.ShowStudentInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// InputScore implements the ServiceImpl interface.
func (s *ServiceImpl) InputScore(ctx context.Context, req *server0.TeacherInputScoreRequest) (resp *server0.TeacherInputScoreResponse, err error) {
	// TODO: Your code here...
	return
}

// ModifyShowCourse implements the ServiceImpl interface.
func (s *ServiceImpl) ModifyShowCourse(ctx context.Context, req *server0.TeacherQueryCourseRequest) (resp *server0.TeacherQueryCourseResponse, err error) {
	// TODO: Your code here...
	return
}

// ModifyShowStudent implements the ServiceImpl interface.
func (s *ServiceImpl) ModifyShowStudent(ctx context.Context, req *server0.ShowStudentInfoRequest) (resp *server0.ShowStudentInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// ModifyScore implements the ServiceImpl interface.
func (s *ServiceImpl) ModifyScore(ctx context.Context, req *server0.TeacherModifyScoreRequest) (resp *server0.TeacherModifyScoreResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryCourse implements the ServiceImpl interface.
func (s *ServiceImpl) QueryCourse(ctx context.Context, req *server0.TeacherQueryCourseRequest) (resp *server0.TeacherQueryCourseResponse, err error) {
	// TODO: Your code here...
	return
}

// ShowStudentScore implements the ServiceImpl interface.
func (s *ServiceImpl) ShowStudentScore(ctx context.Context, req *server0.ShowStudentInfoRequest) (resp *server0.ShowStudentInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// ShowCourseSelection implements the ServiceImpl interface.
func (s *ServiceImpl) ShowCourseSelection(ctx context.Context, req *server0.TeacherQueryCourseRequest) (resp *server0.TeacherQueryCourseResponse, err error) {
	// TODO: Your code here...
	return
}

// StudentCourseSelection implements the ServiceImpl interface.
func (s *ServiceImpl) StudentCourseSelection(ctx context.Context, req *server0.ShowStudentInfoRequest) (resp *server0.ShowStudentInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryStudentInfo implements the ServiceImpl interface.
func (s *ServiceImpl) QueryStudentInfo(ctx context.Context) (resp *server0.AdminQueryStudentInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryTeacherInfo implements the ServiceImpl interface.
func (s *ServiceImpl) QueryTeacherInfo(ctx context.Context) (resp *server0.AdminQueryTeacherInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryCourseInfo implements the ServiceImpl interface.
func (s *ServiceImpl) QueryCourseInfo(ctx context.Context) (resp *server0.AdminQueryCourseInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// AddStudent implements the ServiceImpl interface.
func (s *ServiceImpl) AddStudent(ctx context.Context, req *server0.AdminAddStudentInfoRequest) (resp *server0.AdminAddStudentInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteStudent implements the ServiceImpl interface.
func (s *ServiceImpl) DeleteStudent(ctx context.Context, req *server0.AdminDeleteStudentInfoRequest) (resp *server0.AdminDeleteStudentInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// AddTeacher implements the ServiceImpl interface.
func (s *ServiceImpl) AddTeacher(ctx context.Context, req *server0.AdminAddTeacherInfoRequest) (resp *server0.AdminAddTeacherInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteTeacher implements the ServiceImpl interface.
func (s *ServiceImpl) DeleteTeacher(ctx context.Context, req *server0.AdminDeleteTeacherInfoRequest) (resp *server0.AdminDeleteStudentInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// AddCourse implements the ServiceImpl interface.
func (s *ServiceImpl) AddCourse(ctx context.Context, req *server0.AdminAddCourseInfoRequest) (resp *server0.AdminAddCourseInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteCourse implements the ServiceImpl interface.
func (s *ServiceImpl) DeleteCourse(ctx context.Context, req *server0.AdminDeleteCourseInfoRequest) (resp *server0.AdminDeleteCourseInfoResponse, err error) {
	// TODO: Your code here...
	return
}
