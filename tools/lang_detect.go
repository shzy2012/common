package tools

import (
	"strings"
	"unicode"
)

// 检测语言
func DetectLanguage(s string) string {
	hasHan := false
	hasKana := false
	hasHangul := false
	hasLatin := false
	hasFrench := false
	hasCyrillic := false
	hasArabic := false

	for _, r := range s {
		switch {
		case unicode.Is(unicode.Hangul, r):
			hasHangul = true
		case unicode.Is(unicode.Hiragana, r) || unicode.Is(unicode.Katakana, r):
			hasKana = true
		case unicode.Is(unicode.Han, r):
			hasHan = true
		case unicode.Is(unicode.Arabic, r):
			hasArabic = true
		case unicode.Is(unicode.Cyrillic, r):
			hasCyrillic = true
		case IsFrenchSpec(r):
			hasFrench = true
			hasLatin = true
		case unicode.Is(unicode.Latin, r):
			hasLatin = true
		}
	}

	// 判定优先级
	if hasHangul {
		return "韩语 (Korean)"
	}
	if hasKana {
		return "日语 (Japanese)"
	}
	if hasHan {
		sc, tc := countSCTC(s)
		if tc > sc {
			return "中文 (繁体)"
		}
		return "中文 (简体)"
	}
	if hasArabic {
		return "阿文 (Arabic)"
	}
	if hasCyrillic {
		return "俄文 (Russian)"
	}
	if hasFrench || (hasLatin && IsFrenchByKeywords(s)) {
		return "法文 (French)"
	}
	if hasLatin {
		return "英文 (English)"
	}
	return "未知语言"
}

