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
	m := pdf.NewMaroto(consts.Landscape, consts.A4)
	m.SetPageMargins(10, 10, 10)

	buildHeading(m)

	buildCompletedList(m)
	m.Row(10, func() {
	})
	buildProgressList(m)
	m.Row(10, func() {
	})
	buildPendingList(m)
	m.Row(10, func() {
	})
	buildCheckList(m)

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
			textH1(m, "Rekap Laporan IT Regional Kalimantan")
			textBodyCenter(m, "Dari 7 Jul 08:00 sd 7 Jul 16:00", 12)
		})
		m.ColSpace(2)
	})
}

func buildCompletedList(m pdf.Maroto) {
	tableHeading := []string{"Nama", "Kategori", "Keterangan", "Status", "Update", "Oleh"}
	contents := [][]string{
		{"GATE IN 1 TPKB", "GATE", "restart mingguan", "Completed", "7 Jul 12:25", "muchlis"},
	}
	lightPurpleColor := getLightPurpleColor()

	m.SetBackgroundColor(getTealColor())
	m.Row(9, func() {
		m.Col(12, func() {
			m.Text(" Completed", props.Text{
				Top:             2,
				Family:          consts.Courier,
				Style:           consts.Bold,
				Size:            12,
				Align:           consts.Left,
				VerticalPadding: 0,
				Color:           color.NewWhite(),
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())
	m.TableList(tableHeading, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{2, 1, 5, 1, 2, 1},
		},
		ContentProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{2, 1, 5, 1, 2, 1},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})
}

func buildProgressList(m pdf.Maroto) {
	tableHeading := []string{"Nama", "Kategori", "Keterangan", "Status", "Update", "Oleh"}
	contents := [][]string{
		{"E-BOARDING", "APPLICATION", "aplikasi Roro, kapal niki sae tidak bisa approval karna tgl kedatangan kapal tidak sesuai\n\nsudah di lakukan perubahan tgl kedatangan di permohonan, perencanaan, dan realisasi tetapi di menu pranota tgl masih salah ( di teruskan ke PIC pusat ardian masih di cek) ", "Progress", "06 Jul 22:04", "gema"},
		{"VASA", "APPLICATION", "Ipad pandu pilot00066 an Cecilia Kusuma.  tidak bisa membaca ais plug", "Progress", "6 Jul 16:08", "Muchlis"},
		{"GATE MARBA", "GATE", "Permintaaan pemindahan sementara komputer dan perangkat di gate in marba karna ada proses renovasi Gate (peninggian lantai gate) meja di pindah ke belakang dekat dengan meja gate out atau meja di mundurkan", "Progress", "4 Jul 15:53", "gema"},
		{"CCTV RTG", "OTHER", "Masih Problem RTG16 layar hidup mati ( sudah di cek kemungkinan adaptor layar CCTV bermasalah) tidak ada backup adaptor", "Progress", "29 Jun 23:10", "gema"},
	}
	lightPurpleColor := getLightPurpleColor()

	m.SetBackgroundColor(getOrangeColor())
	m.Row(9, func() {
		m.Col(12, func() {
			m.Text(" Progress", props.Text{
				Top:             2,
				Family:          consts.Courier,
				Style:           consts.Bold,
				Size:            12,
				Align:           consts.Left,
				VerticalPadding: 0,
				Color:           color.NewWhite(),
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())
	m.TableList(tableHeading, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{2, 1, 5, 1, 2, 1},
		},
		ContentProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{2, 1, 5, 1, 2, 1},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})
}

func buildPendingList(m pdf.Maroto) {
	tableHeading := []string{"Nama", "Kategori", "Keterangan", "Status", "Update", "Oleh"}
	contents := [][]string{
		{"PC VOICE RECORDER", "SERVER", "Sengaja dimatikan karena floading ke server cloud pusat", "Pending", "30 Juni 23:38", "Muchlis"},
		{"ALTAI 18", "ALTAI", "ALTAI", "Pending", "29 Jun 14:35", "Muchlis"},
		{"CY5 PARKIRAN BELAKANG - C091", "CCTV", "Indikasi terkena rembesan air hujan pasca ada renovasi di gedung tpkb, cctv sdh diturunkan vendor teknik dan menunggu cctv pengganti\npengecekan menggunakan poe baru indikatornya mati hidup mati hidup lampunya\nunutk koneksi ke server aman", "Pending", "27 Jun 05:52", "Admin"},
	}
	lightPurpleColor := getLightPurpleColor()

	m.SetBackgroundColor(getPinkColor())
	m.Row(9, func() {
		m.Col(12, func() {
			m.Text(" Pending", props.Text{
				Top:             2,
				Family:          consts.Courier,
				Style:           consts.Bold,
				Size:            12,
				Align:           consts.Left,
				VerticalPadding: 0,
				Color:           color.NewWhite(),
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())
	m.TableList(tableHeading, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{2, 1, 5, 1, 2, 1},
		},
		ContentProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{2, 1, 5, 1, 2, 1},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})
}

func buildCheckList(m pdf.Maroto) {
	tableHeading := []string{"Judul", "Lokasi", "Keterangan", "kendala", "Cek"}
	contents := [][]string{
		{"CEK RECORDING SERVICE SERVER CCTV", "SERVER Regional", "record Service aman, milestone web sementara dimatikan sedang ada penghapusan disk F dan G", "Nihil", "06:00"},
		{"CEK BENSIN MOBIL", "LAINNYA Lainnya", "", "Nihil", "07:00"},
	}
	lightPurpleColor := getLightPurpleColor()

	m.SetBackgroundColor(getTealColor())
	m.Row(9, func() {
		m.Col(12, func() {
			m.Text(" CheckList", props.Text{
				Top:             2,
				Family:          consts.Courier,
				Style:           consts.Bold,
				Size:            12,
				Align:           consts.Left,
				VerticalPadding: 0,
				Color:           color.NewWhite(),
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())
	m.TableList(tableHeading, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{4, 1, 4, 1, 2},
		},
		ContentProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{4, 1, 4, 1, 2},
		},
		Align:                consts.Left,
		AlternatedBackground: &lightPurpleColor,
		HeaderContentSpace:   1,
		Line:                 true,
	})
}

func getTealColor() color.Color {
	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}

func getOrangeColor() color.Color {
	return color.Color{
		Red:   255,
		Green: 153,
		Blue:  51,
	}
}

func getPinkColor() color.Color {
	return color.Color{
		Red:   255,
		Green: 51,
		Blue:  153,
	}
}

func getLightPurpleColor() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}

func textH1(m pdf.Maroto, text string) {
	m.Text(text, props.Text{
		Top:         3,
		Style:       consts.Bold,
		Size:        18,
		Align:       consts.Center,
		Extrapolate: true,
		Color:       getDarkColor(),
	})

}

func textBodyCenter(m pdf.Maroto, text string, top float64) {

	m.Text(text, props.Text{
		Top:         top,
		Extrapolate: false,
		Align:       consts.Center,
		Color:       getDarkGreyColor(),
	})

}

func getDarkGreyColor() color.Color {
	return color.Color{
		Red:   83,
		Green: 83,
		Blue:  83,
	}
}

func getDarkColor() color.Color {
	return color.Color{
		Red:   36,
		Green: 36,
		Blue:  36,
	}
}
