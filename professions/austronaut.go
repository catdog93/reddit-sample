package professions

type AstronautService interface {
	GetAstronautInfo() string
	AstronautSpecialSkillsToString() string
}

type Astronaut struct {
	ID            uint64 `bson:"_id"`
	*Employee     `json:"employee"`
	SpecialSkills []AustronautSpecialSkill `json:"specialSkill,omitempty"`
	Experience    Year                     `json:"experience"`
	Spacecrafts   []SpacecraftModel        `json:"spacecraft"`
}

type Year int

type AustronautSpecialSkill string

const (
	Captain        AustronautSpecialSkill = "Captain"
	Programmer     AustronautSpecialSkill = "Programmer"
	SystemEngineer AustronautSpecialSkill = "System engineer"
	Technician     AustronautSpecialSkill = "Technician"
)

type SpacecraftModel string

const (
	ChinaShuttle  = "China Shuttle"
	USShuttle     = "US Space Shuttle 123"
	SovietShuttle = "Buran"
)

/*func (astronaut *Astronaut) GetAstronautInfo() string {
	return astronaut.GetEmployeeInfo() + "\nAstronaut" + astronaut.AstronautSpecialSkillsToString()
}*/

func (astronaut *Astronaut) AstronautSpecialSkillsToString() (skillsString string) {
	for _, skill := range astronaut.SpecialSkills {
		skillsString = skillsString + " " + string(skill)
	}
	return
}
