package main

import (
	"fmt"
	"github.com/xlwh/go-tsz"
	"github.com/xlwh/go-tsz/testdata"
	"math"
	"math/rand"
	"os"
	"text/tabwriter"
)

// collection of 24h worth of minutely points, with different characteristics.
var ConstantZero = make([]testdata.Point, 60*24)
var ConstantOne = make([]testdata.Point, 60*24)
var ConstantPos3f = make([]testdata.Point, 60*24)
var ConstantNeg3f = make([]testdata.Point, 60*24)
var ConstantPos0f = make([]testdata.Point, 60*24)
var ConstantNeg0f = make([]testdata.Point, 60*24)
var ConstantNearMaxf = make([]testdata.Point, 60*24)
var ConstantNearMinf = make([]testdata.Point, 60*24)
var ConstantNearMax0f = make([]testdata.Point, 60*24)
var ConstantNearMin0f = make([]testdata.Point, 60*24)
var Batch100ZeroOne = make([]testdata.Point, 60*24)
var FlappingZeroOne = make([]testdata.Point, 60*24)

var RandomTinyPosf = make([]testdata.Point, 60*24)
var RandomTinyf = make([]testdata.Point, 60*24)
var RandomTinyPos2f = make([]testdata.Point, 60*24)
var RandomTiny2f = make([]testdata.Point, 60*24)
var RandomTinyPos1f = make([]testdata.Point, 60*24)
var RandomTiny1f = make([]testdata.Point, 60*24)
var RandomTinyPos0f = make([]testdata.Point, 60*24)
var RandomTiny0f = make([]testdata.Point, 60*24)

var RandomSmallPosf = make([]testdata.Point, 60*24)
var RandomSmallf = make([]testdata.Point, 60*24)
var RandomSmallPos2f = make([]testdata.Point, 60*24)
var RandomSmall2f = make([]testdata.Point, 60*24)
var RandomSmallPos1f = make([]testdata.Point, 60*24)
var RandomSmall1f = make([]testdata.Point, 60*24)
var RandomSmallPos0f = make([]testdata.Point, 60*24)
var RandomSmall0f = make([]testdata.Point, 60*24)

var Random60kPosf = make([]testdata.Point, 60*24)
var Random60kf = make([]testdata.Point, 60*24)
var Random60kPos2f = make([]testdata.Point, 60*24)
var Random60k2f = make([]testdata.Point, 60*24)
var Random60kPos1f = make([]testdata.Point, 60*24)
var Random60k1f = make([]testdata.Point, 60*24)
var Random60kPos0f = make([]testdata.Point, 60*24)
var Random60k0f = make([]testdata.Point, 60*24)

var SmallTestDataPosf = make([]testdata.Point, 60*24)
var SmallTestDataf = make([]testdata.Point, 60*24)
var SmallTestDataPos0f = make([]testdata.Point, 60*24)
var SmallTestData0f = make([]testdata.Point, 60*24)

var RandomLargePosf = make([]testdata.Point, 60*24)
var RandomLargef = make([]testdata.Point, 60*24)
var RandomLargePos0f = make([]testdata.Point, 60*24)
var RandomLarge0f = make([]testdata.Point, 60*24)
var LargeTestDataPosf = make([]testdata.Point, 60*24)
var LargeTestDataPos0f = make([]testdata.Point, 60*24)
var LargeTestDataf = make([]testdata.Point, 60*24)
var LargeTestData0f = make([]testdata.Point, 60*24)

