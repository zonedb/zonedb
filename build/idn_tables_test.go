package build

import (
	"strings"
	"testing"
)

func TestProcessBasic(t *testing.T) {
	// Sampled from https://www.iana.org/domains/idn-tables/tables/com_latn_1.2.txt
	data := `###############################################################
#
# TLD:            COM
# Script:         Latin
# Version Number: 1.2
# Effective Date: October 25th, 2014
# Registry:       Verisign, Inc.
# Address:        12061 Bluemont Way, Reston VA 20190, USA
# Telephone:      +1 (703) 925-6999
# Email:          info@verisign-grs.com
# URL:            http://www.verisigninc.com
#
###############################################################


###############################################################
#
# Codepoints allowed from the Latin script.
#
###############################################################
U+0061	# LATIN SMALL LETTER A
U+0062	# LATIN SMALL LETTER B
U+0063	# LATIN SMALL LETTER C
U+0064	# LATIN SMALL LETTER D
U+0065	# LATIN SMALL LETTER E
U+0066	# LATIN SMALL LETTER F
U+0067	# LATIN SMALL LETTER G
U+0068	# LATIN SMALL LETTER H
U+0069	# LATIN SMALL LETTER I
U+006A	# LATIN SMALL LETTER J
U+006B	# LATIN SMALL LETTER K
U+006C	# LATIN SMALL LETTER L
U+006D	# LATIN SMALL LETTER M
U+006E	# LATIN SMALL LETTER N
U+006F	# LATIN SMALL LETTER O
U+0070	# LATIN SMALL LETTER P
U+0071	# LATIN SMALL LETTER Q
U+0072	# LATIN SMALL LETTER R
U+0073	# LATIN SMALL LETTER S
U+0074	# LATIN SMALL LETTER T
U+0075	# LATIN SMALL LETTER U
U+0076	# LATIN SMALL LETTER V
U+0077	# LATIN SMALL LETTER W
U+0078	# LATIN SMALL LETTER X
U+0079	# LATIN SMALL LETTER Y
U+007A	# LATIN SMALL LETTER Z
U+00DF	# LATIN SMALL LETTER SHARP S
U+00E0	# LATIN SMALL LETTER A GRAVE
U+00E1	# LATIN SMALL LETTER A ACUTE
U+00E2	# LATIN SMALL LETTER A CIRCUMFLEX
U+00E3	# LATIN SMALL LETTER A TILDE
U+00E4	# LATIN SMALL LETTER A DIAERESIS
U+00E5	# LATIN SMALL LETTER A RING
U+00E6	# LATIN SMALL LETTER A E

U+014D	# LATIN SMALL LETTER O MACRON


###############################################################
#
# Common codepoints
#
###############################################################
U+002D	# HYPHEN-MINUS
U+0030	# DIGIT ZERO
U+0031	# DIGIT ONE
U+0032	# DIGIT TWO
U+0033	# DIGIT THREE
U+0034	# DIGIT FOUR
U+0035	# DIGIT FIVE
U+0036	# DIGIT SIX
U+0037	# DIGIT SEVEN
U+0038	# DIGIT EIGHT
U+0039	# DIGIT NINE
`
	expected := `--09azßæōō`

	CheckTable(t, data, expected, false)
}

func TestProcessNoPrefix(t *testing.T) {
	// Sampled from https://www.iana.org/domains/idn-tables/tables/asia_ko_1.0.txt
	data := `.ASIA KO IDN Language Table

TLD: ASIA
Version: 1.0
Effective Date: 2011-03-11

Summary of IDN Policy Profile:
IDN LangaugeTag: KO
IDN Language Description: Korean
Minimum Length: A-Label: 3; U-Label: 1
Maximum Length: A-Label: 63
Valid Characters: 11,209
Additional Contextual Rules: A Domain Name Applied For must include at least one non-LDH character
IDN Variants: N/A

Registry: DotAsia Organisation
Contact: DotAsia Admin Contact <admin@iana.whois.asia>
Contact: DotAsia Tech Contact <tech@iana.whois.asia>
Address: 15/F, 6 Knutsford Terrace, Tsim Sha Tsui, Hong Kong
TEL: +852.35202635
FAX: +852.35202634
Website: http://www.registry.asia

Relevant Policy Document URLs:
- .ASIA IDN Policies for CJK (Chinese, Japanese & Korean): http://dot.asia/policies/DotAsia-CJK-IDN-
Policies-COMPLETE--2011-03-11.pdf
- .ASIA IDN Sunrise Policies: http://dot.asia/policies/DotAsia-IDN-Sunrise-Policies-COMPLETE--2011-
03-11.pdf


Abstract

This language table is adopted from the KRNIC Korean IDN Language table (ko-KR:
http://www.iana.org/domains/idn-tables/tables/kr_ko-kr_1.0.html) for the implementation of
Korean IDN registrations at the .ASIA gTLD.

DotAsia is committed to maintaining the integrity of .ASIA IDN registrations and understands the
multilingual nature of the .ASIA zone as a gTLD.  In preparation of this language table, DotAsia has
coordinated with KISA (Korea Internet & Security Agency) on the applicability of the KRNIC Korean IDN
Language table for the .ASIA gTLD.

The following table includes 11,172 Hangul Syllables (along with the 36 LDH characters, “a-z”, “0-9” and
“-“) according to the consensus of Korean Internet community.  Hanja characters are not included in
the table.  Mixture of Hanja characters with Hangul characters for domain names is not encouraged.
Registrants interested in registering domain names with Hanja characters only may consider utilizing a
registration based on the ZH table.

The following table uses the RFC3743-defined format, and is in compliance with the ICANN Guidelines
for IDN registration and for publication in the IANA Repository of IDN Practices.

References:

Please refer to the references cited for the ko-KR KRNIC Korean IDN Language table.

Version 1.0 20110311  # Mar. 11, 2011 Version 1.0

002D;002D;
0030;0030;
0031;0031;
0032;0032;
0033;0033;
0034;0034;
0035;0035;
0036;0036;
0037;0037;
0038;0038;
0039;0039;
0061;0061;
0062;0062;
0063;0063;
0064;0064;
0065;0065;
0066;0066;
0067;0067;
0068;0068;
0069;0069;
006A;006A;
006B;006B;
006C;006C;
006D;006D;
006E;006E;
006F;006F;
0070;0070;
0071;0071;
0072;0072;
0073;0073;
0074;0074;
0075;0075;
0076;0076;
0077;0077;
0078;0078;
0079;0079;
007A;007A;
AC00;AC00;
AC01;AC01;
AC02;AC02;
AC03;AC03;
AC04;AC04;
AC05;AC05;
AC06;AC06;
AC07;AC07;
AC08;AC08;
AC09;AC09;
...
AC10;AC10;
AC11;AC11;
AC12;AC12;
AC13;AC13;
AC14;AC14;
AC15;AC15;
AC16;AC16;
AC17;AC17;
AC18;AC18;
AC19;AC19;
AC1A;AC1A;
AC1B;AC1B;
AC1C;AC1C;
AC1D;AC1D;
AC1E;AC1E;
AC1F;AC1F;
`
	expected := `--09az가갉감갟`

	CheckTable(t, data, expected, false)
}

