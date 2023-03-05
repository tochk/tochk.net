// Code generated by qtc from "main.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line main.qtpl:1
package templates

//line main.qtpl:2
import (
	"time"
	"tochkru-golang/internal/app/datastruct"
)

//line main.qtpl:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line main.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line main.qtpl:7
func StreamHeader(qw422016 *qt422016.Writer, title string) {
//line main.qtpl:7
	qw422016.N().S(`
<!DOCTYPE html>
<html>
<head>
    <title>`)
//line main.qtpl:11
	qw422016.E().S(title)
//line main.qtpl:11
	qw422016.N().S(` - tochk.net</title>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8"/>
    <link href="/static/style.css" rel="stylesheet" type="text/css"/>
    <link href="/static/favicon.ico" rel="shortcut icon" type="image/x-icon"/>
    <!-- Global site tag (gtag.js) - Google Analytics -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=G-5PZ91DE9PJ"></script>
    <script>
      window.dataLayer = window.dataLayer || [];
      function gtag(){dataLayer.push(arguments);}
      gtag('js', new Date());

      gtag('config', 'G-5PZ91DE9PJ');
    </script>
    <link href='https://fonts.googleapis.com/css?family=Roboto+Slab&subset=latin,cyrillic' rel='stylesheet'
          type='text/css'>
</head>
<body>
`)
//line main.qtpl:28
}

//line main.qtpl:28
func WriteHeader(qq422016 qtio422016.Writer, title string) {
//line main.qtpl:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line main.qtpl:28
	StreamHeader(qw422016, title)
//line main.qtpl:28
	qt422016.ReleaseWriter(qw422016)
//line main.qtpl:28
}

//line main.qtpl:28
func Header(title string) string {
//line main.qtpl:28
	qb422016 := qt422016.AcquireByteBuffer()
//line main.qtpl:28
	WriteHeader(qb422016, title)
//line main.qtpl:28
	qs422016 := string(qb422016.B)
//line main.qtpl:28
	qt422016.ReleaseByteBuffer(qb422016)
//line main.qtpl:28
	return qs422016
//line main.qtpl:28
}

//line main.qtpl:30
func StreamMenu(qw422016 *qt422016.Writer, tags []datastruct.Tags) {
//line main.qtpl:30
	qw422016.N().S(`
<header>
    <div id="header_for_buttons">
        <div id="logo_name" onclick="window.location.href='/'"><h1>TOCHK.NET</h1></div>
        <div id="buttons">
            <div id="button" onclick="window.location.href='/projects/'">ПРОЕКТЫ</div>
            <div id="button" onclick="window.location.href='/articles/'">СТАТЬИ</div>
            <div id="button" onclick="window.location.href='/'">ГЛАВНАЯ</div>
        </div>
    </div>
</header>
<main>
    <div id="top_menu">
        <div id="header_top"></div>
    </div>
    <div id="left_menu">
        <div id="block_for_left">
            <div id="head_left_menu">TAGS</div>
            <div id="info_left_menu">`)
//line main.qtpl:48
	for idx, tag := range tags {
//line main.qtpl:48
		qw422016.N().S(` `)
//line main.qtpl:48
		qw422016.E().S(tag.Name)
//line main.qtpl:48
		if idx != len(tags)-1 {
//line main.qtpl:48
			qw422016.N().S(`, `)
//line main.qtpl:48
		}
//line main.qtpl:48
	}
//line main.qtpl:48
	qw422016.N().S(`</div>
        </div>
    </div>
    <div id="right_menu">
        <content>
`)
//line main.qtpl:53
}

//line main.qtpl:53
func WriteMenu(qq422016 qtio422016.Writer, tags []datastruct.Tags) {
//line main.qtpl:53
	qw422016 := qt422016.AcquireWriter(qq422016)
//line main.qtpl:53
	StreamMenu(qw422016, tags)
//line main.qtpl:53
	qt422016.ReleaseWriter(qw422016)
//line main.qtpl:53
}

//line main.qtpl:53
func Menu(tags []datastruct.Tags) string {
//line main.qtpl:53
	qb422016 := qt422016.AcquireByteBuffer()
//line main.qtpl:53
	WriteMenu(qb422016, tags)
//line main.qtpl:53
	qs422016 := string(qb422016.B)
//line main.qtpl:53
	qt422016.ReleaseByteBuffer(qb422016)
//line main.qtpl:53
	return qs422016
//line main.qtpl:53
}

