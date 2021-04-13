<template>
  <div class="bottom-tab">
    <span class="tab-item" @click="switchTo('/home')">
      <img :src="'/home' === $route.path ? tabBarImgArr[0].selected : tabBarImgArr[0].normal" alt="">
      <span :class="{on: '/home' === $route.path}">体检商城</span>
    </span>
    <span class="tab-item" @click="switchTo('/me')">
      <img :src="'/me' === $route.path ? tabBarImgArr[1].selected : tabBarImgArr[1].normal" alt="">
      <span :class="{on: '/me' === $route.path}">个人中心</span>
    </span>
  </div>
</template>

<script>
    export default {
        name: "TabBar",
        data(){
          return {
            tabBarImgArr: [
              {normal: require('../../../static/img/store.png'), selected: require('../../../static/img/store_click.png')},
              {normal: require('../../../static/img/me.png'), selected: require('../../../static/img/me_click.png')},
            ]
          }
        },
        methods:{
          switchTo(path){
            //console.log(this.$router);
            this.$router.replace(path);
          }
        }
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
        color black
</style>
