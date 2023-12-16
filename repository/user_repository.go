package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}


// funcの宣言に(ur *userRepository)があるかないかで、
// その関数が「メソッド」として振る舞うか、単なる「関数」として振る舞うかが変わります。
// これはGo言語の「レシーバ」の概念に関連しています。
// この記述があることで、GetUserByEmailはuserRepository型のインスタンスに紐付けられたメソッドとして定義されます。
// メソッド内では、urを通じてレシーバのインスタンス（この場合はuserRepositoryのインスタンス）にアクセスできます。
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}