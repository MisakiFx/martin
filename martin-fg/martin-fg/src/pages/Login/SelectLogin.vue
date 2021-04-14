<template>
  <div class="select-login">
    <div class="bottom">
      <h4>选择注册方式</h4>
    </div>
    <router-link tag="button" to="/login">手机注册</router-link>
  </div>
</template>

<script>
    export default {
        name: "SelectLogin"
    }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  .select-login
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
