package route

import (
	"SGMS/domain/exception"
	"SGMS/domain/face"
	"SGMS/domain/factory/basef"
	"SGMS/domain/util"
	"mime/multipart"
	"regexp"
	"strconv"
	"unicode/utf8"

	"image"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/guregu/null"
	"github.com/kataras/iris"
)

type ValidatorError struct {
	Errors map[string]string
}

func (this *ValidatorError) error() string {
	return "validator errors"
}

type IValidatorContext interface {
	CheckPathParam(name string) *Validator
	CheckQuery(name string) *Validator
	CheckBody(name string) *Validator
	CheckFile(name string) IFileValidator
	CheckAllFile() []FileUrl

	HasError() bool
	AddError(name, tip string)
	GetErrors() map[string]string
	Check()
	GetContext() *iris.Context
}
type ValidatorContext struct {
	ctx    *iris.Context
	errors map[string]string
	goOn   bool
}

func NewValidatorContext(ctx *iris.Context) IValidatorContext {
	return &ValidatorContext{ctx, make(map[string]string), false}

}
func (this *ValidatorContext) CheckBody(name string) *Validator {
	v := new(Validator)
	v.vc = this
	v.Name = name

	v.NameExists = this.ctx.PostArgs().Has(name)
	if !v.NameExists {
		mf, _ := this.ctx.MultipartForm()
		if nil != mf {
			if _, ok := mf.Value[name]; ok {
				v.NameExists = true
			}
		}
	}
	if v.NameExists {
		v.Value = this.ctx.FormValueString(name)
	}
	v.goOn = true
	return v
}

type IFileValidator interface {
	IsEmpty() bool
	//文件可以为空
	Empty(tip ...string) IFileValidator
	//文件不能为空
	NotEmpty(tip ...string) IFileValidator
	//文件最大长度，bytes
	Len(maxBytes int, tip ...string) IFileValidator
	//eg ["jpg","png"]
	SuffixIn(s []string, tip ...string) IFileValidator
	//eg. ["100x200","300x300"]
	Resize(sizes []string, rect *image.Rectangle) (fileName string, result *face.FileSaveImageResult)
	//保存到文件系统，返回文件名和文件路径很
	Save() (fileName string, fileURL string)
	//获取文件字节内容
	Byte() (fileName string, content []byte)
	//获取文件内容
	String() (fileName string, content string)
}
type FileHeader struct {
	FileName, ContentType string
	FileLength            int64
}

type FileValidator struct {
	Validator
	url, fileContent, docHtml, resizedUrls string
	isEmtpy                                bool
	file                                   *multipart.FileHeader
}
type FileUrl struct {
	FileName, FileURL string
}

//一起保持多个文件 ，
func (this *ValidatorContext) CheckAllFile() []FileUrl {
	return nil
}

func (this *ValidatorContext) CheckFile(name string) IFileValidator {
	v := new(FileValidator)
	v.vc = this
	v.Name = name
	mf, _ := this.ctx.MultipartForm()
	v.isEmtpy = true
	if nil != mf {
		if files, ok := mf.File[name]; ok {
			v.NameExists = true
			if len(files) > 0 {
				v.file = files[0]
				v.isEmtpy = false
			}
		}
	}
	v.goOn = true
	return v
}

func (this *ValidatorContext) CheckQuery(name string) *Validator {
	v := new(Validator)
	v.vc = this
	v.Name = name
	v.NameExists = this.ctx.QueryArgs().Has(name)
	if v.NameExists {
		v.Value = this.ctx.URLParam(name)
	}
	v.goOn = true
	return v
}
func (this *ValidatorContext) CheckPathParam(name string) *Validator {
	v := new(Validator)
	v.vc = this
	v.Name = name
	v.Value = this.ctx.GetString(name)
	v.NameExists = "" == v.Value
	v.goOn = true
	return v
}
func (this *ValidatorContext) HasError() bool {
	return 0 < len(this.errors)
}
func (this *ValidatorContext) AddError(name, tip string) {
	this.errors[name] = tip
}
func (this *ValidatorContext) GetErrors() map[string]string {
	if len(this.errors) <= 0 {
		return nil
	}
	return this.errors
}

