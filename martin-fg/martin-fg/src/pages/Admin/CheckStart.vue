<template>
  <div class="check-start">
    <div class="check-head">
      体检登记
    </div>
    <div class="check-field">
      <mt-field label="手机号" placeholder="请输入用户手机号" type="tel" v-model="phone"></mt-field>
    </div>
    <button @click="checkStartButton">提交</button>
  </div>
</template>

<script>
  import {MessageBox, Toast} from 'mint-ui'
  import {checkStart} from "../../api";

  export default {
    name: "CheckStart",
    data(){
      return {
        phone : ''
      }
    },
    methods : {
      async checkStartButton() {
        MessageBox.confirm("是否提交?").then(async action=>{
          let result = await checkStart(this.$store.state.userOpenId, this.phone)
          if (result.code !== 0) {
            Toast(result.msg)
            return
          }
          Toast("登记成功")
          setTimeout(()=>{
            this.$router.replace('/admin/top')
          }, 1000)
        })
      }
    }
  }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  .check-start
    width  100%
    height 100%
    .check-head
      width 100%
      height 50px
      line-height 50px
      font-weight bold
      text-align center
    .check-field
      margin-top 25px
    button
      width 90%
      height 40px
      line-height 40px
      background-color: #e9232c
      text-align center
      margin 100px 5%
      border none
      font-size 16px
      color #fff
      border-radius 10px
</style>
