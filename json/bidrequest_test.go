package json

import (
	"encoding/json"
	"testing"

	"github.com/bsm/openrtb"
	"github.com/json-iterator/go"
	"github.com/pquerna/ffjson/ffjson"
)

var (
	requestData = []byte(`{
    "id": "7979d0c78074638bbdf739ffdf285c7e1c74a691",
    "at": 2,
    "tmax": 143,
    "imp": [{
        "id": "1",
        "tagid": "76334",
        "iframebuster": ["ALL"],
        "banner": {
            "w": 300,
            "h": 250,
            "pos": 1,
            "battr": [9, 1, 14014, 3, 13, 10, 8, 14],
            "api": [3, 1000],
            "topframe": 1
        }
    }],
    "app": {
        "id": "20625",
        "cat": ["IAB1"],
        "name": "com.cheezburger.icanhas",
        "domain": "http://cheezburger.com",
        "privacypolicy": 1,
        "publisher": {
            "id": "8428"
        },
        "ext": {
            "storerating": 1,
            "appstoreid": "457637357"
        }
    },
    "device": {
        "make": "Samsung",
        "model": "SCH-I535",
        "os": "Android",
        "osv": "4.3",
        "ua": "Mozilla/5.0 (Linux; U; Android 4.3; en-us; SCH-I535 Build/JSS15J) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
        "ip": "192.168.1.1",
        "language": "en",
        "devicetype": 1,
        "js": 1,
        "connectiontype": 3,
        "dpidsha1": "F099E6D1C485756C45D1EEACB33C73B55C4BC499",
        "carrier": "Verizon Wireless",
        "geo": {
            "country": "USA",
            "region": "PA",
            "type": 3,
            "ext": {
                "latlonconsent": 1
            }
        }
    },
    "user": {
        "id": "bd5adc55dcbab4bf090604df4f543d90b09f0c88",
        "ext": {
            "sessiondepth": 207
        }
    }
}`)
)

func bidRequest() *openrtb.BidRequest {
	rd := &openrtb.BidRequest{}
	if err := json.Unmarshal(requestData, rd); err != nil {
		panic(err.Error())
	}

	return rd
}

func BenchmarkStandardMarshal(b *testing.B) {
	br := bidRequest()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(br)
		if err != nil {
			b.Fatalf("got marshaler error: %q", err.Error())
		}
	}
}

func BenchmarkStandardUnmarshal(b *testing.B) {
	br := &openrtb.BidRequest{}
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(requestData, br)
		if err != nil {
			b.Fatalf("got unmarshaler error: %q", err.Error())
		}
	}
}

func BenchmarkJsoniterMarshall(b *testing.B) {
	p := bidRequest()
	for i := 0; i < b.N; i++ {
		_, err := jsoniter.Marshal(p)
		if err != nil {
			b.Fatalf("got jsonite-marshaler error: %q", err.Error())
		}
	}
}

func BenchmarkJsoniterUnmarshal(b *testing.B) {
	br := &openrtb.BidRequest{}
	for i := 0; i < b.N; i++ {
		err := jsoniter.Unmarshal(requestData, br)
		if err != nil {
			b.Fatalf("got unmarshaler error: %q", err.Error())
		}
	}
}

func BenchmarkFFJsonMarshal(b *testing.B) {
	p := &FF_BidRequest{
		*bidRequest(),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ffjson.Marshal(p)
		if err != nil {
			b.Fatalf("got marshaler error: %q", err.Error())
		}
	}
}

func BenchmarkFFJsonUnmarshal(b *testing.B) {
	p := &FF_BidRequest{}
	for i := 0; i < b.N; i++ {
		err := ffjson.Unmarshal(requestData, p)
		if err != nil {
			b.Fatalf("got ffjson unmarshaler error: %q", err.Error())
		}
	}
}

func BenchmarkEasyJsonMarshal(b *testing.B) {
	br := &Easy_BidRequest{
		*bidRequest(),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(br)
		if err != nil {
			b.Fatalf("got marshaler error: %q", err.Error())
		}
	}
}

func BenchmarkEasyJsonUnmarshal(b *testing.B) {
	br := &Easy_BidRequest{}
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(requestData, br)
		if err != nil {
			b.Fatalf("got ffjson unmarshaler error: %q", err.Error())
		}
	}
}