func main() {
	for i := 0; i < 60*24; i++ {
		ts := uint32(i * 60)
		ConstantZero[i] = testdata.Point{float64(0), ts}
		ConstantOne[i] = testdata.Point{float64(1), ts}
		ConstantPos3f[i] = testdata.Point{float64(1234.567), ts}
		ConstantNeg3f[i] = testdata.Point{float64(-1234.567), ts}
		ConstantPos0f[i] = testdata.Point{float64(1234), ts}
		ConstantNeg0f[i] = testdata.Point{float64(-1235), ts}
		ConstantNearMaxf[i] = testdata.Point{math.MaxFloat64 / 100, ts}
		ConstantNearMinf[i] = testdata.Point{-math.MaxFloat64 / 100, ts}
		ConstantNearMax0f[i] = testdata.Point{math.Floor(ConstantNearMaxf[i].V), ts}
		ConstantNearMin0f[i] = testdata.Point{math.Floor(ConstantNearMinf[i].V), ts}
		if i%200 < 100 {
			Batch100ZeroOne[i] = testdata.Point{float64(0), ts}
		} else {
			Batch100ZeroOne[i] = testdata.Point{float64(1), ts}
		}
		if i%2 == 0 {
			FlappingZeroOne[i] = testdata.Point{float64(0), ts}
		} else {
			FlappingZeroOne[i] = testdata.Point{float64(1), ts}
		}

		RandomTinyPosf[i] = testdata.Point{rand.ExpFloat64(), ts} // 0-inf, but most vals are very low, mostly between 0 and 2, rarely goes over 10
		RandomTinyf[i] = testdata.Point{rand.NormFloat64(), ts}   // -inf to + inf, as many pos as neg, but similar as above, rarely goes under -10 or over +10
		RandomTinyPos2f[i] = testdata.Point{RoundNum(RandomTinyPosf[i].V, 2), ts}
		RandomTiny2f[i] = testdata.Point{RoundNum(RandomTinyf[i].V, 2), ts}
		RandomTinyPos1f[i] = testdata.Point{RoundNum(RandomTinyPosf[i].V, 1), ts}
		RandomTiny1f[i] = testdata.Point{RoundNum(RandomTinyf[i].V, 1), ts}
		RandomTinyPos0f[i] = testdata.Point{math.Floor(RandomTinyPosf[i].V), ts}
		RandomTiny0f[i] = testdata.Point{math.Floor(RandomTinyf[i].V), ts}

		RandomSmallPosf[i] = testdata.Point{RandomTinyPosf[i].V * 100, ts} // 0-inf, but most vals are very low, mostly between 0 and 200, rarely goes over 1000
		RandomSmallf[i] = testdata.Point{RandomTinyf[i].V * 100, ts}       // -inf to + inf, as many pos as neg, but similar as above, rarely goes under -1000 or over +1000
		RandomSmallPos2f[i] = testdata.Point{RoundNum(RandomSmallPosf[i].V, 2), ts}
		RandomSmall2f[i] = testdata.Point{RoundNum(RandomSmallf[i].V, 2), ts}
		RandomSmallPos1f[i] = testdata.Point{RoundNum(RandomSmallPosf[i].V, 1), ts}
		RandomSmall1f[i] = testdata.Point{RoundNum(RandomSmallf[i].V, 1), ts}
		RandomSmallPos0f[i] = testdata.Point{math.Floor(RandomSmallPosf[i].V), ts}
		RandomSmall0f[i] = testdata.Point{math.Floor(RandomSmallf[i].V), ts}

		Random60kPosf[i] = testdata.Point{rand.Float64() * 60000, ts}
		Random60kf[i] = testdata.Point{Random60kPosf[i].V, ts}
		if rand.Int()%2 == 0 {
			Random60kf[i].V *= -1.0
		}
		Random60kPos2f[i] = testdata.Point{RoundNum(Random60kPosf[i].V, 2), ts}
		Random60k2f[i] = testdata.Point{RoundNum(Random60kf[i].V, 2), ts}
		Random60kPos1f[i] = testdata.Point{RoundNum(Random60kPosf[i].V, 1), ts}
		Random60k1f[i] = testdata.Point{RoundNum(Random60kf[i].V, 1), ts}
		Random60kPos0f[i] = testdata.Point{math.Floor(Random60kPosf[i].V), ts}
		Random60k0f[i] = testdata.Point{math.Floor(Random60kf[i].V), ts}

		SmallTestDataPosf[i] = testdata.Point{float64(testdata.TwoHoursData[i%120].V) * 1.234567, ts} // THD is 650-680, so this is 0-150
		if rand.Int()%2 == 0 {
			SmallTestDataf[i] = testdata.Point{SmallTestDataPosf[i].V, ts} // -150 - 150
		} else {
			SmallTestDataf[i] = testdata.Point{-1 * SmallTestDataPosf[i].V, ts}
		}
		SmallTestDataPos0f[i] = testdata.Point{math.Floor(SmallTestDataPosf[i].V), ts} // 0-150
		SmallTestData0f[i] = testdata.Point{math.Floor(SmallTestDataf[i].V), ts}       // -150 - 150

		RandomLargePosf[i] = testdata.Point{rand.ExpFloat64() * 0.0001 * math.MaxFloat64, ts} // 0-inf, rarely goes over maxfloat/1000
		RandomLargef[i] = testdata.Point{rand.NormFloat64() * 0.0001 * math.MaxFloat64, ts}   // same buth also negative
		RandomLargePos0f[i] = testdata.Point{math.Floor(RandomLargePosf[i].V), ts}
		RandomLarge0f[i] = testdata.Point{math.Floor(RandomLargef[i].V), ts}

		LargeTestDataPosf[i] = testdata.Point{float64(testdata.TwoHoursData[i%120].V) * 0.00001234567 * math.MaxFloat64, ts} // 0-maxfloat/1000
		if rand.Int()%2 == 0 {
			LargeTestDataf[i] = testdata.Point{LargeTestDataPosf[i].V, ts} // -maxfloat/1000 ~maxfloat/1000
		} else {
			LargeTestDataf[i] = testdata.Point{-1 * LargeTestDataPosf[i].V, ts}
		}

		LargeTestDataPos0f[i] = testdata.Point{math.Floor(LargeTestDataPosf[i].V), ts} // 0-maxfloat/1000
		LargeTestData0f[i] = testdata.Point{math.Floor(LargeTestDataf[i].V), ts}       // -mf/1000 ~ mx/1000
	}

	intervals := []int{10, 30, 60, 120, 360, 720, 1440}
	do := func(data []testdata.Point, comment string) string {
		str := ""
		for _, points := range intervals {
			s := tsz.New(data[0].T)
			for _, tt := range data[0:points] {
				s.Push(tt.T, tt.V)
			}
			size := len(s.Bytes())
			BPerPoint := float64(size) / float64(points)
			str += fmt.Sprintf("\033[31m%d\033[39m\t%.2f\t", size, BPerPoint)
		}
		str += comment + "\t"
		return str
	}
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 5, 0, 1, ' ', tabwriter.AlignRight)
	fmt.Println("=== help ===")
	fmt.Println("CS = chunk size in Bytes")
	fmt.Println("BPP = Bytes per point (CS/num-points)")
	fmt.Println("d = integers stored as float64")
	fmt.Println("f = float64's with a bunch of decimal numbers")
	fmt.Println(".Xf = float64's with X decimal numbers")
	fmt.Println("[num1] a - b [num2]: a range between a and b with the occasional outliers up to num1 and num2")
	fmt.Println("=== data ===")
	str := "test"
	for _, points := range intervals {
		str += fmt.Sprintf("\t  \033[39m%dCS\033[39m\t%dBPP", points, points)
	}
	cmtTinyPos := "0 ~ 10 [inf]"
	cmtTinyPosNeg := "[-inf] -10 ~ 10 [inf]"
	cmtSmallPos := "0 ~ 1000 [inf]"
	cmtSmallPosNeg := "[-inf] -1000 ~ 1000 [inf]"
	cmt60kPos := "0 ~60k"
	cmt60kPosNeg := "-60k ~ 60k"
	cmtSmallTestPos := "0~150"
	cmtSmallTestPosNeg := "-150~150"
	cmtRandomLargePos := "0 ~ MaxFloat64/1000 [inf]"
	cmtRandomLargePosNeg := "[-inf] -MaxFloat64/1000 ~ MaxFloat64/1000 [inf]"
	cmtLargeTestPos := "0 ~ MaxFloat64/1000"
	cmtLargeTestPosNeg := "-MaxFloat64/1000 ~ MaxFloat64/1000"
	fmt.Fprintln(w, str+"\tcomment\t")
	fmt.Fprintln(w, "constant zero            d\t"+do(ConstantZero, ""))
	fmt.Fprintln(w, "constant one             d\t"+do(ConstantOne, ""))
	fmt.Fprintln(w, "constant pos           .3f\t"+do(ConstantPos3f, ""))
	fmt.Fprintln(w, "constant neg           .3f\t"+do(ConstantNeg3f, ""))
	fmt.Fprintln(w, "constant pos           .0f\t"+do(ConstantPos0f, ""))
	fmt.Fprintln(w, "constant neg           .0f\t"+do(ConstantNeg0f, ""))
	fmt.Fprintln(w, "constant nearmax         f\t"+do(ConstantNearMaxf, ""))
	fmt.Fprintln(w, "constant nearmin         f\t"+do(ConstantNearMinf, ""))
	fmt.Fprintln(w, "constant nearmax       .0f\t"+do(ConstantNearMax0f, ""))
	fmt.Fprintln(w, "constant nearmin       .0f\t"+do(ConstantNearMin0f, ""))
	fmt.Fprintln(w, "batch100 zero/one        d\t"+do(Batch100ZeroOne, ""))
	fmt.Fprintln(w, "flapping zero/one        d\t"+do(FlappingZeroOne, ""))
	fmt.Fprintln(w, "\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t")
	fmt.Fprintln(w, "random tiny pos       f\t"+do(RandomTinyPosf, cmtTinyPos))
	fmt.Fprintln(w, "random tiny pos/neg   f\t"+do(RandomTinyf, cmtTinyPosNeg))
	fmt.Fprintln(w, "random tiny pos     .2f\t"+do(RandomTinyPos2f, cmtTinyPos))
	fmt.Fprintln(w, "random tiny pos/neg .2f\t"+do(RandomTiny2f, cmtTinyPosNeg))
	fmt.Fprintln(w, "random tiny pos     .1f\t"+do(RandomTinyPos1f, cmtTinyPos))
	fmt.Fprintln(w, "random tiny pos/neg .1f\t"+do(RandomTiny1f, cmtTinyPosNeg))
	fmt.Fprintln(w, "random tiny pos     .0f\t"+do(RandomTinyPos0f, cmtTinyPos))
	fmt.Fprintln(w, "random tiny pos/neg .0f\t"+do(RandomTiny0f, cmtTinyPosNeg))
	fmt.Fprintln(w, "\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t")
	fmt.Fprintln(w, "testdata small pos       f\t"+do(SmallTestDataPosf, cmtSmallTestPos))
	fmt.Fprintln(w, "testdata small pos/neg   f\t"+do(SmallTestDataf, cmtSmallTestPosNeg))
	fmt.Fprintln(w, "testdata small pos     .0f\t"+do(SmallTestDataPos0f, cmtSmallTestPos))
	fmt.Fprintln(w, "testdata small pos/neg .0f\t"+do(SmallTestData0f, cmtSmallTestPosNeg))
	fmt.Fprintln(w, "\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t")
	fmt.Fprintln(w, "random small pos        f\t"+do(RandomSmallPosf, cmtSmallPos))
	fmt.Fprintln(w, "random small pos/neg    f\t"+do(RandomSmallf, cmtSmallPosNeg))
	fmt.Fprintln(w, "random small pos      .2f\t"+do(RandomSmallPos2f, cmtSmallPos))
	fmt.Fprintln(w, "random small pos/neg  .2f\t"+do(RandomSmall2f, cmtSmallPosNeg))
	fmt.Fprintln(w, "random small pos      .1f\t"+do(RandomSmallPos1f, cmtSmallPos))
	fmt.Fprintln(w, "random small pos/neg  .1f\t"+do(RandomSmall1f, cmtSmallPosNeg))
	fmt.Fprintln(w, "random small pos      .0f\t"+do(RandomSmallPos0f, cmtSmallPos))
	fmt.Fprintln(w, "random small pos/neg  .0f\t"+do(RandomSmall0f, cmtSmallPosNeg))
	fmt.Fprintln(w, "\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t")
	fmt.Fprintln(w, "random medium pos       f\t"+do(Random60kPosf, cmt60kPos))
	fmt.Fprintln(w, "random medium pos/neg   f\t"+do(Random60kf, cmt60kPosNeg))
	fmt.Fprintln(w, "random medium pos     .2f\t"+do(Random60kPos2f, cmt60kPos))
	fmt.Fprintln(w, "random medium pos/neg .2f\t"+do(Random60k2f, cmt60kPosNeg))
	fmt.Fprintln(w, "random medium pos     .1f\t"+do(Random60kPos1f, cmt60kPos))
	fmt.Fprintln(w, "random medium pos/neg .1f\t"+do(Random60k1f, cmt60kPosNeg))
	fmt.Fprintln(w, "random medium pos     .0f\t"+do(Random60kPos0f, cmt60kPos))
	fmt.Fprintln(w, "random medium pos/neg .0f\t"+do(Random60k0f, cmt60kPosNeg))
	fmt.Fprintln(w, "\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t")
	fmt.Fprintln(w, "testdata large pos       f\t"+do(LargeTestDataPosf, cmtLargeTestPos))
	fmt.Fprintln(w, "testdata large pos/neg   f\t"+do(LargeTestDataf, cmtLargeTestPosNeg))
	fmt.Fprintln(w, "testdata large pos     .0f\t"+do(LargeTestDataPos0f, cmtLargeTestPos))
	fmt.Fprintln(w, "testdata large pos/neg .0f\t"+do(LargeTestData0f, cmtLargeTestPosNeg))
	fmt.Fprintln(w, "\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t")
	fmt.Fprintln(w, "random large pos        f\t"+do(RandomLargePosf, cmtRandomLargePos))
	fmt.Fprintln(w, "random large pos/neg    f\t"+do(RandomLargef, cmtRandomLargePosNeg))
	fmt.Fprintln(w, "random large pos      .0f\t"+do(RandomLargePos0f, cmtRandomLargePos))
	fmt.Fprintln(w, "random large pos/neg  .0f\t"+do(RandomLarge0f, cmtRandomLargePosNeg))
	w.Flush()
}
