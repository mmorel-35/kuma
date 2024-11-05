package version

import (
	"fmt"
	"runtime"
	"strings"
)

var (
	Product     = "Kuma"
	basedOnKuma = ""
	version     = "unknown"
	gitTag      = "unknown"
	gitCommit   = "unknown"
	buildDate   = "unknown"
	Envoy       = "unknown"
)

type BuildInfo struct {
	Product     string
	Version     string
	GitTag      string
	GitCommit   string
	BuildDate   string
	BasedOnKuma string
}

func (b BuildInfo) FormatDetailedProductInfo() string {
	base := []string{
		"Product:       " + b.Product,
		"Version:       " + b.Version,
		"Git Tag:       " + b.GitTag,
		"Git Commit:    " + b.GitCommit,
		"Build Date:    " + b.BuildDate,
	}
	if b.BasedOnKuma != "" {
		base = append(base, "Based on Kuma: "+b.BasedOnKuma)
	}
	return strings.Join(
		base,
		"\n",
	)
}

func shortCommit(c string) string {
	if len(c) < 7 {
		return c
	}
	return c[:7]
}

func (b BuildInfo) AsMap() map[string]string {
	res := map[string]string{
		"product":    b.Product,
		"version":    b.Version,
		"build_date": b.BuildDate,
		"git_commit": shortCommit(b.GitCommit),
		"git_tag":    b.GitTag,
	}
	if b.BasedOnKuma != "" {
		res["based_on_kuma"] = b.BasedOnKuma
	}
	return res
}

func (b BuildInfo) UserAgent(component string) string {
	commit := shortCommit(b.GitCommit)
	if b.BasedOnKuma != "" {
		commit = fmt.Sprintf("%s/kuma-%s", commit, b.BasedOnKuma)
	}
	return fmt.Sprintf("%s/%s (%s; %s; %s/%s)",
		component,
		b.Version,
		runtime.GOOS,
		runtime.GOARCH,
		b.Product,
		commit)
}

var Build BuildInfo

func init() {
	Build = BuildInfo{
		Product:     Product,
		Version:     version,
		GitTag:      gitTag,
		GitCommit:   gitCommit,
		BuildDate:   buildDate,
		BasedOnKuma: basedOnKuma,
	}
}
