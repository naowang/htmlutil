package htmlutil

import (
	//"bytes"
	"fmt"
	"io/ioutil"

	//"os"
	"testing"
)

func TestInit(t *testing.T) {
	//Init()
}

func TestUrlCalc(t *testing.T) {
	fmt.Println(UrlCalc("eeee__trim__(3324755,5)"))
}

func TestObjDbKv(t *testing.T) {
	/*
								htmldata, _ := ioutil.ReadFile("test.html")
								newhtmldata := HtmlToText(htmldata)
								ioutil.WriteFile("test.html.txt", newhtmldata, 0666)

								htmldata, _ = ioutil.ReadFile("test2.html")
								newhtmldata = HtmlToText(htmldata)
								ioutil.WriteFile("test2.html.txt", newhtmldata, 0666)

								noscripthtml := HtmlRemoveAllScript(string(htmldata))
								ioutil.WriteFile("noscripthtml.html", []byte(noscripthtml), 0666)

								allurl := HtmlFindAllUrl("http://www.sina.com.cn/", string(htmldata))
								fi, _ := os.Create("allurl.txt")
								for _, urlinfo := range allurl {
									fi.Write([]byte(urlinfo[0]))
									fi.Write([]byte("\t"))
									fi.Write([]byte(urlinfo[1]))
									fi.Write([]byte("\t"))
									fi.Write([]byte(urlinfo[2]))
									fi.Write([]byte("\n"))
								}
								fi.Close()

								relinkedhtml := PageRelinkAllToLocal("http://www.sina.com.cn/", string(htmldata), "", "", []string{}, []string{})
								ioutil.WriteFile("relinkedhtml.html", []byte(relinkedhtml), 0666)

								segmemtls := make([][]byte, 3)
								segmemtls[0] = []byte("4ioiojdaslkfjdsajflkajgldfklgjdfklg")
								segmemtls[1] = []byte("890sdalkfklsdajlfjklkjctfopg7uioph89awenlsd")
								segmemtls[2] = []byte("65df65asd6565awe4465rt465rugr4545sd")
								segmemtls2 := SegmentBlock(segmemtls, "(kls.*?)pg7;$1")
								fmt.Println("segmemtls2:", segmemtls2)
								rightbyteresult := [][]byte{[]byte{52, 105, 111, 105, 111, 106, 100, 97, 115, 108, 107, 102, 106, 100, 115, 97, 106, 102, 108, 107, 97, 106, 103, 108, 100, 102, 107, 108, 103, 106, 100, 102, 107, 108, 103}, []byte{107, 108, 115, 100, 97, 106, 108, 102, 106, 107, 108, 107, 106, 99, 116, 102, 111}, []byte{54, 53, 100, 102, 54, 53, 97, 115, 100, 54, 53, 54, 53, 97, 119, 101, 52, 52, 54, 53, 114, 116, 52, 54, 53, 114, 117, 103, 114, 52, 53, 52, 53, 115, 100}}
								for i := 0; i < len(rightbyteresult); i++ {
									if bytes.Compare(rightbyteresult[i], segmemtls2[i]) != 0 {
										panic("result error!")
									} else {
										fmt.Println(string(rightbyteresult[i]))
									}
								}

								segmemtls[0] = []byte("4ioiojdaslkfjdsajflkajgldfklgjdfklg")
								segmemtls[1] = []byte("890sdalkfklsdajlfjklkjctfopg7uioph89awenlsd")
								segmemtls[2] = []byte("65df65asd6565awe4465rt465rugr4545sd")
								segmemtls2 = SegmentBlock(segmemtls, "TTT=kls(.*?pg7);$1")
								fmt.Println("segmemtls2", segmemtls2)
								rightbyteresult = [][]byte{[]byte{52, 105, 111, 105, 111, 106, 100, 97, 115, 108, 107, 102, 106, 100, 115, 97, 106, 102, 108, 107, 97, 106, 103, 108, 100, 102, 107, 108, 103, 106, 100, 102, 107, 108, 103}, []byte{111, 107, 60, 84, 84, 84, 62, 100, 97, 106, 108, 102, 106, 107, 108, 107, 106, 99, 116, 102, 111, 112, 103, 55, 60, 47, 84, 84, 84, 62}, []byte{54, 53, 100, 102, 54, 53, 97, 115, 100, 54, 53, 54, 53, 97, 119, 101, 52, 52, 54, 53, 114, 116, 52, 54, 53, 114, 117, 103, 114, 52, 53, 52, 53, 115, 100}}
								for i := 0; i < len(rightbyteresult); i++ {
									if bytes.Compare(rightbyteresult[i], segmemtls2[i]) != 0 {
										panic("result error!")
									} else {
										fmt.Println(string(rightbyteresult[i]))
									}
								}

								blockdata := HtmlGetBlock([]byte("890243\niosdklfajkdsal;awl;sadl;fkl;dskl;fdsf"), ".*?(l;).*", "bff$1err")
								fmt.Println(string(blockdata))
								if string(blockdata) != "bffl;err" {
									panic("find block error!")
								}

								attrvallist := HtmlBlockFindAttrAndValue([]byte("io3489uidsahhjkdsa,hhh=7,bbb=8\n,hhh=466,bbb=8654\nhfhklq4rfg445uy45e455432df"), "(?ism)uidsahhj(.*?)32d", "$1", "hhh=(\\d+)[^\\n]*bbb=(\\d+)", "hhh=$1;bbbget=$2")
								for i := 0; i < len(attrvallist); i++ {
									fmt.Println("group row:")
									for j := 0; j < len(attrvallist[i]); j++ {
										str := ""
										for k := 0; k < len(attrvallist[i][j]); k++ {
											str += string(attrvallist[i][j][k]) + " "
										}
										fmt.Println("str:", str)
									}

								}

								fulltag := GetFullTag([]byte("dklfjsdalfjdsk<aa>sdk<img src=dsfaldf></img>fj<bb>sdfjsd<cc>sd<aa>fd</cc>sdf<img sdfds />kld<input sdklfj dsfd></input>sf</bb>sd<aa>sd<input lksdf />fdsf</aa>af<script> *&<><></script>df</aa>dfgfdsdaf"), "<(aa)>sdk")
								if string(fulltag) != "<aa>sdk<img src=dsfaldf></img>fj<bb>sdfjsd<cc>sd<aa>fd</cc>sdf<img sdfds />kld<input sdklfj dsfd></input>sf</bb>sd<aa>sd<input lksdf />fdsf</aa>" {
									fmt.Println(string(fulltag))
									panic("error!")
								} else {
									fmt.Println(string(fulltag))
								}

								testtxt3 := `<p class="paragraph">????????????????????????<img src="/uploadfiles/17.jpg" title="" width="100%" onload="if(this.parentNode.parentNode.offsetWidth-15>=this.naturalWidth){this.style.width='auto';}else{this.style.width='100%';}" style="width: 100%;"><br></p>`
								txt3 := HtmlToText([]byte(testtxt3))
								fmt.Println("txt3:", string(txt3))

								testtxt3 = `<p class="paragraph"><span class="highlight" style="background-color:rgb(255, 255, 255)"><span class="colour" style="color:rgb(0, 0, 0)">?????????????????????????????????</span></span></p>`
								txt3 = HtmlToText([]byte(testtxt3))
								fmt.Println("txt3:", string(txt3))

								fmt.Println(string(HtmlToText([]byte("<p>???&hellip;&hellip;???????????????????????????????????????</p>"))))

								pagectt, _ := ioutil.ReadFile("testextractbody1.htm")
								fmt.Println(string(HtmlToText(pagectt)))
								body := ExtractPageBody(pagectt)
								ioutil.WriteFile("testextractboyd1_out.htm", []byte(body), 0666)

								{
									pagectt, _ := ioutil.ReadFile("testextractbody2.htm")
									fmt.Println(string(HtmlToText(pagectt)))
									body := ExtractPageBody(pagectt)
									ioutil.WriteFile("testextractboyd2_out.htm", []byte(body), 0666)
								}

								{
									pagectt, _ := ioutil.ReadFile("testextractbody3.htm")
									fmt.Println(string(HtmlToText(pagectt)))
									body := ExtractPageBody(pagectt)
									ioutil.WriteFile("testextractboyd3_out.htm", []byte(body), 0666)
								}

								pagectt2, _ := ioutil.ReadFile("testextractbody4.htm")
								allurl = HtmlFindAllUrl("http://zgsxd.k618.cn/", string(pagectt2))
								fi, _ = os.Create("testextractbody4_out.txt")
								for _, urlinfo := range allurl {
									fi.Write([]byte(urlinfo[0]))
									fi.Write([]byte("\t"))
									fi.Write([]byte(urlinfo[1]))
									fi.Write([]byte("\t"))
									fi.Write([]byte(urlinfo[2]))
									fi.Write([]byte("\n"))
								}
								fi.Close()

								fmt.Println(ToFullUrl(".././../../aa/html", "http://www.cc.com/aa/bb/cc/dd/ee/df.htm"))
								fmt.Println(ToFullUrl(".///.././/../../aa/html", "http://www.cc.com/aa/bb/cc/dd/ee/df.htm"))
								fmt.Println(ToFullUrl("../../../../sdkfjdslafafsaf.jpg", "http://www.baidu.com"))

								{
									pagectt, _ := ioutil.ReadFile("testextractbody5.htm")
									allurl := HtmlFindAllUrl("http://www.cdpf.org.cn/", string(pagectt))
									fmt.Println(allurl)
								}
								{
									pagectt, _ := ioutil.ReadFile("testextractboyd6.htm")
									bodytxt := ExtractPageBody(pagectt)
									fmt.Println(len(bodytxt))
								}
								{
									pagectt, _ := ioutil.ReadFile("testextractboyd7.htm")
									bodytxt := ExtractPageBody(pagectt)
									fmt.Println(len(bodytxt), string(bodytxt))
								}
								{
									pagectt, _ := ioutil.ReadFile("testextractboyd8.htm")
									bodytxt := ExtractPageBody(pagectt)
									fmt.Println(len(bodytxt), string(bodytxt))
								}
								{
									pagectt, _ := ioutil.ReadFile("testextractboyd9.htm")
									bodytxt := ExtractPageBody(pagectt)
									fmt.Println(len(bodytxt), string(bodytxt), string(HtmlToText(bodytxt)))
								}
								{
									pagectt, _ := ioutil.ReadFile("testextractboyd10.htm")
									bodytxt := ExtractPageBody(pagectt)
									fmt.Println(len(bodytxt), string(bodytxt), string(HtmlToText(bodytxt)))
								}
								{
								pagectt, _ := ioutil.ReadFile("testextractboyd11.htm")
								bodytxt := ExtractPageBody(pagectt)
								fmt.Println(len(bodytxt), string(bodytxt), string(HtmlToText(bodytxt)))
							}
							{
								pagectt, _ := ioutil.ReadFile("testextractboyd12.htm")
								bodytxt := ExtractPageBody(pagectt)
								fmt.Println(len(bodytxt), string(bodytxt), string(HtmlToText(bodytxt)))
							}
							{
							pagectt, _ := ioutil.ReadFile("testextractboyd14.htm")
							bodytxt := ExtractPageBody(pagectt)
							fmt.Println(len(bodytxt), string(bodytxt), string(HtmlToText(bodytxt)))
						}
						{
						pagectt, _ := ioutil.ReadFile("testextractboyd15.htm")
						bodytxt := ExtractPageBody(pagectt)
						fmt.Println(len(bodytxt), string(bodytxt), string(HtmlToText(bodytxt)))
					}
					{
						pagectt, _ := ioutil.ReadFile("testextractboyd16.htm")
						bodytxt := ExtractPageBody(pagectt)
						fmt.Println(len(bodytxt), string(bodytxt), string(HtmlToText(bodytxt)))
					}
					{
					pagectt, _ := ioutil.ReadFile("testextractboyd17.htm")
					bodytxt := ExtractPageBody(pagectt)
					fmt.Println(len(bodytxt), string(bodytxt), string(HtmlToText(bodytxt)))
				}
				{
					pagectt, _ := ioutil.ReadFile("testextractboyd18.htm")
					bodytxt := ExtractPageBody(pagectt)
					fmt.Println(len(bodytxt), string(bodytxt), string(HtmlToText(bodytxt)))
				}
				{
				pagectt, _ := ioutil.ReadFile("longtime-20190505143439")
				bodytxt := ExtractPageBody(pagectt)
				fmt.Println(len(bodytxt), string(bodytxt), string(HtmlToText(bodytxt)))
			}
			{
			pagectt, _ := ioutil.ReadFile("longtime-20190505184204")
			bodytxt := ExtractPageBody(pagectt)
			fmt.Println(len(bodytxt), string(bodytxt), string(HtmlToText(bodytxt)))
		}
	*/

	{
		pagectt, _ := ioutil.ReadFile("longtime-20190505235002")
		bodytxt := ExtractPageBody(pagectt)
		fmt.Println(len(bodytxt), string(bodytxt), string(HtmlToText(bodytxt)))
	}
	//return

	fmt.Println(string(HtmlToText([]byte("&#22823;&#21457;&#24425;&#31080;"))))

	fmt.Println(string(GetTitle([]byte("jldsjfkld<title>dskfjsdklafj</title>sfjskafd"))))

	fmt.Println(ToFullUrl("/?dfs=23432", "http://www.baidu.com/?jsdlkfjsd"))
	fmt.Println(ToFullUrl("product_bangongzhuopingfengkazuo.html", "http://huamingds.com/"))
	fmt.Println(ToFullUrl("product_bangongzhuopingfengkazuo.html", "http://huamingds.com"))
	fmt.Println(ToFullUrl("/kdxyu/17466225.html", "http://www.mgrgmpfu.tw/zzuoia/20190402/"))
	fmt.Println(ToFullUrl("index.php?action=page&pid=usingLibrary/retrieval", "http://www.eq-tsg.cn/index.php?action=page&pid=usingLibrary/retrieval"))

	fmt.Println(GetAllTagALink("http://www.baidu.com/aa/bb/cc/dd.php", []byte(`assadfasf<a dsfsaf href='../zz/dd.php' dsfdsdsfsdf>ds\r\nfdsf</a>dsfdsfadaaa<a href="asdfsaf.php" class="zljdPMAD" title="???????????????.-.[????????????2].??????.(APE)(ED2000.COM).cue">???????????????.-.[????????????2].??????.(APE)(ED2000.COM)\r\n.cue</a>dddds`)))
	html := []byte(`

<!doctype html>
<!--[if lt IE 7]>      <html class="no-js lt-ie9 lt-ie8 lt-ie7" lang=""> <![endif]-->
<!--[if IE 7]>         <html class="no-js lt-ie9 lt-ie8" lang=""> <![endif]-->
<!--[if IE 8]>         <html class="no-js lt-ie9" lang=""> <![endif]-->
<!--[if gt IE 8]><!--> <html class="no-js" lang="zh-hant" prefix="og: http://ogp.me/ns#"> <!--<![endif]-->
<head>
	<script>(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
})(window,document,'script','dataLayer','GTM-PTG5DG');</script>
	<meta charset="UTF-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
	<meta name="viewport" content="width=device-width,minimum-scale=1,initial-scale=1">
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
	<meta name="google" content="notranslate">
	<link rel="apple-touch-icon" sizes="57x57" href="https://www.gov.mo/zh-hant/wp-content/static/favicon/apple-icon-57x57.png">
	<link rel="apple-touch-icon" sizes="60x60" href="https://www.gov.mo/zh-hant/wp-content/static/favicon/apple-icon-60x60.png">
	<link rel="apple-touch-icon" sizes="72x72" href="https://www.gov.mo/zh-hant/wp-content/static/favicon/apple-icon-72x72.png">
	<link rel="apple-touch-icon" sizes="76x76" href="https://www.gov.mo/zh-hant/wp-content/static/favicon/apple-icon-76x76.png">
	<link rel="apple-touch-icon" sizes="114x114" href="https://www.gov.mo/zh-hant/wp-content/static/favicon/apple-icon-114x114.png">
	<link rel="apple-touch-icon" sizes="120x120" href="https://www.gov.mo/zh-hant/wp-content/static/favicon/apple-icon-120x120.png">
	<link rel="apple-touch-icon" sizes="144x144" href="https://www.gov.mo/zh-hant/wp-content/static/favicon/apple-icon-144x144.png">
	<link rel="apple-touch-icon" sizes="152x152" href="https://www.gov.mo/zh-hant/wp-content/static/favicon/apple-icon-152x152.png">
	<link rel="apple-touch-icon" sizes="180x180" href="https://www.gov.mo/zh-hant/wp-content/static/favicon/apple-icon-180x180.png">
	<link rel="icon" type="image/png" sizes="192x192"  href="https://www.gov.mo/zh-hant/wp-content/static/favicon/android-icon-192x192.png">
	<link rel="icon" type="image/png" sizes="32x32" href="https://www.gov.mo/zh-hant/wp-content/static/favicon/favicon-32x32.png">
	<link rel="icon" type="image/png" sizes="96x96" href="https://www.gov.mo/zh-hant/wp-content/static/favicon/favicon-96x96.png">
	<link rel="icon" type="image/png" sizes="16x16" href="https://www.gov.mo/zh-hant/wp-content/static/favicon/favicon-16x16.png">
	<link rel="manifest" href="https://www.gov.mo/zh-hant/wp-content/static/favicon/manifest.json">
	<meta name="msapplication-TileColor" content="#ffffff">
	<meta name="msapplication-TileImage" content="https://www.gov.mo/zh-hant/wp-content/static/favicon/ms-icon-144x144.png">
	<meta name="theme-color" content="#ffffff">
<script>
	var locale = 'zh-hant';
	var timestamp = Date.now();
	var site_url = 'https://www.gov.mo/zh-hant';
	var sign_in_url = 'https://www.gov.mo/zh-hant/sign-in/' + '?' + timestamp;
	var favorite_services_home_part_url = 'https://www.gov.mo/zh-hant/home/my-favorite-services/';
</script>
		<script>
	var is_admin_bar_showing = false;
	</script>
		<title>??????????????????????????????????????? &#8211; ???????????????????????????????????????</title>
<link rel='dns-prefetch' href='//ajax.googleapis.com' />
<link rel='stylesheet' id='owl-carousel-style-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/owl.carousel/assets/owl.carousel.min.css' type='text/css' media='all' />
<link rel='stylesheet' id='owl-carousel-theme-style-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/owl.carousel/assets/owl.theme.min.css' type='text/css' media='all' />
<link rel='stylesheet' id='bootstrap-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/bootstrap.min.css' type='text/css' media='all' />
<link rel='stylesheet' id='materialadmin-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/materialadmin.min.css' type='text/css' media='all' />
<link rel='stylesheet' id='govmo-main-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/govmo.min.css?20200124' type='text/css' media='all' />
<link rel='stylesheet' id='template-home-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/template-home.min.css?20190729' type='text/css' media='all' />
<link rel='stylesheet' id='template-home-20191001-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/template-home-20191001.min.css' type='text/css' media='all' />
<link rel='stylesheet' id='template-home-20200202-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/template-home-20200202.min.css' type='text/css' media='all' />
<link rel='stylesheet' id='template-home-20200214-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/template-home-20200214.min.css' type='text/css' media='all' />
<link rel='stylesheet' id='wp-block-library-css'  href='https://www.gov.mo/zh-hant/wp-includes/css/dist/block-library/style.min.css' type='text/css' media='all' />
<link rel='stylesheet' id='font-awesome-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/font-awesome.min.css' type='text/css' media='all' />
<link rel='stylesheet' id='typhoon-signal-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/typhoon-signal.min.css' type='text/css' media='all' />
<link rel='stylesheet' id='stormsurge-signal-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/stormsurge-signal.min.css' type='text/css' media='all' />
<link rel='stylesheet' id='material-design-iconic-font-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/material-design-iconic-font.min.css' type='text/css' media='all' />
<link rel='stylesheet' id='social-share-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/libs/social-share/social-share.min.css' type='text/css' media='all' />
<link rel='stylesheet' id='govmo-yamm-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/yamm.min.css' type='text/css' media='all' />
<link rel='stylesheet' id='govmo-cookie-consent-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/govmo-cookie-consent.min.css' type='text/css' media='all' />
<link rel='stylesheet' id='govmo-print-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/govmo-print.min.css' type='text/css' media='print' />
<!--[if lt IE 9]>
<link rel='stylesheet' id='govmo-ie8-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/govmo-ie8.min.css' type='text/css' media='1' />
<![endif]-->
<script type='text/javascript' id='wp-sentry-browser-js-extra'>
/* <![CDATA[ */
var wp_sentry = {"dsn":"https:\/\/89c923739b1a4cc2801f992b2ee8fa06@sentry.safp.informac.gov.mo\/10","release":"1.0.0","environment":"production","content":{"tags":{"wordpress":"5.5.3","language":"zh-TW"}}};
/* ]]> */
</script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/plugins/wp-sentry-integration/public/wp-sentry-browser.min.js' id='wp-sentry-browser-js'></script>
<!--[if lt IE 9]>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/utils/html5shiv.min.js' id='html5shiv-js'></script>
<![endif]-->
<link rel="canonical" href="https://www.gov.mo/zh-hant/" />
<link rel='shortlink' href='https://www.gov.mo/zh-hant/' />
<link rel="alternate" hreflang="en-US" href="https://www.gov.mo/en/"><link rel="alternate" hreflang="pt-PT" href="https://www.gov.mo/pt/"><link rel="alternate" hreflang="zh-hans" href="https://www.gov.mo/zh-hans/"><link rel="alternate" hreflang="zh-hant" href="https://www.gov.mo/zh-hant/"><meta property="fb:app_id" content="348547258919852"/>
<meta property="og:url" content="https://www.gov.mo/zh-hant/"/>
<meta property="og:title" content="???????????????????????????????????????"/>
<meta property="og:site_name" content="???????????????????????????????????????"/>
<meta property="og:description" content=""/>
<meta property="og:type" content="website"/>
<meta property="og:image" content="https://www.gov.mo/wp-content/static/img/macau-sar-emblem-128x128.png"/>
<meta property="og:locale" content="zh_hk"/>
</head>
<body class="header-fixed white-enabled yamm-enabled yamm-slim">
	<noscript><iframe src="https://www.googletagmanager.com/ns.html?id=GTM-PTG5DG"
height="0" width="0" style="display:none;visibility:hidden"></iframe></noscript>
	<!--[if lt IE 9]>
			<p class="browserupgrade">???????????????<strong>??????</strong>??????????????????<a href="http://browsehappy.com/" target="_blank">?????????????????????</a>??????????????????</p>
	<![endif]-->
	<!-- BEGIN HEADER-->
<header id="header" class="header  ">
	<div class="headerbar">
		<!-- Brand and toggle get grouped for better mobile display -->
		<div class="headerbar-left">
			<ul class="header-nav header-nav-options">
				<li id="hamburger">
					<a class="btn btn-icon-toggle menubar-toggle" data-toggle="menubar" href="javascript:void(0);">
						<i class="fa fa-bars"></i>
					</a>
				</li>
				<li class="header-nav-brand" >
					<div class="brand-holder">
						<a href="https://www.gov.mo/zh-hant" title="???????????????????????????????????????" class="govmo-logo">
							<!--img src="https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/img/govmo-logo-white.png"-->
							<div class="logo">
								<!--svg width="130" height="24" class="img">
									<image xlink:href="https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/img/gov-mo-logo-output-w-20171110.svg" src="https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/img/govmo-logo-white.png" width="130" height="24" />
								</svg-->
								<div class="text">???????????????????????????????????????</div>
							</div>
						</a>
					</div>
				</li>
			</ul>
		</div>
		<!-- Collect the nav links, forms, and other content for toggling -->
		<div class="headerbar-right">
			<ul class="header-nav header-nav-options">
				<li>
					<!-- Search form -->
															<form id="navbar-search-form" class="navbar-search" role="search" action="https://www.gov.mo/zh-hant/global-search" method="get">
						<div class="form-group">
							<label for="navbar-search-input" class="sr-only">??????</label>
							<input type="text" class="form-control" name="q" placeholder="?????? GOV.MO" pattern=".{2,}" required title="??????????????? 2 ?????????">
						</div>
						<button type="submit" class="btn btn-icon-toggle ink-reaction" title="??????"><i class="fa fa-search"></i></button>
					</form>
									</li>
			</ul><!--end .header-nav-options -->
			<ul class="header-nav header-nav-options">
																													<li class="nav-btn">
						<a href="https://www.gov.mo/zh-hant/about-macau-sar/weather/" title="????????????" class="ink-reaction">
														<i class="pe-is-w-sun-1"></i>
														<span>17??C</span>
						</a>
					</li>
																	<li>
					<a href="https://www.gov.mo/zh-hant/directory/rss-feeds/" title="RSS??????" class="btn btn-icon-toggle ink-reaction">
						<i class="fa fa-rss"></i>
					</a>
				</li>
								<li class="dropdown">
					<a href="javascript:void(0);" class="btn btn-icon-toggle ink-reaction" data-toggle="dropdown" title="????????????" >
						<i class="md md-translate"></i>
					</a>
					<ul id="menu-languages" class="dropdown-menu animation-dock"><li id="menu-item-341186" class="blog-id-4 mlp-language-nav-item menu-item menu-item-type-language menu-item-object-mlp_language mlp-current-language-item menu-item-341186"><a rel="alternate" href="https://www.gov.mo/zh-hant/">????????????</a></li>
<li id="menu-item-341187" class="blog-id-5 mlp-language-nav-item menu-item menu-item-type-language menu-item-object-mlp_language menu-item-341187"><a rel="alternate" href="https://www.gov.mo/zh-hans/">????????????</a></li>
<li id="menu-item-341185" class="blog-id-3 mlp-language-nav-item menu-item menu-item-type-language menu-item-object-mlp_language menu-item-341185"><a rel="alternate" href="https://www.gov.mo/pt/">Portugu??s</a></li>
<li id="menu-item-341184" class="blog-id-2 mlp-language-nav-item menu-item menu-item-type-language menu-item-object-mlp_language menu-item-341184"><a rel="alternate" href="https://www.gov.mo/en/">English</a></li>
</ul>				</li>
			</ul>
			<ul id="header-bar-account" class="header-nav header-nav-profile">
						</ul>
			<ul id="offcanvas-menu-toggle" class="header-nav header-nav-toggle">
				<li>
					<a class="btn btn-icon-toggle ink-reaction" href="#offcanvas-menu" data-toggle="offcanvas" data-backdrop="false">
						<i class="fa fa-bars"></i>
					</a>
				</li>
			</ul>
		</div><!--end #header-navbar-collapse -->
	</div>
			<div class="navbar yamm">
	<div id="navbar-collapse-2" class="navbar-collapse collapse">
		<ul class="nav navbar-nav">
			<li class="nav-btn nav-btn-slim">
				<a class="ink-reaction" href="https://www.gov.mo/zh-hant" title="???????????????????????????????????????"><i class="fa fa-home"></i> ??????</a>
			</li>
			<!--li class="nav-btn nav-btn-slim">
				<a class="ink-reaction" href="https://www.gov.mo/zh-hant/content/egov/" ><i class="fa fa-check-circle"></i> ???????????? </a>
			</li-->
			<li class="nav-btn nav-btn-slim dropdown yamm-fw">
				<a class="ink-reaction" href="#" data-toggle="dropdown" class="dropdown-toggle"><i class="fa fa-info-circle"></i> ???????????? <b class="caret"></b></a>
				<ul class="dropdown-menu">
					<li>
	<div class="yamm-content content-block">
		<div class="container">
			<div class="row">
				<div class="col-xs-12">
					<div class="menu-public-services-container"><ul id="menu-public-services" class="menu"><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62566&entity_id=">?????????????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62579&entity_id=">????????????????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62584&entity_id=">???????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62582&entity_id=">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62577&entity_id=">???????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62591&entity_id=">??????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62598&entity_id=">???????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62620&entity_id=">??????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62615&entity_id=">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62568&entity_id=">??????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62588&entity_id=">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62605&entity_id=">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62573&entity_id=">??????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62570&entity_id=">??????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62611&entity_id=">??????</a> </li></ul></div>				</div>
			</div>
			<div class="row">
				<div class="col-xs-12 explore">
					<a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=&entity_id=">????????????</a>
				</div>
			</div>
		</div>
	</div>
</li>
				</ul>
			</li>
			<li class="nav-btn nav-btn-slim dropdown yamm-fw"><a class="ink-reaction dropdown-toggle" data-toggle="dropdown"href="https://www.gov.mo/zh-hant/about-government/"><i class="fa fa-building"></i> ???????????? <b class="caret"></b></a><ul class="dropdown-menu"><li><div class="yamm-content content-block"><div class="container"><ul class="row"><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/chief-executive-principal-officials-legislature-and-judiciary/">???????????????????????????????????????????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/departments-and-agencies/">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/apm/">?????????????????????????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/apm/apm-entity-index/">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/policy-consultation/">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/content/policy-address/">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-government/special-promotions-index/">????????????</a> </li></ul></div></div></li></ul></li><li class="nav-btn nav-btn-slim dropdown yamm-fw"><a class="ink-reaction dropdown-toggle" data-toggle="dropdown"href="https://www.gov.mo/zh-hant/about-macau-sar/"><i class="fa fa-map-o"></i> ???????????? <b class="caret"></b></a><ul class="dropdown-menu"><li><div class="yamm-content content-block"><div class="container"><ul class="row"><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-macau-sar/weather/">??????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/content/traffic/">??????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-macau-sar/public-holidays/">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/content/culture-and-leisure/">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/content/city-info/">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/about-macau-sar/fact-sheet/">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/content/statistics/">????????????</a> </li></ul></div></div></li></ul></li><li class="nav-btn nav-btn-slim dropdown yamm-fw"><a class="ink-reaction dropdown-toggle" data-toggle="dropdown"href="https://www.gov.mo/zh-hant/announcements/"><i class="fa fa-sticky-note-o"></i> ???????????? <b class="caret"></b></a><ul class="dropdown-menu"><li><div class="yamm-content content-block"><div class="container"><ul class="row"><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/news/">??????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/videos/">??????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="http://www.io.gov.mo/cn/bo/" target="_blank">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="http://www.io.gov.mo/cn/news/list/b/?d=13" target="_blank">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="http://www.io.gov.mo/cn/news/list/b/?d=11" target="_blank">????????????</a> </li></ul></div></div></li></ul></li><li class="nav-btn nav-btn-slim dropdown yamm-fw"><a class="ink-reaction dropdown-toggle" data-toggle="dropdown"href="https://www.gov.mo/zh-hant/content/laws/"><i class="fa fa-balance-scale"></i> ???????????? <b class="caret"></b></a><ul class="dropdown-menu"><li><div class="yamm-content content-block"><div class="container"><ul class="row"><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/content/laws/constitutional-documents/">???????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="http://www.macaolaw.gov.mo/" target="_blank">???????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="http://legismac.safp.gov.mo/legismac/main/main.jsf?lang=zh_TW" target="_blank">??????????????????????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="http://cn.io.gov.mo/Legis/record/2000.aspx" target="_blank">??????????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="http://www.io.gov.mo/cn/legis/int/" target="_blank">????????????</a> </li></ul></div></div></li></ul></li><li class="nav-btn nav-btn-slim dropdown yamm-fw"><a class="ink-reaction dropdown-toggle" data-toggle="dropdown"href="https://www.gov.mo/zh-hant/content/travel/"><i class="fa fa-plane"></i> ???????????? <b class="caret"></b></a><ul class="dropdown-menu"><li><div class="yamm-content content-block"><div class="container"><ul class="row"><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.macaotourism.gov.mo/zh-hant/travelessential" target="_blank">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.macaotourism.gov.mo/zh-hant/sightseeing" target="_blank">??????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.macaotourism.gov.mo/zh-hant/shows-and-entertainment" target="_blank">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.macaotourism.gov.mo/zh-hant/shopping" target="_blank">??????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.macaotourism.gov.mo/zh-hant/events" target="_blank">????????????</a> </li></ul></div></div></li></ul></li><li class="nav-btn nav-btn-slim dropdown yamm-fw"><a class="ink-reaction dropdown-toggle" data-toggle="dropdown"href="https://www.gov.mo/zh-hant/content/business/"><i class="fa fa-suitcase"></i> ??????????????? <b class="caret"></b></a><ul class="dropdown-menu"><li><div class="yamm-content content-block"><div class="container"><ul class="row"><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.ipim.gov.mo/zh-hant/business-investment/" target="_blank">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.ipim.gov.mo/zh-hant/macao-exhibition-and-conference/" target="_blank">??????????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.ipim.gov.mo/zh-hant/smes-business-opportunities-and-services/" target="_blank">????????????????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.ipim.gov.mo/zh-hant/market-information/" target="_blank">????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.economia.gov.mo/zh_TW/web/public/pg_ip?_refresh=true" target="_blank">????????????</a> </li></ul></div></div></li></ul></li><li class="nav-btn nav-btn-slim dropdown yamm-fw"><a class="ink-reaction dropdown-toggle" data-toggle="dropdown"href="https://www.gov.mo/zh-hant/directory/"><i class="fa fa-folder-open-o"></i> ???????????? <b class="caret"></b></a><ul class="dropdown-menu"><li><div class="yamm-content content-block"><div class="container"><ul class="row"><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/directory/mobile-apps/">??????????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/directory/social-media/">??????????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/directory/thematic-websites-directory/">??????????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/directory/rss-feeds/">RSS????????????</a> </li><li class="col-xs-12 col-md-3 col-lg-3"><a href="https://www.gov.mo/zh-hant/directory/forms-download/">????????????</a> </li></ul></div></div></li></ul></li><li class="nav-btn nav-btn-slim dropdown yamm-fw"><a class="ink-reaction dropdown-toggle" data-toggle="dropdown"href="http://www.gov.mo/pcs"><i class="fa fa-users"></i> ???????????? <b class="caret"></b></a><ul class="dropdown-menu"><li><div class="yamm-content content-block"><div class="container"><ul class="row"><li class="col-xs-12 col-md-3 col-lg-3"><a href="http://www.gov.mo/pcs" target="_blank">??????????????????</a> </li></ul></div></div></li></ul></li>		</ul>
    	</div>
</div>
	</header>
<!-- END HEADER-->
	<div id="base-slim">
	<div class="offcanvas">
		<div class="offcanvas">
  <div id="offcanvas-demo-left" class="offcanvas-pane width-6" style="">
    <div class="offcanvas-head">
      <header></header>
      <div class="offcanvas-tools">
        <a class="btn btn-icon-toggle btn-default-light pull-right" data-dismiss="offcanvas">
          <i class="md md-close"></i>
        </a>
      </div>
    </div>
    <div class="nano has-scrollbar" style="height: 939px;"><div class="nano-content" tabindex="0" style="right: -17px;"><div class="offcanvas-body">
    </div></div><div class="nano-pane" style="display: none;"><div class="nano-slider" style="height: 922px; transform: translate(0px, 0px);"></div></div></div>
  </div>
</div>
	</div>
		<div class="offcanvas">
	<!-- BEGIN OFFCANVAS SEARCH -->
	<div id="offcanvas-menu" class="offcanvas-pane width-8" style="">
		<div class="offcanvas-head">
			<header>
				<i class="fa fa-bars"></i> ??????			</header>
			<div class="offcanvas-tools">
				<a class="btn btn-icon-toggle btn-default-light pull-right" data-dismiss="offcanvas">
					<i class="md md-close"></i>
				</a>
			</div>
		</div>
		<div class="nano has-scrollbar" style="height: 605px;">
			<div class="nano-content" tabindex="0" style="right: -17px;">
				<div class="offcanvas-body">
					<ul class="nav navbar-nav nav-pills nav-stacked nav-transparent">
						<li>
							<a class="ink-reaction" href="https://www.gov.mo/zh-hant" title="???????????????????????????????????????"><i class="fa fa-home"></i> ??????</a>
						</li>
						<li>
							<a class="ink-reaction" href="#" data-toggle="dropdown" class="dropdown-toggle"><i class="fa fa-info-circle"></i> ???????????? <b class="caret"></b></a>
							<ul id="menu-public-services-1" class="nav navbar-nav nav-pills nav-stacked nav-transparent dropdown-menu"><li class="ink-reaction"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62566&entity_id=">?????????????????????</a> </li><li class="ink-reaction"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62579&entity_id=">????????????????????????</a> </li><li class="ink-reaction"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62584&entity_id=">???????????????</a> </li><li class="ink-reaction"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62582&entity_id=">????????????</a> </li><li class="ink-reaction"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62577&entity_id=">???????????????</a> </li><li class="ink-reaction"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62591&entity_id=">??????</a> </li><li class="ink-reaction"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62598&entity_id=">???????????????</a> </li><li class="ink-reaction"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62620&entity_id=">??????</a> </li><li class="ink-reaction"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62615&entity_id=">????????????</a> </li><li class="ink-reaction"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62568&entity_id=">??????</a> </li><li class="ink-reaction"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62588&entity_id=">????????????</a> </li><li class="ink-reaction"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62605&entity_id=">????????????</a> </li><li class="ink-reaction"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62573&entity_id=">??????</a> </li><li class="ink-reaction"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62570&entity_id=">??????</a> </li><li class="ink-reaction"><a href="https://www.gov.mo/zh-hant/about-government/service-list/?category_id=62611&entity_id=">??????</a> </li></ul>						</li>
					</ul>
					<ul id="menu-footer-menu-1" class="nav navbar-nav nav-pills nav-stacked nav-transparent"><li><a class="ink-reaction dropdown-toggle" data-toggle="dropdown"href="https://www.gov.mo/zh-hant/about-government/"><i class="fa fa-building"></i> ???????????? <b class="caret"></b></a><ul class="nav nav-pills nav-stacked nav-transparent dropdown-menu"><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/about-government/chief-executive-principal-officials-legislature-and-judiciary/" href="https://www.gov.mo/zh-hant/about-government/chief-executive-principal-officials-legislature-and-judiciary/"><i class=""></i> ???????????????????????????????????????????????????</a></li><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/about-government/departments-and-agencies/" href="https://www.gov.mo/zh-hant/about-government/departments-and-agencies/"><i class=""></i> ????????????</a></li><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/about-government/apm/" href="https://www.gov.mo/zh-hant/about-government/apm/"><i class=""></i> ?????????????????????????????????</a></li><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/about-government/apm/apm-entity-index/" href="https://www.gov.mo/zh-hant/about-government/apm/apm-entity-index/"><i class=""></i> ????????????</a></li><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/about-government/policy-consultation/" href="https://www.gov.mo/zh-hant/about-government/policy-consultation/"><i class=""></i> ????????????</a></li><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/content/policy-address/" href="https://www.gov.mo/zh-hant/content/policy-address/"><i class=""></i> ????????????</a></li><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/about-government/special-promotions-index/" href="https://www.gov.mo/zh-hant/about-government/special-promotions-index/"><i class=""></i> ????????????</a></li></ul></li><li><a class="ink-reaction dropdown-toggle" data-toggle="dropdown"href="https://www.gov.mo/zh-hant/about-macau-sar/"><i class="fa fa-map-o"></i> ???????????? <b class="caret"></b></a><ul class="nav nav-pills nav-stacked nav-transparent dropdown-menu"><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/about-macau-sar/weather/" href="https://www.gov.mo/zh-hant/about-macau-sar/weather/"><i class=""></i> ??????</a></li><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/content/traffic/" href="https://www.gov.mo/zh-hant/content/traffic/"><i class=""></i> ??????</a></li><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/about-macau-sar/public-holidays/" href="https://www.gov.mo/zh-hant/about-macau-sar/public-holidays/"><i class=""></i> ????????????</a></li><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/content/culture-and-leisure/" href="https://www.gov.mo/zh-hant/content/culture-and-leisure/"><i class=""></i> ????????????</a></li><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/content/city-info/" href="https://www.gov.mo/zh-hant/content/city-info/"><i class=""></i> ????????????</a></li><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/about-macau-sar/fact-sheet/" href="https://www.gov.mo/zh-hant/about-macau-sar/fact-sheet/"><i class=""></i> ????????????</a></li><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/content/statistics/" href="https://www.gov.mo/zh-hant/content/statistics/"><i class=""></i> ????????????</a></li></ul></li><li><a class="ink-reaction dropdown-toggle" data-toggle="dropdown"href="https://www.gov.mo/zh-hant/announcements/"><i class="fa fa-sticky-note-o"></i> ???????????? <b class="caret"></b></a><ul class="nav nav-pills nav-stacked nav-transparent dropdown-menu"><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/news/" href="https://www.gov.mo/zh-hant/news/"><i class=""></i> ??????</a></li><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/videos/" href="https://www.gov.mo/zh-hant/videos/"><i class=""></i> ??????</a></li><li><a class="ink-reaction"href="http://www.io.gov.mo/cn/bo/" target="_blank" href="http://www.io.gov.mo/cn/bo/"><i class=""></i> ????????????</a></li><li><a class="ink-reaction"href="http://www.io.gov.mo/cn/news/list/b/?d=13" target="_blank" href="http://www.io.gov.mo/cn/news/list/b/?d=13"><i class=""></i> ????????????</a></li><li><a class="ink-reaction"href="http://www.io.gov.mo/cn/news/list/b/?d=11" target="_blank" href="http://www.io.gov.mo/cn/news/list/b/?d=11"><i class=""></i> ????????????</a></li></ul></li><li><a class="ink-reaction dropdown-toggle" data-toggle="dropdown"href="https://www.gov.mo/zh-hant/content/laws/"><i class="fa fa-balance-scale"></i> ???????????? <b class="caret"></b></a><ul class="nav nav-pills nav-stacked nav-transparent dropdown-menu"><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/content/laws/constitutional-documents/" href="https://www.gov.mo/zh-hant/content/laws/constitutional-documents/"><i class=""></i> ???????????????</a></li><li><a class="ink-reaction"href="http://www.macaolaw.gov.mo/" target="_blank" href="http://www.macaolaw.gov.mo/"><i class=""></i> ???????????????</a></li><li><a class="ink-reaction"href="http://legismac.safp.gov.mo/legismac/main/main.jsf?lang=zh_TW" target="_blank" href="http://legismac.safp.gov.mo/legismac/main/main.jsf?lang=zh_TW"><i class=""></i> ??????????????????????????????</a></li><li><a class="ink-reaction"href="http://cn.io.gov.mo/Legis/record/2000.aspx" target="_blank" href="http://cn.io.gov.mo/Legis/record/2000.aspx"><i class=""></i> ??????????????????</a></li><li><a class="ink-reaction"href="http://www.io.gov.mo/cn/legis/int/" target="_blank" href="http://www.io.gov.mo/cn/legis/int/"><i class=""></i> ????????????</a></li></ul></li><li><a class="ink-reaction dropdown-toggle" data-toggle="dropdown"href="https://www.gov.mo/zh-hant/content/travel/"><i class="fa fa-plane"></i> ???????????? <b class="caret"></b></a><ul class="nav nav-pills nav-stacked nav-transparent dropdown-menu"><li><a class="ink-reaction"href="https://www.macaotourism.gov.mo/zh-hant/travelessential" target="_blank" href="https://www.macaotourism.gov.mo/zh-hant/travelessential"><i class=""></i> ????????????</a></li><li><a class="ink-reaction"href="https://www.macaotourism.gov.mo/zh-hant/sightseeing" target="_blank" href="https://www.macaotourism.gov.mo/zh-hant/sightseeing"><i class=""></i> ??????</a></li><li><a class="ink-reaction"href="https://www.macaotourism.gov.mo/zh-hant/shows-and-entertainment" target="_blank" href="https://www.macaotourism.gov.mo/zh-hant/shows-and-entertainment"><i class=""></i> ????????????</a></li><li><a class="ink-reaction"href="https://www.macaotourism.gov.mo/zh-hant/shopping" target="_blank" href="https://www.macaotourism.gov.mo/zh-hant/shopping"><i class=""></i> ??????</a></li><li><a class="ink-reaction"href="https://www.macaotourism.gov.mo/zh-hant/events" target="_blank" href="https://www.macaotourism.gov.mo/zh-hant/events"><i class=""></i> ????????????</a></li></ul></li><li><a class="ink-reaction dropdown-toggle" data-toggle="dropdown"href="https://www.gov.mo/zh-hant/content/business/"><i class="fa fa-suitcase"></i> ??????????????? <b class="caret"></b></a><ul class="nav nav-pills nav-stacked nav-transparent dropdown-menu"><li><a class="ink-reaction"href="https://www.ipim.gov.mo/zh-hant/business-investment/" target="_blank" href="https://www.ipim.gov.mo/zh-hant/business-investment/"><i class=""></i> ????????????</a></li><li><a class="ink-reaction"href="https://www.ipim.gov.mo/zh-hant/macao-exhibition-and-conference/" target="_blank" href="https://www.ipim.gov.mo/zh-hant/macao-exhibition-and-conference/"><i class=""></i> ??????????????????</a></li><li><a class="ink-reaction"href="https://www.ipim.gov.mo/zh-hant/smes-business-opportunities-and-services/" target="_blank" href="https://www.ipim.gov.mo/zh-hant/smes-business-opportunities-and-services/"><i class=""></i> ????????????????????????</a></li><li><a class="ink-reaction"href="https://www.ipim.gov.mo/zh-hant/market-information/" target="_blank" href="https://www.ipim.gov.mo/zh-hant/market-information/"><i class=""></i> ????????????</a></li><li><a class="ink-reaction"href="https://www.economia.gov.mo/zh_TW/web/public/pg_ip?_refresh=true" target="_blank" href="https://www.economia.gov.mo/zh_TW/web/public/pg_ip?_refresh=true"><i class=""></i> ????????????</a></li></ul></li><li><a class="ink-reaction dropdown-toggle" data-toggle="dropdown"href="https://www.gov.mo/zh-hant/directory/"><i class="fa fa-folder-open-o"></i> ???????????? <b class="caret"></b></a><ul class="nav nav-pills nav-stacked nav-transparent dropdown-menu"><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/directory/mobile-apps/" href="https://www.gov.mo/zh-hant/directory/mobile-apps/"><i class=""></i> ??????????????????</a></li><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/directory/social-media/" href="https://www.gov.mo/zh-hant/directory/social-media/"><i class=""></i> ??????????????????</a></li><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/directory/thematic-websites-directory/" href="https://www.gov.mo/zh-hant/directory/thematic-websites-directory/"><i class=""></i> ??????????????????</a></li><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/directory/rss-feeds/" href="https://www.gov.mo/zh-hant/directory/rss-feeds/"><i class=""></i> RSS????????????</a></li><li><a class="ink-reaction"href="https://www.gov.mo/zh-hant/directory/forms-download/" href="https://www.gov.mo/zh-hant/directory/forms-download/"><i class=""></i> ????????????</a></li></ul></li><li><a class="ink-reaction dropdown-toggle" data-toggle="dropdown"href="http://www.gov.mo/pcs"><i class="fa fa-users"></i> ???????????? <b class="caret"></b></a><ul class="nav nav-pills nav-stacked nav-transparent dropdown-menu"><li><a class="ink-reaction"href="http://www.gov.mo/pcs" target="_blank" href="http://www.gov.mo/pcs"><i class=""></i> ??????????????????</a></li></ul></li></ul>				</div>
				<div class="offcanvas-body style-default-light">
					<div id="header-bar-account-offcanvas">
						<img src="https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/img/spinner-rosetta-gray-26x26.gif" />
					</div>
				</div>
			</div>
			<div class="nano-pane">
				<div class="nano-slider" style="height: 357px; transform: translate(0px, 0px);"></div>
			</div>
		</div>
	</div>
</div>
	<main id="content">
		<div class="yamm-separator"></div>
								<div id="copc-messages-placeholder" data-url="https://www.gov.mo/zh-hant/copc-messages/?alert=1"></div>
				<div id="weather-typhoon-info-placeholder" data-url="https://www.gov.mo/zh-hant/about-macau-sar/weather/typhoon-info/?alert=1"></div>
						<section id="welcome" class="style-default-light">
	<div class="container">
		<div class="row">
			<div class="col-xs-12 col-sm-6">
												<!-- Search form -->
												<form id="search-form" role="search" method="get" action="https://www.gov.mo/zh-hant/global-search">
					<div class="input-group">
						<input type="text" class="form-control" name="q" placeholder="?????? GOV.MO" pattern=".{2,}" required title="??????????????? 2 ?????????">
						<span class="input-group-btn">
							<button type="submit" id="search-btn" class="btn btn-primary ink-reaction" title="??????"><i class="fa fa-search"></i></button>
						</span>
					</div>
				</form>
				<script type="application/ld+json">
				{
					"@context": "http://schema.org",
					"@type": "WebSite",
					"url": "https://www.gov.mo/zh-hant/",
					"potentialAction": {
						"@type": "SearchAction",
						"target": "https://www.gov.mo/zh-hant/global-search?q={search_term_string}",
						"query-input": "required name=search_term_string"
					}
				}
				</script>
							</div>
			<div class="col-xs-12 col-sm-6">
								<div id="quick-links" style="margin-top: 2px">
					<span class="sr-only">???????????????</span>
					<div class="menu-quick-links-container"><ul id="menu-quick-links" class="menu"><li id="menu-item-332319" class="menu-item menu-item-type-post_type menu-item-object-public_services_page menu-item-332319"><a href="https://www.gov.mo/zh-hant/services/ps-1047/">???????????????????????????????????????</a></li>
<li id="menu-item-345499" class="menu-item menu-item-type-post_type menu-item-object-page menu-item-345499"><a href="https://www.gov.mo/zh-hant/app/download/">?????????????????????</a></li>
<li id="menu-item-332400" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-332400"><a target="_blank" rel="noopener noreferrer" href="https://www.gov.mo/egov">????????????????????????</a></li>
<li id="menu-item-316343" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-316343"><a target="_blank" rel="noopener noreferrer" href="https://www.ssm.gov.mo/apps1/PreventCOVID-19/ch.aspx">????????????</a></li>
<li id="menu-item-321147" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-321147"><a target="_blank" rel="noopener noreferrer" href="https://app.ssm.gov.mo/phd/apps/healthdeclaration/?lang=ch">???????????????</a></li>
<li id="menu-item-343225" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-343225"><a target="_blank" rel="noopener noreferrer" href="https://eservice.ssm.gov.mo/rnatestbook/V21/">??????????????????????????????</a></li>
<li id="menu-item-305406" class="menu-item menu-item-type-post_type menu-item-object-page menu-item-305406"><a href="https://www.gov.mo/zh-hant/about-macau-sar/public-holidays/">????????????</a></li>
<li id="menu-item-243067" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-243067"><a target="_blank" rel="noopener noreferrer" href="https://www.gov.mo/zh-hant/apm-entity-page/apm-entity-org-chart">???????????????</a></li>
<li id="menu-item-218795" class="menu-item menu-item-type-post_type menu-item-object-page menu-item-218795"><a href="https://www.gov.mo/zh-hant/about-government/departments-and-agencies/">????????????</a></li>
</ul></div>				</div>
							</div>
		</div>
	</div>
</section>
						<section id="news" class="style-default-light">
	<div class="container">
		<div class="section-head">
			<div class="row">
				<div class="col-xs-10">
					<h2>??????</h2>
				</div>
				<div class="col-xs-2">
					<a class="btn btn-icon-toggle btn-collapse pull-right"><i class="fa fa-angle-down"></i></a>
				</div>
			</div>
		</div>
		<div class="section-body">
			<div class="row news--bg-bottom-left">
				<div class="col-xs-12 col-sm-6 news--bg-top-right">
					<div class="left-col">
																								<h3 class="text-xxl text-light"><a href="https://www.gov.mo/zh-hant/news/360850/" rel="bookmark" title="Permanent Link to ??????????????????????????????   ?????????????????????????????????">??????????????????????????????   ?????????????????????????????????</a></h3>
						<time class="date" datetime="2021-01-18 13:06:00">2021???1???18??? 13:06</time>
																		<a class="caption">
							<img width="1278" height="852" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/p21arf80rh.jpeg" class="attachment-full size-full landscape wp-post-image" alt="" loading="lazy" srcset="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/p21arf80rh.jpeg 1278w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/p21arf80rh-300x200.jpeg 300w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/p21arf80rh-1024x683.jpeg 1024w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/p21arf80rh-768x512.jpeg 768w" sizes="(max-width: 1278px) 100vw, 1278px" />							<span>?????????????????????????????????</span>
						</a>
												<p>????????????1???17????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????<br /><a href="https://www.gov.mo/zh-hant/news/360850/">&hellip;</a></p>
											</div>
				</div>
				<div class="col-xs-12 col-sm-6">
					<div class="right-col">
					<h3 class="date">?????????, 1???18???</h3><ul class="timeline collapse-lg timeline-hairline">					<li class="timeline-inverted">
						<div class="timeline-circ circ-xl style-primary">
															<i class="fa fa-photo"></i>
													</div>
						<div class="timeline-entry">
							<div class="card style-default-bright">
								<div class="card-body small-padding">
									<a href="https://www.gov.mo/zh-hant/news/360596/" rel="bookmark" title="Permanent Link to ???????????????2020?????????">???????????????2020?????????</a><br/>
																		<time class="date" datetime="2021-01-18 12:00:00">2021???1???18??? 12:00</time>
								</div><!--end .card-body -->
							</div><!--end .card -->
						</div><!--end .timeline-entry -->
					</li>
										<li class="timeline-inverted">
						<div class="timeline-circ circ-xl style-primary">
															<i class="fa fa-quote-left"></i>
													</div>
						<div class="timeline-entry">
							<div class="card style-default-bright">
								<div class="card-body small-padding">
									<a href="https://www.gov.mo/zh-hant/news/360805/" rel="bookmark" title="Permanent Link to ?????????????????????????????????????????????????????????????????? (01???17???08:00???01???18???08:00)">?????????????????????????????????????????????????????????????????? (01???17???08:00???01???18???08:00)</a><br/>
																		<time class="date" datetime="2021-01-18 10:15:00">2021???1???18??? 10:15</time>
								</div><!--end .card-body -->
							</div><!--end .card -->
						</div><!--end .timeline-entry -->
					</li>
										<li class="timeline-inverted">
						<div class="timeline-circ circ-xl style-primary">
															<i class="fa fa-photo"></i>
													</div>
						<div class="timeline-entry">
							<div class="card style-default-bright">
								<div class="card-body small-padding">
									<a href="https://www.gov.mo/zh-hant/news/360812/" rel="bookmark" title="Permanent Link to ?????????2,000????????????????????????????????????????????????????????????">?????????2,000????????????????????????????????????????????????????????????</a><br/>
																		<time class="date" datetime="2021-01-18 09:27:00">2021???1???18??? 09:27</time>
								</div><!--end .card-body -->
							</div><!--end .card -->
						</div><!--end .timeline-entry -->
					</li>
										<li class="timeline-inverted">
						<div class="timeline-circ circ-xl style-primary">
															<i class="fa fa-photo"></i>
													</div>
						<div class="timeline-entry">
							<div class="card style-default-bright">
								<div class="card-body small-padding">
									<a href="https://www.gov.mo/zh-hant/news/360799/" rel="bookmark" title="Permanent Link to ????????????????????????????????????????????????????????????">????????????????????????????????????????????????????????????</a><br/>
																		<time class="date" datetime="2021-01-18 09:25:00">2021???1???18??? 09:25</time>
								</div><!--end .card-body -->
							</div><!--end .card -->
						</div><!--end .timeline-entry -->
					</li>
					</ul><h3 class="date">?????????, 1???17???</h3><ul class="timeline collapse-lg timeline-hairline">					<li class="timeline-inverted">
						<div class="timeline-circ circ-xl style-primary">
															<i class="fa fa-photo"></i>
													</div>
						<div class="timeline-entry">
							<div class="card style-default-bright">
								<div class="card-body small-padding">
									<a href="https://www.gov.mo/zh-hant/news/360777/" rel="bookmark" title="Permanent Link to ???????????????2021/01/17?????????????????????????????????????????????">???????????????2021/01/17?????????????????????????????????????????????</a><br/>
																		<time class="date" datetime="2021-01-17 20:19:00">2021???1???17??? 20:19</time>
								</div><!--end .card-body -->
							</div><!--end .card -->
						</div><!--end .timeline-entry -->
					</li>
										</div>
				</div>
			</div>
			<hr />
			<div class="row">
				<div class="col-xs-12 explore">
										<a href="https://www.gov.mo/zh-hant/news/">????????????</a>
				</div>
			</div>
		</div>
	</div>
</section>
		<section id="promotions" class="style-default-bright">
	<div class="container">
		<div class="section-head">
			<div class="row">
				<div class="col-xs-10">
					<h2>????????????</h2>
				</div>
				<div class="col-xs-2">
					<a class="btn btn-icon-toggle btn-collapse pull-right"><i class="fa fa-angle-down"></i></a>
				</div>
			</div>
		</div>
		<div class="section-body">
			<div class="row">
								<div class="col-xs-12">
																			<div id="owl-promotions" class="owl-carousel">
																							<div class="item">
									<div class="promotion--item">
										<div class="card">
																							<div class="card-img">
													<a href="https://www.gov.mo/zh-hant/promotions/358999/" title="????????????">
														<img width="272" height="136" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/icon-????????????_1140X570_150px-272x136.jpg" class="attachment-thumbnail-promotions size-thumbnail-promotions" alt="" loading="lazy" srcset="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/icon-????????????_1140X570_150px-272x136.jpg 272w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/icon-????????????_1140X570_150px-300x150.jpg 300w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/icon-????????????_1140X570_150px-1024x512.jpg 1024w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/icon-????????????_1140X570_150px-768x384.jpg 768w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/icon-????????????_1140X570_150px-1536x769.jpg 1536w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/icon-????????????_1140X570_150px-2048x1025.jpg 2048w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/icon-????????????_1140X570_150px-1140x570.jpg 1140w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/icon-????????????_1140X570_150px-570x285.jpg 570w" sizes="(max-width: 272px) 100vw, 272px" />													</a>
												</div>
																						<div class="card-body mh">
												<div class="promotion--item-title"><a href="https://www.gov.mo/zh-hant/promotions/358999/">????????????</a></div>
											</div>
										</div>
									</div>
								</div>
																							<div class="item">
									<div class="promotion--item">
										<div class="card">
																							<div class="card-img">
													<a href="https://www.gov.mo/zh-hant/promotions/352726/" title="2021???????????????????????????">
														<img width="272" height="136" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/11/promo_banner_b_zh_1140x570-272x136.jpg" class="attachment-thumbnail-promotions size-thumbnail-promotions" alt="" loading="lazy" srcset="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/11/promo_banner_b_zh_1140x570-272x136.jpg 272w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/11/promo_banner_b_zh_1140x570-300x150.jpg 300w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/11/promo_banner_b_zh_1140x570-1024x512.jpg 1024w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/11/promo_banner_b_zh_1140x570-768x384.jpg 768w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/11/promo_banner_b_zh_1140x570-570x285.jpg 570w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/11/promo_banner_b_zh_1140x570.jpg 1140w" sizes="(max-width: 272px) 100vw, 272px" />													</a>
												</div>
																						<div class="card-body mh">
												<div class="promotion--item-title"><a href="https://www.gov.mo/zh-hant/promotions/352726/">2021???????????????????????????</a></div>
											</div>
										</div>
									</div>
								</div>
																							<div class="item">
									<div class="promotion--item">
										<div class="card">
																							<div class="card-img">
													<a href="https://www.gov.mo/zh-hant/promotions/353742/" title="????????????????????????????????????????????????">
														<img width="272" height="136" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/11/????????????????????????????????????-272x136.jpg" class="attachment-thumbnail-promotions size-thumbnail-promotions" alt="" loading="lazy" srcset="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/11/????????????????????????????????????-272x136.jpg 272w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/11/????????????????????????????????????-300x150.jpg 300w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/11/????????????????????????????????????-1024x512.jpg 1024w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/11/????????????????????????????????????-768x384.jpg 768w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/11/????????????????????????????????????-570x285.jpg 570w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/11/????????????????????????????????????.jpg 1140w" sizes="(max-width: 272px) 100vw, 272px" />													</a>
												</div>
																						<div class="card-body mh">
												<div class="promotion--item-title"><a href="https://www.gov.mo/zh-hant/promotions/353742/">????????????????????????????????????????????????</a></div>
											</div>
										</div>
									</div>
								</div>
																							<div class="item">
									<div class="promotion--item">
										<div class="card">
																							<div class="card-img">
													<a href="https://www.gov.mo/zh-hant/promotions/355251/" title="???????????????????????????????????????2021-2030???????????????">
														<img width="272" height="136" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/banner-272x136.jpg" class="attachment-thumbnail-promotions size-thumbnail-promotions" alt="" loading="lazy" srcset="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/banner-272x136.jpg 272w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/banner-300x150.jpg 300w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/banner-1024x512.jpg 1024w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/banner-768x384.jpg 768w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/banner-570x285.jpg 570w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/banner.jpg 1140w" sizes="(max-width: 272px) 100vw, 272px" />													</a>
												</div>
																						<div class="card-body mh">
												<div class="promotion--item-title"><a href="https://www.gov.mo/zh-hant/promotions/355251/">???????????????????????????????????????2021-2030???????????????</a></div>
											</div>
										</div>
									</div>
								</div>
																							<div class="item">
									<div class="promotion--item">
										<div class="card">
																							<div class="card-img">
													<a href="https://www.gov.mo/zh-hant/promotions/355242/" title="2020???????????????_?????????????????????">
														<img width="272" height="136" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/03_???????????????????????????_1140x570px-272x136.jpg" class="attachment-thumbnail-promotions size-thumbnail-promotions" alt="" loading="lazy" srcset="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/03_???????????????????????????_1140x570px-272x136.jpg 272w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/03_???????????????????????????_1140x570px-300x150.jpg 300w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/03_???????????????????????????_1140x570px-1024x512.jpg 1024w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/03_???????????????????????????_1140x570px-768x384.jpg 768w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/03_???????????????????????????_1140x570px-570x285.jpg 570w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/03_???????????????????????????_1140x570px.jpg 1140w" sizes="(max-width: 272px) 100vw, 272px" />													</a>
												</div>
																						<div class="card-body mh">
												<div class="promotion--item-title"><a href="https://www.gov.mo/zh-hant/promotions/355242/">2020???????????????_?????????????????????</a></div>
											</div>
										</div>
									</div>
								</div>
																							<div class="item">
									<div class="promotion--item">
										<div class="card">
																							<div class="card-img">
													<a href="https://www.gov.mo/zh-hant/promotions/358983/" title="??????????????????">
														<img width="272" height="136" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/7_???????????????_1140x570px-272x136.jpg" class="attachment-thumbnail-promotions size-thumbnail-promotions" alt="" loading="lazy" srcset="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/7_???????????????_1140x570px-272x136.jpg 272w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/7_???????????????_1140x570px-300x150.jpg 300w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/7_???????????????_1140x570px-1024x512.jpg 1024w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/7_???????????????_1140x570px-768x384.jpg 768w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/7_???????????????_1140x570px-570x285.jpg 570w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2021/01/7_???????????????_1140x570px.jpg 1140w" sizes="(max-width: 272px) 100vw, 272px" />													</a>
												</div>
																						<div class="card-body mh">
												<div class="promotion--item-title"><a href="https://www.gov.mo/zh-hant/promotions/358983/">??????????????????</a></div>
											</div>
										</div>
									</div>
								</div>
																							<div class="item">
									<div class="promotion--item">
										<div class="card">
																							<div class="card-img">
													<a href="https://www.gov.mo/zh-hant/promotions/343856/" title="??????????????? ???????????????">
														<img width="272" height="136" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/09/????????????_Banner-04-272x136.jpg" class="attachment-thumbnail-promotions size-thumbnail-promotions" alt="" loading="lazy" srcset="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/09/????????????_Banner-04-272x136.jpg 272w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/09/????????????_Banner-04-300x150.jpg 300w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/09/????????????_Banner-04-1024x512.jpg 1024w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/09/????????????_Banner-04-768x384.jpg 768w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/09/????????????_Banner-04-570x285.jpg 570w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/09/????????????_Banner-04.jpg 1141w" sizes="(max-width: 272px) 100vw, 272px" />													</a>
												</div>
																						<div class="card-body mh">
												<div class="promotion--item-title"><a href="https://www.gov.mo/zh-hant/promotions/343856/">??????????????? ???????????????</a></div>
											</div>
										</div>
									</div>
								</div>
																							<div class="item">
									<div class="promotion--item">
										<div class="card">
																							<div class="card-img">
													<a href="https://www.gov.mo/zh-hant/promotions/345500/" title="???????????????????????????????????????">
														<img width="272" height="136" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/09/yihutong-272x136.png" class="attachment-thumbnail-promotions size-thumbnail-promotions" alt="" loading="lazy" srcset="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/09/yihutong-272x136.png 272w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/09/yihutong-300x150.png 300w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/09/yihutong-1024x512.png 1024w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/09/yihutong-768x384.png 768w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/09/yihutong-570x285.png 570w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/09/yihutong.png 1141w" sizes="(max-width: 272px) 100vw, 272px" />													</a>
												</div>
																						<div class="card-body mh">
												<div class="promotion--item-title"><a href="https://www.gov.mo/zh-hant/promotions/345500/">???????????????????????????????????????</a></div>
											</div>
										</div>
									</div>
								</div>
																							<div class="item">
									<div class="promotion--item">
										<div class="card">
																																			<iframe title="2021????????????????????????????????????" src="https://www.youtube.com/embed/videoseries?list=PLDnrjBM1_dFFdsmSelBlq4TPDvlqPQxmu&amp;rel=0" frameborder="0" allow="accelerometer; encrypted-media; gyroscope; picture-in-picture&rel=0&modestbranding=1" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
																						<div class="card-body mh">
												<div class="promotion--item-title"><a href="https://www.gov.mo/zh-hant/promotions/352426/">2021????????????????????????????????????</a></div>
											</div>
										</div>
									</div>
								</div>
																							<div class="item">
									<div class="promotion--item">
										<div class="card">
																																			<iframe title="?????????????????????????????????????????????" src="https://www.youtube.com/embed/sBAW2SekAO0?rel=0&modestbranding=1" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
																						<div class="card-body mh">
												<div class="promotion--item-title"><a href="https://www.gov.mo/zh-hant/promotions/336613/">?????????????????????????????????????????????</a></div>
											</div>
										</div>
									</div>
								</div>
																							<div class="item">
									<div class="promotion--item">
										<div class="card">
																							<div class="card-img">
													<a href="https://www.gov.mo/zh-hant/promotions/314929/" title="????????????????????????????????????????????????????????????">
														<img width="272" height="136" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/01/ncv.macau_01-272x136.jpg" class="attachment-thumbnail-promotions size-thumbnail-promotions" alt="" loading="lazy" srcset="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/01/ncv.macau_01-272x136.jpg 272w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/01/ncv.macau_01-300x150.jpg 300w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/01/ncv.macau_01-1024x512.jpg 1024w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/01/ncv.macau_01-768x384.jpg 768w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/01/ncv.macau_01-570x285.jpg 570w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/01/ncv.macau_01.jpg 1140w" sizes="(max-width: 272px) 100vw, 272px" />													</a>
												</div>
																						<div class="card-body mh">
												<div class="promotion--item-title"><a href="https://www.gov.mo/zh-hant/promotions/314929/">????????????????????????????????????????????????????????????</a></div>
											</div>
										</div>
									</div>
								</div>
																							<div class="item">
									<div class="promotion--item">
										<div class="card">
																							<div class="card-img">
													<a href="https://www.gov.mo/zh-hant/promotions/327202/" title="?????????????????????????????????????????????">
														<img width="272" height="136" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/04/eesn-272x136.jpg" class="attachment-thumbnail-promotions size-thumbnail-promotions" alt="" loading="lazy" srcset="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/04/eesn-272x136.jpg 272w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/04/eesn-300x150.jpg 300w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/04/eesn-1024x512.jpg 1024w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/04/eesn-768x384.jpg 768w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/04/eesn-570x285.jpg 570w, https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/04/eesn.jpg 1140w" sizes="(max-width: 272px) 100vw, 272px" />													</a>
												</div>
																						<div class="card-body mh">
												<div class="promotion--item-title"><a href="https://www.gov.mo/zh-hant/promotions/327202/">?????????????????????????????????????????????</a></div>
											</div>
										</div>
									</div>
								</div>
																											</div>
													</div>
			</div>
			<div class="row">
				<div class="col-xs-12 explore">
					<a href="https://www.gov.mo/zh-hant/about-government/special-promotions-index/?category_id=&entity_id=">????????????</a>
				</div>
			</div>
		</div>
	</div>
</section>
		<section id="events--leisure-and-culture" class="style-default-bright">
	<div class="container">
		<div class="section-head">
			<div class="row">
				<div class="col-xs-10">
					<h2>??????</h2>
				</div>
				<div class="col-xs-2">
					<a class="btn btn-icon-toggle btn-collapse pull-right"><i class="fa fa-angle-down"></i></a>
				</div>
			</div>
		</div>
		<div class="section-body">
			<div class="row">
				<div class="col-xs-12">
																				<div id="owl-events" class="owl-carousel">
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/341703/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/08/NW-182120-300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										9???15??? ??? 7???6???									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/341703/">?????????????????????</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/347044/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/10/IC-182524-300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										10???17??? ??? 8???29???									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/347044/">????????????????????????????????????</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/350344/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/11/NW-182658-300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										11???6??? ??? 3???22???									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/350344/">??????????????????????????????????????????&hellip;</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/345847/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/09/NW-182420-300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										11???29??? ??? 2???6???									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/345847/">?????????????????????</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/355252/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/NW-182865-300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										12???4??? ??? 3???14???									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/355252/">?????????????????????????????????????????????????????????????????????&hellip;</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/356550/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/NW-182908-300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										12???11??? ??? 3???28???									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/356550/">??????????????????????????????</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/356366/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/NW-182892-300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										12???17??? ??? 3???14???									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/356366/">????????????&mdash;&mdash;????????????????????????????????????&hellip;</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/356968/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/NW-182811-700-300x300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										2-19/1 |&nbsp;??????????????????<br />
9:00-11:00???????????????&nbsp;|&nbsp;????????????????????????									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/356968/">???????????????????????????</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/345969/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/09/NW-182440-300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										1???6??? ??? 6???9???									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/345969/">????????????????????????</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/356947/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/NW-182791-700-300x300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										19-21/1&nbsp;|&nbsp;???????????????<br />
20:30-22:00&nbsp;|&nbsp;??????????????????3003???									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/356947/">????????????</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/356956/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/NW-182796-700-300x300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										20-21/1&nbsp;|&nbsp;???????????????<br />
20:00&nbsp;|&nbsp;?????????????????????????????????????????????*???<br />
&nbsp;									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/356956/">????????????</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/357016/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/NW-182874-700-300x300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										20-25/1&nbsp;|&nbsp;???????????????<br />
12:00-20:00&nbsp;| ?????????????????????????????????????????????<br />
<br />
26-31/1&nbsp;|&nbsp;???????????????<br />
12:00-20:00&nbsp;|&nbsp;???????????????									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/357016/">???????????????</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/356935/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/NW-182792-700-300x300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										21/1&nbsp;|&nbsp;?????????<br />
15:00&nbsp;|&nbsp;???????????????									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/356935/">????????????</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/356938/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/NW-182789-700-300x300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										22/1 | ?????????&nbsp;&nbsp;<br />
11:00&nbsp;| ???????????????<br />
15:00&nbsp;| ??????????????? &nbsp;<br />
<br />
23/1 | ?????????<br />
11:00&nbsp;| ???????????????<br />
15:00&nbsp;| ??????????????????									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/356938/">????????????</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/356941/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/NW-182804-700-300x300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										22/1 | ?????????&nbsp;&nbsp;<br />
13:30&nbsp;| ????????????<br />
<br />
23/1 | ?????????<br />
13:30&nbsp;| ?????????????????????????????????									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/356941/">????????????</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/356944/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/NW-182790-700-300x300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										22-23/1&nbsp;|&nbsp;???????????????<br />
20:00&nbsp;|&nbsp;???????????????????????????????????????1???									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/356944/">?????????????????????????????????????????????&hellip;</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/356974/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/NW-182907-700-300x300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										22-24/1&nbsp;|&nbsp;???????????????<br />
21:00&nbsp;|&nbsp;????????????????????????									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/356974/">?????????????????????????????????????????????&hellip;</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/342714/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/09/IC-182195-300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										1???23???									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/342714/">????????????</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/356950/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/NW-182798-700-300x300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										23/1&nbsp;|&nbsp;?????????<br />
21:45-23:45&nbsp;|&nbsp;M?????????									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/356950/">?????????</a></div>
								</div>
							</div>
						</div>
					</div>
															<div class="item">
						<div class="event--item">
							<div class="card">
								<div class="card-head event--item-head">
																											<a href="https://www.gov.mo/zh-hant/event/356959/">
										<div class="event--item-img" style="background-image: url(https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/12/NW-182799-700-300x300.jpg)"></div>
									</a>
																	</div>
								<div class="card-body small-padding mh-events event--item-body">
									<div class="event--item-date">
										23-24/1&nbsp;|&nbsp;???????????????<br />
20:00&nbsp;|&nbsp;???????????????									</div>
									<div class="event--item-title"><a href="https://www.gov.mo/zh-hant/event/356959/">????????????</a></div>
								</div>
							</div>
						</div>
					</div>
									</div>
												</div>
			</div>
			<div class="row">
				<div class="col-xs-12 explore">
					<a href="https://www.gov.mo/zh-hant/event-list/?type=leisure-and-culture&category_id=&entity_id=&start_date=&end_date=">????????????</a>
				</div>
			</div>
		</div>
	</div>
</section>
								<section id="cip" class="style-default-light">
	<div class="container">
		<div class="section-head">
			<div class="row">
				<div class="col-xs-10">
					<h2>????????????</h2>
				</div>
				<div class="col-xs-2">
					<a class="btn btn-icon-toggle btn-collapse pull-right"><i class="fa fa-angle-down"></i></a>
				</div>
			</div>
		</div>
		<div class="section-body">
			<div class="row">
				<div class="col-xs-12 col-md-4 contact content-block">
					<div class="row">
						<div class="col-xs-12 content-block">
							<h3>??????????????????</h3>
							<div class="row">
								<div class="col-xs-12 content-block">
									<p>???????????????????????????????????????????????????????????????????????????????????????</p>
									<p><span class="text-xxl"><i class="fa fa-phone"></i> 8866 8866</span></p>
									<p><span class=""><a href="https://www.safp.gov.mo/zh-hant/gcms/suggestionbox/cip/page" target="_blank"><i class="fa fa-envelope"></i> ????????????</span></a></p>
								</div>
							</div>
						</div>
					</div>
				</div>
				<div class="col-xs-12 col-md-8 top-enquiries content-block">
					<div class="row">
						<div class="col-xs-12 content-block">
							<h3>????????????</h3>
						</div>
					</div>
															<ul>
												<li><a href="https://www.gov.mo/zh-hant/cip-faq/316544/">??????????????????????????????????????????????????????</a></li>
												<li><a href="https://www.gov.mo/zh-hant/cip-faq/316546/">??????????????????????????????????????????????????????</a></li>
												<li><a href="https://www.gov.mo/zh-hant/cip-faq/321279/">?????????????????????????????????????????????</a></li>
												<li><a href="https://www.gov.mo/zh-hant/cip-faq/218623/">?????????????????????????????????????????????????????????</a></li>
												<li><a href="https://www.gov.mo/zh-hant/cip-faq/274444/">??????????????????????????????????????????</a></li>
																	</ul>
										<div class="row">
						<div class="col-xs-12 explore">
							<!--a href="#">View all</a-->
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</section>
		<section id="banners" class="style-default-bright">
	<div class="container">
		<div class="section-head">
			<div class="row">
				<div class="col-xs-10">
					<h2>??????</h2>
				</div>
				<div class="col-xs-2">
					<a class="btn btn-icon-toggle btn-collapse pull-right"><i class="fa fa-angle-down"></i></a>
				</div>
			</div>
		</div>
		<div class="section-body">
			<div class="row">
				<div class="col-xs-12">
									<div id="row-banners">
																																												<div class="row-banner row-banner-1">
																			<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="https://www.gov.mo/zh-hant/content/273811-2/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2019/03/?????????194X64.jpg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="https://photo.gcs.gov.mo/zh/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/06/banner194x64pixels.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="https://www.iam.gov.mo/canil/c/lcsa/detail.aspx" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2020/07/???????????????web_banner_194x64px_C.png" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="https://www.gov.mo/zh-hant/content/275403-2/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2019/03/??????????????????-icon_CN.png" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="https://www.gov.mo/zh-hant/content/framework-agreement-guangdong-macao/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/sign_icon2_tc-1.jpeg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="https://www.gov.mo/zh-hant/content/272645-2/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2019/02/gov.mo-banner_long-wording.jpg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://telecommunications.ctt.gov.mo/web/tc/mediainfo/notice/2046-prepaidcard" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2019/12/?????????_GOV_MO194X64.jpg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="https://telecommunications.ctt.gov.mo/smarthotel/?lc=zh_TW" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2019/11/GOVMO_194X64.jpg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="https://www.easygo.gov.mo" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2018/06/??????2-EasyGo-GOV-banner2_c.jpg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="https://www.gov.mo/zh-hant/services/ps-1047/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2019/01/bannerButton-07a.jpg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.gce.gov.mo/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/20111111_172523_772-1.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.elections.gov.mo/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/20151027_111022_105-1.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.elections.gov.mo/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/ece_zh_TW.jpg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.safp.gov.mo/safptc/servicecitizen/WCM_043362" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/20151027_130223_308-1.jpeg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.basiclaw.gov.mo/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/20150901_115342_458-2.png" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://concurso-uni.safp.gov.mo/zh-hant/content/286" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/20160614_093120_106-2.png" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.safp.gov.mo/zh-hant/gcms/report/table" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/tripicon.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.planocp.gov.mo/2020/default_c.html" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/planocp_2020_c.gif" /></a>
															</div>
																																							</div>
																																			<div class="row-banner row-banner-2">
																			<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="https://www.facebook.com/macaogcs/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/Portal_facebookGcs.jpg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="https://www.bolsas.gov.mo" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/20170421_101135_888-1.jpeg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.platformchinaplp.mo/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2018/01/185x60pixel.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.fsm.gov.mo/psp/pspmonitor/mobile/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2018/01/portal.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.macaupanda.org.mo" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/panda_banner_185_60.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="https://www.cityguide.gov.mo" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2018/12/??????5-CG19Banner_GOV.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.vs.gov.mo/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/image_galleryuuidb85ab902-4195-44e4-b8f4-e1413c7eed91groupid10128t1254888275223-1.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://carreira.safp.gov.mo/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/CarreiraLogo-1.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.cpsp.gov.mo/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/20170331_115932_905-1.jpeg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.fss.gov.mo/zh-hant/rpc/rpc-intro" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/20121018_144546_32-1.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.gov.mo/basiclaw23/public/html.jsf" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/m-1-1-1.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="https://www.csraem.gov.mo/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2019/03/????????????????????????.jpg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.dsat.gov.mo/dsat/realtime.aspx" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/transito_live10-1.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.cip.gov.mo/web/public/index.jsf" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/CIP-1.jpg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.wh.mo/cn/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/bannerwh185C-1.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.dsaj.gov.mo/RSCPL/Rscpl.aspx?modulename=WebModules/Intro.ascx" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/com-banner185x60-1.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.macaolaw.gov.mo/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/MacaoLaw-1.jpg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="https://www.fsm.gov.mo/webticket/default.aspx" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/banner_img02-1.gif" /></a>
															</div>
																																							</div>
																																			<div class="row-banner row-banner-3">
																			<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.mif.com.mo" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/banner_img12-2.jpg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.cepa.gov.mo" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/20100629_173158_774-1.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://macaoideas.ipim.gov.mo/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/macaoideaslogo185x60.jpg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.pprd.org.cn/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/pprd_banner-2.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.decmacau.pt/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/banner_img16-1-e1505895705329-1.jpg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.macauhub.com.mo/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/macauhub_safp_C-1.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://bo.io.gov.mo/bo/i/99/31/codcomcn/" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/c-1-1.jpg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.esigntrust.com" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/eSignTrust-1.jpg" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://www.tedmev.com" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/Tedmev-icon_2010_c_194x64.gif" /></a>
															</div>
																																											<div class="col-xs-3 col-sm-3 col-md-2 no-padding">
																								<a href="http://seps.ctt.gov.mo/chi/main.shtml?locale=zh" target="_blank" rel="nofollow"><img alt="banner" class="banners--item-img" src="https://www.gov.mo/zh-hant/wp-content/uploads/sites/4/2017/09/EPCMlogo-1.jpg" /></a>
															</div>
																													</div>
								</div>
							</div>
		</div>
	</div>
</section>
				<section id="report-page-problem-section" class="style-border-top">
	<div class="container content-block">
		<p><i class="fa fa-question-circle"></i> <a id="report-page-problem-toggle">????????????????????????</a></p>
		<div id="report-page-problem">
			<div class="row">
				<div class="col-xs-12 col-sm-8">
					<h3>??????????????????GOV.MO</h3>
					<form class="form form-validate" id="report-page-problem-form" name="reportpageproblem" method="post" novalidate="novalidate">
						<div class="form-group">
							<textarea name="message_text" id="message_text" class="form-control" rows="3" placeholder="" required></textarea>
							<span id="message_text-error" class="help-block"></span>
							<label for="message_text">??????????????? *</label>
						</div>
						<div class="form-group">
							<input type="text" name="message_name" id="message_name" class="form-control" value="">
							<label for="message_name">??????</label>
						</div>
						<div class="form-group">
							<input type="email" name="message_email" id="message_email" class="form-control" value="">
							<label for="message_email">??????</label>
						</div>
						<div class="form-group">
							<input type="tel" name="message_phone" id="message_phone" class="form-control" value="">
							<label for="message_phone">??????</label>
						</div>
						<p>* ?????????</p>
						<div class="form-group">
							
<!-- BEGIN recaptcha, injected by plugin wp-recaptcha-integration  -->
<div  id="g-recaptcha-0" class="wp-recaptcha" data-theme="light" data-size="normal" data-callback=""></div><noscript>Please enable JavaScript to submit this form.<br></noscript>
<!-- END recaptcha -->
						</div>
												<input type="hidden" name="submitted" value="1">
						<div id="report-page-problem-info"></div>
						<a id="report-page-problem-form-submit-btn" data-url="https://www.gov.mo/zh-hant/report-page-problem/?ajax=1&post_id=9"  class="btn btn-primary ink-reaction">?????? <i class="fa fa-chevron-right"></i></a>
					</form>
				</div>
			</div>
		</div>
	</div>
</section>
			</main>
	<div class="push-slim"></div>
</div>
<footer id="footer" class="style-primary">
	<h2 class="sr-only">????????????</h2>
	<h3 class="sr-only">????????????</h3>
	<div class="row">
		<div id="footer-right" class="col-xs-12 col-md-6 col-md-push-6">
			<div class="footer-nav-language"><ul id="menu-languages-1" class="footer-nav"><li class="blog-id-4 mlp-language-nav-item menu-item menu-item-type-language menu-item-object-mlp_language mlp-current-language-item menu-item-341186"><a rel="alternate" href="https://www.gov.mo/zh-hant/">????????????</a></li>
<li class="blog-id-5 mlp-language-nav-item menu-item menu-item-type-language menu-item-object-mlp_language menu-item-341187"><a rel="alternate" href="https://www.gov.mo/zh-hans/">????????????</a></li>
<li class="blog-id-3 mlp-language-nav-item menu-item menu-item-type-language menu-item-object-mlp_language menu-item-341185"><a rel="alternate" href="https://www.gov.mo/pt/">Portugu??s</a></li>
<li class="blog-id-2 mlp-language-nav-item menu-item menu-item-type-language menu-item-object-mlp_language menu-item-341184"><a rel="alternate" href="https://www.gov.mo/en/">English</a></li>
</ul></div>		</div>
		<h3 class="sr-only">????????????</h3>
		<div id="footer-left" class="col-xs-12 col-md-6 col-md-pull-6">
			<div class="footer-nav-links"><ul id="menu-footer-links" class="footer-nav"><li id="menu-item-186553" class="menu-item menu-item-type-post_type menu-item-object-page menu-item-186553"><a href="https://www.gov.mo/zh-hant/terms-of-use/">????????????</a></li>
<li id="menu-item-186558" class="menu-item menu-item-type-post_type menu-item-object-page menu-item-privacy-policy menu-item-186558"><a href="https://www.gov.mo/zh-hant/privacy-statement/">????????????</a></li>
<li id="menu-item-2614" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-2614"><a target="_blank" rel="noopener noreferrer" href="http://www.safp.gov.mo">???????????????????????????????????????????????????</a></li>
</ul></div>		</div>
	</div>
</footer>
<link rel='stylesheet' id='pe-icon-set-weather-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/libs/pe-icon-set-weather/css/pe-icon-set-weather.css' type='text/css' media='all' />
<link rel='stylesheet' id='pe-icon-set-weather-helper-css'  href='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/css/theme-default/libs/pe-icon-set-weather/css/helper.css' type='text/css' media='all' />
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-includes/js/jquery/jquery.js' id='jquery-core-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/owl.carousel/owl.carousel.min.js' id='owl-carousel-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/jquery-match-height/jquery.matchHeight-min.js' id='jquery-match-height-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/bootstrap/bootstrap.min.js' id='bootstrap-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/bootstrap-toolkit/bootstrap-toolkit.min.js' id='bootstrap-toolkit-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/js-cookie/js.cookie.min.js' id='js-cookie-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/jquery-fitvids/jquery.fitvids.min.js' id='jquery-fitvids-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/jquery-dotdotdot/jquery.dotdotdot.min.js' id='jquery-dotdotdot-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/core/govmo/template-home-v3.min.js?20200324' id='template-home-js'></script>
<script type='text/javascript' src='https://ajax.googleapis.com/ajax/libs/webfont/1.5.18/webfont.js' id='webfont-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/core/govmo/webfont-load.min.js' id='webfont-load-js'></script>
<!--[if lt IE 9]>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/utils/respond.min.js' id='respond-js'></script>
<![endif]-->
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/spin.js/spin.min.js' id='spin-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/autosize/jquery.autosize.min.js' id='autosize-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/nanoscroller/jquery.nanoscroller.min.js' id='nanoscroller-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/headroom/headroom.min.js' id='headroom-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/headroom/jquery.headroom.js' id='jquery-headroom-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/textfit/textFit.min.js' id='textfit-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/core/source/App.min.js' id='app-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/core/source/AppNavigation.min.js' id='app-navigation-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/core/source/AppOffcanvas.min.js' id='app-offcanvas-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/core/source/AppCard.min.js' id='app-card-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/core/source/AppForm.min.js' id='app-form-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/core/source/AppNavSearch.min.js' id='app-nav-search-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/core/source/AppVendor.min.js' id='app-vendor-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/qrcode/qrcode.min.js' id='qrcode-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/clipboard.js/clipboard.min.js' id='clipboardjs-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/social-share/social-share.min.js?20180711f' id='social-share-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/social-share/plugins/wechat.min.js' id='social-share-wechat-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/social-share/plugins/facebook.min.js' id='social-share-facebook-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/social-share/plugins/google-plus.min.js' id='social-share-google-plus-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/social-share/plugins/weibo.min.js' id='social-share-weibo-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/social-share/plugins/twitter.min.js' id='social-share-twitter-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/social-share/plugins/link.min.js?20180711f' id='social-share-link-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/core/govmo/cookie-consent.min.js' id='govmo-cookie-consent-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/core/govmo/main-min.js' id='govmo-main-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-includes/js/wp-embed.min.js' id='wp-embed-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/jquery-validation/dist/jquery.validate.min.js' id='jquery-validation-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/jquery-validation/dist/additional-methods.min.js' id='jquery-validation-additional-methods-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/libs/jquery-validation/dist/localization/messages_zh_TW.min.js' id='jquery-validation-localization-js'></script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/themes/gov-mo-2019/assets/js/core/govmo/template-contact-form-report-page-problem.min.js' id='template-contact-form-report-page-problem-js'></script>
<script type='text/javascript' id='wp-recaptcha-js-extra'>
/* <![CDATA[ */
var wp_recaptcha = {"recaptcha_url":"https:\/\/www.google.com\/recaptcha\/api.js?onload=wp_recaptcha_loaded&render=explicit&hl=zh-TW","site_key":"6Ld1iBgUAAAAAIQQLbEcx6g3PB3FRSLMip7Nmbzf"};
/* ]]> */
</script>
<script type='text/javascript' src='https://www.gov.mo/zh-hant/wp-content/plugins/wp-recaptcha-integration/js/wp-recaptcha.min.js' id='wp-recaptcha-js'></script>
</body>
</html>
`)
	fmt.Println(GetCharsetName([]byte(`\r\r<meta http-equiv="content-type" content="text/html;charset=utf-8">\n\n`)))
	html = RemoveAllTagA([]byte(html))
	fmt.Println(string(html))
	fmt.Println("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz:")
	fmt.Println(string(RemoveSpace(HtmlToText(html))))

	fmt.Println(string(HtmlRemoveAllScriptV2([]byte(`dzz<noscript>Your browser does not support JavaScript!</noscript>uuas
	fs
	af
	dsazz<noscript>Your browser does not support JavaScript!</noscript>yy<script type='text/javascript'>
	var s='Dskfdkslfjs\'\"\\dlkjfadfjasklfj</script>dd\'\"\\sfsafd';
	var d="jdslkfjs\'\"\\ldjfj</script>ddsfl\'\"\\kdsfl";
	` + "var e=`dsfjlkdsajflksdjalfj</script>ddskf\\sjalkfsaf`;" + `
	</script>zz
	f
	sa
	f
	sa
	f`))))

	fmt.Println(GetAllTagALinkAndName("http://www.macity.tk/surprisor/", []byte(`
<a href='ed2k://wewrjrjlkd.dfdsaffa/'>ed2k test</a>sdlkf<a href="HTTP://sadfdsalfdsalf.SD/adsFdsaf#dsfkdsf">dddd</a>
`)))

	fmt.Println(ToFullUrl("./", "http://www.macity.tk/surprisor/"))
}

