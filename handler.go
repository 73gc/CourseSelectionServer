package main

import (
	"context"
	server0 "courseselection/kitex_gen/Server"
	"courseselection/sqlcontroller"
	"errors"
	"fmt"
	"log"
)

// ServiceImpl implements the last service interface defined in the IDL.
type ServiceImpl struct{}

// Login implements the ServiceImpl interface.
func (s *ServiceImpl) Login(ctx context.Context, req *server0.LoginRequest) (resp *server0.LoginResponse, err error) {
	// TODO: Your code here...
	query := fmt.Sprintf("SELECT passwd FROM userinfo WHERE userid=%s", req.Password)
	queryResp, err_ := sqlcontroller.Db.Query(query)
	resp = server0.NewLoginResponse()
	if err_ != nil {
		resp.Message = "出错了"
		log.Println(err_.Error())
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
			err = nil
			resp.Message = "登录成功"
			query = fmt.Sprintf("SELECT adminid FROM admininfo WHERE adminid=%s", req.Username)
			queryResp, _ = sqlcontroller.Db.Query(query)
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
			queryResp, _ = sqlcontroller.Db.Query(query)
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
			queryResp, _ = sqlcontroller.Db.Query(query)
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
	query := fmt.Sprintf("UPDATE userinfo SET passwd='%s' WHERE userid=%s", req.NewPassword_, req.Username)
	sqlResp, err_ := sqlcontroller.Db.Exec(query)
	resp = server0.NewChangePasswordReponse()
	if err_ != nil {
		resp.Message = "发生错误，请稍后重试"
		return
	}
	n, _ := sqlResp.RowsAffected()
	if n != 0 {
		resp.Message = "修改成功"
	} else {
		resp.Message = "新密码和原密码相同"
	}
	err = nil
	return
}

// ShowCourse implements the ServiceImpl interface.
func (s *ServiceImpl) ShowCourse(ctx context.Context, req *server0.StudentShowCourseRequest) (resp *server0.StudentShowCourseReponse, err error) {
	// TODO: Your code here...
	query := fmt.Sprintf("SELECT courseinfo.courseid,courseinfo.coursename,teacherinfo.teachername,courseinfo.credit FROM courseinfo,teacherinfo WHERE courseinfo.teacherid = teacherinfo.teacherid AND courseinfo.courseid NOT IN (SELECT courseid FROM selectioninfo WHERE studentid='%s')", *req.StudentId)
	sqlResp, err_ := sqlcontroller.Db.Query(query)
	if err_ != nil {
		err = errors.New("出错了，稍后重试")
		return
	}
	resp = server0.NewStudentShowCourseReponse()
	resp.Courses = make([]*server0.ShowCourseResponse, 0)
	for sqlResp.Next() {
		var courseId, courseName, teacherName string
		var credit float64
		sqlResp.Scan(&courseId, &courseName, &teacherName, &credit)
		resp.Courses = append(resp.Courses, &server0.ShowCourseResponse{
			CourseId:    courseId,
			CourseName:  courseName,
			TeacherName: teacherName,
			Credit:      credit,
		})
	}
	err = nil
	return
}

// SelectCourse implements the ServiceImpl interface.
func (s *ServiceImpl) SelectCourse(ctx context.Context, req *server0.SelectCourseRequest) (resp *server0.SelectCourseResponse, err error) {
	// TODO: Your code here...
	query := fmt.Sprintf("INSERT INTO selectioninfo (studentid, courseid) VALUES ('%s', '%s')", req.StudentId, req.CourseId)
	_, err_ := sqlcontroller.Db.Exec(query)
	resp = server0.NewSelectCourseResponse()
	if err_ != nil {
		resp.Message = "出错了，稍后重试"
		return
	}
	resp.Message = "添加成功"
	err = nil
	return
}

// QueryScore implements the ServiceImpl interface.
func (s *ServiceImpl) QueryScore(ctx context.Context, req *server0.StudentQueryScoreRequest) (resp *server0.StudentQueryScoreResponse, err error) {
	// TODO: Your code here...
	query := fmt.Sprintf("SELECT courseinfo.courseid, courseinfo.coursename, courseinfo.credit, selectioninfo.score FROM courseinfo, selectioninfo WHERE courseinfo.courseid=selectioninfo.courseid AND selectioninfo.studentid='%s' AND selectioninfo.score IS NOT NULL", req.StudentId)
	sqlResp, err_ := sqlcontroller.Db.Query(query)
	if err_ != nil {
		err = errors.New("出错了，稍后重试")
		return
	}
	resp = server0.NewStudentQueryScoreResponse()
	resp.CourseScore = make([]*server0.StudentScoreResponse, 0)
	for sqlResp.Next() {
		var courseId, courseName string
		var credit, score float64
		sqlResp.Scan(&courseId, &courseName, &credit, &score)
		resp.CourseScore = append(resp.CourseScore, &server0.StudentScoreResponse{
			CourseId:   courseId,
			CourseName: courseName,
			Credit:     credit,
			Score:      score,
		})
	}
	err = nil
	return
}

// QuerySelection implements the ServiceImpl interface.
func (s *ServiceImpl) QuerySelection(ctx context.Context, req *server0.StudentQuerySelectionRequest) (resp *server0.StudentQuerySelectionResponse, err error) {
	// TODO: Your code here...
	// server0.ShowCourseResponse
	query := fmt.Sprintf("SELECT courseinfo.courseid, courseinfo.coursename, teacherinfo.teachername, courseinfo.credit FROM courseinfo, teacherinfo, selectioninfo WHERE courseinfo.teacherid=teacherinfo.teacherid AND selectioninfo.courseid=courseinfo.courseid AND selectioninfo.studentid='%s'", req.StudentId)
	sqlResp, err_ := sqlcontroller.Db.Query(query)
	if err_ != nil {
		err = errors.New("出错了，稍后重试")
		return
	}
	resp = server0.NewStudentQuerySelectionResponse()
	resp.Courses = make([]*server0.ShowCourseResponse, 0)
	for sqlResp.Next() {
		var courseId, courseName, teacherName string
		var credit float64
		sqlResp.Scan(&courseId, &courseName, &teacherName, &credit)
		resp.Courses = append(resp.Courses, &server0.ShowCourseResponse{
			CourseId:    courseId,
			CourseName:  courseName,
			TeacherName: teacherName,
			Credit:      credit,
		})
	}
	err = nil
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
	resp = server0.NewAdminQueryStudentInfoResponse()
	query := "SELECT * FROM studentinfo"
	sqlResp, err_ := sqlcontroller.Db.Query(query)
	if err_ != nil {
		err = errors.New("请求出错")
		return
	}
	resp.Students = make([]*server0.StudentInfo, 0)
	for sqlResp.Next() {
		var studentId, studentName, studentClass string
		sqlResp.Scan(&studentId, &studentName, &studentClass)
		resp.Students = append(resp.Students, &server0.StudentInfo{
			StudentId:     studentId,
			StudentName:   studentName,
			ClassAndGrade: studentClass,
		})
	}
	err = nil
	return
}

// QueryTeacherInfo implements the ServiceImpl interface.
func (s *ServiceImpl) QueryTeacherInfo(ctx context.Context) (resp *server0.AdminQueryTeacherInfoResponse, err error) {
	// TODO: Your code here...
	resp = server0.NewAdminQueryTeacherInfoResponse()
	query := "SELECT * FROM teacherinfo"
	sqlResp, err_ := sqlcontroller.Db.Query(query)
	if err_ != nil {
		err = errors.New("请求出错")
		return
	}
	resp.Teachers = make([]*server0.TeacherInfo, 0)
	for sqlResp.Next() {
		var teacherId, teacherName string
		sqlResp.Scan(&teacherId, &teacherName)
		resp.Teachers = append(resp.Teachers, &server0.TeacherInfo{
			TeacherId:   teacherId,
			TeacherName: teacherName,
		})
	}
	err = nil
	return
}

// QueryCourseInfo implements the ServiceImpl interface.
func (s *ServiceImpl) QueryCourseInfo(ctx context.Context) (resp *server0.AdminQueryCourseInfoResponse, err error) {
	// TODO: Your code here...
	resp = server0.NewAdminQueryCourseInfoResponse()
	query := "SELECT * FROM courseinfo"
	sqlResp, err_ := sqlcontroller.Db.Query(query)
	if err_ != nil {
		err = errors.New("请求出错")
		return
	}
	resp.Courses = make([]*server0.CourseInfo, 0)
	for sqlResp.Next() {
		var courseId, courseName, teacherId, credit string
		sqlResp.Scan(&courseId, &courseName, &teacherId, &credit)
		resp.Courses = append(resp.Courses, &server0.CourseInfo{
			CourseId:    courseId,
			CourseName:  courseName,
			TeacherName: teacherId,
			Credit:      credit,
		})
	}
	err = nil
	return
}

// AddStudent implements the ServiceImpl interface.
func (s *ServiceImpl) AddStudent(ctx context.Context, req *server0.AdminAddStudentInfoRequest) (resp *server0.AdminAddStudentInfoResponse, err error) {
	// TODO: Your code here...
	query := fmt.Sprintf("INSERT INTO studentinfo (studentid, studentname, studentclass) VALUES ('%s', '%s', '%s') ON DUPLICATE KEY UPDATE studentclass='%s'", req.StudentId, req.StudentName, req.ClassAndGrade, req.ClassAndGrade)
	_, err_ := sqlcontroller.Db.Exec(query)
	resp = server0.NewAdminAddStudentInfoResponse()
	if err_ != nil {
		resp.Message = "出错了，请稍后重试"
		return
	}
	query = fmt.Sprintf("INSERT INTO userinfo (userid, passwd) VALUES ('%s', '%s') ON DUPLICATE KEY UPDATE passwd='%s'", req.StudentId, req.StudentId, req.StudentId)
	_, err_ = sqlcontroller.Db.Exec(query)
	if err_ != nil {
		resp.Message = "出错了，请稍后重试"
		return
	}
	resp.Message = "添加成功"
	err = nil
	return
}

// DeleteStudent implements the ServiceImpl interface.
func (s *ServiceImpl) DeleteStudent(ctx context.Context, req *server0.AdminDeleteStudentInfoRequest) (resp *server0.AdminDeleteStudentInfoResponse, err error) {
	// TODO: Your code here...
	query := fmt.Sprintf("DELETE FROM studentinfo WHERE studentid='%s'", req.StudentId)
	_, err_ := sqlcontroller.Db.Exec(query)
	resp = server0.NewAdminDeleteStudentInfoResponse()
	if err_ != nil {
		resp.Message = "出错了，请稍后重试"
		return
	}
	query = fmt.Sprintf("DELETE FROM userinfo WHERE userid='%s'", req.StudentId)
	_, err_ = sqlcontroller.Db.Exec(query)
	if err_ != nil {
		resp.Message = "出错了，请稍后重试"
		return
	}
	resp.Message = "删除成功"
	err = nil
	return
}

// AddTeacher implements the ServiceImpl interface.
func (s *ServiceImpl) AddTeacher(ctx context.Context, req *server0.AdminAddTeacherInfoRequest) (resp *server0.AdminAddTeacherInfoResponse, err error) {
	// TODO: Your code here...
	query := fmt.Sprintf("INSERT INTO teacherinfo (teacherid, teachername) VALUES ('%s', '%s') ON DUPLICATE KEY UPDATE teachername='%s'", req.TeacherId, req.TeacherName, req.TeacherName)
	_, err_ := sqlcontroller.Db.Exec(query)
	resp = server0.NewAdminAddTeacherInfoResponse()
	if err_ != nil {
		resp.Message = "出错了，请稍后重试"
		return
	}
	query = fmt.Sprintf("INSERT INTO userinfo (userid, passwd) VALUES ('%s', '%s') ON DUPLICATE KEY UPDATE passwd='%s'", req.TeacherId, req.TeacherId, req.TeacherId)
	_, err_ = sqlcontroller.Db.Exec(query)
	if err_ != nil {
		resp.Message = "出错了，请稍后重试"
		return
	}
	resp.Message = "添加成功"
	err = nil
	return
}

// DeleteTeacher implements the ServiceImpl interface.
func (s *ServiceImpl) DeleteTeacher(ctx context.Context, req *server0.AdminDeleteTeacherInfoRequest) (resp *server0.AdminDeleteTeacherInfoResponse, err error) {
	// TODO: Your code here...
	query := fmt.Sprintf("DELETE FROM teacherinfo WHERE teacherid='%s'", req.TeacherId)
	_, err_ := sqlcontroller.Db.Exec(query)
	resp = server0.NewAdminDeleteTeacherInfoResponse()
	if err_ != nil {
		resp.Message = "出错了，请稍后重试"
		return
	}
	query = fmt.Sprintf("DELETE FROM userinfo WHERE userid='%s'", req.TeacherId)
	_, err_ = sqlcontroller.Db.Exec(query)
	if err_ != nil {
		resp.Message = "出错了，请稍后重试"
		return
	}
	resp.Message = "删除成功"
	err = nil
	return
}

// AddCourse implements the ServiceImpl interface.
func (s *ServiceImpl) AddCourse(ctx context.Context, req *server0.AdminAddCourseInfoRequest) (resp *server0.AdminAddCourseInfoResponse, err error) {
	// TODO: Your code here...
	query := fmt.Sprintf("INSERT INTO courseinfo (courseid, coursename, teacherid, credit) VALUES ('%s', '%s', '%s', '%g') ON DUPLICATE KEY UPDATE teacherid='%s'", req.CourseId, req.CourseName, req.TeacherId, req.Credit, req.TeacherId)
	_, err_ := sqlcontroller.Db.Exec(query)
	resp = server0.NewAdminAddCourseInfoResponse()
	if err_ != nil {
		resp.Message = "添加失败，稍后重试"
		return
	}
	resp.Message = "添加成功"
	err = nil
	return
}

// DeleteCourse implements the ServiceImpl interface.
func (s *ServiceImpl) DeleteCourse(ctx context.Context, req *server0.AdminDeleteCourseInfoRequest) (resp *server0.AdminDeleteCourseInfoResponse, err error) {
	// TODO: Your code here...
	query := fmt.Sprintf("DELETE FROM courseinfo WHERE courseid='%s'", req.CourseId)
	_, err_ := sqlcontroller.Db.Exec(query)
	resp = server0.NewAdminDeleteCourseInfoResponse()
	if err_ != nil {
		resp.Message = "出错了，请稍后重试"
	}
	resp.Message = "删除成功"
	err = nil
	return
}
