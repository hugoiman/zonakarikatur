package models

import "zonakarikatur/db"

// Faq is class
type Faq struct {
	IDFaq    int    `json:"idFaq"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// Faqs is list of Faq
type Faqs struct {
	Faqs []Faq `json:"faqs"`
}

// GetFaqs is func
func GetFaqs() Faqs {
	con := db.Connect()
	query := "SELECT idFaq, question, answer FROM faq"
	rows, _ := con.Query(query)

	faq := Faq{}
	faqs := Faqs{}

	for rows.Next() {
		rows.Scan(
			&faq.IDFaq, &faq.Question, &faq.Answer,
		)

		faqs.Faqs = append(faqs.Faqs, faq)
	}

	defer con.Close()

	return faqs
}

// GetFaq is func
func GetFaq(idFaq string) (Faq, error) {
	con := db.Connect()
	query := "SELECT idFaq, question, answer FROM faq WHERE idFaq = ?"

	faq := Faq{}

	err := con.QueryRow(query, idFaq).Scan(
		&faq.IDFaq, &faq.Question, &faq.Answer)

	defer con.Close()
	return faq, err
}

// CreateFaq is func
func CreateFaq(data Faq) (int, error) {
	con := db.Connect()
	query := "INSERT INTO faq (question, answer) VALUES (?,?)"
	exec, err := con.Exec(query, data.Question, data.Answer)

	idInt64, _ := exec.LastInsertId()
	idFaq := int(idInt64)
	defer con.Close()

	return idFaq, err
}

// DeleteFaq is func
func DeleteFaq(idFaq string) error {
	con := db.Connect()
	query := "DELETE FROM faq WHERE idFaq = ?"
	_, err := con.Exec(query, idFaq)

	defer con.Close()

	return err
}
