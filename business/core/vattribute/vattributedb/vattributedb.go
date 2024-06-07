package vattributedb

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/PhyoYazar/uas/business/core/vattribute"
	"github.com/PhyoYazar/uas/business/data/order"
	database "github.com/PhyoYazar/uas/business/sys/database/pgx"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Store manages the set of APIs for user database access.
type Store struct {
	log *zap.SugaredLogger
	db  *sqlx.DB
}

// NewStore constructs the api for data access.
func NewStore(log *zap.SugaredLogger, db *sqlx.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

func (s *Store) QueryAttributeWithGaMark(ctx context.Context, filter vattribute.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]vattribute.VAttributeWithGaMark, error) {

	data := map[string]interface{}{
		"offset":        (pageNumber - 1) * rowsPerPage,
		"rows_per_page": rowsPerPage,
	}

	const q = `
		SELECT
			a.attribute_id,
         a.name,
         a.instance,
         a.type,
         m.mark_id,
         m.ga_id,
         m.mark,
			ga.slug ga_slug,
			co.co_id,
			co.name co_name,
			co.instance
		FROM
			attributes a
		LEFT JOIN
	 		marks m ON m.attribute_id = a.attribute_id
		LEFT JOIN
			graduate_attributes ga ON ga.ga_id = m.ga_id
		LEFT JOIN
			co_attributes ca ON ca.attribute_id = a.attribute_id
		LEFT JOIN
			course_outlines co ON co.co_id = ca.co_id`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	orderByClause, err := orderByClause(orderBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(orderByClause)
	buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	rows, err := database.NamedQueryRows(ctx, s.log, s.db, buf.String(), data)
	if err != nil {
		return nil, fmt.Errorf("namedqueryrows: %w", err)
	}

	var attributesMap = make(map[uuid.UUID]vattribute.VAttributeWithGaMark)
	var result []vattribute.VAttributeWithGaMark

	for rows.Next() {
		var (
			attributeID       uuid.UUID
			attributeName     string
			attributeInstance int
			attributeType     string
			markID            sql.NullString
			gaID              sql.NullString
			mark              sql.NullInt64
			gaSlug            sql.NullString
			coID              sql.NullString
			coName            sql.NullString
			coInstance        sql.NullInt64
		)

		err := rows.Scan(
			&attributeID, &attributeName, &attributeInstance, &attributeType,
			&markID, &gaID, &mark, &gaSlug,
			&coID, &coName, &coInstance,
		)
		if err != nil {
			return nil, err
		}

		attribute, ok := attributesMap[attributeID]
		if !ok {
			attribute = vattribute.VAttributeWithGaMark{
				ID:       attributeID,
				Name:     attributeName,
				Instance: attributeInstance,
				Type:     attributeType,
				Marks:    []vattribute.VMark{},
			}
			attributesMap[attributeID] = attribute
		}

		// Append Co if not NULL
		if coID.Valid && coName.Valid && coInstance.Valid {
			co := vattribute.VCo{
				ID:       uuid.MustParse(coID.String),
				Name:     coName.String,
				Instance: int(coInstance.Int64),
			}

			if coIsExist := existInSlice(attribute.Co, co); !coIsExist {
				attribute.Co = append(attribute.Co, co)
			}
		}

		// Append Mark if not NULL
		if markID.Valid && gaID.Valid && mark.Valid && gaSlug.Valid {
			mark := vattribute.VMark{
				ID:     uuid.MustParse(markID.String),
				Mark:   int(mark.Int64),
				GaID:   uuid.MustParse(gaID.String),
				GaSlug: gaSlug.String,
			}

			if markIsExist := existInSlice(attribute.Marks, mark); !markIsExist {
				attribute.Marks = append(attribute.Marks, mark)
			}
		}

		attributesMap[attributeID] = attribute
	}

	for _, value := range attributesMap {
		result = append(result, value)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// Query retrieves a list of existing subjects from the database.
func (s *Store) Query(ctx context.Context, filter vattribute.QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int, subjectID uuid.UUID) ([]vattribute.VAttribute, error) {
	data := struct {
		ID string `db:"subject_id"`
	}{
		ID: subjectID.String(),
	}

	// pdata := map[string]interface{}{
	// 	"offset":        (pageNumber - 1) * rowsPerPage,
	// 	"rows_per_page": rowsPerPage,
	// }

	const q = `
	SELECT
		a.attribute_id, a.name, a.instance, a.type, a.date_created, a.date_updated,
		co.co_id, co.name, co.instance,
		ga.ga_id, ga.name, ga.slug
	FROM
		attributes a
	LEFT JOIN
		co_attributes ca ON ca.attribute_id = a.attribute_id
	LEFT JOIN
		course_outlines co ON co.co_id = ca.co_id
	LEFT JOIN
		marks m ON m.attribute_id = a.attribute_id
	LEFT JOIN
		graduate_attributes ga ON ga.ga_id = m.ga_id
	LEFT JOIN
		co_ga cg ON cg.ga_id = ga.ga_id
	WHERE
		cg.co_id = ca.co_id
	AND
		co.subject_id = :subject_id
	AND
		m.subject_id = :subject_id`

	buf := bytes.NewBufferString(q)
	// s.applyFilter(filter, pdata, buf)

	orderByClause, err := orderByClause(orderBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(orderByClause)
	// buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	rows, err := database.NamedQueryRows(ctx, s.log, s.db, buf.String(), data)
	if err != nil {
		return nil, fmt.Errorf("namedqueryrows: %w", err)
	}

	var attributesMap = make(map[uuid.UUID]vattribute.VAttribute)
	var result []vattribute.VAttribute

	for rows.Next() {
		var (
			attributeID   uuid.UUID
			attributeName string
			instance      int
			attributeType string
			dateCreated   time.Time
			dateUpdated   time.Time
			coID          sql.NullString
			coName        sql.NullString
			coInstance    sql.NullInt64
			gaID          sql.NullString
			gaName        sql.NullString
			gaSlug        sql.NullString
		)

		err := rows.Scan(
			&attributeID, &attributeName, &instance, &attributeType, &dateCreated, &dateUpdated,
			&coID, &coName, &coInstance,
			&gaID, &gaName, &gaSlug,
		)
		if err != nil {
			return nil, err
		}

		// Retrieve or create attribute
		attribute, ok := attributesMap[attributeID]
		if !ok {
			attribute = vattribute.VAttribute{
				ID:          attributeID,
				Name:        attributeName,
				Instance:    instance,
				Type:        attributeType,
				DateCreated: dateCreated,
				DateUpdated: dateUpdated,
				Co:          []vattribute.VCo{},
				Ga:          []vattribute.VGa{},
			}
			attributesMap[attributeID] = attribute
		}

		// Append Co if not NULL
		if coID.Valid && coName.Valid && coInstance.Valid {
			co := vattribute.VCo{
				ID:       uuid.MustParse(coID.String),
				Name:     coName.String,
				Instance: int(coInstance.Int64),
			}

			if coIsExist := existInSlice(attribute.Co, co); !coIsExist {
				attribute.Co = append(attribute.Co, co)
			}
		}

		// Append Ga if not NULL
		if gaID.Valid && gaName.Valid && gaSlug.Valid {
			ga := vattribute.VGa{
				ID:   uuid.MustParse(gaID.String),
				Name: gaName.String,
				Slug: gaSlug.String,
			}

			attribute.Ga = append(attribute.Ga, ga)
		}

		attributesMap[attributeID] = attribute

		// result = append(result, attributesMap[attributeID])
	}

	for _, value := range attributesMap {
		result = append(result, value)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// // QueryByID gets the specified subject from the database.
// func (s *Store) QueryByID(ctx context.Context, coID uuid.UUID) (vco.VCo, error) {
// 	data := struct {
// 		ID string `db:"co_id"`
// 	}{
// 		ID: coID.String(),
// 	}

// 	const q = `
// 	SELECT
// 		s.subject_id, s.name, s.code, s.academic_year, s.instructor, s.semester, co.co_id, co.name,co.instance, ga.ga_id, ga.name, ga.slug, co.date_created, co.date_updated
// 	FROM
// 		course_outlines co
// 	LEFT JOIN
// 		subjects s ON co.subject_id = s.subject_id
// 	LEFT JOIN
// 		co_ga cg ON cg.co_id = co.co_id
// 	LEFT JOIN
// 		graduate_attributes ga ON ga.ga_id = cg.ga_id
// 	WHERE
// 		co.co_id = :co_id`

// 	rows, err := database.NamedQueryRows(ctx, s.log, s.db, q, data)
// 	if err != nil {
// 		if errors.Is(err, database.ErrDBNotFound) {
// 			return vco.VCo{}, fmt.Errorf("db: %w", vco.ErrNotFound)
// 		}
// 		return vco.VCo{}, fmt.Errorf("db: %w", err)
// 	}

// 	defer rows.Close()

// 	// var subject vsubject.VSubject
// 	// coMap := make(map[uuid.UUID]*vsubject.VCo)

// 	// for rows.Next() {
// 	// 	var subjectID, coID, gaID uuid.UUID
// 	// 	var subjectName, subjectCode, academicYear, instructor, semester, coName, gaName, gaSlug sql.NullString

// 	// 	err := rows.Scan(&subjectID, &subjectName, &subjectCode, &academicYear, &instructor, &semester, &coID, &coName, &gaID, &gaName, &gaSlug)
// 	// 	if err != nil {
// 	// 		log.Fatal(err)
// 	// 	}

// 	// 	// Initialize subject if it hasn't been initialized yet
// 	// 	if subject.ID == uuid.Nil {
// 	// 		subject = vsubject.VSubject{
// 	// 			ID:           subjectID,
// 	// 			Name:         subjectName.String,
// 	// 			Code:         subjectCode.String,
// 	// 			AcademicYear: academicYear.String,
// 	// 			Instructor:   instructor.String,
// 	// 			Semester:     semester.String,
// 	// 			Co:           []vsubject.VCo{},
// 	// 		}
// 	// 	}

// 	// 	// Initialize course outline if it doesn't exist in the map
// 	// 	if _, exists := coMap[coID]; !exists {
// 	// 		coMap[coID] = &vsubject.VCo{
// 	// 			ID:   coID,
// 	// 			Name: coName.String,
// 	// 			Ga:   []vsubject.VGa{},
// 	// 		}
// 	// 		subject.Co = append(subject.Co, *coMap[coID])
// 	// 	}

// 	// 	// Add graduate attribute to the course outline
// 	// 	if gaID != uuid.Nil {
// 	// 		ga := vsubject.VGa{
// 	// 			ID:   gaID,
// 	// 			Name: gaName.String,
// 	// 			Slug: gaSlug.String,
// 	// 		}
// 	// 		// Find the course outline in the subject's list and update its Ga field
// 	// 		for i := range subject.Co {
// 	// 			if subject.Co[i].ID == coID {
// 	// 				subject.Co[i].Ga = append(subject.Co[i].Ga, ga)
// 	// 				break
// 	// 			}
// 	// 		}
// 	// 	}
// 	// }

// 	// // Check for errors during row iteration
// 	// if err := rows.Err(); err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// if subject.ID == uuid.Nil {
// 	// 	return vsubject.VSubject{}, errors.New("subject not found")
// 	// }

// 	var co vco.VCo
// 	var gaMap = make(map[uuid.UUID]vco.VGa)

// 	for rows.Next() {
// 		var row vco.CourseOutlineRow
// 		err := rows.Scan(
// 			&row.SubjectID, &row.SubjectName, &row.SubjectCode, &row.AcademicYear, &row.Instructor, &row.Semester,
// 			&row.CoID, &row.CoName, &row.CoInstance,
// 			&row.GaID, &row.GaName, &row.GaSlug,
// 			&row.DateCreated, &row.DateUpdated,
// 		)
// 		if err != nil {
// 			return vco.VCo{}, err
// 		}

// 		if co.ID == uuid.Nil {
// 			co = vco.VCo{
// 				ID:       row.CoID,
// 				Name:     row.CoName,
// 				Instance: row.CoInstance,
// 				Subject: vco.VSubject{
// 					ID:           row.SubjectID,
// 					Name:         row.SubjectName,
// 					Code:         row.SubjectCode,
// 					AcademicYear: row.AcademicYear,
// 					Instructor:   row.Instructor,
// 					Semester:     row.Semester,
// 				},
// 				DateCreated: row.DateCreated,
// 				DateUpdated: row.DateUpdated,
// 				Ga:          []vco.VGa{},
// 			}
// 		}

// 		if row.GaID != uuid.Nil {
// 			if _, exists := gaMap[row.GaID]; !exists {
// 				ga := vco.VGa{
// 					ID:   row.GaID,
// 					Name: row.GaName,
// 					Slug: row.GaSlug,
// 				}
// 				co.Ga = append(co.Ga, ga)
// 				gaMap[row.GaID] = ga
// 			}
// 		}
// 	}

// 	if err := rows.Err(); err != nil {
// 		return vco.VCo{}, err
// 	}

// 	return co, nil
// }

// // Count returns the total number of subjects in the DB.
// func (s *Store) Count(ctx context.Context, filter subject.QueryFilter) (int, error) {
// 	data := map[string]interface{}{}

// 	const q = `
// 	SELECT
// 		count(1)
// 	FROM
// 		subjects`

// 	buf := bytes.NewBufferString(q)
// 	s.applyFilter(filter, data, buf)

// 	var count struct {
// 		Count int `db:"count"`
// 	}
// 	if err := database.NamedQueryStruct(ctx, s.log, s.db, buf.String(), data, &count); err != nil {
// 		return 0, fmt.Errorf("namedquerystruct: %w", err)
// 	}

// 	return count.Count, nil
// }
