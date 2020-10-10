package submission_controller

import (
	http "clx/http-handler"
	"encoding/json"
	terminal "github.com/wayneashleyberry/terminal-dimensions"
	"strconv"

	"gitlab.com/tslocum/cview"
)

const (
	maximumStoriesToDisplay = 30
)

// SubmissionHandler stores submissions and pages
type SubmissionHandler struct {
	Submissions      []Submission
	Pages            *cview.Pages
	PagesRetrieved   int
	CurrentPage      int
	StoriesListed    int
	ScreenHeight     int
	StoriesToDisplay int
}

func NewSubmissionHandler() *SubmissionHandler {
	sh := new(SubmissionHandler)

	y, _ := terminal.Height()
	sh.ScreenHeight = int(y)
	sh.StoriesToDisplay = min(sh.ScreenHeight / 2, maximumStoriesToDisplay)

	return sh
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func (sh *SubmissionHandler) GetStoriesToDisplay() int {
	return sh.StoriesToDisplay
}

func (sh *SubmissionHandler) GetSubmissionInfo(i int) (string, string) {
	submission := sh.GetSubmission(i)

	rank := i + 1
	indentedRank := strconv.Itoa(rank) + "." + GetRankIndentBlock(rank)

	primary := indentedRank + submission.Title + submission.GetDomain()

	secondary := "[::d]" + "    " + submission.GetPoints() + " points by " + submission.Author + " " + submission.Time + " | " + submission.GetComments() + " comments" + "[-:-:-]"

	return primary, secondary
}

func (sh *SubmissionHandler) GetSubmission(i int) Submission {
	return sh.Submissions[i]
}

// Submission represents the JSON structure as
// retrieved from cheeaun's unofficial HN API
type Submission struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Points        int    `json:"points"`
	Author        string `json:"user"`
	Time          string `json:"time_ago"`
	CommentsCount int    `json:"comments_count"`
	URL           string `json:"url"`
	Domain        string `json:"domain"`
	Type          string `json:"type"`
}

func FetchSubmissions(sh *SubmissionHandler) {
	sh.PagesRetrieved++
	p := strconv.Itoa(sh.PagesRetrieved)
	JSON, _ := http.Get("http://node-hnapi.herokuapp.com/news?page=" + p)
	var submissions []Submission
	_ = json.Unmarshal(JSON, &submissions)
	sh.Submissions = append(sh.Submissions, submissions...)
}

func GetDomain(domain string) string {
	if domain == "" {
		return ""
	}
	return "[::d]" + " " + paren(domain) + "[-:-:-]"
}

func (s Submission) GetDomain() string {
	domain := s.Domain
	if domain == "" {
		return ""
	}
	return "[::d]" + " " + paren(domain) + "[-:-:-]"
}

func (s Submission) GetComments() string {
	return strconv.Itoa(s.CommentsCount)
}

func (s Submission) GetPoints() string {
	return strconv.Itoa(s.Points)
}

func paren(text string) string {
	return "(" + text + ")"
}

func GetRankIndentBlock(rank int) string {
	largeIndent := "  "
	smallIndent := " "
	if rank > 9 {
		return smallIndent
	}
	return largeIndent
}
