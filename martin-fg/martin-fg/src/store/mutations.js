import {
  USER_EXAMINATION_INFO,
  USER_INFO, USER_OPEN_ID
} from './mutation-types'

export default {
  [USER_INFO](state, {userInfo}) {
    state.userInfo = userInfo.data.info;
  },
  [USER_OPEN_ID](state, {userOpenId}) {
    state.userOpenId = userOpenId.data.open_id
  },
  [USER_EXAMINATION_INFO](state, {userExamination}) {
    state.userExamination = userExamination.data.info
  }
}
