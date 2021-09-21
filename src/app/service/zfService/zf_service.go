package zfService

import (
	"encoding/json"
	"funnel/app/apis/zf"
	"funnel/app/errors"
	"funnel/app/model"
	"funnel/app/service"
	"funnel/app/utils/fetch"
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
)

// 获得学期课表
func GetLessonsTable(stu *model.User, year string, term string) (interface{}, error) {
	res, err := fetchTermRelatedInfo(stu, zf.ZfClassTable(), year, term)
	if err != nil {
		return nil, err
	}
	var f model.LessonsTableRawInfo
	err = json.Unmarshal([]byte(res), &f)
	return model.TransformLessonTable(&f), err
}

//获得考试信息
func GetExamInfo(stu *model.User, year string, term string) (interface{}, error) {
	res, err := fetchTermRelatedInfo(stu, zf.ZfExamInfo(), year, term)
	if err != nil {
		return nil, err
	}
	var f model.ExamRawInfo
	err = json.Unmarshal([]byte(res), &f)
	return model.TransformExamInfo(&f), err
}
func GetScoreDetail(stu *model.User, year string, term string) (interface{}, error) {
	res, err := fetchTermRelatedInfo(stu, zf.ZfScoreDetail(), year, term)
	if err != nil {
		return nil, err
	}
	var f model.ScoreDetailRawInfo
	err = json.Unmarshal([]byte(res), &f)
	return model.TransformScoreDetailInfo(&f), err
}
func GetScore(stu *model.User, year string, term string) (interface{}, error) {
	res, err := fetchTermRelatedInfo(stu, zf.ZfScore(), year, term)
	if err != nil {
		return nil, err
	}
	var f model.ScoreRawInfo
	err = json.Unmarshal([]byte(res), &f)
	return model.TransformScoreInfo(&f), err
}

func fetchTermRelatedInfo(stu *model.User, requestUrl, year, term string) (string, error) {
	f := fetch.Fetch{}
	f.Init()
	f.Cookie = append(f.Cookie, &stu.Session)
	if term == "上" {
		term = "3"
	} else if term == "下" {
		term = "12"
	} else if term == "短" {
		term = "16"
	}
	requestData := genTermRelatedInfoReqData(year, term)
	s, err := f.PostForm(requestUrl, requestData)

	if len(s) == 0 {
		service.ForgetUserByUsername(service.ZFPrefix, stu.Username)
		return "", errors.ERR_Session_Expired
	}
	if err != nil {
		return "", err
	}
	return string(s), nil
}

func GetTrainingPrograms(stu *model.User) ([]byte, error) {
	f := fetch.Fetch{}
	f.Init()
	f.Cookie = append(f.Cookie, &stu.Session)
	response, err := f.GetRaw(zf.ZfUserInfo())

	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	s, exist := doc.Find("#pyfaxx_id").Attr("value")
	if exist {
		res, _ := f.GetRaw(zf.ZfPY() + s)
		s, _ := ioutil.ReadAll(res.Body)
		return s, nil
	}
	return nil, nil
}

func GetEmptyRoomInfo(stu *model.User, year string, term string, campus string, weekday string, week string, classPeriod string) (string, error) {
	f := fetch.Fetch{}
	f.Init()
	f.Cookie = append(f.Cookie, &stu.Session)
	if term == "上" {
		term = "3"
	} else if term == "下" {
		term = "12"
	} else if term == "短" {
		term = "16"
	}
	if campus == "朝晖" {
		campus = "01"
	} else if campus == "屏峰" {
		campus = "02"
	} else if campus == "莫干山" {
		campus = "A61400B98155D41AE0550113465EF1CF"
	}
	requestData := genEmptyRoomReqData(year, term, campus, week, weekday, classPeriod)
	s, err := f.PostForm(zf.ZfEmptyClassRoom(), requestData)

	if len(s) == 0 {
		service.ForgetUserByUsername(service.ZFPrefix, stu.Username)
		return "", errors.ERR_Session_Expired
	}
	if err != nil {
		return "", err
	}

	return string(s), nil
}

func GetUser(username string, password string) (*model.User, error) {
	user, err := service.GetUser(service.ZFPrefix, username, password)
	if err != nil {
		return login(username, password)
	}
	return user, err
}
