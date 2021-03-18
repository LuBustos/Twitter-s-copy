package models

//Relationship clase para grabar la relacion del usuario con otro usuario
type Relationship struct {
	UserID             string `bson:"userid" json:"userId"`
	UserRelationshipID string `bson:"userrelationshipid" json:"userRelationshipId,omitempty"`
}
