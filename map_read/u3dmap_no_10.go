package map_read

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/smartwalle/math4go"
)

func Test2() {

	//testP := Point{
	//	X: 130,
	//	Y: 0,
	//	Z: 157,
	//}
	//testT := Triangle{
	//	A: Point{
	//		X: 130.1667,
	//		Y: 2.511756,
	//		Z: 156.1667,
	//	},
	//	B: Point{
	//		X: 129.1667,
	//		Y: 2.261756,
	//		Z: 157,
	//	},
	//	C: Point{
	//		X: 130.6667,
	//		Y: 2.511756,
	//		Z: 156.5,
	//	},
	//	maxX: 130.6667,
	//	minX: 129.1667,
	//	maxZ: 157,
	//	minZ: 156.1667,
	//}
	//in := IsPointInTriangle(testP, testT)
	//in2 := simpleCheckNotInTriangle(testP, testT)
	//fmt.Printf("in: %v\n", in)
	//fmt.Printf("in2: %v\n", in2)
	//
	//return

	startPoint := time.Now()

	lines := ReadObj2()
	err := ParseData2(lines)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	DrawPoint2(160, 160)

	//p := Point{
	//	X: 0.7999999999999988,
	//	Y: 0,
	//	Z: -0.7999999999999988,
	//}
	//exist := false
	//for _, t := range Triangles {
	//	if IsPointInTriangle(p, t) {
	//		fmt.Printf("point %v in.", p)
	//}
	//		exist = true
	//		break
	//	}
	//}
	//
	//if !exist {
	//	fmt.Printf("point %v not in.", p)
	//}

	fmt.Printf("花费时间ms：%d\n", time.Now().Sub(startPoint).Milliseconds())

	return
}

func DrawPoint2(wide int32, hight int32) {
	Draw = make([][]int32, hight)
	for idx := int32(0); idx < hight; idx++ {
		Draw[idx] = make([]int32, wide)
	}

	ps := make([]Point, 0)
	for hIdx := 0.0; hIdx > float64(-hight); hIdx -= 1 {
		for wIdx := 0.0; wIdx < float64(wide); wIdx += 1 {

			//note: 排除wide/-hight位置的边缘点
			//w := math4go.Round(wIdx, 1)
			//h := math4go.Round(hIdx, 1)
			//if w >= math4go.Round(float64(wide), 1) || h <= math4go.Round(float64(-hight), 1) {
			//	break
			//}

			//note: 此处不能使用 w/h 给 X/Z 赋值，否则会丢失精度
			ps = append(ps, Point{
				X: wIdx,
				Y: 0.0,
				Z: hIdx,
			})
		}
	}

	fmt.Println("====================")
	fmt.Printf("len(ps): %d\n", len(ps))
	fmt.Println("====================")

	routineCnt := int32(runtime.NumCPU())
	everyLen := int32(len(ps) / runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(int(routineCnt))

	for i := int32(0); i < routineCnt; i++ {

		startIdx := int32(i * everyLen)
		stopIdx := int32(startIdx + everyLen - 1)
		go func(start int32, stop int32) {

			defer func() {
				fmt.Printf("routine start %d exit.\n", start)
				wg.Done()
			}()

			tmpPs := ps[start:stop]

			for _, p := range tmpPs {
				for _, t := range Triangles {

					if simpleCheckNotInTriangle(p, t) {
						continue
					}

					if IsPointInTriangle(p, t) {
						//修改Draw对应坐标值为1
						w := int32(math4go.Round(p.X, 1))
						h := int32(math4go.Round(-p.Z, 1))
						Draw[h][w] = 1
						break
					}
				}
			}

		}(startIdx, stopIdx)
	}

	wg.Wait()

	for _, content := range Draw {
		fmt.Printf("%v\n", content)
	}

	//返回一个矩形
	rectangle := image.Rect(0, 0, int(wide), int(hight))
	rgba := image.NewRGBA(rectangle)

	for z := int32(0); z < hight; z++ {
		for x := int32(0); x < wide; x++ {

			if Draw[z][x] == 1 {
				rgba.Set(int(x), int(z), color.White)
			} else {
				rgba.Set(int(x), int(z), color.Black)
			}
		}
	}

	//创建图片
	file, err := os.Create("./text.jpg")
	if err != nil {
		fmt.Println("os.Open error : ", err)
		return
	}
	defer file.Close()

	// 将图像写入file
	//&jpeg.Options{100} 取值范围[1,100]，越大图像编码质量越高
	jpeg.Encode(file, rgba, &jpeg.Options{100})

	//fmt.Printf("height: %d\n", len(Draw))
	//fmt.Printf("width: %d\n", len(Draw[0]))

}

func ReadObj2() []string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	defer file.Close()

	lines := make([]string, 0)
	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}

	//fmt.Printf("lines: %v\n", lines)

	return lines
}

func ParseData2(lines []string) error {
	var err error

	for _, line := range lines {

		if line == "" {
			break
		}

		if line[0] == 'v' {
			//处理点坐标
			line = line[1:]

			//找到第一个不为空格点坐标
			index := strings.IndexFunc(line, func(r rune) bool {
				if r != rune(' ') {
					return true
				}
				return false
			},
			)
			line = line[index:]

			points := strings.Split(line, " ")
			var p Point
			if p.X, err = strconv.ParseFloat(points[0], 64); err != nil {
				return err
			}
			if p.Y, err = strconv.ParseFloat(points[1], 64); err != nil {
				return err
			}
			if p.Z, err = strconv.ParseFloat(points[2], 64); err != nil {
				return err
			}

			if p.Z > 0 {
				fmt.Println("该代码针对z坐标为负数的obj地图文件，若是解析z坐标为正数的obj地图文件，需要修改 DrawPoint（137/189行）")
				os.Exit(-1)
			}

			TopPoints = append(TopPoints, p)
		}

		if line[0] == 'f' {

			//处理三角形
			line = line[1:]
			//找到第一个不为空格点坐标
			index := strings.IndexFunc(line, func(r rune) bool {
				if r != rune(' ') {
					return true
				}
				return false
			},
			)
			line = line[index:]

			pointIndexs := strings.Split(line, " ")
			var t Triangle
			var pAIndex int64
			var pBIndex int64
			var pCIndex int64
			if pAIndex, err = strconv.ParseInt(pointIndexs[0], 10, 64); err != nil {
				return err
			}
			if pBIndex, err = strconv.ParseInt(pointIndexs[1], 10, 64); err != nil {
				return err
			}
			if pCIndex, err = strconv.ParseInt(pointIndexs[2], 10, 64); err != nil {
				return err
			}
			t.A = TopPoints[pAIndex-1]
			t.B = TopPoints[pBIndex-1]
			t.C = TopPoints[pCIndex-1]
			t.maxX, t.minX, t.maxZ, t.minZ = getRectangleXY(t)

			Triangles = append(Triangles, t)
		}
	}

	for _, triangle := range Triangles {
		fmt.Printf("%v\n", triangle)
	}

	return nil
}
