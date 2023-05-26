package conf

import (
	"fmt"
	"io/ioutil"
	"oa-review/logger"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

type OaReviewConf struct {
	root gjson.Result
	file string
}

type globalConf struct {
	Conf *OaReviewConf
	// others
}

var globalConfInfo *globalConf

// 读取全局配置文件数据
func GetConfig() *globalConf {
	if globalConfInfo == nil {
		logger.Fatal("init global config must be called before GET")
		return nil
	}
	return globalConfInfo
}

func Read(filePath string) (*OaReviewConf, error) {
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %s, err: %s", filePath, err)
	}

	conf, err := Parse(string(dat))

	if err != nil {
		return nil, errors.WithMessagef(err, "in %s", filePath)
	}
	logger.Info("read config file path success")
	conf.file = filePath
	return conf, nil
}

func Parse(jsonData string) (*OaReviewConf, error) {
	if !gjson.Valid(jsonData) {
		return nil, fmt.Errorf("invalid config json")
	}
	res := gjson.Parse(jsonData)

	logger.Info("parse json data success")
	c := &OaReviewConf{
		root: res,
	}
	return c, nil
}

func InitGlobalConfig(conf *OaReviewConf) {
	globalConfInfo = &globalConf{
		Conf: conf,
	}
	logger.Info("Init global config success")
}

func (h *OaReviewConf) GetInt(path string, def int64) int64 {
	val := h.root.Get(path)
	if !val.Exists() {
		return def
	}
	return val.Int()
}

func (h *OaReviewConf) GetBool(path string, def bool) bool {
	val := h.root.Get(path)
	if !val.Exists() {
		return def
	}
	return val.Bool()
}

func (h *OaReviewConf) GetFloat(path string, def float64) float64 {
	val := h.root.Get(path)
	if !val.Exists() {
		return def
	}
	return val.Float()
}

func (h *OaReviewConf) GetString(path string, def string) string {
	val := h.root.Get(path)
	if !val.Exists() {
		return def
	}
	return val.String()
}

func (h *OaReviewConf) MustGetInt(path string) int64 {
	val := h.root.Get(path)
	if !val.Exists() {
		panic(fmt.Sprintf("cannot get config in %s %s", h.file, path))
	}
	if val.Type != gjson.Number {
		panic(fmt.Sprintf("expect number type in %s %s", h.file, path))
	}
	return val.Int()
}

func (h *OaReviewConf) MustGetFloat(path string) float64 {
	val := h.root.Get(path)
	if !val.Exists() {
		panic(fmt.Sprintf("cannot get config in %s %s", h.file, path))
	}
	if val.Type != gjson.Number {
		panic(fmt.Sprintf("expect number type in %s %s", h.file, path))
	}
	return val.Float()
}

func (h *OaReviewConf) MustGetString(path string) string {
	val := h.root.Get(path)
	if !val.Exists() {
		panic(fmt.Sprintf("cannot get config in %s %s", h.file, path))
	}
	if val.Type != gjson.String {
		panic(fmt.Sprintf("expect string type in %s %s", h.file, path))
	}
	return val.String()
}
