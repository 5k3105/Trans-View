package main

import(

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/multimedia"
	"github.com/therecipe/qt/widgets"
)

//*widgets.QWidget
func NewVideoSubtitleWidget(parent *multimedia.QVideoWidget) *widgets.QLabel {
	
	lbl := widgets.NewQLabel(parent, core.Qt__Widget)
	lbl.SetStyleSheet(`background-color: rgba(255, 150, 150, 150);`)
	lbl.SetText("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	
	return lbl
}

func SetVideoSubtitleWidget(txt string) {

	font := gui.NewQFont2("Meiryo", 11, 1, false)
	
	rect := core.NewQRect4(0, 0, 500, 200)
	ivsub := gui.NewQImage2(rect.Size(), gui.QImage__Format_ARGB32_Premultiplied)	
	painter := gui.NewQPainter()
	
	painter.Begin(ivsub)

	painter.FillRect8(rect, core.Qt__white)
	painter.SetRenderHints(gui.QPainter__Antialiasing | gui.QPainter__TextAntialiasing | gui.QPainter__SmoothPixmapTransform, true)

	painter.SetFont(font)
	painter.SetPen2(gui.NewQColor2(core.Qt__black))
	painter.DrawText6(rect, 0x0004 | 0x0040, txt, rect)
//core.Qt__TextWordWrap |core.Qt__AlignHCenter | core.Qt__AlignBottom
	painter.End()
	
	//pixmap := widgets.NewQPixmap()
	
	sublbl.SetPixmap(gui.QPixmap_FromImage(ivsub, core.Qt__AutoColor))

}
