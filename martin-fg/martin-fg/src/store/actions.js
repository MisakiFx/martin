import {
  getUserExaminationInfo,
  getUserInfo, getUserOpenId
} from '../api'

import {
  USER_INFO,
  USER_OPEN_ID,
  USER_EXAMINATION_INFO
} from './mutation-types'

export default {
  //获取用户信息
  async reqUserInfo({commit}, openId){
    const result = await getUserInfo(openId)
    commit(USER_INFO, {userInfo : result})
  },

  //获取用户的openId用作微信鉴权
  async reqUserOpenId({commit}, code) {
    const result = await getUserOpenId(code)
    commit(USER_OPEN_ID, {userOpenId : result})
  },

  //获取用户的体检卡信息
  async reqUserExamination({commit}, openId) {
    const result = await getUserExaminationInfo(openId)
    commit(USER_EXAMINATION_INFO, {userExamination : result})
  }
}
