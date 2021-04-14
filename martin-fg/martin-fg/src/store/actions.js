import {
  getUserInfo
} from '../api'

import {
  USER_INFO
} from './mutation-types'

export default {
  //获取用户信息
  async reqUserInfo({commit}){
    const result = await getUserInfo()
    commit(USER_INFO, {userInfo : result})
  }
}
