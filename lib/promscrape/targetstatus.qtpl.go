// Code generated by qtc from "targetstatus.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line lib/promscrape/targetstatus.qtpl:1
package promscrape

//line lib/promscrape/targetstatus.qtpl:1
import (
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/prompbmarshal"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/promrelabel"
	"time"
)

//line lib/promscrape/targetstatus.qtpl:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line lib/promscrape/targetstatus.qtpl:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line lib/promscrape/targetstatus.qtpl:9
func StreamTargetsResponsePlain(qw422016 *qt422016.Writer, jts []jobTargetsStatuses, emptyJobs []string, showOriginLabels bool) {
//line lib/promscrape/targetstatus.qtpl:11
	for _, js := range jts {
//line lib/promscrape/targetstatus.qtpl:11
		qw422016.N().S(`job=`)
//line lib/promscrape/targetstatus.qtpl:12
		qw422016.N().Q(js.job)
//line lib/promscrape/targetstatus.qtpl:12
		qw422016.N().S(`(`)
//line lib/promscrape/targetstatus.qtpl:12
		qw422016.N().D(js.upCount)
//line lib/promscrape/targetstatus.qtpl:12
		qw422016.N().S(`/`)
//line lib/promscrape/targetstatus.qtpl:12
		qw422016.N().D(js.targetsTotal)
//line lib/promscrape/targetstatus.qtpl:12
		qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:12
		qw422016.N().S(`up)`)
//line lib/promscrape/targetstatus.qtpl:13
		qw422016.N().S(`
`)
//line lib/promscrape/targetstatus.qtpl:14
		for _, ts := range js.targetsStatus {
//line lib/promscrape/targetstatus.qtpl:15
			qw422016.N().S("\t")
//line lib/promscrape/targetstatus.qtpl:15
			qw422016.N().S(`state=`)
//line lib/promscrape/targetstatus.qtpl:16
			if ts.up {
//line lib/promscrape/targetstatus.qtpl:16
				qw422016.N().S(`up`)
//line lib/promscrape/targetstatus.qtpl:16
			} else {
//line lib/promscrape/targetstatus.qtpl:16
				qw422016.N().S(`down`)
//line lib/promscrape/targetstatus.qtpl:16
			}
//line lib/promscrape/targetstatus.qtpl:16
			qw422016.N().S(`,`)
//line lib/promscrape/targetstatus.qtpl:16
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:16
			qw422016.N().S(`endpoint=`)
//line lib/promscrape/targetstatus.qtpl:17
			qw422016.N().S(ts.sw.Config.ScrapeURL)
//line lib/promscrape/targetstatus.qtpl:17
			qw422016.N().S(`,`)
//line lib/promscrape/targetstatus.qtpl:17
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:17
			qw422016.N().S(`labels=`)
//line lib/promscrape/targetstatus.qtpl:18
			qw422016.N().S(promLabelsString(promrelabel.FinalizeLabels(nil, ts.sw.Config.Labels)))
//line lib/promscrape/targetstatus.qtpl:18
			qw422016.N().S(`,`)
//line lib/promscrape/targetstatus.qtpl:18
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:19
			if showOriginLabels {
//line lib/promscrape/targetstatus.qtpl:19
				qw422016.N().S(`originalLabels=`)
//line lib/promscrape/targetstatus.qtpl:19
				qw422016.N().S(promLabelsString(ts.sw.Config.OriginalLabels))
//line lib/promscrape/targetstatus.qtpl:19
				qw422016.N().S(`,`)
//line lib/promscrape/targetstatus.qtpl:19
				qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:19
			}
//line lib/promscrape/targetstatus.qtpl:19
			qw422016.N().S(`scrapes_total=`)
//line lib/promscrape/targetstatus.qtpl:20
			qw422016.N().D(ts.scrapesTotal)
//line lib/promscrape/targetstatus.qtpl:20
			qw422016.N().S(`,`)
//line lib/promscrape/targetstatus.qtpl:20
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:20
			qw422016.N().S(`scrapes_failed=`)
//line lib/promscrape/targetstatus.qtpl:21
			qw422016.N().D(ts.scrapesFailed)
//line lib/promscrape/targetstatus.qtpl:21
			qw422016.N().S(`,`)
//line lib/promscrape/targetstatus.qtpl:21
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:21
			qw422016.N().S(`last_scrape=`)
//line lib/promscrape/targetstatus.qtpl:22
			qw422016.N().FPrec(ts.getDurationFromLastScrape().Seconds(), 3)
//line lib/promscrape/targetstatus.qtpl:22
			qw422016.N().S(`s ago,`)
//line lib/promscrape/targetstatus.qtpl:22
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:22
			qw422016.N().S(`scrape_duration=`)
//line lib/promscrape/targetstatus.qtpl:23
			qw422016.N().D(int(ts.scrapeDuration))
//line lib/promscrape/targetstatus.qtpl:23
			qw422016.N().S(`ms,`)
//line lib/promscrape/targetstatus.qtpl:23
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:23
			qw422016.N().S(`samples_scraped=`)
//line lib/promscrape/targetstatus.qtpl:24
			qw422016.N().D(ts.samplesScraped)
//line lib/promscrape/targetstatus.qtpl:24
			qw422016.N().S(`,`)
//line lib/promscrape/targetstatus.qtpl:24
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:24
			qw422016.N().S(`error=`)
//line lib/promscrape/targetstatus.qtpl:25
			if ts.err != nil {
//line lib/promscrape/targetstatus.qtpl:25
				qw422016.N().S(ts.err.Error())
//line lib/promscrape/targetstatus.qtpl:25
			}
//line lib/promscrape/targetstatus.qtpl:26
			qw422016.N().S(`
`)
//line lib/promscrape/targetstatus.qtpl:27
		}
//line lib/promscrape/targetstatus.qtpl:28
	}
//line lib/promscrape/targetstatus.qtpl:30
	for _, jobName := range emptyJobs {
//line lib/promscrape/targetstatus.qtpl:30
		qw422016.N().S(`job=`)
//line lib/promscrape/targetstatus.qtpl:31
		qw422016.N().Q(jobName)
//line lib/promscrape/targetstatus.qtpl:31
		qw422016.N().S(`(0/0 up)`)
//line lib/promscrape/targetstatus.qtpl:32
		qw422016.N().S(`
`)
//line lib/promscrape/targetstatus.qtpl:33
	}
//line lib/promscrape/targetstatus.qtpl:35
}

