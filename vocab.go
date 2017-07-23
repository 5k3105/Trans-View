package main

import (
	"strings"
	"strconv"
	"bytes"
	"text/template"
	//"github.com/gojp/nihongo/lib/dictionary"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type css struct {
	Fontsize   int
	Fontfamily string
	Color      string
	Lineheight int
}

type span struct {
	Top			css
	Expression 	css
	Reading    	css
	Glossary   	css
}

var theme = span{
	Top: 		css{20, "Meiryo", "black", 12},
	Expression: css{18, "Meiryo", "red", 10},
	Reading:    css{18, "Meiryo", "green", 10},
	Glossary:   css{18, "Meiryo", "blue", 10}}

type Entry struct {
	Word       string `json:"word"`
	Furigana   string `json:"furigana"`
	Definition string `json:"definition"`
	Common     bool   `json:"common,omitempty"`
}

type VocabDoc struct {
	*widgets.QTextBrowser
	vocabDefs	[]string //entry ?
	}


func NewVocabDocument(parent widgets.QWidget_ITF) *VocabDoc {

	var vd = &VocabDoc{
		QTextBrowser: widgets.NewQTextBrowser(parent),
	}

	var txt = buildDefHeader()
	vd.SetHtml(txt)
	vd.SetAcceptDrops(false)
	vd.SetOpenLinks(false)

	var font = gui.NewQFont2("Meiryo", 14, 2, false)
	vd.SetFont(font)
	
	vd.ConnectAnchorClicked(vd.anchorclicked)

	return vd

}

func (vd *VocabDoc) anchorclicked(link *core.QUrl) {

	splt := strings.Split(link.ToString(core.QUrl__None),`:`)
	command  := splt[0]	
	index,_ := strconv.Atoi(splt[1])
	println("anchorclicked",command, index)
	vd.executeVocabCommand(command, index)
	
	}

func(vd *VocabDoc) executeVocabCommand(cmd string, idx int) {
	
	//if idx >= len(vd.vocabDefs) { return }
	
	def := vd.vocabDefs[idx]
	
	if cmd == `copyvocabdef` {
		
		html := buildDefHeader() + def + buildDefFooter()
		
		linedefs.SetHtml(html)

		}

	}