func TestProcessHTMLMultiplePreTags(t *testing.T) {
	// From https://www.iana.org/domains/idn-tables/tables/biz_pt_1.0.html
	data := `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=iso-8859-1">
<title>NeuLevel Portuguese Language Table</title>
</head>
<body bgcolor="#FFFFFF">

<table>
<tr>
<td>
TLD: BIZ<br>
Language Tag: pt<br>
Language Description: Portuguese<br>
Version: 1.0<br>
Effective Date: Jan. 16, 2008<br>
</td>
<td width="30">&nbsp;</td>
<td width="30">&nbsp;</td>
<td>
Registry: NeuLevel<br>
Contact: William Tan <a href="mailto:william.tan@neustar.biz">william.tan@neustar.biz</a><br>
Address: 46000 Center Oak Plaza, Sterling VA 20166 USA<br>
Telephone: +1-571-434-5400<br>
Website: <a href="http://www.neulevel.biz/">http://www.neulevel.biz/</a><br>
</td>
</tr>
</table>

<p>
Relevant Policy Document URL:
<a href="http://www.neulevel.biz/idn/">
http://www.neulevel.biz/idn/
</a>
</p>

<p>
This character repertoire was derived from the following sources:<br>

<pre>
- http://www.iana.org/assignments/idn/br-portuguese.html
- Davies, K. and Schneider, M., "Suggested IDN Code Points For European
  Languages", Version 0.5, 19 November 2003
    https://www.centr.org/docs/2003/11/centr-ga20-idncodepoints.pdf
</pre>

</p>


<p>
This document describes the Internationalized Domain Names (IDN) Character Table
to be used by NeuLevel in the .BIZ TLD registry for the registration of .BIZ
domains in the Portuguese language.
</p>

<p>
We advise registrants to carefully consider the potential confusion over the
perception of equivalence between an accented Latin character and another
sequence of Latin characters (with or without accents.)

There exists no equivalence rule that can be applied universally as any
such mappings are generally dependent on the language, culture and context in
which an IDN is used.
</p>

<p>
This table is in compliance with the ICANN Guidelines for the Implementation
of Internationalized Domain Names and is intended for publication in the IANA
IDN Character Table Registry.
</p>



<br />

<pre>
U+002D	# HYPHEN-MINUS (&#45;)
U+0030	# DIGIT ZERO (&#48;)
U+0031	# DIGIT ONE (&#49;)
U+0032	# DIGIT TWO (&#50;)
U+0033	# DIGIT THREE (&#51;)
U+0034	# DIGIT FOUR (&#52;)
U+0035	# DIGIT FIVE (&#53;)
U+0036	# DIGIT SIX (&#54;)
U+0037	# DIGIT SEVEN (&#55;)
U+0038	# DIGIT EIGHT (&#56;)
U+0039	# DIGIT NINE (&#57;)
U+0061	# LATIN SMALL LETTER A (&#97;)
U+0062	# LATIN SMALL LETTER B (&#98;)
U+0063	# LATIN SMALL LETTER C (&#99;)
U+0064	# LATIN SMALL LETTER D (&#100;)
U+0065	# LATIN SMALL LETTER E (&#101;)
U+0066	# LATIN SMALL LETTER F (&#102;)
U+0067	# LATIN SMALL LETTER G (&#103;)
U+0068	# LATIN SMALL LETTER H (&#104;)
U+0069	# LATIN SMALL LETTER I (&#105;)
U+006A	# LATIN SMALL LETTER J (&#106;)
U+006B	# LATIN SMALL LETTER K (&#107;)
U+006C	# LATIN SMALL LETTER L (&#108;)
U+006D	# LATIN SMALL LETTER M (&#109;)
U+006E	# LATIN SMALL LETTER N (&#110;)
U+006F	# LATIN SMALL LETTER O (&#111;)
U+0070	# LATIN SMALL LETTER P (&#112;)
U+0071	# LATIN SMALL LETTER Q (&#113;)
U+0072	# LATIN SMALL LETTER R (&#114;)
U+0073	# LATIN SMALL LETTER S (&#115;)
U+0074	# LATIN SMALL LETTER T (&#116;)
U+0075	# LATIN SMALL LETTER U (&#117;)
U+0076	# LATIN SMALL LETTER V (&#118;)
U+0077	# LATIN SMALL LETTER W (&#119;)
U+0078	# LATIN SMALL LETTER X (&#120;)
U+0079	# LATIN SMALL LETTER Y (&#121;)
U+007A	# LATIN SMALL LETTER Z (&#122;)
U+00E0	# LATIN SMALL LETTER A WITH GRAVE (&#224;)
U+00E1	# LATIN SMALL LETTER A WITH ACUTE (&#225;)
U+00E2	# LATIN SMALL LETTER A WITH CIRCUMFLEX (&#226;)
U+00E3	# LATIN SMALL LETTER A WITH TILDE (&#227;)
U+00E7	# LATIN SMALL LETTER C WITH CEDILLA (&#231;)
U+00E9	# LATIN SMALL LETTER E WITH ACUTE (&#233;)
U+00EA	# LATIN SMALL LETTER E WITH CIRCUMFLEX (&#234;)
U+00ED	# LATIN SMALL LETTER I WITH ACUTE (&#237;)
U+00F3	# LATIN SMALL LETTER O WITH ACUTE (&#243;)
U+00F4	# LATIN SMALL LETTER O WITH CIRCUMFLEX (&#244;)
U+00F5	# LATIN SMALL LETTER O WITH TILDE (&#245;)
U+00FA	# LATIN SMALL LETTER U WITH ACUTE (&#250;)
U+00FC	# LATIN SMALL LETTER U WITH DIAERESIS (&#252;)

</pre>


</body>
</html>
`
	expected := `--09azàãççéêííóõúúüü`

	CheckTable(t, data, expected, true)
}

