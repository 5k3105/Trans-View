package main

import (
	"github.com/couchbaselabs/vellum"

	"golang.org/x/exp/utf8string"	
	"unicode/utf8"
	"os"
	"fmt"
	"bufio"
	"io/ioutil"
	"strings"
	"strconv"
	
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/core"
)

type LookupLine struct {
	*widgets.QTextBrowser
	scanPosition int
}

var fst *vellum.FST
var dictmap []string
var dict	map[int]string

func init() {

	// load dict
	dict = make(map[int]string)
	readdict(`dict/dict_lines.txt`)


	var err error
	// load fst
	fst, err = vellum.Open(`dict/conjtxt.xyz`)
	if err != nil {
		println(err)
	}
	
	// load dict map
	dictmap = readlines(`dict/txte.txt`)

}

func readlines(f string) []string {

	file, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
	}
	
	reader := bufio.NewReader(file)
	contents, _ := ioutil.ReadAll(reader)

	lf := string([]byte{13}) //,10
	
	lines := strings.Split(string(contents), lf)
	file.Close()

	return lines

}

func readdict(f string)  {

	file, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
	}
	
	reader := bufio.NewReader(file)
	contents, _ := ioutil.ReadAll(reader)

	lf := string([]byte{13}) //,10
	
	lines := strings.Split(string(contents), lf)
	file.Close()

	println("dictlen",len(lines))

	
	for _, v := range lines {
		
		ed := strings.Split(v, `,`)
		
		i,_ := strconv.Atoi(ed[0])
		
		if len(ed) < 2 { continue }
		
		dict[i] = ed[1]
		if i == 1357490 { println("my name is", ed[1])}
		
		}

}

func NewLookupLine(parent widgets.QWidget_ITF) *LookupLine {

	ll := &LookupLine{
		QTextBrowser: widgets.NewQTextBrowser(parent),
		}

	ll.SetMouseTracking(true)
	ll.SetReadOnly(true)
	
	var font = gui.NewQFont2("Meiryo", 14, 2, false)
	ll.SetFont(font)

	ll.ConnectMouseMoveEvent(ll.MouseMove)
	ll.ConnectMousePressEvent(ll.MousePress)

	return ll

}

func (ll *LookupLine) MouseMove(event *gui.QMouseEvent) {
	ll.updateSampleMouseEvent(event)
	}

func (ll *LookupLine) MousePress(event *gui.QMouseEvent) {
	ll.updateSampleMouseEvent(event)	
	}


func (ll *LookupLine) updateSampleMouseEvent(event *gui.QMouseEvent) {
	cursor := ll.CursorForPosition(event.Pos())
	
	if ll.scanPosition == cursor.Position() { return }
	
    ll.scanPosition = cursor.Position()	
	
	if event.Modifiers() == core.Qt__ShiftModifier {
		ll.updateSampleFromPosition()
		}

	}

func (ll *LookupLine) updateSampleFromPosition() {
	
	samplePosStart := ll.scanPosition

	cursor := ll.TextCursor()
	content := ll.ToPlainText()

	sampleLength := 20
	cs := utf8string.NewString(content)

	if cs.RuneCount() < samplePosStart + sampleLength { sampleLength = cs.RuneCount() - samplePosStart }

	println("content             ", content)
	println("content len         ", cs.RuneCount())
	println("samplePosStart      ", samplePosStart)
	println("sampleLength        ", sampleLength)


	var result uint64 
	var contentSample string

	for i := sampleLength; i > 0 ; i-- {
		
		contentSample = cs.Slice(samplePosStart,samplePosStart+i)

		val, exists, err := fst.Get([]byte(contentSample))
		result = val
		
		if err != nil { println(err) }
		if !exists { println("not exists") }

		println("contentSample    ", contentSample)
		
		if exists { break }
		
		}

//contentSample = `こちら`
//res := dict.Search(contentSample, 10)

	if result == 0 { return }
	
	cursor.SetPosition(samplePosStart, gui.QTextCursor__MoveAnchor)
	cursor.SetPosition(samplePosStart + utf8.RuneCountInString(contentSample), gui.QTextCursor__KeepAnchor)

	ll.SetTextCursor(cursor)

	vocabdoc.BuildVocabDefs(contentSample, result)


	}


