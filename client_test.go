package httptool

import (
	"net/http"
	"testing"
)

func TestClientStructure(t *testing.T) {
	client := NewClient()

	if client == nil {
		t.Fatal("Client can't be nil")
	}

	if client.Cookies == nil {
		t.Fatal("Cookies can't be nil")
	}

	if client.Header == nil {
		t.Fatal("Header can't be nil")
	}
}

func TestCookie(t *testing.T) {
	client := NewClient()

	cookie := http.Cookie{Name: "Cookie2", Value: "ttest"}
	client.SetCookie(&cookie)
	if co, ok := client.GetCookie("Cookie2"); ok {
		if co.Value != "ttest" {
			t.Error("Get Cookie value are different")
		}
	} else {
		t.Error("Can't get cookie value")
	}

	cookie = http.Cookie{Name: "Cookie52", Value: "6695"}
	client.SetCookie(&cookie)
	if co, ok := client.GetCookie("Cookie52"); ok {
		if co.Value != "6695" {
			t.Error("Get Cookie value are different in second time")
		}
	} else {
		t.Error("Can't get cookie value in second time")
	}

	client.DelCookie("Cookie52")
	if _, ok := client.GetCookie("Cookie52"); ok {
		t.Error("Del cookie fail")
	}

}

func TestFakeUserAgent(t *testing.T) {
	client := NewClient()

	client.AddFakeUserAgent()

	ua := client.GetHeader("user-agent")
	if ua != Chrome {
		t.Error("Set different user-agent")
	}

	client.SetHeader("user-agent", Edge)
	ua = client.GetHeader("user-agent")
	hs := client.GetHeadersArray("user-agent")
	if len(hs) != 1 || ua != Edge {
		t.Error("Set user-agent by AddHeader fail")
	}
}

func TestHeader(t *testing.T) {
	client := NewClient()

	client.SetHeader("test", "vvv")

	val := client.GetHeader("test")
	hs := client.GetHeadersArray("test")

	if val == "vvv" {
		if len(hs) != 1 {
			t.Error("HttpClient.SetHeader value wrong")
		}
	} else {
		t.Error("HttpClient.SetHeader not set value")
	}

	client.AddHeader("test2", "666")
	val = client.GetHeader("test2")
	hs = client.GetHeadersArray("test2")

	if val == "666" {
		if len(client.Header) != 2 || len(hs) != 1 {
			t.Error("Add different header fail")
		}
	} else {
		t.Error("Add different header lost origin key")
	}

	client.AddHeader("tEst", "kokkk")
	val = client.GetHeader("test")
	hs = client.GetHeadersArray("test")

	if val == "vvv" {
		if len(client.Header) != 2 {
			t.Error("Headers key not correct")
		} else if len(hs) != 2 {
			t.Logf("%d: %v\n", len(val), val)
			t.Logf("header: %v\n", client.Header)
			t.Error("Add same header keys number not correct")
		} else if hs[0] != "vvv" || hs[1] != "kokkk" {
			t.Error("Add same header change original key")
		}
	} else {
		t.Error("Add same header lost origin key")
	}

	client.DelHeader("test")

	val = client.GetHeader("test")
	hs = client.GetHeadersArray("test")

	if val == "test" {
		t.Logf("Wired: %v", val)
		t.Error("Delete header fail")
	} else if len(client.Header) != 1 {
		t.Error("Delete header not complete")
	}

	val = client.GetHeader("test2")
	hs = client.GetHeadersArray("test2")

	if val == "666" {
		if len(hs) != 1 {
			t.Error("Delete header strange")
		}
	} else {
		t.Error("Delete wrong key")
	}
}
