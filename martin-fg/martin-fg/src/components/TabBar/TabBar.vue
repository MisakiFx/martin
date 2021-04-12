<template>
  <div class="bottom-tab">
    <span class="tab-item">
      <img src="../../../static/img/store.png" alt="">
      <span>体检商城</span>
    </span>
    <span class="tab-item">
      <img src="../../../static/img/me.png" alt="">
      <span>个人中心</span>
    </span>
  </div>
</template>

<script>
    export default {
        name: "TabBar"
    }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  .bottom-tab
    width 100%
    height 50px
    background-color #fff
    position fixed
    left 0
    bottom 0
    z-index 999
    display flex
    .tab-item
      display flex
      flex 1
      flex-direction column
      justify-content center
      align-items center
      font-size 12px
      color #666
      img
        width 15%
        margin-bottom 2px
      .on
        color red
</style>
