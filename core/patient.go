package core

type Patient struct {
	// input
	Age                int64  `json:"age"`
	Communication      string `json:"communication"`
	DailyLivingSkill   string `json:"daily_living_skill"`
	Socialization      string `json:"socialization"`
	DangerousBehaviors string `json:"dangerous_behaviors"`

	// logic handle
	Group12MonthTo5YearsAge bool

	CommunicationLow  bool
	CommunicationHigh bool

	DailyLivingSkillLow  bool
	DailyLivingSkillHigh bool

	SocializationLow  bool
	SocializationHigh bool

	// output
	HaveRecommendation        bool
	TraditionalABARecommended string
	TypeOfProgram             string
}

func (p *Patient) DoRecommend1() {
	p.HaveRecommendation = true
	p.TraditionalABARecommended = "Y"
	p.TypeOfProgram = "Comprehensive program - heavier goal weight on communication skills and behavior reduction"
}
