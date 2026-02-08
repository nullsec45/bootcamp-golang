package repositories

import (
	"database/sql"
	// "errors"
	"kasir-api/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetSummary(start string, end string) (*models.SalesSummary, error) {
	var summary models.SalesSummary
	
	err := r.db.QueryRow(`
        SELECT COALESCE(SUM(total_amount), 0), COUNT(id) 
        FROM transactions 
        WHERE created_at BETWEEN $1 AND $2`, start, end).Scan(&summary.TotalRevenue, &summary.TotalTransaksi)

	err = r.db.QueryRow(`
        SELECT products.name, SUM(quantity) as qty 
        FROM transaction_details
		LEFT JOIN products ON transaction_details.product_id = products.id
        GROUP BY products.name ORDER BY qty DESC LIMIT 1`,
    ).Scan(&summary.ProdukTerlaris.Nama, &summary.ProdukTerlaris.QtyTerjual)

	if err != nil {
		return nil, err
	}

    return &summary, nil
}