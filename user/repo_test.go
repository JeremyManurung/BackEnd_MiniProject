package user_test

import (
	"minipro/user"
	"testing"

	"minipro/consts"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Fungsi pembantu jika unit test ingin dipisah nanti
func createUserRepoTest() (user.Repository, user.Service) {
	var dsn = consts.DB_TEST
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	return userRepository, userService
}

func TestLogin(t *testing.T) {
	userRepo, userService := createUserRepoTest()

	newUser := user.RegisterUserInput{
		Nama:      "Test New User",
		Pekerjaan: "Mahasiswa",
		Email:     "example@test.com",
		Password:  "12345678",
	}

	// Email harus belum pernah dipakai
	exist, err := userService.IsEmailAvailable(user.CheckEmailInput{Email: newUser.Email})
	if assert.NoError(t, err) {
		assert.Equal(t, true, exist)
	}

	registered, err := userService.RegisterUser(newUser)
	t.Cleanup(func() {
		// Hapus user yang digunakan saat testing setelah unit testing ini selesai
		// Agar ketika test selanjutnya berjalan, kondisi db tidak punya user testing ini
		userRepo.Delete(registered.ID)
	})

	// Setelah register harus berhasil login dengan email dan password saat mendaftar
	loginRes, err := userService.Login(user.LoginInput{
		Email:    newUser.Email,
		Password: newUser.Password,
	})
	if assert.NoError(t, err) {
		assert.Equal(t, loginRes.Nama, newUser.Nama)
		assert.Equal(t, loginRes.Pekerjaan, newUser.Pekerjaan)
		assert.Equal(t, loginRes.Email, newUser.Email)
		// Password harus bukan plaintext
		assert.NotEqual(t, loginRes.Password, newUser.Password)
	}
}

func TestEmailAvaibility(t *testing.T) {
	userRepo, userService := createUserRepoTest()

	newUser := user.RegisterUserInput{
		Nama:      "Test New User",
		Pekerjaan: "Mahasiswa",
		Email:     "example@test.com",
		Password:  "12345678",
	}

	// Email harus belum pernah dipakai
	exist, err := userService.IsEmailAvailable(user.CheckEmailInput{Email: newUser.Email})
	if assert.NoError(t, err) {
		assert.Equal(t, true, exist)
	}

	registered, err := userService.RegisterUser(newUser)
	assert.NoError(t, err)
	t.Cleanup(func() {
		// Hapus user yang digunakan saat testing setelah unit testing ini selesai
		// Agar ketika test selanjutnya berjalan, kondisi db tidak punya user testing ini
		userRepo.Delete(registered.ID)
	})

	if assert.NoError(t, err) {
		assert.Equal(t, registered.Nama, newUser.Nama)
		assert.Equal(t, registered.Pekerjaan, newUser.Pekerjaan)
		assert.Equal(t, registered.Email, newUser.Email)
		// Password harus bukan plaintext
		assert.NotEqual(t, registered.Password, newUser.Password)
	}

	// Email setelah RegisterUser harus unavailable
	existAfter, err := userService.IsEmailAvailable(user.CheckEmailInput{Email: newUser.Email})
	if assert.NoError(t, err) {
		assert.Equal(t, false, existAfter)
	}
}

func TestGetUserById(t *testing.T) {
	userRepo, userService := createUserRepoTest()

	newUser := user.RegisterUserInput{
		Nama:      "Test New User",
		Pekerjaan: "Mahasiswa",
		Email:     "example@test.com",
		Password:  "12345678",
	}

	// Email harus belum pernah dipakai
	exist, err := userService.IsEmailAvailable(user.CheckEmailInput{Email: newUser.Email})
	if assert.NoError(t, err) {
		assert.Equal(t, true, exist)
	}

	registered, err := userService.RegisterUser(newUser)
	assert.NoError(t, err)
	t.Cleanup(func() {
		// Hapus user yang digunakan saat testing setelah unit testing ini selesai
		// Agar ketika test selanjutnya berjalan, kondisi db tidak punya user testing ini
		userRepo.Delete(registered.ID)
	})

	// GetByUserId harus berhasil
	getUserByIdRes, err := userService.GetUserByID(registered.ID)
	if assert.NoError(t, err) {
		assert.Equal(t, getUserByIdRes.Nama, newUser.Nama)
		assert.Equal(t, getUserByIdRes.Pekerjaan, newUser.Pekerjaan)
		assert.Equal(t, getUserByIdRes.Email, newUser.Email)
		// Password harus bukan plaintext
		assert.NotEqual(t, getUserByIdRes.Password, newUser.Password)
	}
}

func TestRegisterUserFlow(t *testing.T) {
	userRepo, userService := createUserRepoTest()

	newUser := user.RegisterUserInput{
		Nama:      "Test New User",
		Pekerjaan: "Mahasiswa",
		Email:     "example@test.com",
		Password:  "12345678",
	}

	// Email harus belum pernah dipakai
	exist, err := userService.IsEmailAvailable(user.CheckEmailInput{Email: newUser.Email})
	assert.NoError(t, err)
	if assert.NoError(t, err) {
		assert.Equal(t, true, exist)
	}

	registered, err := userService.RegisterUser(newUser)
	t.Cleanup(func() {
		// Hapus user yang digunakan saat testing setelah unit testing ini selesai
		// Agar ketika test selanjutnya berjalan, kondisi db tidak punya user testing ini
		userRepo.Delete(registered.ID)
	})

	if assert.NoError(t, err) {
		assert.Equal(t, registered.Nama, newUser.Nama)
		assert.Equal(t, registered.Pekerjaan, newUser.Pekerjaan)
		assert.Equal(t, registered.Email, newUser.Email)
		// Password harus bukan plaintext
		assert.NotEqual(t, registered.Password, newUser.Password)
	}

	// Email setelah RegisterUser harus unavailable
	existAfter, err := userService.IsEmailAvailable(user.CheckEmailInput{Email: newUser.Email})
	if assert.NoError(t, err) {
		assert.Equal(t, false, existAfter)
	}

	// Setelah register harus berhasil login dengan email dan password saat mendaftar
	loginRes, err := userService.Login(user.LoginInput{
		Email:    newUser.Email,
		Password: newUser.Password,
	})
	if assert.NoError(t, err) {
		assert.Equal(t, loginRes.Nama, newUser.Nama)
		assert.Equal(t, loginRes.Pekerjaan, newUser.Pekerjaan)
		assert.Equal(t, loginRes.Email, newUser.Email)
		// Password harus bukan plaintext
		assert.NotEqual(t, loginRes.Password, newUser.Password)
	}

	// GetByUserId harus berhasil
	getUserByIdRes, err := userService.GetUserByID(registered.ID)
	if assert.NoError(t, err) {
		assert.Equal(t, getUserByIdRes.Nama, newUser.Nama)
		assert.Equal(t, getUserByIdRes.Pekerjaan, newUser.Pekerjaan)
		assert.Equal(t, getUserByIdRes.Email, newUser.Email)
		// Password harus bukan plaintext
		assert.NotEqual(t, getUserByIdRes.Password, newUser.Password)
	}
}
