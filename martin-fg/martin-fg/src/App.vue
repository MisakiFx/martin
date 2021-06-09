<template>
  <div id="app">
    <router-view></router-view>
    <tab-bar v-show="$route.meta.showBottomTabBar"></tab-bar>
  </div>
</template>

<script>
    import TabBar from './components/TabBar/TabBar'
    import {Toast} from 'mint-ui';
    export default {
      name: "App",
      components:{
        TabBar
      },
      methods:{
        async syncGetUserInfo(){
          await this.$store.dispatch('reqUserOpenId', this.$route.query.code)
          await this.$store.dispatch('reqUserInfo', this.$store.state.userOpenId)
          if(!this.$store.state.userInfo.open_id) {
            await this.$router.replace("/me")
          }
        }
      },
      mounted() {
        if(!this.$route.query.code) {
          Toast("无法识别用户身份，请从微信公众号打开")
          return;
        }
        this.syncGetUserInfo()
      }
    }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  #app
    width 100%
    height 100%
    background-color #f5f5f5
</style>
