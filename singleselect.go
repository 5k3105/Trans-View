package main

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/core"
	"io/ioutil"
	"strings"
)

var (

	comboVideo  	*widgets.QComboBox
	comboSub  		*widgets.QComboBox

	VideoTypeFilters = `mkv,avi,mp4,wmv,flv`
	SubTypeFilters = `srt,ass`

	VideoTypeFilter, SubTypeFilter	[]string

)



func init() {
	
	VideoTypeFilter = strings.Split(VideoTypeFilters,`,`)
	SubTypeFilter = strings.Split(SubTypeFilters,`,`)

	}



func NewSingleSelect() *widgets.QWidget {
	var singleSelect = widgets.NewQWidget(nil, 0)
	singleSelect.SetLayout(sss_ui())
	return singleSelect
}

func sss_ui() *widgets.QVBoxLayout { 
	
	var vlayoutPanel = widgets.NewQVBoxLayout()
	
	var hlayoutPanel1 = widgets.NewQHBoxLayout()
	var hlayoutPanel2 = widgets.NewQHBoxLayout()
	
	comboVideo = widgets.NewQComboBox(nil)
	var btnVideosPath = widgets.NewQPushButton3(gui.NewQIcon5(":/icons/folder.png"), "", nil)
	btnVideosPath.SetMaximumWidth(30)
	
	hlayoutPanel1.AddWidget(comboVideo, 0, 0)	
	hlayoutPanel1.AddWidget(btnVideosPath, 0, 0)	
	vlayoutPanel.AddLayout(hlayoutPanel1, 0)	
	

	comboSub = widgets.NewQComboBox(nil)
	var btnSubPath = widgets.NewQPushButton3(gui.NewQIcon5(":/icons/folder.png"), "", nil)
	btnSubPath.SetMaximumWidth(30)
	
	hlayoutPanel2.AddWidget(comboSub, 0, 0)	
	hlayoutPanel2.AddWidget(btnSubPath, 0, 0)	
	vlayoutPanel.AddLayout(hlayoutPanel2, 0)	
	
	var spacerPushUp = widgets.NewQSpacerItem(1, 1, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__MinimumExpanding)
	vlayoutPanel.AddItem(spacerPushUp)


	btnVideosPath.ConnectClicked(func(_ bool) { singleSelect_vDialog() })

	btnSubPath.ConnectClicked(func(_ bool) { singleSelect_sDialog() })







	return vlayoutPanel
	
	}
	
	
func singleSelect_vDialog() {
	
	fd := widgets.NewQFileDialog(window, core.Qt__Dialog)
	d := fd.GetExistingDirectory(window, `Open Directory`, core.QDir_Root().RootPath(), widgets.QFileDialog__DontUseNativeDialog)
	
	if d == `` { return }

	files, err := ioutil.ReadDir(d)
	if err != nil {
		println(err)
	}

	var s []string
	for _, file := range files {
		fn := file.Name()
		if extwithin(fn, VideoTypeFilter) {
			s = append(s, fn)
		}
	}	
	
	comboVideo.AddItems(s)

	}	
	
func singleSelect_sDialog() {
	
	fd := widgets.NewQFileDialog(window, core.Qt__Dialog)
	d := fd.GetExistingDirectory(window, `Open Directory`, core.QDir_Root().RootPath(), widgets.QFileDialog__DontUseNativeDialog)
	
	if d == `` { return }

	files, err := ioutil.ReadDir(d)
	if err != nil {
		println(err)
	}

	var s []string
	for _, file := range files {
		fn := file.Name()
		if extwithin(fn, SubTypeFilter) {
			s = append(s, fn)
		}
	}	
	
	comboSub.AddItems(s)	
	
	}	

func extwithin(s string, c []string) bool {
	
	st := strings.Split(s, `.`)
	s = st[len(st)-1]
	
	for _, str := range c {
		if str == s { return true }
	}
	return false
}

func within(s string, c []string) bool {
	for _, str := range c {
		if str == s { return true }
	}
	return false
}



	
/**/
	//self.populatecombo("theme", self.comboTheme, self.ThemeDir, "Theme")
	//comboTheme.setCurrentIndex(self.comboTheme.findText(self.ThemeFile))
	//comboTheme.currentIndexChanged.connect(self.settheme)	
	
	//self.populatecombo("theme", self.comboTheme, self.ThemeDir, "Theme")
	//comboTheme.setCurrentIndex(self.comboTheme.findText(self.ThemeFile))
	//comboTheme.currentIndexChanged.connect(self.settheme)
