import {
  USER_INFO
} from './mutation-types'

export default {
  [USER_INFO](state, {userInfo}) {
    state.userInfo = userInfo;
  }
}
