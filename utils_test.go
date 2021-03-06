package main

import (
	"reflect"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestZipGeom(t *testing.T) {
	got := ZipGeom([]string{"C", "C", "C", "H", "H"},
		[]float64{
			0.000000000, 0.000000000, -0.888739044,
			0.000000000, 0.662758874, 0.368259613,
			0.000000000, -0.662758874, 0.368259613,
			0.000000000, 1.595272034, 0.906954891,
			0.000000000, -1.595272034, 0.906954891,
		})
	want := `C      0.000000000000      0.000000000000     -0.888739044000
C      0.000000000000      0.662758874000      0.368259613000
C      0.000000000000     -0.662758874000      0.368259613000
H      0.000000000000      1.595272034000      0.906954891000
H      0.000000000000     -1.595272034000      0.906954891000
`
	if got != want {
		t.Errorf("got\n%v, wanted\n%v\n", got, want)
	}
}

func TestSub(t *testing.T) {
	got := Sub(
		[]float64{1, 2, 3},
		[]float64{4, 8, 16},
	)
	want := []float64{-3, -6, -13}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
}

func TestScale(t *testing.T) {
	got := Scale(2.0, []float64{1, 2, 3})
	want := []float64{2, 4, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		a, b float64
		want bool
	}{
		{
			a:    1.0000000000000001,
			b:    1.0,
			want: true,
		},
		{
			a:    1.1,
			b:    1.0,
			want: false,
		},
	}
	for _, test := range tests {
		got := Equal(test.a, test.b)
		if got != test.want {
			t.Errorf("got %v, wanted %v\n", got, test.want)
		}
	}
}

func TestIdentity(t *testing.T) {
	got := Identity(3)
	want := *mat.NewDense(3, 3, []float64{
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	})
	if !reflect.DeepEqual(got.RawMatrix().Data, want.RawMatrix().Data) {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
}

func TestTrimExt(t *testing.T) {
	got := TrimExt("test.in")
	want := "test"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v\n", got, want)
	}
}
