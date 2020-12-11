import {
  correctLongitude,
  formatAddress,
  getShopMarkerIconUrl,
} from "../app/components/Util";

describe.each([
  [0, 0],
  [180, -180],
  [800, 80],
  [-400, -40],
])(".correctLongitude()", (longitude, expected) => {
  test(`correctLongitude(${longitude}) returns ${expected}`, () => {
    expect(correctLongitude(longitude)).toBe(expected);
  });
});

describe.each([
  [null, ""],
  ["東京都新宿区新宿３－３５－８", "東京都新宿区新宿３－３５－８"],
  ["東京都新宿区　　　新宿３－２２－７", "東京都新宿区　新宿３－２２－７"],
  [
    "福岡県福岡市博多区住吉１－２－２２　　キャナルシティ博多４Ｆ",
    "福岡県福岡市博多区住吉１－２－２２　キャナルシティ博多４Ｆ",
  ],
  [
    "北海道滝川市東町２丁目２９－１　　　　　　　イオン滝川店１Ｆ",
    "北海道滝川市東町２丁目２９－１　イオン滝川店１Ｆ",
  ],
])(".formatAddress()", (address, expected) => {
  test(`formatAddress(${address}) returns ${expected}`, () => {
    expect(formatAddress(address)).toBe(expected);
  });
});

describe.each([
  ["タイトーステーション新宿南口ゲームワールド", null],
  ["プリズムストーン 原宿", "/img/marker_prismstone.png"],
])(".getShopMarkerIconUrl()", (shopName, expected) => {
  test(`getShopMarkerIconUrl(${shopName}) returns ${expected}`, () => {
    expect(getShopMarkerIconUrl(shopName)).toBe(expected);
  });
});
