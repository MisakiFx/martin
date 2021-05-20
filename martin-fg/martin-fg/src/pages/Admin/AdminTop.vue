<template>
  <div class="admin-top">
    <div class="bottom">
      <h4>管理员界面</h4>
    </div>
    <router-link tag="button" to="/admin/check_start">体检登记</router-link>
    <router-link tag="button" to="/admin/check_finish">单项体检结束</router-link>
    <router-link tag="button" to="/admin/check_result">体检结果登记</router-link>
  </div>
</template>

<script>
    import {checkAdmin} from "../../api";
    import {Toast} from 'mint-ui';
    export default {
      name: "AdminTop",
      methods : {
        async check() {
          let result = await checkAdmin(this.$store.state.userOpenId)
          if (result.code !== 0) {
            Toast(result.msg)
            setTimeout(()=>{
              this.$router.replace('/home?code=' + this.$route.query.code)
            }, 1000)
          }
        }
      },
      mounted() {
        this.check()
      }
    }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  .admin-top
    background-color #f5f5f5
    width 100%
    height 100%
    position fixed
    left 0
    top 0
    z-index 1000
    display flex
    flex-direction column
    justify-content center
    align-items center
    button
      width 80%
      height 38px
      background-color purple
      border none
      border-radius 5px
      font-size 20px
      color #fff
      margin-top 30px
    .bottom
      position absolute
      top 20%
      font-size 30px
      font-weight bold
      color mediumpurple
      text-align center
</style>
