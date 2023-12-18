package middles

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Encode(password string) string {

	// 生成密码的哈希值
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("加密密码时发生错误:", err)
	}
	// 将哈希值转换为字符串并打印
	hashedPasswordString := string(hashedPassword)
	return hashedPasswordString

	// 验证密码
	//err = bcrypt.CompareHashAndPassword([]byte(hashedPasswordString), []byte(password))
	//if err != nil {
	//	fmt.Println("密码验证失败:", err)
	//	return
	//}
	//
	//fmt.Println("密码验证成功！")
	//mimi := "dfhyuikgfhUKGhkghkgighkkhjhjk"
	//mimei := "fandasgjldkshfjlSHLFjhsjflkj//'sd;lfk;l"
	//newpassword := mimi + password + mimei
	//return newpassword
}
func ValidatePassword(Oldpassword string, newpassword string) string {
	err := bcrypt.CompareHashAndPassword([]byte(Oldpassword), []byte(newpassword))
	if err != nil {
		return "失败"
	}
	return ""
}
