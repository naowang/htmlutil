// htmlutil project htmlutil.go
package htmlutil

import (
	"bytes"
	"charset"
	"crypto/md5"
	"encoding/hex"

	//"fmt"
	"sync"
	"time"
	"toolfunc"

	//"fmt"
	"crypto/sha1"
	"encoding/base64"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var MAXHTMLNAMETAG int = 6

var specialcharmap sync.Map

var htmlutilinitval bool = false
var urlspaceregex *regexp.Regexp

func init() {
	htmlutilinitval = true
	speccharls := []string{"&aacute;", "&acirc;", "&acute;", "&aelig;", "&agrave;", "&alefsym;", "&alpha;", "&amp;", "&and;", "&ang;", "&aring;", "&asymp;", "&atilde;", "&auml;", "&bdquo;", "&beta;", "&brvbar;", "&bull;", "&cap;", "&ccedil;", "&cedil;", "&cent;", "&chi;", "&circ;", "&clubs;", "&cong;", "&copy;", "&crarr;", "&cup;", "&curren;", "&dagger;", "&darr;", "&deg;", "&delta;", "&diams;", "&divide;", "&eacute;", "&ecirc;", "&egrave;", "&empty;", "&emsp;", "&ensp;", "&epsilon;", "&equiv;", "&eta;", "&eth;", "&euml;", "&exist;", "&fnof;", "&forall;", "&frac12;", "&frac14;", "&frac34;", "&frasl;", "&gamma;", "&ge;", "&gt;", "&harr;", "&hearts;", "&hellip;", "&iacute;", "&icirc;", "&iexcl;", "&igrave;", "&image;", "&infin;", "&int;", "&iota;", "&iquest;", "&isin;", "&iuml;", "&kappa;", "&lambda;", "&laquo;", "&larr;", "&lceil;", "&ldquo;", "&le;", "&lfloor;", "&lowast;", "&loz;", "&lrm;", "&lsaquo;", "&lsquo;", "&lt;", "&macr;", "&mdash;", "&micro;", "&middot;", "&minus;", "&mu;", "&nabla;", "&nbsp;", "&ndash;", "&ne;", "&ni;", "&not;", "&notin;", "&nsub;", "&ntilde;", "&nu;", "&oacute;", "&ocirc;", "&oelig;", "&ograve;", "&oline;", "&omega;", "&omicron;", "&oplus;", "&or;", "&ordf;", "&ordm;", "&oslash;", "&otilde;", "&otimes;", "&ouml;", "&para;", "&part;", "&permil;", "&perp;", "&phi;", "&pi;", "&piv;", "&plusmn;", "&pound;", "&prime;", "&prod;", "&prop;", "&psi;", "&quot;", "&radic;", "&raquo;", "&rarr;", "&rceil;", "&rdquo;", "&real;", "&reg;", "&rfloor;", "&rho;", "&rsaquo;", "&rsquo;", "&sbquo;", "&scaron;", "&sdot;", "&sect;", "&shy;", "&sigma;", "&sigmaf;", "&sim;", "&spades;", "&sub;", "&sube;", "&sum;", "&sup1;", "&sup2;", "&sup3;", "&sup;", "&supe;", "&szlig;", "&tau;", "&there4;", "&theta;", "&thetasym;", "&thinsp;", "&thorn;", "&tilde;", "&times;", "&trade;", "&uacute;", "&uarr;", "&ucirc;", "&ugrave;", "&uml;", "&upsih;", "&upsilon;", "&uuml;", "&weierp;", "&xi;", "&yacute;", "&yen;", "&yuml;", "&zeta;", "&zwj;", "&zwnj;", "&lang;", "&rang;", "&apos;", "&euro;", "&rlm;"}
	spechar_sym := []string{"á", "â", "´", "æ", "à", "ℵ", "α", "&", "∧", "∠", "å", "≈", "ã", "ä", "„", "β", "¦", "•", "∩", "ç", "¸", "¢", "χ", "ˆ", "♣", "≅", "©", "↵", "∪", "¤", "†", "↓", "°", "δ", "♦", "÷", "é", "ê", "è", "∅", " ", " ", "ε", "≡", "η", "ð", "ë", "∃", "ƒ", "∀", "½", "¼", "¾", "⁄", "γ", "≥", ">", "↔", "♥", "…", "í", "î", "¡", "ì", "ℑ", "∞", "∫", "ι", "¿", "∈", "ï", "κ", "λ", "«", "←", "⌈", "“", "≤", "⌊", "∗", "◊", "‎", "‹", "‘", "<", "¯", "—", "µ", "·", "−", "μ", "∇", " ", "–", "≠", "∋", "¬", "∉", "⊄", "ñ", "ν", "ó", "ô", "œ", "ò", "‾", "ω", "ο", "⊕", "∨", "ª", "º", "ø", "õ", "⊗", "ö", "¶", "∂", "‰", "⊥", "φ", "π", "ϖ", "±", "£", "′", "∏", "∝", "ψ", "\"", "√", "»", "→", "⌉", "”", "ℜ", "®", "⌋", "ρ", "›", "’", "‚", "š", "⋅", "§", "­", "σ", "ς", "∼", "♠", "⊂", "⊆", "∑", "¹", "²", "³", "⊃", "⊇", "ß", "τ", "∴", "θ", "ϑ", " ", "þ", "˜", "×", "™", "ú", "↑", "û", "ù", "¨", "ϒ", "υ", "ü", "℘", "ξ", "ý", "¥", "ÿ", "ζ", "‍", "‌", "⟨", "⟩", "'", "€", "‏"}
	if len(speccharls) != len(spechar_sym) {
		panic("charset error!")
	}
	for i := 0; i < len(speccharls); i++ {
		specialcharmap.Store(string(speccharls[i]), string(spechar_sym[i]))
		if len(speccharls[i]) > MAXHTMLNAMETAG {
			MAXHTMLNAMETAG = len(speccharls[i])
		}
	}
	urlspaceregex = regexp.MustCompile("[ \r\n\t]")
	return
}

func GetTitle(html []byte) []byte {
	re := regexp.MustCompile("(?ism)<title[^>]*>([^<]*)</title>")
	foundtit := re.FindAll(html, 1)
	if len(foundtit) >= 1 {
		return foundtit[0][bytes.Index(foundtit[0], []byte{'>'})+1 : len(foundtit[0])-8]
	}
	return []byte{}
}

func HtmlToText(in_text_u8 []byte) []byte {
	if !htmlutilinitval {
		panic("htmlutil is not initialize.")
	}
	out_text_u8 := make([]byte, 0, 4096)
	for i := 0; i < len(in_text_u8); i++ {
		switch in_text_u8[i] {
		case '<':
			i++
			tagname := make([]byte, 0, 12)
			var bendtag bool
			if bytes.HasPrefix(in_text_u8[i:], []byte("!--")) {
				i = i + bytes.Index(in_text_u8[i:], []byte("-->")) + 2
			} else {
				if i < len(in_text_u8) && bytes.HasPrefix(in_text_u8[i:], []byte("/")) {
					bendtag = true
					i += 1
				}
				for i < len(in_text_u8) && !unicode.IsLetter(rune(in_text_u8[i])) {
					i++
				}
				for i < len(in_text_u8) && unicode.IsLetter(rune(in_text_u8[i])) {
					tagname = append(tagname, byte(unicode.ToLower(rune(in_text_u8[i]))))
					i++
				}
				bequalpass := false
				for i < len(in_text_u8) {
					if in_text_u8[i] == '=' {
						i++
						bequalpass = true
					} else if bequalpass && in_text_u8[i] == '\'' {
						i++
						for i < len(in_text_u8) && in_text_u8[i] != '\'' {
							i++
						}
						if i < len(in_text_u8) && in_text_u8[i] == '\'' {
							i++
						}
						bequalpass = false
					} else if bequalpass && in_text_u8[i] == '"' {
						i++
						for i < len(in_text_u8) && in_text_u8[i] != '"' {
							i++
						}
						if i < len(in_text_u8) && in_text_u8[i] == '"' {
							i++
						}
						bequalpass = false
					} else if in_text_u8[i] != '>' {
						i++
					} else {
						break
					}
				}
			}
			if bendtag == false && (string(tagname) == "style" || string(tagname) == "script") {
				for i < len(in_text_u8) {
					if i+8 < len(in_text_u8) && toolfunc.Byte1DCompareInsensitive(in_text_u8[i:i+8], []byte("</style>")) || i+9 < len(in_text_u8) && toolfunc.Byte1DCompareInsensitive(in_text_u8[i:i+9], []byte("</script>")) { // || i+11 < len(in_text_u8) && toolfunc.Byte1DCompareInsensitive(in_text_u8[i:i+11], []byte("</textarea>")) {
						i = i + bytes.Index(in_text_u8[i:], []byte{'>'})
						break
					} else {
						i++
					}
				}
			}
			if bendtag == true {
				if string(tagname) == "td" {
					out_text_u8 = append(out_text_u8, '\t')
				}
				if string(tagname) == "tr" || string(tagname) == "p" || string(tagname) == "div" {
					out_text_u8 = append(out_text_u8, '\n')
				}
				if string(tagname) == "span" {
					out_text_u8 = append(out_text_u8, ' ')
				}
			}
			continue
		case '>':
			continue
		case '&':
			htmlsymval := make([]byte, 0)
			if i+1 < len(in_text_u8) && in_text_u8[i+1] == '#' {
				if i+2 < len(in_text_u8) && unicode.ToLower(rune(in_text_u8[i+2])) == 'x' {
					htmlsymval = make([]byte, 0)
					htmlsymval = append(htmlsymval, []byte("&#x")...)
					i += 3
					for i < len(in_text_u8) {
						if unicode.ToLower(rune(in_text_u8[i])) >= 'a' && unicode.ToLower(rune(in_text_u8[i])) <= 'f' || unicode.ToLower(rune(in_text_u8[i])) >= '0' && unicode.ToLower(rune(in_text_u8[i])) <= '9' {
							htmlsymval = append(htmlsymval, byte(unicode.ToLower(rune(in_text_u8[i]))))
							i++
						} else {
							if in_text_u8[i] == ';' {
								htmlsymval = append(htmlsymval, byte(unicode.ToLower(rune(in_text_u8[i]))))
								i++
							}
							break
						}
						if len(htmlsymval) > MAXHTMLNAMETAG {
							break
						}
					}

					spechval, bspechval := specialcharmap.Load(string(htmlsymval))
					if bspechval {
						out_text_u8 = append(out_text_u8, []byte(spechval.(string))...)
					} else {
						indnumsym := bytes.Index(htmlsymval, []byte{byte('#')})
						indsemicolon := bytes.Index(htmlsymval, []byte{byte(';')})
						if indnumsym != -1 && indsemicolon != -1 && indnumsym < indsemicolon {
							hexnum := htmlsymval[indnumsym+2 : indsemicolon]
							ucode, _ := strconv.ParseInt(string(hexnum), 16, 32)
							ucode2 := rune(ucode)
							out_text_u8 = append(out_text_u8, []byte(string(ucode2))...)
						} else {
							out_text_u8 = append(out_text_u8, []byte(htmlsymval)...)
						}
					}
				} else {
					htmlsymval = make([]byte, 0, 8)
					htmlsymval = append(htmlsymval, []byte("&#")...)
					i += 2
					for i < len(in_text_u8) {
						if unicode.ToLower(rune(in_text_u8[i])) >= '0' && unicode.ToLower(rune(in_text_u8[i])) <= '9' {
							htmlsymval = append(htmlsymval, byte(unicode.ToLower(rune(in_text_u8[i]))))
							i++
						} else {
							if in_text_u8[i] == ';' {
								htmlsymval = append(htmlsymval, byte(unicode.ToLower(rune(in_text_u8[i]))))
								i++
							}
							break
						}
						if len(htmlsymval) > MAXHTMLNAMETAG {
							break
						}
					}

					spechrval, bspechrval := specialcharmap.Load(string(htmlsymval))
					if bspechrval {
						out_text_u8 = append(out_text_u8, spechrval.(string)...)
					} else {
						if bytes.Index(htmlsymval, []byte{byte('#')}) != -1 && bytes.Index(htmlsymval, []byte{byte(';')}) != -1 {
							codenum := htmlsymval[bytes.Index(htmlsymval, []byte{byte('#')})+1 : bytes.Index(htmlsymval, []byte{byte(';')})]
							ucode, _ := strconv.ParseInt(string(codenum), 10, 32)
							ucode2 := rune(ucode)
							out_text_u8 = append(out_text_u8, []byte(string(ucode2))...)
						}
					}
				}
			} else {
				htmlsymval = make([]byte, 0)
				htmlsymval = append(htmlsymval, '&')
				i++
				for i < len(in_text_u8) {
					if in_text_u8[i] >= 0x61 && in_text_u8[i] <= 0x7a || in_text_u8[i] >= 0x41 && in_text_u8[i] <= 0x5a || in_text_u8[i] >= 0x30 && in_text_u8[i] <= 0x39 {
						htmlsymval = append(htmlsymval, byte(unicode.ToLower(rune(in_text_u8[i]))))
						i++
						if len(htmlsymval) > MAXHTMLNAMETAG {
							break
						}
					} else {
						if in_text_u8[i] == ';' {
							htmlsymval = append(htmlsymval, byte(unicode.ToLower(rune(in_text_u8[i]))))
							i++
						}
						break
					}
				}
				spechrval, bspechrval := specialcharmap.Load(string(htmlsymval))
				if bspechrval {
					out_text_u8 = append(out_text_u8, []byte(spechrval.(string))...)
				} else {
					out_text_u8 = append(out_text_u8, htmlsymval...)
				}
			}
			i-- //避免for ++ 越过
			break
		default:
			out_text_u8 = append(out_text_u8, in_text_u8[i])
		}
	}
	return out_text_u8
}

func ToFullUrl(urlstr, pageurl string) string {
	var checkprotolen int = 12
	if len(urlstr) < checkprotolen {
		checkprotolen = len(urlstr)
	}
	protocolhead := ""
	for i := 0; i < checkprotolen; i++ {
		if urlstr[i] >= 'a' && urlstr[i] <= 'z' || urlstr[i] >= 'A' && urlstr[i] <= 'Z' || urlstr[i] >= '0' && urlstr[i] <= '9' {
			//
		} else if urlstr[i] == ':' {
			protocolhead = urlstr[:i]
			break
		} else {
			break
		}
	}
	if protocolhead != "" { //protocolhead maybe:javascript,mailto,ed2k,ftp,thunder
		protocolhead = strings.ToLower(protocolhead)
		if protocolhead == "javascript" || protocolhead == "mailto" || protocolhead == "tel" || protocolhead == "skype" || protocolhead == "about" || protocolhead == "sms" || protocolhead == "fax" || protocolhead == "tencent" || protocolhead == "file" || protocolhead == "javasctipt" || protocolhead == "javascrip" || protocolhead == "avascript" || protocolhead == "javacript" || protocolhead == "line" || protocolhead == "steam" {
			return ""
		} else if protocolhead == "http" || protocolhead == "https" {
			//
		} else {
			return urlstr
		}
	}

	//pageurl less than 8 panic is for check error
	if strings.Index(pageurl[8:], "/") == -1 {
		pageurl += "/"
	}
	if len(urlstr) == 0 {
		return urlstr
	}
	if urlstr[0] == '#' {
		return pageurl
	}
	if strings.Index(urlstr, "#") != -1 {
		urlstr = urlstr[:strings.Index(urlstr, "#")]
	}
	if strings.HasPrefix(urlstr, "//") {
		urlstr2 := pageurl[0:strings.Index(pageurl, "/")] + urlstr
		if len(urlstr) < 8 {
			return ""
		}
		if strings.Index(urlstr2[8:], "/") == -1 {
			urlstr2 += "/"
		}
		return toolfunc.UrlEncode(urlstr2)
	}
	if len(urlstr) >= 7 && strings.ToLower(urlstr[0:7]) == "http://" {
		urlstr = strings.Replace(urlstr, "//", "/", -1)
		urlstr = strings.Replace(urlstr, "http:/", "http://", -1)
		if len(urlstr) < 8 {
			return ""
		}
		if strings.Index(urlstr[8:], "/") == -1 {
			urlstr += "/"
		}
		if len(urlstr) < 10 {
			return ""
		}
		urlstr = strings.ToLower(urlstr[0:8+strings.Index(urlstr[8:], "/")]) + urlstr[8+strings.Index(urlstr[8:], "/"):]
		return toolfunc.UrlEncode(urlstr)
	}
	if len(urlstr) >= 8 && strings.ToLower(urlstr[0:8]) == "https://" {
		urlstr = strings.Replace(urlstr, "//", "/", -1)
		urlstr = strings.Replace(urlstr, "//", "/", -1)
		urlstr = strings.Replace(urlstr, "//", "/", -1)
		urlstr = strings.Replace(urlstr, "https:/", "https://", -1)

		if strings.Index(urlstr[8:], "/") == -1 {
			urlstr += "/"
		}
		if len(urlstr) < 11 {
			return ""
		}
		urlstr = strings.ToLower(urlstr[0:strings.Index(urlstr[8:], "/")]) + urlstr[strings.Index(urlstr[8:], "/"):]
		return toolfunc.UrlEncode(urlstr)
	}
	if urlstr == "" {
		return ""
	}
	if urlstr[0] == '.' {
		brightslash := false
		if strings.HasSuffix(urlstr, "/") {
			brightslash = true
		}
		urlpartls := strings.Split(urlstr, "/")
		if strings.Index(pageurl, "?") != -1 {
			pageurl = pageurl[:strings.Index(pageurl, "?")]
		}
		ls2 := strings.Split(pageurl, "/")
		if len(ls2)-1 >= 0 {
			ls2 = ls2[0 : len(ls2)-1]
		}
		for i := 0; i < len(urlpartls); i++ {
			if urlpartls[i] == "." || urlpartls[i] == "" {

			} else if urlpartls[i] == ".." {
				if len(ls2) > 3 {
					ls2 = ls2[0 : len(ls2)-1]
				}
			} else {
				ls2 = append(ls2, urlpartls[i])
			}
		}
		urlstr2 := strings.Join(ls2, "/")
		urlstr2 = strings.Replace(urlstr2, "//", "/", -1)
		urlstr2 = strings.Replace(urlstr2, "//", "/", -1)
		urlstr2 = strings.Replace(urlstr2, "//", "/", -1)
		urlstr2 = strings.Replace(urlstr2, "http:/", "http://", -1)
		urlstr2 = strings.Replace(urlstr2, "https:/", "https://", -1)
		if len(urlstr2) < 11 {
			return ""
		}
		if strings.Index(urlstr2[8:], "/") == -1 {
			urlstr2 += "/"
		}
		if len(urlstr2) > 8 && strings.Index(urlstr2[8:], "/") != -1 {
			urlstr2 = strings.ToLower(urlstr2[0:8+strings.Index(urlstr2[8:], "/")]) + urlstr2[8+strings.Index(urlstr2[8:], "/"):]
			if brightslash {
				if strings.HasSuffix(urlstr2, "/") == false {
					urlstr2 += "/"
				}
			}
			return toolfunc.UrlEncode(urlstr2)
		} else {
			if brightslash {
				if strings.HasSuffix(urlstr2, "/") == false {
					urlstr2 += "/"
				}
			}
			return toolfunc.UrlEncode(urlstr2)
		}
	} else if urlstr[0] == '/' {
		urlstr = strings.Replace(urlstr, "//", "/", -1)
		urlstr2 := pageurl[:8] + pageurl[8:][:strings.Index(pageurl[8:], "/")] + urlstr
		if len(urlstr2) > 8 && strings.Index(urlstr2[8:], "/") != -1 {
			urlstr2 = strings.ToLower(urlstr2[0:8+strings.Index(urlstr2[8:], "/")]) + urlstr2[8+strings.Index(urlstr2[8:], "/"):]
			return toolfunc.UrlEncode(urlstr2)
		} else {
			return toolfunc.UrlEncode(urlstr2)
		}
	} else {
		if strings.Index(pageurl, "?") != -1 {
			pageurl = pageurl[:strings.Index(pageurl, "?")]
		}
		ls2 := strings.Split(pageurl, "/")
		for i := len(ls2) - 1; i >= 0; i-- {
			if ls2[i] == "/" {
				ls2 = append(ls2[0:i], ls2[i+1:]...)
			}
		}
		if ls2[len(ls2)-1] != "" {
			ls2 = ls2[0 : len(ls2)-1]
		}
		ls2 = append(ls2, urlstr)
		urlstr2 := strings.Join(ls2, "/")
		urlstr2 = strings.Replace(urlstr2, "//", "/", -1)
		urlstr2 = strings.Replace(urlstr2, "//", "/", -1)
		urlstr2 = strings.Replace(urlstr2, "//", "/", -1)
		urlstr2 = strings.Replace(urlstr2, "http:/", "http://", -1)
		urlstr2 = strings.Replace(urlstr2, "https:/", "https://", -1)
		if strings.Index(urlstr2[8:], "/") == -1 {
			urlstr2 += "/"
		}
		if len(urlstr2) > 8 && strings.Index(urlstr2[8:], "/") != -1 {
			urlstr2 = strings.ToLower(urlstr2[0:8+strings.Index(urlstr2[8:], "/")]) + urlstr2[8+strings.Index(urlstr2[8:], "/"):]
			return toolfunc.UrlEncode(urlstr2)
		} else {
			return toolfunc.UrlEncode(urlstr2)
		}
	}
	return toolfunc.UrlEncode(urlstr)
}

func WiseGetAttr(attrstr, tagurlattrname string) string {
	urlstr := attrstr
	urlstr = strings.Trim(urlstr[strings.Index(urlstr, tagurlattrname)+len(tagurlattrname):], "\r\n\t ")
	if len(urlstr) > 0 && urlstr[0] == '=' {
		urlstr = strings.Trim(urlstr[1:], "\r\n\t ")
		if len(urlstr) > 0 {
			if urlstr[0] == '\'' {
				urlstr = urlstr[1:]
				if strings.Index(urlstr, "'") != -1 {
					urlstr = urlstr[0:strings.Index(urlstr, "'")]
				} else if strings.Index(urlstr, " ") != -1 {
					urlstr = urlstr[0:strings.Index(urlstr, " ")]
				} else if strings.Index(urlstr, ">") != -1 {
					urlstr = urlstr[0:strings.Index(urlstr, ">")]
				}
			} else if urlstr[0] == '"' {
				urlstr = urlstr[1:]
				if strings.Index(urlstr, "\"") != -1 {
					urlstr = urlstr[0:strings.Index(urlstr, "\"")]
				} else if strings.Index(urlstr, " ") != -1 {
					urlstr = urlstr[0:strings.Index(urlstr, " ")]
				} else if strings.Index(urlstr, ">") != -1 {
					urlstr = urlstr[0:strings.Index(urlstr, ">")]
				}
			} else {
				if strings.Index(urlstr, " ") != -1 && strings.Index(urlstr, ">") != -1 {
					if strings.Index(urlstr, " ") < strings.Index(urlstr, ">") {
						urlstr = urlstr[0:strings.Index(urlstr, " ")]
					} else {
						urlstr = urlstr[0:strings.Index(urlstr, ">")]
					}
				} else if strings.Index(urlstr, " ") != -1 {
					urlstr = urlstr[0:strings.Index(urlstr, " ")]
				} else if strings.Index(urlstr, ">") != -1 {
					urlstr = urlstr[0:strings.Index(urlstr, ">")]
				}
			}
		}
		return urlstr
	} else {
		return ""
	}
}

func HtmlFindUrlWithTag(pageurl, pagehtml, tagregex, tagurlattrname string) [][]string {
	out_newurl := [][]string{}
	newpagectt := pagehtml
	fromtagname := tagregex[strings.Index(tagregex, "<")+1:]
	for i := 0; i < len(fromtagname); i++ {
		if !(fromtagname[i] >= 0x61 && fromtagname[i] <= 0x7a || fromtagname[i] >= 0x41 && fromtagname[i] <= 0x5a) {
			fromtagname = fromtagname[:i]
			break
		}
	}
	fromtagname = strings.ToLower(fromtagname)
	//loop find all link and replace
	atagre := regexp.MustCompile(tagregex)
	spacere := regexp.MustCompile("[ \r\n\t]+")
	maes := atagre.FindAllSubmatchIndex([]byte(newpagectt), -1)
	for i := len(maes) - 1; i >= 0; i -= 1 {
		mastr := pagehtml[maes[i][2]:maes[i][3]]
		urlstr := mastr
		urlstr = strings.Trim(urlstr[strings.Index(urlstr, tagurlattrname)+len(tagurlattrname):], "\r\n\t ")
		if urlstr[0] == '=' {
			urlstr = strings.Trim(urlstr[1:], "\r\n\t ")
			if len(urlstr) > 0 {
				if urlstr[0] == '\'' {
					urlstr = urlstr[1:]
					if strings.Index(urlstr, "'") != -1 {
						urlstr = urlstr[0:strings.Index(urlstr, "'")]
					} else if strings.Index(urlstr, " ") != -1 {
						urlstr = urlstr[0:strings.Index(urlstr, " ")]
					} else if strings.Index(urlstr, ">") != -1 {
						urlstr = urlstr[0:strings.Index(urlstr, ">")]
					}
				} else if urlstr[0] == '"' {
					urlstr = urlstr[1:]
					if strings.Index(urlstr, "\"") != -1 {
						urlstr = urlstr[0:strings.Index(urlstr, "\"")]
					} else if strings.Index(urlstr, " ") != -1 {
						urlstr = urlstr[0:strings.Index(urlstr, " ")]
					} else if strings.Index(urlstr, ">") != -1 {
						urlstr = urlstr[0:strings.Index(urlstr, ">")]
					}
				} else {
					if strings.Index(urlstr, " ") != -1 && strings.Index(urlstr, ">") != -1 {
						if strings.Index(urlstr, " ") < strings.Index(urlstr, ">") {
							urlstr = urlstr[0:strings.Index(urlstr, " ")]
						} else {
							urlstr = urlstr[0:strings.Index(urlstr, ">")]
						}
					} else if strings.Index(urlstr, " ") != -1 {
						urlstr = urlstr[0:strings.Index(urlstr, " ")]
					} else if strings.Index(urlstr, ">") != -1 {
						urlstr = urlstr[0:strings.Index(urlstr, ">")]
					}
				}
				if len(urlstr) > 0 {
					fullurl := ToFullUrl(urlstr, pageurl)
					if len(fullurl) < 512 && strings.HasPrefix(fullurl, "http://") || strings.HasPrefix(fullurl, "https://") {
						//if len(fullurl) > 400 {
						//	fmt.Println("url too long error.")
						//}
						noanchorurl := fullurl
						if strings.Index(noanchorurl, "#") != -1 {
							noanchorurl = noanchorurl[0:strings.Index(noanchorurl, "#")]
						}
						urltitle := ""
						urltime := ""
						if strings.HasPrefix(tagregex, "(?ism)<a") {
							totalstr := pagehtml[maes[i][0]:maes[i][1]]
							allind := regexp.MustCompile("<[^>]*>(.*?)</a>").FindAllSubmatchIndex([]byte(totalstr), -1)
							if len(allind) > 0 {
								urltitle = string(HtmlToText([]byte(totalstr[allind[0][2]:allind[0][3]])))
							}
							urltitle2 := toolfunc.GetRegexGroup1("title\\s*=\\s*['\"](.*?)['\"]", totalstr)
							if len(urltitle2) > len(urltitle) || len(urltitle2) > 0 && strings.HasSuffix(urltitle, "...") {
								urltitle = string(HtmlToText([]byte(urltitle2)))
							}
						} else if strings.HasPrefix(tagregex, "(?ism)<img") {
							totalstr := pagehtml[maes[i][0]:maes[i][1]]
							urltitle = toolfunc.GetRegexGroup1("title\\s*=\\s*['\"](.*?)['\"]", totalstr)
							if urltitle == "" {
								urltitle = toolfunc.GetRegexGroup1("alt\\s*=\\s*['\"](.*?)['\"]", totalstr)
							}
							urltitle = string(HtmlToText([]byte(urltitle)))
						}
						urltitle = strings.Trim(string(spacere.ReplaceAll([]byte(urltitle), []byte(" "))), " \r\n\t")
						linktype := ""
						if fromtagname == "source" {
							linktype = WiseGetAttr(mastr, "type")
						}
						out_newurl = append(out_newurl, []string{noanchorurl, urltitle, urltime, fromtagname, tagurlattrname, linktype})
					}
				}
			}
		}
	}
	return out_newurl
}

func GetAllTagALink(pageurl string, pagehtml []byte) []string {
	out_newurl := []string{}
	atagre := regexp.MustCompile("<a\\s+[^>]*href=([^> \\t]+)[^>]*>.*?</a>")
	maes := atagre.FindAllSubmatchIndex(pagehtml, -1)
	urlset := make(map[string]int8)
	for i := len(maes) - 1; i >= 0; i -= 1 {
		urlstr := string(pagehtml[maes[i][2]:maes[i][3]])
		if urlstr[0] == '\'' {
			urlstr = urlstr[1:]
			if strings.Index(urlstr, "'") != -1 {
				urlstr = urlstr[0:strings.Index(urlstr, "'")]
			} else if strings.Index(urlstr, " ") != -1 {
				urlstr = urlstr[0:strings.Index(urlstr, " ")]
			} else if strings.Index(urlstr, ">") != -1 {
				urlstr = urlstr[0:strings.Index(urlstr, ">")]
			}
		} else if urlstr[0] == '"' {
			urlstr = urlstr[1:]
			if strings.Index(urlstr, "\"") != -1 {
				urlstr = urlstr[0:strings.Index(urlstr, "\"")]
			} else if strings.Index(urlstr, " ") != -1 {
				urlstr = urlstr[0:strings.Index(urlstr, " ")]
			} else if strings.Index(urlstr, ">") != -1 {
				urlstr = urlstr[0:strings.Index(urlstr, ">")]
			}
		}
		fullurl := ToFullUrl(urlstr, pageurl)
		if strings.Index(fullurl, "#") != -1 {
			fullurl = fullurl[0:strings.Index(fullurl, "#")]
		}
		if strings.HasPrefix(fullurl, "http://") || strings.HasPrefix(fullurl, "https://") {
			_, be := urlset[fullurl]
			if be == false && len(fullurl) > 0 {
				out_newurl = append(out_newurl, fullurl)
				urlset[fullurl] = 1
			}
		}
	}
	return out_newurl
}

func GetAllTagALinkAndName(pageurl string, pagehtml []byte) (link, name []string) {
	out_newurl := []string{}
	out_newname := []string{}
	atagre := regexp.MustCompile("<a\\s+[^>]*href=([^> \\t]+)[^>]*>(.*?)</a>")
	maes := atagre.FindAllSubmatchIndex(pagehtml, -1)
	urlset := make(map[string]int8)
	for i := len(maes) - 1; i >= 0; i -= 1 {
		urlstr := pagehtml[maes[i][2]:maes[i][3]]
		namestr := bytes.Trim(HtmlToText(pagehtml[maes[i][4]:maes[i][5]]), " \r\n\t")
		if urlstr[0] == '\'' {
			urlstr = urlstr[1:]
			if bytes.Index(urlstr, []byte{'\''}) != -1 {
				urlstr = urlstr[0:bytes.Index(urlstr, []byte{'\''})]
			} else if bytes.Index(urlstr, []byte{' '}) != -1 {
				urlstr = urlstr[0:bytes.Index(urlstr, []byte{' '})]
			} else if bytes.Index(urlstr, []byte{'>'}) != -1 {
				urlstr = urlstr[0:bytes.Index(urlstr, []byte{'>'})]
			}
		} else if urlstr[0] == '"' {
			urlstr = urlstr[1:]
			if bytes.Index(urlstr, []byte{'"'}) != -1 {
				urlstr = urlstr[0:bytes.Index(urlstr, []byte{'"'})]
			} else if bytes.Index(urlstr, []byte{' '}) != -1 {
				urlstr = urlstr[0:bytes.Index(urlstr, []byte{' '})]
			} else if bytes.Index(urlstr, []byte{'>'}) != -1 {
				urlstr = urlstr[0:bytes.Index(urlstr, []byte{'>'})]
			}
		}
		fullurl := ToFullUrl(string(urlstr), pageurl)
		if strings.Index(fullurl, "#") != -1 {
			fullurl = fullurl[0:strings.Index(fullurl, "#")]
		}
		_, be := urlset[fullurl]
		if be == false && len(fullurl) > 0 {
			out_newurl = append(out_newurl, fullurl)
			out_newname = append(out_newname, string(namestr))
			urlset[fullurl] = 1
		}
	}
	return out_newurl, out_newname
}

func RemoveAllTagA(pagehtml []byte) []byte {
	scripttagre := regexp.MustCompile("<script\\s[^>]*>.*?</script>")
	maes2 := scripttagre.FindAllSubmatchIndex(pagehtml, -1)
	for i := len(maes2) - 1; i >= 0; i -= 1 {
		pagehtml = append(pagehtml[0:maes2[i][0]], pagehtml[maes2[i][1]:]...)
	}
	atagre := regexp.MustCompile("<a\\s+[^>]*href=([^> \\t]+)[^>]*>.*?</a>")
	maes := atagre.FindAllSubmatchIndex(pagehtml, -1)
	for i := len(maes) - 1; i >= 0; i -= 1 {
		pagehtml = append(pagehtml[0:maes[i][0]], pagehtml[maes[i][1]:]...)
	}
	return pagehtml
}

func RemoveSpace(pagehtml []byte) []byte {
	spacere := regexp.MustCompile("[ \r\n\t]{2,}")
	maes2 := spacere.FindAllSubmatchIndex(pagehtml, -1)
	for i := len(maes2) - 1; i >= 0; i -= 1 {
		pagehtml = append(pagehtml[0:maes2[i][0]], pagehtml[maes2[i][1]-1:]...)
		pagehtml[maes2[i][0]] = ' '
	}
	return pagehtml
}

func GetCharsetName(pagehtml []byte) string {
	i := bytes.Index(pagehtml, []byte("<body"))
	if i != -1 {
		pagehtml = pagehtml[:i]
	}
	charsetre := regexp.MustCompile("<meta\\s+charset=['\"]([^'\"']+)['\"]\\s*>")
	allindex := charsetre.FindAllSubmatchIndex(pagehtml, 1)
	if len(allindex) > 0 {
		return string(pagehtml[allindex[0][2]:allindex[0][3]])
	} else {
		charsetre := regexp.MustCompile("<meta\\s+[^>]*charset=([a-zA-Z0-9\\-]+)[^>]*>")
		allindex = charsetre.FindAllSubmatchIndex(pagehtml, 1)
		if len(allindex) > 0 {
			return string(pagehtml[allindex[0][2]:allindex[0][3]])
		}
		return ""
	}
}

func ConvertPageContentToUtf8(pagehtml []byte) []byte {
	codename := ""
	codename = GetCharsetName(pagehtml)
	if codename == "" {
		codename = charset.DetectCharset(pagehtml)
	}
	codename = strings.ToUpper(codename)
	if codename == "GBK2312" || codename == "GB2312" {
		codename = "GB18030"
	}
	if codename != "" && codename != "UTF-8" {
		pagehtml = []byte(charset.Convert(string(pagehtml), codename, "UTF-8"))
	}
	return pagehtml
}

func HtmlRemoveAllScript(pagectt string) string {
	for true {
		ind := regexp.MustCompile("(?ism:<script[^a-zA-Z].*?</script>)").FindStringIndex(pagectt)
		if len(ind) == 2 {
			pagectt = pagectt[0:ind[0]] + pagectt[ind[1]:]
		} else {
			break
		}
	}
	return pagectt
}

func HtmlRemoveAllScriptV2(pagectt []byte) []byte {
	ind := regexp.MustCompile("(?ism:<noscript[^>]*>.*?</noscript>)").FindIndex(pagectt)
	if len(ind) == 2 {
		pagectt = append(pagectt[0:ind[0]], pagectt[ind[1]:]...)
	}
	for true {
		scriptpos := bytes.Index(pagectt, []byte("<script"))
		if scriptpos != -1 {
			scriptposbackup := scriptpos
			bokbreak := false
			for scriptpos < len(pagectt) {
				if pagectt[scriptpos] == '"' {
					scriptpos += 1
					for scriptpos < len(pagectt) {
						if pagectt[scriptpos] == '\\' {
							scriptpos += 2
						} else if pagectt[scriptpos] != '"' {
							scriptpos += 1
						} else {
							scriptpos += 1
							break
						}
					}
				} else if pagectt[scriptpos] == '\'' {
					scriptpos += 1
					for scriptpos < len(pagectt) {
						if pagectt[scriptpos] == '\\' {
							scriptpos += 2
						} else if pagectt[scriptpos] != '\'' {
							scriptpos += 1
						} else {
							scriptpos += 1
							break
						}
					}
				} else if pagectt[scriptpos] == '`' {
					scriptpos += 1
					for scriptpos < len(pagectt) {
						if pagectt[scriptpos] == '\\' {
							scriptpos += 2
						} else if pagectt[scriptpos] != '`' {
							scriptpos += 1
						} else {
							scriptpos += 1
							break
						}
					}
				} else {
					scriptpos += 1
					if scriptpos+len("</script>") < len(pagectt) && bytes.Compare(pagectt[scriptpos:scriptpos+len("</script>")], []byte("</script>")) == 0 {
						pagectt = append(pagectt[:scriptposbackup], pagectt[scriptpos+len("</script>"):]...)
						bokbreak = true
						break
					}
				}
			}
			if bokbreak == false && scriptpos >= len(pagectt) {
				pagectt = pagectt[:scriptposbackup]
				break
			}
		} else {
			break
		}
	}
	return pagectt
}

func HtmlAllScriptIndex(pagectt []byte) [][]int {
	ind := regexp.MustCompile("(?ism:<script[^a-zA-Z].*?</script>)").FindAllSubmatchIndex(pagectt, -1)
	return ind
}

//return url index name:noanchorurl, urltitle, urltime, fromtagname, tagurlattrname
func HtmlFindAllUrl(pageurl, pagehtml string) [][]string {
	pagehtml = HtmlRemoveAllScript(pagehtml)
	tagreattr := []string{"(?ism)<a[^>]*(\\shref\\s*=[^>]*)>.*?</a>", "href", "(?ism)<img[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<link[^>]*(\\shref\\s*=[^>]*)>", "href", "(?ism)<script[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<frame[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<iframe[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<source[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<embed[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<object[^>]*(\\sdata\\s*=[^>]*)>", "data", "(?ism)<audio\\s[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<video\\s[^>]*(\\ssrc\\s*=[^>]*)>", "src"}
	tagreattri := 0
	totalnewurl := [][]string{}
	for true {
		//fmt.Println("HtmlFindAllUrl tagreattri:", tagreattri)
		newurl := HtmlFindUrlWithTag(pageurl, pagehtml, tagreattr[tagreattri], tagreattr[tagreattri+1])
		totalnewurl = append(totalnewurl, newurl...)
		tagreattri += 2
		if tagreattri >= len(tagreattr) {
			break
		}
	}
	return totalnewurl
}

func HtmlFindAllUrlOnlyAFrameAudioVideo(pageurl, pagehtml string) [][]string {
	pagehtml = HtmlRemoveAllScript(pagehtml)
	tagreattr := []string{"(?ism)<a[^>]*(\\shref\\s*=[^>]*)>.*?</a>", "href", "(?ism)<frame[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<iframe[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<source[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<embed[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<object[^>]*(\\sdata\\s*=[^>]*)>", "data", "(?ism)<audio\\s[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<video\\s[^>]*(\\ssrc\\s*=[^>]*)>", "src"}
	tagreattri := 0
	totalnewurl := [][]string{}
	for true {
		//fmt.Println("HtmlFindAllUrl tagreattri:", tagreattri)
		newurl := HtmlFindUrlWithTag(pageurl, pagehtml, tagreattr[tagreattri], tagreattr[tagreattri+1])
		totalnewurl = append(totalnewurl, newurl...)
		tagreattri += 2
		if tagreattri >= len(tagreattr) {
			break
		}
	}
	return totalnewurl
}

func UrlRepath(url string, relinkrootrereplacels, urlmatchurlextls []string) string {
	newpath := ""
	if len(relinkrootrereplacels) == 0 {
		h := md5.New()
		h.Write([]byte(url))
		cipherStr := h.Sum(nil)
		md5fname := hex.EncodeToString(cipherStr)
		ext := ""
		if len(urlmatchurlextls) > 0 {
			urlmatchurlextlsi := 0
			for true {
				if regexp.MustCompile(urlmatchurlextls[urlmatchurlextlsi]).Match([]byte(url)) {
					ext = urlmatchurlextls[urlmatchurlextlsi+1]
					if ext[0] != '.' {
						ext = "." + ext
					}
					break
				} else {
					urlmatchurlextlsi += 2
				}
				if urlmatchurlextlsi >= len(urlmatchurlextls) {
					break
				}
			}
		}
		if ext == "" {
			ext = url
			if strings.LastIndex(ext, ".") != -1 {
				ext = ext[strings.LastIndex(ext, "."):]
			} else {
				ext = ""
			}
			if strings.Index(ext, "?") != -1 {
				ext = ext[0:strings.Index(ext, "?")]
			}
			if ext == ".php" {
				ext = ".html"
			}
			ext = string(regexp.MustCompile("[\\?/\\\\:<>\\|*]").ReplaceAll([]byte(ext), []byte("_")))
			if len(ext) > 6 {
				ext = ext[0:6]
			}
		}

		newpath = md5fname + ext
	} else {
		newpath = string(regexp.MustCompile(relinkrootrereplacels[0]).ReplaceAll([]byte(url), []byte(relinkrootrereplacels[1])))
	}
	return newpath
}

func PageRelinkWithTag(pageurl string, pagehtml []byte, tagregex, tagurlattrname, relinkincludeurlre, relinkexcludeurlre string, relinkrootrereplacels, urlmatchurlextls []string) []byte {
	newpagectt := pagehtml
	pagehtmlscriptind := HtmlAllScriptIndex(pagehtml)
	//loop find all link and replace
	atagre := regexp.MustCompile(tagregex)
	maes := atagre.FindAllSubmatchIndex([]byte(newpagectt), -1)
	for i := len(maes) - 1; i >= 0; i -= 1 {
		mastr := pagehtml[maes[i][2]:maes[i][3]]

		binrange := false
		for _, srng := range pagehtmlscriptind {
			if maes[i][2] >= srng[0] && maes[i][3] <= srng[1] {
				binrange = true
			}
		}

		if binrange {
			continue
		}

		newmastr := []byte{}
		urlstr := mastr
		urlstr = bytes.Trim(urlstr[bytes.Index(urlstr, []byte(tagurlattrname))+len(tagurlattrname):], "\r\n\t ")
		if urlstr[0] == '=' {
			urlstr = bytes.Trim(urlstr[1:], "\r\n\t ")
			if urlstr[0] == '\'' {
				urlstr = urlstr[1:]
				if bytes.Index(urlstr, []byte{'\''}) != -1 {
					urlstr = urlstr[0:bytes.Index(urlstr, []byte{'\''})]
				} else if bytes.Index(urlstr, []byte{' '}) != -1 {
					urlstr = urlstr[0:bytes.Index(urlstr, []byte{' '})]
				} else if bytes.Index(urlstr, []byte{'>'}) != -1 {
					urlstr = urlstr[0:bytes.Index(urlstr, []byte{'>'})]
				}
			} else if urlstr[0] == '"' {
				urlstr = urlstr[1:]
				if bytes.Index(urlstr, []byte{'"'}) != -1 {
					urlstr = urlstr[0:bytes.Index(urlstr, []byte{'"'})]
				} else if bytes.Index(urlstr, []byte{' '}) != -1 {
					urlstr = urlstr[0:bytes.Index(urlstr, []byte{' '})]
				} else if bytes.Index(urlstr, []byte{'>'}) != -1 {
					urlstr = urlstr[0:bytes.Index(urlstr, []byte{'>'})]
				}
			} else {
				if bytes.Index(urlstr, []byte{' '}) != -1 && bytes.Index(urlstr, []byte{'>'}) != -1 {
					if bytes.Index(urlstr, []byte{' '}) < bytes.Index(urlstr, []byte{'>'}) {
						urlstr = urlstr[0:bytes.Index(urlstr, []byte{' '})]
					} else {
						urlstr = urlstr[0:bytes.Index(urlstr, []byte{'>'})]
					}
				} else if bytes.Index(urlstr, []byte{' '}) != -1 {
					urlstr = urlstr[0:bytes.Index(urlstr, []byte{' '})]
				} else if bytes.Index(urlstr, []byte{'>'}) != -1 {
					urlstr = urlstr[0:bytes.Index(urlstr, []byte{'>'})]
				}
			}
			if len(urlstr) > 0 {
				fullurl := ToFullUrl(string(urlstr), pageurl)
				if len(fullurl) >= 7 && toolfunc.StringCompareInsensitive(fullurl[:7], "http://") || len(fullurl) >= 8 && toolfunc.StringCompareInsensitive(fullurl[:8], "https://") {
					var bpass bool
					if relinkincludeurlre == "" && relinkexcludeurlre == "" {
						bpass = true
					} else {
						bpass = false
					}
					if relinkincludeurlre != "" {
						nurlre := regexp.MustCompile(relinkincludeurlre)
						if nurlre.Match([]byte(fullurl)) {
							bpass = true
						}
					}
					if relinkexcludeurlre != "" {
						nurlre := regexp.MustCompile(relinkexcludeurlre)
						if nurlre.Match([]byte(fullurl)) {
							bpass = false
						}
					}

					if bpass {
						noanchorurl := fullurl
						anchor := ""
						if strings.Index(noanchorurl, "#") != -1 {
							anchor = noanchorurl[strings.Index(noanchorurl, "#"):]
							noanchorurl = noanchorurl[0:strings.Index(noanchorurl, "#")]
						}
						urlfilename := UrlRepath(noanchorurl, relinkrootrereplacels, urlmatchurlextls)
						newmastr = bytes.Replace(mastr, urlstr, []byte(urlfilename+anchor), -1)
					}
				}
			}
		}

		if len(newmastr) > 0 {
			st := maes[i][2]
			ed := maes[i][3]
			newpagectt = toolfunc.BytesCombine(newpagectt[0:st], newmastr, newpagectt[ed:])
		}
	}
	return newpagectt
}

func PageRelinkAllToLocal(pageurl string, pagehtml []byte, relinkincludeurlre, relinkexcludeurlre string, relinkrootrereplacels, urlmatchurlextls []string) []byte {
	tagreattr := []string{"(?ism)<a[^>]*(\\shref\\s*=[^>]*)>.*?</a>", "href", "(?ism)<img[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<link[^>]*(\\shref\\s*=[^>]*)>", "href", "(?ism)<script[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<frame[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<iframe[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<source[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<embed[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<object[^>]*(\\sdata\\s*=[^>]*)>", "data"}
	tagreattri := 0
	for true {
		pagehtml = PageRelinkWithTag(pageurl, pagehtml, tagreattr[tagreattri], tagreattr[tagreattri+1], relinkincludeurlre, relinkexcludeurlre, relinkrootrereplacels, urlmatchurlextls)
		tagreattri += 2
		if tagreattri >= len(tagreattr) {
			break
		}
	}
	return pagehtml
}

func PageRelinkFullUrlWithTag(pageurl string, pagehtml []byte, tagregex, tagurlattrname, relinkincludeurlre, relinkexcludeurlre string, relinkrootrereplacels, urlmatchurlextls []string) []byte {
	newpagectt := pagehtml
	pagehtmlscriptind := HtmlAllScriptIndex(pagehtml)
	//loop find all link and replace
	atagre := regexp.MustCompile(tagregex)
	maes := atagre.FindAllSubmatchIndex(newpagectt, -1)
	for i := len(maes) - 1; i >= 0; i -= 1 {
		mastr := pagehtml[maes[i][2]:maes[i][3]]

		binrange := false
		for _, srng := range pagehtmlscriptind {
			if maes[i][2] >= srng[0] && maes[i][3] <= srng[1] {
				binrange = true
			}
		}

		if binrange {
			continue
		}

		newmastr := []byte{}
		urlstr := mastr
		urlstr = bytes.Trim(urlstr[bytes.Index(urlstr, []byte(tagurlattrname))+len(tagurlattrname):], "\r\n\t ")
		if urlstr[0] == '=' {
			urlstr = bytes.Trim(urlstr[1:], "\r\n\t ")
			if len(urlstr) > 0 {
				if urlstr[0] == '\'' {
					urlstr = urlstr[1:]
					if bytes.Index(urlstr, []byte{'\''}) != -1 {
						urlstr = urlstr[0:bytes.Index(urlstr, []byte{'\''})]
					} else if bytes.Index(urlstr, []byte{' '}) != -1 {
						urlstr = urlstr[0:bytes.Index(urlstr, []byte{' '})]
					} else if bytes.Index(urlstr, []byte{'>'}) != -1 {
						urlstr = urlstr[0:bytes.Index(urlstr, []byte{'>'})]
					}
				} else if urlstr[0] == '"' {
					urlstr = urlstr[1:]
					if bytes.Index(urlstr, []byte{'"'}) != -1 {
						urlstr = urlstr[0:bytes.Index(urlstr, []byte{'"'})]
					} else if bytes.Index(urlstr, []byte{' '}) != -1 {
						urlstr = urlstr[0:bytes.Index(urlstr, []byte{' '})]
					} else if bytes.Index(urlstr, []byte{'>'}) != -1 {
						urlstr = urlstr[0:bytes.Index(urlstr, []byte{'>'})]
					}
				} else {
					if bytes.Index(urlstr, []byte{' '}) != -1 && bytes.Index(urlstr, []byte{'>'}) != -1 {
						if bytes.Index(urlstr, []byte{' '}) < bytes.Index(urlstr, []byte{'>'}) {
							urlstr = urlstr[0:bytes.Index(urlstr, []byte{' '})]
						} else {
							urlstr = urlstr[0:bytes.Index(urlstr, []byte{'>'})]
						}
					} else if bytes.Index(urlstr, []byte{' '}) != -1 {
						urlstr = urlstr[0:bytes.Index(urlstr, []byte{' '})]
					} else if bytes.Index(urlstr, []byte{'>'}) != -1 {
						urlstr = urlstr[0:bytes.Index(urlstr, []byte{'>'})]
					}
				}
				fullurl := ToFullUrl(string(urlstr), pageurl)
				if len(fullurl) >= 7 && toolfunc.StringCompareInsensitive(fullurl[:7], "http://") || len(fullurl) >= 8 && toolfunc.StringCompareInsensitive(fullurl[:8], "https://") {
					var bpass bool
					if relinkincludeurlre == "" && relinkexcludeurlre == "" {
						bpass = true
					} else {
						bpass = false
					}
					if relinkincludeurlre != "" {
						nurlre := regexp.MustCompile(relinkincludeurlre)
						if nurlre.Match([]byte(fullurl)) {
							bpass = true
						}
					}
					if relinkexcludeurlre != "" {
						nurlre := regexp.MustCompile(relinkexcludeurlre)
						if nurlre.Match([]byte(fullurl)) {
							bpass = false
						}
					}

					if bpass {
						newmastr = bytes.Replace(mastr, urlstr, []byte(fullurl), -1)
					}
				}
			}
		}

		if len(newmastr) > 0 {
			st := maes[i][2]
			ed := maes[i][3]
			newpagectt = toolfunc.BytesCombine(newpagectt[0:st], newmastr, newpagectt[ed:])
		}
	}
	return newpagectt
}

func PageRelinkAllToFullUrl(pageurl string, pagehtml []byte, relinkincludeurlre, relinkexcludeurlre string, relinkrootrereplacels, urlmatchurlextls []string) []byte {
	tagreattr := []string{"(?ism)<a[^>]*(\\shref\\s*=[^>]*)>.*?</a>", "href", "(?ism)<img[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<link[^>]*(\\shref\\s*=[^>]*)>", "href", "(?ism)<script[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<frame[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<iframe[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<source[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<embed[^>]*(\\ssrc\\s*=[^>]*)>", "src", "(?ism)<object[^>]*(\\sdata\\s*=[^>]*)>", "data"}
	tagreattri := 0
	for true {
		pagehtml = PageRelinkFullUrlWithTag(pageurl, pagehtml, tagreattr[tagreattri], tagreattr[tagreattri+1], relinkincludeurlre, relinkexcludeurlre, relinkrootrereplacels, urlmatchurlextls)
		tagreattri += 2
		if tagreattri >= len(tagreattr) {
			break
		}
	}
	return pagehtml
}

func SplitRepWithToByteLs(ctt []byte) [][]byte {
	repwithls := make([][]byte, 0)
	preseg := []byte("")
	for i := 0; i < len(ctt); i++ {
		if ctt[i] == '$' {
			if unicode.IsDigit(rune(ctt[i+1])) {
				//fmt.Println("preseg:", string(preseg))
				if len(preseg) > 0 {
					//fmt.Println("split regex to ls0:", string(preseg))
					repwithls = append(repwithls, preseg)
					preseg = []byte{}
				}
				//fmt.Println("split regex to ls1:", string([]byte{'$', ctt[i+1]}))
				repwithls = append(repwithls, []byte{'$', ctt[i+1]})
				i++
			} else {
				preseg = append(preseg, ctt[i])
			}
		} else {
			preseg = append(preseg, ctt[i])
		}
	}
	if len(preseg) > 0 {
		//fmt.Println("split regex to ls3:", string(preseg))
		repwithls = append(repwithls, preseg)
		preseg = []byte{}
	}
	return repwithls
}

func MatchGet(ctt []byte, maind []int, cttreplacewith []byte) []byte {
	repls := SplitRepWithToByteLs(cttreplacewith)
	for i := 0; i < len(maind)/2; i += 1 {
		indbt := "$" + strconv.FormatInt(int64(i), 10)
		for j := 0; j < len(repls); j++ {
			if bytes.Compare(repls[j], []byte(indbt)) == 0 {
				repls[j] = ctt[maind[i*2]:maind[i*2+1]]
			}
		}
	}
	return bytes.Join(repls, []byte{})
}

func FindAllGet(blockctt []byte, maallind [][]int, itemreplacewith string) [][][]string {
	retval := make([][][]string, 0)
	itemrepwithls := strings.Split(itemreplacewith, ";")
	//fmt.Println("all maallind:", maallind)
	for i := 0; i < len(maallind); i += 1 {
		rowobj := make([][]string, 0)
		for g := 0; g < len(itemrepwithls); g++ {
			if strings.Index(itemrepwithls[g], "=") != -1 {
				name := itemrepwithls[g][:strings.Index(itemrepwithls[g], "=")]
				replacewith := itemrepwithls[g][strings.Index(itemrepwithls[g], "=")+1:]
				//fmt.Println("all replacewith:", replacewith, maallind[i], itemrepwithls)
				value := MatchGet(blockctt, maallind[i], []byte(replacewith))
				rowobj = append(rowobj, []string{name, string(value)})
			}
		}
		retval = append(retval, rowobj)
	}
	return retval
}

func SplitReplaceWith(ctt string) []string {
	repwithls := make([]string, 0)
	preseg := ""
	for i := 0; i < len(ctt); i++ {
		if ctt[i] == '$' {
			if unicode.IsDigit(rune(ctt[i+1])) {
				if preseg != "" {
					repwithls = append(repwithls, preseg)
					preseg = ""
				}
				repwithls = append(repwithls, "$"+ctt[i+1:i+2])
			}
		} else {
			preseg += ctt[1 : i+1]
		}
	}
	if preseg != "" {
		repwithls = append(repwithls, preseg)
		preseg = ""
	}
	return repwithls
}

//replacewithlist format:name=regex_replace_express;...
func HtmlBlockFindAttrAndValue(ctt []byte, blolckregex, blockregexrepolacewith, itemrowregex, replacewithlist string) [][][]string {
	blockctt := HtmlGetBlock(ctt, blolckregex, blockregexrepolacewith)
	//fmt.Println("blockctt:", string(blockctt), blolckregex)

	//fmt.Println("itemrowregex:", itemrowregex)
	itemregex := regexp.MustCompile(itemrowregex)
	maallind := itemregex.FindAllSubmatchIndex(blockctt, -1)
	//fmt.Println("allrow maallind:", maallind, replacewithlist)
	allitem := FindAllGet(blockctt, maallind, replacewithlist)
	return allitem
}

func HtmlGetBlock(ctt []byte, blolckregex, blockregexrepolacewith string) []byte {
	//fmt.Println("blolckregex:", blolckregex)
	if strings.HasPrefix(blolckregex, "GetFullTagRegex:") {
		return GetFullTag(ctt, blolckregex[len("GetFullTagRegex:"):])
	}
	blockregex := regexp.MustCompile(blolckregex)
	maind := blockregex.FindSubmatchIndex(ctt)
	//fmt.Println("getblock maind:", maind, string(ctt), blolckregex)
	blockctt := MatchGet(ctt, maind, []byte(blockregexrepolacewith))
	return blockctt
}

//segmentexpr=segmenttagname=regex
func SegmentBlock(segmemtls [][]byte, segmentexpr string) [][]byte {
	segmentname := ""
	marl := regexp.MustCompile("([a-zA-Z0-9]+)=(.*?);(.*)").FindSubmatch([]byte(segmentexpr))
	var segmemtregex []byte
	var segmemtregexreplacewith []byte
	var segmemtregexreplacewithls [][]byte
	if len(marl) > 0 {
		segmentname = string(marl[1])
		segmemtregex = marl[2]
		segmemtregexreplacewith = marl[3]
	} else {
		if strings.Index(segmentexpr, ";") != -1 {
			segmemtregex = []byte(segmentexpr[:strings.Index(segmentexpr, ";")])
			segmemtregexreplacewith = []byte(segmentexpr[strings.Index(segmentexpr, ";")+1:])
			if bytes.Index(segmemtregexreplacewith, []byte("@")) != -1 {
				segmemtregexreplacewithls = bytes.Split(segmemtregexreplacewith, []byte("@"))
			}
		} else {
			panic("SegmentBlock express error!")
		}
	}

	segmemtls2 := make([][]byte, 0)
	for i := 0; i < len(segmemtls); i++ {
		if len(segmemtls[i]) > 3 && string(segmemtls[i][:3]) != "ok<" {
			submaind := regexp.MustCompile(string(segmemtregex)).FindAllSubmatchIndex(segmemtls[i], -1)
			if len(submaind) > 0 {
				for j := 0; j < len(submaind); j++ {
					//fmt.Println("segmentname:", segmentname)
					//fmt.Println("submaind[j][2]:", len(submaind), submaind, j)

					//fmt.Println("segment ind:", submaind, string(segmemtls[i][submaind[j][0]:submaind[j][1]]), "segmemtregexreplacewith:", string(segmemtregex), string(segmemtregexreplacewith))
					if len(segmemtregexreplacewithls) != 3 {
						segmemtctt := MatchGet(segmemtls[i], submaind[j], []byte(segmemtregexreplacewith))
						if len(segmentname) > 0 {
							cbbt := BytesCombine([]byte("ok<"), []byte(segmentname), []byte(">"), segmemtctt, []byte("</"), []byte(segmentname), []byte(">"))
							segmemtls2 = append(segmemtls2, cbbt)
						} else {
							segmemtls2 = append(segmemtls2, segmemtctt)
						}
					} else {
						segmemtls2 = append(segmemtls2, MatchGet(segmemtls[i], submaind[j], []byte(segmemtregexreplacewithls[0])))
						if len(segmentname) > 0 {
							segmemtctt := MatchGet(segmemtls[i], submaind[j], []byte(segmemtregexreplacewithls[1]))
							cbbt := BytesCombine([]byte("ok<"), []byte(segmentname), []byte(">"), segmemtctt, []byte("</"), []byte(segmentname), []byte(">"))
							segmemtls2 = append(segmemtls2, cbbt)
						} else {
							segmemtls2 = append(segmemtls2, MatchGet(segmemtls[i], submaind[j], []byte(segmemtregexreplacewithls[1])))
						}
						segmemtls2 = append(segmemtls2, MatchGet(segmemtls[i], submaind[j], []byte(segmemtregexreplacewithls[2])))
					}
				}
			} else {
				segmemtls2 = append(segmemtls2, segmemtls[i])
			}
		} else {
			segmemtls2 = append(segmemtls2, segmemtls[i])
		}
	}
	return segmemtls2
}

func GetFullTagV2And(pagehtml []byte, fulltagstart string) []byte {
	fulltagstartls := strings.Split(fulltagstart, "&")
	fulltagstartls = toolfunc.RemoveAll(fulltagstartls, "").([]string)
	totalctt := []byte{}
	for i := 0; i < len(fulltagstartls); i++ {
		ctt1 := GetFullTagV2(pagehtml, []byte(fulltagstartls[i]))
		if len(ctt1) == 0 {
			return []byte{}
		}
		totalctt = append(totalctt, ctt1...)
	}
	return totalctt
}
func GetFullTagV2(pagehtml []byte, fulltagstart []byte) []byte {
	fulltagstartls := strings.Split(string(fulltagstart), "|")
	var ind int
	for _, fulltagst := range fulltagstartls {
		if len(fulltagst) == 0 {
			continue
		}
		if strings.Index(fulltagst, "&") == -1 {
			ind = strings.Index(string(pagehtml), fulltagst)
			if ind != -1 {
				fulltagstart = []byte(fulltagst)
				break
			}
		} else {
			andrl := GetFullTagV2And(pagehtml, fulltagst)
			if len(andrl) != 0 {
				return andrl
			}
		}
	}
	if ind == -1 {
		return []byte{}
	}
	starttagpos := 0
	tagname := []byte(toolfunc.GetRegexGroup1("<([a-zA-Z]+)\\s", string(fulltagstart)))
	stack := []string{}
	stack = append(stack, string(tagname))
	i := 0
	for i = ind + 1; i < len(pagehtml); i++ {
		if i+1+len(tagname) < len(pagehtml) && bytes.HasPrefix(pagehtml[i:], append([]byte{'<'}, tagname...)) && unicode.IsSpace(rune(pagehtml[i+1+len(tagname)])) {
			stack = append(stack, string(tagname))
		} else if bytes.HasPrefix(pagehtml[i:], append([]byte{'<', '/'}, tagname...)) {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				for i < len(pagehtml) && pagehtml[i] != '>' {
					i += 1
				}
				i += 1
				break
			}
		}
	}
	if len(stack) == 0 {
		return pagehtml[ind+starttagpos : i]
	} else {
		return []byte{}
	}
}