func countSCTC(s string) (sc, tc int) {
	// 简体字符集（只包含简体特有的字符）
	scSet := "发语国时欢点这经学书马门车广长体银义办备补彻尘辞罚范凤购馆归怀继讲鉴节洁惊觉开兰乐离历丽联临陆乱轮罗虑么买卖难宁农欧盘凭弃牵钱强庆穷劝确让扰热认扫审绳圣师识实势视释寿输术树帅双谁说虽随态叹腾听团弯万为务雾习戏吓显现宪县响乡协胁写兴悬选压盐养样钥药页义亿忆议译异隐应拥优邮余鱼愿跃阅云运杂灾脏赞暂责择则泽贼赠扎闸诈斋债斩辗战张涨帐胀这贞针侦诊镇阵争征睁铮筝证织职执纸质钟终种肿众昼诸猪烛属嘱贮铸筑专砖转赚庄装妆壮状锥准浊总纵组钻叙帮标产电进颗库阔里两两梦面亩纳浓旁赔鹏骗贫频评铺凄栖齐骑岂启气弃恰钎铅迁签谦浅谴枪呛墙蔷抢锹桥乔侨窍窃钦亲寝轻氢倾顷请驱权缺却鹊裙群然燃染嚷壤饶绕热认荣绒软锐瑞润洒萨腮赛伞丧嫂杀纱厦筛晒删煽衫闪陕赡伤赏烧舍摄设审婶肾渗声胜圣师诗狮湿时识实蚀驶饰适释收手守首寿受授售兽枢书赎属术树竖数帅双谁税说虽随孙损笋锁态摊贪瘫滩弹坦探汤糖膛趟烫掏逃桃陶淘讨套特疼剔梯踢提题蹄体替添田甜填条调挑跳贴铁帖厅听亭停庭挺艇头投透图徒屠土吐兔颓腿退蜕吞托拖脱驮妥拓袜歪弯湾顽万网枉妄忘旺为围违维委伟伪尾纬卫温纹闻稳问瓮卧无芜舞误误雾习席洗喜戏系细瞎虾匣峡狭吓夏仙先纤鲜闲贤咸显险县现献宪线羡乡详响项消萧硝霄销小晓孝校笑效歇蝎协胁写卸泻谢屑心辛欣新信兴星刑行形型醒幸姓性凶兄汹修羞绣锈须虚嘘叙绪续蓄宣喧悬选癣学勋训讯压押鸭牙芽崖哑亚讶咽阉烟延严言岩沿研盐颜衍掩眼演厌宴艳验谚燕殃央秧羊样洋养氧痒漾吆妖腰邀窑谣摇遥瑶咬舀药耀钥爷也页业叶曳夜液一衣医依仪夷遗疑乙已以艺亿忆义议译异翼翌绎茵荫阴吟淫殷引饮印应英樱缨鹰迎盈营蝇赢影硬映哟拥佣痈庸咏泳勇涌踊用优忧幽悠尤由犹邮油幼诱迂淤娱余鱼渔舆与屿禹宇语玉驭浴预域欲喻遇御誉鸳渊元园员圆缘源远怨院愿约越跃钥岳悦阅云匀陨运蕴杂灾载宰在再攒暂赞脏葬遭凿早枣灶噪燥躁则择泽责贼怎增赠渣扎闸眨栅榨斋债沾斩辗展占战站绽蘸张涨掌仗帐账胀障招昭找沼照罩遮折哲辙者赭褶这蔗针侦珍真诊枕疹阵振震镇争征睁蒸拯整正政帧症郑证支只汁芝枝知织肢脂蜘执直职殖止指纸址志制治质致秩序滞中忠终钟肿种众重舟周洲轴帚昼皱骤朱诸猪蛛竹逐主煮嘱瞩助住贮驻柱祝著筑铸专砖转赚撰篆桩庄装壮状妆撞追准拙捉桌浊酌啄资咨姿滋淄孜紫籽子自字恣渍综棕总纵邹走奏租族阻组钻嘴最罪尊遵左佐坐座"
	// 繁体字符集（只包含繁体特有的字符）
	tcSet := "發語國時歡點這經學書馬門車廣長體銀義辦備補徹塵辭罰範鳳購館歸懷繼講鑑節潔驚覺開蘭樂離歷麗聯臨陸亂輪羅慮麼買賣難寧農歐盤憑棄牽錢強慶窮勸確讓擾熱認掃審繩聖師識實勢視釋壽輸術樹帥雙誰說雖隨態嘆騰聽團彎萬為務霧習戲嚇顯現憲縣響鄉協脅寫興懸選壓鹽養樣鑰藥頁義億憶議譯異隱應擁優郵餘魚願躍閱雲運雜災髒贊暫責擇則澤賊贈紮閘詐齋債斬輾戰張漲帳脹這貞針偵診鎮陣爭徵睜錚箏證織職執紙質鐘終種腫眾晝諸豬燭屬囑貯鑄築專磚轉賺莊裝妝壯狀錐準濁總縱組鑽敘幫標產電進顆庫闊裏兩兩夢面畝納濃旁賠騙貧頻鋪淒棲齊騎豈啟氣棄恰釺鉛遷簽謙錢潛淺譴槍嗆墻薔搶鍬橋喬僑竅竊欽親寢輕氫傾頃請驅權缺卻鵲裙群然燃染嚷壤饒繞熱認榮絨軟銳瑞潤灑薩腮賽傘喪嫂殺紗厦篩曬刪煽衫閃陝贍傷賞燒捨攝設審嬸腎滲聲勝聖師詩獅濕時識實蝕駛飾適釋收手守首壽受授售獸樞書贖屬術樹豎數帥雙誰稅說雖隨孫損筍鎖態攤貪癱灘彈坦探湯糖膛趟燙掏逃桃陶淘討套特疼剔梯踢提題蹄體替添田甜填條調挑跳貼鐵帖廳聽亭停庭挺艇頭投透圖徒屠土吐兔頹腿退蛻吞托拖脫馱妥拓襪歪彎灣頑萬網枉妄忘旺為圍違維委偉偽尾緯衛溫紋聞穩問甕臥無蕪舞誤霧習席洗喜戲系細瞎蝦匣峽狹嚇夏仙先纖鮮閑賢咸顯險縣現獻憲線羨鄉詳響項消蕭硝霄銷小曉孝校笑效歇蠍協脅寫卸瀉謝屑心辛欣新信興星刑行形型醒幸姓性凶兄洶修羞繡鏽須虛噓敘緒續蓄宣喧懸選癬學勛訓訊壓押鴨牙芽崖啞亞訝咽閹煙延嚴言岩沿研鹽顏衍掩眼演厭宴艷驗諺燕殃央秧羊樣洋養氧癢漾吆妖腰邀窯謠搖遙瑤咬舀藥耀鑰爺也頁業葉曳夜液一衣醫依儀夷遺疑乙已以藝億憶義議譯異翼翌繹茵蔭陰吟淫殷引飲印應英櫻纓鷹迎盈營蠅贏影硬映喲擁傭癰庸詠泳勇湧踊用優憂幽悠尤由猶郵油幼誘迂淤娛餘魚漁輿與嶼禹宇語玉馭浴預域欲喻遇御譽鴛淵元園員圓緣源遠怨院願約越躍鑰岳悅閱雲勻隕運蘊雜災載宰在再攢暫贊髒葬遭鑿早棗灶噪燥躁則擇澤責賊怎增贈渣紮閘眨柵榨齋債沾斬輾展占戰站綻蘸張漲掌仗帳賬脹障招昭找沼照罩遮折哲轍者赭褶這蔗針偵珍真診枕疹陣振震鎮爭征睜蒸拯整正政幀症鄭證支只汁芝枝知織肢脂蜘執直職殖止指紙址志制治質致秩序滯中忠終鐘腫種眾重舟周洲軸帚晝皺驟朱諸豬蛛竹逐主煮囑矚助住貯駐柱祝著築鑄專磚轉賺撰篆樁莊裝壯狀妝撞追準拙捉桌濁酌啄資咨姿滋淄孜紫籽子自字恣漬綜棕總縱鄒走奏租族阻組鑽嘴最罪尊遵左佐坐座"

	for _, r := range s {
		if strings.ContainsRune(scSet, r) {
			sc++
		}
		if strings.ContainsRune(tcSet, r) {
			tc++
		}
	}
	return
}

func IsFrenchByKeywords(s string) bool {
	lower := strings.ToLower(s)
	words := strings.FieldsFunc(lower, func(r rune) bool {
		return !unicode.IsLetter(r) && r != '\''
	})
	frenchKeywords := map[string]bool{
		"c'est":   true,
		"bonjour": true,
		"merci":   true,
		"vie":     true,
		"est":     true,
		"sont":    true,
		"vous":    true,
		"nous":    true,
	}
	for _, w := range words {
		if frenchKeywords[w] {
			return true
		}
	}
	return false
}

func IsFrenchSpec(r rune) bool {
	switch r {
	case 'à', 'â', 'æ', 'ç', 'é', 'è', 'ê', 'ë', 'î', 'ï', 'ô', 'œ', 'ù', 'û', 'ü', 'ÿ',
		'À', 'Â', 'Æ', 'Ç', 'É', 'È', 'Ê', 'Ë', 'Î', 'Ï', 'Ô', 'Œ', 'Ù', 'Û', 'Ü', 'Ÿ':
		return true
	}
	return false
}