func TestProcessHTMLIndented(t *testing.T) {
	// From https://www.iana.org/domains/idn-tables/tables/eu_latn_1.0.html
	data := `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
<title>.EU Latin IDN Table</title>
</head>
<body bgcolor="#FFFFFF">
<table>
<tr>
<td>
TLD: EU<br/>
Script Identifier: Latin<br/>
Script Description: Latin Unicode (Basic, Extended-A and Extended-B)<br/>

Version: 1.0<br/>
Effective Date: 10 December 2009<br/>
</td>
<td width="10%">&nbsp;</td>
<td>
Registry: EURid vzw/asbl<br/>
Contact: Marc Van Wesemael <a href="mailto:admin@eurid.eu">admin@eurid.eu</a><br/>
Address: Woluwelaan 150, 1831 Diegem, Belgium<br/>
Telephone: +32-2-4012750<br/>Fax: +32-2-4012751<br/>

Website: <a href="http://www.eurid.eu/">http://www.eurid.eu/</a><br>
</td>
</tr>
</table>
<p>
Relevant Policy Document URL: <a href="http://link.eurid.eu/trm-con">
http://link.eurid.eu/trm-con</a>
</p>
<p>
This document describes the Latin Internationalized Domain Names (IDN) Character Table to be used by
EURid vzw/asbl, the .eu Registry for registration of .eu Internationalized Domain Names in Latin Script.
</p>
<p>
The hyphen (U+002D), as well as the digits zero to nine (U+0030 .. U+0039), for purposes of .eu
Internationalised Domain Name registration are considered to form part of the Latin script.
</p>


<p>
This table is in compliance with the ICANN Guidelines for the Implementation of Internationalized
Domain Names Version 2.2, and is intended for publication in the IANA IDN Character Table Registry.
</p>

</p>
<pre>
Code Point      Character

######
  U+002D  ##  HYPHEN-MINUS
######
  U+0030  ##  DIGIT ZERO
  U+0031  ##  DIGIT ONE
  U+0032  ##  DIGIT TWO
  U+0033  ##  DIGIT THREE
  U+0034  ##  DIGIT FOUR
  U+0035  ##  DIGIT FIVE
  U+0036  ##  DIGIT SIX
  U+0037  ##  DIGIT SEVEN
  U+0038  ##  DIGIT EIGHT
  U+0039  ##  DIGIT NINE
######
###### Basic Latin
  U+0061  ##  a LATIN SMALL LETTER A
  U+0062  ##  b LATIN SMALL LETTER B
  U+0063  ##  c LATIN SMALL LETTER C
  U+0064  ##  d LATIN SMALL LETTER D
  U+0065  ##  e LATIN SMALL LETTER E
  U+0066  ##  f LATIN SMALL LETTER F
  U+0067  ##  g LATIN SMALL LETTER G
  U+0068  ##  h LATIN SMALL LETTER H
  U+0069  ##  i LATIN SMALL LETTER I
  U+006A  ##  j LATIN SMALL LETTER J
  U+006B  ##  k LATIN SMALL LETTER K
  U+006C  ##  l LATIN SMALL LETTER L
  U+006D  ##  m LATIN SMALL LETTER M
  U+006E  ##  n LATIN SMALL LETTER N
  U+006F  ##  o LATIN SMALL LETTER O
  U+0070  ##  p LATIN SMALL LETTER P
  U+0071  ##  q LATIN SMALL LETTER Q
  U+0072  ##  r LATIN SMALL LETTER R
  U+0073  ##  s LATIN SMALL LETTER S
  U+0074  ##  t LATIN SMALL LETTER T
  U+0075  ##  u LATIN SMALL LETTER U
  U+0076  ##  v LATIN SMALL LETTER V
  U+0077  ##  w LATIN SMALL LETTER W
  U+0078  ##  x LATIN SMALL LETTER X
  U+0079  ##  y LATIN SMALL LETTER Y
  U+007A  ##  z LATIN SMALL LETTER Z
###### Latin-1 supplement
  U+00E0  ##  à LATIN SMALL LETTER A WITH GRAVE
  U+00E1  ##  á LATIN SMALL LETTER A WITH ACUTE
  U+00E2  ##  â LATIN SMALL LETTER A WITH CIRCUMFLEX
  U+00E3  ##  ã LATIN SMALL LETTER A WITH TILDE
  U+00E4  ##  ä LATIN SMALL LETTER A WITH DIAERESIS
  U+00E5  ##  å LATIN SMALL LETTER A WITH RING ABOVE
  U+00E6  ##  æ LATIN SMALL LETTER AE
  U+00E7  ##  ç LATIN SMALL LETTER C WITH CEDILLA
  U+00E8  ##  è LATIN SMALL LETTER E WITH GRAVE
  U+00E9  ##  é LATIN SMALL LETTER E WITH ACUTE
  U+00EA  ##  ê LATIN SMALL LETTER E WITH CIRCUMFLEX
  U+00EB  ##  ë LATIN SMALL LETTER E WITH DIAERESIS
  U+00EC  ##  ì LATIN SMALL LETTER I WITH GRAVE
  U+00ED  ##  í LATIN SMALL LETTER I WITH ACUTE
  U+00EE  ##  î LATIN SMALL LETTER I WITH CIRCUMFLEX
  U+00EF  ##  ï LATIN SMALL LETTER I WITH DIAERESIS
  U+00F0  ##  ð LATIN SMALL LETTER ETH
  U+00F1  ##  ñ LATIN SMALL LETTER N WITH TILDE
  U+00F2  ##  ò LATIN SMALL LETTER O WITH GRAVE
  U+00F3  ##  ó LATIN SMALL LETTER O WITH ACUTE
  U+00F4  ##  ô LATIN SMALL LETTER O WITH CIRCUMFLEX
  U+00F5  ##  õ LATIN SMALL LETTER O WITH TILDE
  U+00F6  ##  ö LATIN SMALL LETTER O WITH DIAERESIS
  U+00F8  ##  ø LATIN SMALL LETTER O WITH STROKE
  U+00F9  ##  ù LATIN SMALL LETTER U WITH GRAVE
  U+00FA  ##  ú LATIN SMALL LETTER U WITH ACUTE
  U+00FB  ##  û LATIN SMALL LETTER U WITH CIRCUMFLEX
  U+00FC  ##  ü LATIN SMALL LETTER U WITH DIAERESIS
  U+00FD  ##  ý LATIN SMALL LETTER Y WITH ACUTE
  U+00FE  ##  þ LATIN SMALL LETTER THORN
  U+00FF  ##  ÿ LATIN SMALL LETTER Y WITH DIAERESIS
###### Latin Extended-A
  U+0101  ##  ā LATIN SMALL LETTER A WITH MACRON
  U+0103  ##  ă LATIN SMALL LETTER A WITH BREVE
  U+0105  ##  ą LATIN SMALL LETTER A WITH OGONEK
  U+0107  ##  ć LATIN SMALL LETTER C WITH ACUTE
  U+0109  ##  ĉ LATIN SMALL LETTER C WITH CIRCUMFLEX
  U+010B  ##  ċ LATIN SMALL LETTER C WITH DOT ABOVE
  U+010D  ##  č LATIN SMALL LETTER C WITH CARON
  U+010F  ##  ď LATIN SMALL LETTER D WITH CARON
  U+0111  ##  đ LATIN SMALL LETTER D WITH STROKE
  U+0113  ##  ē LATIN SMALL LETTER E WITH MACRON
  U+0115  ##  ĕ LATIN SMALL LETTER E WITH BREVE
  U+0117  ##  ė LATIN SMALL LETTER E WITH DOT ABOVE
  U+0119  ##  ę LATIN SMALL LETTER E WITH OGONEK
  U+011B  ##  ě LATIN SMALL LETTER E WITH CARON
  U+011D  ##  ĝ LATIN SMALL LETTER G WITH CIRCUMFLEX
  U+011F  ##  ğ LATIN SMALL LETTER G WITH BREVE
  U+0121  ##  ġ LATIN SMALL LETTER G WITH DOT ABOVE
  U+0123  ##  ģ LATIN SMALL LETTER G WITH CEDILLA
  U+0125  ##  ĥ LATIN SMALL LETTER H WITH CIRCUMFLEX
  U+0127  ##  ħ LATIN SMALL LETTER H WITH STROKE
  U+0129  ##  ĩ LATIN SMALL LETTER I WITH TILDE
  U+012B  ##  ī LATIN SMALL LETTER I WITH MACRON
  U+012D  ##  ĭ LATIN SMALL LETTER I WITH BREVE
  U+012F  ##  į LATIN SMALL LETTER I WITH OGONEK
  U+0131  ##  ı LATIN SMALL LETTER DOTLESS I
  U+0135  ##  ĵ LATIN SMALL LETTER J WITH CIRCUMFLEX
  U+0137  ##  ķ LATIN SMALL LETTER K WITH CEDILLA
  U+013A  ##  ĺ LATIN SMALL LETTER L WITH ACUTE
  U+013C  ##  ļ LATIN SMALL LETTER L WITH CEDILLA
  U+013E  ##  ľ LATIN SMALL LETTER L WITH CARON
  U+0140  ##  ŀ LATIN SMALL LETTER L WITH MIDDLE DOT
  U+0142  ##  ł LATIN SMALL LETTER L WITH STROKE
  U+0144  ##  ń LATIN SMALL LETTER N WITH ACUTE
  U+0146  ##  ņ LATIN SMALL LETTER N WITH CEDILLA
  U+0148  ##  ň LATIN SMALL LETTER N WITH CARON
  U+0149  ##  ŉ LATIN SMALL LETTER N PRECEDED BY APOSTROPHE
  U+014B  ##  ŋ LATIN SMALL LETTER ENG
  U+014D  ##  ō LATIN SMALL LETTER O WITH MACRON
  U+014F  ##  ŏ LATIN SMALL LETTER O WITH BREVE
  U+0151  ##  ő LATIN SMALL LETTER O WITH DOUBLE ACUTE
  U+0153  ##  œ LATIN SMALL LIGATURE OE
  U+0155  ##  ŕ LATIN SMALL LETTER R WITH ACUTE
  U+0157  ##  ŗ LATIN SMALL LETTER R WITH CEDILLA
  U+0159  ##  ř LATIN SMALL LETTER R WITH CARON
  U+015B  ##  ś LATIN SMALL LETTER S WITH ACUTE
  U+015D  ##  ŝ LATIN SMALL LETTER S WITH CIRCUMFLEX
  U+0161  ##  š LATIN SMALL LETTER S WITH CARON
  U+0165  ##  ť LATIN SMALL LETTER T WITH CARON
  U+0167  ##  ŧ LATIN SMALL LETTER T WITH STROKE
  U+0169  ##  ũ LATIN SMALL LETTER U WITH TILDE
  U+016B  ##  ū LATIN SMALL LETTER U WITH MACRON
  U+016D  ##  ŭ LATIN SMALL LETTER U WITH BREVE
  U+016F  ##  ů LATIN SMALL LETTER U WITH RING ABOVE
  U+0171  ##  ű LATIN SMALL LETTER U WITH DOUBLE ACUTE
  U+0173  ##  ų LATIN SMALL LETTER U WITH OGONEK
  U+0175  ##  ŵ LATIN SMALL LETTER W WITH CIRCUMFLEX
  U+0177  ##  ŷ LATIN SMALL LETTER Y WITH CIRCUMFLEX
  U+017A  ##  ź LATIN SMALL LETTER Z WITH ACUTE
  U+017C  ##  ż LATIN SMALL LETTER Z WITH DOT ABOVE
  U+017E  ##  ž LATIN SMALL LETTER Z WITH CARON
###### Latin Extended-B
  U+0219  ##  ș LATIN SMALL LETTER S WITH COMMA BELOW
  U+021B  ##  ț LATIN SMALL LETTER T WITH COMMA BELOW
######

</pre>

</body>
</html>
`
	expected := `--09azàöøÿāāăăąąććĉĉċċččďďđđēēĕĕėėęęěěĝĝğğġġģģĥĥħħĩĩīīĭĭįįııĵĵķķĺĺļļľľŀŀłłńńņņňŉŋŋōōŏŏőőœœŕŕŗŗřřśśŝŝššťťŧŧũũūūŭŭůůűűųųŵŵŷŷźźżżžžșșțț`

	CheckTable(t, data, expected, true)
}

