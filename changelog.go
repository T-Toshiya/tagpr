package rcpr

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Songmu/flextime"
)

var (
	versionLinkReg    = regexp.MustCompile(`\*\*Full Changelog\*\*: (https://.*)$`)
	semverFromLinkReg = regexp.MustCompile(`.*[./](v?[0-9]+\.[0-9]+\.[0-9]+)`)
	newContribReg     = regexp.MustCompile(`(?ms)## New Contributors.*\z`)
)

func convertKeepAChangelogFormat(md string) string {
	md = strings.TrimSpace(md)

	var link string
	md = versionLinkReg.ReplaceAllStringFunc(md, func(match string) string {
		m := versionLinkReg.FindStringSubmatch(match)
		link = m[1]
		return ""
	})
	var semvStr string
	if m := semverFromLinkReg.FindStringSubmatch(link); len(m) > 1 {
		semvStr = m[1]
	}
	now := flextime.Now()

	heading := fmt.Sprintf("## [%s](%s) - %s", semvStr, link, now.Format("2006-01-02"))
	md = strings.Replace(md, "## What's Changed", heading, 1)
	md = strings.ReplaceAll(md, "\n* ", "\n- ")
	md = newContribReg.ReplaceAllString(md, "")

	return strings.TrimSpace(md) + "\n"
}
