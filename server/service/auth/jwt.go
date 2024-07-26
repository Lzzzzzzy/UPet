package auth

import (
	"context"

	"github.com/Lzzzzzzy/UPet/server/global"
	"github.com/Lzzzzzzy/UPet/server/model/system/request"
	"github.com/Lzzzzzzy/UPet/server/model/user"
	"github.com/Lzzzzzzy/UPet/server/utils"
	"go.uber.org/zap"
)

type JwtService struct{}

// @function: GetRedisJWT
// @description: 从redis取jwt
// @param: userName string
// @return: redisJWT string, err error
func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.GVA_REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

// @function: SetRedisJWT
// @description: jwt存入redis并设置过期时间
// @param: jwt string, userName string
// @return: err error
func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.GVA_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

// TokenNext 登录以后签发jwt
func (jwtService *JwtService) CreateToken(user *user.User) (string, int64, error) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(request.BaseClaims{
		UserId:   user.ID,
		FamilyId: user.FamilyId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("创建token失败!", zap.Error(err))
		return token, 0, err
	}
	expire := claims.ExpiresAt.Unix() * 1000
	return token, expire, nil
}