func GetFullTag(pagehtml []byte, tagstartregex string) []byte {
	if tagstartregex[0] != '<' {
		return []byte{}
	}
	startind := regexp.MustCompile(tagstartregex).FindSubmatchIndex(pagehtml)
	if len(startind) < 4 {
		return []byte{}
	}
	pagehtmlpos := startind[2]
	tagname := pagehtml[startind[2]:startind[3]]
	return GetFullTagWithPos(pagehtml, pagehtmlpos, tagname)
}
func GetFullTagWithPos(pagehtml []byte, pagehtmlpos int, tagname []byte) []byte {
	tagstack := []string{}
	tagstack = append(tagstack, string(tagname))
	foundtagstartpos := pagehtmlpos
	foundtagendpos := pagehtmlpos
	for true {
		pagehtmlpos2 := bytes.Index(pagehtml[pagehtmlpos:], []byte(">"))
		pagehtmlpos3 := bytes.Index(pagehtml[pagehtmlpos+1:], []byte("<"))
		if pagehtmlpos3 == -1 || pagehtmlpos3 > pagehtmlpos2 {
			if pagehtmlpos2 == -1 {
				return pagehtml[foundtagstartpos:foundtagendpos]
			}
			pagehtmlpos += pagehtmlpos2
			if !(strings.ToLower(tagstack[len(tagstack)-1]) == "pre" || strings.ToLower(tagstack[len(tagstack)-1]) == "script") {
				pagehtmlpos2 = bytes.Index(pagehtml[pagehtmlpos:], []byte("<"))
				if pagehtmlpos2 == -1 {
					return pagehtml[foundtagstartpos:foundtagendpos]
				}
				pagehtmlpos += pagehtmlpos2
			} else {
				pagehtmlpos2 = bytes.Index(pagehtml[pagehtmlpos+1:], []byte("</"+tagstack[len(tagstack)-1]))
				if pagehtmlpos2 == -1 {
					return pagehtml[foundtagstartpos:foundtagendpos]
				}
				pagehtmlpos += 1 + pagehtmlpos2
			}
		} else {
			pagehtmlpos += pagehtmlpos3 + 1
		}
		var isendtag bool
		var bcomment bool
		for pagehtmlpos < len(pagehtml) && (!(pagehtml[pagehtmlpos] >= 0x41 && pagehtml[pagehtmlpos] <= 0x5a || pagehtml[pagehtmlpos] >= 0x61 && pagehtml[pagehtmlpos] <= 0x7a || pagehtml[pagehtmlpos] >= '0' && pagehtml[pagehtmlpos] <= '9')) {
			if pagehtml[pagehtmlpos] == '/' {
				isendtag = true
			}
			if pagehtml[pagehtmlpos] == '!' {
				bcomment = true
				break
			}
			pagehtmlpos += 1
		}
		if bcomment {
			continue
		}

		endtagname := []byte{}
		if isendtag {
			pagehtmlpos3 := pagehtmlpos
			for pagehtmlpos3 < len(pagehtml) && (pagehtml[pagehtmlpos3] >= 0x41 && pagehtml[pagehtmlpos3] <= 0x5a || pagehtml[pagehtmlpos3] >= 0x61 && pagehtml[pagehtmlpos3] <= 0x7a || pagehtml[pagehtmlpos3] >= '0' && pagehtml[pagehtmlpos3] <= '9') {
				endtagname = append(endtagname, pagehtml[pagehtmlpos3])
				pagehtmlpos3 += 1
			}

			endtaglastp := bytes.Index(pagehtml[pagehtmlpos:], []byte(">"))
			if endtaglastp != -1 {
				//fmt.Println("string(endtagname)", string(endtagname))
				if !(strings.ToLower(string(endtagname)) == "img" || strings.ToLower(string(endtagname)) == "input" || strings.ToLower(string(endtagname)) == "link" || strings.ToLower(string(endtagname)) == "meta" || strings.ToLower(string(endtagname)) == "br" || strings.ToLower(string(endtagname)) == "hr") {
					//fmt.Println("tagstack = tagstack[:len(tagstack)-1]", tagstack)
					for l := len(tagstack) - 1; l >= 0; l-- {
						if toolfunc.StringCompareInsensitive(tagstack[l], string(endtagname)) {
							tagstack = tagstack[:l]
							break
						}
					}
				}
				pagehtmlpos += endtaglastp
				if len(tagstack) == 0 {
					foundtagendpos = pagehtmlpos + 1
					break
				}
			} else {
				return pagehtml[foundtagstartpos:foundtagendpos]
			}
			continue
		}
		tagname = []byte{}
		for pagehtmlpos < len(pagehtml) && (pagehtml[pagehtmlpos] >= 0x41 && pagehtml[pagehtmlpos] <= 0x5a || pagehtml[pagehtmlpos] >= 0x61 && pagehtml[pagehtmlpos] <= 0x7a || pagehtml[pagehtmlpos] >= '0' && pagehtml[pagehtmlpos] <= '9') {
			tagname = append(tagname, pagehtml[pagehtmlpos])
			pagehtmlpos += 1
		}
		if pagehtmlpos >= len(pagehtml) {
			icpind := bytes.Index(pagehtml[foundtagstartpos:], []byte("ICP备"))
			if icpind != -1 {
				for icpind > foundtagstartpos && icpind < len(pagehtml) {
					if pagehtml[icpind] == '>' {
						break
					}
					icpind--
				}
				foundtagendpos = foundtagstartpos + icpind
				break
			}
			crind := bytes.Index(pagehtml[foundtagstartpos:], []byte("Copyright &copy;"))
			if crind != -1 {
				for crind > foundtagstartpos && crind < len(pagehtml) {
					if pagehtml[crind] == '>' {
						break
					}
					crind--
				}
				foundtagendpos = foundtagstartpos + crind
				break
			}
			foundtagendpos = pagehtmlpos
			break
		}
		if !(strings.ToLower(string(tagname)) == "img" || strings.ToLower(string(tagname)) == "input" || strings.ToLower(string(tagname)) == "link" || strings.ToLower(string(tagname)) == "meta" || strings.ToLower(string(tagname)) == "br" || strings.ToLower(string(tagname)) == "hr") {
			tagstack = append(tagstack, string(tagname))
		}
	}
	if foundtagendpos >= foundtagstartpos {
		return pagehtml[foundtagstartpos:foundtagendpos]
	} else {
		return []byte{}
	}
}

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

