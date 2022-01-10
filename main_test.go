package main

import (
	"bytes"
	"testing"
)

func TestWriteParams(t *testing.T) {
	var b bytes.Buffer
	p, _ := LoadParams("testfiles/opt.out")
	WriteParams(&b, p)
	got := b.String()
	want := ` ****
 H
F0ss=0.5309794405
ZetaOverlap=1.2686410000
U=-0.4133181015
Beta=-0.3069665138
CoreKO=0.9416560451
GCore=0.0016794859,0.8557539975,3.3750716455
 ****
 C
F0ss=0.4900713060
F0sp=0.4236511293
F0pp=0.3644399818
F2pp=0.1978513158
G1sp=0.0790832954
ZetaOverlap=2.0475580000,1.7028410000
U=-1.8775102017,-1.4676915546
Beta=-0.5653970197,-0.2745883383
CoreKO=1.0202596926
GCore=0.0032154960,0.5881175790,2.5208171714
 ****

`
	if got != want {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
}

func TestWriteCom(t *testing.T) {
	var b bytes.Buffer
	p, _ := LoadParams("testfiles/opt.out")
	WriteCom(&b, []string{"C", "C", "C", "H", "H"},
		[]float64{
			0.0000000000, 0.0000000000, -1.6794733900,
			0.0000000000, 1.2524327590, 0.6959098120,
			0.0000000000, -1.2524327590, 0.6959098120,
			0.0000000000, 3.0146272390, 1.7138963510,
			0.0000000000, -3.0146272390, 1.7138963510,
		}, p)
	got := b.String()
	want := `%mem=1000mb
%nprocs=1
#P PM6=(print,zero,input)

the title

0 1
C      0.000000000000      0.000000000000     -0.888739044306
C      0.000000000000      0.662758874251      0.368259613354
C      0.000000000000     -0.662758874251      0.368259613354
H      0.000000000000      1.595272034246      0.906954890799
H      0.000000000000     -1.595272034246      0.906954890799

 ****
 H
F0ss=0.5309794405
ZetaOverlap=1.2686410000
U=-0.4133181015
Beta=-0.3069665138
CoreKO=0.9416560451
GCore=0.0016794859,0.8557539975,3.3750716455
 ****
 C
F0ss=0.4900713060
F0sp=0.4236511293
F0pp=0.3644399818
F2pp=0.1978513158
G1sp=0.0790832954
ZetaOverlap=2.0475580000,1.7028410000
U=-1.8775102017,-1.4676915546
Beta=-0.5653970197,-0.2745883383
CoreKO=1.0202596926
GCore=0.0032154960,0.5881175790,2.5208171714
 ****



`
	if got != want {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
}