//line lib/promscrape/targetstatus.qtpl:35
func WriteTargetsResponsePlain(qq422016 qtio422016.Writer, jts []jobTargetsStatuses, emptyJobs []string, showOriginLabels bool) {
//line lib/promscrape/targetstatus.qtpl:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promscrape/targetstatus.qtpl:35
	StreamTargetsResponsePlain(qw422016, jts, emptyJobs, showOriginLabels)
//line lib/promscrape/targetstatus.qtpl:35
	qt422016.ReleaseWriter(qw422016)
//line lib/promscrape/targetstatus.qtpl:35
}

//line lib/promscrape/targetstatus.qtpl:35
func TargetsResponsePlain(jts []jobTargetsStatuses, emptyJobs []string, showOriginLabels bool) string {
//line lib/promscrape/targetstatus.qtpl:35
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promscrape/targetstatus.qtpl:35
	WriteTargetsResponsePlain(qb422016, jts, emptyJobs, showOriginLabels)
//line lib/promscrape/targetstatus.qtpl:35
	qs422016 := string(qb422016.B)
//line lib/promscrape/targetstatus.qtpl:35
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promscrape/targetstatus.qtpl:35
	return qs422016
//line lib/promscrape/targetstatus.qtpl:35
}

//line lib/promscrape/targetstatus.qtpl:37
func StreamTargetsResponseHTML(qw422016 *qt422016.Writer, jts []jobTargetsStatuses, emptyJobs []string, onlyUnhealthy bool) {
//line lib/promscrape/targetstatus.qtpl:37
	qw422016.N().S(`<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous"><title>Scrape targets</title><script>function collapse_all() {for (var i = 0; i <`)
//line lib/promscrape/targetstatus.qtpl:47
	qw422016.N().D(len(jts))
//line lib/promscrape/targetstatus.qtpl:47
	qw422016.N().S(`; i++) {var id = "table-" + i;document.getElementById(id).style.display = 'none';}}function expand_all() {for (var i = 0; i <`)
//line lib/promscrape/targetstatus.qtpl:53
	qw422016.N().D(len(jts))
//line lib/promscrape/targetstatus.qtpl:53
	qw422016.N().S(`; i++) {var id = "table-" + i;document.getElementById(id).style.display = 'block';}}</script></head><body class="m-3"><h1>Scrape targets</h1><div style="padding: 3px"><button type="button" class="btn`)
//line lib/promscrape/targetstatus.qtpl:63
	qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:63
	if !onlyUnhealthy {
//line lib/promscrape/targetstatus.qtpl:63
		qw422016.N().S(`btn-primary`)
//line lib/promscrape/targetstatus.qtpl:63
	} else {
//line lib/promscrape/targetstatus.qtpl:63
		qw422016.N().S(`btn-secondary`)
//line lib/promscrape/targetstatus.qtpl:63
	}
//line lib/promscrape/targetstatus.qtpl:63
	qw422016.N().S(`" onclick="location.href='targets'">All</button><button type="button" class="btn`)
//line lib/promscrape/targetstatus.qtpl:66
	qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:66
	if onlyUnhealthy {
//line lib/promscrape/targetstatus.qtpl:66
		qw422016.N().S(`btn-primary`)
//line lib/promscrape/targetstatus.qtpl:66
	} else {
//line lib/promscrape/targetstatus.qtpl:66
		qw422016.N().S(`btn-secondary`)
//line lib/promscrape/targetstatus.qtpl:66
	}
//line lib/promscrape/targetstatus.qtpl:66
	qw422016.N().S(`" onclick="location.href='targets?show_only_unhealthy=true'">Unhealthy</button><button type="button" class="btn btn-primary" onclick="collapse_all()">Collapse all</button><button type="button" class="btn btn-secondary" onclick="expand_all()">Expand all</button></div>`)
//line lib/promscrape/targetstatus.qtpl:76
	for i, js := range jts {
//line lib/promscrape/targetstatus.qtpl:77
		if onlyUnhealthy && js.upCount == js.targetsTotal {
//line lib/promscrape/targetstatus.qtpl:77
			continue
//line lib/promscrape/targetstatus.qtpl:77
		}
//line lib/promscrape/targetstatus.qtpl:77
		qw422016.N().S(`<div><h4>`)
//line lib/promscrape/targetstatus.qtpl:80
		qw422016.E().S(js.job)
//line lib/promscrape/targetstatus.qtpl:80
		qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:80
		qw422016.N().S(`(`)
//line lib/promscrape/targetstatus.qtpl:80
		qw422016.N().D(js.upCount)
//line lib/promscrape/targetstatus.qtpl:80
		qw422016.N().S(`/`)
//line lib/promscrape/targetstatus.qtpl:80
		qw422016.N().D(js.targetsTotal)
//line lib/promscrape/targetstatus.qtpl:80
		qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:80
		qw422016.N().S(`up)<button type="button" class="btn btn-primary" onclick="document.getElementById('table-`)
//line lib/promscrape/targetstatus.qtpl:81
		qw422016.N().D(i)
//line lib/promscrape/targetstatus.qtpl:81
		qw422016.N().S(`').style.display='none'">collapse</button><button type="button" class="btn btn-secondary" onclick="document.getElementById('table-`)
//line lib/promscrape/targetstatus.qtpl:82
		qw422016.N().D(i)
//line lib/promscrape/targetstatus.qtpl:82
		qw422016.N().S(`').style.display='block'">expand</button></h4><div id="table-`)
//line lib/promscrape/targetstatus.qtpl:84
		qw422016.N().D(i)
//line lib/promscrape/targetstatus.qtpl:84
		qw422016.N().S(`"><table class="table table-striped table-hover table-bordered table-sm"><thead><tr><th scope="col">Endpoint</th><th scope="col">State</th><th scope="col" title="scrape target labels">Labels</th><th scope="col" title="total scrapes">Scrapes</th><th scope="col" title="total scrape errors">Errors</th><th scope="col" title="the time of the last scrape">Last Scrape</th><th scope="col" title="the duration of the last scrape">Duration</th><th scope="col" title="the number of metrics scraped during the last scrape">Samples</th><th scope="col" title="error from the last scrape (if any)">Last error</th></tr></thead><tbody>`)
//line lib/promscrape/targetstatus.qtpl:100
		for _, ts := range js.targetsStatus {
//line lib/promscrape/targetstatus.qtpl:102
			endpoint := ts.sw.Config.ScrapeURL
			targetID := getTargetID(ts.sw)
			lastScrapeTime := ts.getDurationFromLastScrape()

//line lib/promscrape/targetstatus.qtpl:106
			if onlyUnhealthy && ts.up {
//line lib/promscrape/targetstatus.qtpl:106
				continue
//line lib/promscrape/targetstatus.qtpl:106
			}
//line lib/promscrape/targetstatus.qtpl:106
			qw422016.N().S(`<tr`)
//line lib/promscrape/targetstatus.qtpl:107
			if !ts.up {
//line lib/promscrape/targetstatus.qtpl:107
				qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:107
				qw422016.N().S(`class="alert alert-danger" role="alert"`)
//line lib/promscrape/targetstatus.qtpl:107
			}
//line lib/promscrape/targetstatus.qtpl:107
			qw422016.N().S(`><td><a href="`)
//line lib/promscrape/targetstatus.qtpl:108
			qw422016.E().S(endpoint)
//line lib/promscrape/targetstatus.qtpl:108
			qw422016.N().S(`" target="_blank">`)
//line lib/promscrape/targetstatus.qtpl:108
			qw422016.E().S(endpoint)
//line lib/promscrape/targetstatus.qtpl:108
			qw422016.N().S(`</a> (<a href="target_response?id=`)
//line lib/promscrape/targetstatus.qtpl:109
			qw422016.E().S(targetID)
//line lib/promscrape/targetstatus.qtpl:109
			qw422016.N().S(`" target="_blank" title="click to fetch target response on behalf of the scraper">response</a>)</td><td>`)
//line lib/promscrape/targetstatus.qtpl:111
			if ts.up {
//line lib/promscrape/targetstatus.qtpl:111
				qw422016.N().S(`UP`)
//line lib/promscrape/targetstatus.qtpl:111
			} else {
//line lib/promscrape/targetstatus.qtpl:111
				qw422016.N().S(`DOWN`)
//line lib/promscrape/targetstatus.qtpl:111
			}
//line lib/promscrape/targetstatus.qtpl:111
			qw422016.N().S(`</td><td><div title="click to show original labels" onclick="document.getElementById('original_labels_`)
//line lib/promscrape/targetstatus.qtpl:113
			qw422016.E().S(targetID)
//line lib/promscrape/targetstatus.qtpl:113
			qw422016.N().S(`').style.display='block'">`)
//line lib/promscrape/targetstatus.qtpl:114
			streamformatLabel(qw422016, promrelabel.FinalizeLabels(nil, ts.sw.Config.Labels))
//line lib/promscrape/targetstatus.qtpl:114
			qw422016.N().S(`</div><div style="display:none" id="original_labels_`)
//line lib/promscrape/targetstatus.qtpl:116
			qw422016.E().S(targetID)
//line lib/promscrape/targetstatus.qtpl:116
			qw422016.N().S(`">`)
//line lib/promscrape/targetstatus.qtpl:117
			streamformatLabel(qw422016, ts.sw.Config.OriginalLabels)
//line lib/promscrape/targetstatus.qtpl:117
			qw422016.N().S(`</div></td><td>`)
//line lib/promscrape/targetstatus.qtpl:120
			qw422016.N().D(ts.scrapesTotal)
//line lib/promscrape/targetstatus.qtpl:120
			qw422016.N().S(`</td><td>`)
//line lib/promscrape/targetstatus.qtpl:121
			qw422016.N().D(ts.scrapesFailed)
//line lib/promscrape/targetstatus.qtpl:121
			qw422016.N().S(`</td><td>`)
//line lib/promscrape/targetstatus.qtpl:123
			if lastScrapeTime < 365*24*time.Hour {
//line lib/promscrape/targetstatus.qtpl:124
				qw422016.N().FPrec(lastScrapeTime.Seconds(), 3)
//line lib/promscrape/targetstatus.qtpl:124
				qw422016.N().S(`s ago`)
//line lib/promscrape/targetstatus.qtpl:125
			} else {
//line lib/promscrape/targetstatus.qtpl:125
				qw422016.N().S(`none`)
//line lib/promscrape/targetstatus.qtpl:127
			}
//line lib/promscrape/targetstatus.qtpl:127
			qw422016.N().S(`<td>`)
//line lib/promscrape/targetstatus.qtpl:128
			qw422016.N().D(int(ts.scrapeDuration))
//line lib/promscrape/targetstatus.qtpl:128
			qw422016.N().S(`ms</td><td>`)
//line lib/promscrape/targetstatus.qtpl:129
			qw422016.N().D(ts.samplesScraped)
//line lib/promscrape/targetstatus.qtpl:129
			qw422016.N().S(`</td><td>`)
//line lib/promscrape/targetstatus.qtpl:130
			if ts.err != nil {
//line lib/promscrape/targetstatus.qtpl:130
				qw422016.E().S(ts.err.Error())
//line lib/promscrape/targetstatus.qtpl:130
			}
//line lib/promscrape/targetstatus.qtpl:130
			qw422016.N().S(`</td></tr>`)
//line lib/promscrape/targetstatus.qtpl:132
		}
//line lib/promscrape/targetstatus.qtpl:132
		qw422016.N().S(`</tbody></table></div></div>`)
//line lib/promscrape/targetstatus.qtpl:137
	}
//line lib/promscrape/targetstatus.qtpl:139
	for _, jobName := range emptyJobs {
//line lib/promscrape/targetstatus.qtpl:139
		qw422016.N().S(`<div><h4><a>`)
//line lib/promscrape/targetstatus.qtpl:142
		qw422016.E().S(jobName)
//line lib/promscrape/targetstatus.qtpl:142
		qw422016.N().S(`(0/0 up)</a></h4><table class="table table-striped table-hover table-bordered table-sm"><thead><tr><th scope="col">Endpoint</th><th scope="col">State</th><th scope="col">Labels</th><th scope="col">Last Scrape</th><th scope="col">Scrape Duration</th><th scope="col">Samples Scraped</th><th scope="col">Error</th></tr></thead></table></div>`)
//line lib/promscrape/targetstatus.qtpl:158
	}
//line lib/promscrape/targetstatus.qtpl:158
	qw422016.N().S(`</body></html>`)
//line lib/promscrape/targetstatus.qtpl:161
}

