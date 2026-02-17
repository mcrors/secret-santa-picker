package repository_test

import (
	"context"
	"time"

	"github.com/mcrors/secret-santa-picker-server/repository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/DATA-DOG/go-sqlmock"
)

var _ = Describe("Group", Label("unit"), func() {

	Describe("Listing Groups", func() {

		Context("when groups exist", func() {
			It("should return a slice which includes all groups", func() {
				db, mock, err := sqlmock.New()
				Expect(err).NotTo(HaveOccurred())
				defer db.Close()

				// Given a database with two records in it
				query := mock.ExpectQuery("SELECT uuid, name, created_at FROM groups")
				rows := sqlmock.NewRows([]string{"uuid", "name", "created_at"}).
					AddRow("4d8bca30-553b-4fea-9738-1090c17b42e8", "houlihans", time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)).
					AddRow("8318f684-ff16-4d69-9de1-e93f486e24b6", "whites", time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))
				query.WillReturnRows(rows)

				// When ListGroups is called
				repo := repository.NewGroupRepository(db)
				groups, err := repo.ListGroups(context.Background())

				// Then the result will include two groups
				Expect(err).NotTo(HaveOccurred())
				Expect(len(groups)).To(Equal(2))

				Expect(mock.ExpectationsWereMet()).To(Succeed())
			})
		})

		Context("when no groups exist", func() {
			It("should return an empty slice", func() {

			})
		})

		Context("when querying fails", func() {
			It("should return an error", func() {

			})
		})

		Context("when scanning fails", func() {
			It("should return an error", func() {

			})
		})
	})
})