func (this *ValidatorContext) Check() {
	if this.HasError() {
		panic(&ValidatorError{this.GetErrors()})
	}
}
func (this *ValidatorContext) GetContext() *iris.Context {
	return this.ctx
}

//获取默认提示消息
func getDefaultTip(tip []string, defaultValue string) string {
	if len(tip) > 0 {
		defaultValue = tip[0]
	}
	return defaultValue
}

type Validator struct {
	vc IValidatorContext
	//	args        *fasthttp.Args
	NameExists, goOn, AllowEmpty bool
	Name, Value                  string
}

func (this *Validator) AddError(tip string) {
	this.Stop()
	this.vc.AddError(this.Name, tip)
}
func (this *Validator) Stop() {
	this.goOn = false
}
func (this *Validator) Optional() *Validator {
	if !this.NameExists {
		this.AllowEmpty = true
		this.Stop()
	}
	return this
}

func (this *Validator) NotEmpty(tip ...string) *Validator {
	if this.goOn && "" == this.Value {
		this.AddError(getDefaultTip(tip, this.Name+"不能为空"))
	}
	return this
}
func (this *Validator) Empty() *Validator {
	if !this.NameExists || "" == this.Value {
		this.Stop()
		this.AllowEmpty = true
	}
	return this
}
func (this *Validator) NotBlank(tip ...string) *Validator {
	if ok, err := regexp.MatchString("^\\s*$", this.Value); ok || nil != err {
		this.AddError(getDefaultTip(tip, this.Name+"不能为空"))
	}
	return this
}
func (this *Validator) Exist(tip ...string) *Validator {
	if this.goOn && !this.NameExists {
		this.AddError(getDefaultTip(tip, this.Name+"不能为空"))
	}
	return this
}
func (this *Validator) Match(reg string, tip ...string) *Validator {
	if this.goOn {
		if ok, err := regexp.MatchString(reg, this.Value); !ok || nil != err {
			this.AddError(getDefaultTip(tip, this.Name+"不能为空"))
		}
	}
	return this
}

func (this *Validator) IsInt(tip ...string) *Validator {
	if this.goOn && !govalidator.IsInt(this.Value) {
		this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
	}
	return this
}
func (this *Validator) IsFloat(tip ...string) *Validator {
	if this.goOn && !govalidator.IsFloat(this.Value) {
		this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
	}
	return this
}

//len in [min,max]
func (this *Validator) Len(min, max int, tip ...string) *Validator {
	charLen := utf8.RuneCountInString(this.Value)
	if this.goOn && (charLen < min || charLen > max) {
		this.AddError(getDefaultTip(tip, this.Name+"长度不正确"))
	}
	return this
}

//type ValidateAsesstFn func(value string)  bool
func (this *Validator) Ensure(assertion func(value string) bool, tip ...string) *Validator {
	if this.goOn && (!assertion(this.Value)) {
		this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
	}
	return this
}

func (this *Validator) EnsureNot(assertion func(value string) bool, tip ...string) *Validator {
	if this.goOn && (assertion(this.Value)) {
		this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
	}
	return this
}

func (this *Validator) IsEmail(tip ...string) *Validator {
	if this.goOn && !govalidator.IsEmail(this.Value) {
		this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
	}
	return this
}
func (this *Validator) IsURL(tip ...string) *Validator {
	if this.goOn && !govalidator.IsURL(this.Value) {
		this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
	}
	return this
}

func (this *Validator) IsNumeric(tip ...string) *Validator {
	if this.goOn && !govalidator.IsNumeric(this.Value) {
		this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
	}
	return this
}
func (this *Validator) IsIP(tip ...string) *Validator {
	if this.goOn && !govalidator.IsIP(this.Value) {
		this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
	}
	return this
}

func (this *Validator) IsIPv4(tip ...string) *Validator {
	if this.goOn && !govalidator.IsIPv4(this.Value) {
		this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
	}
	return this
}

func (this *Validator) IsIPv6(tip ...string) *Validator {
	if this.goOn && !govalidator.IsIPv6(this.Value) {
		this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
	}
	return this
}
func (this *Validator) IsASCII(tip ...string) *Validator {
	if this.goOn && !govalidator.IsASCII(this.Value) {
		this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
	}
	return this
}

