package matcher

import (
	_ "github.com/ekzhu/minhash-lsh"
	_ "github.com/hhatto/gorst"
	_ "github.com/jdkato/prose/chunk"
	"github.com/jfrog/go-license-discovery/utils"
	"github.com/jfrog/gofrog/lru"
	"github.com/jfrog/licenseclassifier/licenseclassifier"
	_ "github.com/pkg/errors"
	_ "golang.org/x/exp/rand"
	_ "gonum.org/v1/gonum/stat/distuv"
	"gopkg.in/goyy/goyy.v0/util/strings"
	_ "gopkg.in/src-d/go-billy.v4/memfs"
	_ "gopkg.in/src-d/go-billy.v4/osfs"
	_ "gopkg.in/src-d/go-git.v4/plumbing"
	_ "gopkg.in/src-d/go-git.v4/plumbing/filemode"
	_ "gopkg.in/src-d/go-git.v4/plumbing/object"
	_ "gopkg.in/src-d/go-git.v4/plumbing/storer"
	_ "gopkg.in/src-d/go-git.v4/storage/filesystem"
	_ "gopkg.in/src-d/go-git.v4/storage/memory"
	"gopkg.in/src-d/go-license-detector.v2/licensedb"
	"gopkg.in/src-d/go-license-detector.v2/licensedb/filer"
	"time"
)

/**
 * @author chenk on 29/08/2017.
 */
var lc *licenseclassifier.License

var extractedLicenseCache *lru.Cache = lru.New(1000, lru.WithExpiry(168*time.Hour))

// read licenses db and instantiate the classifier
func InitLicenseMatcher(licensesFolder string) error {
	licenseclassifier.LicensesDir(licensesFolder + utils.Licenses)
	var err error
	lc, err = licenseclassifier.New(0.85)
	if err != nil {
		return err
	}
	return nil
}

// try to match license txt to known spdx license
// 1st try to match to license classifier and then to license detector
func MatchLicenseTxt(licenseTxt string) []string {
	licSha := utils.LicenseToSha256(licenseTxt)
	if val, ok := extractedLicenseCache.Get(licSha); ok {
		return val.([]string)
	} else {
		n := lc.MultipleMatch(licenseTxt, true)
		set := utils.NewSet()
		if len(n) > 0 {
			for _, m := range n {
				if m.Confidence > 0.85 {
					set.AddString(m.Name)
				}
			}
			if set.Size() > 0 {
				licenses := set.StringValues()
				extractedLicenseCache.Add(licSha, licenses)
				return licenses
			}
		}
		licenses := GetLicenseFromDetector(licenseTxt, licSha)
		if len(licenses) > 0 {
			return licenses
		}
	}
	return []string{utils.LICENSE_UNKNOWN}
}

type DiscoveryFiler struct {
	LicenseContent string
}

func (df DiscoveryFiler) ReadFile(path string) (content []byte, err error) {
	return []byte(df.LicenseContent), nil
}
func (df DiscoveryFiler) ReadDir(path string) ([]filer.File, error) {
	return []filer.File{{Name: "license.txt"}}, nil
}
func (df DiscoveryFiler) Close() {

}

func GetLicenseFromDetector(licenseTxt string, licSha string) []string {
	set := utils.NewSet()
	if len(licSha) == 0 {
		licSha = utils.LicenseToSha256(licenseTxt)
	}
	if val, ok := extractedLicenseCache.Get(licSha); ok {
		return val.([]string)
	} else {
		df := DiscoveryFiler{LicenseContent: licenseTxt}
		licMap, err := licensedb.Detect(df)
		if err != nil {
			return set.StringValues()
		}
		var maxScore float32
		var lic string
		if len(licMap) > 0 {
			for key, value := range licMap {
				if maxScore < value && value > 0.8 {
					maxScore = value
					lic = key
				}
			}
			if len(lic) > 0 {
				extractedLicenseCache.Add(licSha, []string{lic})
				set.AddString(lic)
			}
		}
	}
	return set.StringValues()
}

// read licenses which presented as a comment in the pom file
func GetPomCommentLicense(pomTxt string) string {
	pomComments := utils.ReadPomComments(pomTxt)
	if len(pomComments) == 0 {
		return utils.LICENSE_UNKNOWN
	}
	tLicense := strings.TrimSpace(pomComments)
	lic := GetLicenseFromDetector(tLicense, utils.EMPTY_STRING)
	if len(lic) > 0 {
		return lic[0]
	}

	return utils.LICENSE_UNKNOWN
}
