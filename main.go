package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

// to build db:
// jmdict_parse2, jmdict_parse_te

var (
	window			*widgets.QMainWindow

	singleSelect	*widgets.QWidget
	sessionSelect	*widgets.QWidget
	sessionEdit		*widgets.QWidget
	//themeEdit		*widgets.QWidget
	
	player		*VPlayer
	transcript	*TranscriptList
	vocabdoc	*VocabDoc
	lkpline		*LookupLine
	linedefs	*widgets.QTextBrowser
	sublbl		*widgets.QLabel
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	window = widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Trans-Viewer v0.1")

	// statusbar
	var statusbar = widgets.NewQStatusBar(window)
	window.SetStatusBar(statusbar)

	// docksingleSelect
	var dsingleSelect = widgets.NewQDockWidget("Single Select", window, 0)
	window.AddDockWidget(core.Qt__LeftDockWidgetArea, dsingleSelect)
	singleSelect = NewSingleSelect()
	dsingleSelect.SetWidget(singleSelect)
	dsingleSelect.SetMaximumHeight(100)
	
	// docksessionSelect
	var dsessionSelect = widgets.NewQDockWidget("Session Select", window, 0)
	window.AddDockWidget(core.Qt__LeftDockWidgetArea, dsessionSelect)
	sessionSelect = NewSessionSelect()
	dsessionSelect.SetWidget(sessionSelect)
	dsessionSelect.SetMaximumHeight(120)
	
	// docksessionEdit
	var dsessionEdit = widgets.NewQDockWidget("Session Edit", window, 0)
	window.AddDockWidget(core.Qt__LeftDockWidgetArea, dsessionEdit)
	sessionEdit = NewSessionEdit()
	dsessionEdit.SetWidget(sessionEdit)
	dsessionEdit.SetMaximumHeight(330)


	// dockvideo
	var dvideo = widgets.NewQDockWidget("Video Player", window, 0)
	dvideo.SetMinimumWidth(300)
	window.AddDockWidget(core.Qt__LeftDockWidgetArea, dvideo)

	player = NewVideoPlayer()
	dvideo.SetWidget(player)

	sublbl = NewVideoSubtitleWidget(player.Video)

	transcript = NewTranscriptList(window, player.Media)
	window.SetCentralWidget(transcript.ListWidget)

	var dvocab = widgets.NewQDockWidget("Vocabulary", window, 0)
	window.AddDockWidget(core.Qt__RightDockWidgetArea, dvocab)
	vocabdoc = NewVocabDocument(window)
	dvocab.SetWidget(vocabdoc)

	//var dkanji = widgets.NewQDockWidget("Kanji", window, 0)
	//window.AddDockWidget(core.Qt__RightDockWidgetArea, dkanji)
	//dkanji.Hide()

	linedefs = widgets.NewQTextBrowser(window)
	var dlinedefs = widgets.NewQDockWidget("Saved Definitions", window, 0)
	dlinedefs.SetWidget(linedefs)
	dlinedefs.SetMinimumWidth(250)


	lkpline = NewLookupLine(window)
	var dlookupline = widgets.NewQDockWidget("Transcript Line Lookup", window, 0)
	dlookupline.SetWidget(lkpline)
	dlookupline.SetMaximumHeight(90)

	window.AddDockWidget(core.Qt__RightDockWidgetArea, dlookupline)
	window.AddDockWidget(core.Qt__RightDockWidgetArea, dlinedefs)

	widgets.QApplication_SetStyle2("fusion")
	window.ShowMaximized()
	widgets.QApplication_Exec()
}