func UrlFuncEval(name, param string) (rl string) {
	switch name {
	case "add":
		paramls := strings.Split(param, ",")
		if len(paramls) == 2 {
			v1, _ := strconv.ParseInt(paramls[0], 10, 64)
			v2, _ := strconv.ParseInt(paramls[1], 10, 64)
			rl = strconv.FormatInt(v1+v2, 10)
		} else {
			rl = param
		}
	case "minus":
		paramls := strings.Split(param, ",")
		if len(paramls) == 2 {
			v1, _ := strconv.ParseInt(paramls[0], 10, 64)
			v2, _ := strconv.ParseInt(paramls[1], 10, 64)
			rl = strconv.FormatInt(v1-v2, 10)
		} else {
			rl = param
		}
	case "isotime":
		paramls := strings.Split(param, ",")
		if len(paramls) == 1 {
			v1, _ := strconv.ParseInt(paramls[0], 10, 64)
			rl = time.Unix(v1, 0).Format("2006-01-02T15:04:05")
		} else if len(paramls) > 1 {
			var year, month, day, hour, minute, second int64
			year, _ = strconv.ParseInt(paramls[0], 10, 64)
			if regexp.MustCompile("^\\d+$").Match([]byte(paramls[1])) {
				month, _ = strconv.ParseInt(paramls[1], 10, 64)
			} else {
				month = int64(toolfunc.AbbreviateMonthToNum(paramls[1]))
			}
			if len(paramls) >= 3 {
				day, _ = strconv.ParseInt(paramls[2], 10, 64)
			}
			if len(paramls) >= 4 {
				hour, _ = strconv.ParseInt(paramls[3], 10, 64)
			}
			if len(paramls) >= 5 {
				minute, _ = strconv.ParseInt(paramls[4], 10, 64)
			}
			if len(paramls) >= 6 {
				second, _ = strconv.ParseInt(paramls[5], 10, 64)
			}
			tt := time.Date(int(year), time.Month(month), int(day), int(hour), int(minute), int(second), 0, nil)
			rl = tt.Format("2006-01-02T15:04:05")
		} else {
			rl = param
		}
	case "fixLength":
		paramls := strings.Split(param, ",")
		if len(paramls) == 2 {
			v2, _ := strconv.ParseInt(paramls[1], 10, 64)
			for i := int64(len(paramls[0])); i < v2; i++ {
				paramls[0] = "0" + paramls[0]
			}
			rl = paramls[0]
		} else {
			rl = param
		}
	case "slice":
		paramls := strings.Split(param, ",")
		if len(paramls) == 3 {
			v2, _ := strconv.ParseInt(paramls[1], 10, 64)
			v3, _ := strconv.ParseInt(paramls[2], 10, 64)
			rl = paramls[0][v2:v3]
		} else {
			rl = param
		}
	case "trim":
		paramls := strings.Split(param, ",")
		if len(paramls) == 2 {
			rl = strings.Trim(paramls[0], paramls[1])
		} else {
			rl = param
		}
	case "indexOf":
		paramls := strings.Split(param, ",")
		if len(paramls) == 2 {
			rl = strconv.FormatInt(int64(strings.Index(paramls[0], paramls[1])), 10)
		} else {
			rl = param
		}
	case "md5":
		h := md5.New()
		h.Write([]byte(param))
		cipherStr := h.Sum(nil)
		rl = hex.EncodeToString(cipherStr)
	case "sha1":
		h := sha1.New()
		h.Write([]byte(param))
		bs := h.Sum(nil)
		rl = hex.EncodeToString(bs)
	case "base64":
		rl2 := base64.URLEncoding.EncodeToString([]byte(param))
		rl = string(rl2)
	}
	return rl
}

