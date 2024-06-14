package receipts

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"test/internal/models"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type ReceptsService struct {
}

func NewReceptsService() *ReceptsService {
	return &ReceptsService{}
}

func (s *ReceptsService) CreateAccruals(accrual models.AccrualsDTO) error {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddFont("Helvetica", "", pwd+"/helvetica_1251.json")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 16)
	tr := pdf.UnicodeTranslatorFromDescriptor("/cp1251")

	rand.Seed(time.Now().UnixNano())
	randomNumber := 100000 + rand.Intn(900000)

	pdf.Cell(0, 10, tr(fmt.Sprintf("СЧЕТ-ИЗВЕЩЕНИЯ № %v", randomNumber)))
	pdf.Ln(10)

	pdf.Cell(0, 10, tr(accrual.Addr))
	pdf.Ln(10)
	pdf.Cell(0, 10, tr("Имя:"+accrual.Name))
	pdf.Ln(20)

	// Заголовки таблицы
	columnHeaders := []string{"пред", "посл", "разн", "тариф", "Начислено(тг)"}
	colWidth := 38.0 // Ширина каждого столбца

	// Установка шрифта для заголовков таблицы

	for _, header := range columnHeaders {
		pdf.CellFormat(colWidth, 10, tr(header), "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Данные таблицы

	row := []string{
		strconv.Itoa(accrual.Past),
		strconv.Itoa(accrual.Last),
		strconv.Itoa(accrual.Substract),
		"28.22",
		strconv.Itoa(accrual.Sum),
	}
	for _, data := range row {
		pdf.CellFormat(colWidth, 10, tr(data), "1", 0, "C", false, 0, "")
	}
	pdf.Ln(20)

	pdf.Write(10, tr("Просим произвести оплату в сроки текущего месяца В случае несвоевременной оплаты оставляем за собой право начисления пени, прекращения подачи воды и передачи материалов в суд"))
	pdf.Ln(20)

	err = pdf.OutputFileAndClose(fmt.Sprintf("Счет№%v.pdf", randomNumber))
	if err != nil {
		return err
	}
	return nil
}

func (s *ReceptsService) CreateRes(pay models.PaymentsDTO) error {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddFont("Helvetica", "", pwd+"/helvetica_1251.json")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 16)
	tr := pdf.UnicodeTranslatorFromDescriptor("/cp1251")

	rand.Seed(time.Now().UnixNano())
	randomNumber := 100000 + rand.Intn(900000)

	pdf.Cell(0, 10, tr(fmt.Sprintf("КВИТАНЦИЯ ОБ ОПЛАТЕ № %v", randomNumber)))
	pdf.Ln(10)

	pdf.Cell(0, 10, tr(pay.Addr))
	pdf.Ln(20)
	pdf.Cell(0, 10, tr("Имя:"+pay.Name))
	pdf.Ln(10)
	pdf.Cell(0, 10, tr("начисленно по индевидуальным водомерам:"))
	pdf.Ln(20)

	// Заголовки таблицы
	columnHeaders := []string{"дата оплаты", "сумма оплаты"}
	colWidth := 55.0 // Ширина каждого столбца

	// Установка шрифта для заголовков таблицы

	for _, header := range columnHeaders {
		pdf.CellFormat(colWidth, 10, tr(header), "0", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	// Данные таблицы

	row := []string{
		pay.Date.Format("2006-01-02"),
		strconv.Itoa(pay.Sum),
	}
	for _, data := range row {
		pdf.CellFormat(colWidth, 10, tr(data), "0", 0, "C", false, 0, "")
	}
	pdf.Ln(20)

	pdf.Cell(0, 10, tr("ОПЛАЧЕНО"))

	err = pdf.OutputFileAndClose(fmt.Sprintf("Квитанция%v.pdf", randomNumber))
	if err != nil {
		return err
	}
	return nil
}
