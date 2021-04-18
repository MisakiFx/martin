import ajax from './ajax'

const BASE_URL= 'http://localhost:8080/guardian/api'

//获取用户信息
export const getUserInfo = (openId = 0)=>ajax(BASE_URL + '/user/info', {}, {'open_id':openId})

//发送手机验证码
export const getPhoneCode = (phoneNumber = "")=>ajax(BASE_URL + '/user/login/verification_code', {'phone': phoneNumber}, {'open_id':123})

//获取用户openId
export const getUserOpenId = (code = "")=>ajax(BASE_URL + '/user/open_id/' + code)

//注册用户
export const postUserLogin = (userInfo = {})=>ajax(BASE_URL + '/user/login', userInfo, {}, 'POST')

//更改用户信息
export const postUpdateUserInfo = (userInfo = {})=>ajax(BASE_URL + '/user/update', userInfo, {'open_id':userInfo.open_id}, 'POST')

//购买体检卡
export const postBuyExamination = (examination = {}, userInfo = {})=>ajax(BASE_URL + '/examination/buy', examination, {'open_id' : userInfo.open_id}, 'POST')
