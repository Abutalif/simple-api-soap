package entities

type KHCAllowedProject struct {
	Id              string `json:"id"`
	RespCompNameKaz string `json:"uakiletti_kompaniya_atauy"`
	RespCompNameRus string `json:"naimenovanie_upolnomochennoi_kompanii"`
	BuilderNamerKaz string `json:"kuryltaishy_atauy"`
	BuilderNameRus  string `json:"naimenovanie_zastroishika"`
	ProjectNameKaz  string `json:"zhoba_atauy"`
	ProjactNameRus  string `json:"naimenovanie_proekta"`
}

type AstanaBuilding struct {
	Id            string `json:"id"`
	ApparmentArea string `json:"ploshad_kv"`
	Orderer       string `json:"zakazhik"`
	Type          string `json:"harakt_obekta"`
	LandArea      string `json:"ploshad_zemli"`
	FloorNum      string `json:"kol_etaj"`
	Location      string `json:"mestopolojenie"`
	ApparmentNum  string `json:"kol_kv"`
	Email         string `json:"el.pochta"`
	Phone         string `json:"telefon"`
	Name          string `json:"naimenovanie__jk"`
	Region        string `json:"gorod/rain"`
}

type About struct {
	Description   string
	AvailableData map[string]string
}

type Data interface {
	About | AstanaBuilding | KHCAllowedProject
}