//line lib/promscrape/targetstatus.qtpl:161
func WriteTargetsResponseHTML(qq422016 qtio422016.Writer, jts []jobTargetsStatuses, emptyJobs []string, onlyUnhealthy bool) {
//line lib/promscrape/targetstatus.qtpl:161
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promscrape/targetstatus.qtpl:161
	StreamTargetsResponseHTML(qw422016, jts, emptyJobs, onlyUnhealthy)
//line lib/promscrape/targetstatus.qtpl:161
	qt422016.ReleaseWriter(qw422016)
//line lib/promscrape/targetstatus.qtpl:161
}

//line lib/promscrape/targetstatus.qtpl:161
func TargetsResponseHTML(jts []jobTargetsStatuses, emptyJobs []string, onlyUnhealthy bool) string {
//line lib/promscrape/targetstatus.qtpl:161
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promscrape/targetstatus.qtpl:161
	WriteTargetsResponseHTML(qb422016, jts, emptyJobs, onlyUnhealthy)
//line lib/promscrape/targetstatus.qtpl:161
	qs422016 := string(qb422016.B)
//line lib/promscrape/targetstatus.qtpl:161
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promscrape/targetstatus.qtpl:161
	return qs422016
//line lib/promscrape/targetstatus.qtpl:161
}

