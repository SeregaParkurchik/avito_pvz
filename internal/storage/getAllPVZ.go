package storage

import (
	"avitopvz/internal/models"
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

func (db *AvitoDB) GetAllPVZ(ctx context.Context, listInfo models.GetAllPVZRequest) ([]models.PVZWithReceptions, error) {
	pvzQuery := `
                SELECT id, registration_date, city 
                FROM pvz 
                ORDER BY registration_date 
                LIMIT $1 OFFSET $2`

	pvzRows, err := db.pool.Query(ctx, pvzQuery, listInfo.Limit, (listInfo.Page-1)*listInfo.Limit)
	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос к pvz: %w", err)
	}
	defer pvzRows.Close()

	var result []models.PVZWithReceptions

	for pvzRows.Next() {
		var pvz models.PVZ
		if err := pvzRows.Scan(&pvz.ID, &pvz.RegistrationDate, &pvz.City); err != nil {
			return nil, fmt.Errorf("не удалось сканировать pvz: %w", err)
		}

		receptions, err := db.getReceptionsForPVZ(ctx, pvz.ID, listInfo.StartDate, listInfo.EndDate)
		if err != nil {
			return nil, fmt.Errorf("не удалось получить приемки для pvz %s: %w", pvz.ID, err)
		}

		result = append(result, models.PVZWithReceptions{
			PVZ:        pvz,
			Receptions: receptions,
		})
	}

	if err := pvzRows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка в строках pvz: %w", err)
	}

	return result, nil
}

func (db *AvitoDB) getReceptionsForPVZ(ctx context.Context, pvzID uuid.UUID, startDate, endDate time.Time) ([]models.ReceptionWithProducts, error) {
	receptionQuery := `
                SELECT id, date_time, pvz_id, status 
                FROM acceptances 
                WHERE pvz_id = $1 AND date_time BETWEEN $2 AND $3`

	receptionRows, err := db.pool.Query(ctx, receptionQuery, pvzID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос к приемкам: %w", err)
	}
	defer receptionRows.Close()

	var receptions []models.ReceptionWithProducts

	for receptionRows.Next() {
		var reception models.Receptions
		if err := receptionRows.Scan(&reception.ID, &reception.DateTime, &reception.PVZID, &reception.Status); err != nil {
			return nil, fmt.Errorf("не удалось сканировать приемку: %w", err)
		}

		products, err := db.getProductsForReception(ctx, reception.ID)
		if err != nil {
			return nil, fmt.Errorf("не удалось получить продукты для приемки %s: %w", reception.ID, err)
		}

		receptions = append(receptions, models.ReceptionWithProducts{
			Reception: reception,
			Products:  products,
		})
	}

	if err := receptionRows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка в строках приемки: %w", err)
	}

	return receptions, nil
}

func (db *AvitoDB) getProductsForReception(ctx context.Context, receptionID uuid.UUID) ([]models.Product, error) {
	productQuery := `
                SELECT id, date_time, type, acceptance_id, pvz_id 
                FROM products 
                WHERE acceptance_id = $1`

	productRows, err := db.pool.Query(ctx, productQuery, receptionID)
	if err != nil {
		return nil, fmt.Errorf("не удалось выполнить запрос к продуктам: %w", err)
	}
	defer productRows.Close()

	var products []models.Product

	for productRows.Next() {
		var product models.Product
		if err := productRows.Scan(&product.ID, &product.DateTime, &product.Type, &product.ReceptionsID, &product.PVZID); err != nil {
			return nil, fmt.Errorf("не удалось сканировать продукт: %w", err)
		}
		products = append(products, product)
	}

	if err := productRows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка в строках продукта: %w", err)
	}

	return products, nil
}
