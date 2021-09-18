package config

import (
	"github.com/gdamore/tcell/v2"
)

type Skin struct {
	ModalFgColor         tcell.Color
	MainFgColor          tcell.Color
	SecondaryFgColor     tcell.Color
	OptionalFgColor      tcell.Color
	FieldLabelColor      tcell.Color
	FieldValueColor      tcell.Color
	FieldColonColor      tcell.Color
	BackgroundColor      tcell.Color
	ModalBackgroundColor tcell.Color
	BorderColor          tcell.Color
	FgErrorColor         tcell.Color
	FgSuccessColor       tcell.Color
	TitleColor           tcell.Color
	TableBodyFgColor     tcell.Color
	TableHeaderFgColor   tcell.Color
	TableSelectedFgColor tcell.Color
	StatesColor          map[string]tcell.Color
}
type Skins map[string]Skin

func GetColor(hex string) tcell.Color {
	return tcell.GetColor(hex).TrueColor()
}
func GetSkin() Skin {
	skin := Skin{}
	skin.MainFgColor = GetColor("#86cdf9")
	skin.SecondaryFgColor = GetColor("#FD971F")
	skin.OptionalFgColor = GetColor("#FFFFFF")
	skin.FieldLabelColor = GetColor("#FD971F")
	skin.FieldValueColor = GetColor("#FFFFFF")
	skin.FieldColonColor = GetColor("#FD971F")
	skin.BackgroundColor = GetColor("#000000")
	skin.BorderColor = GetColor("#86cdf9")
	skin.FgErrorColor = GetColor("#F92672")
	skin.FgSuccessColor = GetColor("#F92672")
	skin.TitleColor = GetColor("#02fffe")
	skin.TableBodyFgColor = GetColor("#86cdf9")
	skin.TableHeaderFgColor = GetColor("#FFFFFF")
	skin.TableSelectedFgColor = GetColor("#04f2f2")
	skin.ModalBackgroundColor = GetColor("#5887a3")
	skin.ModalFgColor = GetColor("#FFFFFF")
	skin.StatesColor = map[string]tcell.Color{
		"cancel":     GetColor("#cacaca"),
		"close":      GetColor("#cacaca"),
		"refused":    GetColor("#cacaca"),
		"refuse":     GetColor("#cacaca"),
		"done":       GetColor("#0daa00"),
		"posted":     GetColor("#0daa00"),
		"paid":       GetColor("#0daa00"),
		"approved":   GetColor("#0daa00"),
		"validate":   GetColor("#0daa00"),
		"validated":  GetColor("#0daa00"),
		"to approve": GetColor("#02fffe"),
		"to confirm": GetColor("#02fffe"),
		"waiting":    GetColor("#02fffe"),
		"progress":   GetColor("#02fffe"),
		"to_close":   GetColor("#02fffe"),
		"ready":      GetColor("#02fffe"),
		"pending":    GetColor("#02fffe"),
	}
	return skin
}