//line main.qtpl:56
func StreamFooter(qw422016 *qt422016.Writer) {
//line main.qtpl:56
	qw422016.N().S(`
        </content>
    </div>
</main>
<footer>
    design by lenokh // tochk.net 2013 - `)
//line main.qtpl:61
	qw422016.E().S(time.Now().Format("2006"))
//line main.qtpl:61
	qw422016.N().S(`
</footer>
</body>
</html>
`)
//line main.qtpl:65
}

//line main.qtpl:65
func WriteFooter(qq422016 qtio422016.Writer) {
//line main.qtpl:65
	qw422016 := qt422016.AcquireWriter(qq422016)
//line main.qtpl:65
	StreamFooter(qw422016)
//line main.qtpl:65
	qt422016.ReleaseWriter(qw422016)
//line main.qtpl:65
}

//line main.qtpl:65
func Footer() string {
//line main.qtpl:65
	qb422016 := qt422016.AcquireByteBuffer()
//line main.qtpl:65
	WriteFooter(qb422016)
//line main.qtpl:65
	qs422016 := string(qb422016.B)
//line main.qtpl:65
	qt422016.ReleaseByteBuffer(qb422016)
//line main.qtpl:65
	return qs422016
//line main.qtpl:65
}

//line main.qtpl:68
func StreamProject(qw422016 *qt422016.Writer, project datastruct.Projects, tags map[int]datastruct.Tags, teamMembers map[int]datastruct.TeamMembers) {
//line main.qtpl:68
	qw422016.N().S(`
<div id='block_new_sait' onclick="window.open('/project/`)
//line main.qtpl:69
	qw422016.N().D(project.ID)
//line main.qtpl:69
	qw422016.N().S(`', '_self');">
    <div id='block_in_block_sait'>
    <div id='block_in_block_sait_2'>
    <div id='block_sait_head'>`)
//line main.qtpl:72
	qw422016.E().S(project.Title)
//line main.qtpl:72
	qw422016.N().S(`</div>
    <div id='block_sait_scrin_info'>
    <div class='block_sait_scrin_1'>
    <div class='block_sait_scrin_2' style='background-image: url("`)
//line main.qtpl:75
	qw422016.N().S(project.ImageURL)
//line main.qtpl:75
	qw422016.N().S(`")'></div>
    <div class='block_sait_scrin_3'>`)
//line main.qtpl:76
	qw422016.N().S(project.ShortDescription)
//line main.qtpl:76
	qw422016.N().S(`</div>
    </div></div></div></div><div id='block_sait_info'>
    <div id='block_sait_info_2'>
    <div id='sait_date' style='width: 100px'>`)
//line main.qtpl:79
	qw422016.E().S(project.CreatedAt.Format("2006.01.02"))
//line main.qtpl:79
	qw422016.N().S(`</div>
    <div id='sait_data_about_block' style='width: 180px'>
    <div style='width: 80px' id='sait_author'>Team<br><text_green>`)
//line main.qtpl:81
	for _, teamMemberID := range project.TeamMembersIDs {
//line main.qtpl:81
		qw422016.E().S(teamMembers[teamMemberID].Name)
//line main.qtpl:81
		qw422016.N().S(`<br>`)
//line main.qtpl:81
	}
//line main.qtpl:81
	qw422016.N().S(`</text_green></div>
    <div style='width: 80px' id='sait_cat'>Tags<br><text_green>`)
//line main.qtpl:82
	for _, tagID := range project.TagsIDs {
//line main.qtpl:82
		qw422016.E().S(tags[tagID].Name)
//line main.qtpl:82
		qw422016.N().S(`<br>`)
//line main.qtpl:82
	}
//line main.qtpl:82
	qw422016.N().S(`</text_green></div>
    </div><div id='sait_right_footer'>View</div>
</div></div></div>
`)
//line main.qtpl:85
}

//line main.qtpl:85
func WriteProject(qq422016 qtio422016.Writer, project datastruct.Projects, tags map[int]datastruct.Tags, teamMembers map[int]datastruct.TeamMembers) {
//line main.qtpl:85
	qw422016 := qt422016.AcquireWriter(qq422016)
//line main.qtpl:85
	StreamProject(qw422016, project, tags, teamMembers)
//line main.qtpl:85
	qt422016.ReleaseWriter(qw422016)
//line main.qtpl:85
}