//line lib/promscrape/targetstatus.qtpl:163
func streamformatLabel(qw422016 *qt422016.Writer, labels []prompbmarshal.Label) {
//line lib/promscrape/targetstatus.qtpl:163
	qw422016.N().S(`{`)
//line lib/promscrape/targetstatus.qtpl:165
	for i, label := range labels {
//line lib/promscrape/targetstatus.qtpl:166
		qw422016.E().S(label.Name)
//line lib/promscrape/targetstatus.qtpl:166
		qw422016.N().S(`=`)
//line lib/promscrape/targetstatus.qtpl:166
		qw422016.E().Q(label.Value)
//line lib/promscrape/targetstatus.qtpl:167
		if i+1 < len(labels) {
//line lib/promscrape/targetstatus.qtpl:167
			qw422016.N().S(`,`)
//line lib/promscrape/targetstatus.qtpl:167
			qw422016.N().S(` `)
//line lib/promscrape/targetstatus.qtpl:167
		}
//line lib/promscrape/targetstatus.qtpl:168
	}
//line lib/promscrape/targetstatus.qtpl:168
	qw422016.N().S(`}`)
//line lib/promscrape/targetstatus.qtpl:170
}

//line lib/promscrape/targetstatus.qtpl:170
func writeformatLabel(qq422016 qtio422016.Writer, labels []prompbmarshal.Label) {
//line lib/promscrape/targetstatus.qtpl:170
	qw422016 := qt422016.AcquireWriter(qq422016)
//line lib/promscrape/targetstatus.qtpl:170
	streamformatLabel(qw422016, labels)
//line lib/promscrape/targetstatus.qtpl:170
	qt422016.ReleaseWriter(qw422016)
//line lib/promscrape/targetstatus.qtpl:170
}

//line lib/promscrape/targetstatus.qtpl:170
func formatLabel(labels []prompbmarshal.Label) string {
//line lib/promscrape/targetstatus.qtpl:170
	qb422016 := qt422016.AcquireByteBuffer()
//line lib/promscrape/targetstatus.qtpl:170
	writeformatLabel(qb422016, labels)
//line lib/promscrape/targetstatus.qtpl:170
	qs422016 := string(qb422016.B)
//line lib/promscrape/targetstatus.qtpl:170
	qt422016.ReleaseByteBuffer(qb422016)
//line lib/promscrape/targetstatus.qtpl:170
	return qs422016
//line lib/promscrape/targetstatus.qtpl:170
}
