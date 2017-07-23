package main

import (
	"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/multimedia"
	"github.com/therecipe/qt/widgets"
)

type VPlayer struct {
	*widgets.QWidget
	Media           *multimedia.QMediaPlayer
	Video           *multimedia.QVideoWidget
	openButton      *widgets.QPushButton
	playButton      *widgets.QPushButton
	positionsSlider *widgets.QSlider
	errorLabel      *widgets.QLabel
}

func NewVideoPlayer() *VPlayer {

	v := &VPlayer{}
	v.QWidget = widgets.NewQWidget(nil, 0)
	v.Media = multimedia.NewQMediaPlayer(nil, multimedia.QMediaPlayer__VideoSurface)
	v.Video = multimedia.NewQVideoWidget(nil)
	v.openButton = widgets.NewQPushButton2("Open", nil)
	v.openButton.ConnectClicked(func(_ bool) { v.openFile() })

	v.playButton = widgets.NewQPushButton(nil)
	v.playButton.SetEnabled(false)
	v.playButton.SetIcon(v.Style().StandardIcon(widgets.QStyle__SP_MediaPlay, nil, nil))
	v.playButton.ConnectClicked(func(_ bool) { v.play() })

	v.positionsSlider = widgets.NewQSlider2(core.Qt__Horizontal, nil)
	v.positionsSlider.SetRange(0, 0)
	v.positionsSlider.ConnectSliderMoved(v.setPosition)

	v.errorLabel = widgets.NewQLabel(nil, 0)
	v.errorLabel.SetTextFormat(core.Qt__RichText)
	v.errorLabel.SetSizePolicy2(widgets.QSizePolicy__Preferred, widgets.QSizePolicy__Maximum)

	controlLayout := widgets.NewQHBoxLayout()
	controlLayout.SetContentsMargins(0, 0, 0, 0)
	controlLayout.AddWidget(v.openButton, 0, 0)
	controlLayout.AddWidget(v.playButton, 0, 0)
	controlLayout.AddWidget(v.positionsSlider, 0, 0)

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(v.Video, 0, 0)
	layout.AddLayout(controlLayout, 0)
	layout.AddWidget(v.errorLabel, 0, 0)

	v.SetLayout(layout)

	v.Media.SetVideoOutput(v.Video)
	v.Media.ConnectStateChanged(v.mediaStateChanged)
	v.Media.ConnectPositionChanged(v.positionChanged)
	v.Media.ConnectDurationChanged(v.durationChanged)
	v.Media.ConnectError2(v.handleError)

	v.Media.SetNotifyInterval(50)

	v.Video.ConnectMouseDoubleClickEvent(v.mouseDoubleClickEvent)
	v.Video.ConnectKeyPressEvent(v.keyPressEvent)

	return v
}

func (v *VPlayer) mouseDoubleClickEvent(e *gui.QMouseEvent) {
	v.toggleFullScreen()
}

func (v *VPlayer) toggleFullScreen() {
	if v.Video.IsFullScreen() {
		v.Video.SetFullScreen(false)
	} else {
		v.Video.SetFullScreen(true)
	}
}

func (v *VPlayer) keyPressEvent(e *gui.QKeyEvent) {
	switch int32(e.Key()) {

	case int32(core.Qt__Key_Escape):
		v.toggleFullScreen()

	case int32(core.Qt__Key_Space):
		if v.Media.State() == multimedia.QMediaPlayer__PausedState {
			v.Media.Play()
		} else {
			v.Media.Pause()
		}

	case int32(core.Qt__Key_Right):
		var time = v.Media.Position()
		time += 5000
		v.Media.SetPosition(time)

	case int32(core.Qt__Key_Left):
		var time = v.Media.Position()
		time -= 5000
		v.Media.SetPosition(time)

	}

}

func (v *VPlayer) openFile() {
	v.errorLabel.SetText("")

	var fileName = widgets.QFileDialog_GetOpenFileName(nil, "Open Movie", core.QDir_HomePath(), "", "", 0)
	if fileName != "" {
		v.Media.SetMedia(multimedia.NewQMediaContent2(core.QUrl_FromLocalFile(fileName)), nil)
		v.playButton.SetEnabled(true)
	}
}