/*
    def updateSampleMouseEvent(self, event):
        cursor = self.cursorForPosition(event.pos())
        self.state.scanPosition = cursor.position()
        if event.buttons() & QtCore.Qt.MidButton or event.modifiers() & QtCore.Qt.ShiftModifier:
            self.updateSampleFromPosition()







    def updateSampleFromPosition(self):
        samplePosStart = self.state.scanPosition
        samplePosEnd = self.state.scanPosition + 20 #self.preferences['scanLength']

        cursor = self.textCursor()
        content = unicode(self.toPlainText())
        contentSample = content[samplePosStart:samplePosEnd]
        contentSampleFlat = contentSample.replace(u'\n', unicode())

        if len(contentSampleFlat) == 0 or not japanese.util.isJapanese(contentSampleFlat[0]):
            cursor.clearSelection()
            self.setTextCursor(cursor)
            return

        lengthMatched = 0
        if self.dockVocab.isVisible():
            self.state.vocabDefs, lengthMatched = self.language.findTerm(contentSampleFlat)
            sentence = reader_util.findSentence(content, samplePosStart)
            for definition in self.state.vocabDefs:
                definition['sentence'] = sentence
            self.updateVocabDefs()

        if self.dockKanji.isVisible():
            if lengthMatched == 0:
                self.state.kanjiDefs = self.language.findCharacters(contentSampleFlat[0])
                if len(self.state.kanjiDefs) > 0:
                    lengthMatched = 1
            else:
                self.state.kanjiDefs = self.language.findCharacters(contentSampleFlat[:lengthMatched])
            self.updateKanjiDefs()

        lengthSelect = 0
        for c in contentSample:
            if lengthMatched <= 0:
                break
            lengthSelect += 1
            if c != u'\n':
                lengthMatched -= 1

        cursor.setPosition(samplePosStart, QtGui.QTextCursor.MoveAnchor)
        cursor.setPosition(samplePosStart + lengthSelect, QtGui.QTextCursor.KeepAnchor)
        self.setTextCursor(cursor)






    def findTerm(self, word, wildcards=False):
        self.requireIndex('Terms', 'expression')
        self.requireIndex('Terms', 'reading')

        cursor = self.db.cursor()
        cursor.execute('SELECT * FROM Terms WHERE expression {0} ? OR reading=? LIMIT 100'.format('LIKE' if wildcards else '='), (word, word))

        results = list()
        for expression, reading, glossary, tags in cursor.fetchall():
            results.append({
                'expression': expression,
                'reading': reading,
                'glossary': glossary,
                'tags': tags.split()
            })

        return results






*/

	//cursor.SetPosition(samplePosStart + len(contentSample), gui.QTextCursor__KeepAnchor)	
	//cursor.SetPosition(samplePosStart + len(cs.Slice(samplePosStart,len(contentSample))), gui.QTextCursor__KeepAnchor)
	//cursor.SetPosition(len(contentSample), gui.QTextCursor__KeepAnchor)
	
	
	//println("sample pos start  : ", len(cs.Slice(0,samplePosStart)))
	//println("           end    : ", len(cs.Slice(0,samplePosStart)) + len(contentSample))

	//println("sample pos start  : ", samplePosStart)
	//println("           end    : ", samplePosStart + len(contentSample))
	//println("           end    : ", samplePosStart + len(cs.Slice(samplePosStart,len(contentSample))))
	//println("           end    : ", len(contentSample))

	//cursor.SetPosition(len(cs.Slice(0,samplePosStart)), gui.QTextCursor__MoveAnchor)
	//cursor.SetPosition(len(cs.Slice(0,samplePosStart)) + len(contentSample), gui.QTextCursor__KeepAnchor)

