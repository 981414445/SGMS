package route

import (
	"SGMS/domain/util"
	"encoding/json"
	"fmt"
	"html/template"
	"math"
	"strconv"
	"time"

	"database/sql"
	urlUtil "net/url"

	"sort"

	"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
)

func EhanceTemplate(tplConfig *html.Config) {

	tplConfig.Funcs["json"] = func(obj interface{}) (interface{}, error) {
		b, err := json.Marshal(obj)
		if err != nil {
			return "", err
		}
		return template.JS(b), nil
		// return string(b), nil
	}
	tplConfig.Funcs["encodeurl"] = func(obj interface{}) (string, error) {
		b, err := json.Marshal(obj)
		if err != nil {
			return "", err
		}
		// return template.URL(string(b)), nil
		return strconv.Quote(string(b)), nil
	}
	tplConfig.Funcs["date"] = func(obj interface{}) (string, error) {
		if t, err := obj.(int); err {
			return util.FormatDate(t), nil
		}
		if t, err := obj.(int64); err {
			return util.FormatDate(int(t)), nil
		}
		return "", nil
	}
	tplConfig.Funcs["ToInt"] = func(obj interface{}) (interface{}, error) {
		switch obj.(type) {
		case int:
			return obj, nil
		case int64:
			return int(obj.(int64)), nil
		case float32:
			return int(obj.(float32)), nil
		case float64:
			return int(obj.(float64)), nil
		case bool:
			if obj.(bool) {
				return 1, nil
			}
			return 0, nil
		case string:
			return strconv.Atoi(obj.(string))
		}

		return obj, nil
	}
	tplConfig.Funcs["datetime"] = func(obj interface{}) (string, error) {
		if t, err := obj.(int); err {
			return util.FormatDatetimeShort(t), nil
		}
		if t, err := obj.(int64); err {
			return util.FormatDatetimeShort(int(t)), nil
		}
		return "", nil
	}
	tplConfig.Funcs["dateformat"] = func(obj ...interface{}) (string, error) {
		if 2 > len(obj) {
			return "", nil
		}
		t, ok := obj[0].(int)
		t1, ok1 := obj[0].(int)
		f, ok2 := obj[1].(string)
		if ok && ok2 {
			return util.FormatTimeWith(t, f), nil
		}
		if ok1 && ok2 {
			return util.FormatTimeWith(t1, f), nil
		}
		return "", nil
	}
	tplConfig.Funcs["map"] = func(kvs ...interface{}) map[interface{}]interface{} {
		r := make(map[interface{}]interface{})
		for i := 0; i < len(kvs); i += 2 {
			if i+1 < len(kvs) {
				r[kvs[i]] = kvs[i+1]
			} else {
				r[kvs[i]] = nil
			}
		}
		return r
	}
	tplConfig.Funcs["price"] = func(cent int) string {
		if cent <= 0 {
			return "0.00"
		}
		return fmt.Sprintf("%10.2f", float32(cent)/100)
	}
	tplConfig.Funcs["percent"] = func(percent float64) string {
		if percent == 0 {
			return "0"
		}
		return strconv.Itoa(int(percent*100)) + "%"
	}
	tplConfig.Funcs["unescaped"] = func(x interface{}) interface{} {
		if nil == x {
			return template.HTML("")
		}
		if c, ok := x.(string); ok {
			return template.HTML(c)
		}
		if c, ok := x.(*string); ok {
			if nil == c {
				return template.HTML("")
			}
			return template.HTML(*c)
		}
		if c, ok := x.(sql.NullString); ok {
			return template.HTML(c.String)
		}
		return template.HTML("")

	}
	tplConfig.Funcs["pagnation"] = Pagnation
	tplConfig.Funcs["pagnationx"] = Pagnationx
	tplConfig.Funcs["px"] = Px
	tplConfig.Funcs["select"] = func(options map[int]string, defaultValueObject interface{}, needEmptyOption bool, ps string) interface{} {
		defaultValue := 0
		if defaultValueStr, ok := defaultValueObject.(string); ok {
			defaultValue, _ = strconv.Atoi(defaultValueStr)
		}
		if dvInt, ok := defaultValueObject.(int); ok {
			defaultValue = dvInt
		}
		r := `<select ` + ps + `>`
		if needEmptyOption {
			r += `<option></option>`
		}
		var keys sort.IntSlice = make([]int, len(options))
		i := 0
		for k := range options {
			keys[i] = k
			i++
		}
		keys.Sort()
		for _, k := range keys {
			v := options[k]
			r += `<option value="` + strconv.Itoa(k) + `"`
			if k == defaultValue {
				r += ` selected="selected" `
			}
			r += ">"
			r += v + `</option>`
		}
		r += `</select>`
		return template.HTML(r)
	}

	tplConfig.Funcs["addclassif"] = func(cls string, s1 interface{}, s2 interface{}) (interface{}, error) {
		if s1 == s2 {
			return cls, nil
		}
		return "", nil
	}
	//获取指针值
	tplConfig.Funcs["pval"] = func(v interface{}, defaultValue ...interface{}) interface{} {
		if nil == v {
			if len(defaultValue) > 0 {
				return defaultValue[0]
			}
		}
		if x, ok := v.(*string); ok {
			return *x
		}
		if x, ok := v.(*int); ok {
			return *x
		}
		if x, ok := v.(*int64); ok {
			return *x
		}
		if x, ok := v.(*bool); ok {
			return *x
		}
		return v
	}

	tplConfig.Funcs["add"] = func(d ...interface{}) int {
		r := 0
		for _, v := range d {
			if vi, ok := v.(int); ok {
				r += vi
			}
			if vi, ok := v.(int64); ok {
				r += int(vi)
			}
			if vi, ok := v.(string); ok {
				si, _ := strconv.Atoi(vi)
				r += si
			}

		}
		return r
	}
	tplConfig.Funcs["concat"] = func(ss ...interface{}) string {
		r := ""
		for _, s := range ss {
			ts, _ := util.ToString(s)
			r += ts
		}
		return r
	}

	tplConfig.Funcs["duration"] = func(start int64, end int64) int64 {
		if end-start <= 0 {
			return 0
		}
		return (end - start) / 60
	}

	tplConfig.Funcs["mins"] = func(seconds int64) string {
		if seconds == 0 {
			return ""
		}
		// if seconds > 10 {
		// 	mins := math.Floor(float64(seconds) / 60)
		// 	secs := seconds % 60
		// 	secsStr := ""
		// 	if secs > 9 {
		// 		secsStr = string(secs)
		// 	} else {
		// 		secsStr = "0" + string(secs)
		// 	}
		// 	return strconv.FormatFloat(mins, 'f', -1, 64) + ":" + secsStr
		// }

		// str := "0:0" + string(seconds)
		// return str
		return time.Unix(seconds, 0).Format("04:05")

	}

	tplConfig.Funcs["ifnull"] = func(value, defaultValue interface{}) interface{} {
		if nil == value {
			return defaultValue
		}
		return value
	}
}