//urlsep return from func SeperateUseList;support function:__add__,__minus__,__isotime__,__fixLength__,__mid__,__trim__,__indexOf__,__md5__,__sha1__,__base64__
func UrlCalc(urlsepstr string) string {
	urlsep := toolfunc.SeperateUseList(urlsepstr, []string{"__", "(", ")"})
	urlsepi := 0
	for true {
		if urlsepi >= len(urlsep) {
			break
		}
		stack := []string{}
		stackcur := []int{}
		stackparam := []string{}
		newstack := true
		for urlsepi+2 < len(urlsep) && urlsep[urlsepi] == "__" && urlsep[urlsepi+2] == "__" {
			if newstack {
				stack = append(stack, urlsep[urlsepi+1])
				stackcur = append(stackcur, urlsepi)
				stackparam = append(stackparam, "")
				urlsep[urlsepi] = ""
				urlsep[urlsepi+1] = ""
				urlsep[urlsepi+2] = ""
				urlsep[urlsepi+3] = ""
			}
			gonextstack := false
			for urlsep[urlsepi] != ")" {
				if urlsep[urlsepi] == "__" {
					//to next stack
					gonextstack = true
					break
				}
				stackparam[len(stackparam)-1] += urlsep[urlsepi]
				urlsep[urlsepi] = ""
				urlsepi++
			}
			if gonextstack {
				newstack = true
				continue
			} else {
				rlvalue := UrlFuncEval(stack[len(stack)-1], stackparam[len(stackparam)-1])
				urlsep[urlsepi] = rlvalue
				stack = stack[:len(stack)-1]
				stackcur = stackcur[:len(stackcur)-1]
				stackparam = stackparam[:len(stackparam)-1]
				newstack = false
				if len(stack) == 0 {
					urlsepi += 1
					break
				} else {
					urlsepi = stackcur[len(stackcur)-1]
				}
			}
		}
		if newstack {
			urlsepi++
		}
	}
	//conect as url
	newurl := ""
	for _, sep := range urlsep {
		newurl += sep
	}
	return newurl
}

