package core

type Patient struct {
	// input
	Age               int64
	Communication     string
	DailyLivingSkills string
	Socialization     string

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
	DangerousBehaviors        bool
	TraditionalABARecommended bool
	TypeOfProgram             string
}
