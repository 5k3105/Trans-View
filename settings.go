package main

/*
type session struct {
	sessionName		string
	themeName		int
	assocSetId		int
	assocCurrent	int
	lastTime		int // ?
	}

type assocSet struct {
	id			int
	vid, sub	string
	}

type theme struct {
	themeName			string
	mwh, lh				int
	displayStyleSetId	int
	}

type displayStyleSet struct {
	id						int
	displayType				string
	fs						int
	ft, fc, bc				string
	}
*/



func loadSessions() {}

func saveSession() {}

func deleteSession() {}


func loadThemes() {}

func saveTheme() {}

func deleteTheme() {}


func loadAssocs() {}

func saveAssocs() {}

func deleteAssocSet() {}

func deleteAssocItem() {}








/*
import (
	"github.com/therecipe/qt/widgets"
)

func NewSettings() *widgets.QWidget {
	var settings = widgets.NewQWidget(nil, 0)
	settings.SetLayout(ui())
	return settings
}

func ui() *widgets.QVBoxLayout {

	var vlayout = widgets.NewQVBoxLayout()

	var hlayoutTheme = widgets.NewQHBoxLayout()     // theme
	var hlayoutPanel = widgets.NewQHBoxLayout()     // panel
	var hlayoutWinHt = widgets.NewQHBoxLayout()     // WinHt
	var hlayoutSaveReset = widgets.NewQHBoxLayout() // save reset

	//  # settings profile (load)
	var labelTheme = widgets.NewQLabel(nil, 0)
	labelTheme.SetText(" Theme: ")
	labelTheme.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed)

	hlayoutTheme.AddWidget(labelTheme, 0, 0)

	var comboTheme = widgets.NewQComboBox(nil)
	//self.populatecombo("theme", self.comboTheme, self.ThemeDir, "Theme")
	//comboTheme.setCurrentIndex(self.comboTheme.findText(self.ThemeFile))
	//comboTheme.currentIndexChanged.connect(self.settheme)

	hlayoutTheme.AddWidget(comboTheme, 0, 0)

	var buttonThemeFolder = widgets.NewQPushButton(nil)
	buttonThemeFolder.SetText("+")
	buttonThemeFolder.SetMaximumWidth(50)
	hlayoutTheme.AddWidget(buttonThemeFolder, 0, 0)
	//buttonThemeFolder.clicked.connect(self.showDialogThemeFolder)

	var buttonThemeCreate = widgets.NewQPushButton(nil)
	buttonThemeCreate.SetText("Create")
	buttonThemeCreate.SetMaximumWidth(50)
	hlayoutTheme.AddWidget(buttonThemeCreate, 0, 0)
	//buttonThemeCreate.clicked.connect(self.createtheme)

	var buttonThemeDelete = widgets.NewQPushButton(nil)
	buttonThemeDelete.SetText("Delete")
	buttonThemeDelete.SetMaximumWidth(50)
	hlayoutTheme.AddWidget(buttonThemeDelete, 0, 0)

	vlayout.AddLayout(hlayoutTheme, 0)

	//  # combo set panel
	var comboPanel = widgets.NewQComboBox(nil)
	//comboPanel.addItems(self.panel)
	//self.comboPanel.currentIndexChanged.connect(self.onPanelChanged)
	hlayoutPanel.AddWidget(comboPanel, 0, 0)

	//  # combo set font
	var comboFontFamily = widgets.NewQFontComboBox(nil)
	//self.comboFontFamily.currentFontChanged.connect(self.onFontFamilyChanged)
	hlayoutPanel.AddWidget(comboFontFamily, 0, 0)

	//  # spin font-size
	var spinFontSize = widgets.NewQSpinBox(nil)
	//spinFontSize.valueChanged.connect(self.onFontSizeChanged)
	hlayoutPanel.AddWidget(spinFontSize, 0, 0)

	//  # button fg
	var buttonColorFg = widgets.NewQPushButton(nil)
	//self.buttonColorFg.clicked.connect(self.onButtonColorFgClicked)
	buttonColorFg.SetText("Fg Color")
	buttonColorFg.SetMaximumWidth(50)
	hlayoutPanel.AddWidget(buttonColorFg, 0, 0)

	vlayout.AddLayout(hlayoutPanel, 0)

	//  # (max win ht, bg color) spacer label spin button
	var spacerMaxWinHt = widgets.NewQSpacerItem(40, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	hlayoutWinHt.AddItem(spacerMaxWinHt)

	var labelLineHt = widgets.NewQLabel(nil, 0)
	labelLineHt.SetText("Line Height: ")
	labelLineHt.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed)
	hlayoutWinHt.AddWidget(labelLineHt, 0, 0)
	labelLineHt.Hide()

	var spinLineHt = widgets.NewQSpinBox(nil)
	spinLineHt.SetMaximum(20)
	spinLineHt.SetMinimum(0)
	spinLineHt.SetValue(0)
	//spinLineHt.valueChanged.connect(self.onLineHtChanged)
	hlayoutWinHt.AddWidget(spinLineHt, 0, 0)
	spinLineHt.Hide()

	var labelMaxWinHt = widgets.NewQLabel(nil, 0)
	labelMaxWinHt.SetText("Max Win Height: ")
	labelMaxWinHt.SetSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed)
	hlayoutWinHt.AddWidget(labelMaxWinHt, 0, 0)

	var spinMaxWinHt = widgets.NewQSpinBox(nil)
	spinMaxWinHt.SetMaximum(500)
	spinMaxWinHt.SetMinimum(80)
	spinMaxWinHt.SetDisabled(true)
	spinMaxWinHt.SetValue(0)
	//spinMaxWinHt.valueChanged.connect(self.onWinHtChanged)
	hlayoutWinHt.AddWidget(spinMaxWinHt, 0, 0)

	var buttonColorBg = widgets.NewQPushButton(nil)
	//buttonColorBg.clicked.connect(self.onButtonColorBgClicked)
	buttonColorBg.SetText("Bg Color")
	buttonColorBg.SetMaximumWidth(50)
	hlayoutWinHt.AddWidget(buttonColorBg, 0, 0)

	vlayout.AddLayout(hlayoutWinHt, 0)

	//  # button ok cancel
	var spacerSaveCancel = widgets.NewQSpacerItem(20, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	hlayoutSaveReset.AddItem(spacerSaveCancel)

	var buttonSave = widgets.NewQPushButton(nil)
	var buttonReset = widgets.NewQPushButton(nil)
	//buttonSave.clicked.connect(self.savetheme)
	//buttonReset.clicked.connect(self.loadtheme)
	buttonSave.SetText("Save")
	buttonReset.SetText("Reset")

	hlayoutSaveReset.AddWidget(buttonSave, 0, 0)
	hlayoutSaveReset.AddWidget(buttonReset, 0, 0)

	vlayout.AddLayout(hlayoutSaveReset, 0)

	var spacerPushUp = widgets.NewQSpacerItem(1, 1, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__MinimumExpanding)
	vlayout.AddItem(spacerPushUp)

	return vlayout

}
*/



	
	
	
	
	
	
