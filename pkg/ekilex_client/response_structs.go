package ekilex_client

import (
	"encoding/json"
	"time"
)

// NOTE: Some fields excluded for being out of scope for parent project, overly recursive, or similar.
type WordList struct {
	TotalCount int    `json:"totalCount"`
	Words      []Word `json:"words"`
}

type Word struct {
	ID            int    `json:"wordId"`
	Value         string `json:"wordValue"`
	ValuePrese    string `json:"wordValuePrese"`
	HomonymCount  int    `json:"homonymNr"`
	HomonymsExist bool   `json:"homonymsExist"` // WordDetails Synonyms
	Lang          string `json:"lang"`
	// LexemeRegisterCodes null 	`json:"lexemeRegisterCodes"` // WordDetails Synonyms
	// WordTypeCodes null `json:"wordTypeCodes"` // WordDetails Synonyms
	Prefixoid bool     `json:"prefixoid"`
	Suffixoid bool     `json:"suffixoid"`
	Foreign   bool     `json:"foreign"`
	Datasets  []string `json:"datasetCodes"` // WordList specific
	LexemeID  int      `json:"lexemeId"`     // WordDetails Synonyms specific
	// LexemeLevels	  null  `json:"lexemeLevels"`    // WordDetails Synonyms specific
	LexemesTagNames   []string  `json:"lexemesTagNames"` // WordDetails specific
	LastActivityEvent time.Time `json:"lastActivityEventOn"`
}

type WordDetails struct {
	Word Word `json:"word"`
	// wordTypes [] `json:"wordTypes"`
	Paradigms []Paradigm `json:"paradigms"`
	Lexemes   []Lexeme   `json:"lexemes"`
	// WordEtymology [] `json:"wordEtymology"`
	// OdWordRecommendations [] `json:"odWordRecommendations"`
	WordRelationDetails WordRelationDetails `json:"wordRelationDetails"`
	// FirstDefinitionValue null `json:"firstDefinitionValue"`
	ActiveTagComplete bool `json:"activeTagComplete"`
}

type Paradigm struct {
	ID               int         `json:"paradigmId"`
	Comment          string      `json:"comment"`
	InflectionType   json.Number `type:"integer" required:"true"`
	InflectionTypeNR json.Number `type:"integer" required:"true"`
	WordClass        string      `json:"wordClass"`
	Forms            []Form      `json:"forms"`
	FormsExist       bool        `json:"formsExist"`
}

type Form struct {
	ID          int      `json:"id"`
	Value       string   `json:"value"`
	ValuePrese  string   `json:"valuePrese"`
	Components  []string `json:"components"`
	DisplayForm string   `json:"displayForm"`
	MorphCode   string   `json:"morphCode"`
	MorphValue  string   `json:"morphValue"`
	// MorphFreq null `json:"morphFrequency"`
	// FormFrequency null `json:"formFrequency"`
}

type Lexeme struct {
	WordID           int    `json:"wordId"`
	Value            string `json:"wordValue"`
	ValuePrese       string `json:"wordValuePrese"`
	Lang             string `json:"wordLang"`
	HomonymNR        int    `json:"wordHomonymNr"`
	GenderCode       string `json:"wordGenderCode"`
	AspectCode       string `json:"wordAspectCode"`
	DisplayMorphCode string `json:"wordDisplayMorphCode"`
	// wordTypeCodes null `json:"wordTypeCodes"`
	Prefixoid      bool        `json:"wordPrefixoid"`
	Suffixoid      bool        `json:"wordSuffixoid"`
	Foreign        bool        `json:"wordForeign"`
	LexemeID       int         `json:"lexemeId"`
	MeaningID      int         `json:"meaningId"`
	DatasetName    string      `json:"datasetName"`
	DatasetCode    string      `json:"datasetCode"`
	Level1         int         `json:"level1"`
	Level2         int         `json:"level2"`
	LevelCount     json.Number `type:"integer" required:"true"`
	ValueStateCode string      `json:"valueStateCode"`
	// ValueState null `json:"valueState"`
	Tags       []string `json:"tags"`
	Complexity string   `json:"complexity"`
	Weight     float64  `json:"weight"`
	// Types null `json:"wordTypes"`
	Pos []NameCodeValue `json:"pos"` // TODO: Position or something else?
	// Derivs null `json:"derivs"`
	// Registers null `json:"registers"`
	// Governments [] `json:"governments"`
	// Grammars [] `json:"grammars"`
	Usages []Usage `json:"usages"`
	// Freeforms [] `json:"lexemeFreeforms"`
	// NoteLangGroups [] `json:"lexemeNoteLangGroups"`
	// Relations [] `json:"lexemeRelations"`
	// OdLexemeRecommendations [] `json:"odLexemeRecommendations"`
	CollationPosGroups    []CollationPosGroup `json:"collationPosGroups"`
	SecondaryCollocations []Collocation       `json:"secondaryCollocations"`
	// SourceLinks [] `json:"sourceLinks"`
	Meaning                         Meaning            `json:"meaning"`
	SynonymLangGroups               []SynonymLangGroup `json:"synonymLangGroups"`
	LexemeOrMeaningClassifiersExist bool               `json:"lexemeOrMeaningClassifiersExist"`
	Public                          bool               `json:"public"`
}

type NameCodeValue struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	Value string `json:"value"`
}

type Usage struct {
	ID         int    `json:"id"`
	Value      string `json:"value"`
	Lang       string `json:"lang"`
	Complexity string `json:"complexity"`
	OrderBy    int    `json:"orderBy"`
	TypeCode   string `json:"typeCode"`
	// TypeValue  null    `json:"typeValue"`
	// Translations [] `json:"translations"`
	// Definitions [] `json:"definitions"`
	// OdDefinitions [] `json:"odDefinitions"`
	// OdAlternatives [] `json:"odAlternatives"`
	// Authors [] `json:"authors"`
	// SourceLinks [] `json:"sourceLinks"`
	Public bool `json:"public"`
}

