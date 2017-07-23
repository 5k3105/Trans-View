package main

import (
	"github.com/therecipe/qt/widgets"
)

func NewSessionSelect() *widgets.QWidget {
	var sessionSelect = widgets.NewQWidget(nil, 0)
	sessionSelect.SetLayout(ss_ui())
	return sessionSelect
}

func ss_ui() *widgets.QVBoxLayout { 
	
	var vlayoutPanel = widgets.NewQVBoxLayout()
	var comboSession = widgets.NewQComboBox(nil)
	//self.populatecombo("theme", self.comboTheme, self.ThemeDir, "Theme")
	//comboTheme.setCurrentIndex(self.comboTheme.findText(self.ThemeFile))
	//comboTheme.currentIndexChanged.connect(self.settheme)

	vlayoutPanel.AddWidget(comboSession, 0, 0)
	
	var comboVideo = widgets.NewQComboBox(nil)
	//self.populatecombo("theme", self.comboTheme, self.ThemeDir, "Theme")
	//comboTheme.setCurrentIndex(self.comboTheme.findText(self.ThemeFile))
	//comboTheme.currentIndexChanged.connect(self.settheme)

	vlayoutPanel.AddWidget(comboVideo, 0, 0)	

	var buttonSessionEdit = widgets.NewQPushButton(nil)
	buttonSessionEdit.SetText("Edit")
	//buttonSessionEdit.SetMaximumWidth(50)
	vlayoutPanel.AddWidget(buttonSessionEdit, 0, 1)
	//buttonSessionEdit.clicked.connect(self.createtheme)
	//vlayoutPanel.AddItem(widgets.NewQSpacerItem(1, 1, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum))
	
	var spacerPushUp = widgets.NewQSpacerItem(1, 1, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__MinimumExpanding)
	vlayoutPanel.AddItem(spacerPushUp)
	
	return vlayoutPanel
	
	}
/**/