func (this *Validator) IsAlpha(tip ...string) *Validator {
	if this.goOn && !govalidator.IsAlpha(this.Value) {
		this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
	}
	return this
}
func (this *Validator) IsAlphanumeric(tip ...string) *Validator {
	if this.goOn && !govalidator.IsAlphanumeric(this.Value) {
		this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
	}
	return this
}

func (this *Validator) InRange(min, max int, tip ...string) *Validator {
	intValue := this.ToInt(0, getDefaultTip(tip, this.Name+"取值不正确"))
	if this.goOn && (intValue < min || intValue > max) {
		this.AddError(getDefaultTip(tip, this.Name+"取值不正确"))
	}
	return this
}
func (this *Validator) In(values []string, tip ...string) *Validator {
	if this.goOn {
		s := false
		for _, v := range values {
			if v == this.Value {
				s = true
			}
		}
		if !s {
			this.AddError(getDefaultTip(tip, this.Name+"取值不正确"))
		}
	}
	return this
}

//Sanitizers

func (this *Validator) ToString(defaultValue ...string) string {
	if "" == this.Value && 0 < len(defaultValue) {
		return defaultValue[0]
	}
	return this.Value
}

func (this *Validator) DeHtml(defaultValue ...string) string {
	if "" == this.Value && 0 < len(defaultValue) {
		return defaultValue[0]
	}
	return util.DeHtmlTag(this.Value)
}
func (this *Validator) ToNullString(defaultValue ...string) null.String {
	if "" == this.Value {
		return null.NewString("", false)
	}
	v := this.ToString(defaultValue...)
	return null.NewString(v, true)
}
func (this *Validator) ToInt(defaultValue int, tip ...string) int {
	if "" == this.Value {
		return defaultValue
	}
	this.IsInt(getDefaultTip(tip, this.Name+"格式不正确"))
	if this.goOn {
		v, err := strconv.Atoi(this.Value)
		if nil != err {
			this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
			return 0
		}
		return v
	}
	return 0
}
func (this *Validator) ToNullInt64(defaultValue int64, tip ...string) null.Int {
	if "" == this.Value {
		return null.NewInt(0, false)
	}
	v := this.ToInt64(defaultValue, 10, tip...)
	return null.NewInt(v, true)
}
func (this *Validator) ToInt64(defaultValue int64, radix int, tip ...string) int64 {
	if "" == this.Value {
		return defaultValue
	}
	this.IsInt(getDefaultTip(tip, this.Name+"格式不正确"))
	if this.goOn {
		v, err := strconv.ParseInt(this.Value, radix, 64)
		if nil != err {
			this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
			return defaultValue
		}
		return v
	}
	return defaultValue
}
func (this *Validator) ToFloat(defaultValue float32, tip ...string) float32 {
	if "" == this.Value {
		return defaultValue
	}
	this.IsFloat(getDefaultTip(tip, this.Name+"格式不正确"))
	if this.goOn {
		v, err := strconv.ParseFloat(this.Value, 32)
		if nil != err {
			this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
			return defaultValue
		}
		return float32(v)
	}
	return defaultValue
}
func (this *Validator) ToBool(defaultValue bool, tip ...string) bool {
	if "" == this.Value {
		return defaultValue
	}
	if this.goOn {
		v, err := strconv.ParseBool(this.Value)
		if nil != err {
			this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
			return false
		}
		return v
	}
	return false
}

//ToDate layout :2006-01-02 15:04:05
func (this *Validator) ToDate(layout string, tip ...string) int {
	if "" == this.Value {
		if this.AllowEmpty {
			return -1
		}
		this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
		return -1
	}
	if this.goOn {
		date, err := util.ParseTimeLocal(layout, this.Value)
		if nil != err {
			this.AddError(getDefaultTip(tip, this.Name+"格式不正确"))
			return -1
		}
		return int(date.Unix())
	}
	return -1
}

//ToDateMin 2006-01-02 15:04
func (this *Validator) ToDateMin(tip ...string) int {
	return this.ToDate("2006-01-02 15:04", tip...)
}