type CollationPosGroup struct {
	Code           string          `json:"code"`
	RelationGroups []RelationGroup `json:"relationGroups"`
}

type RelationGroup struct {
	Name         string        `json:"name"`
	Freq         float64       `json:"frequency"`
	Score        float64       `json:"score"`
	Collocations []Collocation `json:"collocations"`
}

type Collocation struct {
	Value      string              `json:"value"`
	Definition string              `json:"definition"`
	Freq       float64             `json:"frequency"`
	Score      float64             `json:"score"`
	Usages     []string            `json:"collocUsages"`
	Members    []CollocationMember `json:"collocMembers"`
}

type CollocationMember struct {
	ID     int     `json:"id"`
	Value  string  `json:"value"`
	Weight float64 `json:"weight"`
}

type Meaning struct {
	ID int `json:"id"`
	// LexemeIDs   null  `json:"lexemeIds"`
	Definitions          []Definition          `json:"definitions"`
	DefinitionLangGroups []DefinitionLangGroup `json:"definitionLangGroups"`
	// LexemeLangGroups [] `json:"lexemeLangGroups"`
	// Domains [] `json:"domains"`
	SemanticTypes []NameCodeValue `json:"semanticTypes"`
	Freeforms     []Freeform      `json:"freeforms"`
	// LearnerComments [] `json:"learnerComments"`
	// Images [] `json:"images"`
	// Medias [] `json:"medias"`
	// NoteLangGroups [] `json:"noteLangGroups"`
	Relations     []Relation   `json:"relations"`
	ViewRelations [][]Relation `json:"viewRelations"`
	// SynonymLangGroups null `json:"synonymLangGroups"`
	ActiveTagComplete   bool      `json:"activeTagComplete"`
	LastActicivityEvent time.Time `json:"lastActivityEventOn"`
	LastApproveEventOn  time.Time `json:"lastApproveEventOn"`
}

type Definition struct {
	ID         int    `json:"id"`
	Value      string `json:"value"`
	Lang       string `json:"lang"`
	Complexity string `json:"complexity"`
	OrderBy    int    `json:"orderBy"`
	TypeCode   string `json:"typeCode"`
	// TypeValue  ?    `json:"typeValue"`
	DatasetCodes []string `json:"datasetCodes"`
	Notes        []string `json:"notes"`
	// SourceLinks [] `json:"sourceLinks"`
	Public bool `json:"public"`
}

type DefinitionLangGroup struct {
	Lang        string       `json:"lang"`
	Selected    bool         `json:"selected"`
	Definitions []Definition `json:"definitions"`
}

type Freeform struct {
	ID         int    `json:"id"`
	ParentID   int    `json:"parentId"`
	Type       string `json:"type"`
	Value      string `json:"value"`
	ValuePrese string `json:"valuePrese"`
	// ValueDate null `json:"valueDate"`
	Lang       string `json:"lang"`
	Complexity string `json:"complexity"`
	OrderBy    int    `json:"orderBy"`
	// Children [] `json:"children"`
	Public bool `json:"public"`
}

type Relation struct {
	ID         int    `json:"id"`
	LexemeID   int    `json:"lexemeId"`
	MeaningID  int    `json:"meaningId"`
	WordID     int    `json:"wordId"`
	Value      string `json:"wordValue"`
	ValuePrese string `json:"wordValuePrese"`
	Lang       string `json:"lang"`
	AspectCode string `json:"aspectCode"`
	// AspectTypeCodes [] `json:"aspectTypeCodes"`
	// WordTypeCodes [] `json:"wordTypeCodes"`
	Prefixoid     bool   `json:"prefixoid"`
	Suffixoid     bool   `json:"suffixoid"`
	Foreign       bool   `json:"foreign"`
	HomonymNR     int    `json:"homonymNr"`
	HomonymsExist bool   `json:"homonymsExist"`
	RelTypeCode   string `json:"relTypeCode"`
	RelTypeLabel  string `json:"relTypeLabel"`
	OrderBy       int    `json:"orderBy"`
	// LexemeValueStateCode null `json:"lexemeValueStateCode"`
	// LexemeRegisterCodes [] `json:"lexemeRegisterCodes"`
	// LexemeGovernmentValues [] `json:"lexemeGovernmentValues"`
	Levels       int      `json:"levels"`
	DatasetCodes []string `json:"datasetCodes"`
	Weight       float64  `json:"weight"`
}

type SynonymLangGroup struct {
	Lang     string    `json:"lang"`
	Selected bool      `json:"selected"`
	Synonyms []Synonym `json:"synonyms"`
}

type Synonym struct {
	Type       string `json:"type"`
	MeaningID  int    `json:"meaningId"`
	RelationID int    `json:"relationId"`
	Words      []Word
	Lang       string  `json:"lang"`
	Weight     float64 `json:"weight"`
	OrderBy    int     `json:"orderBy"`
}

type WordRelationDetails struct {
	// WordSynRelations null `json:"wordSynRelations"`
	PrimaryWordRelationGroups   []WordRelationGroup `json:"primaryWordRelationGroups"`
	SecondaryWordRelationGroups []WordRelationGroup `json:"secondaryWordRelationGroups"`
	// WordGroups [] `json:"wordGroups"`
	GroupRelationExists bool `json:"groupRelationExists"`
}

type WordRelationGroup struct {
	ID        int    `json:"id"`
	TypeCode  string `json:"typeCode"`
	TypeLabel string `json:"typeLabel"`
	// Members null `json:"members"`
}
