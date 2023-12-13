package helper

import (
	"errors"
	"fmt"
	"golang.org/x/exp/constraints"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RemoveError[T any](value T, e error) T {
	check(e)
	return value
}

func ReadTextFile(filePath string) []string {
	file, err := os.ReadFile(filePath)
	check(err)
	return strings.Split(string(file), "\n")
}

func GetKeysOfSetMap[T int | string](inputMap map[T]bool) []T {
	keys := make([]T, len(inputMap))

	i := 0
	for k := range inputMap {
		keys[i] = k
		i++
	}
	return keys
}

func MeasureTime(process string) func() {
	fmt.Printf("Start %s\n", process)
	start := time.Now()
	return func() {
		fmt.Printf("Time taken by %s is %v\n", process, time.Since(start).Nanoseconds())
	}
}

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func FindRepeatedItems[T int | int64](itemgroup1, itemgroup2 []T) []T {
	elementsCompartment1 := map[T]bool{}
	repeatedElemnts := map[T]bool{}
	for _, item := range itemgroup1 {
		elementsCompartment1[item] = true
	}
	for _, item := range itemgroup2 {
		if elementsCompartment1[item] {
			repeatedElemnts[item] = true
		}
	}
	v := make([]T, 0, len(repeatedElemnts))
	for key, _ := range repeatedElemnts {
		v = append(v, key)
	}

	return v
}

func FindMax[T int | int64](slice []T) (m T) {
	for i, e := range slice {
		if i == 0 || e > m {
			m = e
		}
	}
	return
}

func Sum[T int | int64](slice []T) (s T) {
	for _, e := range slice {
		s += e
	}
	return
}

func Contains2Int[T [2]int](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func Contains[T constraints.Ordered](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func Remove[T any](s *[]T, i int) {
	(*s)[i] = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
}

func RemoveOrder[T any](s *[]T, i int) {
	*s = append((*s)[:i], (*s)[i+1:]...)
}

func Insert[T any](s *[]T, v T, i int) {
	*s = append((*s)[:i], append([]T{v}, (*s)[i:]...)...)
}

func Move[T any](s *[]T, srcIndex int, dstIndex int) {
	value := (*s)[srcIndex]
	RemoveOrder(s, srcIndex)
	Insert(s, value, dstIndex)
}

func RemoveElement[T constraints.Ordered](s []T, i T) []T {
	newS := []T{}
	for _, val := range s {
		if val != i {
			newS = append(newS, val)
		}
	}
	return newS
}

func GetValueOf2DMap[T any](location [2]int, map2D *[][]T) (T, error) {
	if location[0] >= len(*map2D) {
		return (*map2D)[0][0], errors.New("First location value too big")
	}
	if location[1] >= len((*map2D)[location[0]]) {
		return (*map2D)[0][0], errors.New("Second location value too big")
	}
	return (*map2D)[location[0]][location[1]], nil
}

func SetValueOf2DMap[T any](location [2]int, value T, map2D *[][]T) int {
	if location[0] >= len(*map2D) {
		return -1
	}
	if location[1] >= len((*map2D)[location[0]]) {
		return -1
	}
	(*map2D)[location[0]][location[1]] = value
	return 0
}

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func GCD(x int, y int) int {
	for y != 0 {
		tempY := y
		y = x % y
		x = tempY
	}
	return x
}

func LCM(x, y int) int {
	return x * y / GCD(x, y)
}

func LCMArray(input []int) int {
	lcm := input[0]
	for i := 1; i < len(input); i++ {
		lcm = LCM(lcm, input[i])
	}
	return lcm
}

func ManHattanDistance(p1, p2 [2]int) int {
	return Abs(p1[0]-p2[0]) + Abs(p1[1]-p2[1])
}

func HeapPermutation(a []int, size int) {
	if size == 1 {
		fmt.Println(a)
	}

	for i := 0; i < size; i++ {
		HeapPermutation(a, size-1)

		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}

func StringSliceToIntSlice(input []string) []int {
	intSlice := []int{}
	for _, str := range input {
		intSlice = append(intSlice, RemoveError(strconv.Atoi(strings.TrimSpace(str))))
	}
	return intSlice
}