func (v *VPlayer) play() {
	switch v.Media.State() {
	case multimedia.QMediaPlayer__PlayingState:
		{
			v.Media.Pause()
		}

	default:
		{
			v.Media.Play()
		}
	}
}

func (v *VPlayer) mediaStateChanged(state multimedia.QMediaPlayer__State) {
	switch state {
	case multimedia.QMediaPlayer__PlayingState:
		{
			v.playButton.SetIcon(v.Video.Style().StandardIcon(widgets.QStyle__SP_MediaPause, nil, nil))
		}

	default:
		{
			v.playButton.SetIcon(v.Video.Style().StandardIcon(widgets.QStyle__SP_MediaPlay, nil, nil))
		}
	}
}

func (v *VPlayer) positionChanged(position int64) {
	v.positionsSlider.SetValue(int(position))

	var time = v.Media.Position()

	ls := transcript.ListState
	lw := transcript.ListWidget

	// if time < ls.Subend //moving back

	if time > ls.Subend {
		var brush = gui.NewQBrush4(core.Qt__gray, core.Qt__SolidPattern)
		lw.Item(ls.Row).SetBackground(brush)

		if time > ls.Nxtsubstart {

			row := ls.Row + 1
			row0 := transcript.Transcript[row]
			row1 := transcript.Transcript[row+1]
			transcript.ListState = TListState{
				Row:         row,
				Substart:    row0.Start,
				Subend:      row0.End,
				Nxtsubstart: row1.Start,
				Nxtsubend:   row1.End}

			brush = gui.NewQBrush4(core.Qt__red, core.Qt__SolidPattern)
			lw.Item(row).SetBackground(brush)
			brush = gui.NewQBrush4(core.Qt__white, core.Qt__SolidPattern)
			lw.Item(row - 1).SetBackground(brush)

			// this doesnt work:
			transcript.ListWidget.ScrollToItem(transcript.ListWidget.Item(transcript.ListWidget.CurrentRow()), widgets.QAbstractItemView__PositionAtCenter) //tl.ListWidget.ScrollHint)
			//v.TL.ListWidget.ScrollToItem(v.TL.ListWidget.Item(v.TL.ListState.Row), widgets.QAbstractItemView__EnsureVisible)
			
			//            # Update LineDefs panel
			//            LineDefs.lookup(subsList.currentRow)

			lkpline.SetText(row0.Text)


		}
	}
}

func (v *VPlayer) durationChanged(duration int64) {
	v.positionsSlider.SetRange(0, int(duration))
}

func (v *VPlayer) setPosition(position int) {
	v.Media.SetPosition(int64(position))
}

func (v *VPlayer) handleError(err multimedia.QMediaPlayer__Error) {
	v.playButton.SetEnabled(false)

	var errString = v.Media.ErrorString()

	if errString == "" {
		switch err {
		case multimedia.QMediaPlayer__NoError:
			{
				errString = "QMediaPlayer::NoError"
			}

		case multimedia.QMediaPlayer__ResourceError:
			{
				errString = "QMediaPlayer::ResourceError"
			}

		case multimedia.QMediaPlayer__FormatError:
			{
				errString = "QMediaPlayer::FormatError"
			}

		case multimedia.QMediaPlayer__NetworkError:
			{
				errString = "QMediaPlayer::NetworkError"
			}

		case multimedia.QMediaPlayer__AccessDeniedError:
			{
				errString = "QMediaPlayer::AccessDeniedError"
			}

		case multimedia.QMediaPlayer__ServiceMissingError:
			{
				errString = "QMediaPlayer::ServiceMissingError"
			}
		}
	}

	v.errorLabel.SetText(fmt.Sprintf("File: %v<br>Error: %v<br>Supported MIME-Types: %v %v %v %v %v %v",
		v.Media.CurrentMedia().CanonicalUrl().ToString(0),
		errString,
		hasSupport("video/avi"),
		hasSupport("video/mp4"),
		hasSupport("video/x-flv"),
		hasSupport("video/3gpp"),
		hasSupport("video/x-ms-wmv"),
		hasSupport("video/x-matroska")))
}

func hasSupport(mimeType string) string {
	if multimedia.QMediaPlayer_HasSupport(mimeType, make([]string, 0), 0) >= multimedia.QMultimedia__MaybeSupported {
		return mimeType + "=<font color=\"green\">true</font>"
	}
	return mimeType + "=<font color=\"red\">false</font>"
}
