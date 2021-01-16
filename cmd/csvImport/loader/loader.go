package loader

import (
	"fmt"
	"io/ioutil"

	"github.com/jszwec/csvutil"
)

type csvRowEntry struct {
	Checked       string `csv:"Checked"`
	OriginalTitle string `csv:"Original title,omitempty"`
	LocalTitle    string `csv:"Local title,omitempty"`
	Year          int16  `csv:"Year,omitempty"`
	MediaType     string `csv:"Media type,omitempty"`
	Type          string `csv:"Type,omitempty"`
	Edition       string `csv:"Edition,omitempty"`
	Country       string `csv:"Country,omitempty"`
	Case          string `csv:"Case,omitempty"`
	Discs         int8   `csv:"Discs,omitempty"`
	SubFi         string `csv:"Sub-fi,omitempty"`
	SubEn         string `csv:"Sub-en,omitempty"`
	AscpectRation string `csv:"Aspect ratio,omitempty"`
	RunningTime   string `csv:"Running time,omitempty"`
	Director      string `csv:"Director,omitempty"`
	Status        string `csv:"Status,omitempty"`
	Condition     string `csv:"Condition,omitempty"`
	Notes         string `csv:"Notes,omitempty"`
	Watched       string `cvs:"Watched,omitempty"`
	Rental        string `csv:"Rental,omitempty"`
	SlipCover     string `csv:"Slip Cover,omitempty"`
	Hologram      string `csv:"Hologram,omitempty"`
	ID            int16  `csv:"Id"`
	Barcode       string `csv:"Barcode,omitempty"`
}

// LoadCollectionItems loads collection items from cvs file in given path
func LoadCollectionItems(filePath string) error {
	var csvRowEntries []csvRowEntry

	csvInput, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err
	}

	if err := csvutil.Unmarshal(csvInput, &csvRowEntries); err != nil {
		fmt.Println("error:", err)
	}

	for _, u := range csvRowEntries {
		fmt.Printf("%+v\n", u)
	}

	return nil
}
