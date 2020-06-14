package models

/*Relacion modelo para grabar la relacion de un usuario con otro */
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioID"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionID"`
}
