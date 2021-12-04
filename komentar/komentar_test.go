package komentar_test

import (
	"minipro/consts"
	"minipro/komentar"
	"minipro/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func CreateKomentarTestSuit() (user.User, komentar.Repository, komentar.Service, func()) {
	// Untuk mengtest komentar kita perlu
	// komentar service, seorang test user, dan komentar repo
	var dsn = consts.DB_TEST
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	newUser := user.RegisterUserInput{
		Nama:      "Test New User",
		Pekerjaan: "Mahasiswa",
		Email:     "example@test.com",
		Password:  "12345678",
	}
	testUser, err := userService.RegisterUser(newUser)
	if err != nil {
		panic(err)
	}

	cleanupFunc := func() {
		userRepository.Delete(testUser.ID)
	}

	compRepo := komentar.NewRepository(db)
	compServ := komentar.NewService(compRepo)

	return testUser, compRepo, compServ, cleanupFunc
}

func TestCreateKomentar(t *testing.T) {
	testUser, komRepo, komServ, cleanup := CreateKomentarTestSuit()
	t.Cleanup(cleanup) // Hapus testUser setelah testing selesai

	newComment := komentar.CreateKomentarInput{
		User:        testUser,
		IsiKomentar: "Sebuah komentar",
	}

	created, err := komServ.CreateKomentar(newComment)
	t.Cleanup(func() {
		komRepo.Delete(created.ID)
	})
	if assert.NoError(t, err) {
		assert.Equal(t, testUser.ID, created.UserID)
		assert.Equal(t, newComment.IsiKomentar, created.IsiKomentar)
	}

}
