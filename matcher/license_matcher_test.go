package matcher

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"io/ioutil"
)

const (
	NoLicense = "no licenes"
)

func TestInitLicenseMatcherWithNoLicensesDB(t *testing.T) {
	err:=InitLicenseMatcher("unknown folder")
	if err == nil{
		t.Fatal(err)
	}
}

func TestInitLicenseMatcherWithLicensesDB(t *testing.T) {
	err:=InitLicenseMatcher("./licenses/")
	if err != nil{
		t.Fatal(err)
	}
}

func TestMatchLicenseTxtNoLicenseFile(t *testing.T){
	err:=InitLicenseMatcher("./licenses/")
	if err != nil{
		t.Fatal(err)
	}
	lics:=MatchLicenseTxt(NoLicense)
	assert.True(t,len(lics) == 0 || lics[0]=="Unknown")

}

func TestMatchLicenseTxtWithClassifier(t *testing.T){
	err:=InitLicenseMatcher("./licenses/")
	if err != nil{
		t.Fatal(err)
	}
	f,err:=os.Open("./fixtures/Multi_LICENSE.txt")
	defer f.Close()
	if err != nil{
		t.Fatal(err)
	}
	data,err:=ioutil.ReadAll(f)
	if err != nil{
		t.Fatal(err)
	}
	lics:=MatchLicenseTxt(string(data))
	assert.True(t,len(lics) == 6)
}

func TestMatchLicenseTxtWithDetector(t *testing.T){
	err:=InitLicenseMatcher("./licenses/")
	if err != nil{
		t.Fatal(err)
	}
	f,err:=os.Open("./fixtures/Partial_LICENSE.txt")
	defer f.Close()

	if err != nil{
		t.Fatal(err)
	}
	data,err:=ioutil.ReadAll(f)
	if err != nil{
		t.Fatal(err)
	}
	lics:=MatchLicenseTxt(string(data))
	assert.True(t,len(lics) == 1)
}

func TestMatchLicenseTxtWithPom(t *testing.T){
	err:=InitLicenseMatcher("./licenses/")
	if err != nil{
		t.Fatal(err)
	}
	f,err:=os.Open("./fixtures/PomWithLicenseAsComment.xml")
	defer f.Close()

	if err != nil{
		t.Fatal(err)
	}
	data,err:=ioutil.ReadAll(f)
	if err != nil{
		t.Fatal(err)
	}
	lic:=GetPomCommentLicense(string(data))
	assert.True(t,lic == "Apache-2.0")
}
