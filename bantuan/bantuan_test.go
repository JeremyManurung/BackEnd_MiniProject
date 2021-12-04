package bantuan_test

import (
	"minipro/bantuan"
	"minipro/consts"
	"minipro/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateBantuanTestSuit() (bantuan.Service, bantuan.Repository, user.User, func()) {
	// Untuk mengtest bantuan kita perlu
	// bantuan service, seorang test user
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

	bantuanRepo := bantuan.NewRepository(db)

	bantuanService := bantuan.NewService(bantuanRepo)

	cleanupFunc := func() {
		userRepository.Delete(testUser.ID)
	}

	return bantuanService, bantuanRepo, testUser, cleanupFunc
}

func TestCreateBantuan(t *testing.T) {
	bantServ, bantRepo, testUser, cleanupFunc := CreateBantuanTestSuit()
	// Setelah selesai testing, hapus user testing yang dibuat
	t.Cleanup(cleanupFunc)

	newBantuan := bantuan.CreateBantuanInput{
		TittleBantuan:    "Test Bantuan",
		DeskripsiSingkat: "Deskripsi singkat bantuan",
		Deskripsi:        "Deskripsi super panjangn bantuan",
		JumlahTarget:     100,
		ListKondisi:      "-----",
		User:             testUser,
	}

	createdBantuan, err := bantServ.CreateBantuan(newBantuan)
	assert.NoError(t, err)
	t.Cleanup(func() {
		// Hapus bantuan yang dibuat untuk test ini
		bantRepo.Delete(createdBantuan.ID)
	})

	if assert.NoError(t, err) {
		assert.Equal(t, newBantuan.TittleBantuan, createdBantuan.TittleBantuan)
		assert.Equal(t, newBantuan.DeskripsiSingkat, createdBantuan.DeskripsiSingkat)
		assert.Equal(t, newBantuan.Deskripsi, createdBantuan.Deskripsi)
		assert.Equal(t, newBantuan.JumlahTarget, createdBantuan.JumlahTarget)
		assert.Equal(t, newBantuan.ListKondisi, createdBantuan.ListKondisi)
	}
}

func TestFindBantuanById(t *testing.T) {
	bantServ, bantRepo, testUser, cleanupFunc := CreateBantuanTestSuit()
	// Setelah selesai testing, hapus user testing yang dibuat
	t.Cleanup(cleanupFunc)

	newBantuan := bantuan.CreateBantuanInput{
		TittleBantuan:    "Test Bantuan",
		DeskripsiSingkat: "Deskripsi singkat bantuan",
		Deskripsi:        "Deskripsi super panjangn bantuan",
		JumlahTarget:     100,
		ListKondisi:      "-----",
		User:             testUser,
	}

	createdBantuan, err := bantServ.CreateBantuan(newBantuan)
	assert.NoError(t, err)
	t.Cleanup(func() {
		// Hapus bantuan yang dibuat untuk test ini
		bantRepo.Delete(createdBantuan.ID)
	})

	found, err := bantServ.FindBantuanByID(createdBantuan.ID)
	if assert.NoError(t, err) {
		assert.Equal(t, newBantuan.TittleBantuan, found.TittleBantuan)
		assert.Equal(t, newBantuan.DeskripsiSingkat, found.DeskripsiSingkat)
		assert.Equal(t, newBantuan.Deskripsi, found.Deskripsi)
		assert.Equal(t, newBantuan.JumlahTarget, found.JumlahTarget)
		assert.Equal(t, newBantuan.ListKondisi, found.ListKondisi)
		assert.Equal(t, newBantuan.User.ID, found.User.ID)
	}
}

func TestFindBantuans(t *testing.T) {
	bantServ, bantRepo, testUser, cleanupFunc := CreateBantuanTestSuit()
	// Setelah selesai testing, hapus user testing yang dibuat
	t.Cleanup(cleanupFunc)

	newBantuan := bantuan.CreateBantuanInput{
		TittleBantuan:    "Test Bantuan",
		DeskripsiSingkat: "Deskripsi singkat bantuan",
		Deskripsi:        "Deskripsi super panjangn bantuan",
		JumlahTarget:     100,
		ListKondisi:      "-----",
		User:             testUser,
	}

	createdBantuan, err := bantServ.CreateBantuan(newBantuan)
	assert.NoError(t, err)
	t.Cleanup(func() {
		// Hapus bantuan yang dibuat untuk test ini
		bantRepo.Delete(createdBantuan.ID)
	})

	bantuanByUser, err := bantServ.FindBantuans(testUser.ID)
	assert.NoError(t, err)
	// Cek apakah bantuan yang baru dibuat ada di FindBantuans berdasarkan user
	found := false
	for _, v := range bantuanByUser {
		if v.ID == createdBantuan.ID && v.TittleBantuan == newBantuan.TittleBantuan {
			found = true
			break
		}
	}

	assert.Equal(t, true, found)
}

func TestUpdateBantuan(t *testing.T) {
	bantServ, bantRepo, testUser, cleanupFunc := CreateBantuanTestSuit()
	// Setelah selesai testing, hapus user testing yang dibuat
	t.Cleanup(cleanupFunc)

	newBantuan := bantuan.CreateBantuanInput{
		TittleBantuan:    "Test Bantuan",
		DeskripsiSingkat: "Deskripsi singkat bantuan",
		Deskripsi:        "Deskripsi super panjangn bantuan",
		JumlahTarget:     100,
		ListKondisi:      "-----",
		User:             testUser,
	}

	createdBantuan, err := bantServ.CreateBantuan(newBantuan)
	assert.NoError(t, err)
	t.Cleanup(func() {
		// Hapus bantuan yang dibuat untuk test ini
		bantRepo.Delete(createdBantuan.ID)
	})

	updatedTitle := "updatedxxxxxx"
	updatedBantuan := bantuan.CreateBantuanInput{
		TittleBantuan:    updatedTitle,
		DeskripsiSingkat: newBantuan.Deskripsi,
		Deskripsi:        newBantuan.Deskripsi,
		JumlahTarget:     newBantuan.JumlahTarget,
		ListKondisi:      newBantuan.ListKondisi,
	}
	_, err = bantServ.UpdateBantuan(createdBantuan.ID, updatedBantuan)
	assert.NoError(t, err)

	foundUpdated, err := bantServ.FindBantuanByID(createdBantuan.ID)
	if assert.NoError(t, err) {
		assert.Equal(t, updatedTitle, foundUpdated.TittleBantuan)
	}

}
