package calc

type Character interface {
	GetClasses() string // enchantress, summoner, ...
	GetType() string // pure_dealer, synergy_dealer, buffer, ...
	GetStatus(inDungeon bool) interface{}
	GetDetails(inDungeon bool) interface{}
	SetEquipment(id int)
	SetEquipmentSet(setId int)
	GetSkillDetail(id int, inDungeon bool) interface{}
}