func StdHtmlDoc(html []byte) []byte {
	html2 := regexp.MustCompile(">\\s+<").ReplaceAll([]byte(html), []byte("><"))
	html2 = regexp.MustCompile(">\\s+").ReplaceAll([]byte(html2), []byte(">"))
	html2 = regexp.MustCompile("\\s+<").ReplaceAll([]byte(html2), []byte("<"))
	return html2
}

func RealFillRuningTimeCheck(objsetname string, ch chan int, ctt []byte) {
	//fmt.Println("RealFillRuningTimeCheck start:", runstr)
	for true {
		select {
		case <-time.After(time.Second * 120):
			toolfunc.WriteFile("longtime-"+toolfunc.TimeForFileName(), ctt)
			return
			//panic("RealFillRuningTimeCheck running time too long at position " + strconv.FormatInt(*pos, 10) + " postime:" + strconv.FormatInt(*postime, 10) + " nowtime:" + strconv.FormatInt(time.Now().Local().Unix(), 10) + " str:" + *runstr)
		case <-ch:
			return
		}
	}
}

func ExtractPageBody(in_text_u8 []byte) (allfulltag []byte) {
	//remove all float div
	//check have , . etc area

	//runch := make(chan int, 0)
	//runpos := int64(0)
	//runpostime := time.Now().Local().Unix()
	//var runstr string
	//go RealFillRuningTimeCheck("", runch, in_text_u8)

	out_text_u8 := make([]byte, 0)
	preend := 0
	bnewtagdonesearch := false
	i := 0
	laststart := -1
	for ; i < len(in_text_u8); i++ {
		switch in_text_u8[i] {
		case '<':
			i++
			tagname := ""
			var bendtag bool
			if bytes.HasPrefix(in_text_u8[i:], []byte("!--")) {
				i = i + bytes.Index(in_text_u8[i:], []byte("-->")) + 2
			} else {
				if bytes.HasPrefix(in_text_u8[i:], []byte("/")) {
					bendtag = true
					i += 1
				}
				for i < len(in_text_u8) && !unicode.IsLetter(rune(in_text_u8[i])) {
					i++
				}
				for i < len(in_text_u8) && unicode.IsLetter(rune(in_text_u8[i])) {
					tagname += string(unicode.ToLower(rune(in_text_u8[i])))
					i++
				}
				bequalpass := false
				for i < len(in_text_u8) {
					if in_text_u8[i] == '=' {
						i++
						bequalpass = true
					} else if bequalpass && in_text_u8[i] == '\'' {
						i++
						for i < len(in_text_u8) && in_text_u8[i] != '\'' {
							i++
						}
						if i < len(in_text_u8) && in_text_u8[i] == '\'' {
							i++
						}
						bequalpass = false
					} else if bequalpass && in_text_u8[i] == '"' {
						i++
						for i < len(in_text_u8) && in_text_u8[i] != '"' {
							i++
						}
						if i < len(in_text_u8) && in_text_u8[i] == '"' {
							i++
						}
						bequalpass = false
					} else if in_text_u8[i] != '>' {
						i++
					} else {
						break
					}
				}
			}
			if bendtag == false && (tagname == "style" || tagname == "script") {
				for i < len(in_text_u8) {
					if i+8 < len(in_text_u8) && toolfunc.Byte1DCompareInsensitive(in_text_u8[i:i+8], []byte("</style>")) || i+9 < len(in_text_u8) && toolfunc.Byte1DCompareInsensitive(in_text_u8[i:i+9], []byte("</script>")) { // || i+11 < len(in_text_u8) && toolfunc.Byte1DCompareInsensitive(in_text_u8[i:i+11], []byte("</textarea>")) {
						i = i + bytes.Index(in_text_u8[i:], []byte{'>'})
						break
					} else {
						i++
					}
				}
			}
			if bendtag == true {
				if tagname == "td" {
					out_text_u8 = append(out_text_u8, '\t')
				}
				if tagname == "tr" || tagname == "p" || tagname == "div" {
					out_text_u8 = append(out_text_u8, '\n')
				}
				if tagname == "span" {
					out_text_u8 = append(out_text_u8, ' ')
				}
			}
			if !(string(tagname) == "br" || string(tagname) == "hr" || string(tagname) == "img" || string(tagname) == "input") {
				bnewtagdonesearch = false
			}
			continue
		case '>':
			continue
		case '&':
			htmlsymval := make([]byte, 0)
			if i+1 < len(in_text_u8) && in_text_u8[i+1] == '#' {
				if i+2 < len(in_text_u8) && unicode.ToLower(rune(in_text_u8[i+2])) == 'x' {
					htmlsymval = make([]byte, 0)
					htmlsymval = append(htmlsymval, []byte("&#x")...)
					i += 3
					for i < len(in_text_u8) {
						if unicode.ToLower(rune(in_text_u8[i])) >= 'a' && unicode.ToLower(rune(in_text_u8[i])) <= 'f' || unicode.ToLower(rune(in_text_u8[i])) >= '0' && unicode.ToLower(rune(in_text_u8[i])) <= '9' {
							htmlsymval = append(htmlsymval, byte(unicode.ToLower(rune(in_text_u8[i]))))
							i++
						} else {
							if in_text_u8[i] == ';' {
								htmlsymval = append(htmlsymval, byte(unicode.ToLower(rune(in_text_u8[i]))))
								i++
							}
							break
						}
						if len(htmlsymval) > MAXHTMLNAMETAG {
							break
						}
					}

					spechval, bspechval := specialcharmap.Load(string(htmlsymval))
					if bspechval {
						out_text_u8 = append(out_text_u8, []byte(spechval.(string))...)
					} else {
						indnumsym := bytes.Index(htmlsymval, []byte{byte('#')})
						indsemicolon := bytes.Index(htmlsymval, []byte{byte(';')})
						if indnumsym != -1 && indsemicolon != -1 && indnumsym < indsemicolon {
							hexnum := htmlsymval[indnumsym+2 : indsemicolon]
							ucode, _ := strconv.ParseInt(string(hexnum), 16, 32)
							ucode2 := rune(ucode)
							out_text_u8 = append(out_text_u8, []byte(string(ucode2))...)
						} else {
							out_text_u8 = append(out_text_u8, []byte(htmlsymval)...)
						}
					}
				} else {
					htmlsymval = make([]byte, 0)
					htmlsymval = append(htmlsymval, []byte("&#")...)
					i += 2
					for i < len(in_text_u8) {
						if unicode.ToLower(rune(in_text_u8[i])) >= '0' && unicode.ToLower(rune(in_text_u8[i])) <= '9' {
							htmlsymval = append(htmlsymval, byte(unicode.ToLower(rune(in_text_u8[i]))))
							i++
						} else {
							if in_text_u8[i] == ';' {
								htmlsymval = append(htmlsymval, byte(unicode.ToLower(rune(in_text_u8[i]))))
								i++
							}
							break
						}
						if len(htmlsymval) > MAXHTMLNAMETAG {
							break
						}
					}

					spechrval, bspechrval := specialcharmap.Load(string(htmlsymval))
					if bspechrval {
						out_text_u8 = append(out_text_u8, spechrval.(string)...)
					} else {
						if bytes.Index(htmlsymval, []byte{byte('#')}) != -1 && bytes.Index(htmlsymval, []byte{byte(';')}) != -1 {
							codenum := htmlsymval[bytes.Index(htmlsymval, []byte{byte('#')})+1 : bytes.Index(htmlsymval, []byte{byte(';')})]
							ucode, _ := strconv.ParseInt(string(codenum), 10, 32)
							ucode2 := rune(ucode)
							out_text_u8 = append(out_text_u8, []byte(string(ucode2))...)
						}

					}
				}
			} else {
				htmlsymval = make([]byte, 0)
				htmlsymval = append(htmlsymval, '&')
				i++
				for i < len(in_text_u8) {
					if in_text_u8[i] >= 0x61 && in_text_u8[i] <= 0x7a || in_text_u8[i] >= 0x41 && in_text_u8[i] <= 0x5a || in_text_u8[i] >= 0x30 && in_text_u8[i] <= 0x39 {
						htmlsymval = append(htmlsymval, byte(unicode.ToLower(rune(in_text_u8[i]))))
						i++
						if len(htmlsymval) > MAXHTMLNAMETAG {
							break
						}
					} else {
						if in_text_u8[i] == ';' {
							htmlsymval = append(htmlsymval, byte(unicode.ToLower(rune(in_text_u8[i]))))
							i++
						}
						break
					}
				}
				spechrval, bspechrval := specialcharmap.Load(string(htmlsymval))
				if bspechrval {
					out_text_u8 = append(out_text_u8, []byte(spechrval.(string))...)
				} else {
					out_text_u8 = append(out_text_u8, htmlsymval...)
				}
			}
			i-- //避免for ++ 越过
			break
		default:
			var bok bool
			if bnewtagdonesearch == false && in_text_u8[i] == ',' || in_text_u8[i] == '.' {
				bok = true
			}
			out_text_u8 = append(out_text_u8, in_text_u8[i])
			if bnewtagdonesearch == false && !bok {
				if bytes.HasSuffix(out_text_u8, []byte("，")) || bytes.HasSuffix(out_text_u8, []byte("。")) {
					bok = true
				}
			}
			if bnewtagdonesearch == false && bok {
				//find up div tag
				recursivedivcnt := 0
				bfindfirsttag := true
				if preend == laststart {
					break
				}
				//fmt.Println("preend", preend, i)
				for t := i; t >= preend; t-- {
					if bfindfirsttag == true {
						if in_text_u8[t] == '<' {
							ts2 := t
							t++
							tagname := []byte{}
							for t < len(in_text_u8) {
								if unicode.IsSpace(rune(in_text_u8[t])) {
									t++
								} else {
									break
								}
							}
							bfendtag := false
							if in_text_u8[t] == '/' {
								bfendtag = true
								t++
							}
							//fmt.Println(string(in_text_u8[t:t+10]), unicode.IsLetter(rune(in_text_u8[t])))
							for s := t; s < len(in_text_u8); s++ {
								if unicode.IsLetter(rune(in_text_u8[s])) || in_text_u8[s] >= '0' && in_text_u8[s] <= '9' {
									tagname = append(tagname, byte(unicode.ToLower(rune(in_text_u8[s]))))
								} else {
									break
								}
							}
							//fmt.Println("string(tagname)", string(tagname))
							if !(bfendtag == false && (string(tagname) == "p" || string(tagname) == "font" || string(tagname) == "span" || string(tagname) == "div" || string(tagname) == "textarea" || string(tagname) == "pre")) { // || string(tagname) == "h1" || string(tagname) == "h2" || string(tagname) == "h3" || string(tagname) == "h4" || string(tagname) == "h5" || string(tagname) == "h6") {
								if !bfendtag { //
									if !(string(tagname) == "br" || string(tagname) == "hr" || string(tagname) == "img" || string(tagname) == "input") {
										break
									} else {
										t = ts2
									}
								} else {
									t -= len(tagname) - 1
									tagstack := []string{string(tagname)}
									for t >= 0 {
										if in_text_u8[t] != '<' {
											t--
										} else {
											ts := t
											t += 1
											if in_text_u8[t] == '/' {
												t++
												tagname = []byte{}
												for t < len(in_text_u8) {
													if unicode.IsLetter(rune(in_text_u8[t])) || in_text_u8[t] >= '0' && in_text_u8[t] <= '9' {
														tagname = append(tagname, byte(unicode.ToLower(rune(in_text_u8[t]))))
														t++
													} else {
														break
													}
												}
												if string(tagname) == "script" {
													for ts >= 0 {
														if ts+7 < len(in_text_u8) && !toolfunc.Byte1DCompareInsensitive(in_text_u8[ts:ts+7], []byte("<script")) {
															ts--
														} else {
															ts += 1
															break
														}
													}
												}
												//fmt.Println("tagname1", t, string(tagname))
												if !(string(tagname) == "br" || string(tagname) == "hr" || string(tagname) == "img" || string(tagname) == "input" || string(tagname) == "script" || string(tagname) == "textarea" || string(tagname) == "style") {
													tagstack = append(tagstack, string(tagname))
												}
												t = ts - 1
											} else {
												tagname = []byte{}
												for t < len(in_text_u8) {
													if unicode.IsLetter(rune(in_text_u8[t])) || in_text_u8[t] >= '0' && in_text_u8[t] <= '9' {
														tagname = append(tagname, byte(unicode.ToLower(rune(in_text_u8[t]))))
														t++
													} else {
														break
													}
												}
												//fmt.Println("tagname2", t, string(tagname))
												if !(string(tagname) == "br" || string(tagname) == "hr" || string(tagname) == "img" || string(tagname) == "input" || string(tagname) == "textarea" || string(tagname) == "script" || string(tagname) == "style") {
													if toolfunc.SliceLastIndex(tagstack, string(tagname)) >= 0 {
														tagstack = tagstack[:toolfunc.SliceLastIndex(tagstack, string(tagname))]
													}
													if len(tagstack) == 0 {
														bfindfirsttag = false
														break
													}
												}
												t = ts - 1
											}
										}
									}
								}
							} else {
								bfindfirsttag = false
							}
						}
					}
					b1ma := false
					b2ma := false
					if t-4 > 0 && toolfunc.Byte1DCompareInsensitive(in_text_u8[t-4:t], []byte("<div")) {
						b1ma = true
					}
					if t-3 > 0 && toolfunc.Byte1DCompareInsensitive(in_text_u8[t-3:t], []byte("<td")) {
						b2ma = true
					}
					if b1ma || b2ma {
						if recursivedivcnt == 0 {
							delta := 0
							if b2ma {
								delta = 1
							}
							//fmt.Println("to GetFullTagWithPos", string(in_text_u8[t-4+delta:t-4+delta+100]))
							fultag := GetFullTagWithPos([]byte(in_text_u8), t-4+delta, in_text_u8[t-3+delta:t])
							//fmt.Println("end GetFullTagWithPos", len(fultag))
							if len(fultag) != 0 {
								i = t - 4 + len(fultag) - 1
								if bytes.Index(fultag, []byte("ICP备")) == -1 && bytes.Index(fultag, []byte("Copyright &copy;")) == -1 {
									allfulltag = append(allfulltag, fultag...)
								}
							}
							break
						} else {
							recursivedivcnt--
						}
					} else if t-6 > 0 && toolfunc.Byte1DCompareInsensitive(in_text_u8[t-6:t], []byte("</div>")) {
						recursivedivcnt++
					}
				}
				if preend == i+1 {
					break
				}
				if bfindfirsttag == false {
					laststart = preend
					preend = i + 1
				}
				bnewtagdonesearch = true
			}
		}
	}
	//runch <- 1
	return allfulltag
}