func TestProcessFiveCharHex(t *testing.T) {
	// Sampled from https://www.iana.org/domains/idn-tables/tables/name_cari_1.1.txt
	data := `###############################################################
#
# TLD:            NAME
# Script:         Carian
# Version Number: 1.1
# Effective Date: July 1st, 2011
# Registry:       Verisign, Inc.
# Address:        12061 Bluemont Way, Reston VA 20190, USA
# Telephone:      +1 (703) 925-6999
# Email:          info@verisign-grs.com
# URL:            http://www.verisigninc.com
#
###############################################################


###############################################################
#
# Codepoints allowed from the Carian script.
#
###############################################################
U+102A0	# CARIAN LETTER A
U+102A1	# CARIAN LETTER P2
U+102A2	# CARIAN LETTER D
U+102A3	# CARIAN LETTER L
U+102A4	# CARIAN LETTER UUU
U+102A5	# CARIAN LETTER R
U+102A6	# CARIAN LETTER LD
U+102A7	# CARIAN LETTER A2
U+102A8	# CARIAN LETTER Q
U+102A9	# CARIAN LETTER B
U+102AA	# CARIAN LETTER M
U+102AB	# CARIAN LETTER O
U+102AC	# CARIAN LETTER D2
U+102AD	# CARIAN LETTER T
U+102AE	# CARIAN LETTER SH
U+102AF	# CARIAN LETTER SH2
U+102B0	# CARIAN LETTER S
U+102B1	# CARIAN LETTER C-18
U+102B2	# CARIAN LETTER U
U+102B3	# CARIAN LETTER NN
U+102B4	# CARIAN LETTER X
U+102B5	# CARIAN LETTER N
U+102B6	# CARIAN LETTER TT2
U+102B7	# CARIAN LETTER P
U+102B8	# CARIAN LETTER SS
U+102B9	# CARIAN LETTER I
U+102BA	# CARIAN LETTER E
U+102BB	# CARIAN LETTER UUUU
U+102BC	# CARIAN LETTER K
U+102BD	# CARIAN LETTER K2
U+102BE	# CARIAN LETTER ND
U+102BF	# CARIAN LETTER UU
U+102C0	# CARIAN LETTER G
U+102C1	# CARIAN LETTER G2
U+102C2	# CARIAN LETTER ST
U+102C3	# CARIAN LETTER ST2
U+102C4	# CARIAN LETTER NG
U+102C5	# CARIAN LETTER II
U+102C6	# CARIAN LETTER C-39
U+102C7	# CARIAN LETTER TT
U+102C8	# CARIAN LETTER UUU2
U+102C9	# CARIAN LETTER RR
U+102CA	# CARIAN LETTER MB
U+102CB	# CARIAN LETTER MB2
U+102CC	# CARIAN LETTER MB3
U+102CD	# CARIAN LETTER MB4
U+102CE	# CARIAN LETTER LD2
U+102CF	# CARIAN LETTER E2
U+102D0	# CARIAN LETTER UUU3


###############################################################
#
# Common codepoints (allowed in all scripts).
#
###############################################################
U+002D	# HYPHEN-MINUS
U+0030	# DIGIT ZERO
U+0031	# DIGIT ONE
U+0032	# DIGIT TWO
U+0033	# DIGIT THREE
U+0034	# DIGIT FOUR
U+0035	# DIGIT FIVE
U+0036	# DIGIT SIX
U+0037	# DIGIT SEVEN
U+0038	# DIGIT EIGHT
U+0039	# DIGIT NINE
U+00B7	# MIDDLE DOT

U+2E2F	# VERTICAL TILDE
`
	expected := `--09··ⸯⸯ𐊠𐋐`

	CheckTable(t, data, expected, false)
}

