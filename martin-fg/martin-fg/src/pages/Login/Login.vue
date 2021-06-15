<template>
  <div class="login-container">
    <!--登录面板内容部分-->
    <div class="login-inner">
      <!--面板头部-->
      <div class="login-header">
        <div class="login-logo">
          用户注册
        </div>
        <!--面板标题-->
      </div>
      <!--面板表单部分-->
      <div class="login-content">
        <form>
          <!--手机验证码登录部分-->
          <div class="current">
            <section class="login-message">
              <input type="tel" maxlength="8" placeholder="姓名" v-model="user_name">
            </section>
            <section class="login-message">
              <input type="tel" maxlength="11" placeholder="手机号" v-model="phone">
              <button
                class="get-verification"
                v-if="!countDown"
                :class="{phone_right: phoneRight}"
                @click.prevent="getVerifyCode()"
              >
                获取验证码
              </button>
              <button
                v-else
                disabled="disabled"
                class="get-verification"
              >
                已发送({{countDown}}s)
              </button>
            </section>
            <section class="login-verification">
              <input type="tel" maxlength="6" placeholder="验证码" v-model="code">
            </section>
            <section class="login-hint">
              温馨提示：未注册体检中心的手机号，登录时将自动注册，且代表已同意
              <a href="javascript:;">服务协议与隐私政策</a>
            </section>
          </div>
          <button class="login-submit" @click.prevent="login()">注册</button>
        </form>
        <button class="login-back" @click="$router.back()">返回</button>
      </div>
    </div>
  </div>
</template>

<script>
    import {getPhoneCode, postUserLogin} from '../../api';
    import {Toast} from 'mint-ui';
    export default {
        name: "Login",
        data() {
          return {
            name:'',
            code:'',
            phone: '', // 手机号码
            countDown: 0, // 倒计时
            user_name: '', // 用户名
            userInfo: {}
          }
        },
        computed:{
          phoneRight() {
            return /^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$/.test(this.phone);
          }
        },
        methods:{
          async getVerifyCode() {
            // 2.1 开启倒计时
            if(this.phoneRight){
              this.countDown = 60;
              // 2.2 设置定时器
              this.intervalId = setInterval(()=>{
                this.countDown --;
                if(this.countDown === 0){
                  clearInterval(this.intervalId);
                }
              }, 1000);

              // 2.3 获取短信验证码
              const result = await getPhoneCode(this.phone);
              console.log(result);

              // 2.4 获取验证码失败
              if(result.code !== 0){
                // console.log(result.message);
                Toast({
                  message: result.msg,
                  position: 'center',
                  duration: 3000
                });

                // 2.5 后续处理
                //setTimeout(()=>{
                //  clearInterval(this.intervalId);
                //  this.countDown = 0;
                //}, 3000);
              }
            }
          },
          async login() {
            if(this.user_name === '') {
              Toast("请输入姓名")
              return;
            }
            if(!this.phone) {
              Toast("请输入手机号码")
              return;
            }
            if (!this.phoneRight) {
              Toast("请输入正确的手机号码")
              return;
            }
            if(!this.code) {
              Toast("请输入验证码")
              return;
            }
            if(!(/^\d{6}$/gi.test(this.code))) {
              Toast("请输入正确的验证码")
              return;
            }

            //调用接口
            const result = await postUserLogin({
              "open_id": this.$store.state.userOpenId,
              "user_name":this.user_name,
              "phone_number": this.phone,
              "user_gender" : 0,
              "verification_code" : this.code
            })

            console.log(result)
            if(result.code !== 0) {
              Toast(result.msg)
              return;
            }
            Toast("注册成功")
            //重新拉取用户信息
            await this.$store.dispatch('reqUserInfo', this.$store.state.userOpenId);
            if(!this.$store.state.userInfo.open_id) {
              Toast("获取用户信息失败:" + this.$store.state.userInfo.msg)
              return;
            }
            //跳转
            this.$router.back()
          }
        }
    }
</script>

<style scoped lang="stylus" ref="stylesheet/stylus">
  @import "../../common/stylus/mixins.styl"
  .login-container
    width 100%
    height 100%
    background #fff

    .login-inner
      padding-top 60px
      width 80%
      margin 0 auto

      .login-header
        .login-logo
          font-size 30px
          font-weight bold
          color mediumpurple
          text-align center

        .login-header-title
          padding-top 40px
          padding-bottom 10px
          text-align center

          > a
            color #333
            font-size 14px
            padding-bottom 4px

            &:first-child
              margin-right 40px

            &.current
              color mediumpurple
              font-weight 700
              border-bottom 2px solid mediumpurple

      .login-content
        > form
          > div
            display none

            &.current
              display block

            input
              width 100%
              height 100%
              padding-left 8px
              box-sizing border-box
              border 1px solid #ddd
              border-radius 4px
              outline 0
              font 400 14px Arial

              &:focus
                border 1px solid mediumpurple

            .login-message
              position relative
              margin-top 16px
              height 48px
              font-size 14px
              background #fff

              .get-verification
                position absolute
                top 50%
                right 10px
                transform translateY(-50%)
                border 0
                color #ccc
                font-size 14px
                background transparent

                &.phone_right
                  color purple

            .login-gender
              position absolute
              top 50%
              right 10px
              display: inline-block;
              color: #0D1529;


            .login-verification
              position relative
              margin-top 16px
              height 48px
              font-size 14px
              background #fff

              .switch-show
                position absolute
                right 10px
                top 12px

                img
                  display none

                img.on
                  display block

            .login-hint
              margin-top 12px
              color #999
              font-size 12px
              line-height 20px

              > a
                color mediumpurple

          .login-submit
            display block
            width 100%
            height 42px
            margin-top 30px
            border-radius 4px
            background mediumpurple
            color #fff
            text-align center
            font-size 16px
            line-height 42px
            border 0

        .login-back
          display block
          width 100%
          height 42px
          margin-top 15px
          border-radius 4px
          background transparent
          border: 1px solid mediumpurple
          color mediumpurple
          text-align center
          font-size 16px
          line-height 42px
</style>
