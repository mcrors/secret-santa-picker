package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/mcrors/secret-santa-picker-server/domain"
	"github.com/mcrors/secret-santa-picker-server/repository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/DATA-DOG/go-sqlmock"
)

var _ = Describe("GroupsRepository", Label("unit"), func() {
	var listQuery = regexp.QuoteMeta("SELECT uuid, name, created_at FROM groups")
	var insertQuery = `^INSERT INTO groups \(uuid, name\) VALUES \(\$1, \$2\) RETURNING uuid, name, created_at$`
	var getQuery = `^SELECT uuid, name, created_at FROM groups WHERE uuid = \$1$`
	var rowColumns = []string{"uuid", "name", "created_at"}

	var id1 = "4d8bca30-553b-4fea-9738-1090c17b42e8"
	var id2 = "8318f684-ff16-4d69-9de1-e93f486e24b6"
	var name1 = "houlihans"
	var name2 = "whites"
	var createdDate = time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	var db *sql.DB
	var mock sqlmock.Sqlmock
	var dbErr error
	var repo repository.Groups

	BeforeEach(func() {
		db, mock, dbErr = sqlmock.New()
		Expect(dbErr).NotTo(HaveOccurred())
		repo = *repository.NewGroupRepository(db)
	})

	AfterEach(func() {
		db.Close()
	})

	Describe("Listing Groups", func() {

		Context("when groups exist", func() {
			It("should return a slice which includes all groups", func() {
				// Given a database with two records in it
				rows := sqlmock.NewRows(rowColumns).
					AddRow(id1, name1, createdDate).
					AddRow(id2, name2, createdDate)
				mock.ExpectQuery(listQuery).WillReturnRows(rows)

				// When ListGroups is called
				groups, err := repo.ListGroups(context.Background())

				// Then the result will include two groups
				Expect(err).NotTo(HaveOccurred())
				Expect(len(groups)).To(Equal(2))
				expected := []domain.Group{
					{ID: uuid.MustParse(id1), Name: name1, CreatedAt: createdDate},
					{ID: uuid.MustParse(id2), Name: name2, CreatedAt: createdDate},
				}

				Expect(expected).To(Equal(groups))
				Expect(mock.ExpectationsWereMet()).To(Succeed())
			})
		})

		Context("when no groups exist", func() {
			It("should return an empty slice", func() {
				// Given a database with zeros records in it
				rows := sqlmock.NewRows(rowColumns)
				mock.ExpectQuery(listQuery).WillReturnRows(rows)

				// When ListGroups is called
				groups, err := repo.ListGroups(context.Background())

				// Then the result will be an empty slice
				Expect(err).NotTo(HaveOccurred())
				Expect(groups).NotTo(BeNil())
				Expect(len(groups)).To(Equal(0))

				Expect(mock.ExpectationsWereMet()).To(Succeed())
			})
		})

		Context("when querying fails", func() {
			It("should return an error", func() {
				// Given a database where the connection drops
				mock.ExpectQuery(listQuery).
					WillReturnError(fmt.Errorf("cannot connect to database"))

				// When ListGroups is called
				_, err := repo.ListGroups(context.Background())

				// Then the results will be an error
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("cannot connect to database"))
				Expect(mock.ExpectationsWereMet()).To(Succeed())
			})
		})

		Context("when row iteration fails (rows.Err)", func() {
			It("should return an error", func() {
				rows := sqlmock.NewRows(rowColumns).
					AddRow(id1, name1, createdDate).
					AddRow(id2, name2, createdDate).
					RowError(1, fmt.Errorf("stream read failed"))

				mock.ExpectQuery(listQuery).WillReturnRows(rows)

				_, err := repo.ListGroups(context.Background())

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("stream read failed"))
				Expect(mock.ExpectationsWereMet()).To(Succeed())
			})
		})

		Context("when scanning fails", func() {
			It("should return an error", func() {
				// Given a row where created_at is the wrong type (string instead of time.Time)
				rows := sqlmock.NewRows(rowColumns).
					AddRow(id1, name1, "not-a-time")

				mock.ExpectQuery(listQuery).WillReturnRows(rows)

				// When ListGroups is called
				_, err := repo.ListGroups(context.Background())

				// Then it should error during scanning/parsing
				Expect(err).To(HaveOccurred())
				Expect(mock.ExpectationsWereMet()).To(Succeed())
			})
		})
	})

	Describe("Creating Groups", func() {

		Context("when inserting data", func() {
			It("should return the created groups", func() {

				// setup
				rows := sqlmock.NewRows(rowColumns).
					AddRow(id1, name1, createdDate)

				mock.ExpectQuery(`^INSERT INTO groups \(uuid, name\) VALUES \(\$1, \$2\) RETURNING uuid, name, created_at$`).
					WithArgs(id1, name1).
					WillReturnRows(rows)

				// Given a group
				group := domain.Group{
					ID:   uuid.MustParse(id1),
					Name: name1,
				}

				// When the group is created by the repo
				result, err := repo.CreateGroup(context.Background(), group)

				// Then the created row is returned correctly
				Expect(err).NotTo(HaveOccurred())
				Expect(group.ID).To(Equal(result.ID))
				Expect(group.Name).To(Equal(result.Name))

				Expect(mock.ExpectationsWereMet()).To(Succeed())
			})
		})

		Context("when insert fails", func() {
			It("should return an error", func() {
				// Given a database where the connection drops
				mock.ExpectQuery(insertQuery).
					WillReturnError(fmt.Errorf("cannot connect to database"))

				// When CreateGroup is called
				_, err := repo.CreateGroup(context.Background(), domain.Group{})

				// Then the results will be an error
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("cannot connect to database"))
				Expect(mock.ExpectationsWereMet()).To(Succeed())

			})
		})

		Context("when inserting duplicate uuid", func() {
			It("should return a conflict error ", func() {
				// Given an inset query the causes a unique violation error
				mock.ExpectQuery(insertQuery).
					WillReturnError(&pq.Error{
						Code: "23505", // unique violations
					})

				// When CreateGroup is called
				_, err := repo.CreateGroup(context.Background(), domain.Group{})

				// Then the result will be a
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(domain.ErrGroupConflict))
				Expect(mock.ExpectationsWereMet()).To(Succeed())
			})
		})

		Context("when scanning fails", func() {
			It("should return an error", func() {
				// Given a row where created_at is the wrong type (string instead of time.Time)
				row := sqlmock.NewRows(rowColumns).
					AddRow(id1, name1, "not-a-time")

				mock.ExpectQuery(insertQuery).WillReturnRows(row)

				// When ListGroups is called
				_, err := repo.CreateGroup(context.Background(), domain.Group{})

				// Then it should error during scanning/parsing
				Expect(err).To(HaveOccurred())
				Expect(mock.ExpectationsWereMet()).To(Succeed())
			})
		})
	})

	Describe("Getting a Group", func() {

		Context("when it is found", func() {
			It("should return the group", func() {
				// Given a database with a specfic group
				row := sqlmock.NewRows(rowColumns).
					AddRow(id1, name1, createdDate)

				mock.ExpectQuery(getQuery).
					WithArgs(id1).
					WillReturnRows(row)

				// When we try to get that group
				group, err := repo.GetGroup(context.Background(), uuid.MustParse(id1))

				// Then that group is return
				Expect(err).NotTo(HaveOccurred())
				Expect(group.ID).To(Equal(uuid.MustParse(id1)))
				Expect(group.Name).To(Equal(name1))
				Expect(group.CreatedAt).To(Equal(createdDate))
				Expect(mock.ExpectationsWereMet()).To(Succeed())
			})
		})

		Context("when it is not found", func() {
			It("should return an not found error", func() {
				// Given a database with no rows
				row := sqlmock.NewRows(rowColumns)

				mock.ExpectQuery(getQuery).
					WithArgs(id1).
					WillReturnRows(row)

				// When we try to get a gropu that doesn't exist
				_, err := repo.GetGroup(context.Background(), uuid.MustParse(id1))

				// Then we should get a not found error
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(domain.ErrGroupNotFound))
				Expect(mock.ExpectationsWereMet()).To(Succeed())
			})
		})

		Context("when the query fails", func() {
			It("should return an error", func() {
				// Given a database where the connection drops
				mock.ExpectQuery(listQuery).
					WithArgs(id1).
					WillReturnError(fmt.Errorf("cannot connect to database"))

				// When ListGroups is called
				_, err := repo.GetGroup(context.Background(), uuid.MustParse(id1))

				// Then the results will be an error
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("cannot connect to database"))
				Expect(mock.ExpectationsWereMet()).To(Succeed())

			})
		})

		Context("when scanning fails", func() {
			It("should return an error", func() {
				// Given a row where created_at is the wrong type (string instead of time.Time)
				row := sqlmock.NewRows(rowColumns).
					AddRow(id1, name1, "not-a-time")

				mock.ExpectQuery(getQuery).
					WithArgs(id1).
					WillReturnRows(row)

				// When ListGroups is called
				_, err := repo.GetGroup(context.Background(), uuid.MustParse(id1))

				// Then it should error during scanning/parsing
				Expect(err).To(HaveOccurred())
				Expect(mock.ExpectationsWereMet()).To(Succeed())
			})
		})
	})

	Describe("Renaming a Group", func() {

		Context("when the group is renamed successfully", func() {
			It("should return a nil error", func() {

			})
		})

		Context("when no records are updated", func() {
			It("should return a not found error", func() {

			})
		})

		Context("when the query fails", func() {
			It("should return an error", func() {

			})
		})

		Context("when the new name doesn't fit the contraints", func() {
			It("should return an error", func() {

			})
		})
	})

	Describe("Deleting a Group", func() {

		Context("when the group is deleted", func() {
			It("should return a nil error", func() {

			})
		})

		Context("when no records are deleted", func() {
			It("should return a not found error", func() {

			})
		})

		Context("when execute query fails", func() {
			It("should return an error", func() {

			})
		})

		Context("when attempting to delete a group with members", func() {
			It("should return an error", func() {

			})
		})
	})
})
