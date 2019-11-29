/*
Version 0.1.0

The program computes Pearson correlation coefficient matrix for a set of observations over multiple variates.

Input:
	single character separated file; default delimiter: tab, use -d flag to change
	Fields:
		variableLabel, [observation ...]
Output:
	adjacency matrix (tab separated file)
	Fields:
		variableLabel1, variableLabel2, correlation
Usage:
	peacor [flags] inFilePath outFilePath
Flags:
	-d string
		input file field delimiter, a single charater string, default: "\t"
	-f
		output full matrix, default: no main diagonal elements
	-a
		output absolute values
	NB:
		flags shouldn't be combined into a single string
Examples:
	peacor -d=' '   var2obs.ssv var2obs.adj
		space separated input
	peacor -a var2obs.tsv var2obs.adj
		output contains absolute values of correlation coefficients
	peacor -f var2obs.tsv var2obs.adj
	output includes the main diogonal of the matrix (var1:var1, var2:var2 ...)
*/
package main

import (
	"fmt"
	"encoding/csv"
	"strconv"
	"math"
	"os"
	"io"
	"log"
	"sync"
	"flag"
	"errors"
	"time"
)

/*
Represents a typical item of numerical data
*/
type Dataln struct {
	id int
	lbl string
	vals []float64
}

/*
	Calculate Pearson correlation = cov(X, Y) / (sigma_X * sigma_Y)
This function efficiently computes the correlation in one pass of the
data and makes use of the algorithm described in:
B. P. Welford, "Note on a Method for Calculating Corrected Sums of
Squares and Products", Technometrics, Vol 4, No 3, 1962.
This paper derives a numerically stable recurrence to compute a sum
of products:
S = sum_{i=1..N} [ (x_i - mu_x) * (y_i - mu_y) ]
with the relation:
S_n = S_{n-1} + ((n-1)/n) * (x_n - mu_x_{n-1}) * (y_n - mu_y_{n-1})
*/
func Correlation (data1 []float64, data2 []float64) (float64, error) {
	var (
		len1 = len(data1)
		len2 = len(data2)
	)
	if len1 != len2 {return 2.0, errors.New("input data lines must be of identical length")}
	n := len1

	var (
		i int
		sum_xsq float64 = 0.0
		sum_ysq float64 = 0.0
		sum_cross float64 = 0.0
		ratio float64
		delta_x float64
		delta_y float64
		mean_x float64
		mean_y float64
		r float64
	)
	/*
	 * Compute:
	 * sum_xsq = Sum [ (x_i - mu_x)^2 ],
	 * sum_ysq = Sum [ (y_i - mu_y)^2 ] and
	 * sum_cross = Sum [ (x_i - mu_x) * (y_i - mu_y) ]
	 * using the above relation from Welford's paper
	 */

	mean_x = data1[0]
	mean_y = data2[0]

	for i = 1; i < n; i++ {
		ratio = float64(i) / float64(i + 1.0)
		delta_x = data1[i] - mean_x
		delta_y = data2[i] - mean_y
		sum_xsq += delta_x * delta_x * ratio
		sum_ysq += delta_y * delta_y * ratio
		sum_cross += delta_x * delta_y * ratio
		mean_x += delta_x / float64(i + 1.0)
		mean_y += delta_y / float64(i + 1.0)
	}
	r = sum_cross / (math.Pow(sum_xsq, 0.5) * math.Pow(sum_ysq, 0.5))
	return r, nil
}

/*
Extract float values from a slice of strings.
*/
func Xfloats( line []string) ([]float64, error) {
	var floats []float64
	for _, vec := range line {
		if f, err := strconv.ParseFloat(vec, 64); err != nil {
			return floats, err
		} else {
			floats = append(floats, f)
		}
	}
	return floats, nil
}

