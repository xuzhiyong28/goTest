package json_iterator

import (
	jsoniter "github.com/json-iterator/go"
	"testing"
)

var str = `{
    "stat": 1,
    "code": 0,
    "msg": "成功",
    "data": {
        "courseInfos": [
            {
                "courseId": "161935",
                "course_name": "高三5科语数英物化提分特训班（20课时）",
                "gradeId": "13",
                "gradeName": "高三",
                "subjectName": "数学",
                "difficultyName": "目标A+",
                "type1Name": "特训班",
                "termIds": "4",
                "schoolTime": "1月29日-1月31日上课（详情见大纲）",
                "price": "799",
                "actualPrice": 20,
                "subjectId": 2,
                "type_1_id": "2064",
                "type_2_id": "2656",
                "type_3_id": "2661"
            }
        ],
        "ext": {
            "price": "799",
            "sale": "20",
            "wxShareObj": "{\"title\":\"语数双科提分特训班\",\"desc\":\"20元抢20课时名师直播好课，下单加送国风限量教辅礼包！\",\"imgUrl\":\"https://activity.xueersi.com/topic/growth/common/images/common/xes-logo.png\",\"miniImgUrl\":\"https://hw.xesimg.com/biz-growth-storage/operations/groupon/20201215/0f4f57c8c6f88b66b059881cfb050527.png\"}",
            "abTestPackage": "{\"h5\":[\"20_20ChineseA_sucaiH5\",\"Azhifudanye_H5\",\"Apintuan_H5\",\"Axueyuanpinglun_ceshi\"],\"smallProgram\":[]}"
        },
        "resourceConfig": {
            "bookImg": [
                {
                    "type": "img",
                    "url": "https://ek.xesimg.com/biz-growth-storage/activity/upload/20201216/86724030cc04b4ea029fe31e4042880c.png",
                    "name": "初高H5&amp;小程序随材图 .png"
                }
            ],
            "detailImg": [
                {
                    "type": "img",
                    "url": "https://oo.xesimg.com/biz-growth-storage/activity/upload/20210126/7e9bd0cb1810c9de1432bc8c570ce3f0.png",
                    "name": "1@2x.png"
                },
                {
                    "type": "img",
                    "url": "https://ek.xesimg.com/biz-growth-storage/activity/upload/20210126/17bdceacf83feda19ee854b50c469152.png",
                    "name": "2@2x.png"
                },
                {
                    "type": "img",
                    "url": "https://hw.xesimg.com/biz-growth-storage/activity/upload/20210126/21189d1f915feeaf0f103a5dfbe77950.png",
                    "name": "3@2x.png"
                },
                {
                    "type": "img",
                    "url": "https://mr.xesimg.com/biz-growth-storage/activity/upload/20210126/fb106f25b740fc7e559bab5568958fe5.png",
                    "name": "4@2x.png"
                },
                {
                    "type": "img",
                    "url": "https://oo.xesimg.com/biz-growth-storage/activity/upload/20210126/93f595fa4cd4fd72683f2faf1b33e86b.png",
                    "name": "5@2x.png"
                },
                {
                    "type": "img",
                    "url": "https://oo.xesimg.com/biz-growth-storage/activity/upload/20210126/b5d789b9aa8833367d4d74af5d20bfc5.png",
                    "name": "7@2x.png"
                },
                {
                    "type": "img",
                    "url": "https://oo.xesimg.com/biz-growth-storage/activity/upload/20210126/389308e953ac695a972840c796443de5.png",
                    "name": "6@2x.png"
                },
                {
                    "type": "img",
                    "url": "https://oo.xesimg.com/biz-growth-storage/activity/upload/20210126/29465ee6a664f36c4ae4b154f60c5b01.png",
                    "name": "矩阵@2x.png"
                }
            ],
            "headImg": [
                {
                    "type": "img",
                    "url": "https://ek.xesimg.com/biz-growth-storage/activity/upload/20210126/45e73133144cb179a1f33d85e5e02c68.png",
                    "name": "头图@2x.png"
                }
            ],
            "bookTextDesc": "多科目组合课程，为保障学习效果，暂不支持调课哦",
            "bookTextWx": "https://ek.xesimg.com/biz-growth-storage/operations/groupon/20201216/cda7dc272b4757efa7a21c3de30f87e8.png",
            "grade_id": "13",
            "feitoufang": "140012,140013,140014",
            "videoInfo": "{\"videoUrl\":\"https://activity.xueersi.com/oss/resource/%E9%AB%98%E4%B8%AD-1611580251682.mp4\",\"videoPoster\":\"https://activity.xueersi.com/oss/resource/%E9%AB%98%E4%B8%AD-1611658207877.png\"}"
        }
    }
}`

type T struct {
	Stat int    `json:"stat"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		CourseInfos []struct {
			CourseId       string `json:"courseId"`
			CourseName     string `json:"course_name"`
			GradeId        string `json:"gradeId"`
			GradeName      string `json:"gradeName"`
			SubjectName    string `json:"subjectName"`
			DifficultyName string `json:"difficultyName"`
			Type1Name      string `json:"type1Name"`
			TermIds        string `json:"termIds"`
			SchoolTime     string `json:"schoolTime"`
			Price          string `json:"price"`
			ActualPrice    int    `json:"actualPrice"`
			SubjectId      int    `json:"subjectId"`
			Type1Id        string `json:"type_1_id"`
			Type2Id        string `json:"type_2_id"`
			Type3Id        string `json:"type_3_id"`
		} `json:"courseInfos"`
		Ext struct {
			Price         string `json:"price"`
			Sale          string `json:"sale"`
			WxShareObj    string `json:"wxShareObj"`
			AbTestPackage string `json:"abTestPackage"`
		} `json:"ext"`
		ResourceConfig struct {
			BookImg []struct {
				Type string `json:"type"`
				Url  string `json:"url"`
				Name string `json:"name"`
			} `json:"bookImg"`
			DetailImg []struct {
				Type string `json:"type"`
				Url  string `json:"url"`
				Name string `json:"name"`
			} `json:"detailImg"`
			HeadImg []struct {
				Type string `json:"type"`
				Url  string `json:"url"`
				Name string `json:"name"`
			} `json:"headImg"`
			BookTextDesc string `json:"bookTextDesc"`
			BookTextWx   string `json:"bookTextWx"`
			GradeId      string `json:"grade_id"`
			Feitoufang   string `json:"feitoufang"`
			VideoInfo    string `json:"videoInfo"`
		} `json:"resourceConfig"`
	} `json:"data"`
}

var Json1 = jsoniter.ConfigDefault
var T1 T
var Tmap map[string]interface{}
var TByte []byte

func init() {
	TByte = []byte(str)
	jsoniter.Unmarshal(TByte, &T1)
	jsoniter.Unmarshal(TByte, &Tmap)
}

func BenchmarkEncodeStructJson1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Json1.Marshal(T1)
	}
}
