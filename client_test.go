package httptool

import "testing"

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

	client.SetCookie("Cookie2", "ttest")
	if val, ok := client.GetCookie("Cookie2"); ok {
		if val != "ttest" {
			t.Error("Get Cookie value are different")
		}
	} else {
		t.Error("Can't get cookie value")
	}

	client.SetCookie("Cookie52", "6695")
	if val, ok := client.GetCookie("Cookie52"); ok {
		if val != "6695" {
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

	if ua, ok := client.GetHeader("user-agent"); ok {
		if ua[0] != Chrome {
			t.Error("Set different user-agent")
		}
	} else {
		t.Error("Set default fake user agent by AddFakeUserAgent fail")
	}

	client.SetHeader("user-agent", Edge)
	if ua, ok := client.GetHeader("user-agent"); ok {
		if len(ua) != 1 || ua[0] != Edge {
			t.Error("Set user-agent by AddHeader fail")
		}
	} else {
		t.Error("Set user-agent fail")
	}
}

func TestHeader(t *testing.T) {
	client := NewClient()

	client.SetHeader("test", "vvv")

	if val, ok := client.GetHeader("test"); ok {
		if len(val) != 1 || val[0] != "vvv" {
			t.Error("HttpClient.SetHeader value wrong")
		}
	} else {
		t.Error("HttpClient.SetHeader not set value")
	}

	client.AddHeader("test2", "666")
	if val, ok := client.GetHeader("test2"); ok {
		if len(client.Header) != 2 || len(val) != 1 || val[0] != "666" {
			t.Error("Add different header fail")
		}
	} else {
		t.Error("Add different header lost origin key")
	}

	client.AddHeader("tEst", "kokkk")
	if val, ok := client.GetHeader("test"); ok {
		if len(client.Header) != 2 {
			t.Error("Headers key not correct")
		} else if len(val) != 2 {
			t.Logf("%d: %v\n", len(val), val)
			t.Logf("header: %v\n", client.Header)
			t.Error("Add same header keys number not correct")
		} else if val[0] != "vvv" || val[1] != "kokkk" {
			t.Error("Add same header change original key")
		}
	} else {
		t.Error("Add same header lost origin key")
	}

	client.DelHeader("test")

	if val, ok := client.GetHeader("test"); ok {
		t.Logf("Wired: %v", val)
		t.Error("Delete header fail")
	} else if len(client.Header) != 1 {
		t.Error("Delete header not complete")
	}

	if val, ok := client.GetHeader("test2"); ok {
		if len(val) != 1 || val[0] != "666" {
			t.Error("Delete header strange")
		}
	} else {
		t.Error("Delete wrong key")
	}
}
