package models

const CodesForForgottenPwdTableName = "codes_for_forgotten_pwd"

type CodesForForgottenPwd struct {
	ID    int64  `db:"id" json:"id"`
	Code  string `db:"code" json:"code"`
	Email string `db:"email" json:"email"`
}

func (CodesForForgottenPwd) TableName() string {
	return CodesForForgottenPwdTableName
}
