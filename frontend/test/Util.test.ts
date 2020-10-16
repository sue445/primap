import { correctLongitude } from "../app/components/Util";

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
