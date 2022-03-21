package prishops

var shopList = []*Shop{
	// c.f. https://www.takaratomy-arts.co.jp/specials/prettyrhythm/pshj/shoplist.php
	// Note. primagiのショップ一覧はprismdbから取得するため不要
	{
		Name:       "プリズムストーン東京駅",
		Prefecture: "東京都",
		Address:    "東京都千代田区丸の内1-9-1 東京駅一番街B1F",
		Series:     []string{"pripara"},
	},
	{
		Name:       "プリズムストーンカフェ原宿",
		Prefecture: "東京都",
		Address:    "東京都渋谷区神宮前3-18-27 2F",
		Series:     []string{"pripara"},
	},
	{
		Name:       "プリズムストーン仙台",
		Prefecture: "宮城県",
		Address:    "宮城県仙台市青葉区中央4-1-1 E BeanS1F アニメガ×ソフマップ仙台駅前店内",
		Series:     []string{"pripara"},
	},
	{
		Name:       "プリズムストーン横浜",
		Prefecture: "神奈川県",
		Address:    "神奈川県横浜市西区南幸2-15-13 横浜ビブレ7F アニメガ×ソフマップ横浜ビブレ店内",
		Series:     []string{"pripara"},
	},
	{
		Name:       "プリズムストーン名古屋",
		Prefecture: "愛知県",
		Address:    "愛知県名古屋市中村区椿町6-9 ビックカメラ 名古屋駅西店 アニメガ×ソフマップ名古屋駅西店内",
		Series:     []string{"pripara"},
	},
	{
		Name:       "プリズムストーン京都",
		Prefecture: "京都府",
		Address:    "京都府京都市南区西九条鳥居口町1番 イオンモールKYOTO Sakura館 4F アニメガ×ソフマップ イオンモールKYOTO店内",
		Series:     []string{"pripara"},
	},
	{
		Name:       "プリズムストーン大阪",
		Prefecture: "大阪府",
		Address:    "大阪府大阪市浪速区日本橋３丁目６−１８ ６F アニメガ×ソフマップなんば店内",
		Series:     []string{"pripara"},
	},
	{
		Name:       "プリズムストーン神戸",
		Prefecture: "兵庫県",
		Address:    "兵庫県神戸市中央区東川崎町1-7-2 umie NORTH MALL内 6F アニメガ×ソフマップ 神戸ハーバーランド店内",
		Series:     []string{"pripara"},
	},
	{
		Name:       "プリズムストーン福岡",
		Prefecture: "福岡県",
		Address:    "福岡県福岡市中央区今泉1-25-1 ビックカメラ天神1号館Bブロック内 2F アニメガ×ソフマップ天神１号館内",
		Series:     []string{"pripara"},
	},
}