func TestProcessHTMLRFC3743(t *testing.T) {
	// Sampled from https://www.iana.org/domains/idn-tables/tables/tw_zh-tw_4.0.1.html
	data := `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html><head>
<meta http-equiv="Content-Type" content="text/html; charset=iso-8859-1">
<title>.tw Traditional Chinese Character Table</title>
</head>
<body bgcolor="#FFFFFF">
<table>
<tr>
    <td width="195"> TLD: TW<br>
      Language Tag: ZH-TW<br>
      Language Description: Traditional Chinese<br>
Version: 4.0.1<br>
      Effective Date: 31 March 2005<br>
</td>
<td width="30">&nbsp;</td>
<td width="33">&nbsp;</td>
    <td width="401"> Registry: Taiwan Network Information Center(TWNIC)<br>
      Contact: Ai-Chin Lu <a href="mailto:tw-admin@twnic.net">tw-admin@twnic.net</a><br>
      Address: 4F-2, No. 9, Sec. 2, Roosevelt Rd., Taipei 100, Taiwan, R.O.C.<br>
      +886-2-2341-1313 #126 FAX: +886-2-2396-8832<br>
Website: <a href="http://www.twnic.net.tw">http://www.twnic.net.tw</a><br>
</td>
</tr>
</table>
<p> Relevant Policy Document URL: <a href="/go/rfc4713">RFC 4713</a></p>
<p> The ZH-TW Language Table defines all the 19520 Chinese characters in
Chinese character variant table for traditional Chinese domain name
registration, and the registration guideline for Chinese domain name
has been expressed as an IETF draft: Registration and Administration
Guideline for Chinese Domain Names.
<p> According to RFC3743, only the list of Valid Code Point in the ZH-TW
Language Table is permitted for registration. If some Chinese character
not included in the ZH-TW Language Table is required to be allowed as a
new Valid Code Point, the ZH-TW Language Table will be modified and the
new version of ZH-TW Language Table must be kept backward compatible
with the old one.
<p> Each registry that is going to support Chinese Domain Name registration
should follow the documents--RFC3743, IETF draft: Registration and
Administration Guideline for Chinese Domain Names and ICANN's Guidelines
for the Implementation of Internationalized Domain Names to implement
the language table.
<p> Acknowledgements
<p>
Thanks for Li-Ming Tseng (Mail address: tsenglm&cc.ncu.edu.tw) who
leading TWNIC Chinese Domain Name Variant Table Working Group to
accomplish the language table.
<p>
<pre>
Reference 0 Unicode 3.2
Reference 1 A Complete Set of Simplified Chinese Characters
Reference 2 Chinese Variants Collation Table
Reference 3 Chinese Big Dictionary
Reference 4 Chinese Relationship Table for Unihan Project
Reference 5 GB2312
Reference 6 General Table for Modern Chinese
Reference 7 International Chinese Standard Big Dictionary
Reference 8 Unihan Database
Reference 9 BIG5

U+002D(0);U+002D(0);
U+0030(0);U+0030(0);
U+0031(0);U+0031(0);
U+0032(0);U+0032(0);
U+0033(0);U+0033(0);
U+0034(0);U+0034(0);
U+0035(0);U+0035(0);
U+0036(0);U+0036(0);
U+0037(0);U+0037(0);
U+0038(0);U+0038(0);
U+0039(0);U+0039(0);
U+0061(0);U+0061(0);
U+0062(0);U+0062(0);
U+0063(0);U+0063(0);
U+0064(0);U+0064(0);
U+0065(0);U+0065(0);
U+0066(0);U+0066(0);
U+0067(0);U+0067(0);
U+0068(0);U+0068(0);
U+0069(0);U+0069(0);
U+006A(0);U+006A(0);
U+006B(0);U+006B(0);
U+006C(0);U+006C(0);
U+006D(0);U+006D(0);
U+006E(0);U+006E(0);
U+006F(0);U+006F(0);
U+0070(0);U+0070(0);
U+0071(0);U+0071(0);
U+0072(0);U+0072(0);
U+0073(0);U+0073(0);
U+0074(0);U+0074(0);
U+0075(0);U+0075(0);
U+0076(0);U+0076(0);
U+0077(0);U+0077(0);
U+0078(0);U+0078(0);
U+0079(0);U+0079(0);
U+007A(0);U+007A(0);
U+3447(0);U+3473(1,3);U+3473(1,3)
U+3473(0);U+3473(0);U+3447(1,3)
U+359E(0);U+558E(1,3,9);U+558E(1,3,9)
U+360E(0);U+361A(1,3);U+361A(1,3)
U+361A(0);U+361A(0);U+360E(1,3)
U+3918(0);U+396E(1);U+396E(1)
U+396E(0);U+396E(0);U+3918(1)
U+39CF(0);U+6386(1,3);U+6386(1,3)
U+39D0(0);U+3A73(1,3);U+3A73(1,3)
U+39DF(0);U+64D3(1,3);U+64D3(1,3)
U+3A73(0);U+3A73(0);U+39D0(1,3)
U+3B4E(0);U+68E1(1,3,9);U+68E1(1,3,9)
U+3C6E(0);U+6BA8(1,3);U+6BA8(1,3)
U+3CE0(0);U+6FBE(1,3);U+6FBE(1,3)
U+4056(0);U+779C(1,3,9);U+779C(1,3,9)
U+415F(0);U+7A47(1,3,9);U+7A47(1,3,9)
U+4337(0);U+7D2C(1,3,9);U+7D2C(1,3,9),U+7EF8(1,3,4),U+7DA2(1,3,4,8,9)
U+43AC(0);U+43B1(1,3);U+43B1(1,3)
U+43B1(0);U+43B1(0);U+43AC(1,3)
U+43DD(0);U+819E(1,3,9);U+819E(1,3,9)
U+44D6(0);U+85ED(1,3,9);U+85ED(1,3,9)
U+464C(0);U+4661(1,3);U+4661(1,3)
U+4661(0);U+4661(0);U+464C(1,3)
U+4723(0);U+8A22(1,3,9);U+8A22(1,3,9)
U+4729(0);U+8B8C(1,3,9);U+8B8C(1,3,9)
U+477C(0);U+477C(0);U+478D(1)
U+478D(0);U+477C(1,3);U+477C(1,3)
U+4947(0);U+4947(0);U+4982(1,3)
U+497A(0);U+91FE(1,3);U+91FE(1,3)
U+497D(0);U+93FA(1,3,9);U+93FA(1,3,9)
U+4982(0);U+4947(1,3);U+4947(1,3)
U+4983(0);U+942F(1,3);U+942F(1,3)
U+4985(0);U+9425(1,3);U+9425(1,3)
U+4986(0);U+9481(1,3,9);U+9481(1,3,9)
U+499B(0);U+499B(0);U+49B6(1,3)
U+499F(0);U+499F(0);U+49B7(1,3)
U+49B6(0);U+499B(1,3);U+499B(1,3)
U+49B7(0);U+499F(1,3);U+499F(1,3)
U+4C77(0);U+4C77(0);U+4CA3(1,3)
U+4C9F(0);U+9BA3(1,3);U+9BA3(1,3)
U+4CA0(0);U+9C06(1,3,9);U+9C06(1,3,9)
U+4CA1(0);U+9C0C(1,3);U+9C0C(1,3),U+9C0D(1,3,9),U+9CC5(1,3,8)
U+4CA2(0);U+9C27(1,3);U+9C27(1,3)
U+4CA3(0);U+4C77(1,3);U+4C77(1,3)
U+4D13(0);U+9CFE(1,3);U+9CFE(1,3)
U+4D14(0);U+9D41(1,3,9);U+9D41(1,3,9)
U+4D15(0);U+9D37(1,3,9);U+9D37(1,3,9)
U+4D16(0);U+9D84(1,3,9);U+9D84(1,3,9)
U+4D17(0);U+9DAA(1,3,9);U+9DAA(1,3,9)
U+4D18(0);U+9DC9(1,3);U+9DC8(3,9),U+9DC9(1,3)
U+4D19(0);U+9E0A(1,3);U+9DFF(3,9),U+9E0A(1,3)
U+4DAE(0);U+9F91(1,3,9);U+9F91(1,3,9)
U+4E00(0);U+4E00(5,9);U+58F9(3,8,9),U+5F0C(3),U+58F1(3,4,8)
U+4E01(0);U+4E01(5,9);
U+4E02(0);U+4E02(0);
U+4E03(0);U+4E03(5,9);
U+4E04(0);U+4E0A(3,8,9);U+4E0A(3,8,9),U+4EE9(3,4,8,9)
U+4E05(0);U+4E0B(3,8,9);U+4E0B(3,8,9)
U+4E06(0);U+4E06(0);
U+4E07(0);U+842C(1,3,8,9);U+842C(1,3,8,9)
U+4E08(0);U+4E08(5,9);
U+4E09(0);U+4E09(5,9);
U+4E0A(0);U+4E0A(5,9);U+4E04(3,8),U+4EE9(4,9)
U+4E0B(0);U+4E0B(5,9);U+4E05(3,8)
U+4E0C(0);U+4E0C(5,9);U+4E93(4,9),U+5176(3,8,9)
U+4E0D(0);U+4E0D(5,9);
U+4E0E(0);U+8207(1,3,8,9);U+8207(1,3,8,9)
U+4E0F(0);U+4E0F(0,9);U+4E10(4,9)
U+4E10(0);U+4E10(5,9);U+4E0F(4,9)
U+4E11(0);U+4E11(1,3,9),U+919C(1,3,4,8,9);U+919C(1,3,4,8,9)
U+4E13(0);U+5C08(1,3,8,9);U+5C08(1,3,8,9)
U+4E14(0);U+4E14(5,9);
U+4E15(0);U+4E15(5,9);
U+4E16(0);U+4E16(5,9);U+4E17(3,8),U+534B(3,8)
U+4E17(0);U+4E16(3,8,9);U+4E16(3,8,9),U+534B(3,8)
U+4E18(0);U+4E18(5,9);U+4E20(3,8),U+5775(2,3,9)
U+4E19(0);U+4E19(5,9);
U+4E1A(0);U+696D(1,3,8,9);U+696D(1,3,8,9)
U+4E1B(0);U+53E2(1,3,4,8,9),U+6B09(3,9);U+53E2(1,3,4,8,9),U+6B09(3,9),U+85C2(3,9)
U+4E1C(0);U+6771(1,3,8,9);U+6771(1,3,8,9)
U+4E1D(0);U+7D72(1,3,8,9);U+7D72(1,3,8,9),U+7CF8(1,3,8,9),U+7CF9(1,3,8),U+7E9F(1,3,8)
</pre>
<p> (22 June 2004)
<p>
</body>
</html>
`
	expected := `--09az㑇㑇㑳㑳㖞㖞㘎㘎㘚㘚㤘㤘㥮㥮㧏㧐㧟㧟㩳㩳㭎㭎㱮㱮㳠㳠䁖䁖䅟䅟䌷䌷䎬䎬䎱䎱䏝䏝䓖䓖䙌䙌䙡䙡䜣䜣䜩䜩䝼䝼䞍䞍䥇䥇䥺䥺䥽䥽䦂䦃䦅䦆䦛䦛䦟䦟䦶䦷䱷䱷䲟䲣䴓䴙䶮䶮一丑专丝`

	CheckTable(t, data, expected, true)
}