//分页
// func Pagnation(url string, si, ps int, total int64) string {
func Pagnation(ctx *iris.Context, total int64) string {
	return Pagnationx(ctx, total, false)
}
func Pagnationx(ctx *iris.Context, total int64, showTotal bool) string {
	ps, ok := ctx.Get("ps").(int)
	if !ok || ps <= 0 {
		ps = 20
	}
	return Px(ctx, total, ps, showTotal, "si")
}
func Px(ctx *iris.Context, total int64, ps int, showTotal bool, siName string) string {
	if nil == ctx {
		return ""
	}
	url := string(ctx.URI().RequestURI())
	si, _ := ctx.URLParamInt(siName)
	if ps <= 0 {
		return ""
	}
	pageCount := int(math.Ceil(float64(total) / float64(ps)))
	current := si / ps
	r := `<div class="pagi-wrapper" >
  <ul class="pagination pagination-warning">
    <li`
	if !hasPreviousPage(si) {
		r += ` class="disabled"`
	}
	r += `>
      <a href="` + getPageUrl(url, si, ps, current-1, pageCount, siName) + `" aria-label="Previous">
        上一页
      </a>
    </li>`
	if pageCount > 0 {
		for _, v := range computePage(current, total) {
			if v > pageCount-1 {
				break
			}
			if -1 == v {
				r += `<li`
				if v == current {
					r += ` class="active" `
				}
				r += `><a href="#">...</a>`
			} else {
				r += `<li`
				if v == current {
					r += ` class="active" `
				}
				r += `><a href="` + getPageUrl(url, si, ps, v, pageCount, siName) + `">` + strconv.Itoa(v+1) + `</a></li>`
			}

		}
	}
	r += `<li`
	if !hasNextPage(si, ps, total) {
		r += ` class="disabled" `
	}
	r += `><a href="` + getPageUrl(url, si, ps, current+1, pageCount, siName) + `" aria-label="Next">
        下一页
      </a>
    </li>
  </ul>`
	if showTotal {
		r += `<span class="pagination pagination-total">共   ` + strconv.Itoa(int(total)) + `   条</span>`
	}
	r += `</div>`
	return r
}
func computePage(current int, total64 int64) []int {
	total := int(total64)
	m0 := []int{0, 1, 2, 3, 4, 5, 6}
	m1 := []int{0, 1, 2, 3, 4, 5, 6, -1, total}
	m2 := []int{0, -1, current - 2, current - 1, current, current + 1, current + 2, -1, total}
	m3 := []int{0, -1, total - 6, total - 5, total - 4, total - 3, total - 2, total - 1, total}
	if total <= 7 {
		return m0
	}
	if current < 5 {
		return m1
	}
	if current > total-5 {
		return m3
	}
	return m2
}

func hasPreviousPage(si int) bool {
	return si > 0
}
func hasNextPage(si, ps int, total int64) bool {
	if total > 0 {
		return total > int64(si+ps)
	}
	return false
}

func getPageUrl(url string, si, ps, index, pageCount int, siName string) string {
	if index < 0 || index >= pageCount {
		return "#"
	}
	urlStruct, _ := urlUtil.ParseRequestURI(url)
	// if -1 != strings.Index(urlStruct.RawQuery, "si") {
	//  siValues, ok := urlStruct.Query().Get() {
	q := urlStruct.Query()
	// if "" == q.Get("si") {
	// 	q.Set("si", "0")
	// }

	q.Set(siName, strconv.Itoa(ps*index))
	// if _, ok := q["si"]; ok {
	// 	q.Set("si", strconv.Itoa(ps*index))
	// 	// url = regexp.MustCompile("si=\\d*").ReplaceAllString(url, "si="+strconv.Itoa(ps*index))
	// } else {
	// 	if -1 == strings.Index(url, "?") {
	// 		url += "?"
	// 	} else {
	// 		url += "&"
	// 	}

	// 	url += "si=" + strconv.Itoa(ps*index)
	// }
	// urlStruct.String()
	// fmt.Println("getPage ", urlStruct.String)
	urlStruct.RawQuery = q.Encode()
	return urlStruct.String()
}
