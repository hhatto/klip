package klip

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var RegexpHighlight = `(ハイライト|Highlight)`
var RegexpBookmark = `(ブックマーク|Bookmark)`
var RegexpNote = `(メモ|Note)`
var RegexpAddedOnEN = ` \| Added on `
var RegexpAddedOnJP = `(: |： )`
var RegexpLocation = `([0-9]{1,})-([0-9]{1,})`
var RegexpLocationNote = `[0-9]{1,}`
var RegexpAddedYMD = `([0-9]{4})年([0-9]{1,2})月([0-9]{1,2})`

type ClipType int

const (
	Highlight ClipType = iota
	Bookmark
	Note
)

const (
	SepLine = iota
	TitleLine
	InfoLine
	BlankLine
	ContentLine
)

type Location struct {
	Start int
	End   int
}

type Meta struct {
	Type     ClipType
	Location Location
}

type KindleClipping struct {
	Content string
	Meta    Meta
	Author  string
	AddedOn time.Time
	Title   string
}

type KindleClippings []*KindleClipping

func parse(fp *os.File) (clips []KindleClipping, err error) {
	lineNum := 0
	scanner := bufio.NewScanner(fp)
	titleSplitter, err := regexp.Compile(` \(`)
	if err != nil {
		log.Fatalf("regexp.Complie error: %v", err)
	}

	var clip KindleClipping
	for scanner.Scan() {
		lineNum += 1
		buf := scanner.Bytes()
		offset := lineNum % 5
		if offset == TitleLine {
			// get title
			clip = KindleClipping{}
			//fmt.Println(offset, lineNum, string(buf))
			t := titleSplitter.Split(string(buf), 5)
			title := t[:len(t)-1]
			clip.Title = strings.Join(title, "")

			// get author
			clip.Author = strings.Split(t[len(t)-1], ")")[0]
		} else if offset == InfoLine {
			// get type
			if regexp.MustCompile(RegexpHighlight).Match(buf) {
				clip.Meta.Type = Highlight
			} else if regexp.MustCompile(RegexpBookmark).Match(buf) {
				clip.Meta.Type = Bookmark
			} else if regexp.MustCompile(RegexpNote).Match(buf) {
				clip.Meta.Type = Note
			}

			// get location
			if clip.Meta.Type == Highlight {
				t := regexp.MustCompile(RegexpLocation).FindSubmatch(buf)
				clip.Meta.Location.Start, _ = strconv.Atoi(string(t[1]))
				clip.Meta.Location.End, _ = strconv.Atoi(string(t[2]))
			} else {
				t := regexp.MustCompile(RegexpLocationNote).FindSubmatch(buf)
				clip.Meta.Location.Start, _ = strconv.Atoi(string(t[0]))
				clip.Meta.Location.End, _ = strconv.Atoi(string(t[0]))
			}

			// get added
			t := regexp.MustCompile(RegexpAddedOnEN).Split(string(buf), -1)
			if len(t) == 2 {
				// en
				clip.AddedOn, err = time.Parse("Monday, January 02, 2006 15:04:05 AM", string(t[1]))
				if err != nil {
					log.Fatalf("time.Parse error: %v", err)
				}
			} else {
				// jp
				t := regexp.MustCompile(RegexpAddedYMD).FindSubmatch(buf)
				y, _ := strconv.Atoi(string(t[1]))
				m, _ := strconv.Atoi(string(t[2]))
				d, _ := strconv.Atoi(string(t[3]))
				t = bytes.Split(buf, []byte(" "))
				ymdhms := fmt.Sprintf("%04d-%02d-%02d %s", y, m, d, t[len(t)-1])
				clip.AddedOn, err = time.Parse("2006-01-02 15:04:05", ymdhms)
				if err != nil {
					log.Fatalf("time.Parse error: %v", err)
				}
			}
		} else if offset == ContentLine {
			clip.Content = string(buf)
			clips = append(clips, clip)
		} else {
			continue
		}
	}
	return clips, err
}

func Load(filename string) (clips []KindleClipping, err error) {
	var fp *os.File
	if fp, err = os.Open(filename); err != nil {
		log.Fatal("ReadFile() error: ", err)
	}

	clips, err = parse(fp)

	return clips, err
}