func TestProcessHTMLRanges(t *testing.T) {
	// From https://www.iana.org/domains/idn-tables/tables/bg_ru-bg_1.0.html
	data := `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=iso-8859-1">
<title>.BG IDN Table</title>
</head>
<body bgcolor="#FFFFFF">

<table>
<tr>
<td>
TLD: BG<br>
Language Identifier: ru-BG<br>
Script Description: Russian<br>

Version: 1.0<br>
Effective Date: 5 September 2009<br>
</td>
<td width="30">&nbsp;</td>
<td width="30">&nbsp;</td>
<td>
Registry: Register.BG<br/>
Contact: Daniel Kalchev daniel@digsys.bg<br/>
Address: 40, Slivnitsa blvd, Varna 9000, Bulgaria<br/>
Telephone: +359 52 694050, +359 52 694060<br/>
Website: https://www.register.bg<br/>
</td>
</tr>
</table>
<p>
	Relevant Policy Document URL:
	https://www.register.bg/user/static/rules/en/index.html
</p>
<p>
	This document presents a character table created by Register.BG for
	Russian language. The document is intended for informational purpose
	only and for exclusive use by the .BG ccTLD Registry.
</p>

<pre>
U+002D HYPEN-MINUS
U+0030..U+0039 DIGIT ZERO .. DIGIT 9
U+0430..U+044F CYRILLIC SMALL LETTER A .. CYRILLIC SMALL LETTER YA
U+0451 CYRILLIC SMALL LETTER IO
</pre>

<p>&nbsp;</p>
</body>
</html>
`
	expected := `--09аяёё`

	CheckTable(t, data, expected, true)
}

func CheckTable(t *testing.T, data string, expected string, isHTML bool) {
	table, err := ProcessIDNTable(strings.NewReader(data), isHTML, "testing")
	if err != nil {
		t.Errorf(`Error processing table: %s`, err)
	}
	out, _ := table.MarshalText()
	outStr := string(out)
	if outStr != expected {
		t.Errorf(`Expected %q, got %q`, expected, outStr)
	}
}
