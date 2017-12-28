package face

import "strconv"

var Faces = make(map[string]interface{})

const (
	SGMS_domain_activity_ActivityItem = "SGMS.domain.activity.ActivityItem"
	SGMS_domain_activity_Activity     = "SGMS.domain.activity.Activity"
	SGMS_domain_round_Round           = "SGMS.domain.round.Round"
	SGMS_domain_paper_IPaper          = "SGMS.domain.paper.IPaper"
)

type PageParam struct {
	Si, Ps  int
	NoLimit bool
}

type OrderKey struct {
	Key   string
	Order bool //true : asc , false : desc
}

type OrderKeys []OrderKey

func NewOrderKeys(keys ...OrderKey) OrderKeys {
	return keys
}
func (o OrderKeys) Sql() string {
	if len(o) <= 0 {
		return ""
	}
	s := " order by"
	for _, i := range o {
		s += i.Key
		if i.Order {
			s += " asc"
		} else {
			s += " desc"
		}
		s += ","
	}
	return s[0 : len(s)-1]
}

func (this *PageParam) PageString() string {
	ps := this.Ps
	si := this.Si
	if this.Si < 0 {
		si = 0
	}
	if this.Ps > 100 {
		ps = 100
	}
	if this.Ps <= 0 {
		ps = 20
	}
	return strconv.Itoa(si) + "," + strconv.Itoa(ps)
}

func (this *PageParam) Limit() string {
	if this.NoLimit {
		return ""
	}
	return " limit " + this.PageString()
}

type UserPageParam struct {
	PageParam
	Uid int
}

type Select2Param struct {
	// 继承PageParam  Si,startIndex Ps，PageSize
	PageParam
	Key string
	Id  int
}

type Select2Result struct {
	Id    int
	Value string
}

const (
	NullDuan     = -100
	NullMonth    = -1
	NullCategory = -1
	NullOnline   = -1
	NullUid      = 0
	NullStatus   = -1
	NullType     = -1
)

type IdStruct struct {
	Id int
}

type IdStructs []IdStruct

func (t IdStructs) Map() map[int]bool {
	r := make(map[int]bool, len(t))
	for _, v := range t {
		r[v.Id] = true
	}
	return r
}
func (t IdStructs) Ids() []int {
	r := make([]int, len(t))
	for i, v := range t {
		r[i] = v.Id
	}
	return r
}
