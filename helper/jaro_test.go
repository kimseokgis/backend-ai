package helper

import (
	"context"
	"fmt"
	"github.com/kimseokgis/backend-ai/model"
	"testing"
)

func TestJaro(t *testing.T) {
	db := SetConnection()
	str1 := "Cara membuat server"
	data, _ := QueriesALL(db, context.TODO())
	var score float64
	var match model.Datasets
	for _, v := range data {
		str2 := v.Question
		scorex := jaroWinkler(str1, str2)
		if score < scorex {
			match = v
			score = scorex
		}
		//fmt.Printf("Jaro-Winkler similarity between '%s' and '%s' is: %f\n", str1, str2, scorex)
	}
	fmt.Printf("Mathes Word: %+v\n Score : %f\n", match, score)

}

func TestStemmer(t *testing.T) {
	sentence := "Rakyat memenuhi halaman gedung untuk menyuarakan isi hatinya. Baca berita selengkapnya di http://www.kompas.com."

	// Ubah kata berimbuhan menjadi kata dasar
	Stringer := Stemmer(sentence)
	fmt.Println(Stringer)
}
