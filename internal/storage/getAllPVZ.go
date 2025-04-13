package storage

import (
	"avitopvz/internal/models"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

func (db *AvitoDB) GetAllPVZ(ctx context.Context, listInfo models.GetAllPVZRequest) ([]models.PVZWithReceptions, error) {
	log.Printf("GetAllPVZ: Starting with listInfo: %+v", listInfo)

	pvzQuery := `
                SELECT id, registration_date, city 
                FROM pvz 
                ORDER BY registration_date 
                LIMIT $1 OFFSET $2`

	pvzRows, err := db.pool.Query(ctx, pvzQuery, listInfo.Limit, (listInfo.Page-1)*listInfo.Limit)
	if err != nil {
		log.Printf("GetAllPVZ: Error querying pvz: %v", err)
		return nil, fmt.Errorf("failed to query pvz: %w", err)
	}
	defer pvzRows.Close()

	var result []models.PVZWithReceptions

	for pvzRows.Next() {
		var pvz models.PVZ
		if err := pvzRows.Scan(&pvz.ID, &pvz.RegistrationDate, &pvz.City); err != nil {
			log.Printf("GetAllPVZ: Error scanning pvz: %v", err)
			return nil, fmt.Errorf("failed to scan pvz: %w", err)
		}

		receptions, err := db.getReceptionsForPVZ(ctx, pvz.ID, listInfo.StartDate, listInfo.EndDate)
		if err != nil {
			log.Printf("GetAllPVZ: Error getting receptions for pvz %s: %v", pvz.ID, err)
			return nil, fmt.Errorf("failed to get receptions for pvz %s: %w", pvz.ID, err)
		}

		result = append(result, models.PVZWithReceptions{
			PVZ:        pvz,
			Receptions: receptions,
		})
	}

	if err := pvzRows.Err(); err != nil {
		log.Printf("GetAllPVZ: Error iterating pvz rows: %v", err)
		return nil, fmt.Errorf("error in pvz rows: %w", err)
	}

	log.Printf("GetAllPVZ: Result: %+v", result)
	return result, nil
}

func (db *AvitoDB) getReceptionsForPVZ(ctx context.Context, pvzID uuid.UUID, startDate, endDate time.Time) ([]models.ReceptionWithProducts, error) {
	log.Printf("getReceptionsForPVZ: pvzID: %s, startDate: %s, endDate: %s", pvzID, startDate, endDate)

	receptionQuery := `
                SELECT id, date_time, pvz_id, status 
                FROM acceptances 
                WHERE pvz_id = $1 AND date_time BETWEEN $2 AND $3`

	receptionRows, err := db.pool.Query(ctx, receptionQuery, pvzID, startDate, endDate)
	if err != nil {
		log.Printf("getReceptionsForPVZ: Error querying acceptances: %v", err)
		return nil, fmt.Errorf("failed to query acceptances: %w", err)
	}
	defer receptionRows.Close()

	var receptions []models.ReceptionWithProducts

	for receptionRows.Next() {
		var reception models.Receptions
		if err := receptionRows.Scan(&reception.ID, &reception.DateTime, &reception.PVZID, &reception.Status); err != nil {
			log.Printf("getReceptionsForPVZ: Error scanning acceptance: %v", err)
			return nil, fmt.Errorf("failed to scan acceptance: %w", err)
		}

		products, err := db.getProductsForReception(ctx, reception.ID)
		if err != nil {
			log.Printf("getReceptionsForPVZ: Error getting products for reception %s: %v", reception.ID, err) // Логирование ошибки
			return nil, fmt.Errorf("failed to get products for reception %s: %w", reception.ID, err)
		}

		receptions = append(receptions, models.ReceptionWithProducts{
			Reception: reception,
			Products:  products,
		})
	}

	if err := receptionRows.Err(); err != nil {
		log.Printf("getReceptionsForPVZ: Error iterating reception rows: %v", err) // Логирование ошибки
		return nil, fmt.Errorf("error in reception rows: %w", err)
	}

	log.Printf("getReceptionsForPVZ: Result: %+v", receptions) // Логирование результата
	return receptions, nil
}

func (db *AvitoDB) getProductsForReception(ctx context.Context, receptionID uuid.UUID) ([]models.Product, error) {
	log.Printf("getProductsForReception: receptionID: %s", receptionID)

	productQuery := `
                SELECT id, date_time, type, acceptance_id, pvz_id 
                FROM products 
                WHERE acceptance_id = $1`

	productRows, err := db.pool.Query(ctx, productQuery, receptionID)
	if err != nil {
		log.Printf("getProductsForReception: Error querying products: %v", err)
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	defer productRows.Close()

	var products []models.Product

	for productRows.Next() {
		var product models.Product
		if err := productRows.Scan(&product.ID, &product.DateTime, &product.Type, &product.ReceptionsID, &product.PVZID); err != nil {
			log.Printf("getProductsForReception: Error scanning product: %v", err)
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	if err := productRows.Err(); err != nil {
		log.Printf("getProductsForReception: Error iterating product rows: %v", err)
		return nil, fmt.Errorf("error in product rows: %w", err)
	}

	log.Printf("getProductsForReception: Result: %+v", products)
	return products, nil
}