//line main.qtpl:85
func Project(project datastruct.Projects, tags map[int]datastruct.Tags, teamMembers map[int]datastruct.TeamMembers) string {
//line main.qtpl:85
	qb422016 := qt422016.AcquireByteBuffer()
//line main.qtpl:85
	WriteProject(qb422016, project, tags, teamMembers)
//line main.qtpl:85
	qs422016 := string(qb422016.B)
//line main.qtpl:85
	qt422016.ReleaseByteBuffer(qb422016)
//line main.qtpl:85
	return qs422016
//line main.qtpl:85
}

//line main.qtpl:87
func StreamArticle(qw422016 *qt422016.Writer, article datastruct.Articles, tags map[int]datastruct.Tags, teamMembers map[int]datastruct.TeamMembers) {
//line main.qtpl:87
	qw422016.N().S(`
<div id='block_news'>
    <div id='block_in_block_news'>
    <div id='block_in_block_news_2'>
    <div id='block_news_head' style='font-size:20px;'><a href='/article/`)
//line main.qtpl:91
	qw422016.N().D(article.ID)
//line main.qtpl:91
	qw422016.N().S(`'><h2 style="font-size: 20px">`)
//line main.qtpl:91
	qw422016.E().S(article.Title)
//line main.qtpl:91
	qw422016.N().S(`</h2></a></div>
    <div id='block_news_content' style='font-size:16px;'>`)
//line main.qtpl:92
	qw422016.N().S(article.ShortText)
//line main.qtpl:92
	qw422016.N().S(`</div>
    </div></div>
    <div id='block_news_info'>
    <div id='block_news_info_2'>
    <div id='block_date' style='font-size: 20px;'>`)
//line main.qtpl:96
	qw422016.E().S(article.CreatedAt.Format("2006.01.02"))
//line main.qtpl:96
	qw422016.N().S(`</div>
    <div id='block_data'>
    <div id='data_author'> Author<br><text_green>`)
//line main.qtpl:98
	for _, teamMemberID := range article.TeamMembersIDs {
//line main.qtpl:98
		qw422016.E().S(teamMembers[teamMemberID].Name)
//line main.qtpl:98
		qw422016.N().S(`<br>`)
//line main.qtpl:98
	}
//line main.qtpl:98
	qw422016.N().S(`<br></text_green></div>
    <div id='data_category'> Tags<br><text_green>`)
//line main.qtpl:99
	for _, tagID := range article.TagsIDs {
//line main.qtpl:99
		qw422016.E().S(tags[tagID].Name)
//line main.qtpl:99
		qw422016.N().S(`<br>`)
//line main.qtpl:99
	}
//line main.qtpl:99
	qw422016.N().S(`<br></text_green></div>
    </div>
    <div id='block_info_footer' style='font-size: 20px;'> <a href='/article/`)
//line main.qtpl:101
	qw422016.N().D(article.ID)
//line main.qtpl:101
	qw422016.N().S(`'>read more</a></div>
    </div></div>
</div>
`)
//line main.qtpl:104
}

//line main.qtpl:104
func WriteArticle(qq422016 qtio422016.Writer, article datastruct.Articles, tags map[int]datastruct.Tags, teamMembers map[int]datastruct.TeamMembers) {
//line main.qtpl:104
	qw422016 := qt422016.AcquireWriter(qq422016)
//line main.qtpl:104
	StreamArticle(qw422016, article, tags, teamMembers)
//line main.qtpl:104
	qt422016.ReleaseWriter(qw422016)
//line main.qtpl:104
}

//line main.qtpl:104
func Article(article datastruct.Articles, tags map[int]datastruct.Tags, teamMembers map[int]datastruct.TeamMembers) string {
//line main.qtpl:104
	qb422016 := qt422016.AcquireByteBuffer()
//line main.qtpl:104
	WriteArticle(qb422016, article, tags, teamMembers)
//line main.qtpl:104
	qs422016 := string(qb422016.B)
//line main.qtpl:104
	qt422016.ReleaseByteBuffer(qb422016)
//line main.qtpl:104
	return qs422016
//line main.qtpl:104
}