/*
Pass a Dataln struct to channel
*/
func ln2ch(line []string, ind int, wg *sync.WaitGroup, ch chan Dataln) error {
	defer wg.Done() // 3. decrementing the counter at the exit of goroutine
	floats, err := Xfloats(line[1:])
	if err != nil {
		log.Println("no data for: ", line[0], ": line: ", ind)
		return err
	}
	var ln = Dataln{ind, line[0], floats}
	ch <- ln
	return nil
}

/*
Compute correlation for 2 vectors and pass it to channel
*/
func cr2ch(ln1 Dataln, ln2 Dataln, wg *sync.WaitGroup, ch chan string, aP *bool) error {
	defer wg.Done() // 3. decrementing the counter at the exit of goroutine
	lbl1 := ln1.lbl
	vec1 := ln1.vals
	lbl2 := ln2.lbl
	vec2 := ln2.vals
	cor, err := Correlation(vec1, vec2)
	if err != nil { return err }
	if *aP && cor < 0 {cor = -cor}
	out := fmt.Sprintf("%s\t%s\t%.3f\n", lbl1, lbl2, cor)
	ch <- out
	return nil
}

func main() {
	start := time.Now()
	dP := flag.String("d", "\t", "field delimiter")
	fP := flag.Bool("f", false, "output full matrix, default: no diagonal elements")
	aP := flag.Bool("a", false, "output absolute values")
	flag.Parse()
	if !flag.Parsed() {log.Fatalln("failed to parse flags")}
	args := flag.Args()
	if len(args) != 2 {log.Fatalln("Usage: [flags] inFile outFile")}
	ipth := args[0]
	opth := args[1]
	log.Println("Started", "-d=", *dP, "-f=", *fP, "-a=", *aP, ipth, opth)
	ifh, err := os.Open(ipth)
	if err != nil {log.Fatalln("Failed to open:", err)}
	defer ifh.Close()
	ofh, err := os.Create(opth)
	if err != nil {log.Fatalln("Failed to create:", err)}
	defer ofh.Close()
	reader := csv.NewReader(ifh)
	//reader := csv.NewReader(os.Stdin)
	reader.Comma = []rune(*dP)[0] // iso default comma <---- here!
	ind := 0
	wg1 := new(sync.WaitGroup) // 1. initiation , pointer
	ch1 := make(chan Dataln)
	for {
		line, err := reader.Read()
		if err == io.EOF { break }
		if err != nil { log.Fatal(err) }
		wg1.Add(1) // 2. incrementing the counter, normally in main() just before go
		go ln2ch(line, ind, wg1, ch1)
		ind ++
	}
// syncing and closing channel
	go func(wg1 *sync.WaitGroup, ch1 chan Dataln) {
		wg1.Wait() // 4. Waiting for the counter to reach 0
		close(ch1)
	}(wg1, ch1)
	datalns := make([]Dataln, 0)
// The range operator reads from the channel until the channel is closed
	for m1 := range ch1 {
		datalns = append(datalns, m1)
	}
	for _, ln1 := range(datalns) {
		wg2 := new(sync.WaitGroup) // 1. initiation , pointer
		ch2 := make(chan string)
		lbl1 := ln1.lbl
		for _, ln2 := range(datalns) {
			lbl2 := ln2.lbl
			if ! *fP && lbl1 == lbl2 {continue}
			wg2.Add(1) // 2. incrementing the counter, normally in main() just before go
			go cr2ch(ln1, ln2, wg2, ch2, aP)
		}
	// syncing and closing channel
		go func(wg2 *sync.WaitGroup, ch2 chan string) {
			wg2.Wait() // 4. Waiting for the counter to reach 0
			close(ch2)
		}(wg2, ch2)
	// The range operator reads from the channel until the channel is closed
		for m2 := range ch2 {
			_, err := ofh.Write([]byte(m2))
		if err != nil {log.Fatalln("Failed to write:", err)}
		}
	}
	log.Println("Done in", time.Since(start))
}
