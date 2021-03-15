package calloption

type peopleOption struct {
	Sex     string
	Age     int32
	Height  float32
	Weight  float32
	Color   string
	Country string
	Address string
}

type PeopleOption func(po *peopleOption)

func WithSex(sex string) PeopleOption {
	return func(po *peopleOption) {
		po.Sex = sex
	}
}

func WithAge(age int32) PeopleOption {
	return func(po *peopleOption) {
		po.Age = age
	}
}

func WithHeight(height float32) PeopleOption {
	return func(po *peopleOption) {
		po.Height = height
	}
}

func WithWeight(weight float32) PeopleOption {
	return func(po *peopleOption) {
		po.Weight = weight
	}
}

func WithColor(color string) PeopleOption {
	return func(po *peopleOption) {
		po.Color = color
	}
}

func WithCountry(country string) PeopleOption {
	return func(po *peopleOption) {
		po.Country = country
	}
}

func WithAddress(address string) PeopleOption {
	return func(po *peopleOption) {
		po.Address = address
	}
}

// People 结构定义
type People struct {
	Id   string
	Name string
	Opt  peopleOption
}

// 默认值
func defaultPeopleOption() peopleOption {
	return peopleOption{
		Sex:     "female",
		Age:     18,
		Height:  170,
		Weight:  100,
		Color:   "yellow",
		Country: "China",
		Address: "HeNan",
	}
}

// 初始化入口
func NewPeople(id string, name string, opts ...PeopleOption) *People {
	cc := People{
		Id:   id,
		Name: name,
		Opt:  defaultPeopleOption(),
	}

	for _, opt := range opts {
		opt(&cc.Opt)
	}

	return &cc
}
