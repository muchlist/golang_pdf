package main

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"log"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	buildHeading(m)
	buildBody(m)
	buildImage(m)
	buildSignature(m)

	err := m.OutputFileAndClose("pdf/test.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Print("PDF sukses dibuat")
}

func buildHeading(m pdf.Maroto) {
	m.Row(10, func() {

	})
	m.Row(20, func() {
		m.Col(2, func() {
			err := m.FileImage("static/images/pelindo3.png", props.Rect{
				Percent: 100,
				Center:  false,
				Top:     3,
			})
			if err != nil {
				fmt.Println("Image tidak dapat di download")
			}
		})
		m.Col(8, func() {
			textH1(m, "Surat Tilang Elektronik TPKB")
			textBodyCenter(m, "Aplikasi ETI Pelindo III Banjarmasin", 12)
		})
		m.ColSpace(2)
	})
}

func buildBody(m pdf.Maroto) {
	m.Row(5, func() {

	})
	m.Row(15, func() {
		m.Col(8, func() {
			textH2(m, "Identifikasi Pelanggaran", 3)
		})
		m.Col(4, func() {
			textH2(m, "     [sangsi]", 3)
		})
	})
	m.Row(40, func() {
		m.Col(4, func() {
			textBody(m, "ID Truck / Nomer Lambung", 0)
			textBody(m, "Nomor Polisi", 5)
			textBody(m, "Lokasi", 10)
			textBody(m, "Tanggal", 15)
			textBody(m, "Pelanggaran ke-", 20)
			textBody(m, "Detail Pelanggaran", 25)
		})
		m.Col(4, func() {
			textBody(m, "0000000", 0)
			textBody(m, "DA 12345 KH", 5)
			textBody(m, "TPKB", 10)
			textBody(m, "29 Desember 2021 , 09:00 Wita", 15)
			textBody(m, "2", 20)
			textBody(m, "Sopir turun dilokasi bongkar , dan sopir tidak menggunakan apd", 25)
		})
		m.ColSpace(1)
		m.Col(3, func() {
			textBody(m, "Diberikan teguran ke dua (2), apabila melakukan pelanggaran 1 kali lagi di tahun yang sama maka akan dilakukan pemblokiran saat gate in pada truck dengan tersenbut", 0)
		})
	})
}

func buildImage(m pdf.Maroto) {
	m.Row(5, func() {

	})
	m.Row(10, func() {
		m.Col(12, func() {
			textH2(m, "Lampiran", 3)
		})
	})
	m.Row(60, func() {
		m.Col(4, func() {
			err := m.FileImage("static/images/download.jpg", props.Rect{
				Percent: 90,
				Center:  true,
			})
			if err != nil {
				fmt.Println("Image tidak dapat di download")
			}
		})
		m.Col(4, func() {
			err := m.FileImage("static/images/download.jpg", props.Rect{
				Percent: 90,
				Center:  true,
			})
			if err != nil {
				fmt.Println("Image tidak dapat di download")
			}
		})
		m.Col(4, func() {
			err := m.FileImage("static/images/portrait.png", props.Rect{
				Percent: 90,
				Center:  true,
			})
			if err != nil {
				fmt.Println("Image tidak dapat di download")
			}
		})
	})
}

func buildSignature(m pdf.Maroto) {
	m.Row(10, func() {

	})
	m.Row(8, func() {
		m.ColSpace(8)
		m.Col(4, func() {
			textBodyCenter(m, "Banjarmain 19 - 06 -2021", 0)
		})
	})
	m.Row(25, func() {
		m.ColSpace(8)
		m.Col(4, func() {
			m.QrCode("testbarcode", props.Rect{
				Top:     0,
				Percent: 100,
				Center:  true,
			})
		})
	})
	m.Row(10, func() {
		m.ColSpace(8)
		m.Col(4, func() {
			textBodyCenter(m, "HSSE TPKB", 5)
		})
	})
}

func textH1(m pdf.Maroto, text string) {
	m.Text(text, props.Text{
		Top:         3,
		Style:       consts.Bold,
		Size:        18,
		Align:       consts.Center,
		Extrapolate: true,
		Color:       getDarkPurpleColor(),
	})

}

func textH2(m pdf.Maroto, text string, top float64) {

	m.Text(text, props.Text{
		Top:         top,
		Extrapolate: false,
		Style:       consts.Bold,
		Size:        14,
		Color:       getDarkPurpleColor(),
	})

}

func textBody(m pdf.Maroto, text string, top float64) {

	m.Text(text, props.Text{
		Top:         top,
		Extrapolate: false,
		Color:       getDarkPurpleColor(),
	})

}

func textBodyCenter(m pdf.Maroto, text string, top float64) {

	m.Text(text, props.Text{
		Top:         top,
		Extrapolate: false,
		Align:       consts.Center,
		Color:       getDarkPurpleColor(),
	})

}

func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}
