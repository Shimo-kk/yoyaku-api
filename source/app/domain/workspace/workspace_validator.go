package workspace

import (
	"unicode/utf8"
	"yoyaku/app/core"
)

// 名称のバリデーション
func validateName(value string) error {
	if value == "" {
		return core.NewError(core.ValidationError, "ワークスペース名は空の値を入力することはできません。")
	}

	if utf8.RuneCountInString(value) > 50 {
		return core.NewError(core.ValidationError, "ワークスペース名は50文字より大きい値を入力できません。")
	}

	return nil
}
