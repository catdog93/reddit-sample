package professions

type SoftwareDeveloper struct {
	ID         uint64 `bson:"_id"`
	*Employee  `json:"employee"`
	SkillLevel Skill `json:"skillLevel"` // intern/trainee, Junior, Middle, Senior developer or Architector etc.
	//Position          Position  `json:"position,omitempty"`          // TeamLead, PM etc.
	TechnologiesStack string    `json:"technologiesStack,omitempty"` // A technology stack, also called a solutions stack or a data ecosystem
	Direction         Direction `json:"direction,omitempty"`         // Back or front-end web development / Mobile or desktop
}

type Skill string

const (
	Trainee     Skill = "trainee"
	Intern      Skill = "intern"
	Junior      Skill = "junior"
	Middle      Skill = "middle"
	Senior      Skill = "senior"
	Architector Skill = "architector"
)

type Position string

const (
	TeamLead       Position = "TeamLead"
	ProjectManager Position = "Project Manager"
)

type Direction string // front, back, fullstack relate to web development

const (
	Frontend  Direction = "front-end"
	Backend   Direction = "back-end"
	FullStack Direction = "full stack"
	Mobile    Direction = "mobile"
	Desktop   Direction = "desktop"
)
