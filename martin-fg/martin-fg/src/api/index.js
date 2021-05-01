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

//获取用户体检卡信息
export const getUserExaminationInfo = (openId = 0)=>ajax(BASE_URL + '/examination/info', {}, {'open_id':openId})

//获取用户消费记录
export const getUserExpense = (openId = 0, page = 1, size = 10)=>ajax(BASE_URL + '/calendar/list', {'page' : page, 'size' : size}, {'open_id':openId})

//预约体检
export const postCheckBooking = (openId = 0, checkBooking = {})=>ajax(BASE_URL + '/check/booking', checkBooking, {'open_id' : openId}, 'POST')

//获取个人体检列表
export const getUserCheckList = (openId = 0, page = 1, size = 10)=>ajax(BASE_URL + '/check/list', {'page' : page, 'size' : size}, {'open_id' : openId})

//取消体检预约
export const postCancelCheck = (openId = 0, bookingId = 0)=>ajax(BASE_URL + '/check/cancel', {'booking_id' : bookingId}, {'open_id' : openId}, 'POST')

//查看体检结果
export const getCheckResult = (openId = 0, bookingId = 0)=>ajax(BASE_URL + '/check/result/' + bookingId, {}, {'open_id' : openId})
