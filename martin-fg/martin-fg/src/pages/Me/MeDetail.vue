<template>
    <div class="user-detail">
      <div class="user-detail-top">基本信息</div>
      <div class="user-detail-group">
        <div class="user-icon">
          <span>头像</span>
          <img src="../../../static/img/myIcon.png" alt="">
        </div>
        <div class="user-item">
          <span>手机</span>
          <span>181****3121</span>
        </div>
        <div class="user-item">
          <span>昵称</span>
          <span><input type="text" v-model="user_name"></span>
        </div>
        <div class="user-item" @click="sheetVisible = true">
          <span>性别</span>
          <span>{{user_sex}}</span>
        </div>
        <button @click="updateUserInfo">修改</button>
      </div>
      <mt-actionsheet
        :actions="action"
        v-model="sheetVisible">
      </mt-actionsheet>
    </div>
</template>

<script>
    import {Toast} from 'mint-ui';
    import {postUpdateUserInfo} from "../../api";

    export default {
      name: "MeDetail",
      data(){
        return {
          user_name : '',
          user_sex: '',

          sheetVisible : false,
          action:[
            {name:'男', method: this.selectSex},
            {name:'女', method: this.selectSex}
          ]
        }
      },
      mounted() {
        this.user_name = this.$store.state.userInfo.user_name
        if (this.$store.state.userInfo.user_gender === 1) {
          this.user_sex = '男'
        } else if (this.$store.state.userInfo.user_gender === 2){
          this.user_sex = '女'
        } else {
          this.user_sex = '未知'
        }
      },
      methods:{
        selectSex(props) {
          this.user_sex = props.name
        },
        async updateUserInfo() {
          let result = await postUpdateUserInfo({
            open_id: this.$store.state.userInfo.open_id,
            user_name: this.user_name,
            user_gender: (this.user_sex === '男' ? 1 : 2)
          })
          if(result.code !== 0) {
            Toast({
              message : result.msg,
              position: 'bottom',
              duration: 2000
            })
            return
          }

          Toast("修改成功")
          await this.$store.dispatch('reqUserInfo', this.$store.state.userInfo.open_id)
          setTimeout(()=>{
            this.$router.replace('/me')
          }, 1000)
        }
      }
    }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  .user-detail
    width  100%
    height 100%
  .user-detail-top
    width 100%
    height 60px
    line-height 60px
    padding -left 10px
    font-weight bold
  .user-detail-group
    .user-icon
      height 60px
      padding 0 10px
      background-color: #fff
      border-bottom 1px sold #e0e0e0
      display flex
      justify-content space-between
      align-items center
      img
        width 50px
        border-radius 50%
    .user-item
      height 40px
      padding 0 10px
      background-color: #fff
      border-bottom 1px solid #e0e0e0
      display flex
      justify-content space-between
      align-items center
      input
        border none
        outline none
        text-align right
    button
      width 90%
      height 40px
      line-height 40px
      background-color: #e9232c
      text-align center
      margin 20px 5%
      border none
      font-size 16px
      color #fff
      border-radius 10px
    .right-title-color
      color #999
      font-size 14px
</style>
