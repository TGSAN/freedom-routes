package routes

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"
)

var _ = fmt.Println

func TestIpString(t *testing.T) {
	ip := Ip{"192.168.1.1", "24", "255.255.255.0"}
	actual := ip.String()
	expect := "192.168.1.1/24"
	if actual != expect {
		t.Errorf("%s; want %s", actual, expect)
	}
}

var cidr2masks = []struct {
	in  int
	out string
}{
	{24, "255.255.255.0"},
	{32, "255.255.255.255"},
	{1, "128.0.0.0"},
}

func TestCidr2mask(t *testing.T) {
	for _, v := range cidr2masks {
		actual := cidr2mask(v.in)
		if actual != v.out {
			t.Errorf("cidr2mask(%d) = %s; want %s", v.in, actual, v.out)
		}
	}
}

func prepareTmpDir() {
	teardownTmpDir()
	os.Mkdir("testtmp", 0755)
}

func teardownTmpDir() {
	os.RemoveAll("testtmp")
}

func verifyGenerate(t *testing.T, path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error(err)
	}
	matched, _ := regexp.Match(`192.168.1.1/24`, data)
	if !matched {
		t.Errorf("Can't find 192.168.1.1/24 in %s", path)
	}

	info, _ := os.Stat(path)
	if info.Mode() != 0755 {
		t.Errorf("File mode is not 0755 - %s", path)
	}
}

func TestGenerate(t *testing.T) {
	prepareTmpDir()
	defer teardownTmpDir()

	ips := []Ip{
		Ip{"192.168.1.1", "24", "255.255.255.0"},
		Ip{"1.1.1.1", "32", "255.255.255.255"},
	}

	Generate("linux", ips, "testtmp")
	verifyGenerate(t, "testtmp/routes-up.sh")
	verifyGenerate(t, "testtmp/routes-down.sh")
}

func verifyIps(t *testing.T, methodName string, ins, outs []Ip) {
	for i, out := range outs {
		in := ins[i]
		if out.String() != in.String() {
			t.Errorf("%s()[%d] = %s, want %s", methodName, i, in, out)
		}
	}
}

func TestFetchLocalIps(t *testing.T) {
	LOCAL_PATH = "testdata/routes"
	outs := []Ip{
		Ip{"192.168.1.1", "24", "255.255.255.0"},
		Ip{"1.1.1.1", "32", "255.255.255.255"},
	}

	verifyIps(t, "FetchLocalIps", FetchLocalIps(), outs)
}

func TestFetchRemoteIps(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `
apnic|CN|ipv4|1.1.1.1|1|20100806|allocated
# comment
apnic|CN|ipv4|192.168.1.1|256|20100806|allocated`)
	}))
	defer ts.Close()

	REMOTE_URL = ts.URL
	outs := []Ip{
		Ip{"1.1.1.1", "32", "255.255.255.255"},
		Ip{"192.168.1.1", "24", "255.255.255.0"},
	}

	verifyIps(t, "FetchRemoteIps", FetchRemoteIps(), outs)
}