//ToDateDay 2006-01-02
func (this *Validator) ToDateDay(tip ...string) int {
	return this.ToDate("2006-01-02", tip...)
}

func (this *Validator) Trim() *Validator {
	this.Value = regexp.MustCompile("^\\s+|\\s+$").ReplaceAllString(this.Value, "")
	return this
}

func (this *Validator) TrimLeft() *Validator {
	this.Value = regexp.MustCompile("^\\s+").ReplaceAllString(this.Value, "")
	return this
}
func (this *Validator) TrimRight() *Validator {
	this.Value = regexp.MustCompile("\\s+$").ReplaceAllString(this.Value, "")
	return this
}

//文件校验器

func (this *FileValidator) NotEmpty(tip ...string) IFileValidator {
	if this.goOn && this.isEmtpy {
		this.AddError(getDefaultTip(tip, this.Name+"不能为空"))
	}
	return this
}
func (this *FileValidator) Empty(tip ...string) IFileValidator {
	if !this.NameExists || this.isEmtpy {
		this.Stop()
		this.AllowEmpty = true
	}
	return this
}
func (this *FileValidator) Len(maxSize int, tip ...string) IFileValidator {
	lenStr := this.file.Header.Get("Content-Length")
	len, _ := strconv.Atoi(lenStr)
	if this.goOn && len > maxSize {
		this.AddError(getDefaultTip(tip, this.Name+"长度超过限度。"))
	}
	return this
}
func (this *FileValidator) SuffixIn(s []string, tip ...string) IFileValidator {
	ext := filepath.Ext(this.file.Filename)
	exists := false
	for _, v := range s {
		if "."+v == ext {
			exists = true
		}
	}
	if this.goOn && !exists {
		this.AddError(getDefaultTip(tip, this.Name+"文件类型不正确。"))
	}
	return this
}
func (this *FileValidator) Resize(sizes []string, rect *image.Rectangle) (fileName string, result *face.FileSaveImageResult) {
	f, err := this.file.Open()
	if nil != err {
		panic(err)
	}
	defer f.Close()
	fileName = this.file.Filename
	result, err = basef.NewFileRepo().SaveImage(f, fileName, sizes, rect)
	if nil != err {
		panic(exception.NewParamError(map[string]string{this.Name: "请添加图片！"}))
	}
	return
}
func (this *FileValidator) Save() (fileName string, fileURL string) {
	f, err := this.file.Open()
	if nil != err {
		panic(err)
	}
	defer f.Close()
	fileName = this.file.Filename
	ext := strings.ToLower(filepath.Ext(fileName))
	if ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".gif" || ext == ".bmp" {
		repo, err := basef.NewFileRepo().SaveImage(f, fileName, nil, nil)
		fileURL = repo.RawImage
		if nil != err {
			panic(exception.NewParamError(map[string]string{this.Name: "请添加图片！"}))
		}
		return
	}
	if ext == ".doc" {
		fileURL, err = basef.NewFileRepo().Save(f, face.FILE_REPO_TYPE_DOC, fileName)
		if nil != err {
			panic(exception.NewParamError(map[string]string{this.Name: "文件不正确！"}))
		}
		return
	}
	fileURL, err = basef.NewFileRepo().Save(f, face.FILE_REPO_TYPE_BIN, fileName)
	if nil != err {
		panic(exception.NewParamError(map[string]string{this.Name: "保存文件失败"}))
	}
	return
}
func (this *FileValidator) Byte() (fileName string, content []byte) {
	f, err := this.file.Open()
	if nil != err {
		panic(err)
	}
	defer f.Close()
	content, err = ioutil.ReadAll(f)
	if nil != err {
		panic(err)
	}
	fileName = this.file.Filename
	return
}
func (this *FileValidator) String() (fileName string, content string) {
	f, err := this.file.Open()
	if nil != err {
		panic(err)
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if nil != err {
		panic(err)
	}
	fileName = this.file.Filename
	content = string(bs)
	return
}
func (this *FileValidator) IsEmpty() bool {
	return this.isEmtpy
}
