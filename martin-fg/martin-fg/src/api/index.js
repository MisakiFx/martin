import ajax from './ajax'

const BASE_URL= 'http://localhost:8080/guardian/api'

export const getUserInfo = ()=>ajax(BASE_URL + '/user/info', {}, {'open_id':123})

//发送手机验证码
export const getPhoneCode = (phoneNumber = "")=>ajax(BASE_URL + '/user/login/verification_code', {'phone': phoneNumber}, {'open_id':123})
