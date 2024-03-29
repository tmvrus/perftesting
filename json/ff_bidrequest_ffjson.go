// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: ff_bidrequest.go

package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bsm/openrtb"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *FF_BidRequest) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *FF_BidRequest) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "id":`)
	fflib.WriteJsonString(buf, string(j.ID))
	buf.WriteByte(',')
	if len(j.Imp) != 0 {
		buf.WriteString(`"imp":`)
		if j.Imp != nil {
			buf.WriteString(`[`)
			for i, v := range j.Imp {
				if i != 0 {
					buf.WriteString(`,`)
				}
				/* Struct fall back. type=openrtb.Impression kind=struct */
				err = buf.Encode(&v)
				if err != nil {
					return err
				}
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if j.Site != nil {
		if true {
			/* Struct fall back. type=openrtb.Site kind=struct */
			buf.WriteString(`"site":`)
			err = buf.Encode(j.Site)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if j.App != nil {
		if true {
			/* Struct fall back. type=openrtb.App kind=struct */
			buf.WriteString(`"app":`)
			err = buf.Encode(j.App)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if j.Device != nil {
		if true {
			/* Struct fall back. type=openrtb.Device kind=struct */
			buf.WriteString(`"device":`)
			err = buf.Encode(j.Device)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if j.User != nil {
		if true {
			/* Struct fall back. type=openrtb.User kind=struct */
			buf.WriteString(`"user":`)
			err = buf.Encode(j.User)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if j.Test != 0 {
		buf.WriteString(`"test":`)
		fflib.FormatBits2(buf, uint64(j.Test), 10, j.Test < 0)
		buf.WriteByte(',')
	}
	buf.WriteString(`"at":`)
	fflib.FormatBits2(buf, uint64(j.AuctionType), 10, j.AuctionType < 0)
	buf.WriteByte(',')
	if j.TMax != 0 {
		buf.WriteString(`"tmax":`)
		fflib.FormatBits2(buf, uint64(j.TMax), 10, j.TMax < 0)
		buf.WriteByte(',')
	}
	if len(j.WSeat) != 0 {
		buf.WriteString(`"wseat":`)
		if j.WSeat != nil {
			buf.WriteString(`[`)
			for i, v := range j.WSeat {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(j.BSeat) != 0 {
		buf.WriteString(`"bseat":`)
		if j.BSeat != nil {
			buf.WriteString(`[`)
			for i, v := range j.BSeat {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(j.WLang) != 0 {
		buf.WriteString(`"wlang":`)
		if j.WLang != nil {
			buf.WriteString(`[`)
			for i, v := range j.WLang {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if j.AllImps != 0 {
		buf.WriteString(`"allimps":`)
		fflib.FormatBits2(buf, uint64(j.AllImps), 10, j.AllImps < 0)
		buf.WriteByte(',')
	}
	if len(j.Cur) != 0 {
		buf.WriteString(`"cur":`)
		if j.Cur != nil {
			buf.WriteString(`[`)
			for i, v := range j.Cur {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(j.Bcat) != 0 {
		buf.WriteString(`"bcat":`)
		if j.Bcat != nil {
			buf.WriteString(`[`)
			for i, v := range j.Bcat {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(j.BAdv) != 0 {
		buf.WriteString(`"badv":`)
		if j.BAdv != nil {
			buf.WriteString(`[`)
			for i, v := range j.BAdv {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(j.BApp) != 0 {
		buf.WriteString(`"bapp":`)
		if j.BApp != nil {
			buf.WriteString(`[`)
			for i, v := range j.BApp {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if j.Source != nil {
		if true {
			/* Struct fall back. type=openrtb.Source kind=struct */
			buf.WriteString(`"source":`)
			err = buf.Encode(j.Source)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if j.Regs != nil {
		if true {
			/* Struct fall back. type=openrtb.Regulations kind=struct */
			buf.WriteString(`"regs":`)
			err = buf.Encode(j.Regs)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if len(j.Ext) != 0 {
		buf.WriteString(`"ext":`)

		{

			obj, err = j.Ext.MarshalJSON()
			if err != nil {
				return err
			}
			buf.Write(obj)

		}
		buf.WriteByte(',')
	}
	if j.Pmp != nil {
		if true {
			/* Struct fall back. type=openrtb.Pmp kind=struct */
			buf.WriteString(`"pmp":`)
			err = buf.Encode(j.Pmp)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtFF_BidRequestbase = iota
	ffjtFF_BidRequestnosuchkey

	ffjtFF_BidRequestID

	ffjtFF_BidRequestImp

	ffjtFF_BidRequestSite

	ffjtFF_BidRequestApp

	ffjtFF_BidRequestDevice

	ffjtFF_BidRequestUser

	ffjtFF_BidRequestTest

	ffjtFF_BidRequestAuctionType

	ffjtFF_BidRequestTMax

	ffjtFF_BidRequestWSeat

	ffjtFF_BidRequestBSeat

	ffjtFF_BidRequestWLang

	ffjtFF_BidRequestAllImps

	ffjtFF_BidRequestCur

	ffjtFF_BidRequestBcat

	ffjtFF_BidRequestBAdv

	ffjtFF_BidRequestBApp

	ffjtFF_BidRequestSource

	ffjtFF_BidRequestRegs

	ffjtFF_BidRequestExt

	ffjtFF_BidRequestPmp
)

var ffjKeyFF_BidRequestID = []byte("id")

var ffjKeyFF_BidRequestImp = []byte("imp")

var ffjKeyFF_BidRequestSite = []byte("site")

var ffjKeyFF_BidRequestApp = []byte("app")

var ffjKeyFF_BidRequestDevice = []byte("device")

var ffjKeyFF_BidRequestUser = []byte("user")

var ffjKeyFF_BidRequestTest = []byte("test")

var ffjKeyFF_BidRequestAuctionType = []byte("at")

var ffjKeyFF_BidRequestTMax = []byte("tmax")

var ffjKeyFF_BidRequestWSeat = []byte("wseat")

var ffjKeyFF_BidRequestBSeat = []byte("bseat")

var ffjKeyFF_BidRequestWLang = []byte("wlang")

var ffjKeyFF_BidRequestAllImps = []byte("allimps")

var ffjKeyFF_BidRequestCur = []byte("cur")

var ffjKeyFF_BidRequestBcat = []byte("bcat")

var ffjKeyFF_BidRequestBAdv = []byte("badv")

var ffjKeyFF_BidRequestBApp = []byte("bapp")

var ffjKeyFF_BidRequestSource = []byte("source")

var ffjKeyFF_BidRequestRegs = []byte("regs")

var ffjKeyFF_BidRequestExt = []byte("ext")

var ffjKeyFF_BidRequestPmp = []byte("pmp")

// UnmarshalJSON umarshall json - template of ffjson
func (j *FF_BidRequest) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *FF_BidRequest) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtFF_BidRequestbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtFF_BidRequestnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyFF_BidRequestApp, kn) {
						currentKey = ffjtFF_BidRequestApp
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyFF_BidRequestAuctionType, kn) {
						currentKey = ffjtFF_BidRequestAuctionType
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyFF_BidRequestAllImps, kn) {
						currentKey = ffjtFF_BidRequestAllImps
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffjKeyFF_BidRequestBSeat, kn) {
						currentKey = ffjtFF_BidRequestBSeat
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyFF_BidRequestBcat, kn) {
						currentKey = ffjtFF_BidRequestBcat
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyFF_BidRequestBAdv, kn) {
						currentKey = ffjtFF_BidRequestBAdv
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyFF_BidRequestBApp, kn) {
						currentKey = ffjtFF_BidRequestBApp
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffjKeyFF_BidRequestCur, kn) {
						currentKey = ffjtFF_BidRequestCur
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffjKeyFF_BidRequestDevice, kn) {
						currentKey = ffjtFF_BidRequestDevice
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyFF_BidRequestExt, kn) {
						currentKey = ffjtFF_BidRequestExt
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyFF_BidRequestID, kn) {
						currentKey = ffjtFF_BidRequestID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyFF_BidRequestImp, kn) {
						currentKey = ffjtFF_BidRequestImp
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffjKeyFF_BidRequestPmp, kn) {
						currentKey = ffjtFF_BidRequestPmp
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffjKeyFF_BidRequestRegs, kn) {
						currentKey = ffjtFF_BidRequestRegs
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyFF_BidRequestSite, kn) {
						currentKey = ffjtFF_BidRequestSite
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyFF_BidRequestSource, kn) {
						currentKey = ffjtFF_BidRequestSource
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyFF_BidRequestTest, kn) {
						currentKey = ffjtFF_BidRequestTest
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyFF_BidRequestTMax, kn) {
						currentKey = ffjtFF_BidRequestTMax
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffjKeyFF_BidRequestUser, kn) {
						currentKey = ffjtFF_BidRequestUser
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffjKeyFF_BidRequestWSeat, kn) {
						currentKey = ffjtFF_BidRequestWSeat
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyFF_BidRequestWLang, kn) {
						currentKey = ffjtFF_BidRequestWLang
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyFF_BidRequestPmp, kn) {
					currentKey = ffjtFF_BidRequestPmp
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyFF_BidRequestExt, kn) {
					currentKey = ffjtFF_BidRequestExt
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyFF_BidRequestRegs, kn) {
					currentKey = ffjtFF_BidRequestRegs
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyFF_BidRequestSource, kn) {
					currentKey = ffjtFF_BidRequestSource
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyFF_BidRequestBApp, kn) {
					currentKey = ffjtFF_BidRequestBApp
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyFF_BidRequestBAdv, kn) {
					currentKey = ffjtFF_BidRequestBAdv
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyFF_BidRequestBcat, kn) {
					currentKey = ffjtFF_BidRequestBcat
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyFF_BidRequestCur, kn) {
					currentKey = ffjtFF_BidRequestCur
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyFF_BidRequestAllImps, kn) {
					currentKey = ffjtFF_BidRequestAllImps
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyFF_BidRequestWLang, kn) {
					currentKey = ffjtFF_BidRequestWLang
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyFF_BidRequestBSeat, kn) {
					currentKey = ffjtFF_BidRequestBSeat
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyFF_BidRequestWSeat, kn) {
					currentKey = ffjtFF_BidRequestWSeat
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyFF_BidRequestTMax, kn) {
					currentKey = ffjtFF_BidRequestTMax
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyFF_BidRequestAuctionType, kn) {
					currentKey = ffjtFF_BidRequestAuctionType
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyFF_BidRequestTest, kn) {
					currentKey = ffjtFF_BidRequestTest
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyFF_BidRequestUser, kn) {
					currentKey = ffjtFF_BidRequestUser
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyFF_BidRequestDevice, kn) {
					currentKey = ffjtFF_BidRequestDevice
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyFF_BidRequestApp, kn) {
					currentKey = ffjtFF_BidRequestApp
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyFF_BidRequestSite, kn) {
					currentKey = ffjtFF_BidRequestSite
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyFF_BidRequestImp, kn) {
					currentKey = ffjtFF_BidRequestImp
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyFF_BidRequestID, kn) {
					currentKey = ffjtFF_BidRequestID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtFF_BidRequestnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtFF_BidRequestID:
					goto handle_ID

				case ffjtFF_BidRequestImp:
					goto handle_Imp

				case ffjtFF_BidRequestSite:
					goto handle_Site

				case ffjtFF_BidRequestApp:
					goto handle_App

				case ffjtFF_BidRequestDevice:
					goto handle_Device

				case ffjtFF_BidRequestUser:
					goto handle_User

				case ffjtFF_BidRequestTest:
					goto handle_Test

				case ffjtFF_BidRequestAuctionType:
					goto handle_AuctionType

				case ffjtFF_BidRequestTMax:
					goto handle_TMax

				case ffjtFF_BidRequestWSeat:
					goto handle_WSeat

				case ffjtFF_BidRequestBSeat:
					goto handle_BSeat

				case ffjtFF_BidRequestWLang:
					goto handle_WLang

				case ffjtFF_BidRequestAllImps:
					goto handle_AllImps

				case ffjtFF_BidRequestCur:
					goto handle_Cur

				case ffjtFF_BidRequestBcat:
					goto handle_Bcat

				case ffjtFF_BidRequestBAdv:
					goto handle_BAdv

				case ffjtFF_BidRequestBApp:
					goto handle_BApp

				case ffjtFF_BidRequestSource:
					goto handle_Source

				case ffjtFF_BidRequestRegs:
					goto handle_Regs

				case ffjtFF_BidRequestExt:
					goto handle_Ext

				case ffjtFF_BidRequestPmp:
					goto handle_Pmp

				case ffjtFF_BidRequestnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_ID:

	/* handler: j.ID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.ID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Imp:

	/* handler: j.Imp type=[]openrtb.Impression kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Imp = nil
		} else {

			j.Imp = []openrtb.Impression{}

			wantVal := true

			for {

				var tmpJImp openrtb.Impression

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJImp type=openrtb.Impression kind=struct quoted=false*/

				{
					/* Falling back. type=openrtb.Impression kind=struct */
					tbuf, err := fs.CaptureField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}

					err = json.Unmarshal(tbuf, &tmpJImp)
					if err != nil {
						return fs.WrapErr(err)
					}
				}

				j.Imp = append(j.Imp, tmpJImp)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Site:

	/* handler: j.Site type=openrtb.Site kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.Site kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Site)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_App:

	/* handler: j.App type=openrtb.App kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.App kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.App)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Device:

	/* handler: j.Device type=openrtb.Device kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.Device kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Device)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_User:

	/* handler: j.User type=openrtb.User kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.User kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.User)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Test:

	/* handler: j.Test type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.Test = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_AuctionType:

	/* handler: j.AuctionType type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.AuctionType = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TMax:

	/* handler: j.TMax type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.TMax = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_WSeat:

	/* handler: j.WSeat type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.WSeat = nil
		} else {

			j.WSeat = []string{}

			wantVal := true

			for {

				var tmpJWSeat string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJWSeat type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJWSeat = string(string(outBuf))

					}
				}

				j.WSeat = append(j.WSeat, tmpJWSeat)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BSeat:

	/* handler: j.BSeat type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.BSeat = nil
		} else {

			j.BSeat = []string{}

			wantVal := true

			for {

				var tmpJBSeat string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJBSeat type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJBSeat = string(string(outBuf))

					}
				}

				j.BSeat = append(j.BSeat, tmpJBSeat)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_WLang:

	/* handler: j.WLang type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.WLang = nil
		} else {

			j.WLang = []string{}

			wantVal := true

			for {

				var tmpJWLang string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJWLang type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJWLang = string(string(outBuf))

					}
				}

				j.WLang = append(j.WLang, tmpJWLang)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_AllImps:

	/* handler: j.AllImps type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.AllImps = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Cur:

	/* handler: j.Cur type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Cur = nil
		} else {

			j.Cur = []string{}

			wantVal := true

			for {

				var tmpJCur string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJCur type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJCur = string(string(outBuf))

					}
				}

				j.Cur = append(j.Cur, tmpJCur)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Bcat:

	/* handler: j.Bcat type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Bcat = nil
		} else {

			j.Bcat = []string{}

			wantVal := true

			for {

				var tmpJBcat string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJBcat type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJBcat = string(string(outBuf))

					}
				}

				j.Bcat = append(j.Bcat, tmpJBcat)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BAdv:

	/* handler: j.BAdv type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.BAdv = nil
		} else {

			j.BAdv = []string{}

			wantVal := true

			for {

				var tmpJBAdv string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJBAdv type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJBAdv = string(string(outBuf))

					}
				}

				j.BAdv = append(j.BAdv, tmpJBAdv)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BApp:

	/* handler: j.BApp type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.BApp = nil
		} else {

			j.BApp = []string{}

			wantVal := true

			for {

				var tmpJBApp string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJBApp type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJBApp = string(string(outBuf))

					}
				}

				j.BApp = append(j.BApp, tmpJBApp)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Source:

	/* handler: j.Source type=openrtb.Source kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.Source kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Source)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Regs:

	/* handler: j.Regs type=openrtb.Regulations kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.Regulations kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Regs)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Ext:

	/* handler: j.Ext type=openrtb.Extension kind=slice quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Ext.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Pmp:

	/* handler: j.Pmp type=openrtb.Pmp kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.Pmp kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Pmp)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