func buildDefHeader() string {
	var txt = ` <html><head><style>
	body { background-color: white }
	span.top { font-size: {{.Top.Fontsize}}px; font-family: '{{.Top.Fontfamily}}'; color: {{.Top.Color}}; line-height: {{.Top.Lineheight}}px }
	span.expression { font-size: {{.Expression.Fontsize}}px; font-family: '{{.Expression.Fontfamily}}'; color: {{.Expression.Color}}; line-height: {{.Expression.Lineheight}}px }
	span.reading { font-size: {{.Reading.Fontsize}}px; font-family: '{{.Reading.Fontfamily}}'; color: {{.Reading.Color}}; line-height: {{.Reading.Lineheight}}px }
	span.glossary { font-size: {{.Glossary.Fontsize}}px; font-family: '{{.Glossary.Fontfamily}}'; color: {{.Glossary.Color}}; line-height: {{.Glossary.Lineheight}}px }
	</style></head><body>`

	buf := new(bytes.Buffer)
	tmpl, err := template.New("txt").Parse(txt)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(buf, theme)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func buildDefFooter() string { return "</body></html>" }



func buildEntry(e string) Entry {


	println("entry is: ", e)

	entry := strings.Split(e, `|`)
	println("entry[0] >> ", entry[0])
	k,_ := strconv.Atoi(entry[0][0:1])
	r,_ := strconv.Atoi(entry[0][1:2])
	p,_ := strconv.Atoi(entry[0][2:3])
	g,_ := strconv.Atoi(entry[0][3:4])		
	if len(entry[0]) == 5 { g,_ = strconv.Atoi(entry[0][3:5]) } 
	

	var ks,rs,ps,gs string

	for i, v := range entry {
		
		if i !=0 {
			
			if k != 0 {
				
				ks += `; ` + v
				k -= 1
				continue
				}
			
			if k == 0 && r != 0 {
				
				rs += `; ` + v
				r -= 1		
				continue
				}
			
			if k == 0 && r == 0 && p != 0 {
				
				ps += `; ` + v
				p -= 1		
				continue
				}		
			
			if k == 0 && r == 0 && p == 0 && g !=0 {
				
				gs += `; ` + v
				g -= 1		
				continue
				}					
			}
		}
		
		if len(ks) > 2 { ks = ks[2:] } 
		if len(rs) > 2 { rs = rs[2:] } 
		if len(ps) > 2 { ps = ps[2:] } 
		if len(gs) > 2 { gs = gs[2:] } 		
	
	return Entry{
			Word:       ks,
			Furigana:   rs,
			Definition: ps + ` | ` + gs,
			//Common:     ps,
		}

}



func (vd *VocabDoc) BuildVocabDefs(contentSample string, result uint64 ) {

	entries := []Entry{}

	println("result: ", result)
	

	if result >= 100000000 {
	
		result = result - 100000000
		de := dictmap[result]
		println("dictmap= ", de)
		des := strings.Split(de, `|`)
		
		for _, e := range des { 
			
			re,_ := strconv.Atoi(e)
			entries = append(entries, buildEntry(dict[re])) 
			
			}

	} else {
		
		entries = append(entries, buildEntry(dict[int(result)]))
	}

	var html string
	vd.vocabDefs = nil
	
	for i, e := range entries {
		
		h := vd.buildVocabDef(e, strconv.Itoa(i))
		
		html += h
		
		vd.vocabDefs = append(vd.vocabDefs, h)
		
		}

	top := strings.Replace(`<span class = "top">[{0}]<br/></span>`, `{0}`, contentSample, -1)

	html = buildDefHeader() + top + html + buildDefFooter()

	vd.SetHtml(html)

	}

func (vd *VocabDoc) buildVocabDef(e Entry, i string) string {

	slink := `<a href = "copyVocabDef:{0}"><img src = "icons/add.png" align = "right" /></a>` // height="24" width="24"
	sreading := `<span class = "reading">[{0}]</span>` // <br/>
	sexpression := `<span class = "expression">[{0}]<br/></span>`
	sglossary := `<span class = "glossary">{0}</span>` //<br/>

	link := strings.Replace(slink, `{0}`, i, -1)
	reading := strings.Replace(sreading, `{0}`, e.Furigana, -1)
	expression := strings.Replace(sexpression, `{0}`, e.Word, -1)
	glossary := strings.Replace(sglossary, `{0}`, e.Definition, -1)			

	htmlbreak := `<hr width="90%" align="left" clear="all">` //`<span><hr width="75%"></span>` //`<br clear = "all"/>`
	
	//html := link + reading + expression + glossary + htmlbreak
	
	html := `<div>`  + link + reading + expression + glossary + htmlbreak + `</div>`

	return html

}




//    def buildDefHeader(self):

//        return u"""
//                <html><head><style>
//                body {{ background-color: {0} }}
//                span.expression {{ font-size: {1}px; font-family: '{2}'; color: {3}; line-height: {10}px }}
//                span.reading {{ font-size: {4}px; font-family: '{5}'; color: {6}; line-height: {11}px }}
//                span.glossary {{ font-size: {7}px; font-family: '{8}'; color: {9}; line-height: {12}px }}
//                </style></head><body>""".format(self.bg, self.efs, self.eft, self.efg, self.rfs, self.rft, self.rfg, self.gfs, self.gft, self.gfg, self.elh, self.rlh, self.glh) #+ html + "</body></html>"

//    def buildDefFooter(self):
//        return '</body></html>'

//    def buildEmpty(self):
//        return u"""
//            <p>No definitions to display.</p>
//            <p>Mouse over text with the <em>middle mouse button</em> or <em>shift key</em> pressed to search.</p>
//            <p>You can also also input terms in the search box below."""








//    def buildVocabDef(self, definition, index, query):
//        reading = unicode()
//        if definition['reading']:
//            reading = u'<span class = "reading">[{0}]<br/></span>'.format(definition['reading'])

//        rules = unicode()
//        if len(definition['rules']) > 0:
//            rules = ' &lt; '.join(definition['rules'])
//            rules = '<span class = "rules">({0})<br/></span>'.format(rules)

//        links = '<a href = "copyVocabDef:{0}"><img src = "img/icon_add_expression.png" align = "right"/></a>'.format(index)
//        if query is not None:
//            if query('vocab', yomi_base.reader_util.markupVocabExp(definition)):
//                links += '<a href = "addVocabExp:{0}"><img src = "://img/img/icon_add_expression.png" align = "right"/></a>'.format(index)
//            if query('vocab', yomi_base.reader_util.markupVocabReading(definition)):
//                links += '<a href = "addVocabReading:{0}"><img src = "://img/img/icon_add_reading.png" align = "right"/></a>'.format(index)

//        html = u"""
//            <span class = "links">{0}</span>
//            <span class = "expression">{1}<br/></span>
//            {2}
//            <span class = "glossary">{3}<br/></span>
//            {4}
//            <br clear = "all"/>""".format(links, definition['expression'], reading, definition['glossary'], rules)
//        #print html
//        return html

//    def buildVocabDefs(self, definitions, query):
//        html = self.buildDefHeader()
//        if len(definitions) > 0:
//            for i, definition in enumerate(definitions):
//                html += self.buildVocabDef(definition, i, query)
//        else:
//            html += self.buildEmpty()

//        return html + self.buildDefFooter()






//    def buildKanjiDef(self, definition, index, query):
//        links = '<a href = "copyKanjiDef:{0}"><img src = "://img/img/icon_copy_definition.png" align = "right"/></a>'.format(index)
//        if query is not None and query('kanji', yomi_base.reader_util.markupKanji(definition)):
//            links += '<a href = "addKanji:{0}"><img src = "://img/img/icon_add_expression.png" align = "right"/></a>'.format(index)

//        readings = ', '.join([definition['kunyomi'], definition['onyomi']])
//        html = u"""
//            <span class = "links">{0}</span>
//            <span class = "expression">{1}<br/></span>
//            <span class = "reading">[{2}]<br/></span>
//            <span class = "glossary">{3}<br/></span>
//            <br clear = "all"/>""".format(links, definition['character'], readings, definition['glossary'])

//        return html

//    def buildKanjiDefs(self, definitions, query):
//        html = self.buildDefHeader()

//        if len(definitions) > 0:
//            for i, definition in enumerate(definitions):
//                html += self.buildKanjiDef(definition, i, query)
//        else:
//            html += self.buildEmpty()

//        return html + self.buildDefFooter()