/*
func TestToFullUrl(t *testing.T) {
	fmt.Println(ToFullUrl("/20181226/ef5d4ec3-b61d-76f4-7f75-6132c1955176.html", "http://sc.cRi.cn/highlights.html"))
	fmt.Println(ToFullUrl("../20181226/ef5d4ec3-b61d-76f4-7f75-6132c1955176.html", "http://sc.cRi.cn/high/lights.html"))
	fmt.Println(ToFullUrl("./20181226/ef5d4ec3-b61d-76f4-7f75-6132c1955176.html#ghg435", "http://sc.cRi.cn/high/lights.html"))
	fmt.Println(ToFullUrl("http://sc.cRi.cn/high/20181226/ef5d4ec3-b61d-76f4-7f75-6132c1955176.html#sdf32=2343", "http://sc.cRi.cn/high/lights.html"))
	fmt.Println(ToFullUrl("//p2.cri.cn/M00/B4/44/CqgNOltVLE6AFx2uAAAAAAAAAAA176.600x302.jpg#sdf32=2343", "http://sc.cRi.cn/high/lights.html"))
	fmt.Println(ToFullUrl("http://P2.cri.cn/M00/B4/44/CqgNOltVLE6AFx2uAAAAAAAAAAA176.600x302.jpg#sdf32=2343", "http://sc.cRi.cn/high/lights.html"))
}
func TestGetFullTagV2(t *testing.T) {
	pa := "pagedata.html"
	ctt, _ := ioutil.ReadFile(pa)
	ctt2 := GetFullTagV2(ctt, []byte("<div id=\"abody\" class=\"abody\""))
	ioutil.WriteFile("pagedatafulltag.txt", ctt2, 0666)

	GetFullTagV2([]byte{}, []byte("jldsf"))
}

/*
func TestPageRelink(t *testing.T) {
	pa := "pagedata2.html"
	ctt, _ := ioutil.ReadFile(pa)
	ctt2 := PageRelinkAllToFullUrl("http://sc.cri.cn/20180723/16a4e4c9-899d-b7a2-59a8-47dd37a5ba7e.html", string(ctt), "", "", []string{}, []string{})
	ioutil.WriteFile("pagedata2relink.txt", []byte(ctt2), 0666)
	ioutil.WriteFile("pagedata2relink_std.txt", StdHtmlDoc([]byte(ctt2)), 0666)
}

/*
*/
