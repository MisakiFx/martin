<template>
  <div class="me-top">
    <router-link tag="div" class="user" to="/detail">
      <img src="../../../static/img/myIcon.png" alt="">
      <p v-if="userInfo.phone_number">{{userInfo.phone_number | phoneFormat}}</p>
      <i class="itlike-3"></i>
      <p>{{userInfo.user_name}}</p>
    </router-link>
    <div class="my-older">
      <div class="older-top">
        <h3>我的钱包</h3>
        <router-link tag="span" to="/expense">消费记录 > </router-link>
      </div>
      <div class="older-bottom">
        <div class="bottom-item">
          <span>{{userExamination.remainder}}元</span>
          <i></i>
          <span>余额</span>
        </div>
        <div class="bottom-item">
          <span>{{userExamination.card_type === 10 ? "无" : userExamination.card_type + "折"}}</span>
          <i></i>
          <span>权益卡</span>
        </div>
        <div class="bottom-item">
          <span>{{userExamination.check_count}}次</span>
          <i></i>
          <span>体检套餐</span>
        </div>
      </div>
    </div>
    <div class="setting">
      <router-link tag="div" to="/check_booking" class="setting-item">
        <i class="itlike-2"></i>
        <span>体检预约</span>
      </router-link>
      <router-link tag="div" to="/checks" class="setting-item">
        <i class="itlike-1"></i>
        <span>我的体检</span>
      </router-link>
      <router-link tag="div" to="/refund" class="setting-item">
        <i class="itlike-3"></i>
        <span>余额退款</span>
      </router-link>
    </div>
  </div>
</template>

<script>
  import {mapState} from 'vuex'
    export default {
      name: "MeTop",
      computed:{
          ...mapState(['userInfo', 'userExamination'])
      },
      mounted() {
        this.$store.dispatch('reqUserExamination', this.$store.state.userInfo.open_id)
      },
      filters:{
        phoneFormat(phone){
          // 1. 字符串转成数组
          let phoneArr = [...phone];
          // 2. 遍历
          let str = '';
          phoneArr.forEach((item, index)=>{
            if(
              index === 3 ||
              index === 4 ||
              index === 5 ||
              index === 6
            ){
              str += '*';
            }else {
              str += item;
            }
          });
          // 3. 返回结果
          return str;
        }
      },
    }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  .me-top
    width 100%
    height 100%
    background-color #f5f5f5
    font-size 14px
    .user
      display flex
      flex-direction row
      align-items center
      padding 20px
      background-color #fff
      margin-bottom 10px
      p
        margin 0 10px
      img
        width 60px
        height 60px
        border-radius 50%
      i
        font-size 25px
    .my-older
      background-color #fff
      .older-top
        display flex
        flex-direction row
        padding 0 10px
        justify-content space-between
        height 44px
        line-height 44px
      .older-bottom
        display flex
        .bottom-item
          flex 1
          height 60px
          display flex
          flex-direction column
          justify-content center
          align-items center
          i
            font-size 30px
            color #E9232C
            margin-bottom 5px
    .setting
      margin-top 10px
      background-color #fff
      display flex
      justify-content space-between
      flex-wrap wrap
      .setting-item
        width 90px
        height 70px
        display flex
        flex-direction column
        justify-content center
        align-items center
        i
          font-size 30px
          color orange
          margin-bottom 5px
</style>
