package main

import (

	"KuroiKitsu/go-ass"
	"github.com/n00bDooD/gosrt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/multimedia"
	"github.com/therecipe/qt/widgets"

	"os"
	"path"
	"sort"
)

type Subtitle struct {
	Start int64
	End   int64
	Text  string
}

type TranscriptSub []Subtitle

func (d TranscriptSub) Len() int           { return len(d) }
func (d TranscriptSub) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d TranscriptSub) Less(i, j int) bool { return d[i].Start < d[j].Start }

type TListState struct {
	Row         int
	Substart    int64
	Subend      int64
	Nxtsubstart int64
	Nxtsubend   int64
}

type TranscriptList struct {
	ListWidget  *widgets.QListWidget
	ListState   TListState
	Transcript  TranscriptSub
	MediaPlayer *multimedia.QMediaPlayer
}

func NewTranscriptList(parent widgets.QWidget_ITF, mp *multimedia.QMediaPlayer) *TranscriptList {

	var file = `C:\[Kamigami][Hotarubi_no_Mori_e].ass` //`C:\[Ringoshii][Hotarubi_no_Mori_e].srt`  //"C:\\Shidonia_No_Kishi_001.ass"
	tl := &TranscriptList{
		ListWidget:  widgets.NewQListWidget(parent),
		ListState:   TListState{},
		Transcript:  loadSubs(file),
		MediaPlayer: mp}

	sort.Sort(tl.Transcript)
	tl.ListState = TListState{0, tl.Transcript[0].Start, tl.Transcript[0].End, tl.Transcript[1].Start, tl.Transcript[1].End}

	for i := range tl.Transcript {
		tl.ListWidget.InsertItem2(i, tl.Transcript[i].Text)
	}

	var font = gui.NewQFont2("Meiryo", 14, 2, false)
	for i := 0; i < tl.ListWidget.Count(); i++ {
		tl.ListWidget.Item(i).SetFont(font)
	}

	var brush = gui.NewQBrush4(core.Qt__red, core.Qt__SolidPattern)
	tl.ListWidget.Item(0).SetBackground(brush)
	
	tl.ListWidget.ConnectKeyPressEvent(tl.keyPressEvent)
	tl.ListWidget.ConnectItemClicked(tl.itemClicked)


	return tl

}

func (tl *TranscriptList) keyPressEvent(event *gui.QKeyEvent) {

	switch int32(event.Key()) {

	case int32(core.Qt__Key_Space):
		if tl.MediaPlayer.State() == multimedia.QMediaPlayer__PausedState {
			tl.MediaPlayer.Play()
		} else {
			tl.MediaPlayer.Pause()
		}

	case int32(core.Qt__Key_Right):
		var time = tl.MediaPlayer.Position()
		time += 5000
		tl.MediaPlayer.SetPosition(time)

	case int32(core.Qt__Key_Left):
		var time = tl.MediaPlayer.Position()
		time -= 5000
		tl.MediaPlayer.SetPosition(time)

	}

}

func (tl *TranscriptList) itemClicked(item *widgets.QListWidgetItem) {

	var brush = gui.NewQBrush4(core.Qt__white, core.Qt__SolidPattern)
	tl.ListWidget.Item(tl.ListState.Row).SetBackground(brush)
	
	var row = tl.ListWidget.CurrentRow()

	row0 := tl.Transcript[row]
	row1 := tl.Transcript[row+1]

	tl.ListState = TListState{
		Row:         row,
		Substart:    row0.Start,
		Subend:      row0.End,
		Nxtsubstart: row1.Start,
		Nxtsubend:   row1.End }

	tl.MediaPlayer.SetPosition(tl.ListState.Substart)

	lkpline.SetText(row0.Text)
	SetVideoSubtitleWidget(row0.Text)
}

func loadSubs(file string) TranscriptSub {
	var subtext = TranscriptSub{}
	switch path.Ext(file) {
	case ".ass":
		subtext = loadASS(file)
	case ".srt":
		subtext = loadSRT(file)
	}
	return subtext
}

func loadASS(file string) TranscriptSub {
	subs, err := ass.ParseFile(file)

	if err != nil {
		panic(err)
	}

	var subtext = TranscriptSub{}
	st := new(Subtitle)

	if events := subs.Section("Events"); events != nil {

		for _, pair := range events.Pairs {
			if pair.Key != "Dialogue" {
				continue
			}

			rawText := pair.Get("Text")
			text := ass.Text(rawText).Readable()
			st.Text = text

			rawTime := pair.Get("Start")
			sta, _ := ass.ParseTime(rawTime)
			st.Start = int64(sta)

			rawTime = pair.Get("End")
			eta, _ := ass.ParseTime(rawTime)
			st.End = int64(eta)

			subtext = append(subtext, *st)

		}
	}

	return subtext

}

func loadSRT(file string) TranscriptSub {
	var f, err = os.Open(file)
	if err != nil {
		panic(err)
	}
	var subs = []gosrt.Subtitle{}

	gosrt.InputValidationStrictness = gosrt.SkipInvalid
	scanner := gosrt.NewScanner(f)

	for scanner.Scan() {
		subs = append(subs, scanner.Subtitle())
	}
	err = scanner.Err()

	var subtext = TranscriptSub{}
	st := new(Subtitle)

	for _, v := range subs {

		st.Text = v.Text
		st.Start = v.Start.Nanoseconds() / 1e6
		st.End = v.End.Nanoseconds() / 1e6
		subtext = append(subtext, *st)
	}

	return subtext

}
