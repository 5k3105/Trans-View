package main

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/gui"
)

func NewSessionEdit() *widgets.QWidget {
	var sessionedit = widgets.NewQWidget(nil, 0)
	sessionedit.SetLayout(se_ui())
	return sessionedit
}

func se_ui() *widgets.QVBoxLayout { 
	
	var vlayoutPanel = widgets.NewQVBoxLayout()
	
	var hlayoutPanel1 = widgets.NewQHBoxLayout()
	
	var hlayoutPanel2 = widgets.NewQHBoxLayout()
	
	
	vlayoutPanel.AddLayout(hlayoutPanel1, 0)
	vlayoutPanel.AddLayout(hlayoutPanel2, 0)
	
	
	
	var vlayoutPanelENSD = widgets.NewQVBoxLayout()
	var hlayoutPanelENSD = widgets.NewQHBoxLayout()

	
	var txtSessionName = widgets.NewQLineEdit(nil)
	vlayoutPanelENSD.AddWidget(txtSessionName, 0, 0)

	hlayoutPanelENSD.AddItem(widgets.NewQSpacerItem(1, 1, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum))
			
	var btnSessionNew = widgets.NewQPushButton(nil)
	btnSessionNew.SetText("New")
	hlayoutPanelENSD.AddWidget(btnSessionNew, 0, 0)

	var btnSessionSave = widgets.NewQPushButton(nil)
	btnSessionSave.SetText("Save")
	hlayoutPanelENSD.AddWidget(btnSessionSave, 0, 0)
	
	var btnSessionDelete = widgets.NewQPushButton(nil)
	btnSessionDelete.SetText("Delete")
	hlayoutPanelENSD.AddWidget(btnSessionDelete, 0, 0)

	vlayoutPanelENSD.AddLayout(hlayoutPanelENSD, 0)
	hlayoutPanel1.AddLayout(vlayoutPanelENSD, 0)

	//var spacer1 = widgets.NewQSpacerItem(1, 1, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	//hlayoutPanel1.AddItem(spacer1)

	var vlayoutPanelENSD2 = widgets.NewQVBoxLayout()
	var hlayoutPanelENSD2 = widgets.NewQHBoxLayout()

	var cmbTheme = widgets.NewQComboBox(nil)
	//self.populatecombo("theme", self.comboTheme, self.ThemeDir, "Theme")
	//comboTheme.setCurrentIndex(self.comboTheme.findText(self.ThemeFile))
	//comboTheme.currentIndexChanged.connect(self.settheme)
	vlayoutPanelENSD2.AddWidget(cmbTheme, 0, 0)

	//var spacer2 = widgets.NewQSpacerItem(1, 1, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	hlayoutPanelENSD2.AddItem(widgets.NewQSpacerItem(1, 1, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum))
		
	var btnThemeNew = widgets.NewQPushButton(nil)
	btnThemeNew.SetText("New")
	hlayoutPanelENSD2.AddWidget(btnThemeNew, 0, 0)

	var btnThemeEdit = widgets.NewQPushButton(nil)
	btnThemeEdit.SetText("Edit")
	hlayoutPanelENSD2.AddWidget(btnThemeEdit, 0, 0)	
	
	vlayoutPanelENSD2.AddLayout(hlayoutPanelENSD2, 0)
	hlayoutPanel1.AddLayout(vlayoutPanelENSD2, 0)


	var listSessions = widgets.NewQListWidget(nil)
	
	hlayoutPanel2.AddWidget(listSessions, 0, 0)	
	
	
	var vlayoutPanelVideos = widgets.NewQVBoxLayout()
	var hlayoutPanelVideos = widgets.NewQHBoxLayout()
		
	var txtVideosPath = widgets.NewQLineEdit(nil)
	hlayoutPanelVideos.AddWidget(txtVideosPath, 0, 0)
		
	var btnVideosPath = widgets.NewQPushButton3(gui.NewQIcon5(":/icons/folder.png"), "", nil)
	//btnVideosPath.SetText("Select")
	hlayoutPanelVideos.AddWidget(btnVideosPath, 0, 0)	
	
	var listVideos = widgets.NewQListWidget(nil)
	
	vlayoutPanelVideos.AddLayout(hlayoutPanelVideos, 0)		
	vlayoutPanelVideos.AddWidget(listVideos, 0, 0)
	
	
	
	var vlayoutPanelSubs = widgets.NewQVBoxLayout()	
	var hlayoutPanelSubs = widgets.NewQHBoxLayout()
		
	var txtSubsPath = widgets.NewQLineEdit(nil)
	hlayoutPanelSubs.AddWidget(txtSubsPath, 0, 0)
		
	var btnSubsPath = widgets.NewQPushButton3(gui.NewQIcon5(":/icons/folder.png"), "", nil)
	//btnSubsPath.SetText("Select")
	hlayoutPanelSubs.AddWidget(btnSubsPath, 0, 0)	
	
	var listSubs = widgets.NewQListWidget(nil)
	
	vlayoutPanelSubs.AddLayout(hlayoutPanelSubs, 0)		
	vlayoutPanelSubs.AddWidget(listSubs, 0, 0)	
	
	hlayoutPanel2.AddLayout(vlayoutPanelVideos, 0)			
	hlayoutPanel2.AddLayout(vlayoutPanelSubs, 0)		

	var spacerPushUp = widgets.NewQSpacerItem(1, 1, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__MinimumExpanding)
	vlayoutPanel.AddItem(spacerPushUp)
		
	return vlayoutPanel
	}
/**/
